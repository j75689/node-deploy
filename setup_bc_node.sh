#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
source ${workspace}/ip2nodeids.sh
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
        #sed -i -e "s/BEP159Height = 9223372036854775807/BEP159Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        #sed -i -e "s/BEP159Phase2Height = 9223372036854775807/BEP159Phase2Height = $((5+BC_BREATHE_BLOCK_INTERVAL))/g" ${workspace}/.local/bc/node${i}/config/app.toml
        #sed -i -e "s/LimitConsAddrUpdateIntervalHeight = 9223372036854775807/LimitConsAddrUpdateIntervalHeight = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP173Height = 9223372036854775807/BEP173Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml

        sed -i -e "s/FixDoubleSignChainIdHeight = 9223372036854775807/FixDoubleSignChainIdHeight = 4\nBEP171Height = 4/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP126Height = 9223372036854775807/BEP126Height = 5/g" ${workspace}/.local/bc/node${i}/config/app.toml
        sed -i -e "s/BEP255Height = 9223372036854775807/BEP255Height = 5/g" ${workspace}/.local/bc/node${i}/config/app.toml
    done

    # generate genesis.json
    ${workspace}/bin/bnbchaind collect-gentxs --acc-prefix tbnb --chain-id ${BC_CHAIN_ID} -i ${workspace}/.local/bc/genTx -o ${workspace}/.local/bc/genesis.json
    sed -i -e "s/\"min_self_delegation\": \"1000000000000\"/\"min_self_delegation\": \"100000000\"/g" ${workspace}/.local/bc/genesis.json
    sed -i -e "s/\"unbonding_time\": \"604800000000000\"/\"unbonding_time\": \"${BC_UNBONDING_TIME}\"/g" ${workspace}/.local/bc/genesis.json
    
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
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/stop_node.sh /server/bc/stop_node.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/start_node.sh /server/bc/start_node.sh && sudo bash /server/bc/stop_node.sh"
    done
}

function cluster_restart() {
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
    echo "enable mirror channel proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done
    sleep 120
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}

    # enable mirror sync channel
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli side-chain submit-channel-manage-proposal --channel-id 5 --enable=true --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "enable mirror channel" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable mirror sync channel proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done
    sleep 120
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}

    # enable mirror channel on bsc
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli params submit-cscParam-change-proposal --key "addOrUpdateChannel" --value "0x04000000000000000000000000000000000000001008" --target 0x0000000000000000000000000000000000002000 --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "enable mirror channel on bsc" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable mirror channel on bsc proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'-delegator" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i}-delegator --proposal-id ${proposal_id} --option Yes --side-chain-id ${BSC_CHAIN_NAME} --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
        sleep 6 #wait for including tx in block
    done

    # enable mirror sync channel on bsc
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli params submit-cscParam-change-proposal --key "addOrUpdateChannel" --value "0x05000000000000000000000000000000000000001008" --target 0x0000000000000000000000000000000000002000 --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "enable mirror sync channel on bsc" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable mirror channel on bsc proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'-delegator" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i}-delegator --proposal-id ${proposal_id} --option Yes --side-chain-id ${BSC_CHAIN_NAME} --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
        sleep 6 #wait for including tx in block
    done
}

function enable_cross_redelegation_channel() {
    # enable cross redelegation channel on bsc
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli params submit-cscParam-change-proposal --key "addOrUpdateChannel" --value "0x11000000000000000000000000000000000000002002" --target 0x0000000000000000000000000000000000002000 --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "enable cross redelegation channel on bsc" --from bsc-operator1 --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable cross redelegation channel on bsc proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 1; i < 7; i++)); do
        # vote
        echo "12345678" | ${workspace}/bin/tbnbcli gov vote --from bsc-operator${i} --proposal-id ${proposal_id} --option Yes --side-chain-id ${BSC_CHAIN_NAME} --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL}
        sleep 6 #wait for including tx in block
    done
}

