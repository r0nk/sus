#!/bin/bash

for i in $(seq 1 100000000); do
	echo $RANDOM;
done > testfile.txt
