#!/usr/bin/env bash

while read d ;
do
    if [ -f $d/Makefile ] ; then
        echo $d;
        make -C $d ci
    fi
done
