#!/bin/bash

source ./content

size=${#array[@]}
echo ""
echo "Message of the Day:"
echo "${array[$(($RANDOM % $size))]}"
echo ""
