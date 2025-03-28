# !/bin/bash
version=''
if (( ${#1} > 0 )); then
	version=$1
else
	echo "请传入需要发布的版本"
	exit
fi
docker pull registry.cn-hangzhou.aliyuncs.com/pf94514203/blog:rpc.user.${version}
docker rm -fv rpc-user
docker run -d --name=rpc-user --restart=always \
      -p 1008:88 \
      --network etcd_network \
      -v /data/golang/blog/deploy/cfg/rpc.user:/cfg:rw \
      -v /data/golang/blog/log:/data:rw \
      --link myredis:redis \
      -e PARAMS="-conf /cfg" \
      registry.cn-hangzhou.aliyuncs.com/pf94514203/blog:rpc.user.${version}