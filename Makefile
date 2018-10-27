cur_dir=$(shell pwd)

full_version=$(shell git describe --long --tags --dirty | awk '{print substr($$1,2)}')
version=$(shell echo ${full_version} | sed -nr "s/^([0-9]+(\.[0-9]+)+)(-([-A-Za-z0-9\.]+))?$$/\1/p")
release=$(shell echo ${full_version} | sed -nr "s/^([0-9]+(\.[0-9]+)+)(-([-A-Za-z0-9\.]+))?$$/\4/p")

program=laudocker
repo=github.com/liuliqiang/${program}
build_name=${program}-${full_version}.linux-amd64


devel: build
	mv build/bin/${program} ${cur_dir}/

build: cmd/${program}/main.go
	mkdir -p build/bin
	go build -o build/bin/${program} ${repo}/cmd/${program}/

clean:
	rm -f laudocker
	rm -f builf/bin/laudocker

fix:
	mount -t proc proc /proc

.PHONY: devel build clean fix
