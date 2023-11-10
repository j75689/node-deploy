#!/bin/bash
rm -rf ./clusterNetwork

bash +x ./setup_bsc_node.sh native_init

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

rm -rf /mnt/efs/bsc-qa/bc-fusion/clusterNetwork
cp -r /server/roshan/bc-fusion-test/clusterNetwork /mnt/efs/bsc-qa/bc-fusion
cp /server/roshan/bc-fusion-test/start_geth.sh /mnt/efs/bsc-qa/bc-fusion/clusterNetwork/.
cp /server/roshan/bc-fusion-test/stop_geth.sh /mnt/efs/bsc-qa/bc-fusion/clusterNetwork/.

for ((i = 0; i < ${#ips[@]}; i++)); do
    dst_id=${ips2ids[${ips[i]}]}
    aws ssm send-command \
        --instance-ids "${dst_id}" \
        --document-name "AWS-RunShellScript" \
        --parameters commands="sudo bash /server/clusterNetwork/stop_geth.sh"
    aws ssm send-command \
        --instance-ids "${dst_id}" \
        --document-name "AWS-RunShellScript" \
        --parameters commands="sudo rm -rf /server/clusterNetwork && sudo cp -r /mnt/efs/bsc-qa/bc-fusion/clusterNetwork /server && sudo chmod +x /server/clusterNetwork/node${i}/geth${i} && sudo bash +x /server/clusterNetwork/start_geth.sh ${i}"

done
