#!/bin/bash
ps -ef | grep bnbchaind | awk '{print $2}' | xargs kill
