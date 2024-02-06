#!/bin/bash

mkdir -p /server/bc/node-dump
rm -rf /server/bc/node-dump/output
rm -rf /server/bc/node-dump/dump.log

yes | cp -rf /mnt/efs/bsc-qa/benchmark/dump_bc_account/dump /server/bc/node-dump/dump

mkdir -p /server/bc/node-dump/output
/server/bc/node-dump/dump export /server/bc/node-dump/output --home /server/bc/node > /server/bc/node-dump/dump.log 2>&1

rm -rf /mnt/efs/bsc-qa/benchmark/dump_bc_account/output
yes | cp -rf /server/bc/node-dump/output /mnt/efs/bsc-qa/benchmark/dump_bc_account/output