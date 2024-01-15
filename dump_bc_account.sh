#!/bin/bash

mkdir -p /server/bc/node-dump
rm -rf /server/bc/node-dump/output
rm -rf /server/bc/node-dump/dump.log

mkdir -p /server/bc/node-dump/output
/mnt/efs/bc-fusion/node-deploy/bin/dump export /server/bc/node-dump/output --home /server/gaiad.0 > /server/bc/node-dump/dump.log 2>&1
