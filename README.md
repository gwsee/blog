# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

计划中：
    后端：golang
    前端：vue3 
    管理：react？目前好像不需要
功能点--现在想到的
    首页--登录;
         游历(简历)
         留言
         随笔(日常记录)
         学习(学习记录)
         面试(随机面试;自己选择tag+其他的)


## 对于单个或者说小团队 bff层不建议


## CMD
```bash
kratos new blog
cd blog

kratos new app/blog --nomod
kratos proto add api/blog/v1/demo.proto
kratos proto client api/blog/v1/demo.proto ##不会生成http文件
kratos proto client api/blog/v1/demo.proto  -- --go-http_opt=omitempty=false ##会生成http文件
kratos proto server api/blog/v1/demo.proto -t internal/service
kratos proto client api/tools/v1/tools.proto  -- --go-http_opt=omitempty=false

protoc --go_out=./ conf.proto --go_opt=paths=source_relative
protoc --go_out=./ page.proto --go_opt=paths=source_relative

protoc --go_out=./ index.proto --go_opt=paths=source_relative
protoc --go_out=./ cfg.proto --go_opt=paths=source_relative
protoc --go_out=./ ./app/user/internal/conf/conf.proto --go_opt=paths=source_relative   
protoc --go_out=./ ./app/tools/internal/conf/conf.proto --go_opt=paths=source_relative   
protoc --go_out=./ ./api/global/cfg.proto --go_opt=paths=source_relative   
```
