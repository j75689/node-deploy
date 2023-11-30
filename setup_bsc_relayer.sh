#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}

function exit_previous() {
	# stop client
    ps -ef | grep bsc-relayer  | grep config |awk '{print $2}' | xargs kill
}

function build_relayer() {
    rm -rf ${workspace}/tmp/bsc-relayer
    cd ${workspace}/tmp/
    git clone https://github.com/bnb-chain/bsc-relayer
    cd bsc-relayer
    make build
}

function init_config(){
    mkdir -p ${workspace}/.local/relayer/
    rm -rf ${workspace}/.local/relayer/bsc_relayer.*
    cp ${workspace}/bsc_relayer.template ${workspace}/.local/relayer/bsc_relayer.json
    sed -i -e "s/{{bsc_chain_id}}/${BSC_CHAIN_ID}/g" ${workspace}/.local/relayer/bsc_relayer.json
    sed -i -e "s/{{private_key}}/${INIT_HOLDER_PRV}/g" ${workspace}/.local/relayer/bsc_relayer.json
}


function prepare_native_config() {
    init_config
    LAN_IP=$(ifconfig |grep 192.168 |awk -F" " '{print $2}' |head -1)
    sed -i -e "s/bc-node-0.bc.svc.cluster.local/${LAN_IP}/g" ${workspace}/.local/relayer/bsc_relayer.json
    sed -i -e "s/bsc-node-0.bsc.svc.cluster.local/${LAN_IP}/g" ${workspace}/.local/relayer/bsc_relayer.json 
    sed -i -e "s:/data/relayer.db:${workspace}/.local/relayer/bsc_relayer.db:g" ${workspace}/.local/relayer/bsc_relayer.json 
}

source ${workspace}/.env
CMD=$1

case ${CMD} in
native_init)
    echo "===== init ===="
    build_relayer
    prepare_native_config
    echo "===== end ===="
    ;;
native_start)
    echo "===== stop native bsc-relayer===="
    exit_previous
    sleep 5
    echo "===== stop native bsc-relayer end ===="

    echo "===== start native node0 ===="
    cp ${workspace}/tmp/bsc-relayer/build/bsc-relayer ${workspace}/.local/relayer/
    nohup ${workspace}/.local/relayer/bsc-relayer --bbc-network-type 0 --config-type local --config-path ${workspace}/.local/relayer/bsc_relayer.json > ${workspace}/.local/relayer/bsc_relayer.log 2>&1 &
    echo "===== start native node0 end ===="
    ;;
native_stop)
    echo "===== stop native node0===="
    exit_previous
    sleep 5
    echo "===== stop native node0 end ===="
    ;;
*)
    echo "Usage: setup_bc_node.sh native_init | native_start | native_stop"
    ;;
esac