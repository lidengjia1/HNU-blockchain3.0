# HyperLedger Cello学习笔记

转载请注明出处：[HyperLedger Cello学习笔记](https://www.cnblogs.com/zooqkl/p/10945177.html)

## 概述

**Hyperledger Cello是Hyperledger下的一个子项目，其主要功能如下:**

1. 管理区块链的生命周期，例如自动创建/启动/停止/删除/保持健康状态。
2. 支持定制化的区块链实例，包括区块链类型，大小，共识机制等，当前主要支持的是fabric。
3. 支持虚拟机，本地Docker主机，swarm集群或Kubernetes作为工作节点。
4. 支持异构体系结构，例如X86，POWER和Z，从裸机服务器到虚拟机云。
5. 通过采用附加组件，扩展监控，日志，运行状况和分析功能。

## 架构图

![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173011168-1836160869.png)

Master节点上的服务有：普通用户控制台、操控者控制台、监控、容器编排引擎、日志管理模块、健康检查模块。

![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173021693-1180670344.png)
Cello采取了一主多从的部署模式，Cello Service部署在Master节点上，提供宿主资源的裸机或虚拟环境称为Host，被Cello管理的区块链服务单元称为Worker。整套环境部署要求至少一个Master与一个Worker。

## Cello部署

环境要求：
docker : 17.0+
docker-compose: 1.8.0~1.12.0

#### 安装Master Node

1. ##### 获取源码

```shell
//获取源码
[centos@baas src]$ git clone https://github.com/hyperledger/cello.git && cd cello
$ git checkout -b v0.9.0
```

1. ##### 初始化 Masetr node

```shell
[centos@baas cello]$ cd scripts/master_node/
[centos@baas cello]$ VERSION=0.9.0 bash setup.sh
```

**此过程将会安装docker及docker-compose,还会下载相关镜像。**

```shell
[centos@baas cello]$ docker images
REPOSITORY                             TAG                             IMAGE ID            CREATED             SIZE
hyperledger/cello-watchdog             latest                          a8d77975c69e        10 days ago         1.04GB
hyperledger/cello-watchdog             x86_64-v0.9.0                   a8d77975c69e        10 days ago         1.04GB
hyperledger/cello-mongo                latest                          06cf2b88461a        10 days ago         1.04GB
hyperledger/cello-mongo                x86_64-0.9.0-snapshot-b6a340d   06cf2b88461a        10 days ago         1.04GB
hyperledger/cello-baseimage            x86_64-0.9.0-snapshot-b6a340d   3bdf82ef579a        10 days ago         1.04GB
rabbitmq                               latest                          5cb7660e7cfe        11 days ago         164MB
hyperledger/cello-operator-dashboard   latest                          296a1e79edb9        12 days ago         1.04GB
hyperledger/cello-operator-dashboard   x86_64-v0.9.0                   296a1e79edb9        12 days ago         1.04GB
hyperledger/cello-engine               latest                          ce8908b6719c        12 days ago         1.04GB
hyperledger/cello-engine               x86_64-v0.9.0                   ce8908b6719c        12 days ago         1.04GB
hyperledger/cello-baseimage            latest                          a81a9fc8f1bd        2 weeks ago         1.04GB
hyperledger/cello-baseimage            x86_64-0.9.0                    a81a9fc8f1bd        2 weeks ago         1.04GB
hyperledger/cello-user-dashboard       latest                          c09f0cd2edca        2 months ago        2.12GB
hyperledger/cello-user-dashboard       x86_64-v0.9.0                   c09f0cd2edca        2 months ago        2.12GB
itsthenetwork/nfs-server-alpine        9                               30f582fb8f6e        12 months ago       51.9MB
mongo                                  3.4.10                          e905a87e116d        17 months ago       360MB
node                                   9.2                             cb4c45f7a9e3        17 months ago       676MB
```

1. 修改准备的文件

   ```yaml
   #cryptogen、configtxgen 这两个可执行文件，需要去对应的fabric源码中去编译
   
   源码地址: https://github.com/hyperledger/fabric.git
   
   # 编译命令
   # cd github.com/hyperledger/fabric 
   # make cryptogen
   # make configtxgen
   # 将会在 github.com/hyperledger/fabric/build/bin/ 目录下生成对应的文件
   # 证书需要重新生成，官方给的证书无法使用。  crypto-config.yaml
   # 生成证书命令:  cryptogen generate --config=./crypto-config.yaml
   # 下面是crypto-config.yaml文件内容
   OrdererOrgs:
     - Name: Orderer
       Domain: example.com
       Specs:
         - Hostname: orderer
   PeerOrgs:
     - Name: Org1
       Domain: org1.example.com
       Template:
      Count: 2
       Users:
         Count: 1
     - Name: Org2
       Domain: org2.example.com
       Template:
         Count: 2
       Users:
         Count: 1
   ```

```shell
# 重新生成创始块及其初始化相关的文件
[centos@baas cello]$ cd src/agent/docker/_compose_files/fabric-1.0/kafka
[centos@baas kafka]$ ln -s ../crypto-config ./crypto-config
# 重新生成channel-artifacts下的文件
$ configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/orderer.genesis.block
$ configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/new_businesschannel.tx -channelID businesschannel
$ configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID businesschannel -asOrg Org1MSP
$ configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID businesschannel -asOrg Org2MSP
```

修改docker-compose.yaml文件

```yaml
operator-dashboard:
    image: hyperledger/cello-operator-dashboard
    container_name: cello-operator-dashboard
    hostname: cello-operator-dashboard
    restart: unless-stopped
    environment:
      - MONGO_URL=mongodb://cello-mongo:27017  
      - MONGO_HOST=mongo
      - MONGO_DB=dev
      - EXPLORER_PORT=8088   # 新增浏览器端口配置
      - MONGODB_PORT=27017
      - DEBUG=False  
      - LOG_LEVEL=DEBUG   # 修改日志等级为DEBUG
      - STATIC_FOLDER=$STATIC_FOLDER
      - TEMPLATE_FOLDER=$TEMPLATE_FOLDER
      - ENABLE_EMAIL_ACTIVE=$ENABLE_EMAIL_ACTIVE
      - BROKER=amqp://$RABBITMQ_DEFAULT_USER:$RABBITMQ_DEFAULT_PASS@rabbitmq:5672/$RABBITMQ_DEFAULT_VHOST
      - BACKEND=amqp://$RABBITMQ_DEFAULT_USER:$RABBITMQ_DEFAULT_PASS@rabbitmq:5672/$RABBITMQ_DEFAULT_VHOST
    ports:
      - "8080:8080"
    volumes:
      - $ROOT_PATH/src/agent/docker/_compose_files:/app/agent/docker/_compose_files/  # 新增映射目录
      - $ROOT_PATH/src/agent/docker/_compose_files:/cello
```

修改fabric-kafka-4.yaml文件

```yaml
# 修改peer环境变量
- CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=cello_net_kafka
# 修改cli的command
command: bash -c 'sleep 15; cd /tmp; source scripts/func.sh; bash scripts/test_channel_create.sh; bash scripts/test_channel_join.sh; bash scripts/test_cc_install.sh; bash scripts/test_cc_instantiate.sh ; while true; do sleep 20180101; done'
# 若是在阿里云上部署，需要在的环境变量里添加
- GODEBUG=netdns=go

# 指定容器网络
networks:
  default:
    external:
      name: cello_net_kafka
```

1. 启动Master节点

```shell
[centos@baas cello]$ make start
//将会启动master节点上的服务
[centos@raft-test--3 cello]$ docker ps
CONTAINER ID        IMAGE                                  COMMAND                  CREATED             STATUS              PORTS                                    NAMES
9f9b9abf94a8        hyperledger/cello-user-dashboard       "/bin/sh -c 'ln -sf …"   8 minutes ago       Up 8 minutes        0.0.0.0:8081->8081/tcp                   cello-user-dashboard
1ad2eca89aaa        mongo:3.4.10                           "docker-entrypoint.s…"   8 minutes ago       Up 8 minutes        27017/tcp                                cello-dashboard_mongo
30c61b815fae        rabbitmq                               "docker-entrypoint.s…"   8 minutes ago       Up 8 minutes        4369/tcp, 5671-5672/tcp, 25672/tcp       cello-dashboard_rabbitmq
beb6bbe7d177        hyperledger/cello-engine               "python restserver.py"   8 minutes ago       Up 8 minutes        0.0.0.0:80->80/tcp                       cello-engine
aa011f819b7d        hyperledger/cello-watchdog             "python watchdog.py"     8 minutes ago       Up 8 minutes                                                 cello-watchdog
9c34f967b9c1        mongo:3.4.10                           "docker-entrypoint.s…"   8 minutes ago       Up 8 minutes        127.0.0.1:27017-27018->27017-27018/tcp   cello-mongo
8ae375ce7593        hyperledger/cello-operator-dashboard   "/bin/sh -c '/etc/in…"   8 minutes ago       Up 8 minutes        0.0.0.0:8080->8080/tcp                   cello-operator-dashboard
b102d502edd2        itsthenetwork/nfs-server-alpine:9      "/usr/bin/nfsd.sh"       3 days ago          Up 3 days           0.0.0.0:2049->2049/tcp                   cello-nfs
```

#### 安装worker节点

1. 获取源码

```shell
//获取源码
[centos@baas src]$ git clone -b v0.9.0 https://github.com/hyperledger/cello.git && cd cello
```

1. 初始化 Worker node

```shell
[centos@baas cello]$ cd cello/scripts/worker_node/
[centos@baas cello]$ WORKER_TYPE=docker MASTER_NODE=x.x.x.x make setup-worker
```

**此过程将会安装docker,还会下载fabric相关镜像,包括版本1.0、1.1、1.2。**

```shell
[centos@baas cello]$ docker images
hyperledger/fabric-ca                                                                                   1.2.0               66cc132bd09c        10 months ago       252 MB
hyperledger/fabric-ca                                                                                   amd64-1.2.0         66cc132bd09c        10 months ago       252 MB
hyperledger/fabric-tools                                                                                1.2.0               379602873003        10 months ago       1.51 GB
hyperledger/fabric-tools                                                                                amd64-1.2.0         379602873003        10 months ago       1.51 GB
hyperledger/fabric-ccenv                                                                                amd64-1.2.0         6acf31e2d9a4        10 months ago       1.43 GB
hyperledger/fabric-orderer                                                                              1.2.0               4baf7789a8ec        10 months ago       152 MB
hyperledger/fabric-orderer                                                                              amd64-1.2.0         4baf7789a8ec        10 months ago       152 MB
hyperledger/fabric-peer                                                                                 1.2.0               82c262e65984        10 months ago       159 MB
hyperledger/fabric-peer                                                                                 amd64-1.2.0         82c262e65984        10 months ago       159 MB
hyperledger/fabric-zookeeper                                                                            1.2.0               2b51158f3898        11 months ago       1.44 GB
hyperledger/fabric-zookeeper                                                                            amd64-0.4.10        2b51158f3898        11 months ago       1.44 GB
hyperledger/fabric-kafka                                                                                1.2.0               936aef6db0e6        11 months ago       1.45 GB
hyperledger/fabric-kafka                                                                                amd64-0.4.10        936aef6db0e6        11 months ago       1.45 GB
hyperledger/fabric-couchdb                                                                              1.2.0               3092eca241fc        11 months ago       1.61 GB
hyperledger/fabric-couchdb                                                                              amd64-0.4.10        3092eca241fc        11 months ago       1.61 GB
hyperledger/fabric-baseimage                                                                            amd64-0.4.10        62513965e238        11 months ago       1.39 GB
hyperledger/fabric-baseos                                                                               amd64-0.4.10        52190e831002        11 months ago       132 MB
itsthenetwork/nfs-server-alpine                                                                         9                   30f582fb8f6e        12 months ago       51.9 MB
hyperledger/fabric-ca                                                                                   1.1.0               72617b4fa9b4        14 months ago       299 MB
hyperledger/fabric-ca                                                                                   x86_64-1.1.0        72617b4fa9b4        14 months ago       299 MB
hyperledger/fabric-tools                                                                                1.1.0               b7bfddf508bc        14 months ago       1.46 GB
hyperledger/fabric-tools                                                                                x86_64-1.1.0        b7bfddf508bc        14 months ago       1.46 GB
hyperledger/fabric-orderer                                                                              1.1.0               ce0c810df36a        14 months ago       180 MB
hyperledger/fabric-orderer                                                                              x86_64-1.1.0        ce0c810df36a        14 months ago       180 MB
hyperledger/fabric-peer                                                                                 1.1.0               b023f9be0771        14 months ago       187 MB
hyperledger/fabric-peer                                                                                 x86_64-1.1.0        b023f9be0771        14 months ago       187 MB
hyperledger/fabric-ccenv                                                                                x86_64-1.1.0        c8b4909d8d46        14 months ago       1.39 GB
hyperledger/fabric-baseimage                                                                            x86_64-0.4.6        dbe6787b5747        15 months ago       1.37 GB
hyperledger/fabric-zookeeper                                                                            1.1.0               92cbb952b6f8        15 months ago       1.39 GB
hyperledger/fabric-zookeeper                                                                            x86_64-0.4.6        92cbb952b6f8        15 months ago       1.39 GB
hyperledger/fabric-kafka                                                                                1.1.0               554c591b86a8        15 months ago       1.4 GB
hyperledger/fabric-kafka                                                                                x86_64-0.4.6        554c591b86a8        15 months ago       1.4 GB
hyperledger/fabric-couchdb                                                                              1.1.0               7e73c828fc5b        15 months ago       1.56 GB
hyperledger/fabric-couchdb                                                                              x86_64-0.4.6        7e73c828fc5b        15 months ago       1.56 GB
hyperledger/fabric-baseos                                                                               x86_64-0.4.6        220e5cf3fb7f        15 months ago       151 MB
yeasy/blockchain-explorer                                                                               0.1.0-preview       d3d781c8c96b        16 months ago       659 MB
hyperledger/fabric-tools                                                                                1.0.5               6a8993b718c8        17 months ago       1.33 GB
hyperledger/fabric-tools                                                                                x86_64-1.0.5        6a8993b718c8        17 months ago       1.33 GB
hyperledger/fabric-couchdb                                                                              1.0.5               9a58db2d2723        17 months ago       1.5 GB
hyperledger/fabric-couchdb                                                                              x86_64-1.0.5        9a58db2d2723        17 months ago       1.5 GB
hyperledger/fabric-kafka                                                                                1.0.5               b8c5172bb83c        17 months ago       1.29 GB
hyperledger/fabric-kafka                                                                                x86_64-1.0.5        b8c5172bb83c        17 months ago       1.29 GB
hyperledger/fabric-zookeeper                                                                            1.0.5               68945f4613fc        17 months ago       1.32 GB
hyperledger/fabric-zookeeper                                                                            x86_64-1.0.5        68945f4613fc        17 months ago       1.32 GB
hyperledger/fabric-orderer                                                                              1.0.5               368c78b6f03b        17 months ago       151 MB
hyperledger/fabric-orderer                                                                              x86_64-1.0.5        368c78b6f03b        17 months ago       151 MB
hyperledger/fabric-peer                                                                                 1.0.5               c2ab022f0bdb        17 months ago       154 MB
hyperledger/fabric-peer                                                                                 x86_64-1.0.5        c2ab022f0bdb        17 months ago       154 MB
hyperledger/fabric-ccenv                                                                                x86_64-1.0.5        33feadb8f7a6        17 months ago       1.28 GB
hyperledger/fabric-ca                                                                                   1.0.5               002c9089e464        17 months ago       238 MB
hyperledger/fabric-ca                                                                                   x86_64-1.0.5        002c9089e464        17 months ago       238 MB
hyperledger/fabric-baseimage                                                                            x86_64-0.3.2        c92d9fdee998        21 months ago       1.26 GB
hyperledger/fabric-baseos                                                                               x86_64-0.3.2        bbcbb9da2d83        21 months ago       129 MB
```

1. 对docker进行设置
   **使docker监听2375端口,并且确认master可以调用此接口。**

```shell
//查看配置文件位于哪里
systemctl show --property=FragmentPath docker

//编辑配置文件内容，接收所有ip请求
sudo vim /lib/systemd/system/docker.service

[Service]
ExecStart=/usr/bin/dockerd -H fd:// -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375 --default-ulimit=nofile=8192:16384 --default-ulimit=nproc=8192:16384
//使修改的配置生效
sudo systemctl daemon-reload 
sudo systemctl restart docker.service
// 注意在云上开通此端口，易被在docker上部署挖矿软件，最好加上防火墙，仅允许master节点调用该端口。由于cello不支持TLS，所以安全性较低。
```

#### 通过cello的web界面部署fabric

1. 打开web界面，登录并添加worker主机，用户名: admin, 密码 pass
   ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173126156-83511131.png)
   添加成功后，即可在主机列表中看到该主机。
   ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173141006-1690389604.png)
