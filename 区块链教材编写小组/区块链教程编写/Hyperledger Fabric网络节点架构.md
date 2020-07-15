# 深度探索Hyperledger技术与应用之超级账本的系统架构

区块链的业务需求多种多样，一些要求在快速达成网络共识及快速确认区块后，才可以将区块加入区块链中。有一些可以接受相对缓慢的处理时间，以换取较低级别的信任。各行各业在扩展性、可信度、合法性、工作流复杂度以及安全性等方面的需求和用途都不尽相同。我们先来看一下在企业级区块链系统中常见的模块构成，如图3-1所示。

![img](https:////upload-images.jianshu.io/upload_images/10818463-294690c76a8d6a84?imageMogr2/auto-orient/strip|imageView2/2/w/478/format/webp)

image







从图中可以看到一些常用的功能模块有：应用程序、成员管理、智能合约、账本、共识机制、事件机制、系统管理等。纵轴代表用户或者开发者更关心的内容，越往上代表用户更关注，比如应用程序和钱包等，越靠下是开发者更关注的模块，比如事件机制。而横轴则是从时间的维度来看的，左边是一开始关注的功能，直到完成所有的功能。

Hyperledger Fabric 1.0是一种通用的区块链技术，其设计目标是利用一些成熟的技术实现分布式账本技术（Distributed Ledger Technology，DLT）。超级账本采用模块化架构设计，复用通用的功能模块和接口。模块化的方法带来了可扩展性、灵活性等优势，会减少模块修改、升级带来的影响，能很好地利用微服务实现区块链应用系统的开发和部署。Hyperledger Fabric 1.0设计有几个特点：

**1）模块插件化：**很多的功能模块（如CA模块、共识算法、状态数据库存储、ESCC、VSCC、BCCSP等）都是可插拔的，系统提供了通用的接口和默认的实现，这满足了大多数的业务需求。这些模块也可以根据需求进行扩展，集成到系统中。

**2）充分利用容器技术：**不仅节点使用容器作为运行环境，链码也默认运行在安全的容器中。应用程序或者外部系统不能直接操作链码，必须通过背书节点提供的接口转发给链码来执行。容器给链码运行提供的是安全沙箱环境，把链码的环境和背书节点的环境隔离开，链码存在安全问题也不会影响到背书节点。

**3）可扩展性：**Hyperledger Fabric 1.0在0.6版本的基础上，对Peer节点的角色进行了拆分，有背书节点（Endorser）、排序服务节点（Orderer）、记账节点（Committer）等，不同角色的节点有不同的功能。节点可以加入到不同的通道（Channel）中，链码可以运行在不同的节点上，这样可以更好地提升并行执行的效率和吞吐量。

**4）安全性：**Hyperledger Fabric 1.0提供的是授权访问的区块链网络，节点共同维护成员信息，MSP（Membership Service Provider）模块验证、授权了最终用户后才能使用区块链网络的功能。多链和多通道的设计容易实现数据隔离，也提供了应用程序和链码之间的安全通道，实现了隐私保护。

# 1系统逻辑架构

下图所示为Hyperledger Fabric 1.0设计的系统逻辑架构图。

![img](https:////upload-images.jianshu.io/upload_images/10818463-c9ca8c3c34c39c03?imageMogr2/auto-orient/strip|imageView2/2/w/470/format/webp)

image

上图所示的系统逻辑架构图是从不同角度来划分的，上层从应用程序的角度，提供了标准的gRPC接口，在API的基础之上封装了不同语言的SDK，包括Golang、Node.js、Java、Python等，开发人员可以利用SDK开发基于区块链的应用。区块链强一致性要求，各个节点之间达成共识需要较长的执行时间，也是采用异步通信的模式进行开发的，事件模块可以在触发区块事件或者链码事件的时候执行预先定义的回调函数。下面分别从应用程序和底层的角度分析应该关注的几个要素。

**1、应用程序角度**

**（1）身份管理**

用户注册和登录系统后，获取到用户注册证书（ECert），其他所有的操作都需要与用户证书关联的私钥进行签名，消息接收方首先会进行签名验证，才进行后续的消息处理。网络节点同样会用到颁发的证书，比如系统启动和网络节点管理等都会对用户身份进行认证和授权。

**（2）账本管理**

授权的用户是可以查询账本数据（ledger）的，这可以通过多种方式查询，包括根据区块号查询区块、根据区块哈希查询区块、根据交易号查询区块、根据交易号查询交易，还可以根据通道名称获取查询到的区块链信息。

**（3）交易管理**

账本数据只能通过交易执行才能更新，应用程序通过交易管理提交交易提案（Proposal）并获取到交易背书（Endorsement）以后，再给排序服务节点提交交易，然后打包生成区块。SDK提供接口，利用用户证书本地生成交易号，背书节点和记账节点都会校验是否存在重复交易。

**（4）智能合约**

实现“可编程的账本”（Programmable Ledger），通过链码执行提交的交易，实现基于区块链的智能合约业务逻辑。只有智能合约才能更新账本数据，其他模块是不能直接修改状态数据（World State）的。

**2.底层角度**

下面的内容是从Hyperledger Fabric 1.0底层的角度来看，如何实现分布式账本技术，给应用程序提供区块链服务。

**（1）成员管理**

MSP（Membership Service Provider）对成员管理进行了抽象，每个MSP都会建立一套根信任证书（Root of Trust Certificate）体系，利用PKI（Public Key Infrastructure）对成员身份进行认证，验证成员用户提交请求的签名。结合Fabric-CA或者第三方CA系统，提供成员注册功能，并对成员身份证书进行管理，例如证书新增和撤销。注册的证书分为注册证书（ECert）、交易证书（TCert）和TLS证书（TLS Cert），它们分别用于用户身份、交易签名和TLS传输。

**（2）共识服务**

在分布式节点环境下，要实现同一个链上不同节点区块的一致性，同时要确保区块里的交易有效和有序。共识机制由3个阶段完成：客户端向背书节点提交提案进行签名背书，客户端将背书后的交易提交给排序服务节点进行交易排序，生成区块和排序服务，之后广播给记账节点验证交易后写入本地账本。网络节点的P2P协议采用的是基于Gossip的数据分发，以同一组织为传播范围来同步数据，提升网络传输的效率。

**（3）链码服务**

智能合约的实现依赖于安全的执行环境，确保安全的执行过程和用户数据的隔离。Hyperledger Fabric采用Docker管理普通的链码，提供安全的沙箱环境和镜像文件仓库。其好处是容易支持多种语言的链码，扩展性很好。Docker的方案也有自身的问题，比如对环境要求较高，占用资源较多，性能不高等，实现过程中也存在与Kubernetes、Rancher等平台的兼容性问题。

**（4）安全和密码服务**

安全问题是企业级区块链关心的问题，尤其在关注国家安全的项目中。其中底层的密码学支持尤其重要，Hyperledger Fabric 1.0专门定义了一个BCCSP（BlockChain Cryptographic Service Provider），使其实现密钥生成、哈希运算、签名验签、加密解密等基础功能。BCCSP是一个抽象的接口，默认是软实现的国标算法，目前社区和较多的厂家都在实现国密的算法和HSM（Hardware Security Module）。

Hyperledger Fabric 1.0在架构上的设计具有很好的可扩展性，目前是众多可见的区块链技术中最为活跃的，值得区块链技术爱好者深入研究。

# 2**网络节点架构**

节点是区块链的通信主体，是一个逻辑概念。多个不同类型的节点可以运行在同一物理服务器上。有多种类型的节点：客户端、Peer节点、排序服务节点和CA节点。下图所示为网络节点架构图。

![img](https:////upload-images.jianshu.io/upload_images/10818463-a5c9d3f6a4acf69e?imageMogr2/auto-orient/strip|imageView2/2/w/559/format/webp)

image

接下来详细地解释图3-3所示的不同节点的类型。

**1.客户端节点**

客户端或者应用程序代表由最终用户操作的实体，它必须连接到某一个Peer节点或者排序服务节点上与区块链网络进行通信。客户端向背书节点（Endorser）提交交易提案（Transaction Proposal），当收集到足够背书后，向排序服务广播交易，进行排序，生成区块。

**2. Peer节点**

所有的Peer节点都是记账节点（Committer），负责验证从排序服务节点区块里的交易，维护状态数据和账本的副本。部分节点会执行交易并对结果进行签名背书，充当背书节点的角色。背书节点是动态的角色，是与具体链码绑定的。每个链码在实例化的时候都会设置背书策略，指定哪些节点对交易背书后才是有效的。也只有在应用程序向它发起交易背书请求的时候才是背书节点，其他时候就是普通的记账节点，只负责验证交易并记账。

图3-2所示的Peer节点还有一种角色是主节点（Leader Peer），代表的是和排序服务节点通信的节点，负责从排序服务节点处获取最新的区块并在组织内部同步。可以强制设置为主节点，也可以动态选举产生。

在图3-2 中还可以看到，有的节点同时是背书节点和记账节点，也可以同时是背书节点、主节点和记账节点，也可以只是记账节点。在后面的章节中，有的地方会用记账节点代表普通的Peer节点。

**3.排序服务节点**

排序服务节点（Ordering Service Node或者Orderer）接收包含背书签名的交易，对未打包的交易进行排序生成区块，广播给Peer节点。排序服务提供的是原子广播（Atomic Broadcast），保证同一个链上的节点接收到相同的消息，并且有相同的逻辑顺序。

排序服务的多通道（MultiChannel）实现了多链的数据隔离，保证只有同一个链的Peer节点才能访问链上的数据，保护用户数据的隐私。

排序服务可以采用集中式服务，也可以采用分布式协议。可以实现不同级别的容错处理，目前正式发布的版本只支持Apache Kafka集群，提供交易排序的功能，只实现CFT（Crash Fault Tolerence，崩溃故障容错），不支持BFT（Byzantine Fault Tolerance，拜占庭容错）。

**4. CA节点**

CA节点是Hyperledger Fabric 1.0的证书颁发机构（Certificate Authority），由服务器和客户端组件组成。CA节点接收客户端的注册申请，返回注册密码用于用户登录，以便获取身份证书。在区块链网络上所有的操作都会验证用户的身份。CA节点是可选的，可以用其他成熟的第三方CA颁发证书。

下期预告：深度探索Hyperledger技术与应用之超级账本的典型交易流程

![img](https:////upload-images.jianshu.io/upload_images/10818463-9050e3afb224cd20?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

image

深度探索区块链

Hyperledger技术与应用

区块链

张增骏，董宁，朱轩彤，陈剑雄 　著

本书由超级账本执行董事Brian Behlendorf领衔推荐，区块链一线落地实践团队、Hyperleger会员智链骨干团对撰写。深入讲解Hyperledger Fabric 1.0的架构、执行逻辑、核心功能实现、从零部署，并以票据案例为例，讲解具体开发实践，穿插开发所需的最佳实践和遇到的问题解决。



作者：宇宙永恒
链接：https://www.jianshu.com/p/8a8c33f4f434
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。





# Hyperledger Fabric 学习二：系统架构

# 1、功能架构如下图所示。

![img](https:////upload-images.jianshu.io/upload_images/13194828-3416e01bef98f91c.png?imageMogr2/auto-orient/strip|imageView2/2/w/738/format/webp)

image.png

从中可以看出包括三大组件：区块链服务（Blockchain）、链码服务（Chaincode）、成员权限管理（Membership）。

# 1.区块链服务

区块链服务提供一个分布式账本平台。一般地，多个交易被打包进区块中，多个区块构成一条区块链。区块链代表的是账本状态机发生变更的历史过程。

- 交易
   交易意味着围绕着某个链码进行操作。
   交易可以改变世界状态。
   交易中包括的内容主要有：
  - 交易类型：目前包括 Deploy、Invoke、Query、Terminate 四种；
  - uuid：代表交易的唯一编号；
  - 链码编号 chaincodeID：交易针对的链码；
  - 负载内容的 hash 值：Deploy 或 Invoke 时候可以指定负载内容；
  - 交易的保密等级 ConfidentialityLevel；
  - 交易相关的 metadata 信息；
  - 临时生成值 nonce：跟安全机制相关；
  - 交易者的证书信息 cert；
  - 签名信息 signature；
  - metadata 信息；
  - 时间戳 timestamp。

交易的数据结构（Protobuf 格式）定义为



```tsx
message Transaction {
    enum Type {
        UNDEFINED = 0;
        // deploy a chaincode to the network and call `Init` function
        CHAINCODE_DEPLOY = 1;
        // call a chaincode `Invoke` function as a transaction
        CHAINCODE_INVOKE = 2;
        // call a chaincode `query` function
        CHAINCODE_QUERY = 3;
        // terminate a chaincode; not implemented yet
        CHAINCODE_TERMINATE = 4;
    }
    Type type = 1;
    //store ChaincodeID as bytes so its encrypted value can be stored
    bytes chaincodeID = 2;
    bytes payload = 3;
    bytes metadata = 4;
    string uuid = 5;
    google.protobuf.Timestamp timestamp = 6;

    ConfidentialityLevel confidentialityLevel = 7;
    string confidentialityProtocolVersion = 8;
    bytes nonce = 9;

    bytes toValidators = 10;
    bytes cert = 11;
    bytes signature = 12;
}
```

- 区块
   区块打包交易，确认交易后的世界状态。
   一个区块中包括的内容主要有：
  - 版本号 version：协议的版本信息；
  - 时间戳 timestamp：由区块提议者设定；
  - 交易信息的默克尔树的根 hash 值：由区块所包括的交易构成；
  - 世界观的默克尔树的根 hash 值：由交易发生后整个世界的状态值构成；
  - 前一个区块的 hash 值：构成链所必须；
  - 共识相关的元数据：可选值；
  - 非 hash 数据：不参与 hash 过程，各个 peer 上的值可能不同，例如本地提交时间、交易处理的返回值等；
     注意具体的交易信息并不存放在区块中。

区块的数据结构（Protobuf 格式）定义为



```go
message Block {
    uint32 version = 1;
    google.protobuf.Timestamp timestamp = 2;
    repeated Transaction transactions = 3;
    bytes stateHash = 4;
    bytes previousBlockHash = 5;
    bytes consensusMetadata = 6;
    NonHashData nonHashData = 7;
}
```

一个真实的区块内容示例：



```json
{
    "nonHashData": {
        "localLedgerCommitTimestamp": {
            "nanos": 975295157,
                "seconds": 1466057539
        },
            "transactionResults": [
            {
                "uuid": "7be1529ee16969baf9f3156247a0ee8e7eee99a6a0a816776acff65e6e1def71249f4cb1cad5e0f0b60b25dd2a6975efb282741c0e1ecc53fa8c10a9aaa31137"
            }
            ]
    },
        "previousBlockHash": "RrndKwuojRMjOz/rdD7rJD/NUupiuBuCtQwnZG7Vdi/XXcTd2MDyAMsFAZ1ntZL2/IIcSUeatIZAKS6ss7fEvg==",
        "stateHash": "TiIwROg48Z4xXFFIPEunNpavMxnvmZKg+yFxKK3VBY0zqiK3L0QQ5ILIV85iy7U+EiVhwEbkBb1Kb7w1ddqU5g==",
        "transactions": [
        {
            "chaincodeID": "CkdnaXRodWIuY29tL2h5cGVybGVkZ2VyL2ZhYnJpYy9leGFtcGxlcy9jaGFpbmNvZGUvZ28vY2hhaW5jb2RlX2V4YW1wbGUwMhKAATdiZTE1MjllZTE2OTY5YmFmOWYzMTU2MjQ3YTBlZThlN2VlZTk5YTZhMGE4MTY3NzZhY2ZmNjVlNmUxZGVmNzEyNDlmNGNiMWNhZDVlMGYwYjYwYjI1ZGQyYTY5NzVlZmIyODI3NDFjMGUxZWNjNTNmYThjMTBhOWFhYTMxMTM3",
            "payload": "Cu0BCAESzAEKR2dpdGh1Yi5jb20vaHlwZXJsZWRnZXIvZmFicmljL2V4YW1wbGVzL2NoYWluY29kZS9nby9jaGFpbmNvZGVfZXhhbXBsZTAyEoABN2JlMTUyOWVlMTY5NjliYWY5ZjMxNTYyNDdhMGVlOGU3ZWVlOTlhNmEwYTgxNjc3NmFjZmY2NWU2ZTFkZWY3MTI0OWY0Y2IxY2FkNWUwZjBiNjBiMjVkZDJhNjk3NWVmYjI4Mjc0MWMwZTFlY2M1M2ZhOGMxMGE5YWFhMzExMzcaGgoEaW5pdBIBYRIFMTAwMDASAWISBTIwMDAw",
            "timestamp": {
                "nanos": 298275779,
                "seconds": 1466057529
            },
            "type": 1,
            "uuid": "7be1529ee16969baf9f3156247a0ee8e7eee99a6a0a816776acff65e6e1def71249f4cb1cad5e0f0b60b25dd2a6975efb282741c0e1ecc53fa8c10a9aaa31137"
        }
    ]
}
```

- 世界观
   世界观用于存放链码执行过程中涉及到的状态变量，是一个键值数据库。典型的元素为 [chaincodeID, ckey]: value 结构。

  为了方便计算变更后的 hash 值，一般采用默克尔树数据结构进行存储。树的结构由两个参数（numBuckets 和 maxGroupingAtEachLevel）来进行初始配置，并由 hashFunction 配置决定存放键值到叶子节点的方式。显然，各个节点必须保持相同的配置，并且启动后一般不建议变动。

  - numBuckets：叶子节点的个数，每个叶子节点是一个桶（bucket），所有的键值被 hashFunction 散列分散到各个桶，决定树的宽度；
  - maxGroupingAtEachLevel：决定每个节点由多少个子节点的 hash 值构成，决定树的深度。
     其中，桶的内容由它所保存到键值先按照 chaincodeID 聚合，再按照升序方式组成。

一般地，假设某桶中包括M个 chaincodeID，对于chaincodeID_i，假设其包括  N 个键值对，则聚合G_i 内容可以计算为：

![G_i = Len(chaincodeID_i) + chaincodeID_i + N + \sum_{1}^{N} {len(key_j) + key_j + len(value_j) + value_j}](https://math.jianshu.com/math?formula=G_i%20%3D%20Len(chaincodeID_i)%20%2B%20chaincodeID_i%20%2B%20N%20%2B%20%5Csum_%7B1%7D%5E%7BN%7D%20%7Blen(key_j)%20%2B%20key_j%20%2B%20len(value_j)%20%2B%20value_j%7D)

该桶的内容则为

![bucket = \sum_{1}^{M} G_i](https://math.jianshu.com/math?formula=bucket%20%3D%20%5Csum_%7B1%7D%5E%7BM%7D%20G_i)

# 2.链码服务（智能合约）

链码包含所有的处理逻辑，并对外提供接口，外部通过调用链码接口来改变世界观。

- 1、接口和操作
   链码需要实现 Chaincode 接口，以被 VP 节点调用。



```tsx
type Chaincode interface {

Init(stub *ChaincodeStub, function string, args []string) ([]byte, error)

Invoke(stub *ChaincodeStub, function string, args []string) ([]byte, error)

Query(stub *ChaincodeStub, function string, args []string) ([]byte, error)
}
```

链码目前支持的交易类型包括：部署（Deploy）、调用（Invoke）和查询（Query）。

- 部署：VP 节点利用链码创建沙盒，沙盒启动后，处理 protobuf 协议的 shim 层一次性发送包含 ChaincodeID 信息的 REGISTER 消息给 VP 节点，进行注册，注册完成后，VP 节点通过 gRPC 传递参数并调用链码 Init 函数完成初始化；
- 调用：VP 节点发送 TRANSACTION 消息给链码沙盒的 shim 层，shim 层用传过来的参数调用链码的 Invoke 函数完成调用；
- 查询：VP 节点发送 QUERY 消息给链码沙盒的 shim 层，shim 层用传过来的参数调用链码的 Query 函数完成查询。

不同链码之间可能互相调用和查询。

- 2、容器
   在实现上，链码需要运行在隔离的容器中，超级账本采用了 Docker 作为默认容器。
   对容器的操作支持三种方法：build、start、stop，对应的接口为 VM。



```csharp
type VM interface { 
  build(ctxt context.Context, id string, args []string, env []string, attachstdin bool, attachstdout bool, reader io.Reader) error 
  start(ctxt context.Context, id string, args []string, env []string, attachstdin bool, attachstdout bool) error 
  stop(ctxt context.Context, id string, timeout uint, dontkill bool, dontremove bool) error 
}
```

链码部署成功后，会创建连接到部署它的 VP 节点的 gRPC 通道，以接受后续 Invoke 或 Query 指令。

- 3、gRPC 消息
   VP 节点和容器之间通过 gRPC 消息来交互。消息基本结构为



```dart
message ChaincodeMessage {

 enum Type { 
    UNDEFINED = 0; 
    REGISTER = 1;
    REGISTERED = 2; 
    INIT = 3; 
    READY = 4; 
    TRANSACTION = 5; 
    COMPLETED = 6;
    ERROR = 7; 
    GET_STATE = 8; 
    PUT_STATE = 9; 
    DEL_STATE = 10;
    INVOKE_CHAINCODE = 11; 
    INVOKE_QUERY = 12; 
    RESPONSE = 13; 
    QUERY = 14; 
    QUERY_COMPLETED = 15; 
    QUERY_ERROR = 16; 
    RANGE_QUERY_STATE = 17;
 }

  Type type = 1; 
  google.protobuf.Timestamp timestamp = 2; 
  bytes payload = 3; 
  string uuid = 4;
}
```

当发生链码部署时，容器启动后发送 REGISTER 消息到 VP 节点。如果成功，VP 节点返回 REGISTERED 消息，并发送 INIT 消息到容器，调用链码中的 Init 方法。

当发生链码调用时，VP 节点发送 TRANSACTION 消息到容器，调用其 Invoke 方法。如果成功，容器会返回 RESPONSE 消息。

类似的，当发生链码查询时，VP 节点发送 QUERY 消息到容器，调用其 Query 方法。如果成功，容器会返回 RESPONSE 消息。

# 3.成员权限管理

通过基于 PKI 的成员权限管理，平台可以对接入的节点和客户端的能力进行限制。

证书有三种，Enrollment，Transaction，以及确保安全通信的 TLS 证书。

- 注册证书 ECert：用于用户身份验证的注册证书，颁发给提供了注册凭证的用户或节点，一般长期有效；
- 交易证书 TCert：用于交易签名的交易证书，颁发给用户，控制每个交易的权限，一般针对某个交易，短期有效。
- 通信证书 TLSCert：加密传输的TSL证书，控制对网络的访问，并且防止窃听。

![img](https:////upload-images.jianshu.io/upload_images/13194828-08da91e11ca18dbb.png?imageMogr2/auto-orient/strip|imageView2/2/w/902/format/webp)

image.png

# 概念术语

- Auditability（审计性）：在一定权限和许可下，可以对链上的交易进行审计和检查。
- Block（区块）：代表一批得到确认的交易信息的整体，准备被共识加入到区块链中。
- Blockchain（区块链）：由多个区块链接而成的链表结构，除了首个区块，每个区块都包括前继区块内容的 hash 值。
- Certificate Authority（CA）：负责身份权限管理，又叫 Member Service 或 Identity Service。
- Chaincode（链上代码或链码）：区块链上的应用代码，扩展自“智能合约”概念，支持 golang、nodejs 等，运行在隔离的容器环境中。
- Committer（提交节点）：1.0 架构中一种 peer 节点角色，负责对 orderer 排序后的交易进行检查，选择合法的交易执行并写入存储。
- Confidentiality（保密）：只有交易相关方可以看到交易内容，其它人未经授权则无法看到。
- Endorser（背书节点）：1.0 架构中一种 peer 节点角色，负责检验某个交易是否合法，是否愿意为之背书、签名。
- Enrollment Certificate Authority（ECA，注册 CA）：负责成员身份相关证书管理的 CA。
- Ledger（账本）：包括区块链结构（带有所有的可验证交易信息，但只有最终成功的交易会改变世界观）和当前的世界观（world state）。Ledger 仅存在于 Peer 节点。
- MSP（Member Service Provider，成员服务提供者）：成员服务的抽象访问接口，实现对不同成员服务的可拔插支持。
- Non-validating Peer（非验证节点）：不参与账本维护，仅作为交易代理响应客户端的 REST 请求，并对交易进行一些基本的有效性检查，之后转发给验证节点。
- Orderer（排序节点）：1.0 架构中的共识服务角色，负责排序看到的交易，提供全局确认的顺序。
- Permissioned Ledger（带权限的账本）：网络中所有节点必须是经过许可的，非许可过的节点则无法加入网络。
- Privacy（隐私保护）：交易员可以隐藏交易的身份，其它成员在无特殊权限的情况下，只能对交易进行验证，而无法获知身份信息。
- Transaction（交易）：执行账本上的某个函数调用。具体函数在 chaincode 中实现。
- Transactor（交易者）：发起交易调用的客户端。
- Transaction Certificate Authority（TCA，交易 CA）：负责维护交易相关证书管理的 CA。
- Validating Peer（验证节点、记账节点）：维护账本的核心节点，参与一致性维护、对交易的验证和执行。
- World State（世界观）：是一个键值数据库，chaincode 用它来存储交易相关的状态。

# 2、网络拓扑结构

![img](https:////upload-images.jianshu.io/upload_images/13194828-09bace88bad9229e.png?imageMogr2/auto-orient/strip|imageView2/2/w/964/format/webp)

image.png

从图中可以看出包含以下节点：客户端节点、CA节点、Peer节点、Orderer节点。

- 客户端节点（应用程序/SDK/命令行工具）
   客户端或应用程序代表由最终用户操作的实体，它必须连接到某一个Peer节点或者排序服务节点上与区块链网络进行通信。客户端向背书节点（Endorser Peer）提交交易提案(Proposal)，当收集到足够背书后，向排序服务节点广播交易，进行排序，生成区块
- Peer节点（Leader主节点、Anchor锚节点、Endorser背书节点、Committer记账节点）
   从上图中可以看出每个组织可以拥有一到多个Peer节点。每个Peer节点可以担任如下多种角色：
  - Endorser Peer（背书结点）
     所谓背书(Endorsement)，就是指特定peer执行交易并向生成交易提案( proposal )的客户端应用程序返回YES/NO响应的过程。
     背书节点是动态的角色，是与具体链码绑定的。每个链码在实例化的时候都会设置背书策略(Endorsement policy)，指定哪些节点对交易背书才有效。
     也只有在应用程序向节点发起交易背书请求时才成为背书节点，其他时候是普通的记账节点，只负责验证交易并记账。
  - Leader Peer（主节点）
     主节点负责和Orderer排序服务节点通信，从排序服务节点处获取最新的区块并在组织内部同步。可以强制设置，也可以选举产生。
  - Committer Peer（记账节点）
     负责验证从排序服务节点接收的区块里的交易，然后将块提交（写入/追加）到其通道账本的副本。记账节点还将每个块中的每个交易标记为有效或无效。
  - Anchor Peer（锚节点）
     在一个通道( channel )上可以被所有其他peer发现的peer，通道上的每个成员都有一个Anchor Peer(或多个Anchor peer 来防止单点故障)，允许属于不同成员的peer发现通道上的所有现有peer。
     注：每个Peer节点必定是一个记账节点，除记账节点外，它也可以担任其它一到多种角色，即某个节点可以同时是记账节点和背书节点，也可以同时是记账节点、背书节点、主节点，锚节点。

- Orderer（排序节点）
   排序服务节点接收包含背书签名的交易，对未打包的交易进行排序生成区块，广播给Peer节点。

  排序服务提供的是原子广播，保证同一个链上的节点为接收到相同的消息，并且有相同的逻辑顺序。

  排序服务独立于peer进程存在并且以先来先服务的方式对网络上的所有信道进行排序交易。排序服务旨在支持超出现有的SOLO和Kafka品种的可插拔实现。排序服务是整个网络的公共绑定; 它包含绑定到每个成员的加密身份材料。

- CA（可选）
   CA节点是fabric的证书颁发节点(Certificate Authority)，由服务器(fabric-ca-server)和客户端(fabric-ca-client)组成。

  CA节点接收客户端的注册申请，返回注册密码用于登录，以便获取身份证书。在区块链网络上所有的操作都会验证用户的身份。

  CA节点是可选的，也可以用其他成熟的第三方CA颁发证书。

Fabric系统是通过组织来划分的，每个组织内都包含承担不同功能的Peer 节点，每个Peer节点又可以担任多种角色。所有的组织共用一个统一的Orderer集群。

# 3、交易流程

![img](https:////upload-images.jianshu.io/upload_images/13194828-c9cc19d0a46b55eb.png?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

image.png

1. 应用程序客户端首先构建交易的预案，预案的作用是调用通道中的链码来读取或者写入账本的数据。应用端使用 Fabric 的 SDK 打包交易预案，并使用用户的私钥对预案进行签名。

   应用打包完交易预案后，接着把预案提交给通道中的背书节点（Endorser），调用证书服务（CA）。
    通道的背书策略定义了哪些节点背书后交易才能有效，应用端根据背书策略选择相应的背书节点，并向它们提交交易预案。

2. 背书（Endorser）节点收到交易预案后，首先校验交易的签名是否合法，然后根据签名者的身份，确认其是否具有权限进行相关交易。此外，背书节点还需要检查交易预案的格式是否正确以及是否之前提交过（防止重放攻击）。

   在所有合法性校验通过后，背书节点按照交易预案，调用链码。链码执行时，读取的数据（键值对）是节点中本地的状态数据库。
    需要指出的是，链码在背书节点中是**模拟执行**，即对数据库的写操作并不会对账本作改变，所有的写操作将归总到一个写入的集合（ Write Set ）中记录下来。

   在链码执行完成之后，将返回链码读取过的数据集（ Read Set ）和链码写入的数据集（ Write Set ）。读集和写集将在确认节点中用于确定交易是否最终写入账本。

3. 背书（Endorser）节点把链码模拟执行后得到的读写集（ Read-Write Set ）等信息签名后发回给预案提交方（应用端）。

4. 应用端在收到背书响应之后，检查背书节点的签名和比较不同节点背书的结果是否一致。
    如果预案是查询账本的请求，则应用端无需提交交易给排序节点。如果是更新账本的请求，应用端在收集到满足背书策略的背书响应数量之后，把背书预案中得到的读写集、所有背书节点的签名和通道号发给排序节点。

5. 排序（Orderers）节点在收到各个节点发来的交易后，并不检查交易的全部内容，而是按照交易中的通道号对交易分类排序，然后把相同通道的交易打包成数据块（ blob ）。

6. 排序（Orderers）节点把打包好的数据块广播给通道中所有的成员。
    数据块的广播有两种触发条件，一种是当通道的交易数量达到某个预设的阈值，另一种是在交易数量没有超过阈值但距离上次广播的时间超过某个特定阈值，也可触发广播数据块。两种方式相结合，使得排序过的交易可以及时广播出去。

7. 记账（Committer）节点收到排序节点发来的交易数据块后，逐笔检查区块中的交易。先检查交易的合法性以及该交易是否曾经出现过。然后调用 VSCC（ Validation System Chaincode ）的系统链码检验交易的背书签名是否合法，以及背书的数量是否满足背书策略的要求。

   接下来进行多版本并发控制 MVCC 的检查，即校验交易的读集（Read Set）是否和当前账本中的版本一致（即没有变化）。如果没有改变，说明交易写集（Write Set）中对数据的修改有效，把该交易标注为有效，交易的写集更新到状态数据库中。

   如果当前账本的数据和读集版本不一致，则该交易被标注为无效，不更新状态数据库。数据块中的交易数据在标注成“有效”或“无效”后封装成区块（block）写入账本的区块链中。

上述的交易流程中，采用了 MVCC 的乐观锁（ optimistic locking ）模型，提高了系统的并发能力。需要注意的是，MVCC 也带来了一些局限性。例如，在同一个区块中若有两个交易先后对某个数据项做更新，顺序在后的交易将失败，因为它的读集版本和当前数据项版本已经不一致（因为之前的交易更新了数据）。



作者：张凯_9908
链接：https://www.jianshu.com/p/59c486fea1f7
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。









# [Hyperledger Fabric网络节点架构](https://www.cnblogs.com/luckray/p/Hyperledger-Fabric-wang-luo-jie-dian-jia-gou.html)

## Fabric区块链网络的组成

![img](https://images2018.cnblogs.com/blog/1420619/201806/1420619-20180620215005197-1999129062.png)￼
区块链网络结构图

### 区块链网络组成

#### 组成区块链网络相关的节点

节点是区块链的通信主体，和区块链网络相关的节点有多种类型：客户端（应用）、Peer节点、排序服务（Orderer）节点、CA节点

##### 客户端（应用程序）节点

客户端必须连接到某一个Peer节点或者排序服务节点才可以与区块链网络通信。

##### Peer节点

Peer节点主要负责通过执行链码（chaincode）实现对账本的读写操作

所有的Peer节点都是记账节点（Committer），负责维护状态数据和账本的副本

部分Peer节点根据背书策略的设定会执行交易并对结果进行签名背书，充当了背书节点（Endorser）的角色。背书节点是动态的角色，每个链码在实例化的时候都会设置背书策略，指定哪些节点对交易背书后才是有效的。只有在应用程序向节点发起交易背书请求的时候该Peer节点才是背书节点，否则它就是普通的记账节点。

一个组织（其实是成员）在一个通道上可以有多个Peer节点，这时候为了提高通信效率，需要选举出来一个主节点（Leader Peer）作为代表和排序服务节点通信，负责从排序服务节点处获取最新的区块并在组织内部同步。

节点可以同时是背书节点和记账节点，也可以同时是背书节点、主节点和记账节点。

##### 排序服务节点（Orderer）

排序服务节点接收包含背书签名的交易，对未打包的交易进行排序生成区块，并广播给Peer节点

一个区块链网络中，只能有一组排序服务，这个排序服务是由多个排序节点组成的。

排序服务启动的时候需要一个整个网络的创世区块，该创世区块中包含了排序节点信息、联盟组织信息、共识算法类型、区块配置信息及访问控制策略。同时在排序服务启动时会创建系统通道，系统通道在网络中有且只有一个，系统通道对应存储了系统账本，其中的有网络排序服务的定义、联盟成员的定义，以及其他初始配置参数，系统通道主要作用就是创建其他通道。

##### CA节点

CA节点是可选的，它主要作为证书颁发机构，也可以用其他成熟的第三方CA颁发证书。

#### 区块链网络与通道

一个区块链网络可以有多个通道，其中系统通道只能有一个。
一个通道对应一个账本，也对应一个链。不同通道的账本和智能合约、策略都是隔离的。

# [Hyperledger Fabric各类节点及其故障分析](https://www.cnblogs.com/preminem/p/8729781.html)

## 1.Client节点

client代表由最终用户操作的实体，它必须连接到某一个peer节点或者orderer节点上与区块链网络通信。客户端向endorser提交交易提案，当收集到足够背书后，向排序服务广播交易，进行排序，生成区块。但是该节点的故障不会影响区块链网络的正常运行。

## 2.CA节点

CA节点是hyperledger 1.0的证书颁发机构，由服务器（fabric-ca-service）和客户端组件（fabric-ca-client）组成。CA节点接收客户端的注册申请，返回注册密码用于用户登入，以便获取身份证书，在区块链网络上所有的操作都会验证用户的身份。因此该节点的故障只会影响到用户的注册申请。

## 3.Orderer节点

orderer负责接收包含背书签名的交易，对未打包的交易进行排序生成区块，广播给peer节点。排序服务可以采用集中式服务（solo，不适合实际生产环境），也可以采用分布式协议（目前正是发布的版本只支持Apache Kafka集群，只能实现崩溃故障容错）。BFT（拜占庭容错）的排序服务会在1.x周期内发布。我们也可以为fabric写一个共识实现，一个共识插件需要实现consensus package中定义的Consenter和Chain接口，我们可以来研究已经针对这些接口构建的plugin（solo和kafka）为自己提供线索。

 

我们分析下现有的基于kafka的排序服务安全性。

首先不用担心排序服务里插入数据带来的安全问题，因为最终记账还需要经过多重检查，比如消息类型的检查和签名验证、记账节点对背书策略的验证和交易内容的验证。但是其他的恶意节点攻击的问题还是无法避免的，比如拒绝服务（丢掉交易）。

kafka集群至少要四个节点，可以保证在1一个节点宕机后还能继续提交交易和排序，并且创建新的通道。zookeeper节点数选择3、5、7都可以。奇数个节点可以避免脑裂，1个节点会存在单点问题，7个以上就太多了。

值得一提的是，官方的example的在kafka集群和orderer之间都是没有安全传输的，切不可直接投入生产环境。我们可以通过使用openssl为kafka集群生成秘钥，增加orderer和kafka的配置以使用基于TLS的安全传输。

## 4.Peer节点

首先所有的peer节点都是committer（记账节点），而又有可能担任的角色有endorser（背书节点）、Leader（主节点）、Anchor（锚节点）。

### Committer

记账节点使用基于Gossip的p2p数据分发，节点会定期跟其他节点交换信息。如果在这个过程中有节点发生故障，则会从存活的节点中删除这个节点的信息。对于故障节点，还会定时检查是否已经恢复，恢复存活的节点会更新到存活节点列表中。如果有新加入的节点，也能通过节点信息的交换获取到，添加到存活列表中，广播给其他节点。由于超级账本采用基于反熵的状态同步，每个节点周期性的和邻居节点交换保存的数据，然后对比本地数据和邻居节点所保存的数据，检查是否有缺失或者过期的数据，然后更新本地节点的数据为最新的数据，因此故障的节点重新连接到网络时会自动恢复数据。这些都能通过Gossip协议学习到，自动调整网络的拓扑结构，适应网络节点的变化，保证整个网络正常运行。并且协议能正确工作的概率不会因为错误数超过f（可靠的广播协议中存在一个f，错误数超过这个值就会出现异常，协议的可靠性等于不超过f个错误的概率）时就快速地降低。（优雅降级）

### Leader

主节点连接到排序服务，负责把接受到的批量区块转发给其他节点。因此主节点与排序服务的稳定连接至关重要。可以强制设置为主节点，也可以动态选举产生。主节点选举的用处是，判断在相同的组织中哪个节点可以作为代表连接排序服务。选举过程在Gossip层实现，一个节点启动的时候它先等网络稳定再开始参与主节点选举，一次主节点选举的有效时间是10s。这样可以有效避免强制设置主节点出现的发生故障无法分发区块的问题。

### Endorser

背书节点为动态的角色与具体的chaincode绑定，背书节点的故障对网络的影响取决于chaincode对应的背书策略，例如背书策略指定只要3个组织其中的2个组织的成员完成背书，该交易就是有效的，那么只有一个组织的成员节点出现故障对交易完成没有影响。

### Anchor

锚节点是在一个channel上可以被所有其他peer发现的peer，channel上的每个成员都有一个anchor Peer(或多个anchor peer 来防止单点故障)，允许属于不同成员的peer发现channel上的所有现有peer。锚节点的配置文件可以通过configtxgen工具生成。



# Hyperledger Fabric节点间的网络架构说明

### 0 导言



```cpp
  在上一讲[《Hyperledger Fabric的逻辑架构是什么样的？》](https://www.jianshu.com/p/2bb78ef729b5)中，我们介绍了Fabric的逻辑架构，就是整个的技术组成部分。从应用程序端来看，包括了SDK、API、事件，通过SDK、API、事件来对底层区块链进行操作：包括**身份管理、账本管理、交易管理、智能合约的部署和调用**，从底层区块链这一端来看，对外提供了以下服务：**成员管理服务、共识服务、链码服务、安全和密码服务**。fabric通过将各个部分分离成不同的模块，做到可插拔性、灵活扩展性。

  通过上一讲我们对fabric的逻辑架构有了一个整体的认识。接下来我们来看看fabric节点间的网络架构是什么样的？我们先来上一张图：
```

![img](https:////upload-images.jianshu.io/upload_images/13024096-92fd788e119a3522.png?imageMogr2/auto-orient/strip|imageView2/2/w/1074/format/webp)

image



```undefined
  从图中可以看出fabric包含以下节点：客户端节点、CA节点、Peer节点、Orderer节点。我们下面来详细介绍一下这些节点。
```

### 1 客户端节点



```undefined
  客户端或应用程序代表由最终用户操作的实体，它必须连接到某一个Peer节点或者排序服务节点上与区块链网络进行通信。客户端向背书节点（Endorser Peer）提交交易提案(Proposal)，当收集到足够背书后，向排序服务节点广播交易，进行排序，生成区块。
```

### 2 CA节点



```undefined
  CA节点是fabric的证书颁发节点(Certificate Authority)，由服务器(fabric-ca-server)和客户端(fabric-ca-client)组成。

  CA节点接收客户端的注册申请，返回注册密码用于登录，以便获取身份证书。在区块链网络上所有的操作都会验证用户的身份。

  CA节点是可选的，也可以用其他成熟的第三方CA颁发证书。
```

### 3 Peer节点



```undefined
  从上图中可以看出每个组织可以拥有一到多个Peer节点。每个Peer节点可以担任如下多种角色：
```

- Endorser Peer（背书结点）
- Leader Peer（主节点）
- Committer Peer（记账节点）
- Anchor Peer（锚节点）
   *注：每个Peer节点必定是一个记账节点，除记账节点外，它也可以担任其它一到多种角色，即某个节点可以同时是记账节点和背书节点，也可以同时是记账节点、背书节点、主节点，锚节点。*

#### 3.1 Endorser Peer（背书结点）



```objectivec
  部分节点会执行交易并对结果进行签名背书，充当背书节点的角色 。

  所谓背书(Endorsement)，就是指特定peer执行交易并向生成交易提案( proposal )的客户端应用程序返回YES/NO响应的过程。

  背书节点是动态的角色，是与具体链码绑定的。每个链码在实例化的时候都会设置背书策略(Endorsement policy)，指定哪些节点对交易背书才有效。

  也只有在应用程序向节点发起交易背书请求时才成为背书节点，其他时候是普通的记账节点，只负责验证交易并记账。
```

#### 3.2 Leader Peer（主节点）



```undefined
  从图中可以看出，主节点负责和Orderer排序服务节点通信，从排序服务节点处获取最新的区块并在组织内部同步。可以强制设置，也可以选举产生。
```

#### 3.3 Committer Peer（记账节点）



```undefined
  负责验证从排序服务节点接收的区块里的交易，然后将块提交（写入/追加）到其通道账本的副本。记账节点还将每个块中的每个交易标记为有效或无效。
```

#### 3.4 Anchor Peer（锚节点）



```undefined
在一个通道( channel )上可以被所有其他peer发现的peer，通道上的每个成员都有一个Anchor Peer(或多个Anchor peer 来防止单点故障)，允许属于不同成员的peer发现通道上的所有现有peer。
```

### 4 Orderer（排序服务节点）



```undefined
 排序服务节点接收包含背书签名的交易，对未打包的交易进行排序生成区块，广播给Peer节点。

 排序服务提供的是原子广播，保证同一个链上的节点为接收到机同的消息，并且有相同的逻辑顺序。

 排序服务独立于peer进程存在并且以先来先服务的方式对网络上的所有信道进行排序交易。排序服务旨在支持超出现有的SOLO和Kafka品种的可插拔实现。排序服务是整个网络的公共绑定; 它包含绑定到每个成员的加密身份材料。
```

### 5 总结



```undefined
 Fabric系统是通过组织来划分的，每个组织内都包含承担不同功能的Peer 节点，每个Peer节点又可以担任多种角色。所有的组织共用一个统一的Orderer集群。因此在设计基于Hyperledger Fabric的系统时需要考虑组织之间的业务关系，以及内部每个模块之间的联系，以此来进行统一的规划。
```

作者：链播学院
 链接：https://www.jianshu.com/p/7b8841ef1a1e
 来源：简书
 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。





[区块链之超级账本]()





作者：小蜗牛爬楼梯
链接：https://www.jianshu.com/p/db95e9277010
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。







# Hyperledger Fabric 共识模型

# Orderers共识(Consenter)

## PBFT共识机制

Fabric 0.6采用的是**PBFT共识机制**，但已经暂时取消该机制改为了Kafka共识。原因如下：

- 交易性能达不到要求
- Fabric 面向的联盟链环境中，因为节点都是有准入控制的，拜赞庭容错的需求不是很强烈，反而是并发性能最重要

在测试开发中较多使用**SOLO共识机制**

> SOLO机制是一个非常容易部署的非生产环境的共识排序节点。由一个为所有客户服务的单一节点组成，所以不需要“共识”，因为有一个中央权威机构。相应地没有高可用性或可扩展性。这使得独立开发和测试很理想，但不适合生产环境部署。order-solo模式作为单节点通信模式，所有从peer收到的消息都在本节点进行排序与生成数据块

在Fabric 1.0往后版本中，采用的是**Kafka共识机制**，将来或采用SBFT（简化拜占庭容错共识）。

## 基于 Kafka 实现的共识

![img](https:////upload-images.jianshu.io/upload_images/14106155-991b038ef504e59c.png?imageMogr2/auto-orient/strip|imageView2/2/w/621/format/webp)

### Kafka共识

一个共识集群由**多个 orderer 节点（OSN）和一个 kafka 集群组**成。orderer 之间并不直接通信，他们仅仅和 Kafka 集群通信。

在 orderer 的实现里，通道(Channel)在 kafka 中是以 主题topic 的形式隔离。
 每个 orderer 内部，针对**每个通道都会建立与 kafka 集群对应 topic 的生产者及消费者**。生产者将 orderer 节点收到的交易发送到 kafka 集群进行排序，在生产的同时，消费者也同步消费排序后的交易。

### 如何鉴别某个交易属于哪个区块

Fabric 的区块结块由两个条件决定，**区块交易量**和**区块时间间隔**。

- 当配置的交易量达到阈值时，无论是否达到时间间隔，都会触发结块操作；
- 另一方面，如果触发了设置的时间间隔阈值，只要有交易就会触发结块操作，也就是说 fabric 中不会有空块。
- 结块操作是由 orderer 节点中的 kafka 生产者发送一条 TTC-X（Time to cut block x）消息到 kafka 集群，
- 当任意 orderer 节点的 kafka 消费者接收到任意节点发出的 TTC-X 消息时，都会将之前收到的交易打包结块，保存在 orderer 本地，之后再分发到各 peer 节点。

> 个人理解：OSN内部有每个通道对应topic的kafka生产者、kafka消费者，OSN里的生产者同时发送交易到集群，OSN同时消费来自集群的交易，共识的交易达到交易量或达到时间间隔后，OSN的kafka生产者会发送TTC-X到集群，OSN的kafka消费者收到后会将之前收到的交易结块，保存在OSN本地，然后分发到peers节点

以上过程可由下图描述：

![img](https:////upload-images.jianshu.io/upload_images/14106155-c0360a772e5a2baa.jpeg?imageMogr2/auto-orient/strip|imageView2/2/w/622/format/webp)

共识过程

### Kafka共识集群结构

#### Kafka简介

> 注，Orderers(OSN)与Consenter(一致性)不一样，Consenter在kafka集群里，是真正实现共识的共识插件，而OSN是与集群cluster通信获得共识结果的节点

**Kafka是一种分布式的，基于发布/订阅的消息系统**。主要设计目标如下：

- 以时间复杂度为O(1)的方式提供消息持久化能力，即使对TB级以上数据也能保证常数时间的访问性能
- 高吞吐率。即使在非常廉价的商用机器上也能做到单机支持每秒100K条消息的传输
- 支持Kafka Server间的消息分区，及分布式消费，同时保证每个partition内的消息顺序传输
- 同时支持离线数据处理和实时数据处理

#### Kafka架构

![img](https:////upload-images.jianshu.io/upload_images/14106155-c59309b55c134335?imageMogr2/auto-orient/strip|imageView2/2/w/781/format/webp)

Kafka架构



![img](https:////upload-images.jianshu.io/upload_images/14106155-615523d4cac745ec?imageMogr2/auto-orient/strip|imageView2/2/w/600/format/webp)

Kafka架构Partition

一个典型的kafka集群中包含若干producer，若干broker，若干consumer group，以及一个Zookeeper集群。

- Kafka通过Zookeeper管理集群配置，选举leader，以及在consumer group发生变化时进行rebalance
- producer使用push模式将消息发布到broker
- consumer使用pull模式从broker订阅并消费消息。

各个角色的功能分别是：

1. Brokers（经纪人）

   代理是负责维护发布数据的简单系统。

   - 每个代理可以每个主题具有零个或多个分区。
   - push模式的目标是尽可能以最快速度传递消息，但是这样很容易造成consumer来不及处理消息，而pull模式下Broker则可以根据consumer的消费能力以适当的速率消费消息。
   - 假设，如果在一个主题和N个代理中有N个分区，每个代理将有一个分区。

2. Producers（生产者）

   生产者是发送给一个或多个Kafka主题的消息的发布者。

   - 生产者向Kafka经纪人发送数据。
   - 每当生产者将消息发布给代理时，代理只需将消息附加到最后一个段文件。实际上，该消息将被附加到分区。
   - 生产者还可以向他们选择的分区发送消息。

3. Consumers（消费者）

   Consumers从经纪人处读取数据。 消费者订阅一个或多个主题，并通过从代理中提取数据来使用已发布的消息。

   - Consumer自己维护消费到哪个offet，线性增加
   - 每个Consumer都有对应的group
   - group内是queue消费模型：各个Consumer消费不同的partition，因此一个消息在group内只消费一次
   - group间是publish-subscribe消费模型：各个group各自独立消费，互不影响，因此一个消息被每个group消费一次

4. Topics（主题）

   属于特定类别的消息流称为主题。

   - 数据存储在主题中。Topic相当于Queue。
   - 主题被拆分成分区。
   - 每个这样的分区包含不可变有序序列的消息。
   - 分区被实现为具有相等大小的一组分段文件。

5. Partition（分区）

   一个Topic可以分成多个Partition，这是为了平行化处理。

   - 每个Partition内部消息有序，其中每个消息都有一个offset序号
   - 一个Partition只对应一个Broker，一个Broker可以管理多个Partition
   - 因为每条消息都被append到该partition中，是顺序写磁盘。因此效率非常高（经验证，顺序写磁盘效率比随机写内存还要高，这是Kafka高吞吐率的一个很重要的保证）

6. Replicas of partition（分区备份）

   - 副本只是一个分区的备份。 副本从不读取或写入数据。 它们用于防止数据丢失。
   - 对于传统的message queue而言，一般会删除已经被消费的消息，而Kafka集群会保留所有的消息，无论其被消费与否。
   - 当然，因为磁盘限制，不可能永久保留所有数据（实际上也没必要），因此Kafka提供两种策略去删除旧数据。一是基于时间，二是基于partition文件大小。

### Ledger账本结构

![img](https:////upload-images.jianshu.io/upload_images/14106155-a68ab0804b0c9797.jpeg?imageMogr2/auto-orient/strip|imageView2/2/w/720/format/webp)

#### 账本结构

Peer节点负责维护区块链的账本（ledger）和状态（State），本地的账本称为PeerLedger。
 **整个区块结构分为文件系统存储的Block结构和数据库维护的State状态**，其中state的存储结构是可以替换的，可选的实现包括各种KV数据库（LEVELDB，CouchDB等）

**账本简单的说，是一系列有序的、不可篡改的状态转移记录日志**

- 状态转移是链码（chaincode）执行（交易）的结果，每个交易都是通过增删改操作提交一系列键值对到账本。一系列有序的交易被打包成块，这样就将账本串联成了区块链。
- 同时，一个状态数据库维护账本当前的状态，因此也被叫做世界状态。

在 1.0 版本的 Fabric 中，**每个通道都有其账本，每个 peer 节点都保存着其加入的通道的账本，包含着交易日志（账本数据库）、状态数据库以及历史数据库。**

账本状态数据库实际上存储的是所有曾经在交易中出现的**键值对的最新值**

- 调用链码执行交易可以改变状态数据，为了高效的执行链码调用，所有数据的最新值都被存放在状态数据库中。
- 就逻辑上来说，状态数据库仅仅是有序交易日志的快照，因此在任何时候都可以根据交易日志重新生成。
- 状态数据库会在peer节点启动的时候自动恢复或重构，未完备前，该节点不会接受新的交易。
- 状态数据库可以使用 LevelDB 或者 CouchDB。
  - 跟 LevelDB 一样，CouchDB 也能够存储任意的二进制数据，
  - CouchDB 额外的支撑 JSON 富文本查询，如果链码的键值对存储的是 JSON，那么可以很好的利用 CouchDB 的富文本查询功能。

Fabric 的账本结构中还有一个**可选的历史状态数据库**，用于查询某个 key 的历史修改记录，需要注意的是，**历史数据库并不存储 key 具体的值，而只记录在某个区块的某个交易里**，某 key 变动了一次。后续需要查询的时候，根据变动历史去查询实际变动的值。

账本数据库是基于文件系统，将区块存储于文件块中，然后在 LevelDB 中存储区块交易对应的文件块及其偏移，也就是将 LevelDB 作为账本数据库的索引。
 现阶段支持的索引有：

- 区块编号
- 区块哈希
- 交易 ID 索引交易
- 区块交易编号
- 交易 ID 索引区块
- 交易 ID 索引交易验证码

### Chaincode链码系统

前面我们提到，状态转移是链码（chaincode）执行（交易）的结果。实际上，链码（chaincode）是 Hyperledger Fabric 提供的智能合约，是上层应用与底层区块链平台交互的媒介。即，合约与交易都与链码相关。

#### CSCC配置系统

介绍一个新的系统chaincode，叫做配置系统chaincode（CSCC），主要负责处理所有的配置相关的事务。CSCC提供方法查询众多的配置数据，包括通道配置。

#### 链码的通道配置

##### 引导

共识服务由1个或多个Orderers组成。 每个Orderer配置有匹配的创世区块，其由引导CLI命令生成，其提供了一些必要的数据，包括一系列可信根节点的列表，Order证书和IP地址的列表，一组特定的共识算法配置以及访问控制策略（谁可以创建信道）。

要启动并连接到共识服务，peer至少需要以下配置：

1. 准入网络的注册证书。 证书可以来自任意CA，只要CA是peer将连接到的共识服务的可信任根的一部分
2. 来自共识服务管理CLI生成的Orderer证书和IP地址的列表
3. 可信任根节点列表
4. peer可以订阅的通道可选列表。 除非明确配置，否则peer在启动时不订阅任何通道

注意，＃2和＃3来自引导的创世区块，我们可以从引导CLI命令获得。
 通过CLI或使用SDK API的应用程序，peer可以订阅已经存在的通道。 orderer通过在通道创建或重新配置期间收到的消息决定谁可以加入通道。

例如，假设peer A和B属于2个不同成员Alice和Bob。 请注意，Alice和Bob可能在网络上有多个Peer，并且他们的任何Peer都可以加入通道。 以下是一个典型的序列：

##### 创建通道

1. 应用程序/ SDK获得A和B的背书用于创建通道“foo”的配置交易。
2. 应用程序/ SDK调用Broadcast RPC，将背书过的配置交易传递给order服务。
3. 应用程序/ SDK然后调用在通道foo上deliver RPC。此RPC将返回一个错误，直到order服务成功创建通道。
4. 当通道最终被创建后，Deliver RPC将返回通道的信息到应用程序/ SDK。在这一时点，通道foo应当仅具有包含相关订阅者的创世区块，并且与该配置交易一起被（或最近的重新配置交易）引导。
5. 应用程序/ SDK在A和B上调用JoinChannel API，将通道foo的创世区块传递给A和B，添加CSCC到通道上。
6. A和B上的CSCC检查创世区块，包括检查区块中的配置交易的背书。如果一切正确，他们调用在通道上的Deliver RPC来开始接收块。
    如果通道已经存在，则参与者列表将被替换。Orderers自动替换订阅者并且将该交易与该通道上的其他交易一起发送给新成员，新成员将会同步完整的块。

##### 关闭通道

应用程序可以通过发送类似于创建通道的配置交易来关闭其创建的通道。 它需要根据应用程序设置的策略从通道参与方得到背书。

peer不会自动销毁相关的账本，但是裁剪进程会在适当的时候处理。
 应用程序可以继续从已关闭的账本中读取数据，只要该账本尚未被删除，但由于通道已被销毁，因此不能执行交易了。

##### 查询通道

通道只能被该通道的成员查询。也就是说，交易发起方的签名能够被存储在账本配置区块中的CA证书验证通过。这是通过发起一个查询交易到CSCC，同时附上链的ID，返回的结果是一个配置区块，里面包含了成员证书和一些其他的配置信息。

##### 链上的交易

一个交易必须包含目标的链ID（链ID =通道ID =账本ID）。
 共识服务将把交易放置在由链ID标识的指定通道上，并且在该通道内被排序，而与其它通道上的交易无关。 最终在该通道上产生一个包含交易的区块并发送到订阅了该通道的那些peer。

注意，每个链都是独立和并行的，因此一个peer可以同时接收和处理不同链上的区块。
 chaincode事务只能操作指定链中的状态变量

##### chaincode限制

1. 从交易调用chaincode总是在交易被发送的链上进行操作
2. 只有系统链上的chaincode可以被私有链上的其他chaincode调用并且是只读的

##### API 配置交易（通道与节点绑定）

1. 在peer上增加一个新的gRPC API和一个新的顶层交易类型。API允许App / SDK通知peer已成功加入的通道。
    **加入通道API的输入是由新创建的通道上的共识服务返回的创世区块**，peer使用此区块设置与通道关联的账本。
2. 新的交易类型称为配置交易，这种类型的交易可以由Orderer和peer处理。 创建或重新配置通道的交易都属于配置交易，其中背书请求是让peer批准和不批准它们创建或重新配置通道。 peer可以通过提案请求返回接受或拒绝。 为了保持灵活性，我们将提供一个系统chaincode来处理通道创建的背书请求，它将自动响应签名请求。
3. chaincode还提供查询此通道上参与成员列表的功能。
4. 新配置交易必须包含所有先前的配置条目，并且所有新/修改的配置条目必须用其包含配置包络的序列号和链ID标记。每个配置条目都具有枚举类型，唯一（按类型划分）ID以及由名称引用的修改策略。Order服务将根据现有配置策略验证配置交易，如果不满足全部修改策略，则拒绝它。
5. SDK可以向API提供进一步的抽象。 例如，它可以提供1个API，创建通道（成员证书列表），它将执行在创建通道部分中讨论的所有6个步骤。
6. 最后，SDK将调用应用程序上的回调，返回创建通道的状态。

#### 链码的合约作用

链码（chaincode）是 Hyperledger Fabric 提供的智能合约，是上层应用与底层区块链平台交互的媒介。现阶段，Fabric 提供 Go、Java 等语言编写的链码

所有的链码都实现两个接口，init 和 invoke。

- init 接口用于初始化合约，在整个链码的生命周期里，该接口仅仅执行一次。
- invoke 接口是编写业务逻辑的唯一入口，虽然只有一个入口，但是可以根据参数传递的不同自由区分不同业务逻辑，灵活性很高。比如应用开发者规定 Invoke 接口的第一个参数是合约方法名，剩余的 Invoke 参数列表是传递给该方法的参数，那么就可以在 Invoke 接口方法体中根据方法名的不同分流不同业务了。

##### 合约里能够获取的内容

- 输入参数获取。这点很好理解，我们只有知道此次调用的输入，才能处理逻辑，推导输出；
- 与状态数据库和历史数据库交互。在合约层，我们可以将区块链底层当做是一个键值对数据库，合约就是对数据库中键值的增删改查；
- 与其他合约的交互。在合约执行的过程中，可以与其他合约交换数据，做到类似跨链的效果。有了这种形式的数据获取方式，其实就可以将联系不紧密的业务逻辑拆分为多个合约，只在必要的时候跨合约调用，非常类似于现在提倡的微服务架构。

编写链码还有一个非常重要的原则：**不要出现任何本地化和随机逻辑**。此处的本地化，不是指语言本地化，而是执行环境本地化。区块链因为是去中心架构，业务逻辑不是只在某一个节点执行，而是在所有的共识节点都执行，如果链码输出与本地化数据相关，那么可能会导致结果差异，从而不能达成共识。

##### 链码部署

Peers是独立实体，通道就是业务载体，链码就是业务；不同的通道即便是运行相同的链码，因为载体不同，可认为是两个不同业务。

1. 创建业务载体通道；
2. 将通道与 peer 节点绑定；
3. 在通道上实例化链码。

##### 通道的管理

通道只有创建，而没有删除功能。但是在使用 kafka 共识的过程中，如果数据操作不当，直接在 kafka 中删除数据，而 orderer 没有逻辑去处理这种异常删除，因此会不断的重试，在达到重试极限后直接崩溃整个进程

##### 没有完善的数据管理方案

在我们的使用场景中，数据增长是很快的，如果使用 CouchDB 作为底层数据引擎，数据更是几何倍数的爆发。现有的解决方案只能是在云上部署节点，提供可持续扩充的云硬盘，再者使用 LevelDB 替换掉 CouchDB，避免使用模糊查询。

### Event事件流

事件框架支持发出2种类型的event(事件)

- Block event
- 自定义/chaincode event（在events.proto中定义的ChaincodeEvent类型）

##### 基本思想

client（event consumers\事件消费者）将注册event类型（block或chaincode）。**在chaincode的情况下，它们可以指定附加的注册标准，即chaincodeID和eventname。**

- ChaincodeID标识client想要查看event的特定Chaincode。
- eventname是Chaincode开发人员，在调用Chaincode中的SetEvent API时嵌入的字符串。

调用transaction是当前唯一可以发出event的操作，并且每个调用，在每个transaction中只能发出一个event。

### 一般Event类型与ChaincodeEvent的关系

Event与event类型相关联。 客户注册他们想要接收event的event类型。
 event类型的生命周期由“block”event来说明

1. 在启动peer时，在支持的event类型中添加“block”
2. client可以与peer（或多个peers）一起注册感兴趣的“block” event类型
3. 创建Block的Peers，向所有注册client发布event
4. 客户收到“block” event并处理Block中的事务

Chaincode event添加了额外的注册过滤级别。 Chaincode event不是注册给定event类型的所有event，而是允许client从特定Chaincode注册特定event。 对于目前的第一个版本，为了简单起见，没有在eventname上实现通配符或正则表达式匹配，但后续会提供该功能

### Fabric通信方式

节点通信、client与节点Api通信：**使用http/2下的gRPC**

- **http**: 基于TCP/IP协议，需要三次握手
- **rpc**：远程进程调用，需要统一的序列化，不适用于频繁连接
- **gRPC**: 使用HTTP/2协议并用ProtoBuf作为序列化工具

##### 与REST比较

- 和REST一样遵循HTTP协议(明确的说是HTTP/2)，但是gRPC提供了全双工流
- 和传统的REST不同的是gRPC使用了静态路径，从而提高性能
- 用一些格式化的错误码代替了HTTP的状态码更好的标示错误

##### gRPC

1. xxx.proto, 定义rpc，输入参数与返回参数的数据命名结构
2. 命令行中protoc编译生成对应的xxx.pb.go源码，编写clientAPI for EventService供客户端使用的接口定义、接口实例、接口实例的初始化函数，和server API for EventService供服务端使用的接口定义，注册函数。

##### HTTP/2特点

1. 将所有传输的信息分割为更小的消息和帧，并对它们采用二进制格式的编码。**在HTTP/2中，数据流以消息的形式发送，而消息由一个或多个帧组成，帧可以在数据流上乱序发送，然后再根据每个帧首部的流标识符重新组装**。二进制分帧是HTTP/2的基石，其他优化都是在这一基础上来实现的。我们先了解几个概念：

   - 帧（Frame）：HTTP/2通信的最小单位，每个帧包含帧首部，至少也会标识出当前帧所属的流。
   - 消息（Message）：由一个或多个帧组合而成，例如请求和响应。
   - 连接（Connection）：与 HTTP/1 相同，都是指对应的 TCP 连接；
   - 流（Stream）：已建立的连接上的双向字节流。

2. 支持请求与响应的多路复用来减少延迟

   - 同域名下所有通信都在单个连接上完成。
   - 单个连接可以承载任意数量的双向数据流。
   - 数据流以消息的形式发送，而消息又由一个或多个帧组成，多个帧之间可以乱序发送，因为根据帧首部的流标识可以重新组装。

   这一特性，性能会有极大的提升，因为：

   - 同个域名只需要占用一个TCP连接，消除了因多个TCP连接而带来的延时和内存消耗。
   - 单个连接上可以并行交错的请求和响应，之间互不干扰。

3. 压缩HTTP首部字段将协议开销降至最低

   - HTTP/2在客户端和服务器端使用“首部表”来跟踪和存储之前发送的键－值对，对于相同的数据，不再通过每次请求和响应发送；
   - 首部表在HTTP/2的连接存续期内始终存在，由客户端和服务器共同渐进地更新;
   - 每个新的首部键－值对要么被追加到当前表的末尾，要么替换表中之前的值。

4. 对请求划分优先级

5. 支持服务端Push消息到客户端

> ProtoBuf:
>
> 1. 一套用于数据存储，网络通信时用于协议编解码的工具库.它和XML和Json数据差不多,把数据已某种形式保存起来.Protobuf相对与XML和Json的不同之处，它是一种二进制的数据格式，具有更高的传输，打包和解包效率.
> 2. 如果使用protobuf实现，首先要写一个proto文件（不妨叫Order.proto），在该文件中添加一个名为"Order"的message结构，用来描述通讯协议中的结构化数据。使用protobuf内置的编译器编译 该proto。

参考资料：

Hyperledger文档：（详细研读）
 https://hyperledger-fabric.readthedocs.io/en/latest/whatis.html

Hyperledger github文档
 https://github.com/hyperledger/fabric/tree/release-1.1/docs/source

常用的Q&A
 https://hyperledger-fabric.readthedocs.io/en/latest/Fabric-FAQ.html?highlight=consenter

知乎专栏翻译：
 https://zhuanlan.zhihu.com/p/23356616

专栏-fabric生命周期
 https://zhuanlan.zhihu.com/p/25119939

hyperledger-MSP event等侧面分析 图解
 https://blog.csdn.net/maixia24/article/category/7507736



15人点赞



[区块链]()





作者：CodingCattwo
链接：https://www.jianshu.com/p/1d472586a3d5
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



#  [[Hyperledger\] Fabric系统中 peer模块的 gossip服务详解](https://www.cnblogs.com/X-knight/p/9909367.html)



**目录**

- [gossip——可最终达到一致的算法:](https://www.cnblogs.com/X-knight/p/9909367.html#_label0)
- [gossip中有三种基本的操作：](https://www.cnblogs.com/X-knight/p/9909367.html#_label1)
- gossip数据传播协议：
  - [以gossip为基础的数据传播协议在Fabric网络上执行三个基础的功能：](https://www.cnblogs.com/X-knight/p/9909367.html#_label2_0)
  - [gossip消息传送：](https://www.cnblogs.com/X-knight/p/9909367.html#_label2_1)

 

**正文**

   最近一直在看fabric系统中的核心模块之一——peer模块。在看peer的配置文件core.yaml的信息时，对其中的gossip配置选项很感兴趣。看了一上午，还是不能明白这个选项到底什么意思呢？表面意思很容易理解：“gossip”——“闲话”。但是在配置选项中为什么要起这么个名字呢？

   后来查阅了一些资料，才知道开发者用意何为？



## **gossip——可最终达到一致**的算法:

 

>    gossip本意是绯闻，流言蜚语，闲谈聊天的意思。而在这里，gossip代表了一种**可最终达到一致**的算法。其灵感来源于办公室八卦：当一个八卦在办公室出现时，在一定阶段内通过散播（dissemination），所有人最终都会知道这个八卦。这样就比较容易理解了，比如peer经过背书签名，将一个有效的交易最终提交。这份交易的写集合是A减100，B加100，因为网络中所有的结点都存有一份账本，因此该交易提交后，在有限的时间内，每个分布在网络中的结点中的账本都会应用这个交易，将自己的账本中的A减去100，B加上100。或者，有新结点加入网络中后，经过一定的时间，该新结点也会存储和其他结点一样的账本数据。
>
>    这里需要注意是，最终一致的另外的含义就是，不保证同时达到一致，也就是在某一指定的时刻，每个结点的账本（也就是状态）不保证一致。同时，gossip不要求节点知道所有其他节点，因此具有去中心化的特性，节点之间完全对等，不需要任何的中心节点，这点也是区块链的显著特征。



## gossip中有三种基本的操作：

 

- push - A节点将数据(key,value,version)及对应的版本号推送给B节点，B节点更新A中比自己新的数据。
- pull - A仅将数据key,version推送给B，B将本地比A新的数据（Key,value,version）推送给A，A更新本地。
- push/pull - 与pull类似，只是多了一步，A再将本地比B新的数据推送给B，B更新本地。



## gossip数据传播协议：

   Fabric通过划分各个执行交易的（背书和提交）peer和ordering结点之间的工作负载来优化区域链网络的执行，安全和可测量性。网络操作的解耦要求一个安全的，可信赖的，可测量的数据传播协议，以保证数据的完整性和机密性。为了达到这些要求，Fabric实现了一个gossip数据传播协议。

 

> **gossip协议：**
>
> peer结点“撬动”gossip以可测量的方式去广播（broadcast）账本和频道数据。gossip消息传送是是持续的，而且在频道中的每个peer不间断的从其它的peer那里接收当前的和一贯的（也就是格式等前后一致）账本数据（ledger data）。每个传播的消息都被签名过，因此“拜占庭的参与者”发送虚假的消息会很容易被识别，把消息发到消息不想到达的目标的分发行为会被阻止。peer会被延迟，网络参与者或者其他造成block丢失的原因所影响，但这些丢失block的peer最终将通过联系持有这些丢失block的peer异步更新到当前账本状态。

 



### 以gossip为基础的数据传播协议在Fabric网络上执行三个基础的功能：

 

- 通过持续性的识别有效成员peer和检测那些已经下线的peer，管理peer的发现（discovery）和频道成员关系。
- 在频道上所有的peer之间传播账本数据。任何持有与频道其他peer结点不同步的数据的peer识别丢失的block并通过拷贝正确的数据来同步自身。
- 通过允许账本数据以peer点对peer点（peer-to-peer）状态传输更新的方式，提高新加入网络的peer结点的同步速度。

   以gossip为基础的广播操作是通过peer从频道中其他peer中接收消息，然后把这些消息传送到频道上一定数量随机选择的peer结点，这个数量是可配置常量。peer结点也能运用一个pull机制，而不是等待一个消息的投递。这个循环重复着，伴随着频道成员关系的结果，账本和状态信息持续保持最新且同步。对于新block的传播，在频道中的领导peer（the leader peer）从ordering服务pull数据并初始化gossip到各个peer的传播。



### gossip消息传送： 

 

   在线的peer结点通过持续的广播“alive”信息来（向leader或其他结点）指示其自身的有效性，每条消息中都包含PKI_ID（the public key infrastructure ID）和发送者的签名。每个peer结点也通过收集这些“alive”消息，来维护自身的频道成员关系（channel membership）。如果没有任何一个peer接收到某一特定的peer的“alive”消息，则这个“dead”peer最终会从频道成员关系中被清除。因为“alive”消息都是加密签名了的，所以恶意的peer会因缺少由root CA认证的签名匙（signing key）而不可能冒充其他正确的peer。

 

   除了自动传输接收的消息（即散播dissemination），一个状态调节进程（state reconciliation process）会通过每个频道上的众多peer结点来与世界状态（world state）同步。每个peer持续性的从频道上的其他peer那里pull来block数据，目的在于，如果（通过与自己的block数据对比）存在差异则修复自身的状态。因为固定的连接（fixed connectivity）不被要求去维护以gossip为基础的数据散播，因此这个进程会可靠的提供私密的和完整的数据到共享的账本，同时包括了对错误结点的容错度。（这里的固定的连接应该这么理解：在网络中没有发生变化的结点集合，比如A，B，C，D四个结点一直没有发生变化，因而四个结点之间的关系也不会发生变化，因此这四个结点之间就不需要去进行gossip散播消息数据。比如一个新加入的E点是与D点发生关系，则只需要D去向E散播消息，而A，B，C，D四者之间仍是不需要互相进行gossip散播的。）

 

   因为多个频道之间是被相互隔离的，所有在一个频道上的peer点不能向其他频道发送消息或分享信息。虽然任一peer都可以属于多个频道，但是依照应用的消息线路选择策略（message routing policies），分配的消息传送禁止把block数据散播到其他频道的peer结点，这里的消息线路选择策略是以peer的频道订阅为基础的。（关于频道订阅，参考出版-订阅消息系统，即一个peer能够接收一个频道中的消息，必须先订阅这个频道的消息。）

 

**注意：**

   点对点（point-to-point）消息的安全性由peer的TLS层来处理，不需要签名。peer结点凭借它们自身的证书获得认证，这些证书由一个CA分配。虽然TLS证书也被使用，但是在gossip层是该peer点的证书被验证授权（而不是TLS的证书）。账本的block由ordering服务签名，然后投递到频道中的leader peer。
认证是由peer的MSP对象管理的。当一个peer第一次连接到频道上，TLS会话（session）同成员身份绑定。这主要是使用在网络和频道中的成员关系去认证每个与新的peer发生的连接。

 


\---------------------

**REFERENCE：**

1.*https://blog.csdn.net/idsuf698987/article/details/77898724* 