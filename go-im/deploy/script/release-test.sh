#!/bin/bash

need_start_server_shell=(
  # "rpc.sh"
  "user-rpc-dev.sh"
  "social-rpc-dev.sh"
  "im-ws-dev.sh"
  "im-rpc-dev.sh"
  # "api.sh"
  "user-api-dev.sh"
  "social-api-dev.sh"
  "im-api-dev.sh"

  # task
  "task-mq-dev.sh"
)

for i in ${need_start_server_shell[*]}; do
  chmod +x $i
  ./$i
done

docker ps
docker exec -it etcd etcdctl get --prefix ""