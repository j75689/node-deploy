#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}

dst_id="i-0d2b8632af953d0f6"

function exit_previous() {
    # stop client
    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="ps -ef | grep bsc-relayer  | grep config |awk '{print $2}' | xargs kill"    
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
    LAN_IP="0.0.0.0"
    sed -i -e "s/bc-node-0.bc.svc.cluster.local/${LAN_IP}/g" ${workspace}/.local/relayer/bsc_relayer.json
    sed -i -e "s/bsc-node-0.bsc.svc.cluster.local/${LAN_IP}/g" ${workspace}/.local/relayer/bsc_relayer.json 
    sed -i -e "s:/data/relayer.db:/server/relayer/bsc_relayer.db:g" ${workspace}/.local/relayer/bsc_relayer.json 
}

function cluster_up() {
    cp ${workspace}/tmp/bsc-relayer/build/bsc-relayer ${workspace}/.local/relayer/
    rm -rf /mnt/efs/bsc-qa/bc-fusion/relayer
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/relayer
    yes | cp -rf ${workspace}/.local/relayer/* /mnt/efs/bsc-qa/bc-fusion/relayer/

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="rm -rf /server/relayer/bsc_relayer.db && mkdir -p /server/relayer/ && yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/relayer/* /server/relayer/ && nohup /server/relayer/bsc-relayer --bbc-network-type 0 --config-type local --config-path /server/relayer/bsc_relayer.json > /server/relayer/bsc_relayer.log 2>&1 &"
}


source ${workspace}/.env
CMD=$1

case ${CMD} in
init)
    echo "===== init ===="
    build_relayer
    prepare_native_config
    echo "===== end ===="
    ;;
start)
    echo "===== stop bsc-relayer===="
    exit_previous
    sleep 5
    echo "===== stop bsc-relayer end ===="

    echo "===== start node0 ===="
    cluster_up
    echo "===== start node0 end ===="
    ;;
stop)
    echo "===== stop node0===="
    exit_previous
    sleep 5
    echo "===== stop node0 end ===="
    ;;
*)
    echo "Usage: setup_bc_node.sh init | start | stop"
    ;;
esac