#! /bin/bash

ETCD_BIN=/home/ubuntu/etcd

declare -a nodes=(172.31.41.106 172.31.41.103 172.31.41.104)

for i in `seq 1 3`;
do
    NODE=${nodes[(($i - 1))]}
    ssh etcd-srv$i "pkill etcd"
done

