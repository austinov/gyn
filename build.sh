#!/bin/bash

set -e

BUILD_DIR=_build
UI_DIR=ui-web
BACKEND_DIR=backend
BACKEND_BIN=backend

pushd "$(dirname "$(readlink -f "$BASH_SOURCE[0]")")" > /dev/null && {

  rm -fr $BUILD_DIR
  mkdir -p $BUILD_DIR/$UI_DIR $BUILD_DIR/$BACKEND_DIR

  cd $UI_DIR

  # Remove previouse build
  rm -fr node_modules dist
  
  # Install npm dependencies
  npm install
  
  # Build ui-web
  npm run dist
  
  # Copy build folder
  cp -R dist/ html/ ../$BUILD_DIR/$UI_DIR/

  cd ../$BACKEND_DIR

  # Build backend 
  if [ -n "$1" ] && [ -n "$2" ]; then
      export GOOS=$1
      export GOARCH=$2
      if [ "$GOOS" = "windows" ]; then
          export CGO_ENABLED=0
          BACKEND_BIN="$BACKEND_BIN.exe"
      fi
      echo "Backend will build for $GOOS-$GOARCH with name \"$BACKEND_BIN\"."
  else
      echo "Backend will build for current OS with name \"$BACKEND_BIN\"."
  fi
  glide up
  go build -o ../$BUILD_DIR/$BACKEND_DIR/$BACKEND_BIN

  # Copy build folder
  cp -R docx/ ../$BUILD_DIR/$BACKEND_DIR/

  echo
  echo "Application was build into $BUILD_DIR directory."
  echo

  popd > /dev/null
}
