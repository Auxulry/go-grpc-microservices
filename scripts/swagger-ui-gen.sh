#!/bin/bash

# Change the SWAGGER_UI_VERSION to the version you want.
SWAGGER_UI_VERSION="v4.15.5"

# Make sure the
REMOTE_REPOSITORY="https://github.com/swagger-api/swagger-ui.git"
CACHE_DIR="./.cache/swagger-ui/$SWAGGER_UI_VERSION"
SWAGGER_DIR="./third_party/swagger-ui"
SWAGGER_FILES_GEN="./api"

set -e

[[ -z "$SWAGGER_UI_VERSION" ]] && echo "missing \$SWAGGER_UI_VERSION" && exit 1

escape_str() {
  echo "$1" | sed -e 's/[]\/$*.^[]/\\&/g'
}

# do caching if there's no cache yet
if [[ ! -d "$CACHE_DIR" ]]; then
  mkdir -p "$CACHE_DIR"
  tmp="$(mktemp -d)"
  echo "Pulling Swagger UI version: $SWAGGER_UI_VERSION"
  git clone --depth 1 --branch "$SWAGGER_UI_VERSION" "$REMOTE_REPOSITORY" "$tmp"
  echo "Success Pull Swagger UI version: $SWAGGER_UI_VERSION"
  cp -r "$tmp/dist/"* "$CACHE_DIR"
  cp -r "$tmp/LICENSE" "$CACHE_DIR"
  rm -rf "$tmp"
  echo "Swagger UI $SWAGGER_UI_VERSION ready for use"
fi

# Populate Swagger Files
tmp="    urls: ["
echo "Populating All Swagger Files."
for i in $(find "$SWAGGER_FILES_GEN" -name "*.swagger.json"); do
  escape_gen_dir="$(escape_str "$SWAGGER_FILES_GEN/")"
  path="$(echo $i | sed -e "s/$escape_gen_dir//g")"
  tmp="$tmp{\"url\":\"/static/$path\",\"name\":\"$path\"},"
done
echo "Success populate all swagger files"

# delete last characters from $tmp
tmp=$(echo "$tmp" | sed 's/.$//')
tmp="$tmp],"

# recreate swagger-ui, delete all except swagger.json
echo "Populate Swagger UI version: $SWAGGER_UI_VERSION"
find "$SWAGGER_DIR" -type f -delete
mkdir -p "$SWAGGER_DIR"
cp -r "$CACHE_DIR/"* "$SWAGGER_DIR"
echo "Success populate Swagger UI version: $SWAGGER_UI_VERSION"

# replace the default URL
echo "Set Urls Swagger UI"
line="$(cat "$SWAGGER_DIR/swagger-initializer.js" | grep -n "url" | cut -f1 -d:)"
escaped_tmp="$(escape_str "$tmp")"
sed -i'' -e "$line s/^.*$/$escaped_tmp/" "$SWAGGER_DIR/swagger-initializer.js"
rm -f "$SWAGGER_DIR/swagger-initializer.js-e"
echo "Finish set Urls Swagger UI"