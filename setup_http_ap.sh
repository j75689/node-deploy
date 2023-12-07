#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh

dst_id="i-0d2b8632af953d0f6"

function start_http_ap() {
    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="cd /server/http-ap/build && nohup ./bnc-http-ap --scheme=http --port=8546 --host=0.0.0.0 2>&1 &"
}

function stop_http_ap() {
    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="ps -ef | grep bnc-http-ap | awk '{print $2}' | xargs kill"
}