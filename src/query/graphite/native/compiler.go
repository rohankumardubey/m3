// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package native

import (
	goerrors "errors"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/m3db/m3/src/query/graphite/lexer"
	"github.com/m3db/m3/src/x/errors"
)

// CompileOptions allows for specifying compile options.
type CompileOptions struct {
	EscapeAllNotOnlyQuotes bool
}

// Compile converts an input stream into the corresponding Expression.
func Compile(input string, opts CompileOptions) (Expression, error) {
	compiler, closer := newCompiler(input, opts)
	defer closer()
	return compiler.compileExpression()
}

// getFirstPathExpression extracts the first matching path expression
// in the input.
func getFirstPathExpression(input string) (string, error) {
	compiler, closer := newCompiler(input, CompileOptions{})
	defer closer()

	// NB(r): Ignore any compilation errors, getFirstPathExpression is meant
	// to be able to compile partial expressions such as "foo.bar, 0.1)" which
	// might be matched as a result of a regular expression match with aliasSub.
	_, _ = compiler.compileExpression()
	if len(compiler.fetches) == 0 {
		return "", compiler.errorf("no fetch expressions")
	}

	return compiler.fetches[0].pathArg.path, nil
}

type closer func()

// A compiler converts an input string into an executable Expression
type compiler struct {
	input   string
	tokens  *tokenLookforward
	fetches []*fetchExpression
}

func newCompiler(input string, opts CompileOptions) (*compiler, closer) {
	booleanLiterals := map[string]lexer.TokenType{
		"true":  lexer.True,
		"false": lexer.False,
	}

	lex, tokens := lexer.NewLexer(input, booleanLiterals, lexer.Options{
		EscapeAllNotOnlyQuotes: opts.EscapeAllNotOnlyQuotes,
	})

	go lex.Run()

	cleanup := closer(func() {
		// Exhaust all tokens until closed or else lexer won't close.
		for range tokens {
		}
	})

	return &compiler{
		input:   input,
		tokens:  newTokenLookforward(tokens),
		fetches: make([]*fetchExpression, 0, 4),
	}, cleanup
}

type tokenLookforward struct {
	lookforward *lexer.Token
	tokens      chan *lexer.Token
}

func newTokenLookforward(tokens chan *lexer.Token) *tokenLookforward {
	return &tokenLookforward{
		tokens: tokens,
	}
}

// get advances the lexer tokens.
func (l *tokenLookforward) get() *lexer.Token {
	if token := l.lookforward; token != nil {
		l.lookforward = nil
		return token
	}

	if token, ok := <-l.tokens; ok {
		return token
	}

	return nil
}

func (l *tokenLookforward) peek() (*lexer.Token, bool) {
	if l.lookforward != nil {
		return l.lookforward, true
	}

	token, ok := <-l.tokens
	if !ok {
		return nil, false
	}

	l.lookforward = token
	return token, true
}

// compileExpression compiles a top level expression
func (c *compiler) compileExpression() (Expression, error) {
	token := c.tokens.get()
	if token == nil {
		return noopExpression{}, nil
	}

	var expr Expression
	switch token.TokenType() {
	case lexer.Pattern:
		expr = c.compileFetchExpression(token.Value())

	case lexer.Identifier:
		fc, err := c.compileFunctionCall(token.Value())
		fetchCandidate := false
		if err != nil {
			var (
				notFuncErr    *errNotFuncCall
				notFoundErr   *errFuncNotFound
				isNotFuncErr  = goerrors.As(err, &notFuncErr)
				isNotFoundErr = goerrors.As(err, &notFoundErr)
			)
			if (isNotFuncErr || isNotFoundErr) && c.canCompileAsFetch(token.Value()) {
				fetchCandidate = true
				expr = c.compileFetchExpression(token.Value())
			} else {
				return nil, err
			}
		}

		if !fetchCandidate {
			expr, err = newFuncExpression(fc)
			if err != nil {
				return nil, err
			}
		}

	default:
		return nil, c.errorf("unexpected value %s", token.Value())
	}

	if token := c.tokens.get(); token != nil {
		return nil, c.errorf("extra data %s", token.Value())
	}

	return expr, nil
}

