#!/bin/bash
# Run all files in `main/', return with
# exit code1 on error.
for file in main/*.go; do
    go run "$file" || exit 1;
done

# determine test coverage with --cover
# (meant for local usage)
if [[ $1 == "--cover" ]]; then
    go test . -v -cover -race -coverprofile=coverage.txt -covermode=atomic
    go tool cover -html=coverage.txt
fi
