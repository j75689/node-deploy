#!/bin/bash

find /qa/node-deploy/*.sh -type f -exec sed -i 's/benchmark/benchmark/g' {} \;
