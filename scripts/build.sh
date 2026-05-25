#!/bin/bash

set -eu

if [ ! -d "bin" ]; then
  mkdir bin
fi

if [ ! -d "cmd" ]; then
  echo "cmd directory is missing"
  exit 1
else
  cd cmd
  files=( ./*.go )

  filesCount=${#files[@]}

  for file in "${files[@]}"; do
    filename=$(basename "$file" .go)

    go build -o "../bin/$filename" "$file"
    echo "file $filename created in bin directory"
  done

  echo "total $filesCount files created in bin directory"
fi