#!/bin/bash
ps -ef | grep "/server/token_approver/approver" |awk '{print $2}' | xargs kill