func (c *compiler) compileFetchExpression(token string) *fetchExpression {
	expr := newFetchExpression(token)
	c.fetches = append(c.fetches, expr)
	return expr
}

// canCompileAsFetch attempts to see if the given term is a non-delimited
// carbon metric; no dots, without any trailing parentheses.
func (c *compiler) canCompileAsFetch(fname string) bool {
	if nextToken, hasNext := c.tokens.peek(); hasNext {
		return nextToken.TokenType() != lexer.LParenthesis
	}

	return true
}

type errFuncNotFound struct{ err error }

func (e *errFuncNotFound) Error() string { return e.err.Error() }

type errNotFuncCall struct{ err error }

func (e *errNotFuncCall) Error() string { return e.err.Error() }

// compileFunctionCall compiles a function call
func (c *compiler) compileFunctionCall(fname string) (*functionCall, error) {
	fn := findFunction(fname)
	if fn == nil {
		return nil, &errFuncNotFound{c.errorf("could not find function named %s", fname)}
	}

	if _, err := c.expectToken(lexer.LParenthesis); err != nil {
		// This could be just a pattern or fetch expression, so return
		// with context that there was no opening parenthesis.
		return nil, &errNotFuncCall{err}
	}

	argTypes := fn.in
	argTypesRequired := len(fn.in)
	if fn.variadic {
		// Variadic can avoid specifying the last arg.
		argTypesRequired--
	}
	var args []funcArg

	// build up arguments for function call
	for {
		// if not variadic, function should be complete after reading len(argTypes) arguments
		if !fn.variadic && len(args) == len(argTypes) {
			_, err := c.expectToken(lexer.RParenthesis)
			if err != nil {
				return nil, err
			}
			break
		}

		argType := argTypes[int(math.Min(float64(len(args)), float64(len(argTypes)-1)))]
		nextArg, foundRParen, err := c.compileArg(fn.name, len(args), argType)
		if err != nil {
			return nil, err
		}
		if foundRParen {
			break
		}

		args = append(args, nextArg)
	}

	// fill in defaults arguments for those not supplied by user explicitly
	for len(args) < len(argTypes) {
		defaultValue, ok := fn.defaults[uint8(len(args)+1)]
		if !ok {
			break
		}

		args = append(args, newConstArg(defaultValue))
	}

	// all required argument types should be filled with values now
	if len(args) < argTypesRequired {
		variadicComment := ""
		if fn.variadic {
			variadicComment = "at least "
		}
		return nil, c.errorf("invalid number of arguments for %s; expected %s%d, received %d",
			fn.name, variadicComment, len(argTypes), len(args))
	}

	return &functionCall{f: fn, in: args}, nil
}

// compileArg parses and compiles a single argument
func (c *compiler) compileArg(
	fname string,
	index int,
	reflectType reflect.Type,
) (arg funcArg, foundRParen bool, err error) {
	token := c.tokens.get()
	if token == nil {
		return nil, false, c.errorf("unexpected eof while parsing %s", fname)
	}

	if token.TokenType() == lexer.RParenthesis {
		return nil, true, nil
	}

	if index > 0 {
		if token.TokenType() != lexer.Comma {
			return nil, false, c.errorf("error parsing %s expected ',' received '%s'",
				fname, token.Value())
		}

		if token = c.tokens.get(); token == nil {
			return nil, false, c.errorf("unexpected eof while parsing %s", fname)
		}
	}

	arg, err = c.convertTokenToArg(token, reflectType)
	if err != nil {
		return nil, false, c.errorf("invalid function call %s, arg %d: %v", fname, index, err)
	}

	if !arg.CompatibleWith(reflectType) {
		return nil, false, c.errorf("invalid function call %s, arg %d: expected a %s, received a %s '%s'",
			fname, index, reflectTypeName(reflectType), reflectTypeName(arg.Type()), arg)
	}

	return arg, false, nil
}

