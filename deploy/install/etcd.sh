docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --name etcd \
  --network etcd_network \
  --restart=always \
  quay.io/coreos/etcd:v3.5.0 \
  /usr/local/bin/etcd \
  -advertise-client-urls http://0.0.0.0:2379 \
  -listen-client-urls http://0.0.0.0:2379 \
  -initial-advertise-peer-urls http://0.0.0.0:2380 \
  -listen-peer-urls http://0.0.0.0:2380 \
  -initial-cluster default=http://0.0.0.0:2380 \
  --log-level info \
  --logger zap \
  --log-outputs stderr