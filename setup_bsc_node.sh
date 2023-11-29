#!/usr/bin/env bash
basedir=$(
    cd $(dirname $0)
    pwd
)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh
size=$((${BSC_CLUSTER_SIZE}))
standalone=true
initIp="172.22.42.13,172.22.42.94,172.22.43.86"

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
        if [ ${standalone} = true ]; then
            continue
        fi
        
        node_dir_index=${i}
        if [ $i -ge ${BC_CLUSTER_SIZE} ]; then
            # echo "${KEYPASS}" | ${workspace}/bin/tbnbcli keys delete node${i}-delegator --home ${workspace}/.local/bc/node0 # for re-entry
            echo "${KEYPASS}" | (echo "${KEYPASS}" | ${workspace}/bin/tbnbcli keys add node${i}-delegator --home ${workspace}/.local/bc/node0)
            node_dir_index=0
        fi
        delegator=$(${workspace}/bin/tbnbcli keys list --home ${workspace}/.local/bc/node${node_dir_index} | grep node${i}-delegator | awk -F" " '{print $3}')
        if [ "$i" != "0" ]; then
            sleep 6 #wait for including tx in block
            echo "${KEYPASS}" | ${workspace}/bin/tbnbcli send --from node0-delegator --to $delegator --amount ${BSC_INIT_DELEGATE_AMOUNT}:BNB --chain-id ${BC_CHAIN_ID} --node ${nodeurl} --home ${workspace}/.local/bc/node0
        fi
        sleep 6 #wait for including tx in block
        echo ${delegator} "balance"
        ${workspace}/bin/tbnbcli account ${delegator}  --chain-id ${BC_CHAIN_ID} --trust-node --home ${workspace}/.local/bc/node${node_dir_index} | jq .value.base.coins
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

function create_validator() {
    rm -rf ${workspace}/clusterNetwork

    for ((i = 0; i < ${size}; i++)); do
        mkdir -p ${workspace}/clusterNetwork/validator${i}
        echo "${KEYPASS}" >${workspace}/clusterNetwork/password.txt

        cons_addr=$(${workspace}/bin/geth account new --datadir ${workspace}/clusterNetwork/validator${i} --password ${workspace}/clusterNetwork/password.txt | grep "Public address of the key:" | awk -F"   " '{print $2}')
        fee_addr=$(${workspace}/bin/geth account new --datadir ${workspace}/clusterNetwork/validator${i}_fee --password ${workspace}/clusterNetwork/password.txt | grep "Public address of the key:" | awk -F"   " '{print $2}')
        mkdir -p ${workspace}/clusterNetwork/bls${i}
        expect create_bls_key.exp ${workspace}/clusterNetwork/bls${i} ${KEYPASS}
        vote_addr=0x$(cat ${workspace}/clusterNetwork/bls${i}/bls/keystore/*json | jq .pubkey | sed 's/"//g')
        if [ ${standalone} = true ]; then
            continue
        fi
    done
}

function generate_static_peers() {
    tool=${workspace}/bin/bootnode
    num=$1
    target=$2
    staticPeers=""
    for ((i = 0; i < $num; i++)); do
        if [ $i -eq $target ]; then
            continue
        fi

        file=${workspace}/clusterNetwork/node${i}/geth/nodekey
        if [ ! -f "$file" ]; then
            $tool -genkey $file
        fi
        port=30311
        domain="bsc-node-${i}.bsc.svc.cluster.local"
        if [ ! -z "$staticPeers" ]; then
            staticPeers+=","
        fi
        staticPeers+='"'"enode://$($tool -nodekey $file -writeaddress)@$domain:$port"'"'
    done

    echo $staticPeers
}

function prepare_config() {
    rm -f ${workspace}/genesis/validators.conf

    for ((i = 0; i < ${size}; i++)); do
        for f in ${workspace}/clusterNetwork/validator${i}/keystore/*; do
            cons_addr="0x$(cat ${f} | jq -r .address)"
        done

        for f in ${workspace}/clusterNetwork/validator${i}_fee/keystore/*; do
            fee_addr="0x$(cat ${f} | jq -r .address)"
        done

        mkdir -p ${workspace}/clusterNetwork/node${i}
        bbcfee_addrs=${fee_addr}
        powers="0x000001d1a94a2000"
        mv ${workspace}/clusterNetwork/bls${i}/bls ${workspace}/clusterNetwork/node${i}/ && rm -rf ${workspace}/clusterNetwork/bls${i}
        vote_addr=0x$(cat ${workspace}/clusterNetwork/node${i}/bls/keystore/*json | jq .pubkey | sed 's/"//g')
        echo "${cons_addr},${bbcfee_addrs},${fee_addr},${powers},${vote_addr}" >>${workspace}/genesis/validators.conf
        echo "validator" ${i} ":" ${cons_addr}
        echo "validatorFee" ${i} ":" ${fee_addr}
        echo "validatorVote" ${i} ":" ${vote_addr}
    done

    cd ${workspace}/genesis/
    git stash push -- contracts/
    sed -i -e "s/'0x' + publicKey.pop()/vs[4]/g" scripts/generate-validator.js
    node scripts/generate-validator.js
    node scripts/generate-initHolders.js --initHolders ${INIT_HOLDER}
    if [ ${standalone} = false ]; then
        initConsensusStateBytes=$(${workspace}/bin/tool -height 1 -rpc ${nodeurl} -network-type 0)
        bash scripts/generate.sh local --chainId ${BSC_CHAIN_ID} --whitelist1Address ${INIT_HOLDER} --initConsensusStateBytes ${initConsensusStateBytes}
    else
        bash scripts/generate.sh local --chainId ${BSC_CHAIN_ID} --whitelist1Address ${INIT_HOLDER}
    fi
}

function initNetwork() {
    cd ${workspace}
    ${workspace}/bin/geth init-network --init.dir ${workspace}/clusterNetwork --init.size=${size} --init.ips ${initIp} --config ${workspace}/config.toml ${workspace}/genesis/genesis.json
    rm -rf ${workspace}/*bsc.log*
    for ((i = 0; i < ${size}; i++)); do
        sed -i -e '/"<nil>"/d' ${workspace}/clusterNetwork/node${i}/config.toml

        cp -R ${workspace}/clusterNetwork/validator${i}/keystore ${workspace}/clusterNetwork/node${i}
        for j in ${workspace}/clusterNetwork/validator${i}/keystore/*; do
            cons_addr="0x$(cat ${j} | jq -r .address)"
        done

        cp ${workspace}/bin/geth ${workspace}/clusterNetwork/node${i}/geth${i}
        # init genesis
        ${workspace}/clusterNetwork/node${i}/geth${i} --datadir ${workspace}/clusterNetwork/node${i} init ${workspace}/genesis/genesis.json
    done
}

CMD=$1
case ${CMD} in
register)
    echo "===== register ===="
    register_validator
    echo "===== end ===="
    ;;
generate)
    echo "===== generate configs ===="
    create_validator
    prepare_config
    initNetwork
    echo "===== end ===="
    ;;
native_init)
    echo "===== register ===="
    create_validator
    echo "===== end ===="
    echo "===== generate configs ===="
    prepare_config
    initNetwork
    echo "===== end ===="
    ;;
*)
    echo "Usage: setup_bsc_node.sh | register | generate | native_init"
    ;;
esac
