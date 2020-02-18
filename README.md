# abango


# How to user example of abango service ; Svc-abango

## Abango
is a golang web framework which supports Kafka, gRpc and RESTful API at the same time in single/multi thread.

To run the server, use this command

#### Step 0: Linux: Ubuntu 16.04.4 LTS \n \l  : recommended Linux version
#### Step 1: Install: Go version go1.11 linux/amd64
#### Step 2: Install: Go libraries
`$ go get -u google.golang.org/grpc`

`$ go get -u github.com/golang/protobuf/protoc-gen-go`

`$ go get xorm.io/xorm`

`$ go get github.com/pilu/fresh`

`$ go get github.com/go-sql-driver/mysql`

`$ go get github.com/dabory/abango `

`$ mkdir -p $GOPATH/bin $GOPATH/src $GOPATH/pkg`

`$ cd $GOPATH/src`

`$ git clone https://github.com/dabory/svc-abango `

To run the grpc server, using this command

`$ cd $GOPATH/src/github.com/dabory/svc-abango`

#### Step 3: Change Open port if you need in conf/ folder files
`$ vi conf/config_select.json`

`$ vi conf/xxx_config.json`  ; liked conf file

#### Step 4: Create MySQL Tables : extract table schema from kangan_db-191125.sql 

#### Step 5: Run go lang server; You should run Kafka first to run Kafka API
#### Step \6: Run docker-compose in Kafka folder

`$ fresh`

