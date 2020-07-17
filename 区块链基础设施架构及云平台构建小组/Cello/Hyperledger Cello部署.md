# 在Kubernetes上部署Hyperledger Fabric第1部分

## **总览**

Fabric是Linux Foundation托管的Hyperledger项目之一。它提供了开发区块链应用程序的框架。自今年7月发布Fabric 1.0以来，人们渴望使用Fabric来构建应用程序以解决其业务问题。但是，由于配置的复杂性，许多人在部署和管理Fabric系统时遇到困难。

为了简化Fabric的操作，我们需要一些工具来帮助我们更好地管理Fabric的分布式系统。为此，Kubernetes似乎是理想的选择。（有趣的是，Kubernetes是CNCF下的旗舰项目，也是Linux Foundation项目。）

首先，将Fabric的位内置到容器映像中。它的chaincode（智能合约）还利用容器在沙箱中运行。Fabric系统由在多个容器中运行的组件组成。另一方面，Kubernetes正在成为控制容器化应用程序的部署，扩展和其他管理自动化的主导平台。两者之间自然契合。

其次，Fabric组件可以通过在Kubernetes上进行部署来实现高可用性。Kubernetes具有称为复制器的功能，该功能监视运行中的Pod并自动启动崩溃的Pod。

第三，Kubernetes支持多租户。我们可以在同一Kubernetes平台上运行多个隔离的Fabric实例。这促进了区块链应用程序的开发和测试。

在以下各节中，我们介绍一种在Kubernetes上部署Fabric的方法。我们假设读者具有Fabric，Docker容器和Kubernetes的基础知识。

**网络拓扑结构**

我们的网络拓扑如图1所示。物理网络用蓝线表示。Kubernetes具有一个或多个主节点和工作节点。除此之外，我们还有一台CMD计算机作为客户端来发布部署命令。NFS服务器用作配置文件和其他数据的共享文件系统。所有这些节点都通过物理网络（例如192.168.0.1/24）连接。

Kubernetes的网络模型使所有Pod都可以直接相互连接，无论它们位于哪个节点上。通过使用Kubernetes的CNI插件（例如Flannel），可以轻松地为此目的创建覆盖网络。如图1中的红线所示（省略了法兰绒组件的一些详细信息），Kubernetes将所有Pod连接到Flannel网络，从而允许这些Pod的容器彼此正确通信。

可以在附加配置文件中指定Flannel网络的IP地址范围以及kube_dns的IP地址。我们需要确保kube_dns的IP地址必须在指定的地址范围内。例如，在图1中，Flannel网络为10.0.0.1/16，而kube_dns地址为10.0.0.10。

