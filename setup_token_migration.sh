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

dst_id="i-001b988ca374e66f1"

function setup_token_recover_contract() {
    merkleRoot=$(cat /mnt/efs/bsc-qa/bc-fusion/dump_bc_account/output/base.json | jq -r '.state_root')
    procter=${TOKEN_RECOVERY_PROTECTOR}
    approver=${TOKEN_RECOVERY_APPROVER}
    operator="0x$(cat ${workspace}/.local/bsc/new_validator0_operator/keystore/* | jq -r '.address')"

    ${workspace}/bin/migrate_tool -bsc_endpoint ${BSC_NODE_URL} -priv_key ${INIT_HOLDER_PRV} \
      -token_recover_merkle_root ${merkleRoot} -token_recover_procter ${procter} -token_recover_approver ${approver} \
      -delegator_vote_operator_addr $operator
}

function setup_approval_service() {
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/approval_service

    yes | cp -rf ${TOKEN_APPROVER_BIN} /mnt/efs/bsc-qa/bc-fusion/approval_service/approver
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