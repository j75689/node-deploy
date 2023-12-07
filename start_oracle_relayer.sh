#!/bin/bash

cmd=$1

mkdir -p /server/relayer/
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/relayer/* /server/relayer/
if [ "${cmd}" == "reset" ]; then
    rm -rf /server/relayer/oracle_relayer.db
fi
rm -rf /server/relayer/oracle_relayer.db

nohup /server/relayer/oracle-relayer --bbc-network 0 --config-type local --config-path /server/relayer/oracle_relayer.json > /server/relayer/oracle_relayer.log 2>&1 &
