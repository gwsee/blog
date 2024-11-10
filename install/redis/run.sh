docker run --name redis -d -p 6379:6379 \
  --restart=always \
  -v /data/redis/config/redis.conf:/etc/redis/redis.conf \
  -v /data/redis/data:/data redis:latest \
  --requirepass OoO0o0ilLi1I1

 ###或者

  docker run \
  --name myredis \
  --restart=always \
  --log-opt max-size=100m \
  --log-opt max-file=2 \
  -p 6379:6379 \
  -v /data/redis/redis.conf:/etc/redis/redis.conf \
  -v /data/redis/data:/data \
  -d redis redis-server /etc/redis/redis.conf \
  --appendonly yes \
  --requirepass 1q2wroot
