#!/bin/bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env

function prepare_bc_binary() {
    mkdir -p ${workspace}/tmp
    rm -rf ${workspace}/tmp/node
    cd ${workspace}/tmp
    git clone --branch ${BC_BRANCH} ${BC_REPO}

    cd node && make build
    cp -f ./build/bnbchaind ${workspace}/bin/bnbchaind
    cp -f ./build/tbnbcli ${workspace}/bin/tbnbcli
    cp -f ./build/bnbcli ${workspace}/bin/bnbcli
}

function prepare_bsc_binary() {
    mkdir -p ${workspace}/tmp
    rm -rf ${workspace}/tmp/bsc
    cd ${workspace}/tmp
    git clone --branch ${BSC_BRANCH} ${BSC_REPO}

    cd bsc && make geth
    go build -o ./build/bin/bootnode ./cmd/bootnode
    cp -f ./build/bin/geth ${workspace}/bin/geth
    cp -f ./build/bin/bootnode ${workspace}/bin/bootnode
}

function prepare_tool_binary() {
    cd ${workspace}
    go build -o ./bin/tool ./tool/main.go
}

CMD=$1
case ${CMD} in
bc)
    prepare_bc_binary
    ;;
bsc)
    prepare_bsc_binary
    ;;
tool)
    prepare_tool_binary
    ;;
all)
    prepare_bc_binary
    prepare_bsc_binary
    prepare_tool_binary
    ;;
*)
    echo "Usage: prepare.sh all | bc | bsc | tool"
    ;;
esac