2. 部署fabric
   ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173153334-1276766713.png)
   在worker主机上，查看cli容器的日志，可以实时查看fabric部署进度。

```shell
[centos@baas cello]$ docker ps |grep cli
[centos@baas cello]$ docker logs -f xxxxxxx_cli
=== Creating channel businesschannel with new_businesschannel.tx... ===
=== Create Channel businesschannel by org 1/peer 0 ===
=== Channel businesschannel is created. ===
=== Created channel businesschannel with new_businesschannel.tx ===

=== Join peers 0 from org 1 into businesschannel... ===
=== Join org 1/peer 0 into channel businesschannel ===
=== org 1/peer 0 joined into channel businesschannel ===
=== Join org 1/peer 1 into channel businesschannel ===
=== org 1/peer 1 joined into channel businesschannel ===
=== Join org 2/peer 0 into channel businesschannel ===
=== org 2/peer 0 joined into channel businesschannel ===
=== Join org 2/peer 1 into channel businesschannel ===
=== org 2/peer 1 joined into channel businesschannel ===
=== Join peers 0 from org 1 into businesschannel Complete ===

=== Installing chaincode exp02 on all 4 peers... ===
=== Install Chaincode exp02:1.0 (examples/chaincode/go/chaincode_example02) on org 1/peer 0 ===
=== Chaincode is installed on remote peer0 ===
=== Install Chaincode exp02:1.0 (examples/chaincode/go/chaincode_example02) on org 1/peer 1 ===
=== Chaincode is installed on remote peer1 ===
=== Install Chaincode exp02:1.0 (examples/chaincode/go/chaincode_example02) on org 2/peer 0 ===
=== Chaincode is installed on remote peer0 ===
=== Install Chaincode exp02:1.0 (examples/chaincode/go/chaincode_example02) on org 2/peer 1 ===
=== Chaincode is installed on remote peer1 ===
=== Install chaincode done ===

=== Instantiating chaincode on channel businesschannel... ===
=== chaincodeInstantiate for channel businesschannel on org 1/peer 0 ====
=== Chaincode Instantiated in channel businesschannel by peer0 ===
=== chaincodeInstantiate for channel businesschannel on org 2/peer 0 ====
=== Chaincode Instantiated in channel businesschannel by peer0 ===
=== Instantiate chaincode on channel businesschannel done ===
```

