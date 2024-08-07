#!/usr/bin/env bash

dist=`uname -a`

if [[ $dist == *"Darwin"* ]]; then
    sed -i '' 's/"type": "null"/"nullable": true/g' $@
else
    sed -i 's/"type": "null"/"nullable": true/g' $@
fi
