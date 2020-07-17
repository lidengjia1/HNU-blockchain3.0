# Tendermint介绍及实战分析

# 1. 活动基本信息

**1）题目：**
 【技术工坊55期】Tendermint介绍及实战分析

**2）议题：**
 今年跨链技术十分火热，知名的跨链项目cosmos是基于Tendermint进行开发的，那么如何使用Tendermint开发一个自己的链呢，本期分享将为你揭晓~

**议题纲要：**

1. 怎么使用Tendermint开发自己的链
2. Tendermint共识算法
3. Tendermint ABCI接口介绍
4. 15分钟开发一个自己的链

**3）嘉宾：**

![img](https:////upload-images.jianshu.io/upload_images/1953055-2a9257dd812ba672.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/750/format/webp)

**4）活动定位**
 区块链技术工坊系列活动，由HiBlock，下笔有神科技，区块链兄弟，HPB芯链联合主办，聚焦于深度分享区块链知识，实现小会技术交友。

区块链技术工坊一直以来坚持4F原则：

- Frency - 每周三晚上一次；
- Focus - 聚焦区块链技术分享；
- Fun - 20人以内会前做自我介绍，分享有深度的技术内容，技术交友；
- Feedback - 会后有活动实录文章和合影照片，深度对接业务交流；

通过技术工坊，连接了广大区块链项目和开发者，搭建了技术交友和知识传播的平台。

# 2.会议实录

![img](https:////upload-images.jianshu.io/upload_images/1953055-3afb6db7fad74c51.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-e6ac8b878717cff8.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



我们将从理论解析和实战操作两个层面为大家介绍 Tendermint，可不要小瞧了这短短两行的目录，接下来的内容可是干货满满，精彩不断。



![img](https:////upload-images.jianshu.io/upload_images/1953055-5acb8aa7ca130906.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)


 提到区块链，大家想必已然不陌生了，不过更多人想到的可能会是众所周知的 Bitcoin 和 Ethereum。的确，两者分别是区块链技术的起源和发展的代表，也是大家广泛传播和深入研究的对象。但是随着 Bitcoin 的不断推进，比特币工作量证明共识机制在速度和扩展性上的不足也逐步展现出来。

Cosmos 的开发团队 Tendermint 其实早在2014年就意识到了其不足，并持续专注于寻求不依赖挖矿等高电力消耗的共识机制，提供快速的交易处理能力，它们的目标是为全世界所有的区块链提供速度、安全和可扩展性。目前，Tendermint 是跨链技术的核心。



![img](https:////upload-images.jianshu.io/upload_images/1953055-fedd66cabb117735.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-0d2e9cfb88524a29.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



Tendrmint 协议与比特币系统有很多相似之处，因为他们都是通过区块来记录信息，但在解决“拜占庭将军”的问题 (共识机制) 上又各采取了不同的解决方案。比特币的协议优化了去中心化的审核机制，而 Tendermint，则是在多节点的广域网（如拥有百万节点的高节点数）中优化了分布式应用及数据处理方面的拜占庭容错。上面两张图生动形象地介绍了 Tendermint 的软件开发层和应用部署网络结构。



![img](https:////upload-images.jianshu.io/upload_images/1953055-bdd970e5faa64cf6.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)


 那相比于老牌的区块链技术，Tendermint 到底拥有哪些开发优势呢？

目前大部分的区块链实现都是采用一种“网络-共识算法-应用”的框架，实现成单一的程序，但是这就很容易出现两个问题：

代码复用困难，代码库的分支管理变得复杂。

限制了应用开发的语言。

针对这两个问题，Tendermint 设计了自己的一套框架，其设计原则是易使用，易理解，高性能，适用于各种分布式应用。它的创新之处在于，将区块链应用（状态）与底层共识进行了分离，将共识引擎和 P2P 网络层封装组成 Tendermint Core。同时提供 ABCI 接口与应用层进行交互，应用逻辑可以用任何语言编写，应用做的事情实际上就是状态机控制。不仅如此，Tendermint 还提供了许多实用的开发库，如密码库，默克尔树，iavl 状态树等。基于这种架构，应用的开发者可以方便地实现自己的区块链。



![img](https:////upload-images.jianshu.io/upload_images/1953055-e1588207a2974d7d.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



接下来，我们将从共识算法，状态机复制及 ABCI 接口三个方面深入挖掘 Tendermint 的细节。



![img](https:////upload-images.jianshu.io/upload_images/1953055-a4100fc09df65827.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)


 下图为 Tendermint 出块节点选择机制及其优缺点。

![img](https:////upload-images.jianshu.io/upload_images/1953055-b418170c0b5081d2.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



参与 Tendermint 共识过程的角色主要有两个：

· Validator（验证者）：网络中的节点，可参与共识过程中的投票，不同的验证者在投票过程可能具备不同的投票权重（vote power）。

· Proposer（提议人）：Tendermint 使用一种确定的非阻塞轮询选择算法从 Validators 中选出 Proposer，该算法根据 Validators 的投票权重所占比例来选择Proposer，投票权重越大的 Validator 被选为 Proposer 的频率越高。

Tendermint 区块链是通过基于 round（回合）的机制来确定下一个区块。每个 round 由三个过程组成：propose（提议），prevote（预投票）和 precommit（预提交）。



![img](https:////upload-images.jianshu.io/upload_images/1953055-90831de015e56253.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



这样的机制优点在于 Proposer 的选择方式是与 Stake 相关的，所以应用层可以实现自己的共识（DPOS）,将应用层 Validator 的权重传递给 Tendermint，Tendermint 就会按照应用层需要的方式选择 Proposer。当然，没有一样技术是完美的，Tendermint 也有其出块机制上的缺点，那就是 Round-robin 策略过于简单，容易被坏人预测到下一个 Validator 是谁，于是可以提前布局，对 Validator 发起 DDoS 攻击或别的攻击。Tendermint 对此给出的解决方法是，把 Validator 节点，全部放在 Sentry Node 后面，对外不暴露 IP 地址。



![img](https:////upload-images.jianshu.io/upload_images/1953055-51aa1f1a0eef4ed9.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)


 Tendermint 是一个易于理解的 BFT 共识协议。此协议遵循一个简单的状态机：验证者轮流对交易的区块提议并对提议的区块进行投票。区块被提交到链上，且每个区块就是一个区块高度。但区块也有可能提交失败，这种情况下协议将选择下一个验证者在相同高度上提议一个新的区块块，并重新开始投票。

如上图所示，想要成功提交一个区块，必须经过两阶段的投票，称为预投票（pre-vote）和预提交（pre-commit）。当超过 2/3 的验证者在同一轮提议中对同一个区块进行了 pre-commit 投票，那么这个区块才会被提交。

由于离线或者网络延迟等原因，可能造成提议人提议区块失败。这种情况在 Tendermint 中也是被允许的，因为验证者会在进入下一轮提议之前等待一定时间，用于接收提议人提议的区块。

假设少于 1/3 的验证人是拜占庭节点，Tendermint 能够保证验证人永远不会在同一高度重复提交区块而造成冲突。为了做到这一点，Tendermint 引入了锁定机制，一旦验证者对一个区块进行了预投票，那么该验证者就会被锁定在这个区块。然后出现以下两种情况：

该验证人必须在预提交的区块进行预投票。

当前一轮预提议和预投票没成功提交区块时，该验证人就会被解锁，然后进行对新块的下一轮预提交。



![img](https:////upload-images.jianshu.io/upload_images/1953055-c94b8867e1def545.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



这么做的优点在于 Tendermint 共识机制拥有无法比拟的共识最终确定性：一旦共识达成就是真实的， 而不像比特币或以太坊的共识是一种概率性质的确定性，并且有可能在将来某个时刻失效。因此在 Tendermint 中不会出现区块链分叉的情况。其缺点是 Tendermint 共识是拜占庭容错的，最多容忍不超过1/3的恶意节点，而比特币最多可以容忍50%的恶意节点。



![img](https:////upload-images.jianshu.io/upload_images/1953055-beaf17e82e0756c7.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)


 上图展示了 Tendermint 的状态机复制过程，其安全性在于 BFT 算法保证了在同一高度只有一个确定的区块。

![img](https:////upload-images.jianshu.io/upload_images/1953055-6f063e86ee04244c.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-586a8ecbd135c172.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-9039fe8d7a98fdff.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-3c64f811629f1660.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



Tendermint Core 通过一个满足 ABCI 标准的 socket 协议与应用程序进行交互。

ABCI 包含3 种消息类型：DeliverTx 消息，CheckTx 消息和 Commit 消息。Tendermint Core 将消息发送给应用，应用根据消息进行相应的回复。

DeliverTx 消息是应用的主要工作流程。链上的每笔交易都通过这个消息进行传送。应用基于当前状态、应用协议和交易的加密证书，验证接收到 DeliverTx 消息的每笔交易。然后，验证后的交易会更新应用状态。

CheckTx 消息类似于 DeliverTx，但它仅用于验证交易。Tendermint Core 的内存池首先使用 CheckTx 检验交易的有效性，并且只将有效交易广播给其他节点。

Commit 消息用于通知应用程序计算当前应用状态的加密保证，该加密保证会被放入下一个区块头部。

一个应用可能与多个 ABCI socket 连接。Tendermint Core 给应用程序创建了三个 ABCI 连接：一个用于内存池广播时的交易验证，一个用于区块提议时的共识引擎，还有一个用于查询应用状态。



![img](https:////upload-images.jianshu.io/upload_images/1953055-eddd06b67b62a729.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



下面将进入我们的实战操作环节，屏幕前的你赶紧打开电脑跟着接下来 PPT 中的指示使用 Tendermint 开发一个属于自己的公链吧！



![img](https:////upload-images.jianshu.io/upload_images/1953055-7f02694820190e62.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-166914b0e866fefa.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-959ff66e9a8a23ae.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-dd60ce7979904751.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-1204e36f226a1ff9.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-d715f5f7ba11c0b6.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-3d724207f483c2b7.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



![img](https:////upload-images.jianshu.io/upload_images/1953055-c5a626a241671a10.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)



看到这里，不知道大家对 Tendermint 的了解有没有更上一层楼呢？



![img](https:////upload-images.jianshu.io/upload_images/1953055-37677fdefbe3535a.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

*本次实录纪要由 杨宁霞Ella ( HiBlock区块链社区上海运营负责人  ；下笔有神科技 社区运营) 整理记录，转发务必注明出处及本段信息。*

------

**现场活动合影照片：**

![img](https:////upload-images.jianshu.io/upload_images/1953055-cfd5b45496efe07b.png?imageMogr2/auto-orient/strip|imageView2/2/w/1095/format/webp)

站最中心位置长得高高的帅哥为本期特邀讲师



# 3. 下期活动宣传

**1）题目：**
 【HiBlock技术工坊56期】预期共识 - Filecoin对共识机制的探索

**2）议题：**
 关于区块链的共识现状，关于Filecoin的共识机制探索；目前的主流共识机制的缺陷在哪里？相关技术研究的重心放在哪里？
 预期共识--一个近乎完美的共识机制，结合技术剖析，聊聊它解决的矛盾，对生态的启示。

**议程纲要：**
 1)共识机制逻辑之领导人选举
 2)逻辑完美的共识机制
 3)预期共识的技术解析

**3）嘉宾：**

![img](https:////upload-images.jianshu.io/upload_images/1953055-c4967a05cdcdb697.png?imageMogr2/auto-orient/strip|imageView2/2/w/217/format/webp)

李昕（Steven Li）



**4）时间/地点：**
 2019-08-28(周三晚上) 18:30 / 上海徐汇区龙华中路596号A座

**5）活动报名海报**
 *敬请报名参会，空降无法着落。*



作者：ella_宁
链接：https://www.jianshu.com/p/c82a020f90fb
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。







# Tendermint共识算法原理和框架流程

![img](https://upload-images.jianshu.io/upload_images/1190574-bab84411e0e0748a.png?imageMogr2/auto-orient/strip|imageView2/2/w/730/format/webp)

# 1. 摘要

Tendermint是跨链Cosmos项目的核心技术。本文主要介绍以下内容：
（1）Tendermint的网络层级框架图
（2）Tendermint模块组成及共识算法原理
（3）Tendermint工作流程

# 2. Tendermint概述

Cosmos的开发团队Tendermint其实早在2014年就开始意识到了其不足，并持续专注于寻求不依赖挖矿等高电力消耗的共识机制，提供快速的交易处理能力，它们的目标是为全事件所有的区块链提供速度、安全和可扩展性。目前，Tendermint加入了微软Azure区块链即服务平台，也成为了以太坊区块链联盟成员之一，同时Tendermint也是跨链技术Cosmos的核心技术。两者大致的关系如下：

![img](https://upload-images.jianshu.io/upload_images/1190574-da7696e086d8ada9.png?imageMogr2/auto-orient/strip|imageView2/2/w/538/format/webp)

图中可以轻松看出Cosmos就是在Tendermint基础上添加一些插件功能来实现的。

## 2.1 Tendermint的概念

Tendermint的概念总结下有以下几点：
（1）Tendermint是一个能够在不同机器上，安全一致复制应用的软件，其中安全性和一致性也是分布式账本的关键概念。
（2）Tendermint具备拜占庭容错能力，是一种拜占庭容错共识算法。
（3）Tendermint主要有两部分组成：
1）Tendermint Core：区块链共识引擎，负责节点之间数据传输以及拜占庭共识。
2）ABCI：区块链应用程序接口(the Application BlockChain Interface )，也是一个协议，支持任何语言的交易处理实现。
总体来讲，Tendermint可以理解为一个模块化的区块链软件框架，支持开发者个性化定制自己的区块链，而又不需要考虑共识以及网络传输的实现。

## 2.2 Tendermint设计原则

区块链是一个具备确定性的状态机，可以在不信任的节点之间进行状态复制，包括应用的状态和改变状态的交易。从架构的层面上，区块链可以简单分为三个概念层：
（1）网络层（Networking）：负责交易和数据传输和同步。
（2）共识算法（Consensus）：负责不同的验证节点处理完交易后，保证状态的一致，也就是将交易打包到区块中。
（3）应用程序（Application）：交易的真正执行者。

大致框架如下：

![img](https://upload-images.jianshu.io/upload_images/1190574-01489aeaf3d12a1e.png?imageMogr2/auto-orient/strip|imageView2/2/w/512/format/webp)

目前大部分的区块链实现都是采用上面的框架，实现成单一的程序，但是这就很容易出现两个问题：
（1）代码复用困难，代码库的分支管理变得复杂。
（2）限制了应用开发的语言。

如何去规避这两个问题呢？Tendermint设计了自己的一套框架，其设计原则是易使用，易理解，高性能，适用于各种分布式应用。它的创新之处在于，将区块链应用（状态）与底层共识进行了分离，将共识引擎和P2P网络层封装组成Tendermint Core。同时提供ABCI接口与应用层进行交互，应用逻辑可以用任何语言编写，应用做的事情实际上就是状态机控制。基于这种架构，应用的开发者可以方便地实现自己的区块链。

![img](https://upload-images.jianshu.io/upload_images/1190574-05a36d48ecebc634.png?imageMogr2/auto-orient/strip|imageView2/2/w/640/format/webp)

Tendermint的框架总体来讲分为ABCI Application以及Tendermint Core两部分，两者通过ABCI连接。

# 3. Tendermint核心模块

## 3.1 ABCI Application

开发者定制开发的区块链应用，开发语言不受限制，可以使用任何语言进行开发，但是必须实现为一个ABCI Server，即需要满足以下几点：

（1）是一个Socket Server，需支持TSP或GRPC两种方式之一。
（2）能够处理ABCI Message。
所有的ABCI消息类型都是通过protobuf来定义的，具体的消息格式可参考[https://github.com/tendermint/abci/blob/master/types/types.proto](https://links.jianshu.com/go?to=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%3A%2F%2Fgithub.com%2Ftendermint%2Fabci%2Fblob%2Fmaster%2Ftypes%2Ftypes.proto)
（3） 实现区块链应用接口（ABCI）。
ABCI是Tendermint中定义的一套Application与Tendermint Core之间交互的协议。详细定义如下（版本：0.10.3）：

![img](https://upload-images.jianshu.io/upload_images/1190574-2b797fd92741718b.png?imageMogr2/auto-orient/strip|imageView2/2/w/640/format/webp)



**ABCI接口可以分为三类：**信息查询、交易校验以及共识相关处理。而Tendermint Core作为ABCI Client在启动时，会与ABCI Server建立三个连接，分别用于这三类接口消息的处理。

在Tendermint Core与Application交互的所有消息类型中，有3种主要的消息类型：
**（1）CheckTx消息**用于验证交易。Tendermint Core中的mempool通过此消息校验交易的合法性，通过之后才会将交易广播给其它节点。
**（2）DeliverTx消息**是应用的主要工作流程，通过此消息真正执行交易，包括验证交易、更新应用程序的状态。
**（3）Commit消息**通知应用程序计算当前的事件状态，并存在下一区块头中。

## 3.2 Tendermint Core

Tendermint共识引擎，包含区块链需要大部分功能实现，主要有：

- **共识算法**：拜占庭POS算法。
- **P2P**：采用gossip算法，默认端口是46656。
- **RPC**：区块链对外接口，默认端口是46657。支持三种访问方式：URI over HTTP、JSONRPC over HTTP、JSONRPC over websockets。详细的RPC接口定义列表可以参考[https://tendermint.github.io/slate](https://links.jianshu.com/go?to=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%3A%2F%2Ftendermint.github.io%2Fslate)
- **其它**：交易缓存池、消息队列等。

### 3.2.1 共识算法

Tendermint是一个易于理解的BFT共识协议。协议遵循一个简单的状态机，如下：

![img](https://upload-images.jianshu.io/upload_images/1190574-80f88e1830c06f07.png?imageMogr2/auto-orient/strip|imageView2/2/w/640/format/webp)

**协议中有两个角色：**
**（1）验证人：**协议中的角色或者节点，不同的验证者在投票过程中具备不同的权力（vote power）。
**（2）提议人：**由验证人轮流产生。
验证人轮流对交易的区块提议并对提议的区块投票。区块被提交到链上，且每个区块就是一个区块高度。但区块也有可能提交失败，这种情况下协议将选择下一个验证人在相同高度上提议一个新块，重新开始投票。

从图中可以看到，成功提交一个区块，必须经过两阶段的投票，称为pre-vote和pre-commit。当超过 2/3 的验证人在同一轮提议中对同一个块进行了pre-commit投票，那么这个区块才会被提交。

由于离线或者网络延迟等原因，可能造成提议人提议区块失败。这种情况在Tendermint中也是允许的，因为验证人会在进入下一轮提议之前等待一定时间，用于接收提议人提议的区块。

假设少于三分之一的验证人是拜占庭节点，Tendermint能够保证验证人永远不会在同一高度重复提交区块而造成冲突。为了做到这一点，Tendermint 引入了锁定机制，一旦验证人预投票了一个区块，那么该验证人就会被锁定在这个区块。然后：
（1）该验证人必须在预提交的区块进行预投票。
（2）当前一轮预提议和预投票没成功提交区块时，该验证人就会被解锁，然后进行对新块的下一轮预提交。

可以看到，Tendermint共识算法和PBFT时非常相似的，可以说是PBFT的变种，那我们来比较一下：

**（1）相同点：**
1）同属BFT体系。
2）抗1/3拜占庭节点攻击。
3）三阶段提交，第一阶段广播交易（区块），后两阶段广播签名（确认）。
4）两者都需要达到法定人数才能提交块。

**（2）不同点：**
1）Tendermint与PBFT的区别主要是在超过1/3节点为拜占庭节点的情况下。
当拜占庭节点数量在验证者数量的1/3和2/3之间时，PBFT算法无法提供保证，使得攻击者可以将任意结果返回给客户端。而Tendermint共识模型认为必须超过2/3数量的precommit确认才能提交块。举个例子，如果1/2的验证者是拜占庭节点，Tendermint中这些拜占庭节点能够阻止区块的提交，但他们自己也无法提交恶意块。而在PBFT中拜占庭节点却是可以提交块给客户端。
简单的说，就是比特币的网络存在分叉的可能，而Tendermint不会发生这种情况。
2）另一个不同点在于拜占庭节点概念不同，PBFT指的是节点数，而Tendermint代表的是节点的权益数，也就是投票权力。
3）最后一点，PBFT需要预设一组固定的验证人，而Tendermint是通过要求超过2/3法定人数的验证人员批准会员变更，从而支持验证人的动态变化。

**锁机制详解**
举个例子，有四个validator 节点，A,B,C,D, 在某个R轮，在propose阶段，
（1）proposer节点广播出了新块blockX；
（2）A的超时时间内没有收到这个新块，向外广播pre-vote nil，B,C,D都收到了，向外广播pre-vote投给blockX；
（3）现在四个节点进入了pre-commit阶段，A处于红色内圈，B,C,D处于蓝色外圈；
（4）假设A由于自身网络不好，又没有在规定时间内收到超过2/3个对blockX的投票，于是只能发出 pre-commit nil投票消息投给空块
（5）D收到了B和C的pre-vote消息，加上自己的，就超过了2/3了，于是D在本机区块链里commit了blockX
（6）假设此时B和C网络出现问题，收不到D在pre-commit消息，这是B和C只能看到2票投给了blockX，一票投给了空块，全部不足2/3，于是B和C都只能 commit空块，高度不变，进人R+1轮，A也只能看到2票投给了blockX，一票投给了空块，也只能commit空块，高度不变，进人R+1轮；
（7）在R+1轮，由于新换了一个proposer, 提议了新的区块blockY，A,B,C 三个个可能会在达成共识，提交blockY，于是在同样的高度，就有blockX和blockY两个块，产生了分叉？其实，Tendermint加上了锁的机制，具体就是，在第7步，即使proposer出了新块blockY，A,B,C只能被锁定在第6步他们的pre-commit块上，即A在第6步投给了空块，那么在第R+1轮，只能继续投给空块，B在第6步投给了blockX，那么在新一轮，永远只能投给blockX，C也是类似。这样在R+1轮，就会有1票投给空块，两票投给blockX，最终达成共识blockX，A,B,C三人都会commit blockX，与D一致，没有产生冲突。

## 3.2.2 P2P网络

Tendermint的P2P网络协议借鉴了比特币的对等发现协议，更准确地说，Tendermint是采用了BTCD的P2P地址簿（Address Book）机制。当连接建立后，新节点将自身的Address信息（包含IP、Port、ID等）发送给相邻节点，相邻节点接收到信息后加入到自己的地址薄，再将此条Address信息，转播给它的相邻节点。

此外为了保证节点之间数据传输的安全性，Tendermint采用了基于Station-to-Station协议的认证加密方案，此协议是一种密钥协商方案，基于经典的DH算法，并提供相互密钥和实体认证。大致的流程如下：

（1）每一个节点都必须生成一对ED25519密钥对作为自己的ID。
（2）当两个节点建立起TCP连接时，两者都会生成一个临时的ED25519密钥对，并把临时公钥发给对方。
（3）两个节点分别将自己的私钥和对方的临时公钥相乘，得到共享密钥。这个共享密钥对称加密密钥。
（4）将两个临时公钥以一定规则进行排序，并将两个临时公钥拼接起来后使用Ripemd160进行哈希处理，后面填充4个0，这样可以得到一个24字节的随机数。
（5）得到的随机数作为加密种子，但为了保证相同的随机数不会被相同的私钥使用两次，我们将随机数最后一个bit置为1，这样就得到了两个随机数，同时约定排序更高的公钥使用反转过的随机数来加密自己的消息，而另外一个用于解密对方节点的消息。
（6）使用排序的临时公钥拼接起来，并进行SHA256哈希，得到一个挑战码。
（7）每个节点都使用自己的私钥对挑战码进行签名，并将自己的公钥和签名发给其它节点校验。
（8）校验通过之后，双方的认证就验证成功了。后续的通信就使用共享密钥和随机数进行加密，保护数据的安全。

# 3.3 应用示例

Tendermint官方项目里内置了ABCI Application的两个简单实现counter以及kvstore。这个两个Demo逻辑非常简单，运行起来也非常简单，以kvstore为例，只需要下面三条简单的指令就可以轻松的跑起来：

> tendermint init

> abci-cli kvstore

> tendermint node

复杂一点，假设想使用Tendermint实现一套类似Ethereum的应用，最终应该是这样：

![img](https://upload-images.jianshu.io/upload_images/1190574-822bc0d084d4cc3e.png?imageMogr2/auto-orient/strip|imageView2/2/w/431/format/webp)

由Tendermint Core负责交易和区块的共享以及共识处理，开发者只需将go-ethereum和ABCI Server集成一个ABCI应用。Ethermint项目就是Tendermint团队开发的一个类似应用，大家可以参考，遗憾的是目前Ethermint目前只支持低版本的abci和go-ethereum。

# 4. Tendermint工作流程

![img](https://upload-images.jianshu.io/upload_images/1190574-ae9f4371801787f8.png?imageMogr2/auto-orient/strip|imageView2/2/w/640/format/webp)

**上图简单描述了Tenermint的工作流。大致为：**
（1）client通过RPC接口broadcast_tx_commit提交交易；
（2）mempool调用ABCI接口CheckTx用于校验交易的有效性，比如交易序号、发送者余额等，同时订阅交易执行后的事件并等待监听。
（3）共识从mempool中获取交易开始共识排序，打包区块，确定之后依次调用ABCI相关接口更新当前的事件状态，并触发事件。
（4）最终将交易信息返回client。

# 5. 参考

本文转载自[《深度解析Tendermint，快速融入Cosmos生态》](https://links.jianshu.com/go?to=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F38252058)。

更多Tendermint资料参考：
（1）拜占庭共识Tendermint介绍及简单入门
[https://blog.csdn.net/niyuelin1990/article/details/80537329](https://links.jianshu.com/go?to=https%3A%2F%2Fblog.csdn.net%2Fniyuelin1990%2Farticle%2Fdetails%2F80537329)
（2）Tendermint 说明文档
[https://tendermint.readthedocs.io/en/master/](https://links.jianshu.com/go?to=https%3A%2F%2Ftendermint.readthedocs.io%2Fen%2Fmaster%2F)
（3）Tendermint GIT地址
[https://github.com/tendermint/tendermint](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Ftendermint%2Ftendermint)
（4）深度解析Tendermint，快速融入Cosmos生态[质量高]
[https://zhuanlan.zhihu.com/p/38252058](https://links.jianshu.com/go?to=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F38252058)
（5）区块链框架 Tendermint 入门教程
[https://hbliu.coding.me/2018/04/02/tendermint-introduction-1/](https://links.jianshu.com/go?to=https%3A%2F%2Fhbliu.coding.me%2F2018%2F04%2F02%2Ftendermint-introduction-1%2F)
（6）详解Tendermint共识算法
[https://www.odaily.com/post/5134145](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.odaily.com%2Fpost%2F5134145)
（7）分布式一致性协议介绍（Paxos、Raft）
[https://www.cnblogs.com/zhang-qc/p/8688258.html](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.cnblogs.com%2Fzhang-qc%2Fp%2F8688258.html)

[辉哥的技术投资之路](https://www.jianshu.com/nb/2423685)