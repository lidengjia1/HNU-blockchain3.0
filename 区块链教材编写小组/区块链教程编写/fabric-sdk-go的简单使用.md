> # fabric-sdk-go的简单使用
>
> 使用fabric-sdk-go之前,需要安装好go环境,docker以及docker compose环境,还有 hyperledger fabric 环境.

### 一.   创建crypto-config.yaml

使用fabric提供的cryptogen工具生成文件模板
 `$ cryptogen showtemplate > crypto-config.yaml`
 进行修改,添加一个组织,一个orderer节点.



```css
OrdererOrgs:

  - Name: Orderer
    Domain: xq.com
    Specs:
      - Hostname: orderer

PeerOrgs:

  - Name: Travle
    Domain: travle.xq.com
    EnableNodeOUs: false

    Template:
      Count: 2

    Users:
      Count: 2
```

根据crypto-config.yaml文件生成证书文件:
 `$ cryptogen generate --config=crypto-config.yaml`

查看生成的证书文件夹结构:



```css
.
├── ordererOrganizations
│   └── xq.com
│       ├── ca
│       │   ├── 1d8deec1977f7abf81692e72c06861e811c00b34278b48c9fe44dc51238d8621_sk
│       │   └── ca.xq.com-cert.pem
│       ├── msp
│       │   ├── admincerts
│       │   ├── cacerts
│       │   └── tlscacerts
│       ├── orderers
│       │   └── orderer.xq.com
│       ├── tlsca
│       │   ├── 5b189397d9390ae07146a411f0e7b6c5bf1fdb809d4b4d51b6f25b739a7cc127_sk
│       │   └── tlsca.xq.com-cert.pem
│       └── users
│           └── Admin@xq.com
└── peerOrganizations
    └── travle.xq.com
        ├── ca
        │   ├── 8ae1b694ae88d19d07fd698b0990b8d4f3bc46782fbfb3620e3e6e3fb0263e1c_sk
        │   └── ca.travle.xq.com-cert.pem
        ├── msp
        │   ├── admincerts
        │   ├── cacerts
        │   └── tlscacerts
        ├── peers
        │   ├── peer0.travle.xq.com
        │   └── peer1.travle.xq.com
        ├── tlsca
        │   ├── d6e2633fe6f35345eca10bc0231a2b34f127c2b9ebc8a86b582b412e875e513c_sk
        │   └── tlsca.travle.xq.com-cert.pem
        └── users
            ├── Admin@travle.xq.com
            ├── User1@travle.xq.com
            └── User2@travle.xq.com
```

### 二. 生成创世区块文件和 通道

需要从fabric的源码案例中拷贝configtx.yaml文件
 `$ cp $GOPATH/src/github.com/hyperledger/fabric-samples/first-network/configtx.yaml ./`
 对configtx.yaml文件进行修改.
 修改之前,创建一个文件夹,来保存即将创建的创世区块文件



```ruby
$ mkdir channel-artifacts
```

将创建区块文件和通道的命令写到一个脚本中! generate.sh



```bash
rm -rf ./channel-artifacts/*
rm -rf ./crypto-config/*

#根据crypto-config.yaml文件生成证书
cryptogen generate --config=./crypto-config.yaml

# 生成创始块文件
echo "---------------- Create genesis.block file BEGIN --------------------"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
echo "---------------- Create genesis.block file END --------------------"

# 生成 travlechannel 文件
echo "---------------- Create travlechannel.tx file BEGIN -------------------"
configtxgen -profile TravleOrgsChannel -outputCreateChannelTx ./channel-artifacts/travlechannel.tx -channelID travlechannel
echo "---------------- Create travlechannel.tx file END -------------------"

# 生成更新锚节点文件
echo "---------------- Create TravleMSPanchors.tx file BEGIN -------------------"
configtxgen -profile TravleOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/TravleMSPanchors.tx -channelID travlechannel -asOrg TravleMSP
echo "---------------- Create TravleMSPanchors.tx file END -------------------"
```

