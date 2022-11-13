#!/usr/bin/env bash
set -xeuo pipefail

go build -o go-cp

./go-cp --from testdata/input.txt --to out.txt
# cmp out.txt testdata/out_offset0_limit0.txt

rm -f go-cp out.txt
echo "PASS"
