#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
source ${workspace}/ip2nodeids.sh
bc_node_ips=(${BC_NODE_IP})

function exit_previous() {
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        # stop client
        aws ssm send-command \
          --instance-ids "${dst_id}" \
          --document-name "AWS-RunShellScript" \
          --parameters commands="cp /mnt/efs/bsc-qa/bc-fusion/relayer/stop_oracle_relayer.sh /server/relayer/ && bash /server/relayer/stop_oracle_relayer.sh"
    done
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
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        cp ${workspace}/oracle_relayer.template ${workspace}/.local/relayer/oracle_relayer-${i}.json

        sed -i -e "s/{{bsc_chain_id}}/${BSC_CHAIN_ID}/g" ${workspace}/.local/relayer/oracle_relayer-${i}.json
        mnemonic="\"$(cat ${workspace}/.local/bc/node${i}/operator.info |tail  -1)\""
        sed -i -e "s/{{bbc_mnemonic}}/${mnemonic}/g" ${workspace}/.local/relayer/oracle_relayer-${i}.json
        sed -i -e "s:/data/relayer.db:/server/relayer/oracle_relayer.db:g" ${workspace}/.local/relayer/oracle_relayer-${i}.json
    done
}

function cluster_up() {
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 20000000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block
    done

    cp ${workspace}/tmp/oracle-relayer/build/oracle-relayer ${workspace}/.local/relayer/
    rm -rf /mnt/efs/bsc-qa/bc-fusion/relayer
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/relayer
    yes | cp -rf ${workspace}/.local/relayer/* /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/stop_bsc_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/start_bsc_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/stop_oracle_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/
    yes | cp -rf ${workspace}/start_oracle_relayer.sh /mnt/efs/bsc-qa/bc-fusion/relayer/

    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        
        aws ssm send-command \
          --instance-ids "${dst_id}" \
          --document-name "AWS-RunShellScript" \
          --parameters commands="cp /mnt/efs/bsc-qa/bc-fusion/relayer/start_oracle_relayer.sh /server/relayer/ && bash /server/relayer/start_oracle_relayer.sh ${i} reset"
    done
}

source ${workspace}/.env
CMD=$1

case ${CMD} in
init)
    echo "===== init ===="
    build_relayer
    init
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