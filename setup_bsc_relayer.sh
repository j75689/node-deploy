#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}

dst_id="i-0870976185c85a30d"

function exit_previous() {
    # stop client
    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="cp /mnt/efs/bsc-qa/benchmark/relayer/stop_bsc_relayer.sh /server/relayer/ && bash /server/relayer/stop_bsc_relayer.sh"    
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
    rm -rf /mnt/efs/bsc-qa/benchmark/relayer
    mkdir -p /mnt/efs/bsc-qa/benchmark/relayer
    yes | cp -rf ${workspace}/.local/relayer/* /mnt/efs/bsc-qa/benchmark/relayer/
    yes | cp -rf ${workspace}/stop_bsc_relayer.sh /mnt/efs/bsc-qa/benchmark/relayer/
    yes | cp -rf ${workspace}/start_bsc_relayer.sh /mnt/efs/bsc-qa/benchmark/relayer/
    yes | cp -rf ${workspace}/stop_oracle_relayer.sh /mnt/efs/bsc-qa/benchmark/relayer/
    yes | cp -rf ${workspace}/start_oracle_relayer.sh /mnt/efs/bsc-qa/benchmark/relayer/

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="cp /mnt/efs/bsc-qa/benchmark/relayer/start_bsc_relayer.sh /server/relayer/ && bash /server/relayer/start_bsc_relayer.sh reset"
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
    echo "Usage: setup_bsc_relayer.sh init | start | stop"
    ;;
esac
