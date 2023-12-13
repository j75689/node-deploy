#!/usr/bin/env bash
basedir=$(cd `dirname $0`; pwd)
workspace=${basedir}
source ${workspace}/.env
source ${workspace}/utils.sh

function start() {
    ps -ef | grep bsc_tx_bot.sh |awk '{print $2}' | xargs kill
}

function stop() {
    rm -rf ${workspace}/tmp/txbot
    mkdir -p ${workspace}/tmp/txbot
    cd ${workspace}/tmp/txbot
    nohup bash ${workspace}/bsc_tx_bot.sh > ${workspace}/tmp/txbot/txbot.log 2>&1 &
}

source ${workspace}/.env
CMD=$1

case ${CMD} in
start)
    echo "===== stop ===="
    stop
    sleep 5
    echo "===== stop end ===="

    echo "===== start ===="
    start
    echo "===== start end ===="
    ;;
stop)
    echo "===== stop ===="
    stop
    sleep 5
    echo "===== stop end ===="
    ;;
*)
    echo "Usage: setup_bsc_tx_bot.sh start | stop"
    ;;
esac