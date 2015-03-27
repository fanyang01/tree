#!/bin/bash

for dir in $(ls); do
	if [[ ! -d $dir ]]; then
		continue
	fi
	cd $dir
	echo $dir
	if [[ "$1" != "" ]]; then
		mv ${dir}_test.go ${dir}_test.go.tmp
		sed "/^func Benchmark/a     b.N = $1" ${dir}_test.go.tmp > ${dir}_test.go
		go test -bench .
		mv ${dir}_test.go.tmp ${dir}_test.go
	else
		go test -bench .
	fi
	cd ..
done
