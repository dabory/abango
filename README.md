# abango
## Microservices Framework to support Apache Kafka, gRpc and REST API

### Abango
is a golang web framework which supports Kafka, gRpc and RESTful API at the same time in single/multi thread.


#### Step 0: Linux: Ubuntu 16.04.4 LTS \n \l  : recommended Linux version
#### Step 1: Install: Go version go1.11 linux/amd64
#### Step 2: Install: Go libraries
`$ go get -u google.golang.org/grpc`

`$ go get -u github.com/golang/protobuf/protoc-gen-go`

`$ go get xorm.io/xorm`

`$ go get github.com/pilu/fresh`

`$ go get github.com/go-sql-driver/mysql`

`$ go get github.com/dabory/abango`

`$ mkdir -p $GOPATH/bin $GOPATH/src $GOPATH/pkg`

`$ go get github.com/dabory/svc-abango `

`$ go get github.com/dabory/end-abango `

`$ cd $GOPATH/src/github.com/dabory/end-abango`

`$ cd $GOPATH/src/github.com/dabory/kafka-docker`

#### Step 3: You should run Kafka API first if you want to use Kafka

`$ cd $GOPATH/src/github.com/dabory/kafka-docker`

`$ docker-compose up`


#### Step 4: Run abango service 
`$ cd $GOPATH/src/github.com/dabory/svc-abango`

`$ vi conf/config_select.json` ; specify config.json file

`$ vi conf/xxx_config.json`  ; change parameter values liked conf file

#### Step 4: Change Open port if you need in conf/ folder files





#### Step 5: Run docker-composein Kafka folder
To run the server, use this command

`$ fresh`

