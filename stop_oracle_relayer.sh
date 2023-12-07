#!/bin/bash
ps -ef  | grep oracle-relayer | grep config |awk '{print $2}' | xargs kill