至此，fabric已经部署完成。

#### 通过cello的用户界面操作fabric

1. 打开界面，并登陆。

![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173210877-499187773.png)

1. 创建链

![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173222622-493537335.png)

1. 查看链详情

![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173233668-791532185.png)

1. 部署合约

- 上传合约

  ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173243022-405749259.png)

  - 安装及部署合约

  ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173305848-1573862008.png)

  - 执行交易

  ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173316192-1316769753.png)

  - channel详情界面可以看到操作记录

  ![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173326267-769891812.png)

转载请注明出处：[HyperLedger Cello学习笔记](https://www.cnblogs.com/zooqkl/p/10945177.html)



root@Hypercello:~/cello# ls
CHANGELOG.md  docker-compose-files  docs      LICENSE         Makefile    mongo  README.md      scripts  test        tox.ini
docker        dockerhub             env.tmpl  MAINTAINERS.md  mkdocs.yml  pip    release_notes  src      thirdparty  user-dashboard
root@Hypercello:~/cello# git checkout -b v0.9.0
fatal: A branch named 'v0.9.0' already exists.
root@Hypercello:~/cello# cd cello/scripts/master_node/
-bash: cd: cello/scripts/master_node/: No such file or directory
root@Hypercello:~/cello# cd scripts/master_node/
root@Hypercello:~/cello/scripts/master_node# VERSION=0.9.0 bash setup.sh
user: root, distro: ubuntu, db_dir: /opt/cello/mongo
Make sure have installed: python-pip, tox, curl and docker-engine
Install software: pip
No pip found, try installing
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  binutils binutils-common binutils-x86-64-linux-gnu build-essential cpp cpp-7 dpkg-dev fakeroot g++ g++-7 gcc gcc-7 gcc-7-base libalgorithm-diff-perl
  libalgorithm-diff-xs-perl libalgorithm-merge-perl libasan4 libatomic1 libbinutils libc-dev-bin libc6-dev libcc1-0 libcilkrts5 libdpkg-perl libexpat1-dev libfakeroot
  libfile-fcntllock-perl libgcc-7-dev libgomp1 libisl19 libitm1 liblsan0 libmpc3 libmpx2 libpython-all-dev libpython-dev libpython-stdlib libpython2.7 libpython2.7-dev
  libpython2.7-minimal libpython2.7-stdlib libquadmath0 libstdc++-7-dev libtsan0 libubsan0 linux-libc-dev make manpages-dev python python-all python-all-dev
  python-asn1crypto python-cffi-backend python-crypto python-cryptography python-dbus python-dev python-enum34 python-gi python-idna python-ipaddress python-keyring
  python-keyrings.alt python-minimal python-pip-whl python-pkg-resources python-secretstorage python-setuptools python-six python-wheel python-xdg python2.7 python2.7-dev
  python2.7-minimal
Suggested packages:
  binutils-doc cpp-doc gcc-7-locales debian-keyring g++-multilib g++-7-multilib gcc-7-doc libstdc++6-7-dbg gcc-multilib autoconf automake libtool flex bison gdb gcc-doc
  gcc-7-multilib libgcc1-dbg libgomp1-dbg libitm1-dbg libatomic1-dbg libasan4-dbg liblsan0-dbg libtsan0-dbg libubsan0-dbg libcilkrts5-dbg libmpx2-dbg libquadmath0-dbg
  glibc-doc bzr libstdc++-7-doc make-doc python-doc python-tk python-crypto-doc python-cryptography-doc python-cryptography-vectors python-dbus-dbg python-dbus-doc
  python-enum34-doc python-gi-cairo gnome-keyring libkf5wallet-bin gir1.2-gnomekeyring-1.0 python-fs python-gdata python-keyczar python-secretstorage-doc
  python-setuptools-doc python2.7-doc binfmt-support
The following NEW packages will be installed:
  binutils binutils-common binutils-x86-64-linux-gnu build-essential cpp cpp-7 dpkg-dev fakeroot g++ g++-7 gcc gcc-7 gcc-7-base libalgorithm-diff-perl
  libalgorithm-diff-xs-perl libalgorithm-merge-perl libasan4 libatomic1 libbinutils libc-dev-bin libc6-dev libcc1-0 libcilkrts5 libdpkg-perl libexpat1-dev libfakeroot
  libfile-fcntllock-perl libgcc-7-dev libgomp1 libisl19 libitm1 liblsan0 libmpc3 libmpx2 libpython-all-dev libpython-dev libpython-stdlib libpython2.7 libpython2.7-dev
  libpython2.7-minimal libpython2.7-stdlib libquadmath0 libstdc++-7-dev libtsan0 libubsan0 linux-libc-dev make manpages-dev python python-all python-all-dev
  python-asn1crypto python-cffi-backend python-crypto python-cryptography python-dbus python-dev python-enum34 python-gi python-idna python-ipaddress python-keyring
  python-keyrings.alt python-minimal python-pip python-pip-whl python-pkg-resources python-secretstorage python-setuptools python-six python-wheel python-xdg python2.7
  python2.7-dev python2.7-minimal
0 upgraded, 75 newly installed, 0 to remove and 9 not upgraded.
Need to get 80.1 MB of archives.
After this operation, 239 MB of additional disk space will be used.
Get:1 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libpython2.7-minimal amd64 2.7.17-1~18.04ubuntu1 [335 kB]
Get:2 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 python2.7-minimal amd64 2.7.17-1~18.04ubuntu1 [1,294 kB]
Get:3 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-minimal amd64 2.7.15~rc1-1 [28.1 kB]
Get:4 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libpython2.7-stdlib amd64 2.7.17-1~18.04ubuntu1 [1,915 kB]
Get:5 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 python2.7 amd64 2.7.17-1~18.04ubuntu1 [248 kB]
Get:6 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libpython-stdlib amd64 2.7.15~rc1-1 [7,620 B]
Get:7 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python amd64 2.7.15~rc1-1 [140 kB]
Get:8 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 binutils-common amd64 2.30-21ubuntu1~18.04.3 [196 kB]
Get:9 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libbinutils amd64 2.30-21ubuntu1~18.04.3 [488 kB]
Get:10 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 binutils-x86-64-linux-gnu amd64 2.30-21ubuntu1~18.04.3 [1,839 kB]
Get:11 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 binutils amd64 2.30-21ubuntu1~18.04.3 [3,388 B]
Get:12 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libc-dev-bin amd64 2.27-3ubuntu1 [71.8 kB]
Get:13 http://mirrors.aliyun.com/ubuntu bionic-proposed/main amd64 linux-libc-dev amd64 4.15.0-107.108 [997 kB]
Get:14 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libc6-dev amd64 2.27-3ubuntu1 [2,587 kB]
Get:15 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 gcc-7-base amd64 7.5.0-3ubuntu1~18.04 [18.3 kB]
Get:16 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libisl19 amd64 0.19-1 [551 kB]
Get:17 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libmpc3 amd64 1.1.0-1 [40.8 kB]
Get:18 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 cpp-7 amd64 7.5.0-3ubuntu1~18.04 [8,591 kB]
Get:19 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 cpp amd64 4:7.4.0-1ubuntu2.3 [27.7 kB]
Get:20 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libcc1-0 amd64 8.4.0-1ubuntu1~18.04 [39.4 kB]
Get:21 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libgomp1 amd64 8.4.0-1ubuntu1~18.04 [76.5 kB]
Get:22 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libitm1 amd64 8.4.0-1ubuntu1~18.04 [27.9 kB]
Get:23 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libatomic1 amd64 8.4.0-1ubuntu1~18.04 [9,192 B]
Get:24 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libasan4 amd64 7.5.0-3ubuntu1~18.04 [358 kB]                                                            
Get:25 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 liblsan0 amd64 8.4.0-1ubuntu1~18.04 [133 kB]                                                            
Get:26 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libtsan0 amd64 8.4.0-1ubuntu1~18.04 [288 kB]                                                            
Get:27 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libubsan0 amd64 7.5.0-3ubuntu1~18.04 [126 kB]                                                           
Get:28 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libcilkrts5 amd64 7.5.0-3ubuntu1~18.04 [42.5 kB]                                                        
Get:29 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libmpx2 amd64 8.4.0-1ubuntu1~18.04 [11.6 kB]                                                            
Get:30 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libquadmath0 amd64 8.4.0-1ubuntu1~18.04 [134 kB]                                                        
Get:31 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libgcc-7-dev amd64 7.5.0-3ubuntu1~18.04 [2,378 kB]                                                      
Get:32 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 gcc-7 amd64 7.5.0-3ubuntu1~18.04 [9,381 kB]                                                             
Get:33 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 gcc amd64 4:7.4.0-1ubuntu2.3 [5,184 B]                                                                  
Get:34 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libstdc++-7-dev amd64 7.5.0-3ubuntu1~18.04 [1,471 kB]                                                   
Get:35 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 g++-7 amd64 7.5.0-3ubuntu1~18.04 [9,697 kB]                                                             
Get:36 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 g++ amd64 4:7.4.0-1ubuntu2.3 [1,568 B]                                                                  
Get:37 http://mirrors.aliyun.com/ubuntu bionic/main amd64 make amd64 4.1-9.1ubuntu1 [154 kB]                                                                               
Get:38 http://mirrors.aliyun.com/ubuntu bionic-updates/main amd64 libdpkg-perl all 1.19.0.5ubuntu2.3 [211 kB]                                                              
Get:39 http://mirrors.aliyun.com/ubuntu bionic-updates/main amd64 dpkg-dev all 1.19.0.5ubuntu2.3 [607 kB]                                                                  
Get:40 http://mirrors.aliyun.com/ubuntu bionic/main amd64 build-essential amd64 12.4ubuntu1 [4,758 B]                                                                      
Get:41 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libfakeroot amd64 1.22-2ubuntu1 [25.9 kB]                                                                        
Get:42 http://mirrors.aliyun.com/ubuntu bionic/main amd64 fakeroot amd64 1.22-2ubuntu1 [62.3 kB]                                                                           
Get:43 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libalgorithm-diff-perl all 1.19.03-1 [47.6 kB]                                                                   
Get:44 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libalgorithm-diff-xs-perl amd64 0.04-5 [11.1 kB]                                                                 
Get:45 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libalgorithm-merge-perl all 0.08-3 [12.0 kB]                                                                     
Get:46 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libexpat1-dev amd64 2.2.5-3ubuntu0.2 [122 kB]                                                           
Get:47 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libfile-fcntllock-perl amd64 0.22-3build2 [33.2 kB]                                                              
Get:48 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libpython2.7 amd64 2.7.17-1~18.04ubuntu1 [1,053 kB]                                                     
Get:49 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libpython2.7-dev amd64 2.7.17-1~18.04ubuntu1 [28.3 MB]                                                  
Get:50 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libpython-dev amd64 2.7.15~rc1-1 [7,684 B]                                                                       
Get:51 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libpython-all-dev amd64 2.7.15~rc1-1 [1,092 B]                                                                   
Get:52 http://mirrors.aliyun.com/ubuntu bionic/main amd64 manpages-dev all 4.15-1 [2,217 kB]                                                                               
Get:53 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-all amd64 2.7.15~rc1-1 [1,076 B]                                                                          
Get:54 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 python2.7-dev amd64 2.7.17-1~18.04ubuntu1 [279 kB]                                                      
Get:55 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-dev amd64 2.7.15~rc1-1 [1,256 B]                                                                          
Get:56 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-all-dev amd64 2.7.15~rc1-1 [1,100 B]                                                                      
Get:57 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-asn1crypto all 0.24.0-1 [72.7 kB]                                                                         
Get:58 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-cffi-backend amd64 1.11.5-1 [63.4 kB]                                                                     
Get:59 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-crypto amd64 2.6.1-8ubuntu2 [244 kB]                                                                      
Get:60 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-enum34 all 1.1.6-2 [34.8 kB]                                                                              
Get:61 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-idna all 2.6-1 [32.4 kB]                                                                                  
Get:62 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-ipaddress all 1.0.17-1 [18.2 kB]                                                                          
Get:63 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-six all 1.11.0-2 [11.3 kB]                                                                                
Get:64 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 python-cryptography amd64 2.1.4-1ubuntu1.3 [221 kB]                                                     
Get:65 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-dbus amd64 1.2.6-1 [90.2 kB]                                                                              
Get:66 http://mirrors.aliyun.com/ubuntu bionic-updates/main amd64 python-gi amd64 3.26.1-2ubuntu1 [197 kB]                                                                 
Get:67 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-secretstorage all 2.3.1-2 [11.8 kB]                                                                       
Get:68 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-keyring all 10.6.0-1 [30.6 kB]                                                                            
Get:69 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-keyrings.alt all 3.0-1 [16.7 kB]                                                                          
Get:70 http://mirrors.aliyun.com/ubuntu bionic-updates/universe amd64 python-pip-whl all 9.0.1-2.3~ubuntu1.18.04.1 [1,653 kB]                                              
Get:71 http://mirrors.aliyun.com/ubuntu bionic-updates/universe amd64 python-pip all 9.0.1-2.3~ubuntu1.18.04.1 [151 kB]                                                    
Get:72 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-pkg-resources all 39.0.1-2 [128 kB]                                                                       
Get:73 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python-setuptools all 39.0.1-2 [329 kB]                                                                          
Get:74 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 python-wheel all 0.30.0-0.2 [36.4 kB]                                                                        
Get:75 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 python-xdg all 0.25-4ubuntu1 [31.3 kB]                                                                       
Fetched 80.1 MB in 16s (4,951 kB/s)                                                                                                                                        
Extracting templates from packages: 100%
Selecting previously unselected package libpython2.7-minimal:amd64.
(Reading database ... 102965 files and directories currently installed.)
Preparing to unpack .../0-libpython2.7-minimal_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking libpython2.7-minimal:amd64 (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package python2.7-minimal.
Preparing to unpack .../1-python2.7-minimal_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking python2.7-minimal (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package python-minimal.
Preparing to unpack .../2-python-minimal_2.7.15~rc1-1_amd64.deb ...
Unpacking python-minimal (2.7.15~rc1-1) ...
Selecting previously unselected package libpython2.7-stdlib:amd64.
Preparing to unpack .../3-libpython2.7-stdlib_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking libpython2.7-stdlib:amd64 (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package python2.7.
Preparing to unpack .../4-python2.7_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking python2.7 (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package libpython-stdlib:amd64.
Preparing to unpack .../5-libpython-stdlib_2.7.15~rc1-1_amd64.deb ...
Unpacking libpython-stdlib:amd64 (2.7.15~rc1-1) ...
Setting up libpython2.7-minimal:amd64 (2.7.17-1~18.04ubuntu1) ...
Setting up python2.7-minimal (2.7.17-1~18.04ubuntu1) ...
Linking and byte-compiling packages for runtime python2.7...
Setting up python-minimal (2.7.15~rc1-1) ...
Selecting previously unselected package python.
(Reading database ... 103713 files and directories currently installed.)
Preparing to unpack .../00-python_2.7.15~rc1-1_amd64.deb ...
Unpacking python (2.7.15~rc1-1) ...
Selecting previously unselected package binutils-common:amd64.
Preparing to unpack .../01-binutils-common_2.30-21ubuntu1~18.04.3_amd64.deb ...
Unpacking binutils-common:amd64 (2.30-21ubuntu1~18.04.3) ...
Selecting previously unselected package libbinutils:amd64.
Preparing to unpack .../02-libbinutils_2.30-21ubuntu1~18.04.3_amd64.deb ...
Unpacking libbinutils:amd64 (2.30-21ubuntu1~18.04.3) ...
Selecting previously unselected package binutils-x86-64-linux-gnu.
Preparing to unpack .../03-binutils-x86-64-linux-gnu_2.30-21ubuntu1~18.04.3_amd64.deb ...
Unpacking binutils-x86-64-linux-gnu (2.30-21ubuntu1~18.04.3) ...
Selecting previously unselected package binutils.
Preparing to unpack .../04-binutils_2.30-21ubuntu1~18.04.3_amd64.deb ...
Unpacking binutils (2.30-21ubuntu1~18.04.3) ...
Selecting previously unselected package libc-dev-bin.
Preparing to unpack .../05-libc-dev-bin_2.27-3ubuntu1_amd64.deb ...
Unpacking libc-dev-bin (2.27-3ubuntu1) ...
Selecting previously unselected package linux-libc-dev:amd64.
Preparing to unpack .../06-linux-libc-dev_4.15.0-107.108_amd64.deb ...
Unpacking linux-libc-dev:amd64 (4.15.0-107.108) ...
Selecting previously unselected package libc6-dev:amd64.
Preparing to unpack .../07-libc6-dev_2.27-3ubuntu1_amd64.deb ...
Unpacking libc6-dev:amd64 (2.27-3ubuntu1) ...
Selecting previously unselected package gcc-7-base:amd64.
Preparing to unpack .../08-gcc-7-base_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking gcc-7-base:amd64 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package libisl19:amd64.
Preparing to unpack .../09-libisl19_0.19-1_amd64.deb ...
Unpacking libisl19:amd64 (0.19-1) ...
Selecting previously unselected package libmpc3:amd64.
Preparing to unpack .../10-libmpc3_1.1.0-1_amd64.deb ...
Unpacking libmpc3:amd64 (1.1.0-1) ...
Selecting previously unselected package cpp-7.
Preparing to unpack .../11-cpp-7_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking cpp-7 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package cpp.
Preparing to unpack .../12-cpp_4%3a7.4.0-1ubuntu2.3_amd64.deb ...
Unpacking cpp (4:7.4.0-1ubuntu2.3) ...
Selecting previously unselected package libcc1-0:amd64.
Preparing to unpack .../13-libcc1-0_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libcc1-0:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libgomp1:amd64.
Preparing to unpack .../14-libgomp1_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libgomp1:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libitm1:amd64.
Preparing to unpack .../15-libitm1_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libitm1:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libatomic1:amd64.
Preparing to unpack .../16-libatomic1_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libatomic1:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libasan4:amd64.
Preparing to unpack .../17-libasan4_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking libasan4:amd64 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package liblsan0:amd64.
Preparing to unpack .../18-liblsan0_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking liblsan0:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libtsan0:amd64.
Preparing to unpack .../19-libtsan0_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libtsan0:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libubsan0:amd64.
Preparing to unpack .../20-libubsan0_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking libubsan0:amd64 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package libcilkrts5:amd64.
Preparing to unpack .../21-libcilkrts5_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking libcilkrts5:amd64 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package libmpx2:amd64.
Preparing to unpack .../22-libmpx2_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libmpx2:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libquadmath0:amd64.
Preparing to unpack .../23-libquadmath0_8.4.0-1ubuntu1~18.04_amd64.deb ...
Unpacking libquadmath0:amd64 (8.4.0-1ubuntu1~18.04) ...
Selecting previously unselected package libgcc-7-dev:amd64.
Preparing to unpack .../24-libgcc-7-dev_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking libgcc-7-dev:amd64 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package gcc-7.
Preparing to unpack .../25-gcc-7_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking gcc-7 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package gcc.
Preparing to unpack .../26-gcc_4%3a7.4.0-1ubuntu2.3_amd64.deb ...
Unpacking gcc (4:7.4.0-1ubuntu2.3) ...
Selecting previously unselected package libstdc++-7-dev:amd64.
Preparing to unpack .../27-libstdc++-7-dev_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking libstdc++-7-dev:amd64 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package g++-7.
Preparing to unpack .../28-g++-7_7.5.0-3ubuntu1~18.04_amd64.deb ...
Unpacking g++-7 (7.5.0-3ubuntu1~18.04) ...
Selecting previously unselected package g++.
Preparing to unpack .../29-g++_4%3a7.4.0-1ubuntu2.3_amd64.deb ...
Unpacking g++ (4:7.4.0-1ubuntu2.3) ...
Selecting previously unselected package make.
Preparing to unpack .../30-make_4.1-9.1ubuntu1_amd64.deb ...
Unpacking make (4.1-9.1ubuntu1) ...
Selecting previously unselected package libdpkg-perl.
Preparing to unpack .../31-libdpkg-perl_1.19.0.5ubuntu2.3_all.deb ...
Unpacking libdpkg-perl (1.19.0.5ubuntu2.3) ...
Selecting previously unselected package dpkg-dev.
Preparing to unpack .../32-dpkg-dev_1.19.0.5ubuntu2.3_all.deb ...
Unpacking dpkg-dev (1.19.0.5ubuntu2.3) ...
Selecting previously unselected package build-essential.
Preparing to unpack .../33-build-essential_12.4ubuntu1_amd64.deb ...
Unpacking build-essential (12.4ubuntu1) ...
Selecting previously unselected package libfakeroot:amd64.
Preparing to unpack .../34-libfakeroot_1.22-2ubuntu1_amd64.deb ...
Unpacking libfakeroot:amd64 (1.22-2ubuntu1) ...
Selecting previously unselected package fakeroot.
Preparing to unpack .../35-fakeroot_1.22-2ubuntu1_amd64.deb ...
Unpacking fakeroot (1.22-2ubuntu1) ...
Selecting previously unselected package libalgorithm-diff-perl.
Preparing to unpack .../36-libalgorithm-diff-perl_1.19.03-1_all.deb ...
Unpacking libalgorithm-diff-perl (1.19.03-1) ...
Selecting previously unselected package libalgorithm-diff-xs-perl.
Preparing to unpack .../37-libalgorithm-diff-xs-perl_0.04-5_amd64.deb ...
Unpacking libalgorithm-diff-xs-perl (0.04-5) ...
Selecting previously unselected package libalgorithm-merge-perl.
Preparing to unpack .../38-libalgorithm-merge-perl_0.08-3_all.deb ...
Unpacking libalgorithm-merge-perl (0.08-3) ...
Selecting previously unselected package libexpat1-dev:amd64.
Preparing to unpack .../39-libexpat1-dev_2.2.5-3ubuntu0.2_amd64.deb ...
Unpacking libexpat1-dev:amd64 (2.2.5-3ubuntu0.2) ...
Selecting previously unselected package libfile-fcntllock-perl.
Preparing to unpack .../40-libfile-fcntllock-perl_0.22-3build2_amd64.deb ...
Unpacking libfile-fcntllock-perl (0.22-3build2) ...
Selecting previously unselected package libpython2.7:amd64.
Preparing to unpack .../41-libpython2.7_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking libpython2.7:amd64 (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package libpython2.7-dev:amd64.
Preparing to unpack .../42-libpython2.7-dev_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking libpython2.7-dev:amd64 (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package libpython-dev:amd64.
Preparing to unpack .../43-libpython-dev_2.7.15~rc1-1_amd64.deb ...
Unpacking libpython-dev:amd64 (2.7.15~rc1-1) ...
Selecting previously unselected package libpython-all-dev:amd64.
Preparing to unpack .../44-libpython-all-dev_2.7.15~rc1-1_amd64.deb ...
Unpacking libpython-all-dev:amd64 (2.7.15~rc1-1) ...
Selecting previously unselected package manpages-dev.
Preparing to unpack .../45-manpages-dev_4.15-1_all.deb ...
Unpacking manpages-dev (4.15-1) ...
Selecting previously unselected package python-all.
Preparing to unpack .../46-python-all_2.7.15~rc1-1_amd64.deb ...
Unpacking python-all (2.7.15~rc1-1) ...
Selecting previously unselected package python2.7-dev.
Preparing to unpack .../47-python2.7-dev_2.7.17-1~18.04ubuntu1_amd64.deb ...
Unpacking python2.7-dev (2.7.17-1~18.04ubuntu1) ...
Selecting previously unselected package python-dev.
Preparing to unpack .../48-python-dev_2.7.15~rc1-1_amd64.deb ...
Unpacking python-dev (2.7.15~rc1-1) ...
Selecting previously unselected package python-all-dev.
Preparing to unpack .../49-python-all-dev_2.7.15~rc1-1_amd64.deb ...
Unpacking python-all-dev (2.7.15~rc1-1) ...
Selecting previously unselected package python-asn1crypto.
Preparing to unpack .../50-python-asn1crypto_0.24.0-1_all.deb ...
Unpacking python-asn1crypto (0.24.0-1) ...
Selecting previously unselected package python-cffi-backend.
Preparing to unpack .../51-python-cffi-backend_1.11.5-1_amd64.deb ...
Unpacking python-cffi-backend (1.11.5-1) ...
Selecting previously unselected package python-crypto.
Preparing to unpack .../52-python-crypto_2.6.1-8ubuntu2_amd64.deb ...
Unpacking python-crypto (2.6.1-8ubuntu2) ...
Selecting previously unselected package python-enum34.
Preparing to unpack .../53-python-enum34_1.1.6-2_all.deb ...
Unpacking python-enum34 (1.1.6-2) ...
Selecting previously unselected package python-idna.
Preparing to unpack .../54-python-idna_2.6-1_all.deb ...
Unpacking python-idna (2.6-1) ...
Selecting previously unselected package python-ipaddress.
Preparing to unpack .../55-python-ipaddress_1.0.17-1_all.deb ...
Unpacking python-ipaddress (1.0.17-1) ...
Selecting previously unselected package python-six.
Preparing to unpack .../56-python-six_1.11.0-2_all.deb ...
Unpacking python-six (1.11.0-2) ...
Selecting previously unselected package python-cryptography.
Preparing to unpack .../57-python-cryptography_2.1.4-1ubuntu1.3_amd64.deb ...
Unpacking python-cryptography (2.1.4-1ubuntu1.3) ...
Selecting previously unselected package python-dbus.
Preparing to unpack .../58-python-dbus_1.2.6-1_amd64.deb ...
Unpacking python-dbus (1.2.6-1) ...
Selecting previously unselected package python-gi.
Preparing to unpack .../59-python-gi_3.26.1-2ubuntu1_amd64.deb ...
Unpacking python-gi (3.26.1-2ubuntu1) ...
Selecting previously unselected package python-secretstorage.
Preparing to unpack .../60-python-secretstorage_2.3.1-2_all.deb ...
Unpacking python-secretstorage (2.3.1-2) ...
Selecting previously unselected package python-keyring.
Preparing to unpack .../61-python-keyring_10.6.0-1_all.deb ...
Unpacking python-keyring (10.6.0-1) ...
Selecting previously unselected package python-keyrings.alt.
Preparing to unpack .../62-python-keyrings.alt_3.0-1_all.deb ...
Unpacking python-keyrings.alt (3.0-1) ...
Selecting previously unselected package python-pip-whl.
Preparing to unpack .../63-python-pip-whl_9.0.1-2.3~ubuntu1.18.04.1_all.deb ...
Unpacking python-pip-whl (9.0.1-2.3~ubuntu1.18.04.1) ...
Selecting previously unselected package python-pip.
Preparing to unpack .../64-python-pip_9.0.1-2.3~ubuntu1.18.04.1_all.deb ...
Unpacking python-pip (9.0.1-2.3~ubuntu1.18.04.1) ...
Selecting previously unselected package python-pkg-resources.
Preparing to unpack .../65-python-pkg-resources_39.0.1-2_all.deb ...
Unpacking python-pkg-resources (39.0.1-2) ...
Selecting previously unselected package python-setuptools.
Preparing to unpack .../66-python-setuptools_39.0.1-2_all.deb ...
Unpacking python-setuptools (39.0.1-2) ...
Selecting previously unselected package python-wheel.
Preparing to unpack .../67-python-wheel_0.30.0-0.2_all.deb ...
Unpacking python-wheel (0.30.0-0.2) ...
Selecting previously unselected package python-xdg.
Preparing to unpack .../68-python-xdg_0.25-4ubuntu1_all.deb ...
Unpacking python-xdg (0.25-4ubuntu1) ...
Setting up libquadmath0:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up libgomp1:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up libatomic1:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up python-pip-whl (9.0.1-2.3~ubuntu1.18.04.1) ...
Setting up libcc1-0:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up make (4.1-9.1ubuntu1) ...
Setting up libtsan0:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up linux-libc-dev:amd64 (4.15.0-107.108) ...
Setting up libdpkg-perl (1.19.0.5ubuntu2.3) ...
Setting up liblsan0:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up gcc-7-base:amd64 (7.5.0-3ubuntu1~18.04) ...
Setting up binutils-common:amd64 (2.30-21ubuntu1~18.04.3) ...
Setting up libfile-fcntllock-perl (0.22-3build2) ...
Setting up libmpx2:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up libfakeroot:amd64 (1.22-2ubuntu1) ...
Setting up libalgorithm-diff-perl (1.19.03-1) ...
Setting up libmpc3:amd64 (1.1.0-1) ...
Setting up libc-dev-bin (2.27-3ubuntu1) ...
Setting up manpages-dev (4.15-1) ...
Setting up libc6-dev:amd64 (2.27-3ubuntu1) ...
Setting up libitm1:amd64 (8.4.0-1ubuntu1~18.04) ...
Setting up libpython2.7-stdlib:amd64 (2.7.17-1~18.04ubuntu1) ...
Setting up libisl19:amd64 (0.19-1) ...
Setting up libasan4:amd64 (7.5.0-3ubuntu1~18.04) ...
Setting up libbinutils:amd64 (2.30-21ubuntu1~18.04.3) ...
Setting up libcilkrts5:amd64 (7.5.0-3ubuntu1~18.04) ...
Setting up libubsan0:amd64 (7.5.0-3ubuntu1~18.04) ...
Setting up python2.7 (2.7.17-1~18.04ubuntu1) ...
Setting up fakeroot (1.22-2ubuntu1) ...
update-alternatives: using /usr/bin/fakeroot-sysv to provide /usr/bin/fakeroot (fakeroot) in auto mode
Setting up libgcc-7-dev:amd64 (7.5.0-3ubuntu1~18.04) ...
Setting up cpp-7 (7.5.0-3ubuntu1~18.04) ...
Setting up libstdc++-7-dev:amd64 (7.5.0-3ubuntu1~18.04) ...
Setting up libpython-stdlib:amd64 (2.7.15~rc1-1) ...
Setting up libalgorithm-merge-perl (0.08-3) ...
Setting up libalgorithm-diff-xs-perl (0.04-5) ...
Setting up libpython2.7:amd64 (2.7.17-1~18.04ubuntu1) ...
Setting up libexpat1-dev:amd64 (2.2.5-3ubuntu0.2) ...
Setting up libpython2.7-dev:amd64 (2.7.17-1~18.04ubuntu1) ...
Setting up python2.7-dev (2.7.17-1~18.04ubuntu1) ...
Setting up python (2.7.15~rc1-1) ...
Setting up python-xdg (0.25-4ubuntu1) ...
Setting up binutils-x86-64-linux-gnu (2.30-21ubuntu1~18.04.3) ...
Setting up python-idna (2.6-1) ...
Setting up cpp (4:7.4.0-1ubuntu2.3) ...
Setting up libpython-dev:amd64 (2.7.15~rc1-1) ...
Setting up python-asn1crypto (0.24.0-1) ...
Setting up python-crypto (2.6.1-8ubuntu2) ...
Setting up python-dev (2.7.15~rc1-1) ...
Setting up python-wheel (0.30.0-0.2) ...
Setting up libpython-all-dev:amd64 (2.7.15~rc1-1) ...
Setting up python-pkg-resources (39.0.1-2) ...
Setting up python-cffi-backend (1.11.5-1) ...
Setting up python-gi (3.26.1-2ubuntu1) ...
Setting up python-six (1.11.0-2) ...
Setting up python-enum34 (1.1.6-2) ...
Setting up binutils (2.30-21ubuntu1~18.04.3) ...
Setting up python-dbus (1.2.6-1) ...
Setting up python-ipaddress (1.0.17-1) ...
Setting up python-pip (9.0.1-2.3~ubuntu1.18.04.1) ...
Setting up python-all (2.7.15~rc1-1) ...
Setting up python-setuptools (39.0.1-2) ...
Setting up gcc-7 (7.5.0-3ubuntu1~18.04) ...
Setting up g++-7 (7.5.0-3ubuntu1~18.04) ...
Setting up gcc (4:7.4.0-1ubuntu2.3) ...
Setting up python-keyrings.alt (3.0-1) ...
Setting up dpkg-dev (1.19.0.5ubuntu2.3) ...
Setting up python-all-dev (2.7.15~rc1-1) ...
Setting up python-cryptography (2.1.4-1ubuntu1.3) ...
Setting up g++ (4:7.4.0-1ubuntu2.3) ...
update-alternatives: using /usr/bin/g++ to provide /usr/bin/c++ (c++) in auto mode
Setting up python-secretstorage (2.3.1-2) ...
Setting up python-keyring (10.6.0-1) ...
Setting up build-essential (12.4ubuntu1) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Processing triggers for mime-support (3.60ubuntu1) ...
Processing triggers for libc-bin (2.27-3ubuntu1) ...
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  javascript-common keyutils libjs-jquery libjs-sphinxdoc libjs-underscore libnfsidmap2 libtirpc1 python3-distutils python3-lib2to3 python3-pluggy python3-py
  python3-setuptools python3-virtualenv rpcbind virtualenv
Suggested packages:
  apache2 | lighttpd | httpd watchdog subversion python3-pytest python-setuptools-doc
The following NEW packages will be installed:
  javascript-common keyutils libjs-jquery libjs-sphinxdoc libjs-underscore libnfsidmap2 libtirpc1 nfs-common python3-distutils python3-lib2to3 python3-pluggy python3-py
  python3-setuptools python3-virtualenv rpcbind tox virtualenv
0 upgraded, 17 newly installed, 0 to remove and 9 not upgraded.
Need to get 1,444 kB of archives.
After this operation, 8,746 kB of additional disk space will be used.
Get:1 http://mirrors.aliyun.com/ubuntu bionic/main amd64 javascript-common all 11 [6,066 B]
Get:2 http://mirrors.aliyun.com/ubuntu bionic/main amd64 keyutils amd64 1.5.9-9.2ubuntu2 [47.9 kB]
Get:3 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libjs-jquery all 3.2.1-1 [152 kB]
Get:4 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libjs-underscore all 1.8.3~dfsg-1 [59.9 kB]
Get:5 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libjs-sphinxdoc all 1.6.7-1ubuntu1 [85.6 kB]
Get:6 http://mirrors.aliyun.com/ubuntu bionic/main amd64 libnfsidmap2 amd64 0.25-5.1 [27.2 kB]
Get:7 http://mirrors.aliyun.com/ubuntu bionic-security/main amd64 libtirpc1 amd64 0.2.5-1.2ubuntu0.1 [75.7 kB]
Get:8 http://mirrors.aliyun.com/ubuntu bionic/main amd64 rpcbind amd64 0.2.3-0.6 [40.6 kB]
Get:9 http://mirrors.aliyun.com/ubuntu bionic-updates/main amd64 nfs-common amd64 1:1.3.4-2.1ubuntu5.2 [205 kB]
Get:10 http://mirrors.aliyun.com/ubuntu bionic-updates/main amd64 python3-lib2to3 all 3.6.9-1~18.04 [77.4 kB]
Get:11 http://mirrors.aliyun.com/ubuntu bionic-updates/main amd64 python3-distutils all 3.6.9-1~18.04 [144 kB]
Get:12 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 python3-pluggy all 0.6.0-1 [12.3 kB]
Get:13 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 python3-py all 1.5.2-1 [65.2 kB]
Get:14 http://mirrors.aliyun.com/ubuntu bionic/main amd64 python3-setuptools all 39.0.1-2 [248 kB]
Get:15 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 python3-virtualenv all 15.1.0+ds-1.1 [43.4 kB]
Get:16 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 virtualenv all 15.1.0+ds-1.1 [4,476 B]
Get:17 http://mirrors.aliyun.com/ubuntu bionic/universe amd64 tox all 2.5.0-1 [149 kB]
Fetched 1,444 kB in 3s (505 kB/s)
Selecting previously unselected package javascript-common.
(Reading database ... 110232 files and directories currently installed.)
Preparing to unpack .../00-javascript-common_11_all.deb ...
Unpacking javascript-common (11) ...
Selecting previously unselected package keyutils.
Preparing to unpack .../01-keyutils_1.5.9-9.2ubuntu2_amd64.deb ...
Unpacking keyutils (1.5.9-9.2ubuntu2) ...
Selecting previously unselected package libjs-jquery.
Preparing to unpack .../02-libjs-jquery_3.2.1-1_all.deb ...
Unpacking libjs-jquery (3.2.1-1) ...
Selecting previously unselected package libjs-underscore.
Preparing to unpack .../03-libjs-underscore_1.8.3~dfsg-1_all.deb ...
Unpacking libjs-underscore (1.8.3~dfsg-1) ...
Selecting previously unselected package libjs-sphinxdoc.
Preparing to unpack .../04-libjs-sphinxdoc_1.6.7-1ubuntu1_all.deb ...
Unpacking libjs-sphinxdoc (1.6.7-1ubuntu1) ...
Selecting previously unselected package libnfsidmap2:amd64.
Preparing to unpack .../05-libnfsidmap2_0.25-5.1_amd64.deb ...
Unpacking libnfsidmap2:amd64 (0.25-5.1) ...
Selecting previously unselected package libtirpc1:amd64.
Preparing to unpack .../06-libtirpc1_0.2.5-1.2ubuntu0.1_amd64.deb ...
Unpacking libtirpc1:amd64 (0.2.5-1.2ubuntu0.1) ...
Selecting previously unselected package rpcbind.
Preparing to unpack .../07-rpcbind_0.2.3-0.6_amd64.deb ...
Unpacking rpcbind (0.2.3-0.6) ...
Selecting previously unselected package nfs-common.
Preparing to unpack .../08-nfs-common_1%3a1.3.4-2.1ubuntu5.2_amd64.deb ...
Unpacking nfs-common (1:1.3.4-2.1ubuntu5.2) ...
Selecting previously unselected package python3-lib2to3.
Preparing to unpack .../09-python3-lib2to3_3.6.9-1~18.04_all.deb ...
Unpacking python3-lib2to3 (3.6.9-1~18.04) ...
Selecting previously unselected package python3-distutils.
Preparing to unpack .../10-python3-distutils_3.6.9-1~18.04_all.deb ...
Unpacking python3-distutils (3.6.9-1~18.04) ...
Selecting previously unselected package python3-pluggy.
Preparing to unpack .../11-python3-pluggy_0.6.0-1_all.deb ...
Unpacking python3-pluggy (0.6.0-1) ...
Selecting previously unselected package python3-py.
Preparing to unpack .../12-python3-py_1.5.2-1_all.deb ...
Unpacking python3-py (1.5.2-1) ...
Selecting previously unselected package python3-setuptools.
Preparing to unpack .../13-python3-setuptools_39.0.1-2_all.deb ...
Unpacking python3-setuptools (39.0.1-2) ...
Selecting previously unselected package python3-virtualenv.
Preparing to unpack .../14-python3-virtualenv_15.1.0+ds-1.1_all.deb ...
Unpacking python3-virtualenv (15.1.0+ds-1.1) ...
Selecting previously unselected package virtualenv.
Preparing to unpack .../15-virtualenv_15.1.0+ds-1.1_all.deb ...
Unpacking virtualenv (15.1.0+ds-1.1) ...
Selecting previously unselected package tox.
Preparing to unpack .../16-tox_2.5.0-1_all.deb ...
Unpacking tox (2.5.0-1) ...
Setting up libjs-jquery (3.2.1-1) ...
Setting up libnfsidmap2:amd64 (0.25-5.1) ...
Setting up libjs-underscore (1.8.3~dfsg-1) ...
Setting up python3-py (1.5.2-1) ...
Setting up keyutils (1.5.9-9.2ubuntu2) ...
Setting up libjs-sphinxdoc (1.6.7-1ubuntu1) ...
Setting up python3-pluggy (0.6.0-1) ...
Setting up libtirpc1:amd64 (0.2.5-1.2ubuntu0.1) ...
Setting up javascript-common (11) ...
Setting up python3-lib2to3 (3.6.9-1~18.04) ...
Setting up python3-distutils (3.6.9-1~18.04) ...
Setting up rpcbind (0.2.3-0.6) ...
Created symlink /etc/systemd/system/multi-user.target.wants/rpcbind.service → /lib/systemd/system/rpcbind.service.
Created symlink /etc/systemd/system/sockets.target.wants/rpcbind.socket → /lib/systemd/system/rpcbind.socket.
Setting up nfs-common (1:1.3.4-2.1ubuntu5.2) ...

Creating config file /etc/idmapd.conf with new version
Adding system user `statd' (UID 111) ...
Adding new user `statd' (UID 111) with group `nogroup' ...
Not creating home directory `/var/lib/nfs'.
Created symlink /etc/systemd/system/multi-user.target.wants/nfs-client.target → /lib/systemd/system/nfs-client.target.
Created symlink /etc/systemd/system/remote-fs.target.wants/nfs-client.target → /lib/systemd/system/nfs-client.target.
nfs-utils.service is a disabled or a static unit, not starting it.
Setting up python3-virtualenv (15.1.0+ds-1.1) ...
Setting up python3-setuptools (39.0.1-2) ...
Setting up virtualenv (15.1.0+ds-1.1) ...
Setting up tox (2.5.0-1) ...
Processing triggers for libc-bin (2.27-3ubuntu1) ...
Processing triggers for systemd (237-3ubuntu10.41) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Processing triggers for ureadahead (0.100.0-21) ...
Add existing user root to docker group
Checking existing containers
Download required Docker images for Cello Services...
Downloading the docker images for Cello services: VERSION=0.9.0 ARCH=x86_64
Check node:9.2 image.
pulling node:9.2
9.2: Pulling from library/node
f49cf87b52c1: Pull complete 
7b491c575b06: Pull complete 
b313b08bab3b: Pull complete 
51d6678c3f0e: Pull complete 
da59faba155b: Pull complete 
7f84ea62c1fd: Pull complete 
70daba82d737: Pull complete 
e58cf1fcd0d3: Pull complete 
Digest: sha256:1d4a8dbe3817d65b5915de8c5df1c6b223986514b286275490cb15d55438e8b6
Status: Downloaded newer image for node:9.2
docker.io/library/node:9.2
Pulling hyperledger/cello-baseimage:x86_64-0.9.0 from dockerhub
x86_64-0.9.0: Pulling from hyperledger/cello-baseimage
c7b7d16361e0: Pulling fs layer 
b7a128769df1: Pulling fs layer 
1128949d0793: Pulling fs layer 
667692510b70: Waiting 
bed4ecf88e6a: Waiting 
94d1c1cbf347: Waiting 
f59f6b55cd0f: Waiting 
6513a2441bbb: Waiting 
792e28117005: Waiting 
de833a28e748: Waiting 
67dc298b39ec: Waiting 
77cec3cff815: Waiting 
a48745f8002c: Waiting 
ee9fd47693eb: Waiting 
Get https://registry-1.docker.io/v2/: net/http: TLS handshake timeout
Error response from daemon: No such image: hyperledger/cello-baseimage:x86_64-0.9.0
Pulling hyperledger/cello-engine:x86_64-0.9.0 from dockerhub
Error response from daemon: Get https://registry-1.docker.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
Error response from daemon: No such image: hyperledger/cello-engine:x86_64-0.9.0
Pulling hyperledger/cello-mongo:x86_64-0.9.0 from dockerhub
Error response from daemon: Get https://registry-1.docker.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
Error response from daemon: No such image: hyperledger/cello-mongo:x86_64-0.9.0
Pulling hyperledger/cello-operator-dashboard:x86_64-0.9.0 from dockerhub
Error response from daemon: Get https://registry-1.docker.io/v2/: net/http: TLS handshake timeout
Error response from daemon: No such image: hyperledger/cello-operator-dashboard:x86_64-0.9.0
Pulling hyperledger/cello-user-dashboard:x86_64-0.9.0 from dockerhub
Error response from daemon: Get https://registry-1.docker.io/v2/: net/http: TLS handshake timeout
Error response from daemon: No such image: hyperledger/cello-user-dashboard:x86_64-0.9.0
Pulling hyperledger/cello-watchdog:x86_64-0.9.0 from dockerhub
Error response from daemon: Get https://registry-1.docker.io/v2/: net/http: TLS handshake timeout
Error response from daemon: No such image: hyperledger/cello-watchdog:x86_64-0.9.0
3.4.10: Pulling from library/mongo
c4bb02b17bb4: Pull complete 
3f58e3bb3be4: Pull complete 
a229fb575a6e: Pull complete 
8f5ddc533743: Pull complete 
5e9d2af6e206: Pull complete 
8b17a0553222: Pull complete 
1723351cb8b4: Pull complete 
78dd8437755a: Pull complete 
61012c043572: Pull complete 
76aeb242d479: Pull complete 
bdc8171e92f4: Pull complete 
Digest: sha256:b84baeffd0f14bebaf057b36de9414ee41584a897351795f4a3889257cf19b6d
Status: Downloaded newer image for mongo:3.4.10
docker.io/library/mongo:3.4.10
9: Pulling from itsthenetwork/nfs-server-alpine
2fdfe1cd78c2: Pull complete 
0d5cf19ef7ec: Pull complete 
e71117978fa3: Pull complete 
1740ce97b1ca: Pull complete 
741c5bcecf21: Pull complete 
e8bfe8fb8b55: Pull complete 
58e2f7c38e16: Pull complete 
f58765ded7b4: Pull complete 
235bcf0c3941: Pull complete 
62adbce8ad35: Pull complete 
0c0fb61ba6a8: Pull complete 
Digest: sha256:270e12213da124fdd02d491d24694dd1562f8de02a6e75c5ff02d736ee66160d
Status: Downloaded newer image for itsthenetwork/nfs-server-alpine:9
docker.io/itsthenetwork/nfs-server-alpine:9
All Image downloaded 
Checking local mounted database path /opt/cello/mongo...
Local database path /opt/cello/mongo not existed, creating one
Setup done, please logout and login again.
It's safe to run this script repeatedly.

root@Hypercello:~/cello/scripts/master_node# docker images
REPOSITORY                        TAG                 IMAGE ID            CREATED             SIZE
itsthenetwork/nfs-server-alpine   9                   30f582fb8f6e        2 years ago         51.9MB
mongo                             3.4.10              e905a87e116d        2 years ago         360MB
node                              9.2                 cb4c45f7a9e3        2 years ago         676MB

**cello部署分为master节点部署和worker节点部署,** 

**master节点为管理baas平台的节点,worker节点为工作节点,**

**worker节点可以跟master节点部署在一起,也可以分开.**

 

### 系统要求

- Hardware: 8c16g100g
- Linux Kernel >= 3.0.0
- Docker engine: 1.10.0+ (Docker 18.0+ support is experimental)
- docker-compose: 1.10.0+(切记默认apt下载的版本比此版本要小)

其它环境,默认下载不再赘述,说下怎么下载最新的docker-compose

**公共部分**

#### 1下载docker-compose

```
$ sudo curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
```

#### 2授权

```
$ sudo chmod +x /usr/local/bin/docker-compose
```

#### 3查看版本信息

```
$ docker-compose --version
```

#### 4克隆项目

```
git clone https://github.com/hyperledger/cello.git
cd cello
```

 

#### **5运行安装程序**

首次运行时

```
cd ~/cello/scripts/worker_node
sudo vi setup_worker_node_docker.sh
```

找到第28行,MASTER_NODE=" "赋予MASTER节点的ip

**非公共部分**

#### **6 master配置**


pull运行所需的镜像,此过程一直在pull镜像，需等待几分钟

```
make setup-master
```

##### 启动

```
SERVER_PUBLIC_IP=x.x.x.x make start
```

 

注:浏览器访问master_ip:8080进入的是operator_dashboard，master_ip:8081进入的是user_dashboard
至此master安装完毕

##### 停止

```
make stop
```

##### 重新开始

```
make restart
```

#### 7 worker配置

(本文只有docker版本的)

docker开放外网IP访问，此处开启docker外网访问，开放2375端口。

```
vim /lib/systemd/system/docker.service
其中的
ExecStart=/usr/bin/dockerd -H fd://
改为
ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375 --default-ulimit=nofile=8192:16384 --default-ulimit=nproc=8192:16384
```

重新配置文件并重启docker

```
systemctl daemon-reload
systemctl restart docker.service
```

安装nfs服务(服务端挂挂载文件到本地的一个工具)

```
apt-get install nfs-common
```

pull Worker端所需的镜像,拉取镜像，worker端已clone过代码，所以直接

```
cd cello
WORKDER_TYPE=docker MASTER_NODE=master_ip make setup-worker
```

判断 nfs 是否挂载成功

```
ls /opt/cello
fabric-1.0 fabric-1.1 fabric-1.2
```

注:出现上面显示,少一个都不行,说明成功挂载nfs服务器

注:docker没有开启外网访问的话,不能添加节点

注:nfs没有挂载成功的话,添加chaincode会失败

 原文链接：https://www.cnblogs.com/sfgoto/p/10750948.html
