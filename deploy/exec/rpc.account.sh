# !/bin/bash
version=''
if (( ${#1} > 0 )); then
	version=$1
else
	echo "请传入需要发布的版本"
	exit
fi
docker pull registry.cn-chengdu.aliyuncs.com/pf94514203/blog:rpc.account.${version}
docker rm -fv rpc-account
docker run -d --name=rpc-account --restart=always \
      -p 8088:88 \
      --network etcd_network \
      -v /data/golang/blog/deploy/cfg/rpc.account:/cfg:rw \
      -v /data/golang/blog/log:/data:rw \
      --link redis:redis \
      -e PARAMS="-conf /cfg" \
      registry.cn-chengdu.aliyuncs.com/pf94514203/yunchuang:rpc.account.${version}