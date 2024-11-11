#!/bin/bash

# 检查是否传入了足够的参数
if [ $# -lt 3 ]; then
    echo "用法: $0 <镜像名称> <标签> <Dockerfile路径>"
    exit 1
fi

# 接收参数
IMAGE_NAME=$1   # 镜像名称
TAG=$2          # 标签
DOCKERFILE_PATH=$3  # Dockerfile 路径
REPO_URL="registry.cn-chengdu.aliyuncs.com/pf94514203"  # 固定的 Docker 仓库地址

# 构建 Docker 镜像
echo "正在构建 Docker 镜像..."
docker build -t $IMAGE_NAME:$TAG -f $DOCKERFILE_PATH .

# 检查构建是否成功
if [ $? -ne 0 ]; then
    echo "Docker 镜像构建失败"
    exit 1
fi

echo "构建成功: $IMAGE_NAME:$TAG"

# 给镜像打标签
echo "正在为镜像打标签..."
docker tag $IMAGE_NAME:$TAG $REPO_URL/$IMAGE_NAME:$TAG

# 检查打标签是否成功
if [ $? -ne 0 ]; then
    echo "镜像打标签失败"
    exit 1
fi

echo "打标签成功: $REPO_URL/$IMAGE_NAME:$TAG"

# 推送镜像到仓库
echo "正在推送镜像到仓库..."
docker push $REPO_URL/$IMAGE_NAME:$TAG

# 检查推送是否成功
if [ $? -ne 0 ]; then
    echo "镜像推送失败"
    exit 1
fi

echo "镜像推送成功: $REPO_URL/$IMAGE_NAME:$TAG"

# 完成
echo "任务完成"

## ./deploy/dockerfile/run.sh blog bff.001 ./deploy/dockerfile/bff.Dockerfile
## ./deploy/dockerfile/run.sh blog rpc.account.001 ./deploy/dockerfile/rpc.account.Dockerfile
## ./deploy/dockerfile/run.sh blog rpc.blogs.001 ./deploy/dockerfile/rpc.blogs.Dockerfile