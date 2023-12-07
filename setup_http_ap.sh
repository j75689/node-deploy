#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh

dst_id="i-0d2b8632af953d0f6"

function start_http_ap() {
    rm -rf /mnt/efs/bsc-qa/bc-fusion/http_ap
    mkdir -p /mnt/efs/bsc-qa/bc-fusion/http_ap
    yes | cp -rf ${workspace}/stop_http_ap.sh /mnt/efs/bsc-qa/bc-fusion/http_ap/
    yes | cp -rf ${workspace}/start_http_ap.sh /mnt/efs/bsc-qa/bc-fusion/http_ap/

    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="mkdir -p /server/http-ap && yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/http_ap/start_http_ap.sh /server/http-ap/start_http_ap.sh && bash /server/http-ap/start_http_ap.sh"
}

function stop_http_ap() {
    aws ssm send-command \
      --instance-ids "${dst_id}" \
      --document-name "AWS-RunShellScript" \
      --parameters commands="mkdir -p /server/http-ap && yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/http_ap/stop_http_ap.sh /server/http-ap/stop_http_ap.sh && bash /server/http-ap/stop_http_ap.sh"
}

source ${workspace}/.env
CMD=$1

case ${CMD} in
start)
    echo "===== stop ===="
    stop_http_ap
    sleep 5
    echo "===== stop end ===="

    echo "===== start ===="
    start_http_ap
    echo "===== start end ===="
    ;;
stop)
    echo "===== stop ===="
    stop_http_ap
    sleep 5
    echo "===== stop end ===="
    ;;
*)
    echo "Usage: setup_http_ap.sh start | stop"
    ;;
esac