# !/bin/bash
version=''
if (( ${#1} > 0 )); then
	version=$1
else
	echo "请传入需要发布的版本"
	exit
fi
docker pull registry.cn-hangzhou.aliyuncs.com/pf94514203/blog:bff.${version}
docker rm -fv bff
docker run -d --name=bff --restart=always \
      -p 8080:88 \
      --network etcd_network \
      -v /data/golang/blog/deploy/cfg/bff:/cfg:rw \
      -v /data/golang/blog/log:/data:rw \
      --link myredis:redis \
      -e PARAMS="-conf /cfg" \
      registry.cn-hangzhou.aliyuncs.com/pf94514203/blog:bff.${version}
