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

# need a clean bc without stakings
function register_validator() {
    rm -rf ${workspace}/.local/bsc

    for ((i=0;i<${size};i++));do
        mkdir -p ${workspace}/.local/bsc/validator${i}
        echo "${KEYPASS}" > ${workspace}/.local/bsc/password.txt
        
        cons_addr=$(${workspace}/bin/geth account new --datadir ${workspace}/.local/bsc/validator${i} --password ${workspace}/.local/bsc/password.txt | grep "Public address of the key:" | awk -F"   " '{print $2}')
        fee_addr=$(${workspace}/bin/geth account new --datadir ${workspace}/.local/bsc/validator${i}_fee --password ${workspace}/.local/bsc/password.txt | grep "Public address of the key:" | awk -F"   " '{print $2}')
        mkdir -p ${workspace}/.local/bsc/bls${i}
        expect create_bls_key.exp ${workspace}/.local/bsc/bls${i} ${KEYPASS}
        vote_addr=0x$(cat ${workspace}/.local/bsc/bls${i}/bls/keystore/*json| jq .pubkey | sed 's/"//g')

        node_dir_index=${i}
        if [ $i -ge ${BC_CLUSTER_SIZE} ]; then
            # echo "${KEYPASS}" | ${workspace}/bin/tbnbcli keys delete node${i}-delegator --home ${workspace}/.local/bc/node0 # for re-entry
            echo "${KEYPASS}" | (echo "${KEYPASS}" | ${workspace}/bin/tbnbcli keys add node${i}-delegator --home ${workspace}/.local/bc/node0)
            node_dir_index=0
        fi
        delegator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${node_dir_index} | grep node${i}-delegator | awk -F" " '{print $3}')
        if [ "$i" != "0" ]; then
            sleep 6 #wait for including tx in block
            echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $delegator --amount ${BSC_INIT_DELEGATE_AMOUNT}:BNB --chain-id ${BC_CHAIN_ID} --node ${BC_NODE_URL} --home ${workspace}/.local/bc/node0
        fi
        sleep 6 #wait for including tx in block
        echo ${delegator} "balance"
        ${workspace}/bin/tbnbcli account ${delegator}  --chain-id ${BC_CHAIN_ID} --node ${BC_NODE_URL} --trust-node --home ${workspace}/.local/bc/node${node_dir_index} | jq .value.base.coins
        echo "${KEYPASS}" | ${workspace}/bin/tbnbcli staking bsc-create-validator \
            --side-cons-addr "${cons_addr}" \
            --side-vote-addr "${vote_addr}" \
            --bls-wallet ${workspace}/.local/bsc/bls${i}/bls/wallet \
            --bls-password "${KEYPASS}" \
            --side-fee-addr "${fee_addr}" \
            --address-delegator "${delegator}" \
            --side-chain-id ${BSC_CHAIN_NAME} \
            --amount ${BSC_INIT_DELEGATE_AMOUNT}:BNB \
            --commission-rate 80000000 \
            --commission-max-rate 95000000 \
            --commission-max-change-rate 3000000 \
            --moniker "${cons_addr}" \
            --details "${cons_addr}" \
            --identity "${delegator}" \
            --from node${i}-delegator \
            --chain-id "${BC_CHAIN_ID}" \
            --node ${BC_NODE_URL} \
            --home ${workspace}/.local/bc/node${node_dir_index}
    done
}

function prepare_config() {
    cd ${workspace}/genesis/
    git add . && git stash

    cd ${workspace}
    for ((i=0;i<${size};i++));do
        for f in ${workspace}/.local/bsc/validator${i}/keystore/*;do
            cons_addr="0x$(cat ${f} | jq -r .address)"
        done
        
        for f in ${workspace}/.local/bsc/validator${i}_fee/keystore/*;do
            fee_addr="0x$(cat ${f} | jq -r .address)"
        done
        
        mkdir -p ${workspace}/.local/bsc/clusterNetwork/node${i}
        bbcfee_addrs=${fee_addr}
        powers="0x000001d1a94a2000"
        if [ ${standalone} = false ]; then
            bbcfee_addrs=`${workspace}/bin/tbnbcli staking side-top-validators ${size} --side-chain-id=${BSC_CHAIN_NAME} --node="${nodeurl}" --chain-id=${BC_CHAIN_ID} --trust-node --output=json| jq -r ".[${i}].distribution_addr" |xargs ${workspace}/bin/tool -network-type 0 -addr`
            powers=`${workspace}/bin/tbnbcli staking side-top-validators ${size} --side-chain-id=${BSC_CHAIN_NAME} --node="${nodeurl}" --chain-id=${BC_CHAIN_ID} --trust-node --output=json| jq -r ".[${i}].tokens" |xargs ${workspace}/bin/tool -network-type 0 -power`
        fi
        mv ${workspace}/.local/bsc/bls${i}/bls ${workspace}/.local/bsc/clusterNetwork/node${i}/ && rm -rf ${workspace}/.local/bsc/bls${i}
        vote_addr=0x$(cat ${workspace}/.local/bsc/clusterNetwork/node${i}/bls/keystore/*json| jq .pubkey | sed 's/"//g')
        echo "${cons_addr},${bbcfee_addrs},${fee_addr},${powers},${vote_addr}" >> ${workspace}/genesis/validators.conf
        echo "validator" ${i} ":" ${cons_addr}
        echo "validatorFee" ${i} ":" ${fee_addr}
        echo "validatorVote" ${i} ":" ${vote_addr}
    done

    cd ${workspace}/genesis/
    npm install
    sed -i -e "s/address public constant WHITELIST_1 = 0xA904540818AC9c47f2321F97F1069B9d8746c6DB;/address public constant WHITELIST_1 = ${INIT_HOLDER};/g" ${workspace}/genesis/contracts/RelayerHub.template
    sed -i -e "s/dues = INIT_DUES;/dues = INIT_DUES;\n        whitelistInit();/g" ${workspace}/genesis/contracts/RelayerHub.template
    sed -i -e "s/function whitelistInit() external/function whitelistInit() public/g" ${workspace}/genesis/contracts/RelayerHub.template
    node generate-validator.js
    if [ ${standalone} = false ]; then
        initConsensusStateBytes=$(${workspace}/bin/tool -height 1 -rpc ${nodeurl} -network-type 0)
        node generate-genesis.js --chainid 714 --bscChainId 02ca --network 'local' --initConsensusStateBytes  ${initConsensusStateBytes}
    else
        node generate-genesis.js --chainid 714 --bscChainId 02ca --network 'local'
    fi

}
function generate_static_peers() {
    tool=${workspace}/bin/bootnode
    num=$1
    target=$2
    staticPeers=""
    for ((i=0;i<$num;i++)); do
        if [ $i -eq $target ]
        then
           continue
        fi

        file=${workspace}/.local/bsc/clusterNetwork/node${i}/geth/nodekey
        if [ ! -f "$file" ]; then
            $tool -genkey $file
        fi
        port=30311
        node_ip=${bsc_node_ips[${i}]}
        if [ ! -z "$staticPeers" ]
        then
            staticPeers+=","
        fi
        staticPeers+='"'"enode://$($tool -nodekey $file -writeaddress)@$node_ip:$port"'"'
    done

    echo $staticPeers
}

function initNetwork() {
    cd ${workspace}
    ${workspace}/bin/geth init-network --init.dir ${workspace}/.local/bsc/clusterNetwork --init.size=${size} --config ${workspace}/config.toml ${workspace}/genesis/genesis.json
    rm -rf  ${workspace}/*bsc.log*
    for ((i = 0; i < ${size}; i++)); do
        staticPeers=$(generate_static_peers ${size} ${i})
        line=`grep -n -e 'StaticNodes' ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml | cut -d : -f 1`
        head -n $((line-1)) ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml >> ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml-e
        echo "StaticNodes = [${staticPeers}]" >> ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml-e
        tail -n +$(($line+1)) ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml >> ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml-e
        rm -f ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml
        mv ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml-e ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml

        sed -i -e '/"<nil>"/d' ${workspace}/.local/bsc/clusterNetwork/node${i}/config.toml

        cp -r ${workspace}/.local/bsc/validator${i}/keystore ${workspace}/.local/bsc/clusterNetwork/node${i}
        cp ${workspace}/bin/geth ${workspace}/.local/bsc/clusterNetwork/node${i}/geth${i}
        # init genesis
        ${workspace}/.local/bsc/clusterNetwork/node${i}/geth${i} --datadir ${workspace}/.local/bsc/clusterNetwork/node${i} init ${workspace}/genesis/genesis.json
    done
}

function cluster_up() {
    declare -A ips2ids
    ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
    ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
    ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"

    rm -rf /mnt/efs/bsc-qa/bc-fusion/bsc_cluster
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/bsc_cluster
    cp -r ${workspace}/.local/bsc/* /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/
    cp -r ${workspace}/stop_geth.sh /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/
    cp -r ${workspace}/start_geth.sh /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/
    cp -f ${workspace}/bin/geth /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/

    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
    done
    sleep 10
    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo bash +x /server/bsc/start_geth.sh ${i} reset"
    done
}

function cluster_down() {
    declare -A ips2ids
    ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
    ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
    ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"

    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
    done
}

function cluster_restart() {
    declare -A ips2ids
    ips2ids["172.22.42.13"]="i-0d2b8632af953d0f6"
    ips2ids["172.22.42.94"]="i-001b988ca374e66f1"
    ips2ids["172.22.43.86"]="i-0d36ebf557138f8e5"

    yes | cp -r ${workspace}/stop_geth.sh /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/
    yes | cp -r ${workspace}/start_geth.sh /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/

    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
    done
    sleep 10
    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="sudo bash +x /server/bsc/start_geth.sh ${i}"
    done
}

function fyenman_hardfork(){
    rm -rf ${workspace}/tmp/bsc-feynman
    mkdir -p ${workspace}/tmp/bsc-feynman
    cd ${workspace}/tmp/bsc-feynman
    git clone --branch ${BSC_FYENMAN_BRANCH} ${BSC_FYENMAN_REPO}

    rm -rf ${workspace}/tmp/bsc_fyenman_bytecode
    mkdir -p ${workspace}/tmp/bsc_fyenman_bytecode
    cp -r ${workspace}/bsc_fyenman_bytecode/* ${workspace}/tmp/bsc_fyenman_bytecode/

    init_validator_addr=""
    init_vote_addr=""
    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        acc=$(${workspace}/bin/geth account list --datadir ${workspace}/.local/bsc/clusterNetwork/node${i} | grep 'Account #0' | awk -F'--' '{print $3}' | awk -F'/' '{print $NF}')
        sed -i -e "s/_validator${i}_/${acc}/g" ${workspace}/tmp/bsc_fyenman_bytecode/StakeHubContractByteCode
        bls=$(${workspace}/bin/geth bls account list --datadir ${workspace}/.local/bsc/clusterNetwork/node${i} --blspassword ${workspace}/.local/bsc/password.txt | grep -E -o '0x[0-9a-fA-F]+' | sed 's/0x//')
        sed -i -e "s/_vote${i}_/${bls}/g" ${workspace}/tmp/bsc_fyenman_bytecode/StakeHubContractByteCode
    done
    
    
    genesis_hash=$(curl -s ${BSC_NODE_URL} \
        -X POST \
        -H "Content-Type: application/json" \
        --data '{"method":"eth_getBlockByNumber","params":["0x0",false],"id":1,"jsonrpc":"2.0"}' | jq -r '.result.hash')
    echo "genesis_hash:" ${genesis_hash}
    current_block=$(curl -s ${BSC_NODE_URL} \
        -X POST \
        -H "Content-Type: application/json" \
        --data '{"method":"eth_blockNumber","params":[],"id":1,"jsonrpc":"2.0"}' | jq -r '.result')
    hertz_upgrade_block=$((${current_block} + ${BSC_WAIT_BLOCK_FOR_HERTZ}))

    upgrade_time=$1
    # Check if the input is empty
    if [ -z "$upgrade_time" ]; then
      echo "Please provide a timestamp as upgrade_time."
      exit 1
    fi
    echo "upgrade_time:" ${upgrade_time}

    sed -i -e "s/0xee835a629f9cf5510b48b6ba41d69e0ff7d6ef10f977166ef939db41f59f5501/${genesis_hash}/g" ${workspace}/tmp/bsc-feynman/bsc/params/config.go
    
    sed -i -e "s/_hertz_upgrade_block_/big.NewInt(${hertz_upgrade_block})/g" ${workspace}/tmp/bsc-feynman/bsc/params/config.go
    sed -i -e "s/_rialto_upgrade_height_/newUint64(${upgrade_time})/g" ${workspace}/tmp/bsc-feynman/bsc/params/config.go
    
    sed -i -e "s/ValidatorContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/ValidatorContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/SlashContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/SlashContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/TokenHubContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/TokenHubContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    
    sed -i -e "s/StakingContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/StakingContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/StakeHubContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/StakeHubContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/StakeCreditContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/StakeCreditContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/GovernorContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/GovernorContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/GovTokenContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/GovTokenContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/TimelockContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/TimelockContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/TokenRecoverPortalContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/TokenRecoverPortalContractByteCode)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go

    cd ${workspace}/tmp/bsc-feynman/bsc && make geth
    yes | cp -f ${workspace}/tmp/bsc-feynman/bsc/build/bin/geth ${workspace}/bin/geth_feynman
    yes | cp -f ${workspace}/tmp/bsc-feynman/bsc/build/bin/geth /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/

    echo $upgrade_time
}

function clean() {
    rm -rf ${workspace}/.local/bsc
}

function wait_for_hardfork(){
    # Input the target timestamp
    target_timestamp="$1"

    # Check if the input is empty
    if [ -z "$target_timestamp" ]; then
      echo "Please provide a timestamp as input."
      exit 1
    fi

    # Get the current timestamp
    current_timestamp=$(date +%s)

    # Calculate the waiting time in seconds
    wait_seconds=$((target_timestamp - current_timestamp))

    # Check if waiting is necessary
    if [ $wait_seconds -le 0 ]; then
      echo "The input timestamp has already passed."
      exit 1
    fi

    # Calculate the number of iterations
    iterations=$((wait_seconds / 5))

    # Run the for loop
    for ((i = 0; i < iterations; i++)); do
      echo "Waiting... (remaining seconds: $((target_timestamp - (current_timestamp + i * 5))))"
      sleep 5
    done

    echo "Time has elapsed!"
}

function migrate_validator() {
    validator_index=$1
    # Check if the input is empty
    if [ -z "$validator_index" ]; then
      echo "Please provide a validator index as input."
      exit 1
    fi

    validator_addr=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${validator_index} | grep node${validator_index} | awk '$1 == "node'${validator_index}'-delegator" {print $3}')
    validator_addr=$(echo ${validator_addr} | xargs ${workspace}/bin/tool -network-type 0 -bsc-val-addr)
    ${workspace}/bin/tbnbcli staking bsc-unbond \
     --chain-id ${BC_CHAIN_ID} \
     --side-chain-id ${BSC_CHAIN_NAME} \
     --from node${validator_index}-delegator \
     --validator ${validator_addr} \
     --amount ${BSC_INIT_DELEGATE_AMOUNT}:BNB \
     --home ${workspace}/.local/bc/node${validator_index}

}

CMD=$1
case ${CMD} in
cluster_up)
    echo "===== generate configs ===="
    clean
    register_validator
    prepare_config
    initNetwork
    cluster_up
    echo "===== end ===="
    ;;
cluster_down)
    echo "===== stop native bsc-relayer===="
    cluster_down
    sleep 5
    echo "===== stop native bsc-relayer end ===="
    ;;
cluster_restart)
    echo "===== cluster_restart ===="
    cluster_restart
    echo "===== end ===="
    ;;
fyenman_hardfork)
    echo "===== fyenman_hardfork ===="
    current_time=$(date +%s)
    hardfork_time=$((${current_time} + ${BSC_WAIT_SEC_FOR_FYENMAN}))
    fyenman_hardfork ${hardfork_time}
    cluster_restart
    wait_for_hardfork ${hardfork_time}
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_bsc_node.sh cluster_up | cluster_down | cluster_restart | fyenman_hardfork"
    ;;
esac
