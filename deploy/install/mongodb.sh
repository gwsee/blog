 docker run -d --name mongodb --restart always  --network etcd_network -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin mongo mongod --auth
