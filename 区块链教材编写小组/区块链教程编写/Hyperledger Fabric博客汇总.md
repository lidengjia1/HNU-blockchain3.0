# [Hyperledger Fabric博客汇总](https://www.cnblogs.com/cbkj-xd/p/12170903.html)

# Hyperledger Fabric

1. [Hyperledger Fabric1.4环境搭建](https://ifican.top/2019/11/23/blog/fabric/Fabric环境搭建/)
2. [深入解析Hyperledger Fabric搭建的全过程](https://ifican.top/2019/11/23/blog/fabric/深入解析Fabric搭建的全过程/)
3. [Hyperledger Fabric动态添加组织到网络中](https://ifican.top/2019/12/08/blog/fabric/Hyperledger_Fabric动态添加组织到网络中/)
4. [Hyperledger Fabric多机部署](https://ifican.top/2019/11/23/blog/fabric/Fabric1.4多机部署/)
5. [Hyperledger Fabric动态配置Raft节点](https://ifican.top/2019/12/31/blog/fabric/动态配置Raft节点/)
6. [Hyperledger Fabric 开启TLS调用Java SDK](https://ifican.top/2019/12/28/blog/fabric/TLS_SDK调用/)
7. [Hyperledger Fabric 最简单的方式测试你的链码](https://ifican.top/2019/11/27/blog/fabric/链码测试/)
8. [Hyperledger Fabric私有数据](https://ifican.top/2019/12/04/blog/fabric/私有数据/)
9. [Hyperledger Fabric使用硬件安全模块(HSM)](https://ifican.top/2019/12/24/blog/fabric/使用硬件安全模块/)
10. [Hyperledger Fabric链码作为外部服务](https://ifican.top/2019/12/27/blog/fabric/链码作为外部服务/)
11. [Hyperledger Fabric外部链码构建与运行](https://ifican.top/2019/12/24/blog/fabric/外部链码构建和运行/)
12. [Hyperledger Fabric-CA](https://ifican.top/2019/12/08/blog/fabric/Hyperledger_Fabric_CA/)
13. [Hyperledger Fabric手动生成CA证书搭建Fabric网络](https://ifican.top/2019/12/08/blog/fabric/Hyperledger_Fabric手动生成CA证书搭建Fabric网络/)

# Raft

1. [Raft算法之Leader选举](https://ifican.top/2020/01/04/blog/consensus/raft-election/)
2. [Raft算法之日志复制](https://ifican.top/2020/01/05/blog/consensus/raft-log/)
3. [Raft算法之成员关系变化](https://ifican.top/2020/01/06/blog/consensus/raft-relationship/)
4. [Raft算法之日志压缩](https://ifican.top/2020/01/07/blog/consensus/raft-snapshot/)

转载请注明作者与出处：https://www.cnblogs.com/cbkj-xd/ 个人网站主页：https://ifican.top



# Hyperledger Fabric动态配置Raft节点

最近看官方文档发现新的共识算法etcdRaft允许动态添加或删除排序节点，所以也花了一天时间操作了以下，写篇文章把整个过程记录一下。
初始网络本文设置了4个Orderer节点，1个Peer节点(用于更新配置文件以及测试用),然后动态添加第五个Orderer节点。
本文分成两个部分:

1. 第一部分是手动通过Fabric-CA生成每一个节点的证书文件
2. 第二部分是更新Fabric网络配置添加新的Orderer节点。

本文基于**Fabric v2.0.0-beta**版本。版本号只要高于1.4.1就行

## 1 搭建定制化的Fabric网络

前提条件是成功跑起来Fabric的示例网络，可以看这里->[Hyperledger Fabric环境搭建](https://ifican.top/2019/11/23/blog/fabric/Fabric环境搭建/)

首先在`$GOPATH`下(本文路径地址为`$GOPATH/src/github.com/hyperledger/fab`)建立如下几个文件夹用于之后的操作:

```
.  # 这里是根目录fab
├── ca    # 用于生成CA证书的ca配置文件的文件夹
│   ├── org1
│   │   └── fabric-ca-server-config.yaml
│   └── server
│       └── fabric-ca-server-config.yaml
├── channel-artifacts    #用于保存创世区块以及通道配置文件
├── configtx.yaml      #配置文件：用于生成创世区块以及通道配置文件
├── crypto-config     #存储生成的证书文件
├── docker      # Fabric网络节点通过Docker启动，用于启动节点的Docker文件
│   ├── base.yaml
│   ├── docker-compose-addOrderer5.yaml
│   ├── docker-compose-ca.yaml
│   ├── docker-compose-orderers.yaml
│   └── docker-compose-peer.yaml
└── store    #存储区块等信息
```

**以下所有操作默认都在根目录文件夹内！**

### 1.1CA配置文件

直接在这里贴出来:`org1/fabric-ca-server-config.yaml`:

```
version: 1.2.0

# Server's listening port (default: 7054)
port: 7054

# Enables debug logging (default: false)
debug: false

crlsizelimit: 512000

tls:
  # Enable TLS (default: false)
  enabled: true
  certfile:
  keyfile:
  clientauth:
    type: noclientcert
    certfiles:

ca:
  # Name of this CA
  name: Org1CA
  keyfile:
  certfile:
  chainfile:

crl:
  expiry: 24h

registry:

  maxenrollments: -1

  identities:
     - name: admin
       pass: adminpw
       type: client
       affiliation: ""
       attrs:
          hf.Registrar.Roles: "*"
          hf.Registrar.DelegateRoles: "*"
          hf.Revoker: true
          hf.IntermediateCA: true
          hf.GenCRL: true
          hf.Registrar.Attributes: "*"
          hf.AffiliationMgr: true

db:
  type: sqlite3
  datasource: fabric-ca-server.db
  tls:
      enabled: false
      certfiles:
      client:
        certfile:
        keyfile:

ldap:

   enabled: false
   url: ldap://<adminDN>:<adminPassword>@<host>:<port>/<base>
   tls:
      certfiles:
      client:
         certfile:
         keyfile:
   attribute:
      names: ['uid','member']
      converters:
         - name:
           value:
      maps:
         groups:
            - name:
              value:

affiliations:
   org1:
      - department1
      - department2
   org2:
      - department1

signing:
    default:
      usage:
        - digital signature
      expiry: 8760h
    profiles:
      ca:
         usage:
           - cert sign
           - crl sign
         expiry: 43800h
         caconstraint:
           isca: true
           maxpathlen: 0
      tls:
         usage:
            - signing
            - key encipherment
            - server auth
            - client auth
            - key agreement
         expiry: 8760h

csr:
   cn: ca.org1.example.com
   names:
      - C: US
        ST: "North Carolina"
        L: "Durham"
        O: org1.example.com
        OU:
   hosts:
     - localhost
     - org1.example.com
   ca:
      expiry: 131400h
      pathlength: 1

bccsp:
    default: SW
    sw:
        hash: SHA2
        security: 256
        filekeystore:
            keystore: msp/keystore

cacount:

cafiles:

intermediate:
  parentserver:
    url:
    caname:

  enrollment:
    hosts:
    profile:
    label:

  tls:
    certfiles:
    client:
      certfile:
      keyfile:
      
```

以及`server/fabric-ca-server-config.yaml:`:

```
# Version of config file
version: 1.2.0
# Server's listening port (default: 7054)
port: 7054
# Enables debug logging (default: false)
debug: false
# Size limit of an acceptable CRL in bytes (default: 512000)
crlsizelimit: 512000
tls:
  # Enable TLS (default: false)
  enabled: true
  # TLS for the server's listening port
  certfile:
  keyfile:
  clientauth:
    type: noclientcert
    certfiles:

ca:
  # Name of this CA
  name: OrdererCA
  keyfile:
  certfile:
  chainfile:

crl:
  expiry: 24h

registry:
  maxenrollments: -1

  identities:
     - name: admin
       pass: adminpw
       type: client
       affiliation: ""
       attrs:
          hf.Registrar.Roles: "*"
          hf.Registrar.DelegateRoles: "*"
          hf.Revoker: true
          hf.IntermediateCA: true
          hf.GenCRL: true
          hf.Registrar.Attributes: "*"
          hf.AffiliationMgr: true

db:
  type: sqlite3
  datasource: fabric-ca-server.db
  tls:
      enabled: false
      certfiles:
      client:
        certfile:
        keyfile:

ldap:
   enabled: false
   url: ldap://<adminDN>:<adminPassword>@<host>:<port>/<base>
   tls:
      certfiles:
      client:
         certfile:
         keyfile:
   attribute:
      names: ['uid','member']
      converters:
         - name:
           value:
      maps:
         groups:
            - name:
              value:

affiliations:
   org1:
      - department1
      - department2
   org2:
      - department1

signing:
    default:
      usage:
        - digital signature
      expiry: 8760h
    profiles:
      ca:
         usage:
           - cert sign
           - crl sign
         expiry: 43800h
         caconstraint:
           isca: true
           maxpathlen: 0
      tls:
         usage:
            - signing
            - key encipherment
            - server auth
            - client auth
            - key agreement
         expiry: 8760h

csr:
   cn: ca.example.com
   names:
      - C: US
        ST: "New York"
        L: "New York"
        O: example.com
        OU:
   hosts:
     - localhost
     - example.com
   ca:
      expiry: 131400h
      pathlength: 1

bccsp:
    default: SW
    sw:
        hash: SHA2
        security: 256
        filekeystore:
            keystore: msp/keystore

cacount:
cafiles:

intermediate:
  parentserver:
    url:
    caname:

  enrollment:
    hosts:
    profile:
    label:

  tls:
    certfiles:
    client:
      certfile:
      keyfile:
```

`docker-compose-ca.yaml`文件:

```
version: '2'

services:
  ca:
    image: hyperledger/fabric-ca:1.4.4
    environment:
      - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
      - FABRIC_CA_SERVER_CA_NAME=ca-orderer
      - FABRIC_CA_SERVER_TLS_ENABLED=true
      - FABRIC_CA_SERVER_PORT=9054
    ports:
      - "9054:9054"
    command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
    volumes:
      - ../ca/server:/etc/hyperledger/fabric-ca-server
    container_name: ca_orderer

  ca0:
    image: hyperledger/fabric-ca:1.4.4
    environment:
      - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
      - FABRIC_CA_SERVER_CA_NAME=ca-org1
      - FABRIC_CA_SERVER_TLS_ENABLED=true
      - FABRIC_CA_SERVER_PORT=7054
    ports:
      - "7054:7054"
    command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
    volumes:
      - ../ca/org1:/etc/hyperledger/fabric-ca-server
    container_name: ca_org1
```

将以上三个文件保存到指定的路径，然后使用以下命令启动CA服务器：

```
docker-compose -f docker/docker-compose-ca.yaml up -d
```

服务器会自动读取上面的两个配置文件，并初始化CA服务器。
当然，服务器配置文件将自动生成在`ca/server/`子文件夹内，其中最主要使用到的是`tls-cert.pem`文件。

### 1.2 注册Orderer节点

首先配置环境变量并登陆管理员账号:

```
#创建存储Order节点证书的子文件夹。
mkdir -p crypto-config/orderOrganization/example.com
export FABRIC_CA_CLIENT_HOME=${PWD}/crypto-config/orderOrganization/example.com
fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

生成节点类型分类配置文件(不知道这个文件应该称作什么，暂且使用这个名字称呼好了):

```
  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml
```

之后注册网络中初始的4个Orderer节点:

```
fabric-ca-client register -u https://admin:adminpw@localhost:9054 --caname ca-orderer --id.name orderer1 --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client register -u https://admin:adminpw@localhost:9054 --caname ca-orderer --id.name orderer2 --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client register -u https://admin:adminpw@localhost:9054 --caname ca-orderer --id.name orderer3 --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client register -u https://admin:adminpw@localhost:9054 --caname ca-orderer --id.name orderer4 --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

注册`Admin`节点:

```
fabric-ca-client register -u https://admin:adminpw@localhost:9054 --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

### 1.3 获取Orderer证书文件

为刚刚创建的几个用户创建各自的文件夹用于存储证书文件:

```
mkdir -p crypto-config/orderOrganization/example.com/orderers
mkdir -p crypto-config/orderOrganization/example.com/orderers/orderer1.example.com
mkdir -p crypto-config/orderOrganization/example.com/orderers/orderer2.example.com
mkdir -p crypto-config/orderOrganization/example.com/orderers/orderer3.example.com
mkdir -p crypto-config/orderOrganization/example.com/orderers/orderer4.example.com
```

接下来获取每一个Orderer节点的`MSP`证书文件:

```
fabric-ca-client enroll -u https://orderer1:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/msp --csr.hosts orderer1.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client enroll -u https://orderer2:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/msp --csr.hosts orderer2.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client enroll -u https://orderer3:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/msp --csr.hosts orderer3.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client enroll -u https://orderer4:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/msp --csr.hosts orderer4.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

还有每一个节点的`TLS`证书:

```
fabric-ca-client enroll -u https://orderer1:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls --enrollment.profile tls --csr.hosts orderer1.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client enroll -u https://orderer2:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls --enrollment.profile tls --csr.hosts orderer2.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client enroll -u https://orderer3:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls --enrollment.profile tls --csr.hosts orderer3.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
fabric-ca-client enroll -u https://orderer4:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls --enrollment.profile tls --csr.hosts orderer4.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

将之前生成的节点类型分类配置文件拷贝到每一个节点的`MSP`文件夹:

```
cp ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/msp/config.yaml
cp ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/msp/config.yaml
cp ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/msp/config.yaml
cp ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/msp/config.yaml
```

然后为每一个节点的`TLS`证书以及秘钥文件修改名字，方便之后的使用:

```
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/ca.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/signcerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/server.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/keystore/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/server.key

cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/ca.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/signcerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/server.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/keystore/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/server.key

cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/ca.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/signcerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/server.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/keystore/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/server.key

cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/ca.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/signcerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/server.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/keystore/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/server.key
```

然后在`MSP`文件夹内创建`tlscacerts`文件夹，并将`TLS`文件拷贝过去:

```
mkdir ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/msp/tlscacerts
mkdir ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/msp/tlscacerts
mkdir ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/msp/tlscacerts
mkdir ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/msp/tlscacerts

cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

复制TLS根证书:

```
mkdir -p ${PWD}/crypto-config/orderOrganization/example.com/msp/tlscacerts
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

最后是`Admin`节点的证书文件:

```
#首先也是创建文件夹
mkdir -p crypto-config/orderOrganization/example.com/users
mkdir -p crypto-config/orderOrganization/example.com/users/Admin@example.com
#获取证书文件
fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/users/Admin@example.com/msp --tls.certfiles ${PWD}/ca/server/tls-cert.pem
cp ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml ${PWD}/crypto-config/orderOrganization/example.com/users/Admin@example.com/msp/config.yaml
```

到这里Orderer节点证书已经生成完毕(可以根据实际需要修改Orderer节点数量，最少不能低于3个)，接下来是网络中唯一的`peer`节点的配置文件生成:

### 1.4 注册Peer节点

和上面步骤相同，首先创建子文件夹用于存储证书文件:

```
mkdir -p crypto-config/peerOrganizations/org1.example.com/
```

配置环境变量并登陆管理员身份:

```
export FABRIC_CA_CLIENT_HOME=${PWD}/crypto-config/peerOrganizations/org1.example.com/
fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
```

生成节点类型分类配置文件:

```
echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/crypto-config/peerOrganizations/org1.example.com/msp/config.yaml
```

虽然网络中只有一个peer节点，但是我们需要注册三个用户:`peer0,user1,org1admin`，其中第一个是必需的，第二个是用于测试的，第三个为`Admin`用户，安装和实例化链码需要`Admin`用户的证书:

```
fabric-ca-client register -u https://admin:adminpw@localhost:7054 --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --id.attrs '"hf.Registrar.Roles=peer"' --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
fabric-ca-client register -u https://admin:adminpw@localhost:7054 --caname ca-org1 --id.name user1 --id.secret user1pw --id.type client --id.attrs '"hf.Registrar.Roles=client"' --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
fabric-ca-client register -u https://admin:adminpw@localhost:7054 --caname ca-org1 --id.name org1admin --id.secret org1adminpw --id.type admin --id.attrs '"hf.Registrar.Roles=admin"' --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
```

### 1.5 获取Peer节点证书文件

节点注册完毕，获取他们的证书文件:
创建子文件夹:

```
mkdir -p crypto-config/peerOrganizations/org1.example.com/peers
mkdir -p crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.co
```

获取证书文件:

```
#MSP文件
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp --csr.hosts peer0.org1.example.com --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
#TLS证书
fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls --enrollment.profile tls --csr.hosts peer0.org1.example.com --csr.hosts localhost --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
```

拷贝节点分类配置文件:

```
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/msp/config.yaml ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp/config.yaml
```

修改证书以及秘钥文件，方便之后使用:

```
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/tlscacerts/* ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/signcerts/* ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/keystore/* ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key
```

将TLS相关证书复制一份:

```
mkdir ${PWD}/crypto-config/peerOrganizations/org1.example.com/msp/tlscacerts
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/tlscacerts/* ${PWD}/crypto-config/peerOrganizations/org1.example.com/msp/tlscacerts/ca.crt

mkdir ${PWD}/crypto-config/peerOrganizations/org1.example.com/tlsca
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/tlscacerts/* ${PWD}/crypto-config/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem

mkdir ${PWD}/crypto-config/peerOrganizations/org1.example.com/ca
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp/cacerts/* ${PWD}/crypto-config/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem
```

获取`user`与`Admin`用户证书文件:

```
#创建子文件夹
mkdir -p crypto-config/peerOrganizations/org1.example.com/users
mkdir -p crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com
mkdir -p crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com
#获取证书文件
fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-org1 -M ${PWD}/crypto-config/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-org1 -M ${PWD}/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp --tls.certfiles ${PWD}/ca/org1/tls-cert.pem
cp ${PWD}/crypto-config/peerOrganizations/org1.example.com/msp/config.yaml ${PWD}/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/config.yaml
```

### 1.6 启动网络之前的准备

到这里我们已经生成了所有需要的证书文件，接下来是生成用于启动网络的创世区块,生成创世区块需要一个文件`configtx.yaml`,直接复制过来:

```
Organizations:
    - &OrdererOrg
        Name: OrdererOrg
        ID: OrdererMSP
        MSPDir: ./crypto-config/orderOrganization/example.com/msp   #这里路径需要对应！！！
        Policies:
            Readers:
                Type: Signature
                Rule: "OR('OrdererMSP.member')"
            Writers:
                Type: Signature
                Rule: "OR('OrdererMSP.member')"
            Admins:
                Type: Signature
                Rule: "OR('OrdererMSP.admin')"

    - &Org1  #如果需要更多组织节点，可以按照该模板在下面添加
        Name: Org1MSP
        ID: Org1MSP
        MSPDir: ./crypto-config/peerOrganizations/org1.example.com/msp  #这里路径需要对应！！！
        Policies:
            Readers:
                Type: Signature
                Rule: "OR('Org1MSP.admin', 'Org1MSP.peer', 'Org1MSP.client')"
            Writers:
                Type: Signature
                Rule: "OR('Org1MSP.admin', 'Org1MSP.client')"
            Admins:
                Type: Signature
                Rule: "OR('Org1MSP.admin')"
            Endorsement:
                Type: Signature
                Rule: "OR('Org1MSP.peer')"
        AnchorPeers:
              Port: 7051

Capabilities:
    Channel: &ChannelCapabilities
        V2_0: true

    Orderer: &OrdererCapabilities
        V2_0: true

    Application: &ApplicationCapabilities
        V2_0: true

Application: &ApplicationDefaults
    Organizations:
    Policies:
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
        LifecycleEndorsement:
            Type: ImplicitMeta
            Rule: "MAJORITY Endorsement"
        Endorsement:
            Type: ImplicitMeta
            Rule: "MAJORITY Endorsement"
    Capabilities:
        <<: *ApplicationCapabilities

Orderer: &OrdererDefaults
    OrdererType: etcdraft
    
    Addresses:
        - orderer1.example.com:7050
    BatchTimeout: 2s
    BatchSize:
        MaxMessageCount: 10
        AbsoluteMaxBytes: 99 MB
        PreferredMaxBytes: 512 KB
    Organizations:
    Policies:
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
        BlockValidation:
            Type: ImplicitMeta
            Rule: "ANY Writers"
Channel: &ChannelDefaults
    Policies:
        # Who may invoke the 'Deliver' API
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        # Who may invoke the 'Broadcast' API
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        # By default, who may modify elements at this config level
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
    Capabilities:
        <<: *ChannelCapabilities

Profiles:
    
    TwoOrgsChannel:   #用于生成通道配置文件
        Consortium: SampleConsortium
        <<: *ChannelDefaults
        Application:
            <<: *ApplicationDefaults
            Organizations:
                - *Org1
            Capabilities:
                <<: *ApplicationCapabilities

    SampleMultiNodeEtcdRaft:   #用于生成系统通道创世区块
        <<: *ChannelDefaults
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            OrdererType: etcdraft   #指定使用etcdraft共识算法
            EtcdRaft:
                Consenters:
                - Host: orderer1.example.com
                  Port: 7050
                  ClientTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/server.crt
                  ServerTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/server.crt
                - Host: orderer2.example.com
                  Port: 8050
                  ClientTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/server.crt
                  ServerTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/server.crt
                - Host: orderer3.example.com
                  Port: 9050
                  ClientTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/server.crt
                  ServerTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/server.crt
                - Host: orderer4.example.com
                  Port: 10050
                  ClientTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/server.crt
                  ServerTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/server.crt
#                    - Host: orderer5.example.com
#                      Port: 11050
#                      ClientTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/server.crt
#                      ServerTLSCert: ./crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/server.crt
            Addresses:
                - orderer1.example.com:7050
                - orderer2.example.com:8050
                - orderer3.example.com:9050
                - orderer4.example.com:10050
#                - orderer5.example.com:11050
         
            Organizations:
            - *OrdererOrg
            Capabilities:
                <<: *OrdererCapabilities
        Application:
            <<: *ApplicationDefaults
            Organizations:
            - <<: *OrdererOrg
        Consortiums:
            SampleConsortium:
                Organizations:
                - *Org1
```

将该文件保存到指定位置，接下来生成创世区块:

```
export FABRIC_CFG_PATH=$PWD
configtxgen -profile SampleMultiNodeEtcdRaft -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
# 生成通道配置文件
export CHANNEL_NAME=mychannel
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/mychannel.tx -channelID $CHANNEL_NAME
```

### 1.7 启动网络

首先写包含所有节点的Docker文件,这里直接贴出来:

```
version: '2'

services:
  orderer-base:
    image: hyperledger/fabric-orderer:2.0.0-beta
    environment:
      - FABRIC_LOGGING_SPEC=INFO
#      - FABRIC_LOGGING_SPEC=DEBUG
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_BOOTSTRAPMETHOD=file
      - ORDERER_GENERAL_BOOTSTRAPFILE=/var/hyperledger/orderer/orderer.genesis.block
      - ORDERER_GENERAL_LOCALMSPID=OrdererMSP
      - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
      # enabled TLS
      - ORDERER_GENERAL_TLS_ENABLED=true
      - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
      - ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_CLUSTER_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
```

然后是Orderer节点的Docker文件:

```
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

version: '2'

volumes:
  orderer1.example.com:
  orderer2.example.com:
  orderer3.example.com:
  orderer4.example.com:

networks:
  byfn:

services:
  
  orderer1.example.com:
    extends:
      file: base.yaml
      service: orderer-base
    environment:
      - ORDERER_GENERAL_LISTENPORT=7050
    container_name: orderer1.example.com
    networks:
      - byfn
    volumes:
      - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ../crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/msp:/var/hyperledger/orderer/msp
      - ../crypto-config/orderOrganization/example.com/orderers/orderer1.example.com/tls/:/var/hyperledger/orderer/tls
      - ../store/o1:/var/hyperledger/production/orderer
    ports:
      - 7050:7050



  orderer2.example.com:
    extends:
      file: base.yaml
      service: orderer-base
    environment:
      - ORDERER_GENERAL_LISTENPORT=8050
    container_name: orderer2.example.com
    networks:
      - byfn
    volumes:
      - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ../crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/msp:/var/hyperledger/orderer/msp
      - ../crypto-config/orderOrganization/example.com/orderers/orderer2.example.com/tls/:/var/hyperledger/orderer/tls
      - ../store/o2:/var/hyperledger/production/orderer
    ports:
      - 8050:8050
  
  orderer3.example.com:
    extends:
      file: base.yaml
      service: orderer-base
    environment:
      - ORDERER_GENERAL_LISTENPORT=9050
    container_name: orderer3.example.com
    networks:
      - byfn
    volumes:
      - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ../crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/msp:/var/hyperledger/orderer/msp
      - ../crypto-config/orderOrganization/example.com/orderers/orderer3.example.com/tls/:/var/hyperledger/orderer/tls
      - ../store/o3:/var/hyperledger/production/orderer
    ports:
      - 9050:9050
  
  orderer4.example.com:
    extends:
      file: base.yaml
      service: orderer-base
    environment:
      - ORDERER_GENERAL_LISTENPORT=10050
    container_name: orderer4.example.com
    networks:
      - byfn
    volumes:
      - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ../crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/msp:/var/hyperledger/orderer/msp
      - ../crypto-config/orderOrganization/example.com/orderers/orderer4.example.com/tls/:/var/hyperledger/orderer/tls
      - ../store/o4:/var/hyperledger/production/orderer
    ports:
      - 10050:10050
```

最后一个是peer节点的Docker文件：

```
version: '2'

volumes:
  peer0.org1.example.com:

networks:
  byfn:

services:

  peer0.org1.example.com:
    container_name: peer0.org1.example.com
    image: hyperledger/fabric-peer:2.0.0-beta
    environment:
      #Generic peer variables
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      # the following setting starts chaincode containers on the same
      # bridge network as the peers
      # https://docs.docker.com/compose/networking/
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=${COMPOSE_PROJECT_NAME}_byfn
      - FABRIC_LOGGING_SPEC=INFO
      #- FABRIC_LOGGING_SPEC=DEBUG
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
      # Peer specific variabes
      - CORE_PEER_ID=peer0.org1.example.com
      - CORE_PEER_ADDRESS=peer0.org1.example.com:7051
      - CORE_PEER_LISTENADDRESS=0.0.0.0:7051
      - CORE_PEER_CHAINCODEADDRESS=peer0.org1.example.com:7052
      - CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:7052
      - CORE_PEER_GOSSIP_BOOTSTRAP=peer0.org1.example.com:7051
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org1.example.com:7051
      - CORE_PEER_LOCALMSPID=Org1MSP
      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb0:5984
      # The CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME and CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD
      # provide the credentials for ledger to connect to CouchDB.  The username and password must
      # match the username and password set for the associated CouchDB.
      - CORE_LEDGER_STATE_COUCHDBCONFIG_USERNAME=
      - CORE_LEDGER_STATE_COUCHDBCONFIG_PASSWORD=
    volumes:
      - /var/run/:/host/var/run/
      - ../crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp:/etc/hyperledger/fabric/msp
      - ../crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls:/etc/hyperledger/fabric/tls
      - ../store/p1:/var/hyperledger/production
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    depends_on:
      - couchdb0
    ports:
      - 7051:7051
    networks:
      - byfn

  couchdb0:
    container_name: couchdb0
    image: couchdb:2.3
    # Populate the COUCHDB_USER and COUCHDB_PASSWORD to set an admin user and password
    # for CouchDB.  This will prevent CouchDB from operating in an "Admin Party" mode.
    environment:
      - COUCHDB_USER=
      - COUCHDB_PASSWORD=
    # Comment/Uncomment the port mapping if you want to hide/expose the CouchDB service,
    # for example map it to utilize Fauxton User Interface in dev environments.
    ports:
      - "5984:5984"
    networks:
      - byfn

  cli:
    container_name: cli
    image: hyperledger/fabric-tools:2.0.0-beta
    tty: true
    stdin_open: true
    environment:
      - GOPATH=/opt/gopath
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
#      - FABRIC_LOGGING_SPEC=DEBUG
      - FABRIC_LOGGING_SPEC=INFO
      - CORE_PEER_ID=peer0.org1.example.com
      - CORE_PEER_ADDRESS=peer0.org1.example.com:7051
      - CORE_PEER_LOCALMSPID=Org1MSP
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
      - CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: /bin/bash
    volumes:
      - /var/run/:/host/var/run/
      - ./../../chaincode/:/opt/gopath/src/github.com/hyperledger/fabric-samples/chaincode
      - ../crypto-config:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/
      - ../channel-artifacts:/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts
    depends_on:
      - peer0.org1.example.com
    networks:
      - byfn
```

将以上文件保存到指定位置后，使用以下命令直接启动:

```
docker-compose -f docker/docker-compose-orderers.yaml -f docker/docker-compose-peer.yaml up -d
```

启动完成后可以查看每个节点的日志确认节点成功运行:

```
docker logs orderer1.example.com
...
docker logs peer0.org1.example.com
```

如果没有错误的话就可以进行第二部分了，如果出现错误则要回去检查是不是哪里漏掉了。

### 1.8 简单测试

先进行第一部分的测试，看一下创建通道，加入通道是否成功:

```
#进入CLI容器
docker exec -it cli bash
#配置环境变量
export CHANNEL_NAME=mychannel
export ORDERER_CA=${PWD}/crypto/orderOrganization/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export PEER0_ORG1_CA=${PWD}/crypto/peerOrganization/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
export CORE_PEER_MSPCONFIGPATH=${PWD}/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
```

创建通道:

```
peer channel create -o orderer1.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --outputBlock ./channel-artifacts/${CHANNEL_NAME}.block
```

加入通道:

```
peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block
```

如果一切顺利的话，网络就成功搭建起来了，至于链码就不再测试了。
直接到第二部分，动态添加一个Orderer节点。

## 2 动态添加Raft节点

主要步骤如下：

1. 为该节点生成证书文件
2. 获取当前网络的配置文件
3. 将证书文件添加到配置文件中
4. 更新配置文件
5. 启动新的Orderer节点

### 2.1 生成证书文件

#### 2.1.1 注册该节点身份

```
fabric-ca-client register -u https://admin:adminpw@localhost:9054 --caname ca-orderer --id.name orderer5 --id.secret ordererpw --id.type orderer --id.attrs '"hf.Registrar.Roles=orderer"' --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

为该节点创建存储证书的文件夹:

```
mkdir -p crypto-config/orderOrganization/example.com/orderers/orderer5.example.com
```

#### 2.1.2 获取该节点证书

```
#MSP
fabric-ca-client enroll -u https://orderer5:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/msp --csr.hosts orderer5.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
#TLS
fabric-ca-client enroll -u https://orderer5:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls --enrollment.profile tls --csr.hosts orderer5.example.com --tls.certfiles ${PWD}/ca/server/tls-cert.pem
```

复制节点分类配置文件:

```
cp ${PWD}/crypto-config/orderOrganization/example.com/msp/config.yaml ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/msp/config.yaml
```

修改证书与秘钥文件名称:

```
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/ca.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/signcerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/server.crt
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/keystore/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/server.key
```

创建文件夹并拷贝TLS证书文件:

```
mkdir ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/msp/tlscacerts
cp ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/tlscacerts/* ${PWD}/crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

### 2.2 获取网络配置文件

将节点添加进网络，首先需要将该节点添加到系统通道内，所以先获取系统通道的配置文件:
进入`cli`容器:

```
docker exec -it cli bash
```

配置环境变量，需要使用Orderer节点的身份信息:

```
export CORE_PEER_LOCALMSPID="OrdererMSP"
export ORDERER_CA=${PWD}/crypto/orderOrganization/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/crypto/ordererOrganization/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/orderOrganization/example.com/users/Admin@example.com/msp
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
```

获取系统通道配置文件:

```
peer channel fetch config channel-artifacts/config_block.pb -o orderer1.example.com:7050 -c byfn-sys-channel --tls --cafile $ORDERER_CA
```

解码该配置文件:

```
configtxlator proto_decode --input channel-artifacts/config_block.pb --type common.Block | jq .data.data[0].payload.data.config > channel-artifacts/config.json
```

### 2.3将证书文件添加到配置文件中

退出容器，可以在`channel-artifacts`文件内找到`config.json`文件。将该文件复制一份并在`channel-artifacts`文件夹下保存为`update_config.json`,使用编辑工具打开，并搜索`.example.com`字段如下：
字段一部分：

```
  {
    "client_tls_cert": "一连串的字符串",
    "host": "orderer1.example.com",
    "port": 7050,
    "server_tls_cert": "一连串的字符串"
  }
```

以及匹配到的第二部分的字段:

```
      "OrdererAddresses": {
        "mod_policy": "/Channel/Orderer/Admins",
        "value": {
          "addresses": [
            "orderer1.example.com:7050",
            "orderer2.example.com:8050",
            "orderer3.example.com:9050",
            "orderer4.example.com:10050"
          ]
        },
        "version": "0"
    }
```

在字段一部分，需要将我们生成的新的节点的证书添加上去，其中证书文件地址为:

```
crypto-config/ordererOrganizations/example.com/orderers/orderer5.example.com/tls/server.crt
```

使用`BASE64`转码:

```
cat crypto-config/ordererOrganizations/example.com/orderers/orderer5.example.com/tls/server.crt | base64 > cert.txt
```

在`update_config.json`文件中字段一的部分下面按照字段一的格式添加相同的代码块，并进行修改：
将`cert.txt`文件中的内容复制到字段一的`client_tls_cert,server_tls_cert`对应部分，并修改`host`对应部分为`orderer5.example.com`，`port`为`11050`.

### 2.4更新配置文件

接下来进入`cli`容器:

```
docker exec -it cli bash
```

对原有的配置文件与更新的配置文件进行编码:

```
configtxlator proto_encode --input channel-artifacts/config.json --type common.Config > channel-artifacts/config.pb
configtxlator proto_encode --input channel-artifacts/update_config.json --type common.Config > channel-artifacts/config_update.pb
```

计算出两个文件的差异:

```
configtxlator compute_update --channel_id byfn-sys-channel --original channel-artifacts/config.pb --updated channel-artifacts/config_update.pb > channel-artifacts/updated.pb
```

对该文件进行解码，并添加用于更新配置的头部信息:

```
configtxlator proto_decode --input channel-artifacts/updated.pb --type common.ConfigUpdate > channel-artifacts/updated.json
echo '{"payload":{"header":{"channel_header":{"channel_id":"byfn-sys-channel", "type":2}},"data":{"config_update":'$(cat channel-artifacts/updated.json)'}}}' | jq . > channel-artifacts/updated_envelope.json
```

编码为`Envelope`格式的文件:

```
configtxlator proto_encode --input channel-artifacts/updated_envelope.json --type common.Envelope > channel-artifacts/updated_envelope.pb
```

对该文件进行签名操作，用于更新配置:

```
peer channel signconfigtx -f channel-artifacts/updated_envelope.pb
```

提交更新通道配置交易:

```
peer channel update -f channel-artifacts/updated_envelope.pb -c byfn-sys-channel -o orderer1.example.com:7050 --tls true --cafile $ORDERER_CA
```

如果没有错误的话，新的Orderer节点证书已经成功添加到网络配置中，接下来可以启动新的节点了:

### 2.5 启动新的Orderer节点

写一下新的Orderer节点的Docker文件:

```
version: '2'

volumes:
  orderer5.example.com:

networks:
  byfn:

services:
  orderer5.example.com:
    extends:
      file: base.yaml
      service: orderer-base
    environment:
      - ORDERER_GENERAL_LISTENPORT=11050
    container_name: orderer5.example.com
    networks:
      - byfn
    volumes:
      - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ../crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/msp:/var/hyperledger/orderer/msp
      - ../crypto-config/orderOrganization/example.com/orderers/orderer5.example.com/tls/:/var/hyperledger/orderer/tls
      - ../store/o5:/var/hyperledger/production/orderer
    ports:
      - 11050:11050
```

直接通过命令启动它:

```
docker-compose -f docker-compose-addOrderer5.yaml up -d
```

可以查看新节点的日志确认新的节点已经成功加入了网络。

到这里，本文成功把新的Orderer节点添加进了网络，但是只将该节点添加到了系统通道内，对于应用通道`mychannel`来说，新的节点并没有添加进来，将新的节点添加进`mychannel`通道和以上步骤相同，只需要将通道名称由系统通道修改为`mychannel`即可。本文便不再说明了。
而动态删除节点的过程与添加相似，只不过是从配置文件中删除节点证书。

转载请注明作者与出处：https://www.cnblogs.com/cbkj-xd/ 个人网站主页：https://ifican.top