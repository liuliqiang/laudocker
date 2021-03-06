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
	env GOOS=linux GOARCH=amd64 \
		go build -ldflags '-X main.Version=${full_version}' -o build/bin/${program} ${repo}/cmd/${program}/

release: cmd/${program}/main.go
	mkdir -p build/bin
	env GOOS=linux GOARCH=amd64 \
		go build -ldflags '-X main.Version=${version}' -o build/bin/${program} ${repo}/cmd/${program}/

clean:
	rm -f laudocker
	rm -f build/bin/laudocker

upload: build
	scp -P 2200 -i ${vagrant_key} build/bin/${program} vagrant@127.0.0.1:/vagrant/${program}

fix:
	mount -t proc proc /proc

install: release
	cp build/bin/${program} ${GOPATH}/bin/${program}

.PHONY: devel build clean fix upload
