#!/bin/sh

ROOTSRC=$(cd $(dirname $0) && pwd -P)

function make_gopath() {
	echo "$(go version)"
    export GOPATH=$ROOTSRC
}

function make() {
    go install -ldflags "-X 'main.BUILD_TIME=`date "+%Y-%m-%d %H:%M:%S"`'" algo
}

function build() {
	make_gopath

	make
}

function make_output() {
    rm -rf output

    mkdir -p output/bin
	mkdir -p output/conf
    mkdir -p output/log

}

cd $ROOTSRC
build
make_output
cd -
exit 0
