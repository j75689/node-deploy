#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env

while true; do
  $workspace/bin/migrate_tool -bsc_endpoint ${BSC_NODE_URL} \
    -amount 1000000000000000000 -gas_price 10000000000000 \
    -to ${BSC_TX_BOT_ADDR} \
    -priv_key=${BSC_TX_BOT_ADDR_PRV}
done
