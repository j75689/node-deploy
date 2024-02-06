#!/bin/bash

mkdir -p /server/token_approver/
yes | cp -rf /mnt/efs/bsc-qa/benchmark/approval_service/approver /server/token_approver/approver
yes | cp -rf /mnt/efs/bsc-qa/benchmark/dump_bc_account/output/accounts.json /server/token_approver/accounts.json
yes | cp -rf /mnt/efs/bsc-qa/benchmark/dump_bc_account/output/proofs.json /server/token_approver/merkle_proofs.json


export CHAIN_ID="Binance-Chain-Ganges"
merkleRoot=$(cat /mnt/efs/bsc-qa/benchmark/dump_bc_account/output/base.json | jq -r '.state_root')
export MERKLE_ROOT=${merkleRoot}
export HTTP_PORT=8546
export LOGGER_LEVEL=DEBUG
export LOGGER_FORMAT=json
export SECRET_LOCAL_SECRET_PRIVATE_KEY=faec46db1c5b777dfce177421d639a432eef2644d1e79866136b4d80b1e41453
export STORE_MEMORY_STORE_ACCOUNTS=/server/token_approver/accounts.json
export STORE_MEMORY_STORE_MERKLE_PROOFS=/server/token_approver/merkle_proofs.json

nohup /server/token_approver/approver > /server/token_approver/approver.log 2>&1 &