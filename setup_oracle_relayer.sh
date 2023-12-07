#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
dst_id="i-0d2b8632af953d0f6"

function exit_previous() {
	# stop client
    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="cp /mnt/efs/bsc-qa/bc-fusion/relayer/stop_oracle_relayer.sh /server/relayer/ && bash /server/relayer/stop_oracle_relayer.sh"
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
    sed -i -e "s:/data/relayer.db:/server/relayer/oracle_relayer.db:g" ${workspace}/.local/relayer/oracle_relayer.json 
}

function cluster_up() {
    cp ${workspace}/tmp/oracle-relayer/build/oracle-relayer ${workspace}/.local/relayer/
    rm -rf /mnt/efs/bsc-qa/bc-fusion/relayer
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/relayer
    yes | cp -rf ${workspace}/.local/relayer/* /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/stop_oracle_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="rm -rf /server/relayer/oracle_relayer.db && mkdir -p /server/relayer/ && yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/relayer/* /server/relayer/ && nohup /server/relayer/oracle-relayer --bbc-network 0 --config-type local --config-path /server/relayer/oracle_relayer.json > /server/relayer/oracle_relayer.log 2>&1 &"
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
    echo "===== stop native oracle-relayer===="
    exit_previous
    sleep 5
    echo "===== stop native oracle-relayer end ===="

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
    echo "Usage: setup_oracle_relayer.sh init | start | stop"
    ;;
esac