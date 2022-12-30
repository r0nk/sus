#!/bin/bash

go build .
time cat testfile.txt | sort | uniq -c | sort -rn > /dev/null
time cat testfile.txt | ./sus > /dev/null
