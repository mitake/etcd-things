#! /bin/bash

ETCD_BIN=/home/ubuntu/etcd

declare -a nodes=(172.31.41.106 172.31.41.103 172.31.41.104)

for i in `seq 1 3`;
do
    NODE=${nodes[(($i - 1))]}
    ssh etcd-srv$i "rm -rf infra$i.etcd"
    ssh -f etcd-srv$i "$ETCD_BIN --name infra$i --listen-client-urls http://$NODE:12379 --advertise-client-urls http://$NODE:12379 --listen-peer-urls http://$NODE:12380 --initial-advertise-peer-urls http://$NODE:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://172.31.41.106:12380,infra2=http://172.31.41.103:12380,infra3=http://172.31.41.104:12380' --initial-cluster-state new"
done

