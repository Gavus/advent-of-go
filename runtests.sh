#!/usr/bin/env bash

go test $(find . -name "*_test.go" -exec dirname {} \;)