脚本文件和配置文件的目录结构:



```css
├── channel-artifacts
├── configtx.yaml
├── crypto-config
│   ├── ordererOrganizations
│   └── peerOrganizations
├── crypto-config.yaml
└── generate.sh
```

执行generate.sh文件生成创世区块文件和通道,其实只有一个组织,也没必要生成锚节点更新文件..
 `$ ./generate.sh`

### 三. 通过 docker-compose 启动容器

配置docker-compose文件:



```jsx
version: '2'

networks:
  xq_travle:

services:
  orderer.xq.com:
    image: hyperledger/fabric-orderer:latest
    container_name: orderer.xq.com
    environment:
      - ORDERER_GENERAL_LOGLEVEL=debug
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_LISTENPORT=7050
      - ORDERER_GENERAL_GENESISPROFILE=Orderer
      - ORDERER_GENERAL_GENESISMETHOD=file
      - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block
      - ORDERER_GENERAL_LOCALMSPID=xq.com
      - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
      - ORDERER_GENERAL_TLS_ENABLED=true
      - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]

    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
    volumes:
      - ./channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ./crypto-config/ordererOrganizations/xq.com/orderers/orderer.xq.com/msp:/var/hyperledger/orderer/msp
      - ./crypto-config/ordererOrganizations/xq.com/orderers/orderer.xq.com/tls:/var/hyperledger/orderer/tls
    ports:
      - 7050:7050
    networks:
      - xq_travle

  peer0.travle.xq.com:
    image: hyperledger/fabric-peer:latest
    container_name: peer0.travle.xq.com
    environment:
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_ATTACHSTDOUT=true
      - CORE_LOGGING_LEVEL=DEBUG
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=/var/hyperledger/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/var/hyperledger/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/var/hyperledger/tls/ca.crt
      - CORE_PEER_ID=peer0.travle.xq.com
      - CORE_PEER_ADDRESSAUTODETECT=true
      - CORE_PEER_ADDRESS=peer0.travle.xq.com:7051
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.travle.xq.com:7051
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_GOSSIP_SKIPHANDSHAKE=true
      - CORE_PEER_LOCALMSPID=travle.xq.com
      - CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp
      - CORE_PEER_TLS_SERVERHOSTOVERRIDE=peer0.travle.xq.com
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    volumes:
      - /var/run/:/host/var/run/
      - ./crypto-config/peerOrganizations/travle.xq.com/peers/peer0.travle.xq.com/msp:/var/hyperledger/msp
      - ./crypto-config/peerOrganizations/travle.xq.com/peers/peer0.travle.xq.com/tls:/var/hyperledger/tls
    ports:
      - 7051:7051
      - 7053:7053
    depends_on:
      - orderer.xq.com
    links:
      - orderer.xq.com
    networks:
      - xq_travle

  peer1.travle.xq.com:
    image: hyperledger/fabric-peer:latest
    container_name: peer1.travle.xq.com
    environment:
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_ATTACHSTDOUT=true
      - CORE_LOGGING_LEVEL=DEBUG
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=/var/hyperledger/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/var/hyperledger/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/var/hyperledger/tls/ca.crt
      - CORE_PEER_ID=peer1.travle.xq.com
      - CORE_PEER_ADDRESSAUTODETECT=true
      - CORE_PEER_ADDRESS=peer1.travle.xq.com:7051
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer1.travle.xq.com:7051
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_GOSSIP_SKIPHANDSHAKE=true
      - CORE_PEER_LOCALMSPID=travle.xq.com
      - CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp
      - CORE_PEER_TLS_SERVERHOSTOVERRIDE=peer1.travle.xq.com
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    volumes:
      - /var/run/:/host/var/run/
      - ./crypto-config/peerOrganizations/travle.xq.com/peers/peer1.travle.xq.com/msp:/var/hyperledger/msp
      - ./crypto-config/peerOrganizations/travle.xq.com/peers/peer1.travle.xq.com/tls:/var/hyperledger/tls
    ports:
      - 8051:7051
      - 8053:7053
    depends_on:
      - orderer.xq.com
    links:
      - orderer.xq.com
    networks:
      - xq_travle
```

