#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
source ${workspace}/ip2nodeids.sh
size=$((${BC_CLUSTER_SIZE}))
bc_node_ips=(${BC_NODE_IP})
p2p_port=8557

function start_server() {
        for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo bash /server/relayer/start_oracle_relayer && sudo bash /server/relayer/start_bsc_relayer && sudo bash /server/bsc/start_geth.sh && sudo bash /server/bc/start_node.sh"
    done
    sleep 10
}

function recover_bsc() { 
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo rm -rf /server/bsc && sudo tar -I lz4 -xvf /mnt/efs/bsc-qa/bsc${i}.tar && mv bsc /server/ "
    done
    sleep 10

}

function recover_bc() {
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo rm -rf /server/bc && sudo tar -I lz4 -xvf bc${i}.tar  && mv bc /server/"
    done
    sleep 10
}

function recover_relayer() {
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo rm -rf /server/relayer &&  sudo tar -I lz4 -xvf relayer${i}.tar  && mv relayer /server/relayer"
    done
    sleep 10
}

CMD=$1
case ${CMD} in
all)
    start_server
    recover_bc
    recover_bsc
    recover_relayer
    ;;
bc)
    recover_bc
    ;;
bsc)
    recover_bsc
    ;;
relayer)
    recover_relayer
    ;;
*)
    echo "Usage: recover_bc_bsc.sh all | bc | bsc | relayer"
    ;;
esac

