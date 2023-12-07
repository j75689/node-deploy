#!/bin/bash
ps -ef | grep bnc-http-ap | awk '{print $2}' | xargs kill