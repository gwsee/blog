# !/bin/bash
version=''
if (( ${#1} > 0 )); then
	version=$1
else
	echo "请传入需要发布的版本"
	exit
fi
docker pull registry.cn-hangzhou.aliyuncs.com/pf94514203/blog:rpc.tools.${version}
docker rm -fv rpc-tools
docker run -d --name=rpc-tools --restart=always \
      -p 998:88 \
      --network etcd_network \
      -v /data/golang/blog/deploy/cfg/rpc.tools:/cfg:rw \
      -v /data/golang/blog/log:/data:rw \
      --link myredis:redis \
      -e PARAMS="-conf /cfg" \
      registry.cn-hangzhou.aliyuncs.com/pf94514203/blog:rpc.tools.${version}