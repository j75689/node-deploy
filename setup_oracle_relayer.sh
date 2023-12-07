#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
bc_node_ips=(${BC_NODE_IP})
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
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 20000000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done

    cp ${workspace}/tmp/oracle-relayer/build/oracle-relayer ${workspace}/.local/relayer/
    rm -rf /mnt/efs/bsc-qa/bc-fusion/relayer
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/relayer
    yes | cp -rf ${workspace}/.local/relayer/* /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/stop_oracle_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/start_oracle_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="cp /mnt/efs/bsc-qa/bc-fusion/relayer/start_oracle_relayer.sh /server/relayer/ && bash /server/relayer/start_oracle_relayer.sh reset"
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