[![img](http://www.think-foundry.com/wp-content/uploads/2017/11/networkTopo-1024x333.png)](http://www.think-foundry.com/wp-content/uploads/2017/11/networkTopo.png)

图1

**将结构组件映射到Kubernetes Pod**

[![img](http://www.think-foundry.com/wp-content/uploads/2017/11/fabricDeployment-1024x483.png)](http://www.think-foundry.com/wp-content/uploads/2017/11/fabricDeployment.png)

图2

Fabric是包含多个节点的分布式系统。这些节点可以属于不同的实体。如图2所示，每个组织都有自己的一组节点（为简单起见，未显示所有节点）。订购者还提供公共共识服务。要将Fabric部署到Kubernetes上，我们需要将所有组件转换为Pod进行部署，并使用命名空间隔离组织。

在Kubernetes中，名称空间是一个重要的概念。它用于在多个用户之间分配群集资源。对于Fabric，可以将组织映射到名称空间，以便拥有专用资源。完成此映射后，可以按域名区分每个组织的对等方。此外，我们可以通过设置网络策略来隔离不同的组织（本博客中未介绍）。

如图2所示，假设在Fabric网络中有N个对等组织和M个订购者组织。这是我们如何在Kubernetes上划分它们的方法：

**A）**如果是Fabric，我们为第N个对等组织分配名称*orgN*。在Kubernetes中其对应的命名空间也称为*orgN*。Fabric *orgN的*所有组件都将放入*Kubernetes*中的名称空间*orgN*中。每个组织的命名空间下都有多个Pod。吊舱是Kubernetes中的一个部署单元，它由一个或多个容器组成。我们可以将每个组织的Fabric容器捆绑到几个吊舱中。这些吊舱类型如下：

- **Peer Pod：**包括Fabric对等方，ouchDB（可选），代表组织的对等节点。每个组织可以有一个或多个对等豆荚。
- **CA Server Pod：**组织的Fabric CA Server节点。通常，组织中需要一个pod。
- **CLI Pod ：（**可选）为命令行工具提供一个环境来操纵组织的节点。Fabric的对等环境变量在此窗格中配置。

[![img](http://www.think-foundry.com/wp-content/uploads/2017/11/peerNS.png)](http://www.think-foundry.com/wp-content/uploads/2017/11/peerNS.png)

图3

**B）** Fabric中可能有一个或多个订购者。我们将第M个*订购*者组织的名称设置为*orgordererM*。它在Kubernetes上的对应命名空间是*orgordererM*。它具有一个或多个Pod以运行订购者节点。

![img](http://www.think-foundry.com/wp-content/uploads/2017/11/ordererNS.png)

图4

**C）**如果将Kafka用于共识过程，我们可以将Kafka放在单独的命名空间中。它仅用于运行和管理Zookeeper和Kafka容器。

综上所述，总体部署如下所示：

[![img](http://www.think-foundry.com/wp-content/uploads/2017/11/k8spods-1024x593.png)](http://www.think-foundry.com/wp-content/uploads/2017/11/k8spods.png)

图5

**共享存储**

在部署Fabric之前，我们需要准备其组件（例如对等方和订购者）的配置文件。这是一个非常复杂的过程，并且容易出错。幸运的是，我们创建了一个工具来自动生成这些配置文件。生成的文件存储在NFS之类的共享文件系统中。

以后启动Fabric的Pod时，我们将配置文件的不同子集安装到Pod中，以便它们具有特定于它们所属组织的配置。

在Kubernetes中，我们可以使用持久卷（PV）和持久卷声明（PVC）将文件或目录挂载到Pod中。我们为Fabric中的每个组织创建PV和PVC，以实现资源隔离。每个组织只能在NFS服务器中看到其自己的目录。

创建PV之后，我们定义PVC，以便Fabric节点可以使用PV来访问相应的目录和文件。

以对等组织org1为例。首先，我们创建一个名称空间org1及其PV。PV映射到NFS上的目录*/ opt / share / crypto-config / peerOrganizations / org1*。其次，我们创建PVC来消耗PV。命名空间org1下的所有Pod使用相同的PVC。但是，我们仅通过在Pod配置文件中指定安装路径将必需的文件映射到每个Pod中。

图6显示了Pod及其NFS共享目录之间的关系。变量$ PVC表示PVC安装点，在此示例中为*/ opt / share / crypto-config / peerOrganizations / org1*。

[![img](http://www.think-foundry.com/wp-content/uploads/2017/11/NFSmapping-1024x622.png)](http://www.think-foundry.com/wp-content/uploads/2017/11/NFSmapping.png)

图6

**Fabric组件之间的通信**

当所有Fabric组件都放入Kubernetes的Pod中时，我们需要考虑这些Pod之间的网络连接。Kubernetes中的每个Pod都有一个内部IP地址，但是由于IP地址是Pod临时的，因此很难使用IP和端口在Pod之间进行通信。当Pod重新启动时，其IP地址也将更改。因此，有必要在Kubernetes中为Pod创建服务，以便它们可以通过服务名称相互通信。服务的命名应遵循以下原则，以显示其绑定到的Pod信息：

**1）**服务的名称空间和pod应当一致。
**2）**服务名称应与容器在容器中的ID一致。

例如，组织org1的Fabric的peer0映射到名称空间org1下名为peer0的容器。与其绑定的服务应命名为peer0.org1，其中peer0是服务的名称，而org1是服务的名称空间。其他Pod可以通过服务名称peer0.org1连接到org1的peer0，该名称显示为peer0的主机名。

**解决chaincode沙箱**

当Fabric中的同位体实例化一个链码时，它将创建一个在其中运行链码的Docker容器。它调用以创建容器的Docker API端点为unix：///var/run/docker.sock。只要对等容器和chaincode容器由同一Docker引擎管理，此机制就可以很好地工作。但是，在Kubernetes中，链码容器是由对等方创建的而不通知Kubernetes。因此，链码和对等容器无法相互连接，这在实例化链码时导致失败。

要解决此问题，我们需要在每个工作程序节点的Docker引擎中添加Kube_dns IP地址。在Docker引擎的配置文件中添加以下选项。在下面的示例中，10.0.0.10是kube_dns pod的IP地址。在您的环境中将其替换为正确的值。

```
“ --dns = 10.0.0.10 --dns = 192.168.0.1 --dns-search \
default.svc.cluster.local --dns-搜索\
svc.cluster.local --dns-opt ndots：2 --dns-opt \
超时：2 --dns-opt尝试：2“
```

到目前为止，我们已经说明了将Fabric部署到Kubernetes上的关键点。在下一篇文章中，我们将描述部署的详细步骤。对于迫不及待的人，请下载我们的Fling“ **[vSphere上的区块链](https://labs.vmware.com/flings/blockchain-on-vsphere)** ”，以了解其工作原理。它是一个自动化工具，可让您以最少的配置在Kubernetes上部署Fabric。如果不使用vSphere运行Kubernetes，则可以为Kubernetes实例选择任何基础架构。只需跳过该部分即可在vSphere上部署Kubernetes。

**继续第2部分：**









# [Hyperledger Cello部署](https://www.cnblogs.com/sfgoto/p/10750948.html)

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

#### 1，下载docker-compose

```
$ sudo curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
```

#### 2，授权

```
$ sudo chmod +x /usr/local/bin/docker-compose
```

#### 3，查看版本信息

```
$ docker-compose --version
```

### 4.克隆项目

```
git clone https://github.com/hyperledger/cello.git
cd cello
```

 

**5.运行安装程序**

首次运行时

```
cd ~/cello/scripts/worker_node
vim setup_worker_node_docker.sh
```

找到第28行,MASTER_NODE=" "赋予MASTER节点的ip

**非公共部分**

**6. master配置**


pull运行所需的镜像,此过程一直在pull镜像，需等待几分钟

```
cd ~/cello

make setup-master
```

启动

```
SERVER_PUBLIC_IP=x.x.x.x make start
```

 

注:浏览器访问master_ip:8080进入的是operator_dashboard，master_ip:8081进入的是user_dashboard
至此master安装完毕

停止

```
make stop
```

重新开始

```
make restart
```

**7. worker配置(本文只有docker版本的)**

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

 

# HyperLedger Cello学习笔记

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

1. 获取源码

```shell
//获取源码
[centos@baas src]$ git clone -b v0.9.0 https://github.com/hyperledger/cello.git && cd cello
```

1. 初始化 Masetr node

```shell
[centos@baas cello]$ cd cello/scripts/master_node/
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
sudo vim /usr/lib/systemd/system/docker.service

[Service]
ExecStart=/usr/bin/dockerd -H fd:// -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375 --default-ulimit=nofile=8192:16384 --default-ulimit=nproc=8192:16384
//使修改的配置生效
sudo systemctl daemon-reload; sudo systemctl restart docker.service
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





```
注意点：

1、证书需要重新生成，官方给的证书无法使用。
2、重新生成证书后，channel-artifacts下的文件也需要重新生成。
3、operator-dashboard 容器的环境变量需新增浏览器端口。
日志等级设置为DEBUG，方便查错。
- EXPLORER_PORT=8088
- LOG_LEVEL=DEBUG
volumes 里需新增映射
- $ROOT_PATH/src/agent/docker/_compose_files:/app/agent/docker/_compose_files/
4、若是在阿里云上部署，需要在orderer和peer的环境变量里添加
- GODEBUG=netdns=go
5、容器网络最好写死，否则找不到网络。
# 指定容器网络
networks:
default:
external:
name: cello_net_kafka
# 修改peer环境变量
- CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=cello_net_kafka
```



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



![image-20200612112029182](Hyperledger Cello部署.assets/image-20200612112029182.png)





Master节点上的服务有：普通用户控制台、操控者控制台、监控、容器编排引擎、日志管理模块、健康检查模块。

![img](https://img2018.cnblogs.com/blog/1592847/201905/1592847-20190529173021693-1180670344.png)
Cello采取了一主多从的部署模式，Cello Service部署在Master节点上，提供宿主资源的裸机或虚拟环境称为Host，被Cello管理的区块链服务单元称为Worker。整套环境部署要求至少一个Master与一个Worker。

## Cello部署

环境要求：
docker : 17.0+
docker-compose: 1.8.0~1.12.0

#### 安装Master Node

1. 获取源码

```shell
//获取源码
[centos@baas src]$ git clone -b v0.9.0 https://github.com/hyperledger/cello.git && cd cello
```

1. 初始化 Masetr node

```shell
[centos@baas cello]$ cd cello/scripts/master_node/
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
sudo systemctl daemon-reload; sudo systemctl restart docker.service
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



# 如何在Kubernetes上部署Hyperledger Fabric

最初由Henry Zhang 发表于 2017年11月8日![img](https://hackernoon.com/hn-images/0*XWJnYYVtYsV71Jwc.png)

[](https://hackernoon.com/u/zhanghenry)

### 总览

Fabric是Linux Foundation托管的Hyperledger项目之一。它提供了开发区块链应用程序的框架。自今年7月发布Fabric 1.0以来，人们渴望使用Fabric来构建应用程序以解决其业务问题。但是，由于配置的复杂性，许多人在部署和管理Fabric系统时遇到困难。

为了简化Fabric的操作，我们需要一些工具来帮助我们更好地管理Fabric的分布式系统。为此，Kubernetes似乎是理想的选择。（有趣的是，Kubernetes是CNCF下的旗舰项目，也是Linux Foundation项目。）

首先，将Fabric的位内置到容器映像中。它的chaincode（智能合约）还利用容器在沙箱中运行。Fabric系统由在多个容器中运行的组件组成。另一方面，Kubernetes正在成为控制容器化应用程序的部署，扩展和其他管理自动化的主导平台。两者之间自然契合。

其次，Fabric组件可以通过在Kubernetes上进行部署来实现高可用性。Kubernetes具有称为复制器的功能，该功能监视运行中的Pod并自动启动崩溃的Pod。

第三，Kubernetes支持多租户。我们可以在同一Kubernetes平台上运行多个隔离的Fabric实例。这促进了区块链应用程序的开发和测试。

在以下各节中，我们介绍一种在Kubernetes上部署Fabric的方法。我们假设读者具有Fabric，Docker容器和Kubernetes的基础知识。

### **网络拓扑结构**

我们的网络拓扑如图1所示。物理网络用蓝线表示。Kubernetes具有一个或多个主节点和工作节点。除此之外，我们还有一台CMD计算机作为客户端来发布部署命令。NFS服务器用作配置文件和其他数据的共享文件系统。所有这些节点都通过物理网络（例如192.168.0.1/24）连接。

Kubernetes的网络模型使所有Pod都可以直接相互连接，无论它们位于哪个节点上。通过使用Kubernetes的CNI插件（例如Flannel），可以轻松地为此目的创建覆盖网络。如图1中的红线所示（省略了法兰绒组件的一些详细信息），Kubernetes将所有Pod连接到Flannel网络，从而允许这些Pod的容器彼此正确通信。

可以在附加配置文件中指定Flannel网络的IP地址范围以及kube_dns的IP地址。我们需要确保kube_dns的IP地址必须在指定的地址范围内。例如，在图1中，Flannel网络为10.0.0.1/16，而kube_dns地址为10.0.0.10。

![img](https://hackernoon.com/hn-images/0*XWJnYYVtYsV71Jwc.png)

图1

### **将结构组件映射到Kubernetes Pod**

![img](https://hackernoon.com/hn-images/0*RsyzLqCkvXB25BJb.png)

图2

Fabric是包含多个节点的分布式系统。这些节点可以属于不同的实体。如图2所示，每个组织都有自己的一组节点（为简单起见，未显示所有节点）。订购者还提供公共共识服务。要将Fabric部署到Kubernetes上，我们需要将所有组件转换为Pod进行部署，并使用命名空间隔离组织。

在Kubernetes中，名称空间是一个重要的概念。它用于在多个用户之间分配群集资源。对于Fabric，可以将组织映射到名称空间，以便拥有专用资源。完成此映射后，可以按域名区分每个组织的对等方。此外，我们可以通过设置网络策略来隔离不同的组织（本博客中未介绍）。

如图2所示，假设在Fabric网络中有N个对等组织和M个订购者组织。这是我们如何在Kubernetes上划分它们的方法：

**A）**如果是Fabric，我们为第N个对等组织分配名称*orgN*。在Kubernetes中其对应的命名空间也称为*orgN*。Fabric *orgN的*所有组件都将放入*Kubernetes*中的名称空间*orgN*中。每个组织的命名空间下都有多个Pod。吊舱是Kubernetes中的一个部署单元，它由一个或多个容器组成。我们可以将每个组织的Fabric容器捆绑到几个吊舱中。这些吊舱类型如下：

- **Peer Pod：**包括Fabric对等方，ouchDB（可选），代表组织的对等节点。每个组织可以有一个或多个对等豆荚。
- **CA Server Pod：**组织的Fabric CA Server节点。通常，组织中需要一个pod。
- **CLI Pod ：（**可选）为命令行工具提供一个环境来操纵组织的节点。Fabric的对等环境变量在此窗格中配置。

![img](https://hackernoon.com/hn-images/0*HUUjVZ7oox7AKu0k.png)

图3

**B）** Fabric中可能有一个或多个订购者。我们将第M个*订购*者组织的名称设置为*orgordererM*。它在Kubernetes上的对应命名空间是*orgordererM*。它具有一个或多个Pod以运行订购者节点。

![img](https://hackernoon.com/hn-images/0*w7rSoan-9fFdGbnn.png)

图4

**C）**如果将Kafka用于共识过程，我们可以将Kafka放在单独的命名空间中。它仅用于运行和管理Zookeeper和Kafka容器。

综上所述，总体部署如下所示：

![img](https://hackernoon.com/hn-images/0*3gDQxe7Rpj2zZSJD.png)

图5

### **共享存储**

在部署Fabric之前，我们需要准备其组件（例如对等方和订购者）的配置文件。这是一个非常复杂的过程，并且容易出错。幸运的是，我们创建了一个工具来自动生成这些配置文件。生成的文件存储在NFS之类的共享文件系统中。

以后启动Fabric的Pod时，我们将配置文件的不同子集安装到Pod中，以便它们具有特定于它们所属组织的配置。

在Kubernetes中，我们可以使用持久卷（PV）和持久卷声明（PVC）将文件或目录挂载到Pod中。我们为Fabric中的每个组织创建PV和PVC，以实现资源隔离。每个组织只能在NFS服务器中看到其自己的目录。

创建PV之后，我们定义PVC，以便Fabric节点可以使用PV来访问相应的目录和文件。

以对等组织org1为例。首先，我们创建一个名称空间org1及其PV。PV映射到NFS上的目录*/ opt / share / crypto-config / peerOrganizations / org1*。其次，我们创建PVC来消耗PV。命名空间org1下的所有Pod使用相同的PVC。但是，我们仅通过在Pod配置文件中指定安装路径将必需的文件映射到每个Pod中。

图6显示了Pod及其NFS共享目录之间的关系。变量$ PVC表示PVC安装点，在此示例中为*/ opt / share / crypto-config / peerOrganizations / org1*。

![img](https://hackernoon.com/hn-images/0*UpuaCjyWOaeS7hcG.png)

图6

### **Fabric组件之间的通信**

当所有Fabric组件都放入Kubernetes的Pod中时，我们需要考虑这些Pod之间的网络连接。Kubernetes中的每个Pod都有一个内部IP地址，但是由于IP地址是Pod临时的，因此很难使用IP和端口在Pod之间进行通信。当Pod重新启动时，其IP地址也将更改。因此，有必要在Kubernetes中为Pod创建服务，以便它们可以通过服务名称相互通信。服务的命名应遵循以下原则，以显示其绑定到的Pod信息：

**1）**服务的名称空间和pod应当一致。
**2）**服务名称应与容器在容器中的ID一致。

例如，组织org1的Fabric的peer0映射到名称空间org1下名为peer0的容器。与其绑定的服务应命名为peer0.org1，其中peer0是服务的名称，而org1是服务的名称空间。其他Pod可以通过服务名称peer0.org1连接到org1的peer0，该名称显示为peer0的主机名。

### **解决chaincode沙箱**

当Fabric中的同位体实例化一个链码时，它将创建一个在其中运行链码的Docker容器。它调用以创建容器的Docker API端点为unix：///var/run/docker.sock。只要对等容器和chaincode容器由同一Docker引擎管理，此机制就可以很好地工作。但是，在Kubernetes中，链码容器是由对等方创建的而不通知Kubernetes。因此，链码和对等容器无法相互连接，这在实例化链码时导致失败。

要解决此问题，我们需要在每个工作程序节点的Docker引擎中添加Kube_dns IP地址。这样可以确保链码容器可以使用Kube_dns服务正确解析对等方的主机名（服务名）。为此，请在Docker引擎的配置文件中添加以下选项，该文件在Ubuntu 14.04中通常为*/ etc / default /* docker。如果该文件不存在，则可能需要创建它。请注意，在最新版本的Docker和某些Linux Distro中，设置docker daemon选项的方式可能有所不同。有关更多详细信息，请参阅Docker文档。

在下面的示例中，10.0.0.10是kube_dns pod的IP地址。在您的环境中将其替换为正确的值。

```
DOCKER_OPTS=“ --dns = 10.0.0.10 --dns = 192.168.0.1 --dns-search \ 
default.svc.cluster.local --dns-search \ 
svc.cluster.local --dns-opt ndots：2 --dns -opt \ 
timeout：2 --dns-opt尝试次数：2“
```

到目前为止，我们已经说明了将Fabric部署到Kubernetes上的关键点。在下一篇文章中，我们将描述部署的详细步骤。对于迫不及待的人，请下载我们的Fling“ [**Kubernetes上的区块链**](https://labs.vmware.com/flings/blockchain-on-kubernetes) ”，以了解其工作原理。它是一种自动化工具，可让您以最少的配置在Kubernetes上部署Fabric。文件介绍了如何在vSphere上部署Kubernetes。如果不使用vSphere，则可以为Kubernetes实例选择任何基础架构。只需跳过在vSphere上部署Kubernetes的步骤。随时让我们知道您的想法。

**续第2部分：**

https://medium.com/@zhanghenry/how-to-deploy-hyperledger-fabric-on-kubernetes-2-751abf44c807

**相关文章：**

[**使用Helm Chart的Hyperledger Fabric部署**](http://www.think-foundry.com/hyperledger-fabric-deployment-using-helm-chart/)

### **关于作者：**

**Henry Zhang：** VMware中国研发部首席架构师，[Project Harbor](https://github.com/vmware/harbor)（[https://github.com/goharbor](https://github.com/vmware/harbor)）的创始人，[Project Harbor](https://github.com/vmware/harbor)是一个开源容器注册服务器。亨利是《**区块链技术指南**》一书的合著者。他还是[Hyperledger Cello项目](https://github.com/hyperledger/cello)的贡献者。推特：@zhanghaining

**Luke Chen：** VMware中国研发部工程师。他拥有广州大学的硕士学位。他是Hyperledger Cello项目的维护者。