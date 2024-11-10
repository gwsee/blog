docker pull mysql:5.7

touch /data/mysql/my.conf

docker run -p 3316:3306 --name mysql01 \
--restart=always \
-v /data/mysql/conf:/etc/mysql/conf.d \
-v /data/mysql/my.conf:/etc/mysql/my.conf.d/mysqld.cnf \
-v /data/mysql/logs:/logs \
-v /data/mysql/data:/var/lib/mysql \
-e TZ=Asia/Shanghai \
-e MYSQL_ROOT_PASSWORD=rzco.otabcfl \
-d mysql:5.7

// mysql 第二版安装
mkdir -p /data/mysql/conf && mkdir -p /data/mysql/data && mkdir -p /data/mysql/log

docker run --name mysql01 -p 3316:3306 \
-v /data/mysql/conf/my.cnf:/etc/mysql/my.cnf \
-v /data/mysql/data:/var/lib/mysql \
-v /data/mysql/log:/var/log/mysql \
-e MYSQL_ROOT_PASSWORD=rzco.otabcfl --restart=always \
-d mysql:5.7

##
