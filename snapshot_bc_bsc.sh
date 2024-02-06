#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
source ${workspace}/ip2nodeids.sh
size=$((${BC_CLUSTER_SIZE}))
bc_node_ips=(${BC_NODE_IP})
p2p_port=8557

function stop_server() {
        for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo bash /server/relayer/stop_oracle_relayer && sudo bash /server/relayer/stop_bsc_relayer && sudo bash /server/bsc/stop_geth.sh && sudo bash /server/bc/stop_node.sh"
    done
    sleep 10
}

function backup_bsc() { 
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo tar -I lz4 -cvf bsc${i}.tar /server/bsc && mv bsc${i}.tar /mnt/efs/bsc-qa"
    done
    sleep 10

}

function backup_bc() {
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo tar -I lz4 -cvf bc${i}.tar /server/bc && mv bc${i}.tar /mnt/efs/bsc-qa"
    done
    sleep 10
}

function backup_relayer() {
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo tar -I lz4 -cvf relayer${i}.tar /server/relayer && mv relayer${i}.tar /mnt/efs/bsc-qa"
    done
    sleep 10
}

CMD=$1
case ${CMD} in
all)
    stop_server
    backup_bsc
    backup_bc
    backup_relayer
    ;;
backup_bc)
    backup_bc
    ;;
backup_bsc)
    backup_bsc
    ;;
backup_relayer)
    backup_relayer
    ;;
*)
    echo "Usage: snapshot_bc_bsc.sh all | bc | bsc | relayer"
    ;;
esac
