#!/usr/bin/env bash

printf "package deribit\n\n" > rpc_methods.go

shopt -s nullglob
for p in private public; do
    for f in examples/${p}/*.request.json; do
        name=$(basename "${f%.*}")
        go run cmd/gen/main.go --method=${name%.*} --namespace=${p} >> rpc_methods.go
    done
done