启动容器, 启动后查看容器运行情况
 `$ docker-compose up -d`
 `$ docker-compose ps`



```rust
       Name               Command       State                       Ports                     
----------------------------------------------------------------------------------------------
orderer.xq.com        orderer           Up      0.0.0.0:7050->7050/tcp                        
peer0.travle.xq.com   peer node start   Up      0.0.0.0:7051->7051/tcp, 0.0.0.0:7053->7053/tcp
peer1.travle.xq.com   peer node start   Up      0.0.0.0:8051->7051/tcp, 0.0.0.0:8053->7053/tcp
```

在这里,创建两个脚本文件,用于docker容器的管理
 **clear_docker.sh** 文件:



```ruby
sudo docker rm -f $(sudo docker ps -aq) # 清除容器们
sudo docker network prune # 来清理没有再被任何容器引用的networks
sudo docker volume prune  # 清理挂载卷
```

**restart.sh** 文件:



```bash
./clear_docker.sh    # 执行clear_docker.sh脚本文件
docker-compose up -d
```

### 四. 创建sdk配置文件

创建配置文件的时候,有两个文件可以进行参考...



```bash
$GOPATH/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/config/config_test.yaml
```



```bash
$GOPATH/src/github.com/hyperledger/fabric-sdk-go/pkg/core/config/testdata/template/config.yaml
```

修改后的sdk配置文件:



```tsx
name: "driver-service-network"

version: 1.1.0

client:
  organization: Travle
  logging:
    level: info
  cryptoconfig:
    path: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config
  credentialStore:
    path: /tmp/driverStore

  BCCSP:
    security:
      enabled: true
      default:
        provider: "SW"
      hashAlgorithm: "SHA2"
      softVerify: true
      level: 256

  tlsCerts:
    systemCertPool: true
    client:
      keyfile: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/peerOrganizations/travle.xq.com/users/User1@travle.xq.com/tls/client.key
      certfile: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/peerOrganizations/travle.xq.com/users/User1@travle.xq.com/tls/client.crt

channels:
  travlechannel:
    orderers:
      - orderer.xq.com

    peers:
      peer0.travle.xq.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true

      peer1.travle.xq.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true

    policies:
      queryChannelConfig:
        minResponses: 1
        maxTargets: 1
        retryOpts:
          attempts: 5
          initialBackoff: 500ms
          maxBackoff: 5s
          backoffFactor: 2.0

organizations:
  travle:
    # configtx.yaml organizations -> ID
    mspid: travle.xq.com

    cryptoPath: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/peerOrganizations/travle.xq.com/users/{username}@travle.xq.com/msp
    peers:
    - peer0.travle.xq.com
    - peer1.travle.xq.com

  ordererorg:
    mspID: xq.com
    cryptoPath: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/ordererOrganizations/xq.com/users/Admin@xq.com/msp

orderers:
  orderer.xq.com:
    url: localhost:7050
    grpcOptions:
      ssl-target-name-override: orderer.xq.com
      keep-alive-time: 0s
      keep-alive-timeout: 20s
      keep-alive-permit: false
      fail-fast: false
      allow-insecure: false

    tlsCACerts:
      path: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/ordererOrganizations/xq.com/tlsca/tlsca.xq.com-cert.pem
peers:
  peer0.travle.xq.com:
    url: grpcs://localhost:7051
    eventUrl: localhost:7053
    grpcOptions:
      ssl-target-name-override: peer0.travle.xq.com
      keep-alive-time: 0s
      keep-alive-timeout: 20s
      keep-alive-permit: false
      fail-fast: false
      allow-insecure: false
    tlsCACerts:
      path: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/peerOrganizations/travle.xq.com/tlsca/tlsca.travle.xq.com-cert.pem

  peer1.travle.xq.com:
    url: grpcs://localhost:8051
    eventUrl: localhost:8053
    grpcOptions:
      ssl-target-name-override: peer1.travle.xq.com
      keep-alive-time: 0s
      keep-alive-timeout: 20s
      keep-alive-permit: false
      fail-fast: false
      allow-insecure: false
    tlsCACerts:
      path: /Users/xq_mac/Go/src/driverFabricDemo/conf/crypto-config/peerOrganizations/travle.xq.com/tlsca/tlsca.travle.xq.com-cert.pem
```

