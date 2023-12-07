#!/bin/bash

nohup /server/http-ap/build/bnc-http-ap --scheme=http --port=8546 --host=0.0.0.0 2>&1 &