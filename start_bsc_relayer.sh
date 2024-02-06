#!/bin/bash

cmd=$1

mkdir -p /server/relayer/
yes | cp -rf /mnt/efs/bsc-qa/benchmark/relayer/* /server/relayer/
if [ "${cmd}" == "reset" ]; then
    rm -rf /server/relayer/bsc_relayer.db
fi

sudo nohup /server/relayer/bsc-relayer --bbc-network-type 0 --config-type local --config-path /server/relayer/bsc_relayer.json > /server/relayer/bsc_relayer.log 2>&1 &
