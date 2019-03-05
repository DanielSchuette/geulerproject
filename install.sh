#!/bin/bash
# Run all files in `main/', return with
# exit code1 on error.
for file in main/*.go; do
    go run "$file" || exit 1;
done
