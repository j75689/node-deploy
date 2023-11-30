#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}

function exit_previous() {
	# stop client
    ps -ef  | grep oracle-relayer | grep config |awk '{print $2}' | xargs kill
}

function build_relayer() {
    rm -rf ${workspace}/tmp/oracle-relayer
    cd ${workspace}/tmp/
    git clone https://github.com/bnb-chain/oracle-relayer
    cd oracle-relayer
    make build
    cd build && mv relayer oracle-relayer
}

function init_config(){
    mkdir -p ${workspace}/.local/relayer/
    rm -rf ${workspace}/.local/relayer/oracle_relayer.*
    cp ${workspace}/oracle_relayer.template ${workspace}/.local/relayer/oracle_relayer.json

    sed -i -e "s/{{bsc_chain_id}}/${BSC_CHAIN_ID}/g" ${workspace}/.local/relayer/oracle_relayer.json
    mnemonic="\"$(cat ${workspace}/.local/bc/node0/operator.info |tail  -1)\""
    sed -i -e "s/{{bbc_mnemonic}}/${mnemonic}/g" ${workspace}/.local/relayer/oracle_relayer.json
}

function prepare_native_config() {
    init_config
    sed -i -e "s:/data/relayer.db:${workspace}/.local/relayer/oracle_relayer.db:g" ${workspace}/.local/relayer/oracle_relayer.json 
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
    echo "===== stop native oracle-relayer===="
    exit_previous
    sleep 5
    echo "===== stop native oracle-relayer end ===="

    echo "===== start native node0 ===="
    cp ${workspace}/tmp/oracle-relayer/build/oracle-relayer ${workspace}/.local/relayer/
    nohup ${workspace}/.local/relayer/oracle-relayer --bbc-network 0 --config-type local --config-path ${workspace}/.local/relayer/oracle_relayer.json > ${workspace}/.local/relayer/oracle_relayer.log 2>&1 &
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