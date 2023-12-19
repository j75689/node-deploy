#!/bin/bash

mkdir -p /server/bc/node-dump
rm -rf /server/bc/node-dump/output
rm -rf /server/bc/node-dump/dump.log
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/dump_bc_account/dump /server/bc/node-dump/dump

/server/bc/node-dump/dump export /server/bc/node-dump/output --home /server/bnbchaind/disaster/node/gaiad > /server/bc/node-dump/dump.log 2>&1