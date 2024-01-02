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

dst_id="i-0f501565decf3b921"

function setup_token_recover_contract() {
    merkleRoot=$(cat /mnt/efs/bsc-qa/bc-fusion/dump_bc_account/output/base.json | jq -r '.state_root')
    procter=${TOKEN_RECOVERY_PROTECTOR}
    approver=${TOKEN_RECOVERY_APPROVER}
    operator="0x$(cat ${workspace}/.local/bsc/new_validator0_operator/keystore/* | jq -r '.address')"

    ${workspace}/bin/migrate_tool -bsc_endpoint ${BSC_NODE_URL} -priv_key ${INIT_HOLDER_PRV} \
      -token_recover_merkle_root ${merkleRoot} -token_recover_procter ${procter} -token_recover_approver ${approver} \
      -delegator_vote_operator_addr $operator
}

function start_approval_service() {
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/approval_service

    yes | cp -rf ${TOKEN_APPROVER_BIN} /mnt/efs/bsc-qa/bc-fusion/approval_service/approver
    yes | cp -rf ${workspace}/start_token_approver.sh /mnt/efs/bsc-qa/bc-fusion/approval_service/start_token_approver.sh
    yes | cp -rf ${workspace}/stop_token_approver.sh /mnt/efs/bsc-qa/bc-fusion/approval_service/stop_token_approver.sh

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="mkdir -p /server/token_approver/ && yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/approval_service/start_token_approver.sh /server/token_approver/start_token_approver.sh && bash /server/token_approver/start_token_approver.sh"
}

function stop_approval_service() {
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/approval_service

    yes | cp -rf ${TOKEN_APPROVER_BIN} /mnt/efs/bsc-qa/bc-fusion/approval_service/approver
    yes | cp -rf ${workspace}/start_token_approver.sh /mnt/efs/bsc-qa/bc-fusion/approval_service/start_token_approver.sh
    yes | cp -rf ${workspace}/stop_token_approver.sh /mnt/efs/bsc-qa/bc-fusion/approval_service/stop_token_approver.sh

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="mkdir -p /server/token_approver/ && yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/approval_service/stop_token_approver.sh /server/token_approver/stop_token_approver.sh && bash /server/token_approver/stop_token_approver.sh"
}

CMD=$1

case ${CMD} in
setup_token_recover_contract)
    echo "===== setup_token_recover_contract ===="
    setup_token_recover_contract
    echo "===== end ===="
    ;;
start)
    echo "===== start_approval_service ===="
    stop_approval_service
    sleep 5
    start_approval_service
    echo "===== end ===="
    ;;
stop)
    echo "===== stop_approval_service ===="
    stop_approval_service
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_token_migration.sh setup_token_recover_contract | start | stop"
    ;;
esac