// reflectTypeName will dereference any pointer types to their base name
// so that function call or fetch expression can be referenced by their name.
func reflectTypeName(reflectType reflect.Type) string {
	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType.Name()
}

// convertTokenToArg converts the given token into the corresponding argument
func (c *compiler) convertTokenToArg(token *lexer.Token, reflectType reflect.Type) (funcArg, error) {
	switch token.TokenType() {
	case lexer.Number:
		n, err := strconv.ParseFloat(token.Value(), 64)
		if err != nil {
			return nil, err
		}

		if reflectType.Kind() == reflect.Int {
			return newIntConst(int(n)), nil
		}

		return newFloat64Const(n), nil
	case lexer.String:
		return newStringConst(token.Value()), nil
	case lexer.Pattern:
		return c.compileFetchExpression(token.Value()), nil
	case lexer.True, lexer.False:
		b, err := strconv.ParseBool(token.Value())
		if err != nil {
			return nil, err
		}
		return newBoolConst(b), nil
	case lexer.Identifier:
		currentToken := token.Value()

		// handle named arguments
		nextToken, hasNextToken := c.tokens.peek()
		if !hasNextToken {
			return nil, c.errorf("unexpected eof, %s should be followed by = or (", currentToken)
		}

		if nextToken.TokenType() == lexer.Equal {
			// TODO: check if currentToken matches the expected parameter name
			_ = c.tokens.get() // Consume the peeked equal token.
			tokenAfterNext := c.tokens.get()
			if tokenAfterNext == nil {
				return nil, c.errorf("unexpected eof, named argument %s should be followed by its value", currentToken)
			}
			return c.convertTokenToArg(tokenAfterNext, reflectType)
		}

		fc, err := c.compileFunctionCall(currentToken)
		if err != nil {
			var (
				notFuncErr    *errNotFuncCall
				notFoundErr   *errFuncNotFound
				isNotFuncErr  = goerrors.As(err, &notFuncErr)
				isNotFoundErr = goerrors.As(err, &notFoundErr)
			)
			if (isNotFuncErr || isNotFoundErr) && c.canCompileAsFetch(currentToken) {
				return c.compileFetchExpression(currentToken), nil
			}
			return nil, err
		}

		return fc, nil
	default:
		return nil, c.errorf("%s not valid", token.Value())
	}
}

// expectToken reads the next token and confirms it is the expected type before returning it
func (c *compiler) expectToken(expectedType lexer.TokenType) (*lexer.Token, error) {
	token, ok := c.tokens.peek()
	if !ok || token == nil {
		return nil, c.errorf("expected %v but encountered eof", expectedType)
	}

	if token.TokenType() != expectedType {
		return nil, c.errorf("expected %v but encountered %s", expectedType, token.Value())
	}

	// Consume the token as it matches.
	_ = c.tokens.get()
	return token, nil
}

// errorf returns a formatted error vfrom the compiler
func (c *compiler) errorf(msg string, args ...interface{}) error {
	return errors.NewInvalidParamsError(fmt.Errorf("invalid expression '%s': %s", c.input, fmt.Sprintf(msg, args...)))
}

// ExtractFetchExpressions extracts timeseries fetch expressions from the given query
func ExtractFetchExpressions(s string) ([]string, error) {
	expr, err := Compile(s, CompileOptions{})
	if err != nil {
		return nil, err
	}

	var targets []string
	extractFetchExpressions(expr, &targets)
	return targets, nil
}

func extractFetchExpressions(expr Expression, targets *[]string) {
	switch v := expr.(type) {
	case *funcExpression:
		extractFetchExpressionsFromFuncCall(v.call, targets)
	case *fetchExpression:
		*targets = append(*targets, v.pathArg.path)
	}
}

func extractFetchExpressionsFromFuncCall(call *functionCall, targets *[]string) {
	for _, arg := range call.in {
		switch varg := arg.(type) {
		case *functionCall:
			extractFetchExpressionsFromFuncCall(varg, targets)
		case Expression:
			extractFetchExpressions(varg, targets)
		}
	}
}
