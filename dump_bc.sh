#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh

dst_id="i-0870976185c85a30d"

function dump_bc_account() {
    rm -rf /mnt/efs/bsc-qa/benchmark/dump_bc_account
    mkdir -p /mnt/efs/bsc-qa/benchmark/dump_bc_account
    yes | cp -rf ${BC_DUMP_TOOL_BINARY} /mnt/efs/bsc-qa/benchmark/dump_bc_account/dump
    yes | cp -rf ${workspace}/dump_bc_account.sh /mnt/efs/bsc-qa/benchmark/dump_bc_account/

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="mkdir -p /server/bc/node-dump && yes | cp -rf /mnt/efs/bsc-qa/benchmark/dump_bc_account/dump_bc_account.sh /server/bc/node-dump/dump_bc_account.sh && bash /server/bc/node-dump/dump_bc_account.sh"
}

dump_bc_account