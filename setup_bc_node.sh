#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
size=$((${BC_CLUSTER_SIZE}))
bc_node_ips=(${BC_NODE_IP})
p2p_port=8557

function init() {
    rm -rf ${workspace}/.local/bc/*
    mkdir -p ${workspace}/.local/bc/genTx
    node_ids=""
    
    for ((i=0;i<${size};i++));do
        mkdir -p ${workspace}/.local/bc/node${i}

        # make node info
        ${workspace}/bin/bnbchaind init --home ${workspace}/.local/bc/node${i} --chain-id ${BC_CHAIN_ID} --moniker node${i} --kpass "${KEYPASS}" > ${workspace}/.local/bc/node${i}/node.info
        node_ip=${bc_node_ips[${i}]}
        node_ids="$(${workspace}/bin/bnbchaind tendermint show-node-id --home ${workspace}/.local/bc/node${i})@${node_ip}:${p2p_port} ${node_ids}"

        # create delegator and operator account
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli keys add node${i}-delegator --home ${workspace}/.local/bc/node${i} > ${workspace}/.local/bc/node${i}/delegator.info
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli keys add node${i} --home ${workspace}/.local/bc/node${i} > ${workspace}/.local/bc/node${i}/operator.info

        # create validator
        nodeID=$(cat ${workspace}/.local/bc/node${i}/node.info | jq -r '.node_id')
        pubKey=$(cat ${workspace}/.local/bc/node${i}/node.info | jq -r '.pub_key')
        delegator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i}-delegator | awk -F" " '{print $3}')
        ${workspace}/bin/tbnbcli staking create-validator --chain-id=${BC_CHAIN_ID} \
            --from node${i} --pubkey ${pubKey} --amount=1000000000:BNB \
            --moniker=node${i} --address-delegator=${delegator} --commission-rate=0 \
            --commission-max-rate=0 --commission-max-change-rate=0 --proposal-id=0 \
            --node-id=${nodeID} --genesis-format --home ${workspace}/.local/bc/node${i} \
            --generate-only > ${workspace}/.local/bc/node${i}/node${i}-unsigned.json 
        
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli sign \
            ${workspace}/.local/bc/node${i}/node${i}-unsigned.json \
            --name "node${i}-delegator" --home ${workspace}/.local/bc/node${i} \
            --chain-id=${BC_CHAIN_ID} --offline > ${workspace}/.local/bc/node${i}/node${i}-signed-t.json
        
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli sign \
            ${workspace}/.local/bc/node${i}/node${i}-signed-t.json \
            --name "node${i}" --home ${workspace}/.local/bc/node${i} \
            --chain-id=${BC_CHAIN_ID} --offline > ${workspace}/.local/bc/genTx/node${i}.json

        # modify configs
        sed -i -e "s/bscChainId = \"bsc\"/bscChainId = \"${BSC_CHAIN_NAME}\"/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/bscIbcChainId = 2/bscIbcChainId = ${BSC_CHAIN_ID}/g" ${workspace}/.local/bc/node${i}/config/app.toml

        sed -i -e "s/timeout_commit = \"1s\"/timeout_commit = \"${BC_BLOCK_TIMEOUT}\"/g" ${workspace}/.local/bc/node${i}/config/config.toml
    
        sed -i -e "s/breatheBlockInterval = 0/breatheBlockInterval = ${BC_BREATHE_BLOCK_INTERVAL}/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/bech32PrefixAccAddr = \"bnb\"/bech32PrefixAccAddr = \"tbnb\"/g" ${workspace}/.local/bc/node${i}/config/app.toml
        
        sed -i -e "s/BEP6Height = 1/BEP6Height = 2/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP9Height = 1/BEP9Height = 2/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP10Height = 1/BEP10Height = 2/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP19Height = 1/BEP19Height = 2/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP12Height = 1/BEP12Height = 2/g" ${workspace}/.local/bc/node${i}/config/app.toml

        sed -i -e "s/BEP3Height = 1/BEP3Height = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/FixSignBytesOverflowHeight = 1/FixSignBytesOverflowHeight = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/LotSizeUpgradeHeight = 1/LotSizeUpgradeHeight = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/ListingRuleUpgradeHeight = 1/ListingRuleUpgradeHeight = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/FixZeroBalanceHeight = 1/FixZeroBalanceHeight = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/LaunchBscUpgradeHeight = 1/LaunchBscUpgradeHeight = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP8Height = 1/BEP8Height = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP67Height = 1/BEP67Height = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP70Height = 1/BEP70Height = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP67Height = 1/BEP67Height = 3/g" ${workspace}/.local/bc/node${i}/config/app.toml
        
        sed -i -e "s/BEP82Height = 9223372036854775807/BEP82Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP84Height = 9223372036854775807/BEP84Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP87Height = 9223372036854775807/BEP87Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/FixFailAckPackageHeight = 9223372036854775807/FixFailAckPackageHeight = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/EnableAccountScriptsForCrossChainTransferHeight = 9223372036854775807/EnableAccountScriptsForCrossChainTransferHeight = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP128Height = 9223372036854775807/BEP128Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP151Height = 9223372036854775807/BEP151Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP153Height = 9223372036854775807/BEP153Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP173Height = 9223372036854775807/BEP173Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml

        sed -i -e "s/FixDoubleSignChainIdHeight = 9223372036854775807/FixDoubleSignChainIdHeight = 4\nBEP171Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP126Height = 9223372036854775807/BEP126Height = 5/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP255Height = 9223372036854775807/BEP255Height = 5/g" ${workspace}/.local/bc/node${i}/config/app.toml
    done

    # generate genesis.json
    ${workspace}/bin/bnbchaind collect-gentxs --acc-prefix tbnb --chain-id ${BC_CHAIN_ID} -i ${workspace}/.local/bc/genTx -o ${workspace}/.local/bc/genesis.json
    sed -i -e "s/\"min_self_delegation\": \"1000000000000\"/\"min_self_delegation\": \"100000000\"/g" ${workspace}/.local/bc/genesis.json
    
    # copy genesis file and set persistent peers
    persistent_peers=$(joinByString ',' ${node_ids})
    for ((i=0;i<${size};i++));do
        rm -rf ${workspace}/.local/bc/node${i}/config/genesis.json
        cp ${workspace}/.local/bc/genesis.json ${workspace}/.local/bc/node${i}/config/genesis.json

        sed -i -e "s/persistent_peers = \".*\"/persistent_peers = \"${persistent_peers}\"/g" ${workspace}/.local/bc/node${i}/config/config.toml
        sed -i -e "s/0.0.0.0:26656/0.0.0.0:${p2p_port}/g" ${workspace}/.local/bc/node${i}/config/config.toml
    done
}

function start_cluster() {
    declare -A ips2ids
    ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
    ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
    ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"

    rm -rf /mnt/efs/bsc-qa/bc-fusion/bc_cluster
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/bc_cluster
    cp -r ${workspace}/.local/bc/* /mnt/efs/bsc-qa/bc-fusion/bc_cluster/
    cp -r ${workspace}/stop_node.sh /mnt/efs/bsc-qa/bc-fusion/bc_cluster/
    cp -r ${workspace}/start_node.sh /mnt/efs/bsc-qa/bc-fusion/bc_cluster/
    cp -f ${workspace}/bin/bnbchaind /mnt/efs/bsc-qa/bc-fusion/bc_cluster/

    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/stop_node.sh /server/bc/stop_node.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/start_node.sh /server/bc/start_node.sh && sudo bash /server/bc/stop_node.sh"
    done
    sleep 10
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo bash +x /server/bc/start_node.sh ${i} reset"
    done
}

function cluster_down() {
    declare -A ips2ids
    ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
    ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
    ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"

    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/stop_node.sh /server/bc/stop_node.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/start_node.sh /server/bc/start_node.sh && sudo bash /server/bc/stop_node.sh"
    done
}

function cluster_restart() {
    declare -A ips2ids
    ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
    ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
    ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"

    yes | cp -r ${workspace}/stop_node.sh /mnt/efs/bsc-qa/bc-fusion/bc_cluster/
    yes | cp -r ${workspace}/start_node.sh /mnt/efs/bsc-qa/bc-fusion/bc_cluster/
    yes | cp -f ${workspace}/bin/bnbchaind /mnt/efs/bsc-qa/bc-fusion/bc_cluster/

    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/stop_node.sh /server/bc/stop_node.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/start_node.sh /server/bc/start_node.sh && sudo bash /server/bc/stop_node.sh"
    done
    sleep 10
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo bash +x /server/bc/start_node.sh ${i}"
    done
}

function enable_mirror_channel() {
    # enable mirror channel
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli side-chain submit-channel-manage-proposal --channel-id 4 --enable=true --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "enable mirror channel" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    sleep 6
    echo "enable mirror channel proposal_id: ${proposal_id}"
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --side-chain-id ${BSC_CHAIN_NAME} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done
    sleep 120
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}

    # enable mirror sync channel
    sleep 6
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli side-chain submit-channel-manage-proposal --channel-id 5 --enable=true --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "enable mirror channel" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable mirror sync channel proposal_id: ${proposal_id}"
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --side-chain-id ${BSC_CHAIN_NAME} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done
    sleep 120
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}
}

CMD=$1
case ${CMD} in
init)
    echo "===== init ===="
    init
    echo "===== end ===="
    ;;
cluster_up)
    echo "===== cluster_up ===="
    start_cluster
    echo "===== end ===="
    ;;
cluster_down)
    echo "===== cluster_down ===="
    cluster_down
    echo "===== end ===="
    ;;
cluster_restart)
    echo "===== cluster_restart ===="
    cluster_restart
    echo "===== end ===="
    ;;
enable_mirror_channel)
    echo "===== enable_mirror_channel ===="
    enable_mirror_channel
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_bc_node.sh init | cluster_up | cluster_down | cluster_restart | enable_mirror_channel"
    ;;
esac