### 五. 实例化sdk

- 5.1 在工程中定义一个模型,来维护实例化sdk所需要的变量



```cpp
type FabricModel struct {
    ConfigFile  string //sdk的配置文件路径
    ChainCodeID string // 链码名称
    ChaincodePath string  // 链码在工程中的存放目录
    ChaincodeGoPath string  // GOPATH
    OrgAdmin string // 组织的管理员用户
    OrgName  string // config.yaml ---> organizations ---> travle
    OrgID    string // 组织id
    UserName string // 组织的普通用户
    ChannelID string // 通道id
    ChannelConfigPath string //组织的通道文件路径
    OrdererName string  // config.yaml ---> orderers ---> orderer.xq.com // 将组织添加到通道时候使用!
    Sdk        *fabsdk.FabricSDK // 保存实例化后的sdk
    ResMgmtCli *resmgmt.Client // 资源管理客户端,也需要在安装链码时候的使用
    Channelclient     *channel.Client // 通道客户端
    HasInit  bool // 是否已经初始化了sdk
}
```

创建出一个模型对象,给其赋值,并开始初始化sdk



```go
func init() {
    fs := FabricModel{
        OrdererName: "orderer.xq.com",
        ChannelID:     "travlechannel",
        ChannelConfigPath: os.Getenv("GOPATH") + "/src/driverFabricDemo/conf/channel-artifacts/travlechannel.tx",
        ChainCodeID:     "mycc",
        ChaincodeGoPath: os.Getenv("GOPATH"),
        ChaincodePath:   "driverFabricDemo/chaincode",
        OrgAdmin:        "Admin",
        OrgName:         "travle",
        ConfigFile:      "conf/config.yaml",
        UserName: "User1",
    }
    // 实例化SDK  创建通道, 将组织节点加入到通道
    fs.Initialization()
}
```

使用 pkg/fabsdk/fabsdk.go中的New()方法进行实例化

```kotlin
sdk, err := fabsdk.New(config.FromFile(this.ConfigFile))
```

- 5.2   根据实例创建资源管理客户端

```kotlin
resCliProvider := sdk.Context(fabsdk.WithUser(this.OrgAdmin),fabsdk.WithOrg(this.OrgName))
resClient, err := resmgmt.New(resCliProvider)
```

- 5.3  创建channel. 得到一个操作链代码的客户端
   使用pkg/client/resmgmt/resmgmt.go文件中的方法

```kotlin
// 先创建一个 创建channel的请求
chanReq := resmgmt.SaveChannelRequest{
        ChannelID: this.ChannelID,
        ChannelConfigPath: this.ChannelConfigPath,
}
// 利用 resClient 创建 channel
chanRsp, err := resClient.SaveChannel(chanReq)
```

- 5.4 把组织添加到 channel 中, 一般制定一些重试的策略,和指定 orderer节点的网络位置

```kotlin
err = resClient.JoinChannel(
        this.ChannelID,
        resmgmt.WithRetry(retry.DefaultResMgmtOpts),
        resmgmt.WithOrdererEndpoint(this.OrdererName),
)
```

### 六. 安装链码,实例化链码