function change_largeTransferLockPeriod() {
    # enable cross redelegation channel on bsc
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli params submit-cscParam-change-proposal --key "largeTransferLockPeriod" --value "0x0000000000000000000000000000000000000000000000000000000000000078" --target 0x0000000000000000000000000000000000001004 --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "change transfer lock period" --from bsc-operator1 --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable cross redelegation channel on bsc proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 1; i < 7; i++)); do
        # vote
        echo "12345678" | ${workspace}/bin/tbnbcli gov vote --from bsc-operator${i} --proposal-id ${proposal_id} --option Yes --side-chain-id ${BSC_CHAIN_NAME} --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL}
        sleep 6 #wait for including tx in block
    done
}

function disable_staking_channel() {
    # disable staking channel
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli side-chain submit-channel-manage-proposal --channel-id 8 --enable=false --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "disable staking channel" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    echo "disable staking channel proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done
    sleep 120
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}
}

function enable_staking_channel() {
    # disable staking channel
    proposal_id=$(echo "${KEYPASS}" | ${workspace}/bin/tbnbcli side-chain submit-channel-manage-proposal --channel-id 8 --enable=true --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "disable staking channel" --from node0-delegator --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --home ${workspace}/.local/bc/node0 --json=true | jq -r '.Response.data' | base64 -d)
    echo "enable staking channel proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        operator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${i} | grep node${i} | awk '$1 == "node'${i}'" {print $3}')
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $operator --amount 200000000:BNB --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        sleep 6 #wait for including tx in block

        # vote
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli gov vote --from node${i} --proposal-id ${proposal_id} --option Yes --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node${i}
    done
    sleep 120
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}
}

function change_sc_unbounding_time() {
    # change_sc_unbounding_time
    proposal_id=$(echo "12345678" | ${workspace}/bin/tbnbcli params submit-sc-change-proposal --sc-param-file ${workspace}/side-params.json --deposit 200000000000:BNB --voting-period 100 --side-chain-id ${BSC_CHAIN_NAME} --title "change_sc_unbounding_time" --from bsc-operator1 --node ${BC_NODE_URL} --trust-node --chain-id ${BC_CHAIN_ID} --json=true | jq -r '.Response.data' | base64 -d)
    echo "change sc unbounding time proposal_id: ${proposal_id}"
    sleep 6
    for ((i = 1; i < 7; i++)); do
        # vote
        echo "12345678" | ${workspace}/bin/tbnbcli gov vote --from bsc-operator${i} --proposal-id ${proposal_id} --option Yes --side-chain-id ${BSC_CHAIN_NAME} --chain-id ${BC_CHAIN_ID} --trust-node --node ${BC_NODE_URL}
        sleep 6 #wait for including tx in block
    done
}

function first_sunset_hardfork() {
    current_height=$(curl -sL ${BC_NODE_URL}/abci_info | jq -r '.result.response.last_block_height')

    expect_hardfork_height=$((${current_height} + ${WAIT_BLOCK_FOR_HARDFORK}))
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        sed -i -e "s/FirstSunsetHeight = 9223372036854775807/FirstSunsetHeight = ${expect_hardfork_height}/g" ${workspace}/.local/bc/node${i}/config/app.toml
        yes | cp -f ${workspace}/.local/bc/node${i}/config/app.toml /mnt/efs/bsc-qa/bc-fusion/bc_cluster/node${i}/config/app.toml
    done
    echo $expect_hardfork_height
}

function second_sunset_hardfork() {
    current_height=$(curl -sL ${BC_NODE_URL}/abci_info | jq -r '.result.response.last_block_height')

    expect_hardfork_height=$((${current_height} + ${WAIT_BLOCK_FOR_HARDFORK}))
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        sed -i -e "s/SecondSunsetHeight = 9223372036854775807/SecondSunsetHeight = ${expect_hardfork_height}/g" ${workspace}/.local/bc/node${i}/config/app.toml
        yes | cp -f ${workspace}/.local/bc/node${i}/config/app.toml /mnt/efs/bsc-qa/bc-fusion/bc_cluster/node${i}/config/app.toml
    done
    echo $expect_hardfork_height
}

