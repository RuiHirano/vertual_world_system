#!/bin/sh

if [ "$1" = "" ]; then
    echo "input point_id"
    echo "bash set_signage.sh [point_id] [neighbor_id] [rate]"
    echo "ex. bash set_signage.sh 0 1 0.5"
    exit 1
elif [ "$2" = "" ]; then
    echo "input neignbor_id"
    echo "bash set_signage.sh [point_id] [neighbor_id] [rate]"
    echo "ex. bash set_signage.sh 0 1 0.5"
    exit 1
elif [ "$3" = "" ]; then
    echo "input rate"
    echo "bash set_signage.sh [point_id] [neighbor_id] [rate]"
    echo "ex. bash set_signage.sh 0 1 0.5"
    exit 1
else 
    curl -X POST -d '{"point_id": 1, "neighbor_id": 2, "ratio": 0.0}'  -H 'Content-Type: application/json'  http://localhost:8000/set/signage
fi