#!/usr/bin/env bash
set -e
HUGO_DOCKER_IMAGE="klakegg/hugo:ext-alpine"
ASSET_DIR=${1:-src/query/generated/assets/openapi}
# First generate old versions
DOCS_VERSIONS=$(git tag -l 'docs/*')

# Now generate latest
# Copy over OpenAPI doc.
mkdir -p site/static/openapi
rsync -a --exclude=*.go "$ASSET_DIR"/* site/static/openapi
# Create .bak file and then delete it to make sed work for both GNU and Mac versions
sed -i.bak "s#spec-url='.*'#spec-url='spec.yml'#g" site/static/openapi/index.html
rm -f site/static/openapi/index.html.bak
# Now run hugo
if [[ -n "${HUGO_DOCKER:-}" ]]; then
        docker run --rm -it -v "$PWD"/site:/src "${HUGO_DOCKER_IMAGE}"
else
        cd site
        hugo -v
        cd ..
fi

for docVersion in $DOCS_VERSIONS
do
        IFS='/'
        read -ra version <<< "$docVersion"
        echo "Building $docVersion"
        git archive --output "site/${version[1]}.zip" "$docVersion:site/" && unzip -d "site/${version[1]}" "site/${version[1]}.zip" && rm -f "site/${version[1]}.zip"

        # Copy over OpenAPI doc.
        # mkdir -p site/static/openapi
        # rsync -a --exclude=*.go "$ASSET_DIR"/* site/static/openapi
        # # Create .bak file and then delete it to make sed work for both GNU and Mac versions
        # sed -i.bak "s#spec-url='.*'#spec-url='spec.yml'#g" site/static/openapi/index.html
        # rm -f site/static/openapi/index.html.bak
        # echo "public/${docVersion//docs\/}"
        # Now run hugo
        cd "site/${version[1]}"
        hugo -v -d "../public/${version[1]}"   
done