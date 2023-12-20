#!/bin/bash
ps -ef | grep approver |awk '{print $2}' | xargs kill
