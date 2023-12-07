#!/bin/bash
ps -ef | grep bsc-relayer  | grep config |awk '{print $2}' | xargs kill