- 6.1 在工程中创建一个chaincode文件夹,用来存在链码,写一个最简单的链码,实现写入数据和读取数据的功能

```go
package main

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
    "fmt"
)

type ChaincodeXQ struct {

}

func (this * ChaincodeXQ)Init(stub shim.ChaincodeStubInterface) pb.Response  {

    return shim.Success([]byte("初始化成功"))
}

func (this * ChaincodeXQ)Invoke(stub shim.ChaincodeStubInterface) pb.Response  {

    // 获取方法和参数
    fname, params := stub.GetFunctionAndParameters()
    if fname == "set" { // 插入操作
        return set(stub,params[0],params[1])
    }else if fname == "get" { //获取操作

        return get(stub,params[0])
    }
    return shim.Error("Invoke 操作失败!!")
}

func set(stub shim.ChaincodeStubInterface, key string, value string) pb.Response {

    err := stub.PutState(key,[]byte(value))
    if err != nil {
        return shim.Error("PutState 操作失败!!")
    }
    return shim.Success([]byte("PutState 操作成功!!"))
}

func get(stub shim.ChaincodeStubInterface, key string) pb.Response {

    data, err := stub.GetState(key)
    if err != nil {
        return shim.Error("GetState 操作失败!!")
    }
    // If the key does not exist in the state database, (nil, nil) is returned.
    if data == nil {
        return shim.Error("GetState 操作失败!!, data == nil")
    }
    return shim.Success(data)
}

func main()  {

    err := shim.Start(new(ChaincodeXQ))
    if err != nil {
        fmt.Println("开启链码失败,err:",err)
        return
    }
    fmt.Println("开启链码成功")
}
```

- 6.2 安装链码
   安装链码之前需要创建请求

```kotlin
InstallCCRequest := resmgmt.InstallCCRequest{
        Name:this.ChainCodeID, // 链码名称
        Path:this.ChaincodePath, //链码在工程中的路径
        Version:"0",
        Package:ccp,
}
```

创建请求之前,需要使用 gopackager.NewCCPackage 方法生成一个resource.CCPackage 对象,传递两个参数,一个是链码的路径(相对于工程的路径), 一个是GOPATH的路径.

```kotlin
ccp, err := gopackager.NewCCPackage(this.ChaincodePath,this.ChaincodeGoPath)
```

安装链码,使用pkg/client/resmgmt/resmgmt.go文件中的方法

```kotlin
req2Arr, err := this.ResMgmtCli.InstallCC(InstallCCRequest)
```

- 6.3 实例化链码
   实例化链码之前需要创建请求

```kotlin
rsp1 := resmgmt.InstantiateCCRequest{
        Name:this.ChainCodeID,// 链码名称
        Path:this.ChaincodeGoPath,//链码在工程中的路径
        Version:"0",
        Args:nil,
        Policy:ccpolity,
}
```

创建请求之前,需要生成一个*cb.SignaturePolicyEnvelope类型的对象,使用 **third_party/github.com/hyperledger/fabric/common/cauthdsl/cauthdsl_builder.go**  文件中的方法即可,提供了好几个方法, 使用任意一个即可.这里使用  **SignedByAnyMember** 方法: 需要传入所属组织ID

```cpp
ccpolity := cauthdsl.SignedByAnyMember([]string{this.OrgID})
```

实例化链码

```kotlin
txID,err := this.ResMgmtCli.InstantiateCC(this.ChannelID,rsp1)
```

- 6.4 再创建一个通道的客户端,后面的调用链码需要使用这个客户端
   使用的是 **pkg/client/channel/chclient.go** 中的方法

```kotlin
// 创建上下文
clientContext := this.Sdk.ChannelContext(this.ChannelID, fabsdk.WithUser(this.UserName))
// 创建channel客户端
channelclient, err := channel.New(clientContext)
```

### 七. 调用链码