function final_sunset_hardfork() {
    current_height=$(curl -sL ${BC_NODE_URL}/abci_info | jq -r '.result.response.last_block_height')

    expect_hardfork_height=$((${current_height} + ${WAIT_BLOCK_FOR_HARDFORK}))
    for ((i = 0; i < ${#bc_node_ips[@]}; i++)); do
        sed -i -e "s/FinalSunsetHeight = 9223372036854775807/FinalSunsetHeight = ${expect_hardfork_height}/g" ${workspace}/.local/bc/node${i}/config/app.toml
        yes | cp -f ${workspace}/.local/bc/node${i}/config/app.toml /mnt/efs/bsc-qa/bc-fusion/bc_cluster/node${i}/config/app.toml
    done

    echo $expect_hardfork_height
}

function wait_for_hardfork() {
    expect_hardfork_height=$1
    while true; do
        current_height=$(curl -sL ${BC_NODE_URL}/abci_info | jq -r '.result.response.last_block_height')
        echo "current_height: ${current_height}, expect_hardfork_height: ${expect_hardfork_height}"
        if [ ${current_height} -ge ${expect_hardfork_height} ]; then
            break
        fi
        sleep 10
    done
}

function get_channel_permission(){
    ${workspace}/bin/tbnbcli side-chain show-channel-permissions --node ${BC_NODE_URL} --trust-node --side-chain-id ${BSC_CHAIN_NAME}
}

CMD=$1
case ${CMD} in
cluster_up)
    echo "===== cluster_up ===="
    init
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
get_channel_permission)
    echo "===== get_channel_permission ===="
    get_channel_permission
    echo "===== end ===="
    ;;
enable_mirror_channel)
    echo "===== enable_mirror_channel ===="
    enable_mirror_channel
    echo "===== end ===="
    ;;
enable_staking_channel)
    echo "===== enable_staking_channel ===="
    enable_staking_channel
    echo "===== end ===="
    ;;
disable_staking_channel)
    echo "===== disable_staking_channel ===="
    disable_staking_channel
    echo "===== end ===="
    ;;
enable_cross_redelegation_channel)
    echo "===== enable_cross_redelegation_channel ===="
    enable_cross_redelegation_channel
    echo "===== end ===="
    ;;
change_sc_unbounding_time)
    echo "===== change_sc_unbounding_time ===="
    change_sc_unbounding_time
    echo "===== end ===="
    ;;
change_largeTransferLockPeriod)
    echo "===== change_largeTransferLockPeriod ===="
    change_largeTransferLockPeriod
    echo "===== end ===="
    ;;
first_sunset_hardfork)
    echo "===== first_sunset_hardfork ===="
    target=$(first_sunset_hardfork)
    cluster_restart
    sleep 10
    wait_for_hardfork ${target}
    echo "===== end ===="
    ;;
second_sunset_hardfork)
    echo "===== second_sunset_hardfork ===="
    target=$(second_sunset_hardfork)
    cluster_restart
    sleep 10
    wait_for_hardfork ${target}
    echo "===== end ===="
    ;;
final_sunset_hardfork)
    echo "===== final_sunset_hardfork ===="
    target=$(final_sunset_hardfork)
    cluster_restart
    sleep 10
    wait_for_hardfork ${target}
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_bc_node.sh cluster_up | cluster_down | cluster_restart | get_channel_permission | enable_mirror_channel | enable_cross_redelegation_channel | enable_staking_channel | disable_staking_channel | change_sc_unbounding_time | first_sunset_hardfork | second_sunset_hardfork | final_sunset_hardfork"
    ;;
esac