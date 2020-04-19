#### Generate gRPC with protocol buffer:
```shell
protoc blogpb/blog.proto --go_out=plugins=grpc:.
```
### start mongodb
```shell
docker-compose up -d
```