使用 **pkg/client/channel/chclient.go** 中的 **Execute()**方法,来进行数据写入的操作:
 `rsp, err := model.Channelclient.Execute(req)`
 写入之前,要创建请求:

```go
req := channel.Request{
        ChaincodeID:model.ChainCodeID, // 链码名称
        Fcn:args[0], // 方法名:  get
        Args:tempArgs, // 传递的参数,是一个二维字符切片  [[49] [49 49 49 49]]  
}
```

tempArgs是要传给链码的参数,可以做下封装,就不受参数个数的限制了

```go
var tempArgs [][]byte
    for i := 1; i < len(args); i++ {
        tempArgs = append(tempArgs, []byte(args[i]))
}
```

使用 **pkg/client/channel/chclient.go** 中的 **Query()**方法,来进行数据查询的操作: 查询之前,同样需要创建请求.

```go
req := channel.Request{
        ChaincodeID: model.ChainCodeID, 
        Fcn: args[0], 
        Args: tempArgs,
    }
rsp, err := model.Channelclient.Query(req)
```

### 遇到的问题

> 1  链码路径错误

```bash
Chaincode status Code: (500) UNKNOWN. Description: error starting container: error starting container: Failed to generate platform-specific docker build: Error returned from build: 1 "can't load package: package chaincode: cannot find package "chaincode" in any of:
        /opt/go/src/chaincode (from $GOROOT)
        /chaincode/input/src/chaincode (from $GOPATH)
        /opt/gopath/src/chaincode
```

链码在工程中的路径应该是 **工程名/chaincode文件夹**
 比如:
 `driverFabricDemo/chaincode`
 而不应该省略掉工程名这样写:`chaincode`

> 2 go版本的问题



```go
Failed to build the application: # github.com/cloudflare/cfssl/csr
../github.com/cloudflare/cfssl/csr/csr.go:272:26: cert.URIs undefined (type *x509.Certificate has no field or method URIs)
../github.com/cloudflare/cfssl/csr/csr.go:387:7: tpl.URIs undefined (type x509.CertificateRequest has no field or method URIs)
```

错误原因:cert.URIs 和 tpl.URIs 这两个字段没有被定义.
 进入tpl对象中,`/usr/local/go/src/crypto/x509/x509.go` 是个结构体,并没有发现`URIs`字段

对go版本进行升级,从1.9.3升级到1.11.3, 再次进入 `/usr/local/go/src/crypto/x509/x509.go` 文件中,查看结构体内容:



```cpp
 // Subject Alternate Name values.
   DNSNames       []string
   EmailAddresses []string
   IPAddresses    []net.IP
   URIs           []*url.URL
```

> 3 写入数据到区块链时错误



```bash
Transaction processing for endorser [localhost:7051]: Chaincode status Code: (500) UNKNOWN. 
Description: Received unknow function invocatio
```

在执行sdk的Excute()方法时报错.
 方法不存在,一般是由于链码的Invoke方法中的方法名和Excute()方法传入的方法名不一样.
 但是可以肯定的是,链码的Invoke方法中的方法名和,项目中执行Excute()方法时传入的方法名是完全一样的! 但是很奇怪了,为什么会出现这个错误呢? 使用`docker rmi` 删除掉`dev-peerx.travle.xq.com` 的镜像,再重新运行即可.

> 4 Cannot use str (type *cb.SignaturePolicyEnvelope) as type *common.SignaturePolicyEnvelope

在创建实例化链码请求的时候



```kotlin
ccpolity := cauthdsl.SignedByAnyMember([]string{this.OrgID})
rsp1 := resmgmt.InstantiateCCRequest{
   Name:this.ChainCodeID,
   Path:this.ChaincodeGoPath,
   Version:"1.0",
   Args:nil,
   Policy:ccpolity,
}
```

