# anomaly-tracker-proto
Contains the proto definitions of the various dependencies for the anomaly tracker projects.

## Updating
Make changes to the `.proto` files (the `.pb.go` are generated, so don't edit them. Then use one of the following methods to compile it to Go:

#### Local
```bash
$ protoc --go_out=. -I. *.proto
```

#### Docker
```bash
$ docker-compose up
```
