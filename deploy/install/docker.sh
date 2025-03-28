docker本身

cd /etc/docker/
vi daemon.json ##修改这个文件添加这个配置
sudo systemctl daemon-reload
sudo systemctl restart docker



{
    "registry-mirrors": [
        "https://kz8sdu63.mirror.aliyuncs.com"
    ]
}