#!/bin/bash
basedir=$(
    cd $(dirname $0)
    pwd
)
workspace=${basedir}/..

declare -A ips2ids
ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"
ips="172.22.42.13 172.22.42.94 172.22.43.86"

for ((i = 0; i < ${#ips[@]}; i++)); do
    dst_id=${ips2ids[${ips[i]}]}
    aws ssm send-command \
        --instance-ids "${dst_id}" \
        --document-name "AWS-RunShellScript" \
        --parameters commands="sudo bash /server/clusterNetwork/stop_geth.sh"
done
