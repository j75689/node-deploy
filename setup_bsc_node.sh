#!/usr/bin/env bash
basedir=$(
    cd $(dirname $0)
    pwd
)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
source ${workspace}/ip2nodeids.sh
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
        ${workspace}/bin/geth bls account new --datadir ${workspace}/.local/bsc/bls${i} --blspassword ${workspace}/.local/bsc/password.txt
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
    forge install --no-git --no-commit foundry-rs/forge-std@v1.1.1
    sed -i -e "s/0xA904540818AC9c47f2321F97F1069B9d8746c6DB/${INIT_HOLDER}/g" ${workspace}/genesis/scripts/generate-relayerHub.sh
    rm -rf ${workspace}/genesis/scripts/init_holders.js
    yes | cp -f ${workspace}/init_holder.js ${workspace}/genesis/scripts/init_holders.js
    node ${workspace}/genesis/scripts/generate-validator.js
    
    # tmp fix for bsc-genesis
    mv ${workspace}/genesis/scripts/generate-crossChain.sh ${workspace}/genesis/scripts/generate-crosschain.sh
    sed -i -e "s/176405560900000000000000000/17640556090000000000000000000/g" ${workspace}/genesis/scripts/generate-genesis.js

    initConsensusStateBytes=$(${workspace}/bin/tool -height 1 -rpc ${nodeurl} -network-type 0)
    sed -i -e "s/42696e616e63652d436861696e2d4e696c650000000000000000000000000000000000000000000229eca254b3859bffefaf85f4c95da9fbd26527766b784272789c30ec56b380b6eb96442aaab207bc59978ba3dd477690f5c5872334fc39e627723daa97e441e88ba4515150ec3182bc82593df36f8abb25a619187fcfab7e552b94e64ed2deed000000e8d4a51000/${initConsensusStateBytes}/g" ${workspace}/genesis/scripts/generate.sh
    bash scripts/generate.sh local
    sed -i -e "s/\"period\": 3/\"period\": ${BSC_BLCOK_INTERVAL}/g" ${workspace}/genesis/genesis.json
    sed -i -e "s/\"epoch\": 200/\"epoch\": ${BSC_EPOCH}/g" ${workspace}/genesis/genesis.json
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
    rm -rf /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster
    mkdir -p /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster
    cp -r ${workspace}/.local/bsc/* /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/
    cp -r ${workspace}/stop_geth.sh /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/
    cp -r ${workspace}/start_geth.sh /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/
    cp -f ${workspace}/bin/geth /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/

    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
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
    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
    done
}

function cluster_restart() {
    yes | cp -r ${workspace}/stop_geth.sh /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/
    yes | cp -r ${workspace}/start_geth.sh /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/

    for ((i = 0; i < ${#bsc_node_ips[@]}; i++)); do
        dst_id=${ips2ids[${bsc_node_ips[i]}]}
        aws ssm send-command \
            --instance-ids "${dst_id}" \
            --document-name "AWS-RunShellScript" \
            --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
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
        sed -i -e "s/_validator${i}_/${acc}/g" ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002002.txt
        bls=$(${workspace}/bin/geth bls account list --datadir ${workspace}/.local/bsc/clusterNetwork/node${i} --blspassword ${workspace}/.local/bsc/password.txt | grep -E -o '0x[0-9a-fA-F]+' | sed 's/0x//')
        sed -i -e "s/_vote${i}_/${bls}/g" ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002002.txt
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

    sed -i -e "s/_rialto_parlia_period_/${BSC_BLCOK_INTERVAL}/g" ${workspace}/tmp/bsc-feynman/bsc/params/config.go
    sed -i -e "s/_rialto_parlia_epoch_/${BSC_EPOCH}/g" ${workspace}/tmp/bsc-feynman/bsc/params/config.go
    
    sed -i -e "s/ValidatorContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000001000.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/SlashContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000001001.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/SystemRewardContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000001002.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/TokenHubContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000001004.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/GovHubContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000001007.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/StakingContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002001.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/StakeHubContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002002.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/StakeCreditContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002003.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/GovernorContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002004.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/GovTokenContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002005.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/TimelockContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000002006.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go
    sed -i -e "s/TokenRecoverPortalContractByteCode/$(cat ${workspace}/tmp/bsc_fyenman_bytecode/0x0000000000000000000000000000000000003000.txt)/g" ${workspace}/tmp/bsc-feynman/bsc/core/systemcontracts/upgrade.go

    cd ${workspace}/tmp/bsc-feynman/bsc && make geth
    yes | cp -f ${workspace}/tmp/bsc-feynman/bsc/build/bin/geth ${workspace}/bin/geth_feynman
    yes | cp -f ${workspace}/tmp/bsc-feynman/bsc/build/bin/geth /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/

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

    # create new validator
    rm -rf ${workspace}/.local/bsc/new_validator${validator_index}
    rm -rf ${workspace}/.local/bsc/new_validator${validator_index}_operator
    mkdir -p ${workspace}/.local/bsc/new_validator${validator_index}
    mkdir -p ${workspace}/.local/bsc/new_validator${validator_index}_operator
    cons_addr=$(${workspace}/bin/geth account new --datadir ${workspace}/.local/bsc/new_validator${validator_index} --password ${workspace}/.local/bsc/password.txt | grep "Public address of the key:" | awk -F"   " '{print $2}')
    operator_addr=$(${workspace}/bin/geth account new --datadir ${workspace}/.local/bsc/new_validator${validator_index}_operator --password ${workspace}/.local/bsc/password.txt | grep "Public address of the key:" | awk -F"   " '{print $2}')
    ${workspace}/bin/geth bls account new --datadir ${workspace}/.local/bsc/new_validator${validator_index} --blspassword ${workspace}/.local/bsc/password.txt
    vote_addr=0x$(cat ${workspace}/.local/bsc/new_validator${validator_index}/bls/keystore/*json| jq .pubkey | sed 's/"//g')
    vote_addr_proof="$(${workspace}/bin/geth_feynman bls account generate-proof --datadir ${workspace}/.local/bsc/new_validator${validator_index} --chain-id ${BSC_CHAIN_ID} --blspassword ${workspace}/.local/bsc/password.txt ${vote_addr} | grep -E -o '0x[0-9a-fA-F]+')"

    transfer_amt=${BSC_CREATE_DELEGATE_AMOUNT}
    ${workspace}/bin/migrate_tool -priv_key ${BSC_TX_BOT_ADDR_PRV} -bsc_endpoint ${BSC_NODE_URL} \
     -amount ${transfer_amt} -to ${operator_addr}
    transfer_amt=10000000000000000000
    ${workspace}/bin/migrate_tool -priv_key ${BSC_TX_BOT_ADDR_PRV} -bsc_endpoint ${BSC_NODE_URL} \
     -amount ${transfer_amt} -to ${operator_addr}

    operator_priv_file=""
    for f in ${workspace}/.local/bsc/new_validator${validator_index}_operator/keystore/*;do
        operator_priv_file=${f}
    done
    operator_priv=$(${workspace}/bin/migrate_tool -secret ${operator_priv_file} -password ${KEYPASS})

    ${workspace}/bin/migrate_tool -priv_key ${operator_priv} -bsc_endpoint ${BSC_NODE_URL} \
     -delegation ${BSC_CREATE_DELEGATE_AMOUNT} \
     -consensus_addr ${cons_addr} -vote_addr ${vote_addr} -vote_bls_proof ${vote_addr_proof} \
     -commission_rate 800 -commission_max_rate 950 -commission_max_change_rate 300 \
     -moniker "Nval${validator_index}" -details ${cons_addr} -identity ${operator_addr}

    while true; do
        if ${workspace}/bin/migrate_tool -get_validator_set -bsc_endpoint "${BSC_NODE_URL}" | grep -q "${cons_addr}"; then
            echo "New validator is ready"
            break
        fi
        echo "Wait for new validator to be ready"
        sleep 3
    done

    rm -rf /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/clusterNetwork/node${validator_index}/keystore
    rm -rf /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/clusterNetwork/node${validator_index}/bls
    mkdir -p /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/clusterNetwork/node${validator_index}/keystore
    mkdir -p /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/clusterNetwork/node${validator_index}/bls
    yes | cp -rf  ${workspace}/.local/bsc/new_validator${validator_index}/keystore/* /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/clusterNetwork/node${validator_index}/keystore/
    yes | cp -rf  ${workspace}/.local/bsc/new_validator${validator_index}/bls/* /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/clusterNetwork/node${validator_index}/bls/

    dst_id=${ips2ids[${bsc_node_ips[validator_index]}]}
    aws ssm send-command \
        --instance-ids "${dst_id}" \
        --document-name "AWS-RunShellScript" \
        --parameters commands="mkdir -p /server/bsc/ && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/stop_geth.sh /server/bsc/stop_geth.sh && yes | cp -f /mnt/efs/bsc-qa/bc-fusion-gov-env/bsc_cluster/start_geth.sh /server/bsc/start_geth.sh && sudo bash /server/bsc/stop_geth.sh"
    sleep 10
    aws ssm send-command \
        --instance-ids "${dst_id}" \
        --document-name "AWS-RunShellScript" \
        --parameters commands="sudo bash +x /server/bsc/start_geth.sh ${validator_index}"

    # wait epoch sync
    echo "wait for epoch sync validator set $((BSC_BLCOK_INTERVAL * BSC_EPOCH + 60))"
    sleep $((BSC_BLCOK_INTERVAL * BSC_EPOCH + 60))
}

function unbond_validator_on_bc() {
    validator_index=$1
    # Check if the input is empty
    if [ -z "$validator_index" ]; then
      echo "Please provide a validator index as input."
      exit 1
    fi

    # unbound old validator on bc
    validator_addr=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${validator_index} | grep node${validator_index} | awk '$1 == "node'${validator_index}'-delegator" {print $3}')
    validator_addr=$(echo ${validator_addr} | xargs ${workspace}/bin/tool -network-type 0 -bsc-val-addr)
    echo "${KEYPASS}" | ${workspace}/bin/tbnbcli staking bsc-unbond \
     --chain-id ${BC_CHAIN_ID} \
     --side-chain-id ${BSC_CHAIN_NAME} \
     --from node${validator_index}-delegator \
     --validator ${validator_addr} \
     --amount ${BSC_INIT_DELEGATE_AMOUNT}:BNB \
     --node ${BC_NODE_URL} --trust-node \
     --home ${workspace}/.local/bc/node${validator_index}

    # bc_block_interval=$(echo "${BC_BLOCK_TIMEOUT}" | sed 's/s$//')
    # t=$((BC_BREATHE_BLOCK_INTERVAL * bc_block_interval))
    # echo "wait for breath block $t s"
    # sleep $t
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
    cluster_restart
    wait_for_hardfork ${hardfork_time}
    echo "===== end ===="
    ;;
migrate_validator)
    echo "===== migrate_validator ===="
    migrate_validator $2
    echo "===== end ===="
    ;;
migrate_all_validator)
    echo "===== migrate_all_validator ===="
    unbond_validator_on_bc 0
    migrate_validator 0
    unbond_validator_on_bc 1
    migrate_validator 1
    migrate_validator 2
    echo "===== end ===="
    ;;
unbond_validator_on_bc)
    echo "===== unbond_validator_on_bc ===="
    unbond_validator_on_bc $2
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_bsc_node.sh cluster_up | cluster_down | cluster_restart | fyenman_hardfork | migrate_validator | migrate_all_validator | unbond_validator_on_bc"
    ;;
esac
