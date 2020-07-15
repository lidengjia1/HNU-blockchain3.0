# Hyperledger-Fabric-SDK-设计

## 概要

### 区块链网络使用

- gRPC 协议
- Protocol Buffers(格式的 API)

### 使用的协议

- gRPC

------

Protocol Buffers 格式的 API 包括

- 交易处理
- 安全的成员关系服务
- 区块链遍历
- 事件处理

------

### SDK 提供的 API 最小集合实现

#### 目标

- 要解决客户在区块链网络中直接面对的一些原始需求
- 做一些顶层的合理抽象以方便开发人员开发
- 在本地应用代码中使用这些 API 后方便开发工作的进行。

使用 SDK 并不会妨碍应用直接调用 gRPC 的操作。

**注意：** 当前的 REST API 并不被推荐使用，SDK 不应该基于使用 REST API 而创建。跳过 REST 直接使用 gRPC 的原因是：能够控制 全双向的数据流（大部分的 API 调用都是异步的），更好的性能。此外，自从节点内部模块通信使用了 gRPC 之后，便不必再开放更多的 HTTP/HTTPS 端口了

**例外的情况** 新的成员服务方式（CA）是按 REST API 的规则实现的

作为参考，Fabric v1.0的架构和交互模型可以在下面 [这个链接](https://link.jianshu.com?t=https%3A%2F%2Fjira.hyperledger.org%2Fbrowse%2FFAB-37) 所指向的工作单元的附件文档里面查看

------

## 目标

### 应用开发

为开发人员提供编写应用程序的多种操作区块链网络的方式。

- 部署/执行 chaincode
- 监听网络中产生的事件
- 接收块信息
- 把交易存储到账本中，等等

### chaincode 开发

可以为开发人员提供编写 chaincode 单元测试代码
 开发人员应该能够在不把 chaincode 部署到网络上的情况下快速测试 chaincode

## 原则

### 良好的 API 文档，数据模型，示例代码

提供清晰的文档说明

### 便于使用

便于使用

- chaincode 开发人员和应用程序开发人员关注于书写业务逻辑代码。
- SDK 编译时间上不应该对 fabric project 有任何依赖。（除非原始文件定义了多种合约？？**需要看英文原文**）
- SDK packages/jars/libraries 应该在常用的资源库网站上可获取 **引用的包需要常用且方便查找**

### 性能

- 高吞吐量
- 水平扩展能力
- 更低的延迟

SDK 应该是一个组件状态无关化的理想实现，或者允许应用实例数据库共享状态

### 版本控制

统一的版本控制下，建议将 SDK 通过多种语言实现，并且多种实现的 SDK 之间在功能上能够保持互通

### 可服务性

本 SDK 应该方便于在可服务性支持中添加插件，即日志系统。**具体细节不是很理解。**

它应该允许SDK使用方的应用添加一个日志处理的实例。**应用倾向于在SDK内部和外部使用一个通用的日志系统，这是非常有用的。**一个IT组织为了监视和分析通常会安装日志脚本，比如对一个标准日志格式的需求。本SDK应该有一个编译进版本的日志系统以便于开发人员通过默认方式获取日志。但是它必须允许设置一个拥有一系列标准集合的外部日志系统来记录API日志。

## 场景

有许多种可能场景，但是让我们关注一少部分来说明 SDK 的整体功能

### 成员的注册和登记



```undefined
需要有身份认证的证书，证书的提供者，证书可被验证
```

- 为了跟 fabric 交互，需要具备一个专有的身份认证，它被编码进 ECert (登记证书) 的成员证书里
- 证书可能是外部认证机构提供，或者由 fabric 成员服务提供
- ECert可以被fabric组件在建立链的信任链时验证即可

------

如果用一个标准的CA来管理分配ECert，交易会被ECert签名。这将使得每一个交易对于每一个客户端都是可验证的。

> TCert 怎么产生的？
>  MSP 在 fabric 中的设计是，利用一个加密算法从 ECert 的密钥对导出一个密钥对，因此为每一个交易产生一个假名证书

一个 TCert 不能代表一个认证的身份，但是能通过成员服务追踪到原始的 ECert。这个技术在许多业务场景下隐藏交易认证是非常有用的。

Fabric 提供了 MSP 接口的一种实现，命名为 “COP” 立意来源于它的作用就像警察，而不是来源于首字母。在github 的组织超级账本里面的代码库 [fabric-cop](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fhyperledger-archives%2Ffabric-cop) 就是它。（注：现在这个仓库在 github 上的 hyperledger-archives 目录下，目前此仓库的状态为只读状态，已经弃用）

给一个基于 fabric 的网络配置一个外部的 CA，用户注册也会在外部维护。这个外部的用户注册系统负责认证用户。一个已经被认证的用户为了获取 ECert 可以使用 CA 请求登记。

就 SDK 而言，应该有两种 API

- 支持任何 CA（包括内建的和外部的）通用 API
  - enroll(): 使用 SDK 的应用程序执行最基本的操作 比如关联用户身份，为认证用户获取 ECert 的程序
  - getTCert(): 获取交易证书允许用户提交交易
- 专门为 fabric 的内建成员服务提供的 API

[客户端有关可选的成员服务实现（即COP）的设计需要单独的文档来说明](https://link.jianshu.com?t=https%3A%2F%2Fdocs.google.com%2Fdocument%2Fd%2F1TO-sdHGpn8ifB9C_pH_y54H75Rx0RdXvZih7-lhsLII%2Fedit)

### 链的创建

Hyperledger Fabric 所设计的支持隐私性和机密性的通道和账本，统一被称为一条『链』。

[通道和账本设计更多内容请查看](https://link.jianshu.com?t=https%3A%2F%2Fdocs.google.com%2Fdocument%2Fd%2F1eRNxxQ0P8yp4Wh__Vi6ddaN_vhN2RQHP-IruHNUwyhc%2Fedit%23)

- 通信（传输中的数据）被限制在参与方之间的通道内
- 同时块内数据（已被持久化的数据）保存在一个私有的仅仅在这些参与方之间分布式存储的账本中。
- 不属于链的网络节点不了解通道和私有账本内部任何数据，也不能利用这些数据。

一旦一个链被创建，应用可以发送交易给链上以私有方式存在的节点，并且已提交被验证的交易到私有账本。**这句话不通**

**创建链的责任落在应用程序身上。**通过 SDK 应用程序初始化链的组织团体（网络成员的代表）构成并且给排序服务。**为什么是给排序服务**

- 在 SDK 里，通道和它本身关联的账本的绑定由链类体现。
- 应用和指定的排序节点第一次准备一个新的通道并且获取一个包含关于新链的密钥信息的创世区块，
  - 参与者信息（URL 和证书）、
  - 排序节点信息（URL 和证书）。

> Q: 应用通过什么来协调新通道与参与节点的邀请?
>
> A: 应用通过 目标是配置系统 chaincode 的配置化交易 来协调新通道对参与节点的邀请。（这句话还是不够直白，看英文自己理解吧~）

![img](https:////upload-images.jianshu.io/upload_images/107769-336b6ca8da86ca85.png?imageMogr2/auto-orient/strip|imageView2/2/w/1014/format/webp)

sdk-image01

### 交易支持

一个交易包含两个明确的步骤：

- 背书：是请求节点对交易的结果进行运行并背书（签名）
- 提交：是在交易和交易的背书内容的有效期内请求共识 **这个是什么意思**

下图说明交易时客户端SDK和节点之间的交互。更多信息请查看 [架构文档](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric%2Fblob%2Fmaster%2Fproposals%2Fr1%2FNext-Consensus-Architecture-Proposal.md)

流程简介：

首先 SDK 和应用协同工作，使用 **应用** 的（或者被认证用户的）**私钥给信息签名**（所有的外部信息都需要签名认证）。然后它根据 **背书策略**（通过验证系统链码或者VSCC实现的）给一个或多个节点 **发送信息**。它获得 **异步的反馈** 并且决定是否执行创建和提交交易到共识服务。执行交易提交的决定是基于背书策略谓词（就像 2 out of 3）基于带外知识。一旦提交，交易的处理流程便是异步的，所以SDK监听提交事件来通知应用交易处理流程的完成或者拒绝的结果。

![img](https:////upload-images.jianshu.io/upload_images/107769-c584d6ac731d23f3.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

image02

上图是对交易流程非常上层的描述。SDK 中有详细的执行网络和环境配置，包括 获取签名密钥来安全的管理成员，处理交易和事件流，（依赖应用的）多种共识通道。

## 客户端服务器 API 参考

下面的产指向 gRPC 与 fabric （节点，排序者和成员服务）沟通的信息和服务定义

- Proposal (部署或者交易调用的提案)
- ProposalResponse (全局的提案回应)
- Chaincode (chaincode, 调用，部署，等细节)
- ChaincodeProposal (chaincode-specific proposal header and payload)
- Transaction (transaction invocation)
- ChaincodeTransaction (chaincode-specific transaction action payload)

新的COP API是基于json的，需要添加到新的成员服务设计中。** COP 已弃用**

消息定义应该是设计 SDK APIs 的一个灵感来源。因为 SDK 可以使用 **智能缺省** 和 **状态信息** 来最小化请求的参数。

## 详细说明

这里讨论设计的原则和架构思路

总的来说，我们有不同等级的（数值越小表示等级越高）一些模块：

### package: Hyperledger Fabric Client

| 模块             | 等级 |                             功能                             |
| :--------------- | :--: | :----------------------------------------------------------: |
| Client           |  0   | 主要的入口模块。它必须允许用户创建需要的任何对象来执行所有支持的操作，例如直接连接网络，chaincode 部署，交易执行，多种查询。另外，基于编码规范和普遍的社区练习，每一种语言的实现也能决定是否添加方便的方法，如 sendTransaction(chain, tx) |
| Chain            |  1   | 一个链代表一些节点特别形成的一个网络，启动一个共识的通道，在通道中交易 可以被独立的处理。一个网络可能有一个或多个链。链上的节点维护一个单独的账本包含交易在链上派发，包括成员关系的任何配置。所有的交易都是在链上发送，一个应用可能操作多个链。 |
| Peer             |  2   | 代表网络上的计算节点。节点的角色有背书节点和提交节点，它们都在维护着账本。应用可能连接到一定数量的可用的节点。 |
| Orderer          |  2   | 类似节点，不同的是它代表排序服务的终端，可能是一个单独的节点（开发时本地安装）或者一个网络排序者的代理节点。基于区块链网络的 fabric 会有一个由多个排序者节点组成的单独的排序服务。应用可以选择信任特定的的排序者，或者一部分排序者，或者设置代理去给排序节点广播交易。 |
| User             |  2   | 代表在网络上交易的用户。用户实例可以基于登记证书被初始化。证书只可以从成员服务或者外部 CA 获取。理论上，这种用户也能代表网络上的节点成员。然而，这与应用程序无关（这更像是网络管理方面的问题），所以在这个设计中没有开放。 |
| Proposal         |  3   | 登记的用户可以向节点列表提出交易提案来背书交易。一旦接收到背书响应，应用程序可以决定是否已经获取背书签名，是否需要执行提交交易到共识服务。这是关于提案原始的 GRPC 消息的包装类，它提供了便利的创建方法。 |
| ProposalResponse |  3   | 提案调用背书节点的响应，打包背书结果（是或否），签名，等等。这是关于提案响应原始的 GRPC 消息包装类，它提供了便利的方法来利用它自己的内容（背书，签名，等等） |
| Transaction      |  3   | 登记用户收集了背书之后可以提交交易。交易请求包含背书签名和 MVCC + post-image, 并且使用排序服务 API。交易有两种类型：部署和执行。这是交易有关原始 GRPC 消息的包装类，它提供了便利的创建方法。 |
| CryptoSuite      |  3   | 加密模块打包了数字签名算法，非对称加密的密钥对，对称加密的密钥消息，安全的 hash 和 MAC。 |

Package: Member Service

|     模块      | 等级 |                             功能                             |
| :-----------: | :--: | :----------------------------------------------------------: |
| MemberService |  0   | 这是 fabric 可选模块的客户端，成员服务。本模块的主要功能是从成员服务获取用户登记证书。另外，这个模块本身或它的扩展类也应该能在 fabric 默认的成员服务的实现中提供可用的额外的功能，如用户注册功能等。 |

以上各模块的关系，我们给出了以下 UML 图：

![img](https:////upload-images.jianshu.io/upload_images/107769-d37b2fb0d64c141c.png?imageMogr2/auto-orient/strip|imageView2/2/w/934/format/webp)

sdk-image03

### 客户端

和终端用户主要的交互处理器。客户端实例提供一个和 网络上的节点，排序者，可选成员服务交互的处理器。应用程序使用 SDK 需要和多个网络交互，分别通过单独的客户端实例进行。

每个客户端被创建时，应该是使用来自于共识服务的配置数据初始化创建，这些数据包含一个被信任的根的列表，排序节点证书和IP地址，还有一个节点证书列表和可使用的IP地址。这必须是作为应用程序环境的一部分进行的。**应用程序负责维护客户端的配置，因为 SDK 不持久地保存这个对象**

- 共识服务的配置数据包括：
  - 一个被信任的根的列表
  - 排序节点证书
  - IP地址
  - 一个节点证书列表
  - 可使用的IP地址

每个客户端实例可以维护几条链代表通道和相关的账本。

- new_chain （创建一个新链）

根据参数给出的名字创建一个链的实例。这实际上代表“通道”（正如上面解释的），这个调用返回一个空对象。初始化这个通道，这个返回的对象上必须配置一个参与方的背书者列表和排序者节点。

Params (参数)



```rust
+ name(str): 链的名称，推荐使用命名空间防止碰撞
```

Return（返回值）



```undefined
+ （Chain instance）:未初始化的链的实例
```

- get_chain (获取链)

获取链的实例。本接口允许保存已存在的链实例供之后的检索，并且在应用程序实例之间共享。记录链的信息是应用程序或者 SDK 负责的事情。如果应用程序不能查看存储中链的信息，它可以调用另外一种 API 查询一个或多个节点来获得这些信息。

Params



```rust
+ name (str): The name of the chain （链的名称）
```

Returns



```python
+ (Chain instance or None): the chain instance for the name. （以入参命名的链的实例）
```

Error:



```bash
+ The state store has not been set (还没有设置存储状态)
+ A chain does not exist under that name (不存在该名称命名的链)
```

- query_chain_info（查询链的信息）

这是一个网络调用，用来查询指定的节点上链的信息。目标节点必须是属于目标链，才能够返回请求的信息。

Params



```cpp
+ name (str): The name of the chain （链名）
+ peers (array of Peer instances): target Peers to query（查询的目标节点）
```

Returns



```python
+ (Chain instance or None): the chain instance for the name.（以入参命名的链的实例。）
```

Error:



```undefined
+ The target Peer(s) does not know anything about the chain（目标节点不了解链的信息）
```

**有没有可能就没有这条链**

- set_state_store（设置状态的存储）

SDK 应该有一个内建的键值存储的实现（建议是基于文件的实现，以便于在开发中设置）。但是生产系统需要通过数据库为更多的稳定存储和聚簇存储备份，所以多种应用程序实例可以通过数据库共享应用状态（备注：应用不需要丰富的状态）。这个 API 使得存储模块具有可插拔特性，所以应用程序可以选择不同的存储实现。

Params



```csharp
+ store (KeyValueStore): instance of an alternative KeyValueStore implementation provided by the consuming app.（使用方应用程序提供的键值存储实现的实例）
```

Returns



```python
+ None
```

- get_state_store（获取状态存储）

为 client 提供的获取状态存储对象的便利方法

Params



```python
+ None
```

Returns
 \+ (KeyValueStore instance): The KeyValueStore implementation object set within this Client, or null if it does not exist（返回设置到client中的按键值对方式实现的存储对象，如果不存在返回空）

- set_crypto_suite (设置加密模块)

设置一个加密模块的实例，该实例是按 CryptoSuite 接口实现的。一个加密模块打包了数字签名的算法和使用非对称密钥对的加密功能，使用对称密钥加密的消息，安全的 hashing 处理和 MAC

Params



```csharp
+ Suite (object): an instance of a crypto suite implementation（按接口实现的加密模块的实例）
```

- get_crypto_suite (获取加密模块)

Client获取加密模块对象的便利方法。

Params



```python
+ None
```

Returns



```kotlin
+ (CryptoSuite instance): The CryptoSuite implementation object set within this Client, or null if it does not exist（加密模块的对象）
```

- set_user_context（设置用户上下文）

根据 client 实例的安全的上下文设置用户类的实例。用户的资格证书会被用来执行交易并且查询区块链网络。如果状态存储已经设置到了 client 实例上，根据设定的用户上下文，SDK 可以把对象保存在一个持久化的缓存中。如果没有设置状态存储，当应用程序崩溃或者被覆盖时，缓存不能被创建，并且应用程序需要再次设置用户上下文。

Params



```kotlin
+ user (User): an instance of the User class encapsulating the authenticated user’s signing materials (private key and enrollment certificate) （参数是用户，用户类的实例，它把被认证用户拥有的签名材料（私钥和背书证书）打包在内）
```

- get_user_context（获取用户上下文）

正如上面所解释的，client 实例可以拥有一个可选的状态存储。SDK 保存已注册用户到可以被应用程序的已认证用户（认证的工作在 SDK 之外由应用程序完成）利用的存储中。本方法试图通过本地存储中的名称（通过键值存储接口获取）加载用户。已加载的用户对象必须代表一个已注册用户，并且该用户拥有一个已信任的 CA（如 COP 服务）签名的可用的背书证书。

Params



```rust
+ name (str): The name of the user （用户名称）
```

Returns



```csharp
+ (User instance): The user object corresponding to the name, or null if the user does not exist or if the state store has not been set（返回匹配名称的用户对象，如果用户不存在或者状态存储未设置返回空）
```

### 链

“链”对象从通道获取设置项，由排序者节点创建，与排序者给通道上参与的节点派发交易的行为相隔离。根据节点列表和排序者列表配置链之后，它必须被初始化。**初始化过程给排序者节点发送一个配置交易来创建特定的通道，并且询问节点加入通道。**

- add_peer（添加节点）

给链对象添加节点，**这是纯本地操作。**

Params

peer (Peer): an instance of the Peer class that has been initialized with URL, TLC certificate, and enrollment certificate（入参：被 URL，TLC 证书，和 背书证书 初始化之后的 Peer 类的对象。）

- remove_peer (移除节点)

从链对象移除节点，这是一个纯本地操作

Params



```kotlin
+ peer (Peer): an instance of the Peer class（节点实例）
```

- get_peers (获取节点)

从链的本地信息获取节点。

Params



```python
+ None
```

Returns



```cpp
+ (Peer list): The peer list on the chain（返回链上的节点列表。）
```

- add_orderer（添加排序节点）

给链对象添加排序者节点，**这是纯本地操作。** 链实例可以选择使用单个的排序者节点，这个排序者负责向排序者网络中其他排序者广播请求。或者如果应用不信任排序者节点，它可以选择使用更多的排序者，仅需要向链对象添加它们即可。有关某个排序者的所有API会同时广播给所有的排序者。

Params



```kotlin
+ orderer (Orderer): an instance of the Orderer class （Orderer实例）
```

- remove_orderer（移除排序节点）

从链对象移除排序者节点，这是纯本地操作。

Params



```kotlin
+ orderer (Orderer): an instance of the Orderer class
```

get_orderers（获取排序节点）

获取链的排序节点，这是纯本地操作。

Params



```python
+ None
```

Returns



```cpp
+ (Orderer list): The orderer list on the chain
```

- initialize_chain（初始化链）

**调用排序者来开始创建新的链**，创建动作本身是开放的新的消息流和连接参与节点的组合。这是一个耗时的处理。只有一个应用程序实例需要调用这个方法。一旦链被成功创建，其他的应用程序实例仅仅需要调用 get_chain() 来获取有关链的信息。

Params



```python
+ None
```

Returns （链的初始化动作是否执行成功）



```cpp
+ (bool): whether the chain initialization process was successful
```

- update_chain（更新链）

**调用排序者节点来更新已经存在的链**。这允许给已存在的链添加或删除节点，也是基于证书重建对节点证书信息进行更新。

Params



```python
+ None
```

Returns （更新链的操作是否成功）



```cpp
+ (bool): whether the chain update process was successful
```

- is_readonly（是否只读）

获取链的状态来查看底层通道是否已经被终止，创建一个只读链，可以查询信息（交易和状态）但是不可以提交新的交易。

Params



```python
+ None
```

Returns （是否只读）



```csharp
+ (bool): is ready-only (true) or not
```

- query_info（查询信息）

查询链的状态（高度，已知的节点）中多种有用信息。

Params



```swift
+ none
```

Returns



```dart
+ (ChainInfo) with height, currently the only useful info (高度，当前有用的信息)
```

- query_block（查询块）

根据块号查询块。

Params



```tsx
+ blockNumber (number)
```

Returns



```dart
+ Object containing the block
```

- query_transaction（查询交易）

根据交易号查询交易。

Params



```undefined
+ transactionID
```

Returns



```undefined
+ TransactionInfo containing the transaction (交易包含的交易信息)
```

- create_deploy_proposal（创建部署提案）

创建交易提案。通过数据（chaincodeID，chaincode 调用，等）装配提案，并且使用匹配 ECert 签名的私钥签名提案。

Params



```tsx
+ chaincode_path (string): path to the chaincode to deploy

+ chaincode_name (string): a custom name to identify the chaincode on the chain

+ fcn (string): name of the chaincode function to call after deploy to

+ initiate the state（chaincode 的方法名，部署后初始化状态时要调用它）

+ args (string[]): arguments for calling the init function designated by “fcn”（上面 fcn 的参数）

+ sign (Bool): Whether to sign the transaction, default to True
```

Returns



```python
+ (Proposal): The created Proposal instance or None.
```

- create_transaction_proposal（创建交易提案）

为交易创建一个提案。通过数据（chaincode 名称，需要调用的方法，参数等）装配提案，并且使用匹配ECert签名的私钥给提案签名。

Params



```csharp
+ chaincode_name (string): The name given to the target chaincode to invoke（要执行的目标chaincode的名称）

+ args (string[]): arguments for calling the “invoke” method on the chaincode（正在调用执行的chaincode上的方法的参数）

+ Sign (Bool): Whether to sign the transaction, default to True
```

Returns



```python
+ (Transaction_Proposal instance): The created Transaction_Proposal instance or None.
```

- send_transaction_proposal（发送交易提案）

把创建好的提案发送给节点去背书。

Params



```kotlin
+ transaction_proposal (Transaction_Proposal): The transaction proposal data（交易提案的数据）

+ chain: The target chain whose peers the proposal will be sent to（提案将要发送给的节点所在的目标链）

+ retry (Number): Times to retry when failure, by default to 0 (no retry)（重试次数，默认0）
```

Returns



```undefined
+ (Transaction_Proposal_Response response): The response to send proposal request.（发送提案请求后的响应）
```

- create_transaction (创建交易)

遵从背书策略根据提案的响应信息创建交易。

Params



```cpp
+ proposal_responses ([Transaction_Proposal_Response]): The array of proposal responses received in the proposal call.（在提案的调用中返回的响应信息的数组）
```

Returns



```csharp
+ (Transaction instance): The created transaction object instance.（创建的交易对象实例）
```

- send_transaction（发送交易）

给链的排序服务（由一个或多个排序者节点组成）发送交易，然后做共识和提交到账本的工作。

**本调用是异步的**，并且交易提交成功信息是通过块或者 chaincode 事件通知的。（本方法必须给应用程序提供一个响应机制关联事件监听器处理“交易已提交”，“交易完成”，“错误”等事件。）

下面有两种有关 fabric 后端的沟通方式，触发不同的事件回调给应用程序处理器

排序者服务的 gRPC 客户端在“广播”调用的请求或响应方法中使用常规的无状态的 HTTP 连接。这个方法的实现需要在响应中接收到成功确认时发送“交易已提交”事件，或者在接收到错误时发送“错误”事件
 为了支持 fabric “BLOCK” ”, “CHAINCODE” 和 “TRANSACTION” 事件，这个方法的实现需要和作为内部事件枢纽机制一部分的链事件源节点维护一个持久的连接。这些事件应该引发方法给应用程序发送 “完成” 或 “错误” 事件。

Params



```csharp
+ transaction (Transaction): The transaction object constructed above（上面创建的交易对象。）
```

Returns（一个事件处理，可以关联到应用程序事件处理器的）



```csharp
+ result (EventEmitter): an handle to allow the application to attach event handlers on“submitted”, “complete”, and “error”.
```

### 用户

用户类代表已经通过 **注册证书注册** 和 **签名密钥签名** 的已登记用户。注册证书必须通过区块链网络配置信任的 CA 证书签名。已注册用户（拥有已签名密钥和注册证书）可以引导 chaincode 部署，交易和链上查询。

用户注册证书可以事先作为部署应用程序的一部分从 CA 获取，或者可以通过它本身的登记程序从可选的 fabric COP ( **已弃用** ) 服务获取。

有时用户认证和节点认证会混淆。因为用户可以利用私钥，所以用户身份有代理签名的能力，而节点在应用程序 /SDKs 上下文中只有验证签名的认证能力。**应用程序不能使用节点身份来签名因为应用程序不能利用节点身份的私钥**。

> 什么是用户认证，什么是节点认证？

> 用户可以利用私钥，所以用户身份有代理签名的能力
>
> 节点在应用程序 /SDKs 上下文中只有验证签名的认证能力

- get_name（获取名称）

获取成员名称。从对象实例请求参数。

Returns (str):



```undefined
+ The name of the user
```

- get_roles（获取角色）

  获取用户的角色。它可能是“客户端”“审计员”这类值的数组。成员服务多定义了两个的角色用来保存节点成员关系，如“节点”和“验证者”，这两个没有暴露给应用程序。

Returns (str[]):

The roles for this user

- get_enrollment_certificate（获取背书证书）

返回底层认证证书代表的用户身份证书。

Params:



```swift
+ none
```

Returns:



```cpp
+ Certificate in PEM format signed by the trusted CA（返回已信任的CA按PEM格式签名的证书）
```

- set_name（设置名称）

设置用户的名称/ID。

Params:



```cpp
+ name (string[): The user name / id.
```

- set_roles（设置角色）

按上面定义的角色值设置用户角色。

Params:



```cpp
+ Roles (string[]): The list of roles for the user（入参用户的角色列表。）
```

- set_enrollment_certificate（设置背书证书）

设置用户的背书证书。

Params:



```cpp
+ Certificate : The certificate in PEM format signed by the trusted CA（入参是被已信任的CA按PEM格式签名的证书）
```

- generate_tcerts（生成交易证书）

获取一系列 TCert，可以在交易中使用。TCert 和交易之间是一对一的关系。TCert 可以在本地通过 SDK 使用用户的加密数据项生成。

Params



```cpp
+ count (number): how many in the batch to obtain?（获取的交易证书的数量）

+ Attributes (string[]): list of attributes to include in the TCert（TCert包含的属性的列表）
```

Returns (TCert[]):



```cpp
+ An array of TCerts
```

### 节点

节点类代表了远程节点和它本身网络成员的数据，即用来验证签名的 ECert。**节点成员代表组织**，不像用户成员代表个体。

当节点被创建后，只要配置一个名为 “eventSourceUrl” 的属性，节点实例就可以被指定为一个事件源。允许 SDK 自动关联交易事件监听器事件流。

需要说明的是 **节点事件流功能在 节点层次，不在 链 和 chaincode 的层次**。

- connectEventSource（连接事件源）

由于几乎所有节点都是事件的生产者，当创建一个节点实例时，应用可以指定它作为应用程序的事件源。**只需要链上的一个节点成为事件源，因为链上的所有节点产生的事件相同**。本方法告诉 SDK 对于客户端应用程序来说哪一个节点作为事件源使用。管理与节点的 EventHub 连接的生命周期是 SDK 的责任。理解并通知选择的节点想接收哪种事件类型以及想使用哪个回调方法，是客户端应用程序的责任。

> SDK: 管理与节点的 EventHub 连接的生命周期
>  客户端应用程序：理解并通知选择的节点想接收哪种事件类型以及想使用哪个回调方法

Params:



```python
+ None
```

Result:



```jsx
+ Promise/Future: this gives the app a handle to attach “success” and “error” listeners（应用获得一个处理程序来关联“成功”或“错误”的监听器）
```

- is_event_listened（事件是否已被监听的标记）

网络调用可以显示出是否至少有一个监听器已经连接到事件的目标节点。这能帮助应用实例在崩溃后恢复中或者多个实例部署的情况下决定是否需要连接事件源。

备注：**这个请求对节点上事件的生成者有强化作用**。**不是很理解**

Params:



```swift
+ eventName (string): required （必需的）
+ chain (Chain): optional （可选的）
```

Result:（是否已经被链上的一些应用实例监听）



```csharp
+ (boolean): whether the said event has been listened on by some application instance on that chain
```

- addListener（添加监听器）

方法为连接到事件源的节点提供，监听器注册在 EventCallBack，用以接收事件类型集的回调。添加监听器的方法可以被执行多次来支持不同的 EventCallBack 方法接收不同类型的事件。

说明：以下的参数在某些语言里是可选的，比如 Java，为监听器接口创建一个实例，并把该实例作为参数。

Params:



```csharp
+ eventType : ie. Block, Chaincode, Transaction （事件类型）

+ eventTypeData : Object Specific for event type as necessary, currently needed for “Chaincode” event type, specifying a matching pattern to the event name set in the chaincode(s) being executed on the target Peer, and for “Transaction” event type, specifying the transaction ID （事件类型数据：事件类型有必要有一个特定的对象，对于“chaincode”事件类型要求是，在目标节点上执行的chaincode中设置一个指定的事件名匹配范式；对于“Transaction”事件类型，要指定交易ID。）

+ eventCallback : Client Application class registering for the callback.（事件回调：客户端应用程序类为回调而注册的。）
```

Returns:



```csharp
+ [event-listener-ref] a reference to the event listener, some language uses an ID (javascript), others uses object reference (Java)（返回值：事件监听器的引用，指向事件监听器的引用，有些语言（javascript）使用ID，其他的语言（Java）使用对象引用。）
```

- removeListener (移除监听器)

注销一个监听器

Params:（SDK返回事件监听器的引用）



```csharp
+ [event-listener-ref] : reference returned by SDK for event listener
```

Returns:



```undefined
+ statusFlag: Success / Failure
```

- get_name (获取名称)

获取节点名称。需要对象实例的参数。

Returns (str):



```undefined
+ The name of the Peer （Peer 的名称）
```

- set_name (设置名称)

Set the Peer name / id. （设置 Peer 的名称/id）

Params:



```cpp
+ Name (string): The unique name / id of this Peer.
```

- get_roles（获取角色）

获取节点参与的用户的角色。可能的结果是“client”和“auditor”的数组。成员服务多定义了两个角色来保存节点成员关系：“peer”和“validator”，这两个没有暴露给应用程序。

Returns (str[]):



```kotlin
+ The roles for this user (这个用户的角色)
```

- set_roles（设置角色）

设置节点参与的用户角色。值遵从上面预定的几个值。

Params:



```cpp
+ Roles (string[]): The list of roles for the user（用户角色列表）
```

- get_enrollment_certificate（获取背书证书）

返回底层代表用户身份的ECert。

Params:



```swift
+ none
```

Returns:



```cpp
+ Certificate in PEM format signed by the trusted CA
```

- set_enrollment_certificate（设置背书证书）

Set the Peer’s enrollment certificate. （设置 Peer 的背书证书）

Params:



```cpp
+ Certificate: Certificate in PEM format signed by the trusted CA （入参是被已信任的CA按PEM格式签名的证书）
```

### 键值存储（接口）

区块链应用程序需要保存状态，包含用户登记材料（私钥，CA签名的证书）。这些状态需要被持久化。“KeyValueStore” 的接口为 SDK 自动保存状态提供了一个简单的机制，这对应用程序是有益的。 如果应用使用基于软件密钥生成器即 CryptoSuite 的实现，那么它需要键值存储。如果应用程序还没有设置一个存储系统，SDK 应该默认使用一个内建的实现，比如一个基于本地文件系统的实现。

SDK 也可以在这个键值存储系统里面以可选缓存的方式保存用户登记材料。但是，如果应用程序没有配置一个键值存储系统，SDK 会把应用程序理解成选择了总是为会话设置上下文，并且没有意图使用默认的键值存储。

- set_value（获取值）

Retrieves a value given a key （获取一个入参的 key 的值）

Params



```rust
+ key (str): The name of the key
```

Returns

Result (Object):



```csharp
+ The value
```

- set_value（存入值）

Sets the value （存入值）

Params



```csharp
+ Key

+ value
```

Returns



```csharp
+ Acknowledgement of successful storage of the value （确认了值的成功存储）
```

### 加密模块（接口）

加密模块打包了 **数字签名算法** 和 **非对称加密方法**，**消息对称加密方法**，和 **hash** 以及 **MAC**。这是为区块链加密服务提供者接口的一个镜像设计，被 fabric 加密团队发表。

默认实现当前 peer 和 COP 的计划,并且还有 SDK 的默认实现：

- ECDSA: curves “secp256r1” and “secp384r1”
- AES: AES128, AES256 with CBC/CTR/GCM mode key length = 128 bits
- SHA: SHA256, SHA384, SHA3_256, SHA3_384
- generate_key（生成密钥）

基于这些选项生成一个 key。输出可以非对称加密算法的一个密钥对，或者对称加密算法的一个密钥

Params



```csharp
+ opts (Object): an object that encapsulates two properties, “algorithm” and“ephemeral”.
```

Returns



```csharp
+ Result (Key): The key object
```

- deriveKey (导出密钥)

Derives a key from k using opts.(从用上述选项中生成的 k 中获取 一个 key)

Params



```dart
+ k (Key)

+ opts (Object)
```

Returns



```undefined
+ (Key) derived key (导出 密钥)
```

- importKey（导入密钥）

Imports a key from its raw representation using opts.

Params



```dart
+ k (Key)

+ opts (Object)
```

Returns



```kotlin
+ (Key) An instance of the Key class wrapping the raw key bytes
```

- getKey（获取密钥）

Returns the key this CSP associates to the Subject Key Identifier ski .

Params



```csharp
+ ski (byte[])
```

Returns



```kotlin
+ (Key) An instance of the Key class corresponding to the ski
```

- hash（hash处理）

Hashes messages msg using options opts .

Params



```csharp
+ msg (byte[])

+ opts (Object) an object that encapsulates property “algorithm” with values for hashing algorithms such as “SHA2” or “SHA3”
```

Returns



```kotlin
+ (Key) An instance of the Key class corresponding to the ski
```

- encrypt（加密）

Encrypt plain text.

Params



```csharp
+ key (Key) public encryption key

+ plainText (byte[])

+ opts (Object)
```

Returns



```csharp
+ (byte[]) Cipher text
```

- decrypt（解密）

Decrypt cipher text.

Params



```csharp
+ key (Key) private decryption key

+ cipherText (byte[])

+ opts (Object)
```

Returns



```csharp
+ (byte[]) Plain text
```

- sign（签名）

Sign the data.

Params



```csharp
+ Key (Key) private signing key

+ digest (byte[]) fixed-length digest of the target message to be signed

+ opts (function) hashing function to use
```

Returns



```csharp
+ Result(Object):Signature object
```

- verify（验证）

Verify the signature.

Params



```csharp
+ key (Key) public verification key

+ signature (byte[]) signature

+ digest (byte[]) original digest that was signed
```

Returns



```cpp
+ (bool): verification successful or not
```

## 处理网络错误

客户端 SDK 和 fabric 用两个方法沟通：无状态的 HTTP 连接和 HTTP 长连接。

发送提案和发送交易调用在请求/响应中是无状态的。如果出现网络错误，调用会超时。SDK 应该有一个可配置的超时时间，方便应用程序基于对客户端应用和 fabric 节点之间的网络特征的理解控制调用的动作。

另外，这些方法可能有一个重试次数，以便于 SDK 根据超时错误自动试图重连 HTTP 调用。这些重试都尝试并且仍然发生超时错误之后，连接方法应该返回一个错误。

另一方面，事件流连接 SDK 和事件源节点是长连接。特别地，由于事件流接口被 fabric 定义，连接是双向的，允许消息发送到两方。对于网络错误，以防丢失事件和不能触发 client 应用注册的监听器的结果，连接会被断开。

为应用的利益考虑SDK应该试图重新发起连接。但是，如果在尝试了重连次数之后不能恢复连接，应该用一个高级别的严重错误通知应用这种情况。

## 参考

1. [Next-Consensus-Architecture_Proposal](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric%2Fblob%2Fmaster%2Fproposals%2Fr1%2FNext-Consensus-Architecture-Proposal.md)
2. [Consensus endorsing, consenting, and committing model](https://link.jianshu.com?t=https%3A%2F%2Fjira.hyperledger.org%2Fbrowse%2FFAB-37)
3. [Node.js SDK](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric-sdk-node)
4. [Fabric-Cop Design](https://link.jianshu.com?t=https%3A%2F%2Fdocs.google.com%2Fdocument%2Fd%2F1EznCYLJw3wKhANEm5l_SsPjYd5_rWNEpil8hTLmXXU8%2Fedit%3FcopiedFromTrash) **这篇文档已弃用**
5. [Next Hyperledger-Fabric Architecture Protocol messages](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fhyperledger%2Ffabric%2Fblob%2Fmaster%2Fproposals%2Fr1%2FNext-Consensus-Architecture-Proposal.md)



4人点赞



[Hyperledger Fabric translate]()





作者：简闻
链接：https://www.jianshu.com/p/bba58e9a281b
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。





# 关于HyperLedger Fabric基础内容

### 1、fabric V1.x的架构图

![img](https:////upload-images.jianshu.io/upload_images/5525735-7a77ba4adea9495d?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

HyperLedger Fabric



架构主要包括几个部分：

- 应用sdk：用于和区块链网络进行通信，提供了包括安全认证，交易申请等功能
- 节点：负责背书，验证，提交交易等功能，每个节点都维护了一个或是多个账本，同时通过gossip网络对其他节点保持通信。Fabric里每个节点都是无状态的，
- order服务，负责打包，排序和分发交易

##### 上图也简要概括了fabric一个交易的完整周期。

1. 客户端首先通过进行身份认证等安全操作，进入区块链网络
2. 客户端之后创建一个交易申请，发送给背书节点进行背书操作。
3. 背书节点执行对应的链码，基于应用的key操作生成读写操作集，并把背书结果返回给客户端
4. 客户端收到背书返回后，把交易发送给节点，由节点转发到order服务
5. order服务排序交易，把交易封装到区块里，并广播给每个节点（或是说每个提交节点）
6. 节点对交易进行背书策略验证，身份认证，区块中所有交易的有效性认证。之后写入账本，并返回交易结果

### 2、节点内交易流程

在Fabric中，为了方便共识模块化和数据访问权限的控制，引入了ordering服务和通道(Channel)的概念。所以Fabric中我们说的链路，其实包含了三部分：节点，通道和ordering服务，也就是说ordering服务和通道决定了哪些节点是在同一条链路上。



![img](https:////upload-images.jianshu.io/upload_images/5525735-6ce22e28812b2ad2?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

链路

> 上图包含的一条完整的链为：节点1.1，节点1.2，节点2.1，节点2.3。而节点1.3 和节点2.2不在这条链上。在节点 1.1，1.2，2.1，2.3中，都维护了同一份账本，这几个节点由于处在同一个通道和ordering服务中，所以他们之间的所有数据和消息都是共享和透明的，他们中的任何一个节点发送了状态变化，其他三个节点都能知道，但是这些信息对于节点1.3，2.2来说就是完全不透明的。同样的，节点1.3，2.2的状态变化或是账本信息变化对于处于链路上的节点(节点1.1，节点1.2，节点2.1，节点2.3)也是完全独立，不可见。比如，链路上有笔交易申请发生在节点1.1上，节点1.1会把这个交易申请发给ordering服务，ordering服务会根据发过来的请求判断这个请求是发生在哪个通道上，之后会把封装了这个交易的块发给这个通道上的其他所有节点。这个特性能让fabric在实际的应用更好的方便数据的访问控制。

接下来我们简单了解下链码(CC -- ChainCode)。链码是Fabric中的智能合约，相比于以太坊上用solidity写智能合约，链码可以直接采用java，go语言来编写，更容易上手。链码分为两种

- 系统链码(SCC)
- 用户链码

###### 系统链码

用来实现系统层面的功能，包括系统的配置，用户链码的部署、升级，用户交易的签名和验证策略等，运行在节点进程中。Fabric里有很多的系统链码，比如，用于背书的背书系统链码（ESCC），验证交易合法性的验证系统链码（VSCC）

##### 用户链码

用于实现用户的应用功能。开发者编写链码应用程序并将其部署到网络上。终端用户通过与网络节点交互的客户端应用程序调用链码，运行在独立的Docker容器中。

下面是一个节点内的交易流程图



![img](https:////upload-images.jianshu.io/upload_images/5525735-e6a838df25227641?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

image

这里的交易提交规则为：
 1.如果一个节点收到的2f（f为可容忍的拜占庭节点数）个其它节点发来的摘要都和自己相等，就向全网广播一条提交消息。
 2.如果一个节点收到2f+1条提交消息，即可提交新区块及其交易到本地的区块链和状态数据库。

### 3、交易背书

**什么是背书**
 背书可以理解为一种签名，签署行为。在fabric的意思是：告诉其他节点或是客户端，我这个节点模拟执行这个交易后的结果是什么，并附上自己的签名。客户端拿到到这个背书签名，就知道哪个节点执行了这个交易后的结果是什么。

> **不要把fabric里的背书行为当成是共识算法。这是错误的。**

本文将详细阐述交易背书的三个流程。
 **1.客户端创建一个交易背书申请并发送给背书节点**

我们这里说的客户端是指上链后的客户端，也就是说这个客户端一定连接到区块链的某个节点上的。为了发起一个交易，客户端需要给一组背书节点(客户端自己选择哪些背书节点)发送PROPOSE消息。那客户端是怎么知道背书节点的呢？在一条链路上，每个背书节点都会对外暴露部署在自己节点上的链码ID(chaincodeID), 这样通过交易中指明的chaincodeID，就可以找到所有部署了这个chaincodeID的背书节点。

PROPOSE消息格式为<PROPOSE,tx,[anchor]>,tx 是必须有的，anchor是可选。
 tx=<clientID,chaincodeID,txPayload,timestamp,clientSig>

> 其中clientID 是发起交易的客户端ID。chaincodeID这个交易所用到的链码ID。txPayload 交易申请的payload。timestamp  单调递增的整形数值，由客户端维护，一个交易对应一个。clientSig  对上述字段值的签名

anchor 版本相关信息，包含了读集合，更明确的是说一组key-version对(version是有序的版本号)。该KV对必须是KVS(key/value 存储，用于表示状态的一种数据结构)中的值。

在fabric中，有两种交易：部署交易和调用交易。部署交易是用于部署一个新的链码到一个区块链上。而调用交易是对已经部署在区块链上的链码的一个操作或是链码上一个方法调用。不同类型的交易，上面的payload也不同。

对于调用交易：
 txpayload = <operation,metadata>
 其中operation 值链码上的方法和参数。metadata 调用的相关属性。

对于部署交易：
 txpayload = <source,metadata,policies>
 其中source 链码的代码。metadata 链码和应用的相关属性。policies 包含了这个链码的相关功能，比如背书策略，当然txpayload里不会包含背书策略，而是指包含背书策略的ID和所需要参数。

每个交易都有tid,是全局唯一的，通过对tx进行加密hash计算得到。客户端会把tid存储在内存中，等待背书签名的返回。

由于PROPOSE消息格式中anchor是可选的，所以有两种方式发送PROPOSE消息给背书节点。一种是先发送<PROPOSE,tx> 到单个背书节点，这个背书节点会产生一个anchor，客户端拿到anchor后，之后可以发送<PROPOSE,tx,anchor>消息到其他的背书节点上。还有一种是直接发送<PROPOSE,tx>到所有的背书节点上。具体用哪种，由客户端自己选择。

**2.背书节点模拟一个交易并产生一个背书签名**
 当背书节点(id为epID)接收到从客户端发来的一条背书请求消息(<PROPOSE,tx,[anchor]>)，会先验证客户端的签名(clientSig)，然后模拟一个交易。所谓模拟一个交易就是利用交易上的链码ID(chaincodeID)调用相关链码，并且拷贝背书节点本地的状态信息。通过这种方式尝试的执行一个交易(txPayload).模拟交易执行的结果就是 背书节点计算得到 readset 和 writeset 两个集合。

状态信息包含键值对，并且所有的键值对都有版本控制。也就是说，每个键值对都包含了一个有序的版本信息，每当键值对的值发生变化时候，版本信息都会自动增加。背书节点通过链码把这个交易转换成所有的键值对，不管是读或是写，但是节点的状态信息在这个时候是没有更新的。更详细的为：

假定背书节点执行交易前的状态为s，对于交易的每个键k,读取这个键对应的值s(k).value，把 (k,s(k).value) 存于读集合(readset)中。

对每个被该交易更改后的键k所对应的新值v’, (k,v’)被加入到写集合(writeset). 当然也可以存储原来旧的值和当前新值的差值。

通过上述描述，我们知道背书节点会自己默默的计算并保存这个交易前后的状态集合变化，通过这种方式来模拟一次交易。这些操作对外是不可见的。

之后，节点会把内部交易请求(tran-proposal)转发到节点的背书逻辑，实现对交易的背书签名操作，默认情况下，节点的背书逻辑接收到交易请求后，只是简单的对这个交易请求进行签名。当然背书逻辑有可以进行额外的一些操作，比如，通过交易请求和tx 作为输入来判断是否要对这个交易进行背书。

如果背书逻辑决定了要对交易进行背书签名操作，背书逻辑会发送
 <TANSACTION-ENDORSED,tid,tran-proposal,epSig>

消息到对应的客户端(tx.clientID)：
 tran-proposal =<epID,tid,chaincodeID,txContentBlob,readset,writeset>
 epID节点ID。tid 交易ID。txContentBlob 指链码或是交易信息，txContentBlob =tx.txPayload。readset, writeset 模拟交易后的读写集合。 epSig 对tran-proposal 的背书节点签名。

如果背书节点拒绝对一个交易进行背书签名，背书节点会发送<TRANSACTION-INVALID,tid,REJECTED>到客户端。

> 注意，在整个背书过程，背书节点都是没有改变状态

**3.客户端收到背书结果并通过\**\**ordering\**\**服务广播给其他节点**
 客户端发送背书请求后，一直等待，直到客户端接收到了足够多的消息和在(TRANSACTION-ENDORSED,tid,*,*)状态上的签名之后，才认为交易已经经过背书签名操作。客户端具体需要多少消息才是足够多的消息呢？ 这个数字取决于链码的背书策略。如果收到的消息满足了背书策略，那么就认为交易进行过背书签名操作。注意，这里只是说交易被背书签名，而不是说交易被提交。从背书节点发过来的并且满足背书策略的签名 TRANSACTION-ENDORSED集合,就是 背书签名。如果客户端在这个过程中没有收到节点返回的有效背书签名，重试一定次数后还是失败就终止当前交易。

当收到节点返回的背书签名后，我们就开始使用ordering服务了。客户端通过使用broadcast(blob)的方式（blob为背书签名）来调用ordering服务。如果客户端没有直接调用ordering服务的能力，可以通过所连接的节点来代理广播背书签名给ordering服务。这个代理节点对于客户端来说是可信的，默认这个节点不会篡改背书签名或是伪造背书签名。否则这个交易就无效了。
 当事件deliver(seqno,prevhash,blob)发生，节点已经对所有低于seqno信息已经处理完成(seqno是有序递增的)，节点接下来：

- 根据该交易相关的链码(blob.tran-proposal.chaincodeID)上的背书策略，检查该交易背书签名(blob.endorsement)是否有效
- 同时验证背书签名相关性(blob.endordement.tran-proposal.readset)是否有被篡改。更复杂的验证方式也可以通过背书策略来验证 tran-proposal的字段是否有效
- 同时验证背书签名相关性(blob.endordement.tran-proposal.readset)是否有被篡改。更复杂的验证方式也可以通过背书策略来验证 tran-proposal的字段是否有效

验证背书签名相关性有多种实现方式，通过一致性或是状态更新的隔离保证都可以。串行化是默认的一种隔离保证方式，链码上的背书策略也可以指明了另外一种隔离保证方式。串行化可以通过要求在readset中每个key的版本号必须要等于状态中（KVS）的key的版本号来实现。如果不满足这个条件，就直接拒绝该交易。

> **需要注意，背书验证失败后，虽然会更新账本中失败交易的bitmask,但是不会更新区块链状态。**

这些节点处理好一个给定序列号的广播事件后，结果是这些节点都拥有一样的状态。也就是说，通过ordering服务的保证，所有正确的节点都会接受到来自ordering服务转发的同一个序列号事件(deliver(deqno,prevhash,blob))。由于背书策略和readset中的版本相关性验证都是确定的，所有的节点对于同一个交易处理结果也是一样的，不管这个交易是否有效。因此，所有的节点提交交易和更新状态的行为是一致的。
 下面用图来表示上面描述的一种通用交易流程。



![img](https:////upload-images.jianshu.io/upload_images/5525735-2258acd416dcb718?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

image



说明：背书节点是提交节点的一个子集，这里为了表示有提交功能，所以图示上显示了一个提交节点。其实所有的背书节点也都是提交节点，具有提交功能。但是如果一个节点不部署背书链码，那么它就不具有背书功能。

### 4、账本

账本提供了一个可证实的历史记录，它记录了对系统操作期间发生的所有成功交易和所有失败交易。账本由ordering服务生成，是一个完全排序的交易块（失败交易和成功交易）的哈希链。链上的每个节点都持有一份账本，部分ordering服务也可以持有账本。如果是ordering服务持有的账本，我们称为这种账本为orderer服务账本，而节点持有的账本称为节点账本，节点账号和orderer账号的区别在于节点账本上维持了一个位掩码来区分哪些交易是有效的，哪些是无效的。同时账本也允许节点回放所有的交易并且重新构造状态。

随着系统运行时间越来越长，无效交易也会随之变多，导致节点账本上存放了一堆无效交易，额外的增加了存储空间，如果这个时间点有新的节点加入到系统，在同步有效交易的同时，也会同步一大堆无效的交易，这也会导致新节点的同步时间增长，同时也会导致验证这些交易时间变长。于是，为了减少存储空间和新节点加入系统的时间和成本，fabric引入了有效账本的概念。

> 所谓的有效账本是指除了状态和账本外，节点持有一个逻辑账本，只包含有有效并且是已提交的交易。这个哈希链是通过过滤账本上的所有有效交易得到。
>  生成一个有效账本上的有效交易块的过程如下：

![img](https:////upload-images.jianshu.io/upload_images/5525735-ab629c5871c39425?imageMogr2/auto-orient/strip|imageView2/2/w/1080/format/webp)

image



当一个交易在变成有效交易块之前会判断该交易是否有效，如果是无效交易，就被剔除，如果是有效交易则加入进一个有效块(vBlock)中。所有的节点都会在本地进行这样的操作，比如通过使用节点账本的位掩码来过滤。一个有效交易块中不能包含无效交易块，所有无效交易都已经被剔除。这种有效交易块(vBlock)的长度是不固定的。每个节点上的有效交易块(vBlock)被连接起来成为一个哈希链，也就成了一个有效账本。有效账本的每个块包含：

- 前一个有效块的哈希
- 有效块的编号
- 从上一个有效交易块生成后到现在该节点提交的所有有效交易的有序列表
- 在该节点账本中，派生出当前有效交易块的交易块的哈希
   所有上述信息都被节点进行哈希索引
   账本包含有无效交易，虽然这些交易没有记录的必要，但是节点不会简单的就丢弃节点账本上的这些交易块。因此一旦节点账本生成了相关有效节点，就会对节点账本进行删减。也就是说，在这种情况下，如果有一个新节点计入到这个网络，其他节点就不会发送那些被剔除的交易块到这个新的节点上，也不需要新加入的节点去验证他们的有效交易块。那么节点什么时候生成有效账本呢？于是有了检查点机制，检查点机制通过检查点协议让节点知道什么时候生成有效交易块并去剔除无效的交易块。

检查点协议如下：
 节点对每个CHK块周期性地执行检查点，CHK是一个可配置参数。为了初始化一个检查点，节点通过gossip网络广播检查点消息到其他节点，检查点消息为：
 <CHECKPOINT,blocknohash,blockno,stateHash,peerSig>
 其中blockno 是当前交易块号，blocknohash为当前交易块的哈希，stateHash 为最近状态的哈希，peerSig 节点对检查点消息中其他字段的签名。

节点不停的接收检查点消息，直到有足够多的正确签名信息，可以通过这些信息中的blockno，blocknohash和stateHash建立一个有效的检查点。

对块号为blockno ,带有blocknohash的块创建了一个有效检查点后，节点首先检查blockno是否大于最新有效检查点的blockno,如果是，就把最新有效检查点的blockno改成blockno. 之后，存储由各自节点的签名组成了一个有效的检查点到latestValidCheckpointProof。同时存储stateHash响应的状态到latestValidCheckpointProof。最后可选的是否需要修剪节点账本。



作者：神奇的考拉
链接：https://www.jianshu.com/p/d33f414566d6
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。