总是提示
 `Cannot use str (type *cb.SignaturePolicyEnvelope) as type *common.SignaturePolicyEnvelope less... (⌘F1) Inspection info: Reports composite literals with incompatible types and values`
 明明是相同的类型,却总是报错,应该是IDE的问题.把vendor文件夹删除后,就不会有提示了. 再使用vendor对工程进行init 和 add +external 就好了!!

> 5 CONNECTION_ Description: dialing connection timed out

出现这个错误,一般都是配置文件哪个地方写错了,需要细心检查



作者：小李飞刀无情剑
链接：https://www.jianshu.com/p/26cf5e4a9de9
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



# Fabric-sdk-go案例-访问yaml配置并创建SDK实例

Fabric-sdk-go案例
准备工作
创建案例目录
编写案例源程序demo1.go
运行结果
准备工作
创建案例目录

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200514093515598.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200514093651623.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE4ODA3MDQz,size_16,color_FFFFFF,t_70)

config_e2e.yaml是fabric-sdk-go本身提供的一个配置文件，搜索得到，复制到案例目录。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200514093935952.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE4ODA3MDQz,size_16,color_FFFFFF,t_70)

编写案例源程序demo1.go

第18行，从当前目录读配置文件，即config_e2e.yaml
第19行，configOpt是1个函数，调用它得到1个切片backend
第20行，在切片backend中查询version，也就是config_e2e.yaml里的version

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200514094311816.png)

第22行，在切片backend中查询client.cryptoconfig.path

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200514094333737.png)

第24行，使用config_e2e.yaml创建1个fabric sdk实例
第30行，关闭SDK

运行结果

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200514094605823.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE4ODA3MDQz,size_16,color_FFFFFF,t_70)

————————————————
版权声明：本文为CSDN博主「看聊效」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_18807043/article/details/106112918



# 简化Fabric-sdk-go的测试案例

Fabric-sdk-go文档表明的测试案例有：make unit-test、make integration-test。这些测试案例检查、配置环境耗费很长时间，而且每次测试都重复这些检查。对这些测试案例进行简化，加快测试进度。
1、简化unit-test
在Makefile中找到unit-test, 用#取消掉unit-test的依赖项。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103247303.png)

执行make crypto-gen，根据cryptogen.yaml生成组织的数字证书。该文件的位置如下图。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103322140.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103336495.png)

执行make channel-config-gen，产生创世纪块、通道等环境。用到的配置文件是configtx.yaml。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103358158.png)

执行make unit-test，略过诸多检查、配置，直接调用unit.sh脚本。

![在这里插入图片描述](https://img-blog.csdnimg.cn/2020042110341772.png)

2、简化integration-test
在Makefile中找到integration-test, 用#取消掉integration-test的依赖项。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103436891.png)

integration-test可以调用多个版本的集成测试，例如integration-tests-stable。在Makefile中找到integration-tests-stable, 用#取消掉integration-tests-stable的依赖项。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103455550.png)

如果没有执行make crypto-gen和make channel-config-gen，则执行并产生环境和配置。
执行make integration-test，这时会略过很多检查、下载，进度加快。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103513932.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200421103524718.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE4ODA3MDQz,size_16,color_FFFFFF,t_70)

————————————————
版权声明：本文为CSDN博主「看聊效」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_18807043/article/details/105652526



# 编译Fabric V2.0.0

看聊效 2020-04-14 14:53:31  27  收藏
展开
下载Fabric源码之后，进入Fabric目录，检出代码git checkout v2.0.0

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200414144832774.png)

分别执行make clean; make release

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200414144910451.png)

编译完毕之后，将在release目录下产生Fabric的二进制文件，共计7个

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200414144957551.png)

以peer为例，main.go的位置是cmd/peer

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200414145058173.png)

peer的maingo很简短，与其下一级命令选项对应

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200414145144554.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzE4ODA3MDQz,size_16,color_FFFFFF,t_70)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200414145155260.png)

要修改peer的代码，找到对应的import项。
————————————————
版权声明：本文为CSDN博主「看聊效」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_18807043/article/details/105511971