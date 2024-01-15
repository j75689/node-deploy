#!/bin/bash

mkdir -p /server/bc/node-dump
rm -rf /server/bc/node-dump/output
rm -rf /server/bc/node-dump/dump.log

mkdir -p /server/bc/node-dump/output
/mnt/efs/bc-fusion/node-deploy/bin/dump export /server/bc/node-dump/output --home /server/gaiad.0 > /server/bc/node-dump/dump.log 2>&1

rm -rf /server/token_approver.env
echo CHAIN_ID=Binance-Chain-Ganges >> /server/token_approver.env
merkleRoot=$(cat /mnt/efs/bsc-qa/bc-fusion/dump_bc_account/output/base.json | jq -r '.state_root')
echo MERKLE_ROOT=${merkleRoot} >> /server/token_approver.env
echo HTTP_PORT=8088 >> /server/token_approver.env
echo LOGGER_LEVEL=DEBUG >> /server/token_approver.env
echo LOGGER_FORMAT=json >> /server/token_approver.env
echo SECRET_LOCAL_SECRET_PRIVATE_KEY=faec46db1c5b777dfce177421d639a432eef2644d1e79866136b4d80b1e41453 >> /server/token_approver.env
echo STORE_MEMORY_STORE_ACCOUNTS=/server/token_approver/accounts.json >> /server/token_approver.env
echo STORE_MEMORY_STORE_MERKLE_PROOFS=/server/token_approver/merkle_proofs.json >> /server/token_approver.env
