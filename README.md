# Zookeeper REST API

For managing Zookeeper configuration data, start up this thin REST API server and connect.

## Installing

```
go get github.com/freenowtech/zk-rest-api
```

## Usage

```
Usage of ./zk-rest-api:
  -addr=":8080": address and port to listen on
  -zk=[]:        comma-separated list of hosts to zookeeper.
```

## Example

```
zk-rest-api -addr localhost:8001 -zk zookeeper-1,zookeeper-2:2081
```

## License

MIT
