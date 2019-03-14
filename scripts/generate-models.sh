#!/usr/bin/env bash

shopt -s nullglob
for p in private public; do
    for f in examples/${p}/*.request.json; do
        name=$(basename "${f%.*}")
        struct=$(echo "${name%.*}_request" | perl -pe 's/(^|_)./uc($&)/ge;s/_//g')
        cat ${f} | jq .params | gojson -name="${struct}" -pkg=${p} -o=models/${p}/${name%.*}_request.go
        if [[ $? -ne 0 ]]; then
            printf "package ${p}\n\ntype ${struct} string\n" > models/${p}/${name%.*}_request.go
        fi
    done
    for f in examples/${p}/*.response.json; do
        name=$(basename "${f%.*}")
        struct=$(echo "${name%.*}_response" | perl -pe 's/(^|_)./uc($&)/ge;s/_//g')
        cat ${f} | jq .result | gojson -name="${struct}" -pkg=${p} -o=models/${p}/${name%.*}_response.go
        if [[ $? -ne 0 ]]; then
            printf "package ${p}\n\ntype ${struct} string\n" > models/${p}/${name%.*}_response.go
        fi
    done
done
for f in examples/notifications/*.json; do
    name=$(basename "${f%.*}")
    name=${name//./_}
    struct=$(echo "${name}" | perl -pe 's/(^|_|\.)./uc($&)/ge;s/_|\.//g')
    cat ${f} | jq .params.data[0] | gojson -name="${struct}" -pkg=notifications -o=models/notifications/${name}.go
done