#!/usr/bin/env bash
basedir=$(
    cd $(dirname $0)
    pwd
)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
size=$((${BSC_CLUSTER_SIZE}))
standalone=false
bsc_node_ips=(${BSC_NODE_IP})
nodeurl=${BC_NODE_URL}

declare -A ips2ids
ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"


function setup_token_recover_contract() {
    
}

CMD=$1

case ${CMD} in
setup_token_recover_contract)
    echo "===== setup_token_recover_contract ===="
    setup_token_recover_contract
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_token_migration.sh setup_token_recover_contract"
    ;;
esac