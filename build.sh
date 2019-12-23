#!/bin/bash

tag="$(git describe --tags $(git rev-list --tags --max-count=1))"
filename="pi-temp-$tag-arm64-linux"

env GOOS=linux GOARCH=arm GOARM=6 go build -o release/$filename cmd/main.go