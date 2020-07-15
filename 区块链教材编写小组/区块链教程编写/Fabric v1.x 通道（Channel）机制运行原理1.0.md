# 1.Fabric v1.x 通道（Channel）机制运行原理

## 文章目录

一、Channel实现原理
1.1 System Channel
1.2 创建新channel
1.3 普通交易（Normal Transaction）
1.4 配置交易（Config Transaction）
1.5 总结
二、Channel数据隔离
三、Channel上部署chaincode

## 1、Channel实现原理

### 1.1 System Channel

channel是Orderer的一个模块，Fabric的启动会创建一个内建的system channel，是系统的一个默认链，用于管理其他的user channel。
orderer启动的时候必须要有该channel的genesis block，genesis block里规定了所有关于system channel的配置，因此所有的orderer都必须拿到相同的genesis block才能启动。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091311366.png)

### 1.2 创建新channel

创建一个新的user channel时，其实是向system channel发送一个“New Channel” transaction（包括要创建的channel的名字，channel的配置信息，包括哪些组织，出块属性等），这个transaction会被提交到sysem channel，然后orderer中会创建一个新的user channel，刚刚发送的transction里的信息作为user channel的genesis block，这样user channel就创建完成了。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091400967.png)



### 1.3 普通交易（Normal Transaction）

join channel的操作就是把genesis block拿下来，送给peer，让peer知道去哪儿找到orderer节点（genesis block里含有orderer的address信息），peer也会知道链里包含哪些组织（peer一定是属于这个channel允许的组织），这时候peer就可以参与transaction了，下图中的“Normal”就是用户的普通交易。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091505292.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)



### 1.4 配置交易（Config Transaction）

当需要修改channel的属性（新增一个组织、修改batchsize等配置），其实是向链发送一个“Config”transaction，该transaction无需遵循出块规则，会立即生成一个块，因为配置的信息需要尽快生效。修改一个链的属性的时候，本身也是需要共识的，共识产生块后，提交到orderer的本身的账本后，进行真正的属性更新。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091545677.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)



### 1.5 总结

使用上述同样的方法，system channel可以创建一个新的user channel B，system channel只负责创建user channel，创建完之后，user channel完全独立，与system channel也没有任何关系了。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091611633.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)



system channel与user channel在实现上没有区别，同样样可以提交一个“Config”transaction，共识产生块后更新配置属性。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091638679.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

Fabric的channel之间是相互隔离的，仅有的联系是user channel需要通过system channel进行创建，但是创建完之后，这些channel相互之间没有任何影响了。以此类推，新的user channel B同样可以通过提交“Config” transaction修改配置属性，或者参与业务交易；通过向system channel提交“New Channel”交易创建user channel C。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200222091714951.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)



## 2、Channel数据隔离

创建channel的交易在system channel共识完成后，channel A就会被创建，里面包含OrgA和OrgB两个组织，如下图所示：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200218155457580.png)

使用上述方法，可以创建出channel B和channel C，组织OrgA同时在channelA和channelB中，组织OrgB同时在channelA、channelB、channelC中，组织OrgC在channelB、channelC中，如下图所示：

![在这里插入图片描述](https://img-blog.csdnimg.cn/2020021815561063.png)

在这种情况下，channelB中包含全部3个组织，3个组织都可以在channelB中互相交易，而组织OrgC无法获取到channelA中的交易，组织OrgA无法获取到channelC中的交易，这样就实现了数据隔离和隐私。

## 3、Channel上部署chaincode

channel建立完成后，就可以在channel上部署chaincode，如下图所示，在channelA上部署了“mycc_v1.0”的chaincode，部署该chaincode时指定了背书策略为“AND(OrgA, OrgB)”，组织OrgA和OrgB都必须给该chaincode的交易背书才会被写入账本，这也意味着组织OrgA和OrgB都必须部署“mycc_v1.0”chaincode，如下图所示：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200218155733137.png)



利用规定的语法可以组成一些比较复杂的背书策略，如下图所示，在channelB部署“mycc_v1.0”chaincode的背书策略为“OR(OrgA, AND(OrgB, OrgC))”，表示组织OrgA单独背书即可，或者组织OrgB和OrgC一起背书，channelC的背书策略为“OutOf(1, OrgB, OrgC)”，表示要求除组织B和组织C外的另一个组织进行背书。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200218160109275.png)



另外，任何channel上都可以部署多个chaincode，并且这些chaincode可以相互调用，在同一个channel上的智能合约互相调用可以修改数据，不同channel上的智能合约也可以互相调用（取决于ACL），但是只能读取数据，不能修改数据，不同的chaincode组织起来可以完成一个复杂的业务逻辑。如下图所示，channelB和channelC上还部署了“evmcc_v0.1*”chaincode，用于运行以太坊的智能合约，channelA上还部署了“hercc_v2.0”chaincode。



![在这里插入图片描述](https://img-blog.csdnimg.cn/20200218160156965.png)

————————————————
版权声明：本文为CSDN博主「ice-fire-x」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/ice_fire_x/article/details/104433733



# 2.Fabric v1.x Peer节点与交易流程分析

## 文章目录

- [一、Fabric节点类型](https://uzshare.com/view/816666#Fabric_1)

- [二、Peer节点内部模块](https://uzshare.com/view/816666#Peer_9)

- [三、交易流程](https://uzshare.com/view/816666#_22)

- - [3.1 client提出交易](https://uzshare.com/view/816666#31_client_27)
  - [3.2 Endorsers仿真执行交易提案](https://uzshare.com/view/816666#32_Endorsers_30)
  - [3.3 client接收提案响应](https://uzshare.com/view/816666#33_client_33)
  - [3.4 client将响应提交给Ordering](https://uzshare.com/view/816666#34_clientOrdering_36)
  - [3.5 Orderer分发给Committing Peers](https://uzshare.com/view/816666#35_OrdererCommitting_Peers_39)
  - [3.6 Committing peers验证交易](https://uzshare.com/view/816666#36_Committing_peers_42)
  - [3.7 Committing peers通知client](https://uzshare.com/view/816666#37_Committing_peersclient_45)



## 1、Fabric节点类型

Fabric网络中包括Peer节点和排序节点（Ordering Node），Peer节点分为记账节点（Committing Peer）和背书节点（Endorsing Peer），这三种节点发挥的作用如下：

- **Committing Peer：**
  负责维护账本和世界状态以及将交易提交到账本并更新世界状态，可能部署有链码；
- **Endorsing Peer：**
  负责接收申请背书的交易提案，仿真执行交易，验证交易内容是否遵守智能合约，然后回复授权或拒绝背书，背书者对合约进行签名；背书节点必须部署链码；
- **Ordering Node：**
  负责将交易打包成区块，并且与committing peer和endorsing peer进行通信，控制记账内容以确保账本是一致的；无需部署智能合约，无需维护账本。

## 2、Peer节点内部模块

Peer nodes是区块链网络的基础，是账本和智能合约的载体。通过智能合约，账本以不可篡改的方式记录交易的全过程。Peer的内部模块如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200219150307884.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

- **Channels**
  对一个公司来说有不同的业务，需要和不同的公司进行相关联，这就需要创建多个链，因此就需要创建多个channel；channel是多个成员之间以机密交易为目的而建立的私网，可以很好的隔离数据；
- **Ledgers**
  每个channel可以维护一个或多个账本；一个Peer可以join多个channel，每个channel都要维护自己的账本，不同channel的账本之间互相隔离；
- **Chaincode**
  Peer需要维护安装好的智能合约；还要管理运行时Docker容器以实例化链码；
- **LocalMSP**
  Membership Services Provider：提供身份认证、加密、签名等服务
- **Events**
  给客户端Application发出Events通知

## 3、交易流程

下图展示了Fabric v1.x的架构，交易由client主动发起，经Endoser背书，然后经由Ordering排序打包，最后由Committer来记入账本，将结果的events通知给客户端。
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200219152139645.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)
下面的交易流程图展示了各个角色是如何交互的，以及交易流的全过程，后面将详细分析这七步交易过程。
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200219153537292.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.1 client提出交易

client应用提交一个智能合约A的交易提案，根据背书策略必须发送给背书节点{E0, E1, E2}，{P3、P4}则不在策略要求范围内，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/2020021814575937.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.2 Endorsers仿真执行交易提案

E0、E1、E2将各自仿真执行提出的交易，这些执行都不会更新账本，每次执行都会产生一组读写集，包含读取和需要写入的数据，这些数据将在fabric中流动；Endorsers还会对读写集进行签名，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/2020021814594172.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.3 client接收提案响应

读写集被异步返回给client，读写集由每个Endorser签名，并且还包含每个记录的版本号，此信息将会在以后的共识过程中进行检查，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200218150916214.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.4 client将响应提交给Ordering

client提交响应给Ordering，内容是一个要被排序的交易，Ordering本身不会查看交易内容，与其他client提交的交易并行地在fabric中进行排序，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200218151147633.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.5 Orderer分发给Committing Peers

排序服务按照一定的规则将交易打包到区块中，以分发给Committing peers（peer调用deliever接口pull区块）。然后Peers还可以分发给同一级的其他Peers，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200218151913889.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.6 Committing peers验证交易

每个Committing peer根据背书策略进行验证，还要检查读写集的版本号在当前世界状态下是否仍然合法；验证过的交易将应用于世界状态，并保存到账本中；非法的交易也被保存到账本中，但是不会更新到世界状态，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200218152042437.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3.7 Committing peers通知client

client可以注册事件触发器，当交易成功或失败、区块被添加到账本中时，client可以从被连接的的每个peer处得到确切的通知，如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200218152144978.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)
**根据以上交易流程，可以分析出Fabric中如何解决区块链中常见的双花问题：**
双花问题指的是恶意用户将同一笔钱花费两次。client在同一个读集version下，发两笔转账交易，在Endorser阶段会进行正常背书，但在最后验证阶段，第一个交易执行成功后，读集版本会发生变化，第二个交易的读集版本会对应不上，就会被认为是非法交易。

 

# 3.Fabric v1.x Idemix（Identity Mixer）介绍

## 文章目录

- [一、Idemix是什么](https://uzshare.com/view/818488#Idemix_1)
- [二、Idemix的实现](https://uzshare.com/view/818488#Idemix_10)
- [三、Idemix的特性](https://uzshare.com/view/818488#Idemix_25)

## 一、Idemix是什么

Idemix（Identity Mixer）的核心是零知识证明（Zero Knowledge Proof），用户无需暴露私有数据以及任何有用的信息，也能证明自己拥有这些私有数据，对方能够进行有效验证，这就是零知识证明。
Idemix是一个密码协议套件（X.509+加密算法），保留隐私实现匿名性，交易时不用透露交易者的身份，而且交易间是无关联的，不可往前追溯。

Identity包含三个角色，包括用户（User）、发行者（Issuer）、验证者（Verifier），各自作用如下：

- 用户：通过Idemix生成一个proof，证明自己知道某个秘密
- 发行者：（fabric CA 或 idemixgen工具）验证用户的隶属属性，然后颁发一个证书
- 验证者：（fabric MSP）验证proof
  ![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200303201433335.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 二、Idemix的实现

Fabric中Peer通过Fabric CA进行Enroll、Register、Revoke的操作，还可以通过Identity Mixer对交易进行签名验签的操作；
Identity Mixer和Fabric CA都需要调用加密包进行具体的流程。
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200303201842351.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)
**带有Identity Mixer的MSP**
除了X.509证书外，还可以通过发行idemix凭据来实现MSP。fabric CA初始化的时候，生成两个发行者的key，包括：

- IssuerPublicKey
- IssuerRevocationPublicKey。

Fabric CA或idemixgen工具可以作为发行者（Issuer），Idemix MSP作为验证者（Verifier），发行者为用户颁发数字证书，用户生成Idemix凭据后向验证者提供proof，用于验证属性是否正确。目前Idemix凭据只支持3个属性，包括：

- OU
- IsAdmin
- Enrollment ID
  ![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200303203430353.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 三、Idemix的特性

**Idemix与X.509的相同点如下：**

- 一组属性被签名，且签名不可伪造
- 凭证通过密码学的方式绑定到一个密钥

**Idemix与X.509的不同点如下：**

- Idemix通过零知识证明来确保不会泄露知识或信息，并且用户拥有凭证密钥；而X.509通过最初签名的公钥来验证，知道私钥的人才能生成证明；
- Idemix是各个信息间是无关联的，且不可往回追溯；而X.509显示所有属性，因此所有用于签发交易的X.509证书使用都是关联的。

Identity Mixer与X.509的对比如下图所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200303210041293.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)
Idemix与X.509是可以共存的，可以通过在configtx.yam里面指定“msptype: idemix”来支持Idemix，如下所示：
![在这里插入图片描述](https://uzshare.com/_p?https://img-blog.csdnimg.cn/20200303212249581.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)
**Idemix仍然有一些局限性：**

- 只支持固定的属性，例如OU、Role attribute、Enrollment ID、Revocation Handle attribute等；
- 不支持Idemix的撤销
- Peers还不能使用Idemix来进行背书，目前Peers的Idemix MSP只是用来验证签名，Idemix签名只能通过客户端SDK来进行；
- 建议每个channel或每个网络只使用一个基于Idemix的MSP，因为Idemix当前仅提供同一组织（MSP）中client的匿名性。



# 4.Identity Mixer MSP配置生成器（idemixgen）工具用法

本文章描述了idemixgen工具的用法，该工具可用于为基于MSP的身份混合器创建配置文件。 有两个命令可用，一个用于创建新的CA密钥对，另一个用于使用以前生成的CA密钥创建MSP配置。

## 一、目录结构

idemixgen工具将创建具有以下结构的目录：



```jsx
- /ca/
    IssuerSecretKey
    IssuerPublicKey
    RevocationKey
- /msp/
    IssuerPublicKey
    RevocationPublicKey
- /user/
    SignerConfig
```

ca目录包含颁发者密钥（包括撤销密钥），并且只应存在于CA. msp目录包含设置MSP验证idemix签名所需的信息。 用户目录指定默认签名者。

## 二、CA密钥生成

可以使用命令idemixgen ca-keygen创建适用于身份混合器的CA（颁发者）密钥。 这将在工作目录中创建目录ca和msp。

## 三、添加默认签名者

使用idemixgen ca-keygen生成ca和msp目录后，可以使用idemixgen signerconfig将user目录中指定的默认签名者添加到配置中。



```kotlin
$ idemixgen signerconfig -h
usage: idemixgen signerconfig [<flags>]

Generate a default signer for this Idemix MSP

Flags:
    -h, --help               Show context-sensitive help (also try --help-long and --help-man).
    -u, --org-unit=ORG-UNIT  The Organizational Unit of the default signer
    -a, --admin              Make the default signer admin
    -e, --enrollment-id=ENROLLMENT-ID
                             The enrollment id of the default signer
    -r, --revocation-handle=REVOCATION-HANDLE
                             The handle used to revoke this signer
```

例如，我们可以使用以下命令创建一个默认签名者，该签名者是组织单元“OrgUnit1”的成员，注册标识为“johndoe”，撤销句柄为“1234”，这是一个管理员：



```bash
idemixgen signerconfig -u OrgUnit1 --admin -e "johndoe" -r 1234
```



作者：时间里的小恶魔
链接：https://www.jianshu.com/p/da9a73708d17
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。





## 四、Fabric MSP 文件结构-案例分析

前置知识 [fabric msp 基础](https://zhuanlan.zhihu.com/p/35683522)

![img](https:////upload-images.jianshu.io/upload_images/15915219-dedf6dc9fed3d0f4.png?imageMogr2/auto-orient/strip|imageView2/2/w/637/format/webp)

节点构成.png



目录结构中的所有文件，就是我们用cryptogen工具来生成的MSP需要用到的证书和相关文件。主要包括各种证书和相关的签名。



```ruby
org1.example.com/
├── ca     # 存放组织Org1的根证书和对应的私钥文件，默认采用EC算法，证书为自签名。组织内的实体将基于该根证书作为证书根。
│   ├── ca.org1.example.com-cert.pem
│   └── dfb841b77804d726eea25231ae5e89a31901ca0538688a6d764731148f0bdc5b_sk
├── msp    # 存放代表该组织的身份信息。
│   ├── admincerts         # 组织管理员的身份验证证书，被根证书签名。
│   │   └── Admin@org1.example.com-cert.pem
│   ├── cacerts # 组织的根证书，同ca目录下文件。
│   │   └── ca.org1.example.com-cert.pem
│   └── tlscacerts          # 用于TLS的CA证书，自签名。
│       └── tlsca.org1.example.com-cert.pem
├── peers   # 存放属于该组织的所有Peer节点
│   ├── peer0.org1.example.com    # 第一个peer的信息，包括其msp证书和tls证书两类。
│   │   ├── msp # msp相关证书   
│   │   │   ├── admincerts  # 组织管理员的身份验证证书。Peer将基于这些证书来认证交易签署者是否为管理员身份。
│   │   │   │   └── Admin@org1.example.com-cert.pem
│   │   │   ├── cacerts     # 存放组织的根证书
│   │   │   │   └── ca.org1.example.com-cert.pem
│   │   │   ├── keystore    # 本节点的身份私钥，用来签名
│   │   │   │   └── 59be216646c0fb18c015c58d27bf40c3845907849b1f0671562041b8fd6e0da2_sk
│   │   │   ├── signcerts   # 验证本节点签名的证书，被组织根证书签名 
│   │   │   │   └── peer0.org1.example.com-cert.pem
│   │   │   └── tlscacerts  # TLS连接用到身份证书，即组织TLS证书
│   │   │       └── tlsca.org1.example.com-cert.pem
│   │   └── tls # tls相关证书
│   │       ├── ca.crt      # 组织的根证书
│   │       ├── server.crt  # 验证本节点签名的证书，被组织根证书签名
│   │       └── server.key  # 本节点的身份私钥，用来签名
│   └── peer1.org1.example.com    # 第二个peer的信息，结构类似。（此处省略。）
│       ├── msp
│       │   ├── admincerts
│       │   │   └── Admin@org1.example.com-cert.pem
│       │   ├── cacerts
│       │   │   └── ca.org1.example.com-cert.pem
│       │   ├── keystore
│       │   │   └── 82aa3f8f9178b0a83a14fdb1a4e1f944e63b72de8df1baeea36dddf1fe110800_sk
│       │   ├── signcerts
│       │   │   └── peer1.org1.example.com-cert.pem
│       │   └── tlscacerts
│       │       └── tlsca.org1.example.com-cert.pem
│       └── tls
│           ├── ca.crt
│           ├── server.crt
│           └── server.key
├── tlsca    # 存放tls相关的证书和私钥。
│   ├── 00e4666e5f56804274aadb07e2192db2f005a05f2f8fcfd8a1433bdb8ee6e3cf_sk
│   └── tlsca.org1.example.com-cert.pem
└── users    # 存放属于该组织的用户的实体
    ├── Admin@org1.example.com    # 管理员用户的信息，其中包括msp证书和tls证书两类。
    │   ├── msp # msp相关证书
    │   │   ├── admincerts     # 组织根证书作为管理员身份验证证书 
    │   │   │   └── Admin@org1.example.com-cert.pem
    │   │   ├── cacerts        # 存放组织的根证书
    │   │   │   └── ca.org1.example.com-cert.pem
    │   │   ├── keystore       # 本用户的身份私钥，用来签名
    │   │   │   └── fa719a7d19e7b04baebbe4fa3c659a91961a084f5e7b1020670be6adc6713aa7_sk
    │   │   ├── signcerts      # 管理员用户的身份验证证书，被组织根证书签名。要被某个Peer认可，则必须放到该Peer的msp/admincerts目录下
    │   │   │   └── Admin@org1.example.com-cert.pem
    │   │   └── tlscacerts     # TLS连接用的身份证书，即组织TLS证书
    │   │       └── tlsca.org1.example.com-cert.pem
    │   └── tls # 存放tls相关的证书和私钥。
    │       ├── ca.crt       # 组织的根证书
    │       ├── server.crt   # 管理员的用户身份验证证书，被组织根证书签名
    │       └── server.key   # 管理员用户的身份私钥，被组织根证书签名。
    └── User1@org1.example.com    # 第一个用户的信息，包括msp证书和tls证书两类
        ├── msp # msp证书相关信息
        │   ├── admincerts   # 组织根证书作为管理者身份验证证书。
        │   │   └── User1@org1.example.com-cert.pem
        │   ├── cacerts      # 存放组织的根证书
        │   │   └── ca.org1.example.com-cert.pem
        │   ├── keystore     # 本用户的身份私钥，用来签名
        │   │   └── 97f2b74ee080b9bf417a4060bfb737ce08bf33d0287cb3eef9b5be9707e3c3ed_sk
        │   ├── signcerts    # 验证本用户签名的身份证书，被组织根证书签名
        │   │   └── User1@org1.example.com-cert.pem
        │   └── tlscacerts   # TLS连接用的身份证书，被组织根证书签名。
        │       └── tlsca.org1.example.com-cert.pem
        └── tls # 组织的根证书
            ├── ca.crt       # 组织的根证书
            ├── server.crt   # 验证用户签名的身份证书，被根组织证书签名
            └── server.key   # 用户的身份私钥用来签名。
```

总结下来，MSP的peer&Orderer节点配置就是，在每一个Peer节点和排序服务节点上设置MSP目录，放入相关的证书和签名公私钥，然后在对应的配置文件中，设置msp文件路径。Peer节点和排序服务节点就有了签名证书，在通道节点之间传输数据时，要验证节点的签名，从而实现证书的验证和签名。



作者：DUO1510
链接：https://www.jianshu.com/p/61ca3af17a67
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。







# 5.Hyperledger Fabric中的零知识证明

https://www.51discuss.com/categories/区块链)

 Fabric 1.3中的新增的idemixer(Identity Mixer)以前不大懂zero-knowledge proof(零知识证明),原本觉得PKI基础的MSP是比较常用和稳健的方式,新加个验证方式是不是有点增加复杂性。

最近有时间整理下,才发现零知识证明也真是个黑科技。

## 1. 零知识证明入门

### 1.1 零知识证明例子

网上这篇文章写得蛮好的http://www.elecfans.com/blockchain/1015964.html

这里以Fabric给出的例子: 假设Alice需要向Bob(门店职员)证明她DMV(车管所)颁发的合法驾照。这个场景,Alice就是下图的user/用户, DMV车管所则是issuer/证书颁发者,Bob则是verifier验证者。![idemix-overview.png](https://www.51discuss.com/images/20191012/idemix-overview.png)

Alice为了证明自己是合法的司机,大多时候她会把自己的驾照交给Bob检查和验证，但这样做Bob就可以知道Alice的很多额外的隐私信息,例如名字,地址,年龄等。

如果使用idemixer和零正式证明的方式, 我们只允许Bob知道当前这个女用户是个合法的司机,其它信息都保密。 即使下次Alice再来门店,Alice应该提供给Bob不同的证明,保证Bob不会知道这个证明是同一个用户。

即零知识证明可提供匿名性和无关联性。

### 1.2 零知识证明用处

elecfans的文章总结得很好了,常见的是以下两点。 - 数据隐私保护和身份验证,如Alice和Bob的例子所示, - 减少计算和扩容,同样的多次计算可以使用零知识证明压缩和减少,最新的以太坊可是大力推崇

## 2. 如果使用Fabric的idemixer

### 2.1 测试开发环境使用idemixgen命令行

具体参看https://hyperledger-fabric.readthedocs.io/en/latest/idemixgen.html

### 2.2 生产环境使用Fabric CA 1.3以上版本

(1) Fabric CA配置 fabric-ca-server init会生成IssuerPublicKey和IssuerRevocationPublicKey两个文件

(2) MSP初始化配置 configtx.yaml定义Org1Idemix和Org2Idemix两个组织, 注意msptype类型为idemix, mspdir对应文件应使用Fabric CA生成的IssuerPublicKey和IssuerRevocationPublicKey构成。

```yaml
Organizations:
   - &Org1Idemix
        # defaultorg defines the organization which is used in the sampleconfig
        # of the fabric.git development environment
        name: idemixMSP1

        # id to load the msp definition as
        id: idemixMSPID1

        msptype: idemix
        mspdir: crypto-config/peerOrganizations/org3.example.com


    - &Org2Idemix
        # defaultorg defines the organization which is used in the sampleconfig
        # of the fabric.git development environment
        name: idemixMSP2

        # id to load the msp definition as
        id: idemixMSPID2

        msptype: idemix
        mspdir: crypto-config/peerOrganizations/org4.example.com

profiles:
    TwoOrgsOrdererGenesis_v13:
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            Organizations:
                - *OrdererOrg
            Capabilities:
                <<: *OrdererCapabilities
        Consortiums:
            SampleConsortium:
                Organizations:
                    - *Org1
                    - *Org2
                    - *Org1Idemix
                    - *Org2Idemix
        Application:
            <<: *ApplicationDefaults
            Organizations:
                - *OrdererOrg
            Capabilities:
                <<: *ApplicationCapabilities1_3

    TwoOrgsChannel_v13:
        Consortium: SampleConsortium
        Application:
            <<: *ApplicationDefaults
            Organizations:
                - *Org1
                - *Org2
                - *Org1Idemix
                - *Org2Idemix
            Capabilities:
                <<: *ApplicationCapabilities1_3
```

(3) CA客户端生成用户

```java
IdemixEnrollment idemixEnrollment = hfcaClient.idemixEnroll(x509enrollment, "idemixMSPID1");
```

(4) 验证者链码如何获取idemixer信息 暂时go的链码的cid(Client Identity)库才支持获取idemixer证书信息。

```go
func GetAttributeValue(stub ChaincodeStubInterface, attrName string) (value string, found bool, err error)

type ChaincodeStubInterface interface {
    // GetCreator returns `SignatureHeader.Creator` (e.g. an identity)
    // of the `SignedProposal`. This is the identity of the agent (or user)
    // submitting the transaction.
    GetCreator() ([]byte, error)
}

type ClientIdentity interface {

    // GetID returns the ID associated with the invoking identity.  This ID
    // is guaranteed to be unique within the MSP.
    GetID() (string, error)

    // Return the MSP ID of the client
    GetMSPID() (string, error)

    // GetAttributeValue returns the value of the client's attribute named `attrName`.
    // If the client possesses the attribute, `found` is true and `value` equals the
    // value of the attribute.
    // If the client does not possess the attribute, `found` is false and `value`
    // equals "".
    GetAttributeValue(attrName string) (value string, found bool, err error)

    // AssertAttributeValue verifies that the client has the attribute named `attrName`
    // with a value of `attrValue`; otherwise, an error is returned.
    AssertAttributeValue(attrName, attrValue string) error

    // GetX509Certificate returns the X509 certificate associated with the client,
    // or nil if it was not identified by an X509 certificate.
    GetX509Certificate() (*x509.Certificate, error)
}
```

暂时idemixer的GetAttributeValue只支持ou和role - ou 是identity’s affiliation (e.g. “org1.department1”) - role 是‘member’ 或 ‘admin’.

具体调用的go链码

```go
package main

import (
	"fmt"
	"log"
    "os"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)
// Invoke makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	Info.Println("########### example_cc Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	// Getting attributes from an idemix credential
	ou, found, err := cid.GetAttributeValue(stub, "ou");
	if err != nil {
		return shim.Error("Failed to get attribute 'ou'")
	}
	if !found {
		return shim.Error("attribute 'ou' not found")
	}
	if !strings.HasSuffix(ou, "department1") {
		return shim.Error(fmt.Sprintf("Incorrect 'ou' returned, got '%s' expecting to end with 'department1'", ou))
	}
	role, found, err := cid.GetAttributeValue(stub, "role");
	if err != nil {
		return shim.Error("Failed to get attribute 'role'")
	}
	if !found {
		return shim.Error("attribute 'role' not found")
	}
	if role != "member" {
		return shim.Error(fmt.Sprintf("Incorrect 'role' returned, got '%s' expecting 'member'", role))
	}

	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}

	if function == "move" {
		// Deletes an entity from its state
		return t.move(stub, args)
	}

	Error.Printf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}
```

### 2.3 idemixer的限制

还不大完善,基本现阶段还是推荐用传统的MSP方式,具体参考https://hyperledger-fabric.readthedocs.io/en/latest/idemix.html#current-limitations

> 零知识证明在以太坊是推崇的,它的应用场景实际蛮广的,fabric尚需努力,不过貌似2.0那么久还没release或者是推广得不好。不过区块链的前途是光明的,道路是曲折的,多了解下这些基础不是坏事。

- **原文作者：**[Zealot](https://www.51discuss.com/)
- **原文链接：**https://www.51discuss.com/posts/fabric-zero-knowledge-proof/
- **版权声明：**本作品采用[知识共享署名-非商业性使用-禁止演绎 4.0 国际许可协议](https://creativecommons.org/licenses/by-nc-nd/4.0/)进行许可，非商业转载请注明出处（作者，原文链接），商业转载请联系作者获得授权。



## See Also

- [Fabric链码最佳实践](https://www.51discuss.com/posts/fabric-chaincode-best-pratise/)
- [Hyperledger Composer技术分享](https://www.51discuss.com/posts/hyperledger-composer-training/)
- [Fabric 1.4.1基于raft的排序服务](https://www.51discuss.com/posts/fabric-raft-cluster/)
- [Hyperledger Fabric 2.0 Alpha新特性](https://www.51discuss.com/posts/fabric-2.0-alpha/)
- [Hyperledger Fabric 1.4运维服务](https://www.51discuss.com/posts/fabric-operation-service/)

- [fabric](https://www.51discuss.com/tags/fabric)



# 6.Fabric CA 1.4入门

## 1.简介

Fabric CA基于开源项目CFSSL开发, 主要为fabric网络提供PKI证书服务,是MSP生成的基础。
可能有人会问, 官方不是有cryptogen工具批量生成MSP吗? cryptogen实际是辅助测试工具,默认不同orderer,org都有不同的CA, 如果一个org要追加个peer或user, cryptogen就不管用了。生产环境我们建议使用fabric ca全面管理证书, 如果想简单来而区块链组织,节点和用户基本不会变, cryptogen也没问题。

## 2.架构

![t_bfdd38ffe95f40d5b5f6ae4b109d8788.png](https://www.javatree.cn/file-server/e/20190219/t_bfdd38ffe95f40d5b5f6ae4b109d8788.png)

我们在MSP已有类似提过。
https://www.javatree.cn/news/db31d44e278f40179116ad9f3a5618cc
Fabric CA默认部署为restful服务, 单点默认使用sqlite, 也可使用nginx, ha-proxy, keepalive实现高可用和均衡负载, 支持mysql等数据库持久化。

## 3.Fabric-CA启动

可以参考fabric-sample的basic-network例子启动ca。 有点要注意, 默认镜像是设置了账号为admin/adminpw.

```shell
docker inspect hyperledger/fabric-ca:1.4.0
```

会发现默认会执行

```shell
Cmd fabric-ca-server start -b admin:adminpw
```

编写Docker-compose文件时command最好覆盖

```shell
command: sh -c 'fabric-ca-server start -b yourAccount:yourPassword'
```

貌似不大好, CA账户密码都暴露出来了。是否可以像传统数据库那样启动, 账号就保存在数据库里面。

默认sqlite是可以处理下, docker起个容器, 用清空配置, 使用fabric-ca-server init -b重新指定账号密码, 重新编辑生成的fabric-ca-server-config.yaml, 全部配置好使用新的这套配置启动fabric-ca即可。

使用mysql就还是得在fabric-ca-server-config.yaml配置数据库链接的账号密码, 绕不过。但是一个原则就是fabric-ca启动了， 务必不能让使用默认账号admin/adminpw, 就好像ssh暴露root账号一样危险。

## 4.Fabric-ca配置

实际上ca启动优先是命令行传入的参数, 其次是docker-compose环境变量, 再是fabric-ca-server-config.yaml配置, 参数较多, 我们就以yaml讲解具体配置， 里面也有不少注释，讲几个注意的点。

CA开启HTTPS需要配置enabled=true, 或者通过环境变量FABRIC_CA_TLS_ENABLED=true,
签名证书certfile,私钥keyfile, 可以使用互联网正式CA颁发的或自签名, 看使用场景。

```yaml
tls:
  # Enable TLS (default: false)
  enabled: false
  # TLS for the server's listening port
  certfile:
  keyfile:
  clientauth:
    type: noclientcert
    certfiles:
```

数据库默认是sqlite3, 也可以配置为mysql, datasource例如为

```shell
root:rootpw@tcp(localhost:3306)/fabric_ca?parseTime=true&tls=custom
```

如果用mysql5.7要保证日期格式允许’0000-00-00’, 即my.cnf删除NO_ZERO_DATE。

```yaml
db:
  type: sqlite3
  datasource: fabric-ca-server.db
  tls:
      enabled: false
      certfiles:
      client:
        certfile:
        keyfile:
```

Affiliations也是组织机构部门, CA register用户使用, 客户端可创建新的affiliation, affiliation笔者也有些疑惑,后面注册时说明。

```yaml
affiliations:
   org1:
      - department1
      - department2
   org2:
      - department1
```

签名有时效性,默认是1年半年, 按照实际需求适当的调整时间。

```yaml
signing:
    default:
      usage:
        - digital signature
      expiry: 87600h
    profiles:
      ca:
         usage:
           - cert sign
           - crl sign
         expiry: 438000h
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
         expiry: 87600h
```

CSR涉及到CA根证书的创建，可适当修改里面的名称, CA根证书默认15年有效。

```yaml
csr:
   cn: fabric-ca-server
   keyrequest:
     algo: ecdsa
     size: 256
   names:
      - C: US
        ST: "North Carolina"
        L:
        O: Hyperledger
        OU: Fabric
   hosts:
     - 28e9b53aaa88
     - localhost
   ca:
      expiry: 131400h
      pathlength: 1
```

还有不少配置，参考官方文档
https://hyperledger-fabric-ca.readthedocs.io/en/release-1.1/users-guide.html
相对方便些应该还是把这些YAML配置项设置为对应的环境变量方便些。

## 5.注册用户

一些固定的用户我们可以使用fabric-ca-client生成。
默认的admin账号获取到证书

```shell
CA_ADMIN_ID="admin:adminpw"
fabric-ca-client enroll -H ca-admin-home -d   -u http://$CA_ADMIN_ID@localhost:7054
```

ca-admin-home结构如下, 主要是msp

```yaml
[root@k8s-master ca-admin-home]# tree 
.
├── fabric-ca-client-config.yaml
└── msp
    ├── cacerts
    │   └── localhost-7054.pem
    ├── IssuerPublicKey
    ├── IssuerRevocationPublicKey
    ├── keystore
    │   └── 570aaaaadb9b185e68d6e6ae9c102bc2f04326b96f340568400e45a4fd9af71e_sk
    ├── signcerts
    │   └── cert.pem
└── user
```

假设创建orderer

```shell
CA_ADMIN_ID="admin@:adminpw"
ORDERER_NAME="orderer.example.com"
ORDERER_SECRET="123456"
fabric-ca-client register -H ca-admin-home  --id.name $ORDERER_NAME --id.secret $ORDERER_SECRET  --id.type orderer  -u  http://$CA_ADMIN_ID@localhost:7054
fabric-ca-client enroll  -H orderer-home --csr.names C=cn,ST=hubei,L=wuhan,O=example.com  -u http://$ORDERER_NAME:$ORDERER_SECRET@localhost:7054
```

生成目录, 里面是没tlsca的内容的, 可能需要额外再起创建账号生成一套。

```shell
[root@k8s-master orderer-home]# tree
.
├── fabric-ca-client-config.yaml
└── msp
    ├── cacerts
    │   └── localhost-7054.pem
    ├── IssuerPublicKey
    ├── IssuerRevocationPublicKey
    ├── keystore
    │   └── bb9de343d9bb7f92a02fa74b3be9e87fd8e3d6df78aa4cec17f0dcc6d2e2d259_sk
    ├── signcerts
    │   └── cert.pem
    └── user
```

时不时也要用openssl查看下证书内容, 尽量和cryptogen类似
openssl x509 -in cert.pem -noout -text

官方给出的很多例子, 都是通过O, OU等实际去确定角色, 例如first-network例子里org1的MSP里的config.yaml

```yaml
NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/ca.org2.example.com-cert.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/ca.org2.example.com-cert.pem
    OrganizationalUnitIdentifier: peer
```

而Java版fabric-ca-client SDK注册的时候可以传入affiliation, 命令行实际也可以设置, 但只是在生成证书的最后一段属性有生成, 可能hf.Type等就是默认直接对应到OU的角色, client, admin等, 但是cryptogen生成的内容是不包含这些attrs的，有点诡异，这些估计只能去跟源码了。

```shell
[root@k8s-master signcerts]# openssl x509 -in cert.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            1a:65:3f:d1:27:60:fd:80:b8:50:c7:6a:36:0c:c9:6a:1e:e9:d0:83
    Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=US, ST=California, L=San Francisco, O=org1.example.com, CN=ca.org1.example.com
        Validity
            Not Before: Feb 19 11:36:00 2019 GMT
            Not After : Feb 19 11:41:00 2020 GMT
        Subject: C=cn, ST=hubei, L=wuhan, O=example.com, OU=orderer, CN=orderer.example.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub: 
                    04:0f:5f:d3:82:d1:41:7f:5a:30:1e:5c:33:85:23:
                    af:4c:56:20:ad:38:86:e6:cd:59:53:9f:61:e4:ac:
                    c8:1f:79:11:d2:c1:75:e6:7c:2d:94:d4:eb:73:e1:
                    7f:25:36:73:61:c2:03:7c:8e:01:42:bd:65:dd:25:
                    2f:ad:df:a1:e1
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier: 
                92:EF:D0:BB:EE:81:C0:4B:0F:F1:38:40:63:0D:A8:C7:53:D8:25:4C
            X509v3 Authority Key Identifier: 
                keyid:42:39:AA:0D:CD:76:DA:EE:B8:BA:0C:DA:70:18:51:D1:45:04:D3:1A:AD:1B:2D:DD:DB:AC:6A:57:36:5E:49:7C

            X509v3 Subject Alternative Name: 
                DNS:k8s-master
            1.2.3.4.5.6.7.8.1: 
                {"attrs":{"hf.Affiliation":"","hf.EnrollmentID":"orderer.example.com","hf.Type":"orderer"}}
    Signature Algorithm: ecdsa-with-SHA256
         30:44:02:20:5d:a5:09:bb:16:7e:89:0b:d7:00:f8:fc:17:fe:
         e3:7b:06:68:2b:0c:cf:d5:e6:5a:95:82:5b:cf:e7:06:7b:05:
         02:20:44:65:90:da:34:1e:e3:b1:b7:09:a7:64:27:d8:6b:6d:
         6c:e1:56:14:db:5d:68:aa:80:b4:3a:a0:2e:a6:f1:70
```

而java fabric-ca-client SDK用法我们在开源项目有相应代码, 也可以参考下用SDK注册用户
https://github.com/zealzeng/fabric-rest

## 6.小结

先来个入门，官方的文档实际还有不少内容不知道如何用在fabric的部署当中用到，1.3的id mixer还没用上，把东西做复杂不是自豪的事情，但有那么多大公司撑着fabric, 将就用着吧。





# 7.浅析Hyperledger Fabric架构原理

## 1.Hyperledger Fabric概述

Hyperledger Fabric是由IBM公司主导开发的一个面向企业级客户的开源项目。与比特币和以太坊这类公有链不同，Hyperledger Fabric网络中的节点必须经过授权认证后才能加入，从而避免了POW资源开销，大幅提高了交易处理效率，满足企业级应用对处理性能的诉求。同时，为了满足灵活多变的应用场景，Hyperledger Fabric采用了高度模块化的系统设计理念，将权限认证模块（MSP）、共识服务模块（Ordering Service）、背书模块（Endorsing peers）、区块提交模块（committing peers）等进行分离部署，使开发者可以根据具体的业务场景替换模块，实现了模块的插件式管理（plug-in/plug-out）。所以，Hyperledger Fabric是一个私有链／联盟链的开发框架，而且系统的运行不需要token支持。

## 2.关键组件

Channel：是一种数据隔离机制，保证交易信息只有交易参与方可见，每个channel是一个独立的区块链，这使得多个用户可以共用同一个区块链系统而不用担心信息泄露问题。

Chaincode：也叫智能合约，将资产定义和资产处理逻辑封装成接口，当其被用户调用的时候，改变账本的状态。

Ledger：区块链账本，保存交易信息和智能合约代码。

Network：交易处理节点之间的P2P网络，用于维持区块链账本的一致性。

Ordering service：利用kafka、SBTF等共识算法对所有交易信息进行排序并打包成区块，发给committing peers节点，写入区块链中。

World state：显示当前资产数据的状态，底层通过LevelDB和CouchDB数据库将区块链中的资产信息组织起来，提供高效的数据访问接口。

Membership service provider（MSP）：管理认证信息，为client和peers提供授权服务。

![img](http://youqudian.com/uploads/allimg/200412/04302Va0-0.png)



## 3.Hyperledger Fabric Network中的角色

在Hyperledger中，由三种类型的角色：

Client：应用客户端，用于将终端用户的交易请求发送到区块链网络；

Peers：负责维护区块链账本，分为endoring peers和committing peers，其中，endorser为交易做背书（验证交易并对交易签名），committer接收打包好的区块，然后写入区块链中。Peers节点是一个逻辑的概念，endorser和committer可以同时部署在一台物理机上。

Ordering Service：接收交易信息，并将其排序后打包成区块，放入区块链，最后将结果返回给committer peers。



## 4.Hyperledger交易流程

1、客户端通过SDK接口，向endorsing peer节点发送交易信息：

![img](http://youqudian.com/uploads/allimg/200412/04302W9E-1.png)

2、每个endorsing peer节点模拟处理交易，此时并不会将交易信息写入账本。然后，endorser peer会验证交易信息的合法性，并对交易信息签名后，返回给client。此时的交易信息只是在client和单个endorser peer之间达成共识，并没有完成全网共识，各个client的交易顺序没有确定，可能存在双花问题，所以还不能算是一个有效的交易。同时，client需要收到大多数endorser peer的验证回复后，才算验证成功，具体的背书策略由智能合约代码控制，可以由开发者自由配置。

![img](http://youqudian.com/uploads/allimg/200412/04302R2G-2.png)

3、client将签名后的交易信息发送给order service集群进行交易排序和打包。Order service集群通过共识算法，对所有交易信息进行排序，然后打包成区块。Order service的共识算法是以组件化形态插入Hyperledger系统的，也就是说开发者可以自由选择合适的共识算法。

![img](http://youqudian.com/uploads/allimg/200412/04302QP6-3.png)

4、ordering service将排序打包后的区块广播发送给committing peers，由其做最后的交易验证，并写入区块链。ordering service只是决定交易处理的顺序，并不对交易的合法性进行校验，也不负责维护账本信息。只有committing peers才有账本写入权限。

![img](http://youqudian.com/uploads/allimg/200412/04302R1Y-4.png)

## 5.Hyperledger Fabric Network的共识算法

在所有peers中，交易信息必须按照一致的顺序写入账本（区块链的基本原则）。例如，比特币通过POW机制，由最先完成数学难题的节点决定本次区块中的信息顺序，并广播给全网所有节点，以此来达成账本的共识。而Hyperledger Fabric采用了更加灵活、高效的共识算法，以适应企业场景下，对高TPS的要求。目前，Hyperledger Fabric有三种交易排序算法可以选择。

SOLO：只有一个order服务节点负责接收交易信息并排序，这是最简单的一种排序算法，一般用在实验室测试环境中。Sole属于中心化的处理方式。

Kafka：是Apache的一个开源项目，主要提供分布式的消息处理／分发服务，每个kafka集群由多个服务节点组成。Hyperledger Fabric利用kafka对交易信息进行排序处理，提供高吞吐、低延时的处理能力，并且在集群内部支持节点故障容错。

SBFT：简单拜占庭算法，相比于kafka，提供更加可靠的排序算法，包括容忍节点故障以及一定数量的恶意节点。目前，Hyperledger Fabric社区正在开发该算法。

## 6.交易流程总结

区块链的账本由peer节点维护，并不是由ordering service集群维护，所以，只有peer节点上可以找到完整的区块链信息，而order service集群只负责对交易进行排序，只保留处理过程中的一部分区块链信息。Hyperledger Fabric系统中的节点是一个逻辑的概念，并不一定是一个台物理设备，但是对于生产环境的设计者来说，peer节点不能和order节点部署在一台机器上，而enduring peers和committing peers可以部署在同一台机器上，这种设计主要是为了系统架构的解耦，提高扩展性，以及通过主机隔离提高安全性。 Endorsing peer校验客户端的签名，然后执行智能合约代码模拟交易。交易处理完成后，对交易信息签名，返回给客户端。客户端收到签名后的交易信息后，发给order节点排序。Order节点将交易信息排序打包成区块后，广播发给committing peers，写入区块链中。一个完整的交易处理流程如下图所示：

![img](http://youqudian.com/uploads/allimg/200412/04302R095-5.png)

## 7.Channel的概念

Channels能够让上层不同的用户业务共享同一个区块链系统资源，主要包括网络、计算、存储资源。从本质上来说，channels是通过不同的区块链账本来为上层业务服务，而且，这些区块链统一部署在peers节点上，统一通过ordering service进行交易排序和打包区块。Channels之间通过权限隔离控制，不同channel内的成员，无法访问对方的交易信息，只能访问所属channel的交易信息。

![img](http://youqudian.com/uploads/allimg/200412/04302U207-6.png)

channel可以理解为系统资源的逻辑单元，每个channel都包含peers资源、order资源、网络资源等等，而且这些资源有可能是和其它channel所共享。

## 8.State Database

状态数据库保存了账本所有资产的最新状态（例如，账户A拥有某种资产的总量），同时，为智能合约提供了丰富的资产查询语义。所有的资产信息最终以文件形式记录在区块链账本中，而数据库是区块链账本的视图表现形式，能够让智能合约更加高效的和账本信息进行交互。数据库自动从底层区块链账本中更新或者恢复数据，默认的状态数据库是LevelDB，也可以替换为CouchDB。

LevelDB：Hyperledger Fabric的默认数据库，简单的存储键值对信息；

CouchDB：提供更加丰富的查询语义，可以保存JSON对象，以及范围key的查询。

![img](http://youqudian.com/uploads/allimg/200412/04302V596-7.png)

## 9.Smart Contract

智能合约就是一段部署在区块链账本中的计算机程序，用于执行交易以及修改资产的状态。在Hyperledger Fabric中，智能合约被称作chaincode，使用Go语言编写。

## 10.Membership Service Provider（MSP）

Hyperledger Fabric是一种permissioned blockchain，所有的节点都是必须经过授权后才能访问区块链网络（比特币属于permissionless blockchain）。MSP是Hyperledger Fabric中的身份认证模块，用于对用户身份的校验、授权以及网络访问权限控制。默认的MSP接口是Fabric-CA API，同时，开发者可以根据自身的业务需要，实现自己的身份认证接口，对接MSP。Hyperledger Fabric网络可以被多个MSP控制，用以满足各个组织的需要。

![img](http://youqudian.com/uploads/allimg/200412/04302Q543-8.png)

## 11.Hyperledger Fabric的商业价值

随着比特币、以太坊等一系列虚拟货币的疯涨，区块链一度被认为是可以颠覆互联网的下一代革命性技术。但是，我们需要清楚的意识到，虚拟货币的价格并不能代表区块链技术能够为人类创造的实际价值，而且，大多数区块链应用本身并不需要token激励机制。所以，区块链未来的发展取决于能否广泛应用在商业领域，为人类的生活生产降低成本、提高效率。

Hyperledger Fabric的意义在于，迈出了区块链向商业领域进军的第一步，未来会有更多的区块链项目在各个行业中创造价值。

本文由作者极简主义首发于资讯，未经允许不得转载。



# 8.超全总结Hyperledger Fabric架构原理

## 1.概述

Hyperledger Fabric是由IBM公司主导开发的一个面向企业级客户的开源项目。与比特币和以太坊这类公有链不同，Hyperledger Fabric网络中的节点必须经过授权认证后才能加入，从而避免了POW资源开销，大幅提高了交易处理效率，满足企业级应用对处理性能的诉求。同时，为了满足灵活多变的应用场景，Hyperledger Fabric采用了高度模块化的系统设计理念，将联盟成员权限认证模块（MSP）、共识服务模块（Ordering Service）、背书模块（Endorsing peers）、区块提交模块（committing peers）等进行分离部署，使开发者可以根据具体的业务场景替换模块，实现了模块的插件式管理（plug-in/plug-out）。所以，Hyperledger Fabric是一个私有链／联盟链的开发框架，而且系统的运行不需要token支持。

## 2.基本概念

### Ledger

账本，节点维护的区块链和状态数据库。

### World state

世界状态，经过数次交易后最新的键值对，世界状态是一个数据库，存储的是账本当前值。作用是能够快速获取账本最新值而不必根据交易日志从头开始计算。世界状态中还有一个属性——版本号，版本号从0开始，每当状态更新时版本号就递增。状态更新时会首先检查版本号，以确保当前状态的版本与背书时的版本一致（避免并发更新）。

### Channel

通道，私有的子网络，通道中的节点共同维护账本，实现数据的隔离和保密。 每个channel对应一个账本，由加入该channel的peer维护，一个peer可以加入多个channel，维护多个账本。

### Org

Orginazation，管理一系列成员的组织。一个channel内可以有多个组织。

### Chaincode

链码(智能合约)，运行在节点内的程序，提供业务逻辑接口，对账本进行查询或更新

### Endorse

背书，指一个节点执行了一个交易并对结果进行签名后返回响应的过程。
Ordering Service：排序服务，将交易排序后放入区块中，并广播给网络各节点

### PKI

Public Key Infrastructure，一种遵循标准的利用公钥加密技术为电子商务的开展提供一套安全基础平台的技术和规范

### MSP

Membership Service Provider，成员管理服务，基于PKI实现，为网络成员生成证书，并管理身份

### 节点

在Fabric v1.x中把节点分为peer节点（维护state和ledger）,Order节点（负责共识或者排序账本中的交易）和背书节点（负责执行交易和链码）而在Fabric v0.6中则没有这些概念，只有一个peer节点，peer节点同时完成上述所有的功能。

这样设计的优势
链码（Chaincode）执行信任的可伸缩性，将用户自己开发的链码和系统提供的Order服务拆分，用户开发的链码和系统提供的Order服务不再是一一对应的关系，Order也可以适当容忍错误的出现，增强了系统的鲁棒性。
可扩展性，拆分链码和Order的串行执行，在原有架构中，当链码执行非常耗时的时候，Order将会处于闲置状态，不利于提高系统的吞吐量，拆分以后链码和order可以并行执行。
共识机制可以单独实现（Order）。

## 3.共识算法

在所有peers中，交易信息必须按照一致的顺序写入账本（区块链的基本原则）。例如，比特币通过POW机制，由最先完成数学难题的节点决定本次区块中的信息顺序，并广播给全网所有节点，以此来达成账本的共识。而Hyperledger Fabric采用了更加灵活、高效的共识算法，以适应企业场景下，对高TPS的要求。目前，Hyperledger Fabric有三种交易排序算法可以选择。

### SOLO

只有一个order服务节点负责接收交易信息并排序，这是最简单的一种排序算法，一般用在实验室测试环境中。Sole属于中心化的处理方式。

### Kafka

是Apache的一个开源项目，主要提供分布式的消息处理／分发服务，每个kafka集群由多个服务节点组成。Hyperledger Fabric利用kafka对交易信息进行排序处理，提供高吞吐、低延时的处理能力，并且在集群内部支持节点故障容错。

### SBFT

简单拜占庭算法，相比于kafka，提供更加可靠的排序算法，包括容忍节点故障以及一定数量的恶意节点。目前，Hyperledger Fabric社区正在开发该算法。

## 4.节点类型

节点（peer）是区块链的通信实体，是一个逻辑概念，不同类型的多个节点可以运行在同一个物理服务器上。节点主要有以下四种：

### 客户端节点

客户端必须连接到某一个peer节点或排序服务节点上才能与区块链网络进行通信。客户端向背书节点（endorser）提交交易提案（transaction proposal），当收集到足够背书后，向排序服务广播交易提案，进行排序，生成区块。

### 普通节点peer

![image](https://imgconvert.csdnimg.cn/aHR0cHM6Ly90dmExLnNpbmFpbWcuY24vbGFyZ2UvMDA4MzFyU1RseTFnZDcxd25jeG45ajMxMXkwcmVqdTEuanBn?x-oss-process=image/format,png)


peer节点根据所承担的角色又可以分为记账节点（committer）、背书节点（endorser）、主节点（leader）和锚节点（anchor）。

### 记账节点

所有的peer节点都是记账节点（committer），负责验证排序服务节点区块里的交易，维护状态和总账（Ledger）的副本。该节点会定期从orderer节点获取包含交易的区块，在对这些区块进行核发验证之后，会把这些区块加入到区块链中。committer节点无法通过配置文件配置，需要在当前客户端或者命令行发起交易请求的时候手动指定相关的committer节点。记账节点可以有多个。

### 背书节点

部分节点还会执行交易并对结果进行签名背书，充当背书节点（endorser）的角色。背书节点是动态的角色，是与具体链码绑定的。每个链码在实例化的时候都会设置背书策略，指定哪些节点对交易背书后交易才是有效的。并且只有应用程序向它发起交易背书请求的时候才是背书节点，其他时候都是普通的记账节点，只负责验证交易并记账。背书节点也无法通过配置文件指定，而是由发起交易请求的客户端指定。背书节点可以有多个。

### 主节点

peer节点还可以是主节点（leader peer），能与排序服务节点通信，负责从排序服务节点获取最新的区块并在组织内部同步。主节点在整个组织中只能有一个。

### 锚节点

peer节点还可以是锚节点（anchor peer），锚节点主要负责代表组织和其他组织进行信息交换。每个组织都有一个锚节点，锚节点对于组织来说非常重要，如果锚节点出现问题，当前组织就会与其他组织失去联系。锚节点的配置信息是在configtxgen模块的配置文件configtx.yaml中配置的。锚节点只能有一个。

### 排序服务节点orderer

接收包含背书签名的交易，对未打包的交易进行排序生成区块，广播给peer节点。排序服务提供的是原子广播，保证同一个链上的节点接收到相同的信息，并且有相同的逻辑顺序。

### CA节点

fabric1.0的证书颁发机构，由服务器和客户端组成。CA节点接收客户端的注册申请，返回注册密码用于用户登录，以便获取身份证书。区块链上的所有操作都需要验证用户身份。

原文链接：https://blog.csdn.net/ASN_forever/article/details/86538915



## 5.交易流程

以下是fabric的经典交易流程，所有涉及到对账本数据更新的操作都是基于这个交易流程来完成的。

![image](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9waWM0LnpoaW1nLmNvbS84MC92Mi00OTRjMDY2Y2VlZTkzMzVkYjdlNmE0Yzk1YjFhODRhM18xNDQwdy5qcGc?x-oss-process=image/format,png)

### 1.发送交易提案

客户端发送交易提案（Proposal）到背书节点（peer节点），提案中包含交易所需参数。

### 2.模拟执行交易提案

背书节点会调用链码模拟执行交易提案(Proposal)，这些执行不会更新账本。

每个执行都会产生对状态数据读出和写入的数据集合，叫做读写集（RWsets），读写集是交易中记录的主要内容。

### 3.返回提案响应

背书节点会对读写集进行背书(Endorse)签名，生成提案响应(Proposal response)并返回给应用程序。

### 4.交易排序

应用程序根据接收到的提案响应生成交易，并发送给排序服务节==（Order节点）==。排序服务打包一组交易到一个区块后，分发给各记账节点。

### 5.交易验证并提交

每个节点会对区块中的所有交易进行验证，包括验证背书策略以及版本冲突验证（防止双花），验证不通过的交易会被标记会无效（Invalid）。

账本更新：节点将读写集更新到状态数据库 ，将区块提交到区块链上。

### 6.通知交易结果给客户端

各记账节点通知应用程序交易的成功与否，交易完成。

在Fabric v1.0中，仅在网络上发送签名和读写集，所以可伸缩性和性能得到了优化，因为只有背书者和提交者能看到交易，所以整个网络所需要的信任更低，提供了更高的安全性。

每个ChainCode在部署时，都需要和背书（ESCC）和验证（VSCC）的系统ChainCode关联。

ESCC决定如何对proposal进⾏背书。

VSCC决定事务的有效性（包括背书的正确性）。

### 7.交易两种类型

链码的部署是通过一个带参数的交易来进行的，当执行成功则一个链码被部署到区块链上
调用部署好的链码修改相应的状态并且返回
即两种，部署交易（deploy transaction）和调用交易（invoke transaction）。

## 6.Blockchain数据结构

### 状态（State）

区块链最新的未被持久化的状态是带版本号的键值对（KVS），这些状态被智能合约通过put，get方法操作，并且会被日志记录，如果有其他的RDBMS解决方案都可以灵活的替换。

智能合约可以根据key的名字来识别是属于哪个智能合约，原则上一个合约可以读取所有属于它自身的所有key。跨合约交易（cross-transactions）修改state目前还不支持，属于v1后续版本。

### 账本（Ledger）

账本提供了所有state成功修改和对state不成功修改的尝试的记录。

账本通过order把所有有效和无效的交易构建了一个完全有序的链，而其中的每一个block中的交易又是完全有序的，这样就保证了所有交易的有序性。

账本保存在Peer或者可以保存在order的子集中，二者唯一的区别是Peer中保存的账本可以通过bitmask来区分交易的有效性。

Ledger允许重做所有transactions的历史记录，并且重建state。

## 7.Fabric系统逻辑结构图

![image](https://img-blog.csdnimg.cn/20190217094400824.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0FTTl9mb3JldmVy,size_16,color_FFFFFF,t_70)



### 通道channel

fabric的数据存储结构被设计成多账本体系，每个账本在fabric中被称为channel。每个channel中都有一个完全独立的账本。同一个channel中的所有peer节点都保存一份相同的数据。

通道是两个或多个特定网络成员之间进行通信的专用“子网”，用于进行私有和机密事务。通道由成员（组织）、每个成员的锚节点、账本、链码应用程序和排序服务节点定义。网络上的每个交易都是在一个通道上执行的，在该通道上，每一方都必须经过身份验证和授权才能在该通道上进行交易。加入通道的每一个peer都有其自己的身份，由成员服务提供者（MSP）提供。

为通道上的每个成员选择主节点，用于代表该成员与排序服务进行通信。算法会自动选出主节点。共识服务对交易进行排序打包成区块，并把区块它发送给主节点。然后主节点根据gossip协议将区块分发给 通道中的成员。

尽管一个锚节点可以属于多个通道，因此能够维护多个账本，但是任何账本数据都不能从一个通道传递到另一个通道。这种用通道对账本进行隔离的方式是由配置链码、成员身份服务和gossip数据分发协议共同定义和实现的。通道上只有具备验证成员资格功能的节点才有权限进行数据分发。这种将peers和账本数据进行通道隔离的方式使得网络成员能够在同一个区块链中进行私密交易（只有自己通道的成员才知道交易信息）。

https://zhuanlan.zhihu.com/p/55341714
https://zhuanlan.zhihu.com/p/38289914

### 网络及配置

网络则主要包括两大部分：用户证书和链码。

接口主要是fabric sdk，比如fabric java sdk和fabric node sdk。

网络基础配置=证书+yaml文件

网络的基本单元是容器，网络是由一个个的容器组成的，如peer容器，orderer容器等。其中，容器是由yaml文件和证书两部分组成的。

yaml文件和docker联系紧密，所以你需要掌握docker的使用方法。yaml文件描述了peer的配置，而peer是网络的重要组成部分，所以如何实现动态的增加用户(peer\user)，还得需要使用kubernates进行管理。

在用户证书这部分，由于官方1.1版本的的例子中使用的是cryptogen二进制工具，意思就是说官方1.1版本的例子只能安装事前配置的用户的数量生成相应数量的用户证书，所以没法实现动态生成peer\user证书从而动态增加用户(peer\user)。于是出现了fabric-ca，fabric-ca即是解决这样的需求的。通过使用fabric-ca动态生成证书，你将对fabric证书体系的理解更上一个层次。

在网络配置上运行着的是链码
就好像电线架好了，电就可以传输了一样，基础网络配置好了后，那在网络上运行着的便是链码了。

在fabric中链码和peer是独立的容器，因此你可以独立编辑的你的链码文件，只要实现了规定的链码接口，那么你这个链码便可以在peer上运行了，现在链码文件大都是go语言实现的。
————————————————
版权声明：本文为CSDN博主「yhc166188」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/yhc166188/article/details/105110270



# 9.Fabric v1.x 账本与状态数据库



## 1、Fabric账本

Fabric账本是有序的、不可篡改的状态转换记录，包括区块链（Blockchain）和世界状态（World stat）两部分。

区块链中保存着不可变的顺序记录，包含配置记录，例如channel的配置；还包含全部交易记录；
世界状态中维护账本的当前状态，方便Appication快速查询

## 2、区块链

区块链是一个历史交易记录，记录着所有数据对象是如何到达当前状态的。
下图中有4个区块B0、B1、B2、B3，第一个区块B0为创世区块（genesis block），保存一些配置信息，包括Order、peer的信息和证书信息；后面的区块B1、B2、B3则保存着后续交易信息：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200219173206119.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

### 2.1 区块信息

区块分为3部分，分别为区块头（Block header）、区块数据（Block Data）、区块元数据（Block Metadata）：

区块头里面包含区块序号（Block number）、当前区块哈希（Current Block Hash）、上一个区块哈希（Previous Block Hash），
区块数据就是一系列交易数据；
区块元数据主要包含区块写入时间、写入的人以及签名。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200219173452393.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

### 2.2 交易信息

交易（Transaction）保存着捕获的关于交易的一些基本元数据：

transaction header：chaincode名字、版本等信息
signature：client的签名
proposal：client的提案，主要是一些输入参数
response：智能合约执行前后的数据
endorsement：所有背书节点各自返回的背书结果，由背书策略而定，可能有多个

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200219173629999.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3、世界状态

世界状态维护着一组业务对象的当前状态。采用KV数据库的形式，每组数据包含key、value以及version，每次修改数据时version都会增加。查询数据时，直接通过状态数据库就可以快速地获取数据。如下所示：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200219175402121.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

下图列出了账本中的数据示例：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200219175809123.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

————————————————
版权声明：本文为CSDN博主「ice-fire-x」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/ice_fire_x/article/details/104362219



# 10.Fabric v1.x 私有数据（PrivateData）模型分析

## 1、私有数据机制介绍

对于一些组织来说，总会有一些机密数据，比如产品价格、用户电话、用户地址等信息，这些数据可能希望只被某几个组织知道，或者仅仅自己知道。可以将这些机密数据存储到私有数据库中，然后通过私有数据收集策略来定义有哪些peer有权获得这些私有数据，这些授权peer同样是将这些数据存储到私有数据库里，对于ordering来说，也只能看到私有数据的哈希。
另外私有数据同样是通过Gossip协议在peer间传输的。

## 2、私有数据的交易流程

在之前一篇文章中的交易流程章节分析了普通交易的详细流程。私有数据的交易流程与之相比有些许不同，具体流程如下：

首先由client请求关于私有数据的交易提案；
然后背书节点仿真执行交易，并把私有数据存储到一个临时数据库里；
背书节点同时将这些私有数据通过gossip协议传播给其他有权限的记账节点，这些记账节点同样将私有数据存储到临时数据库里；
传播到指定的数量后背书节点将结果进行签名，返回给client端的数据不包含私有数据，只会返回签名和私有数据key和value的哈希值；
client将提案的响应发送给ordering后，ordering无法看到私有数据的内容，只能看到一串哈希值；
然后ordering将交易区块传输给记账节点，记账节点会进行一些验证，除了验证背书策略，还要验证ordering发送的哈希值与本地临时数据库里保存的私有数据是否匹配，如果验证成功，记账节点就会将临时数据库中的数据正式存储到私有数据库里；
记账节点存储完成后会通知客户端记账的情况。
根据上述的步骤整理出交易流程图：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200221111135648.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

## 3、私有数据集合配置

在chaincode实例化的时候需要指定私有数据收集策略配置文件路径，需要在配置文件中指定名称、收集策略、传播的节点数量、私有数据保存时间等选项，具体实例如下所示：

```
[
    {
        "name": "collectionMarbles",
        "policy": "OR('Org1MSP.member', 'Org2MSP.member')", // Org1和Org2都有权限访问
        "requiredPeerCount": 0, // Endorsing peer返回执行结果前必须传播给Committing peer的节点最小数量
        "maxPeerCount": 3,    // 最多要传播的节点数量
        "blockToLive":1000000,    // 私有数据需要在DB里保存的时间，超过这个期限私有数据会自动销毁
        "memberOnlyRead": true    // 只有策略定义的member可以访问
    },
    {
        "name": "collectionMarblePrivateDetails",
        "policy": "OR('Org1MSP.member')",    // 只有Org1可以访问
        "requiredPeerCount": 0,
        "maxPeerCount": 3,
        "blockToLive":3,
        "memberOnlyRead": true
    }
]
```



## 4、私有数据存储的代码分析

下面是私有数据保存的代码实现，第一段代码使用“collectionMarbles”配置来存储私有数据，包括Name、Color、Size、Owner这些隐私级别较低的数据，根据上面的配置可知，组织Org1MSP和Org2MSP的成员都可以访问这些数据，具体代码如下所示：

四、私有数据存储的代码分析
下面是私有数据保存的代码实现，第一段代码使用“collectionMarbles”配置来存储私有数据，包括Name、Color、Size、Owner这些隐私级别较低的数据，根据上面的配置可知，组织Org1MSP和Org2MSP的成员都可以访问这些数据，具体代码如下所示：

```
// ==== Create marble object, marshal to JSON, and save to state ====
marble := &marble{
    ObjectType: "marble",
    Name: marbleInput.Name,
    Color: marbleInput.Color,
    Size: marbleInput.Size,
    Owner: marbleInput.Owner,
}
marbleJSONasBytes, err := json.Marshal(marble)
// === Save marble to state ===
err = stub.PutPrivateData("collectionMarbles", marbleInput.Name, marbleJSONasBytes)
```


第二段代码使用“collectionMarblePrivateDetails”配置存储私有数据，包括price这种隐私级别较高的数据，根据上面的配置可知，只有Org1MSP的成员可以访问，具体代码如下所示：

第二段代码使用“collectionMarblePrivateDetails”配置存储私有数据，包括price这种隐私级别较高的数据，根据上面的配置可知，只有Org1MSP的成员可以访问，具体代码如下所示：

```
// ==== Create marble private details object with price, marshal to JSON, and save to state ====
marblePrivateDetails := &marblePrivateDetails{
    ObjectType: "marblePrivateDetails",
    Name: marbleInput.Name,
    Price: marbleInput.Price,
}
marblePrivateDetailsBytes, err := json.Marshal(marblePrivateDetails)
err = stub.PutPrivateData("collectionMarblePrivateDetails", marbleInput.Name, marblePrivateDetailsBytes)
```


最终org1和org2的peer成员都会存储marble对象，而marblePrivateDetails对象只会存储到org1的peer节点中，两个组织的存储图如下所示：

最终org1和org2的peer成员都会存储marble对象，而marblePrivateDetails对象只会存储到org1的peer节点中，两个组织的存储图如下所示：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200221130751514.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200221130803112.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)



————————————————
版权声明：本文为CSDN博主「ice-fire-x」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/ice_fire_x/article/details/104362297



# 11.Fabric v1.x MSP的结构和使用方法

## 1、MSP是什么

### 1.1 相关概念

#### 证书：

证书是fabric权限管理的基础，采用ecdsa算法，符合X.509标准，通过CA来签发证书，每个identity或organization都可以拥有自己的身份证书。

#### 组织：

证书遵循组织结构，通过组织实现灵活的权限管理，组织一般包含名称、ID、MSP、管理策略、anchor peer节点位置信息。

#### 成员：

peer和client都是成员，peer提供节点服务，在跟不同组织交互的时候，同一个组织的成员节点一般被认为是同一个身份，然后代表组织进行签名。

#### 联盟：

由多个组织组成，在联盟中，每个组织都要有自己的MSP ID

### 1.2 MSP（Membership Service Provider）

MSP（Membership Service Provider）是从一个从fabric CA里enroll的一个identity实体到一个organization的映射，确定如何将一个节点分配一个特定角色，并且获得适当的访问权限，抽象了颁发和验证证书（MSP向clients提供认证交易的凭证，以及向peers提供背书的凭证），以及身份验证背后的加密机制和协议。MSP是可插拔的接口，一个MSP是可以通过管理员的配置来更换或插拔，可以定义一些验证和签名的规则。Fabric可由多个MSP管理，提供模块化的成员操作，保证不同成员标准和结构的互操作性。但是同一个channel上MSP ID必须是唯一的，如果有两个相同的启动会失败。

### 1.3 MSP的验证元素

MSP的验证元素包括：

```
MSP Identifier
Root CAs
Intermediate CAs
Admin CAs
OU List
CRLs
```



## 2、MSP的结构

Fabric需要在每个peer、orderer、channel上指定MSP，启用这些节点的身份验证、签名验证的功能。

**MSP的参数来自于RFC5280，包括：**

```
自签名的X.509证书列表构成信任根，以及TLS证书的TLS信任根
X.509证书列表，需要带有可验证的证书路径，以表示管理员的MSP，并有权请求对此MSP的配置进行更改（例如Root CA、intermediate CA）
中间证书由信任根中的其中一个证书精确认证
OU列表
CRL
```

**MSP的身份合法条件：**

```
符合X.509证书标准，提供可验证的证书
MSP不能包含在任何CRL中
在其X.509证书结构的OU字段中可以设置一个或多个MSP的OU（组织单元）
```

**MSP文件树如下：**

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303130118339.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

Root CAs、Intermediate CAs是签发证书的CA的证书；
Organizational Units是可选的，默认不会生成，需要在msp目录下的config.yaml文件里定义organization；
Revoked Certificates、Administrators、TLS Root CAs、TLS Intermediate CAs同样是可选的；
Signing Certificates是节点本身的证书；
Keystore就是私钥，用于签名；

**组织与MSP的映射关系**
一般是在定义好了组织结构之后，再来命名MSP，这样就是比较节省资源。
一个organization可以不同的department，具有不同的功能和职责，可能在区块链网络里是属于不同channel的，一般需要建立多个MSP，来划分组织权限结构。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303131422129.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

**MSP在channel上的定义**

Local MSP：对应的是peer节点，在节点中给对应组织生成MSP。主要用来管理peer节点、orderer节点；
Global MSP（Channel MSP）：channel里所有organization共享的MSP集合 ，会把所有配置拷贝到Peer节点。主要用来管理channel资源、账本、智能合约等。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303132000332.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

**MSP的层级**
MSP必须指定在某个层级上，这是由于channel MSP和Local MSP的分割造成的，用来管理组织和本地资源，层级从高到底有网络层、channel的Global MSP、Peer节点、orderer节点。层级化的MSP要求最好事先做好organization的设置。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303132745213.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

Peer和Orderer的MSP配置
Peer和Orderer的MSP配置基本是一样的，示例如下：

Path on the file system where peer will find MSP local configurations

    mspConfigPath: msp
    
    # Identifier of the local MSP
    # ----!!!!IMPORTANT!!!-!!!IMPORTANT!!!-!!!IMPORTANT!!!!----
    # Deployers need to change the value of the localMspId string.
    # In particular, the name of the local MSP ID of a peer needs
    # to match the name of one of the MSPs in each of the channel
    # that this peer is a member of. Otherwise this peer's messages
    # will not be identified as valid by other nodes.
    localMspId: SampleOrg
Type for the local MSP - by default it's of type bccsp

    localMspType: bccsp

MSP的生命周期
用Fabric CA Client做enroll的时候，就会自动生成一个MSP的目录结构，也就是生成了一个MSP；

对于X.509的CRL：MSP的信息如果被加到CRL中，用户在做验证的时候会检查CRL，如果验证的这个人在黑名单里，就不应该有任何权限了；
Idemix CRL：新的revoke功能，还没实现完成。

## 3、MSP实践

推荐一个组织对应一个MSP，逻辑非常清晰，实际中也可以1个组织对应多个MSP，或多个组织对应一个MSP。
授予不同部门访问不同channel的不同访问权限；
需要将管理者Admin的证书和CA证书分开；
可以把某个被攻击的中间CA加到黑名单，这样该CA颁发的所有证书都会失效；
推荐把CA和TLS CA负责各自的功能，不要放在一起。

## 4、TLS通信的证书

channel上节点间通信都是基于TLS的，以保证数据安全和可信性。

单向（只有服务器）
双向（服务器-客户端）
Peer节点既是TLS服务器又是TLS客户端

TLS server：其他peer节点、application、客户端都会连接它
TLS client：连接其他peer节点或orderer
TLS client的身份验证默认是关闭的，如果需要双边通信则要启用该功能。
在channel里，channel成员的root CA证书链是从config区块里读取的，然后添加到TLS客户端和服务器的Root CAs的数据结构。因此，peer与peer、peer与orderer之间的通信是无缝的。

TLS通信的握手过程如下图所示：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303151845303.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ljZV9maXJlX3g=,size_16,color_FFFFFF,t_70)

————————————————
版权声明：本文为CSDN博主「ice-fire-x」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/ice_fire_x/article/details/104628346







# Fabric v1.x 应用开发指南

文章目录
1 技术栈
1.1 经典软件工程管理
1.2 PKI密码体系
1.3 chaincode开发语言
1.4 devOps运维开发
2 chaincode API解析
2.1 chaincode代码示例
2.2 chaincode接口（shim.ChaincodeStubInterface）
获取参数（Params）
访问世界状态（Access world states）
历史记录查询（Witness the immutable）
键值范围查询（List for NoSQL: key range）
复合key（Composite key）
富查询（CouchDB or default levelDB）
2.3 复合API
跨智能合约调用（Cross chaincode invoke）
客户端身份解析（Client Identity Chaincode Library）
获得Transient对象（GetTransient）
私有数据（PrivateData）
Key级别的背书（Key level endorsement）
3 Init函数
4 Invoke VS Query
5 开发模式
6 开发周期
6.1 peer部分
(1) 生成加密物料（Generate crypto material）
(2) 生成创世区块（Generate genesis block file）
(3) 启动orderer和peer（Start orderer and peer）
6.2 channel部分
(1) 生成channel引导文件（Generate channel bootstrap file）
(2) 创建channel（Create channel）
(3) Peer节点加入到channel（Peer join channel）
6.3 chaincode生命周期
(1) 打包与安装chaincode（Package, install chaincode）
(2) 实例化chaincode（Instantiate chaincode）
(3) 升级chaincode（Upgrade chaincode）
7、 黑盒开发



## 1 技术栈

### 1.1 经典软件工程管理

Fabric应用开发者需要了解经典软件工程管理：

- 依赖管理（调用第三方库）：govendor/dep, npm/yarn, gradle/maven, pip
- 异常处理：defer, saync/await

- 测试流水线：Smoke, Unit/Mock, SI tests

### 1.2 PKI密码体系

Fabric应用开发者需要精通PKI密码体系，包括：

- ECDSA
- X.509
- HSM和pkcs11

### 1.3 chaincode开发语言

chaincode开发支持的语言包括：

- Golang
- Node.js
- Java

### 1.4 devOps运维开发

devOps运维开发技术栈包括：

- 受认证的Fabric管理员（Certified Hyperledger Fabric Administrator）：CLI in Unix/Linux
- SDK语言：Java/Golang/Node.js/Python
- Docker
- gRPC

## 2 chaincode API解析

### 2.1 chaincode代码示例

chaincode的package包名必须是main，因为容器里会去找main函数。
必须实现Init和Invoke两个函数，这样才会构成一个chaincode。在初始化和升级chaincode时，会调用到Init函数；query以及invoke的第一阶段，会调用到Invoke函数。

```
package main

type StressChaincode struct{}

//called when initialize, upgrade
func (t *StressChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

//called when query, phase-1 of invoke
func (t *StressChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

func main() {
    err:=shim.Start(new(StressChaincode))
}

```

### 2.2 chaincode接口（shim.ChaincodeStubInterface）

### 获取参数（Params）

从外部传入chaincode的参数列表可以通过下面的两个接口获取到。GetFunctionAndParameters可以将第一个元素转换成功能名称。

```
GetArgs() [][]byte

GetFunctionAndParameters() (string, []string) 
//returns 1st argument as function，"fcn", rest as params
```



# 超级账本Fabric框架介绍（一）

## Gaining Network Membership

> Hyperledger Fabric is a permissioned network, meaning that only participants who have been approved can gain entry to the network. To handle network membership and identity, membership service providers (MSP) manage user IDs, and authenticate all the participants in the network. A Hyperledger Fabric blockchain network can be governed by one or more MSPs. This provides modularity of membership operations, and interoperability across different membership standards and architectures.
> In our scenario, the regulator, the approved fishermen, and the approved restaurateurs should be the only ones allowed to join the network. To achieve this, a membership service provider (MSP) is defined to accommodate membership for all members of this supply chain. In configuring this MSP, certificates and membership identities are created. Policies are then defined to dictate the read/write policies of a channel, or the endorsement policies of a chaincode.
> Our scenario has two separate chaincodes, which are run on three separate channels. The two chaincodes are: one for the price agreement between the fisherman and the restaurateur, and one for the transfer of tuna. The three channels are: one for the price agreement between Sarah and Miriam; one for the price agreement between Sarah and Carl; and one for the transfer of tuna. Each member of this network knows about each other and their identity. The channels provide privacy and confidentiality of transactions.
> In Hyperledger Fabric, MSPs also allow for dynamic membership to add or remove members to maintain integrity and operation of the supply chain. For example, if Sarah was found to be catching her fish illegally, she can have her membership revoked, without compromising the rest of the network. This feature is critical, especially for enterprise applications, where business relationships change over time.

Hyperledger Fabric是一个基于授权的网络，这意味着只有经过批准的参与者才能进入网络。为了处理网络成员身份和身份，成员服务提供者(MSP)管理用户id，并对网络中的所有参与者进行身份验证。一个Fabric区块链网络可以由一个或多个MSPs控制。这提供了成员操作的模块化，以及跨不同成员标准和体系结构的互操作性。

在我们的方案中，监管机构、被批准的渔民和批准的餐馆老板应该是唯一允许加入网络的角色。为了实现这一点，定义了一个成员服务提供者(MSP)，以管理该供应链所有成员的成员资格。在配置这个MSP时，将创建证书和成员身份标识。然后定义策略来指定通道的读/写策略，或chaincode的支持策略。

我们的场景有两个单独的链码，它们分别运行在三个不同的通道上。这两个链码是:一个是渔民和餐馆老板之间的价格协议，另一个是金枪鱼的转让。这三个通道是: 一个是莎拉和米里亚姆之间的价格协议; 一个是莎拉和卡尔之间的价格协议; 还有一个用来转移金枪鱼。这个网络的每个成员都知道对方的身份和身份。这些通道为交易提供了隐私和机密性。

在Hyperledger Fabric中，MSPs还允许动态成员添加或删除成员以维护供应链的完整性和操作。例如，如果萨拉被发现非法捕捞她的鱼，她可以让她的成员被撤销，而不损害其他的网络。这个特性非常重要，特别是对于企业应用程序来说，业务关系随着时间而变化。

## Summary of Demonstrated Scenario

> Below is a summary of the tuna catch scenario presented in this section:
> \1. Sarah catches a tuna and uses the supply chain application’s user interface to record all the details about the catch to the ledger. Before it reaches the ledger, the transaction is passed to the endorsing peers on the network, where it is then endorsed. The endorsed transaction is sent to the ordering service, to be ordered into a block. This block is then sent to the committing peers in the network, where it is committed after being validated.
> \2. As the tuna is passed along the supply chain, regulators may use their own application to query the ledger for details about specific catches (excluding price, since they do not have access to the price-related chaincode).
> \3. Sarah may enter into an agreement with a restaurateur Carl, and agree on a price of 80 per pound. They use the blue channel for the chaincode contract stipulating 80/lb. The blue channel's ledger is updated with a block containing this transaction.
> \4. In a separate business agreement, Sarah and Miriam agree on a special price of 50 per pound. They use the red channel's chaincode contract stipulating 50/lb. The red channel's ledger is updated with a block containing this transaction.



![img](https://pic1.zhimg.com/80/v2-61bbd7eb188c20cc6428391379df45e0_720w.jpg)

下面是本节介绍的金枪鱼供应链场景的摘要:

莎拉抓住了一条金枪鱼，并使用了供应链应用的用户界面，记录了所有关于捕获的细节。在它到达分布式账本之前，交易被传递给网络上的背书节点，然后被背书。已被背书的事务被发送到排序节点，将被排序到一个块中。然后该块被发送到网络中的提交节点，在经过验证后提交。
当金枪鱼在供应链中被传递时，监管机构可能会使用他们自己的应用程序来查询关于特定渔获物的详细信息(不包括价格，因为他们无法获得与价格相关的链码)。
莎拉可能会和一个餐馆老板达成协议，并同意每磅80美元的价格。他们使用蓝色通道的chaincode合同规定$80/lb。蓝色通道的分类帐将被更新为包含该事务的区块。
在另一份商业协议中，莎拉和米利安同意每磅50美元的特价。他们使用红色通道的chaincode合同规定50美元/磅。红色通道的分类帐将更新为包含该事务的区块。

## Roles within a Hyperledger Fabric Network

> There are three different types of roles within a Hyperledger Fabric network:
> \- - **Clients**
> Clients are applications that act on behalf of a person to propose transactions on the network.
> \- **Peers**
> Peers maintain the state of the network and a copy of the ledger. There are two different types of peers: **endorsing** and **committing**
> peers. However, there is an overlap between endorsing and committing
> peers, in that endorsing peers are a special kind of committing peers.
> All peers commit blocks to the distributed ledger.
> \- *Endorsers* simulate and endorse transactions
> \- *Committers* verify endorsements and validate transaction results, prior to committing transactions to the blockchain.
> \- **Ordering Service**
> The ordering service accepts endorsed transactions, orders them into a block, and delivers the blocks to the committing peers.

在`Fabric`的网络中有三个不同的角色：

1. 客户端：为现实社会中的参与者提供与区块链网络交互的接口
2. 节点：节点维持着整个网络中账本的备份。节点分为两种：背书节点和提交节点，两个功能可能在物理上从属于同一个节点。背书节点的功能为对交易进行模拟及背书；提交节点的功能为校验背书状态和交易的合法性，并提交交易到超级账本。
3. 排序节点：排序节点接收背书之后的交易，将所有交易排序至一个区块，将区块发送至提交节点写入超级账本。

## How to Reach Consensus

> In a distributed ledger system, **consensus** is the process of reaching agreement on the next set of transactions to be added to the ledger. In Hyperledger Fabric, consensus is made up of three distinct steps:
> \- - Transaction endorsement
> \- Ordering
> \- Validation and commitment.
> These three steps ensure the policies of a network are upheld. We will explore how these steps are implemented by exploring the transaction flow.

共识机制用于分布式账本进行下一次生成区块之前达成共识的方法。在`Fabric`中，共识机制分为三步：

1. 交易背书
2. 交易排序
3. 交易校验及提交

## Transaction Flow (Step 1)

> Within a Hyperledger Fabric network, transactions start out with client applications sending transaction proposals, or, in other words, proposing a transaction to endorsing peers.

![img](https://pic1.zhimg.com/80/v2-21431955acd5b7888ca8d393c94deaf8_720w.jpg)



> **Client applications** are commonly referred to as **applications** or **clients**, and allow people to communicate with the blockchain network. Application developers can leverage the Hyperledger Fabric network through the application SDK.

在`Fabric`的架构中，交易起始于客户端应用发送的交易申请（提交交易至背书节点）。客户端应用允许参与者与区块链进行通信。应用的开发者可以利用 SDK 和超级账本进行交互。

## Transaction Flow (Step 2)

> Each endorsing peer simulates the proposed transaction, without updating the ledger. The endorsing peers will capture the set of **R**ead and **W**ritten data, called **RW Sets**. These RW sets capture what was read from the current world state while simulating the transaction, as well as what would have been written to the world state had the transaction been executed. These RW sets are then signed by the endorsing peer, and returned to the client application to be used in future steps of the transaction flow.

![img](https://pic1.zhimg.com/80/v2-13e5a6a80c0e150f46d45ec0634b86b8_720w.jpg)

> Endorsing peers must hold smart contracts in order to simulate the transaction proposals.

每一个背书节点对交易申请进行模拟，模拟交易将要读取和写入超级账本的内容（但不实际修改账本），将其生成一份`RW Sets`, 背书节点对`RW Sets`进行签名并将签名后的内容返回给客户端应用。

**注意**：背书节点需要持有交易所用的链码用于交易申请的模拟工作。

## Transaction Endorsement

> A transaction endorsement is a signed response to the results of the simulated transaction. The method of transaction endorsements depends on the endorsement policy which is specified when the chaincode is deployed. An example of an endorsement policy would be "the majority of the endorsing peers must endorse the transaction". Since an endorsement policy is specified for a specific chaincode, different channels can have different endorsement policies.

交易的背书指交易申请被背书节点模拟并签名之后的响应。交易背书的方法依赖于链码部署时指定的背书策略。比如：“主要的背书节点必须对交易申请进行背书”。因此，背书策略被链码定制，进而，不同的通道也可以有不同的背书策略。

## Transaction Flow (Step 3)

> The application then submits the endorsed transaction and the RW sets to the ordering service. Ordering happens across the network, in parallel with endorsed transactions and RW sets submitted by other applications.

![img](https://pic2.zhimg.com/80/v2-b6e7b13624d1cff4152e2c223538c355_720w.jpg)

客户端将背书交易和`RW Sets`发送给排序节点。排序节点会收集整个网络里各个客户端的背书交易来进行排序。

## Transaction Flow (Step 4)

> The ordering service takes the endorsed transactions and RW sets, orders this information into a block, and delivers the block to all committing peers.

![img](https://pic2.zhimg.com/80/v2-eeb54ce57f8a6018443e22f34b3ebad9_720w.jpg)

> The **ordering service**, which is made up of a cluster of orderers, does not process transactions, smart contracts, or maintains the shared ledger. The ordering service accepts the endorsed transactions and specifies the order in which those transactions will be committed to the ledger. The Fabric v1.0 architecture has been designed such that the specific implementation of 'ordering' (Solo, Kafka, BFT) becomes a pluggable component. The default ordering service for Hyperledger Fabric is Kafka. Therefore, the ordering service is a modular component of Hyperledger Fabric.

排序节点将所有的背书交易排好序封装入区块后，将该区块发送给所有的提交节点。排序节点由一个节点集群构成，排序节点不包含处理交易，智能合约和写入账本的功能。排序节点的功能仅仅是接受各个客户端应用发送过来的背书交易，并把所有交易进行排序打包写入区块中，然后将区块发送给提交节点。`Fabric`实现了`Solo,Kafka,SBFT`等排序的机制，并且将其作为模块化的形式与区块链网络通信。

## Ordering Service

> [排序节点视频介绍](https://link.zhihu.com/?target=https%3A//courses.edx.org/courses/course-v1%3ALinuxFoundationX%2BLFS171x%2B3T2017/courseware/f0db5224eb0e4bbb8cc1e93a6819012c/5ebddaca983d4d6d952e83f95ea9e281/%3Fchild%3Dfirst%23transcript-end-b130c6deb2764111ac0fad8e37f967dc)
> The ordering service is actually something that we conceived of as a function of the initial rollout of Fabric 0.6, last year,
> in the sense that... we determined that, in order to improve the performance of the consensus computation,
> that, if we separated out the ordering aspects of consensus, where typically, whether it's Bitcoin or Ethereum, the minors are determining the order of transactions in a block,
> if we instead make that in an independent service, and apply the fault tolerance to the ordering service itself,
> we can actually get a significant improvement in performance and throughput of the overall system.
> And so, we've implemented, to date, two ordering services.
> One is called SOLO - it's really a toy; I mean, it's intended to be used for development purposes, or initial testing of new functions, and so forth.
> And then, another one is based on an implementation of Kafka.
> And, over time, as we go forward, we plan on introducing various forms of fault tolerance too to that ordering service.
> And so, the initial one is going to be based on RAFT consensus, which isn't byzantine fault tolerant, but it is crash fault tolerant,
> and then, there is ongoing work on something we call Simplified Byzantine Fault Tolerance,
> and that, we should have probably in the first half of 2018.

`Fabric`将排序节点作为一个独立模块的方式，不同于比特币和以太坊将共识机制嵌入到整个公有链系统。最大的好处就是可以针对不同的业务场景自行选择共识的算法，以达到对不同业务场景需求下性能的最优化。

## Ordering (Part I)

> *Transactions within a timeframe are sorted into a block and are committed in sequential order.*
>
> In a blockchain network, transactions have to be written to the shared ledger in a consistent order. The order of transactions has to be established to ensure that the updates to the world state are valid when they are committed to the network. Unlike the Bitcoin blockchain, where ordering occurs through the solving of a cryptographic puzzle, or *mining*, Hyperledger Fabric allows the organizations running the network to choose the ordering mechanism that best suits that network. This modularity and flexibility makes Hyperledger Fabric incredibly advantageous for enterprise applications.

排序节点，实现了区块链上所有的交易排序之后被计入分布式账本的功能，交易排序为了保证交易被写入账本时的合法性。

## Ordering (Part II)

> Hyperledger Fabric provides three ordering mechanisms: SOLO, Kafka, and Simplified Byzantine Fault Tolerance (SBFT), the latter of which has not yet been implemented in Fabric v1.0.
> \- - **SOLO** is the Hyperledger Fabric ordering mechanism most typically used by developers experimenting with Hyperledger Fabric networks. SOLO involves a single ordering node.
> \- **Kafka** is the Hyperledger Fabric ordering mechanism that is recommended for production use. This ordering mechanism utilizes Apache Kafka, an open source stream processing platform that provides a unified, high-throughput, low-latency platform for handling real-time data feeds. In this case, the data consists of endorsed transactions and RW sets. The Kafka mechanism provides a crash fault-tolerant solution to ordering.
> \- **SBFT** stands for Simplified Byzantine Fault Tolerance. This ordering mechanism is both crash fault-tolerant and byzantine fault-tolerant, meaning that it can reach agreement even in the presence of malicious or faulty nodes. The Hyperledger Fabric community has not yet implemented this mechanism, but it is on their roadmap.
> These three ordering mechanisms provide alternate methodologies for agreeing on the order of transactions.

目前`Fabric`提供了三个排序算法实现：

1. `Solo`用于开发阶段的调试。
2. `Kafka`建议商用。
3. 其他如`Raft`,`SFBT`等实现仍在开发阶段。

## Transaction Flow (Step 5)

> The committing peer validates the transaction by checking to make sure that the RW sets still match the current world state. Specifically, that the Read data that existed when the endorsers simulated the transaction is identical to the current world state. When the committing peer validates the transaction, the transaction is written to the ledger, and the world state is updated with the Write data from the RW Set.

![img](https://pic1.zhimg.com/80/v2-b05e5430900cf5e414e307d2f99de088_720w.jpg)

> If the transaction fails, that is, if the committing peer finds that the RW set does not match the current world state, the transaction ordered into a block will still be included in that block, but it will be marked as invalid, and the world state will not be updated.
> Committing peers are responsible for adding blocks of transactions to the shared ledger and updating the world state. They may hold smart contracts, but it is not a requirement.

提交节点在接收到排序节点排序完成的交易之后，通过验证`RW Sets`是否还满足当前的区块链账本的状态的方式来判断交易是否依旧合法。特别的，对于`Read Data`而言，需要判断背书节点进行模拟交易时的读取数据与当前区块链账本的数据是否完全一致。完成交易验证之后，提交节点将按照`Write Data`的内容对区块链账本进行更新。

如果验证交易失败，该交易仍旧会在区块中，但会被标记为非法状态并且整个的账本状态不会更新。

提交节点用于接收排序节点的输出，并将区块写入区块链账本，对校验失败的交易标志非法。提交节点可以持有智能合约但非必须。

## Transaction Flow (Step 6)

> Lastly, the committing peers asynchronously notify the client application of the success or failure of the transaction. Applications will be notified by each committing peer.

![img](https://pic1.zhimg.com/80/v2-ba380b73a55eff97c85da3abdc1d86e8_720w.jpg)

最后，各个提交节点异步的通知客户端应用，其提交的交易是否成功。

## Identity Verification

> In addition to the multitude of endorsement, validity, and versioning checks that take place, there are also ongoing identity verifications happening during each step of the transaction flow. Access control lists are implemented on the hierarchical layers of the network (from the ordering service down to channels), and payloads are repeatedly signed, verified, and authenticated as a transaction proposal passes through the different architectural components.

在整个交易流程中，除了上述操作，同时发生的还有身份鉴权及权限控制。访问控制列表是在网络结构的各层上实现的(从排序节点到通道)，其相关的开销（签名，验证，授权）同交易申请一起在交易流程中出现。

## Transaction Flow Summary

> It is important to note that the state of the network is maintained by peers, and not by the ordering service or the client. Normally, you will design your system such that different machines in the network play different roles. That is, machines that are part of the ordering service should not be set up to also endorse or commit transactions, and vice versa. However, there is an overlap between endorsing and committing peers on the system. Endorsing peers must have access to and hold smart contracts, in addition to fulfilling the role of a committing peer. Endorsing peers do commit blocks, but committing peers do not endorse transactions.
> Endorsing peers verify the client signature, and execute a chaincode function to simulate the transaction. The output is the chaincode results, a set of key/value versions that were read in the chaincode (Read set), and the set of keys/values that were written by the chaincode. The proposal response gets sent back to the client, along with an endorsement signature. These proposal responses are sent to the orderer to be ordered. The orderer then orders the transactions into a block, which it forwards to the endorsing and committing peers. The RW sets are used to verify that the transactions are still valid before the content of the ledger and world state is updated. Finally, the peers asynchronously notify the client application of the success or failure of the transaction.

整个的区块链网络是由提交节点提供支撑的。一般情况下，设计的区块链系统中，不同功能的节点会出现在不同的宿主机上。但提交节点和背书节点的功能会有重合，背书节点必须持有链码，但提交节点不一定。背书节点可以背书交易，但提交节点不可以。

背书节点需要验证客户端的签名，利用持有的链码逻辑来进行交易模拟。模拟的结果就是链码的执行结果。包括一个读的 `K/V` 列表以及一个写的 `K/V` 列表。背书节点将上述`RW Sets`以及背书节点自身的签名一起返回给客户端应用。客户端应用再将背书交易发送至排序节点。排序节点介绍全网的背书交易，并将其排序放入区块中，之后将区块发送给各个提交节点。最后，各个提交节点利用`RW Sets`验证交易，并执行写入区块链账本的操作。最后提交节点异步的通知客户端应用其发送的交易是否合法。

## Channels

> Channels allow organizations to utilize the same network, while maintaining separation between multiple blockchains. Only the members of the channel on which the transaction was performed can see the specifics of the transaction. In other words, channels partition the network in order to allow transaction visibility for stakeholders only. This mechanism works by delegating transactions to different ledgers. Only the members of the channel are involved in consensus, while other members of the network do not see the transactions on the channel.

![img](https://pic4.zhimg.com/80/v2-b23a6aaaa627620a0ab161c556ff87b3_720w.jpg)

> The diagram above shows three distinct channels -- blue, orange, and grey. Each channel has its own application, ledger, and peers.
> Peers can belong to multiple networks or channels. Peers that do participate in multiple channels simulate and commit transactions to different ledgers. The ordering service is the same across any network or channel.
> A few things to remember:
> \- - The network setup allows for the creation of channels.
> \- The same chaincode logic can be applied to multiple channels.
> \- A given user can participate in multiple channels.

通道技术允许不同的组织公用同一套区块链网络，却可以实现数据层面的相对独立。只有交易所在的通道的参与方才能够获取有关该通道的信息。通道的这种机制通过委托交易至不同的账本实现。只有通道的参与方才需要针对该交易达成共识，而非该通道的参与方则无法看见该交易。

上面的示意图所示，蓝色橘色和灰色代表不同的通道，不同的通道包含不同的客户端应用，账本和节点。

节点可以同时属于多个通道。节点对来自不同通道的业务执行模拟交易，背书，提交交易到不同账本的功能。但排序节点是所有通道公用的。

**注意**：

1. 区块链网络通信允许创建通道。
2. 同样的链码可以部署给不同的通道。
3. 一个确定的用户可以参与到多个不同的通道中。

## State Database

> The current state data represents the latest values for all assets in the ledger. Since the current state represents all the committed transactions on the channel, it is sometimes referred to as world state.
> Chaincode invocations execute transactions against the current state data. To make these chaincode interactions extremely efficient, the latest key/value pairs for each asset are stored in a state database. The state database is simply an indexed view into the chain’s committed transactions. It can therefore be regenerated from the chain at any time. The state database will automatically get recovered (or generated, if needed) upon peer startup, before new transactions are accepted. The default state database, **LevelDB**, can be replaced with **CouchDB**.
> \- - LevelDB is the default key/value state database for Hyperledger Fabric, and simply stores key/value pairs.
> \- CouchDB is an alternative to LevelDB. Unlike LevelDB, CouchDB stores JSON objects. CouchDB is unique in that it supports keyed, composite, key range, and full data-rich queries.
> Hyperledger Fabric’s LevelDB and CouchDB are very similar in their structure and function. Both LevelDB and CouchDB support core chaincode operations, such as getting and setting key assets, and querying based on these keys. With both, keys can be queried by range, and composite keys can be modeled to enable equivalence queries against multiple parameters. But, as a JSON document store, CouchDB additionally enables rich query against the chaincode data, when chaincode values (e.g. assets) are modeled as JSON data.

![img](https://pic1.zhimg.com/80/v2-9fa87a2726077cff05169f85584224ac_720w.jpg)

区块链当前的数据状态代表了当前时刻分布式账本上记录的所有资产的状态信息。因此当前的数据状态代表了在通道上所有已经被提交的交易，有时会把这个状态称为世界观。

链码对当前状态数据通过调用执行交易的方式进行通信。为了使这些链码的交互非常高效，每个资产的最新密钥/值对存储在一个状态数据库中。状态数据库只是对区块链的提交事务的索引视图。因此，它可以在任何时候从链中重新生成，在新事务被接受之前，状态数据库将自动恢复(或生成，如果需要的话)。默认的状态数据库是 `LevelDB`，可以用`CouchDB`来替换。

`LevelDB` 是一个简单的存储`K/V`的数据库

`CouchDB`则存储`json`对象

两个数据库都支持通过核心链码 `API`来操作。

## Smart Contracts

> As a reminder, smart contracts are computer programs that contain logic to execute transactions and modify the state of the assets stored within the ledger. Hyperledger Fabric smart contracts are called **chaincode** and are written in Go. The chaincode serves as the business logic for a Hyperledger Fabric network, in that the chaincode directs how you manipulate assets within the network. We will discuss more about chaincode in the *Understanding Chaincode* section.

智能合约是包含执行交易，修改分布式账本中资产状态的代码逻辑, 在超级账本中被称为链码。链码用`GO`语言编写。在`Fabric`中，链码用于实现具体的业务逻辑，包括如何在区块链网络中处置你的资产。

## Membership Service Provider (MSP)

> The membership service provider, or MSP, is a component that defines the rules in which identities are validated, authenticated, and allowed access to a network. The MSP manages user IDs and authenticates clients who want to join the network. This includes providing credentials for these clients to propose transactions. The MSP makes use of a *Certificate Authority*, which is a pluggable interface that verifies and revokes user certificates upon confirmed identity. The default interface used for the MSP is the **Fabric-CA API**, however, organizations can implement an External Certificate Authority of their choice. This is another feature of Hyperledger Fabric that is modular. Hyperledger Fabric supports many credential architectures, which allows for many types of External Certificate Authority interfaces to be used. As a result, a single Hyperledger Fabric network can be controlled by multiple MSPs, where each organization brings their favorite.

`MSP`用于超级账本的成员管理，其管理着联盟链的准入权，以及所有已经加入联盟链的所有用户的`id`。 默认基于`CA`为加入的用户颁发/取消授权，默认的接口为`Fabric-CA API`。但开发者可以根据自己的需求选则其他方案作为`CA`模块。

## What Does the MSP Do?

> To start, users are authenticated using a certificate authority. The certificate authority identifies the application, peer, endorser, and orderer identities, and verifies these credentials. A signature is generated through the use of a *Signing Algorithm* and a *Signature Verification Algorithm*.
> Specifically, generating a signature starts with a *Signing Algorithm*, which utilizes the credentials of the entities associated with their respective identities, and outputs an endorsement. A signature is generated, which is a byte array that is bound to a specific identity. Next, the *Signature Verification Algorithm* takes the identity, endorsement, and signature as inputs, and outputs 'accept' if the signature byte array corresponds with a valid signature for the inputted endorsement, or outputs 'reject' if not. If the output is 'accept', the user can see the transactions in the network and perform transactions with other actors in the network. If the output is 'reject', the user has not been properly authenticated, and is not able to submit transactions to the network, or view any previous transactions.

![img](https://pic4.zhimg.com/80/v2-2fe3f7dc2fa52699a96ef7948432113b_720w.jpg)

`CA`鉴权用于节点，背书，排序和验证的各个环节。主要包含“签名算法“和”验签算法“两个。

具体地说，生成签名首先使用一个签名算法，该算法利用签名者的数字身份对一段信息进行签名，并输出一个背书。生成签名，表面是绑定到特定标识的字节数组。接下来，签名验证算法将身份、背书和签名作为输入，如果签名字节数组对应于输入的背书的有效签名则输出“通过”，反之输出“拒绝”，。如果输出是“通过”，则用户可以看到网络中的交易，并与网络中的其他参与者执行交易。如果输出是“拒绝”，则用户没有经过正确的身份验证，不能向网络提交交易，也不能查看任何以前的交易。

# FastFabric:提升Hyperledger Fabric性能到20000TPS

## 摘要

预计区块链技术将对各种行业产生重大影响。然而，阻碍它们的一个问题是它们有限的交易吞吐量，特别是与诸如分布式数据库系统之类的已建立的解决方在本文中，我们重新构建了一个现代许可的区块链系统Hyperledger Fabric，以将交易吞吐量从每秒3,000次增加到20,000次。我们专注于超出共识机制的性能瓶颈，我们提出架构更改，以减少交易排序和验证期间的计算和I / O开销，从而大大提高吞吐量。值得注意的是，我们的优化是完全即插即用的，不需要对Hyperledger Fabric进行任何界面更改。

## 1. 介绍

区块链等分布式账本技术提供了一种以安全和可验证的方式进行交易的方法，而无需可信任的第三方。因此，人们普遍认为区块链将对金融、房地产、公共管理、能源和交通行业产生重大影响[1]。但是，为了在实践中可行，区块链必须支持与现有数据库管理系统支持的交易率相当的交易率，这可以提供一些相同的交易保证。

与不限制网络成员资格的无权限区块链相比，我们专注于许可的区块链，其中所有参与节点的身份都是已知的。许可的区块链适用于许多应用领域，包括融资;例如，Ripple区块链旨在提供类似于当前SWIFT系统的货币兑换和跨银行交易的支付网络。

从技术角度来看，我们观察到这两种区块链之间的重要区别。无权区块链的无信任性质要求工作或利益证明，或昂贵的全球规模的拜占庭 - 容错共识机制[2]。最近的许多工作都集中在使这些更有效。另一方面，许可的区块链通常将共识和交易验证委托给选定的节点组，从而减轻了共识算法的负担。虽然在这种情况下，共识仍然是一个瓶颈，但最近正在解决这个问题https：//[http://ripple.com](https://link.zhihu.com/?target=http%3A//ripple.com)工作[3]，[4]，激励我们超越共识来确定进一步的绩效改进。

在本文中，我们批判性地研究了Hyperledger Fabric 1.2的设计，因为据报道它是最快的开源许可区块链[5]。虽然已经有一些关于优化Hyperledger Fabric的工作，例如，使用积极的缓存[6]，但我们并不知道任何先前关于重新构建系统的工作2。因此，我们基于通用系统设计技术设计并实现了几种架构优化，这些技术将端到端交易吞吐量提高了近7倍，从每秒3,000到20,000个交易处理，同时减少了块延迟。我们的具体贡献如下：

1）从数据中分离元数据：Fabric中的共识层接收整个交易作为输入，但只需要交易ID来决定交易顺序。我们重新设计Fabric的交易排序服务，仅使用交易ID，从而大大提高了吞吐量。

2）并行和缓存：交易验证的某些方面可以并行化，而其他方面可以从缓存数据中受益。我们重新设计了Fabric的验证服务，通过积极缓存调度器中的未编组块并通过并行化尽可能多的验证步骤，包括认可策略验证和语法验证。

3）利用存储器层次结构在关键路径上进行快速数据访问：Fabric维护世界状态的键值存储可以用轻量级内存数据结构代替，其缺乏耐久性保证可以通过区块链本身进行补偿。我们围绕轻量级哈希表重新设计Fabric的数据管理层，该表可以更快地访问关键交易验证路径上的数据，从而将不可变块的存储推迟到写优化存储集群。 4）资源分离：提交者和背书者的对等角色争夺资源。我们介绍了一种将这些角色转移到单独硬件的架构。

重要的是，我们的优化不会违反Fabric的任何API或模块化边界，因此它们可以合并到Fabric版本2.0的计划版本中[7]。

2有关相关工作的详细调查，请参阅第V节。

我们还概述了未来工作的几个方向，与我们提出的优化一起，有可能达到Visa等信用卡公司所要求的每秒50,000笔交易[2]。

在本文的其余部分，第二部分简要概述了Hyperledger Fabric，第三部分介绍了我们的改进设计，第四部分讨论了实验结果，第五部分介绍了我们在先前工作中的贡献，第六部分总结了未来的发展方向工作。

## 2. Fabric架构

作为由Linux Foundation托管的开源Hyperledger项目的一个项目，Fabric是最活跃开发的许可区块链系统之一[8]。由于Androulaki等[9]详细描述了交易流程，因此我们仅提供一个简短的概要，重点介绍我们在第III节中提出改进的系统部分。

为了避免智能合约确定性的陷阱并允许即插即用的系统组件替换，Fabric的结构与其他常见的区块链系统不同。交易遵循执行顺序提交流模式而不是公共顺序执行提交模式。客户端交易首先在沙箱中执行以确定它们的读写集，即由交易读取和写入的健值对集。然后，交易由排序服务排序，并最终验证并提交给区块链。此工作流程由分配了特定角色的节点实现，如下所述。

A. 节点类型

客户端发起交易，即对区块链的读取和写入，发送到Fabric节点。节点是Peer或排序者Orderer;一些Peer也是背书者。所有Peer都将块提交到区块链的本地副本，并将相应的更改应用于维护当前世界状态快照的状态数据库。允许背书者节点根据链码，Fabric的智能合约版本中捕获的业务规则来证明交易有效。排序者仅负责决定交易顺序，而不是正确性或有效性。

B. 交易流程

客户将其交易发送给一些背书者。每个背书者在沙箱中执行交易，并计算相应的读写集以及访问的每个键的版本号。每个背书者还使用业务规则来验证交易的正确性。客户等待足够数量的认可，然后将这些响应发送给排序服务的Orderer。排序者首先就进入的交易的顺序达成共识，然后将消息队列分成块。块被传递给Peer节点，然后Peer验证并提交它们。由于其许可性质，所有节点必须已知并向成员服务提供商（MSP）注册，否则其他节点将忽略它们。

C. 实现细节

为了讨论第III节中的改进，我们现在仔细研究一下orderer和peer架构。

1）排序者：在收到来自背书者的回复后，客户端创建包含标题和有效负载的交易提议。标题包括交易ID和通道ID。有效负载包括读写集和相应的版本号，以及支持的Peer节点的签名。交易提议使用客户的凭证签署并发送到排序服务。

排序服务的两个目标是（a）在交易顺序上达成共识，以及（b）将包含有序交易的块交付给提交者Peer节点。Fabric目前使用Apache Kafka（基于ZooKeeper [10]）来实现容错崩溃的共识。

排序者收到交易提案时，它会检查客户是否有权提交交易。如果是，则orderer将交易提议发布到Kafka集群，其中每个Fabric通道都映射到Kafka主题以创建相应的不可变的交易序列顺序。然后，每个排序者根据每个块允许的最大交易数或块超时周期将从Kafka接收的交易组装成块。使用orderer的凭证对块进行签名，并使用gRPC [11]将其传递给Peer节点。

2）Peer：从排序服务接收消息时，Peer首先从区块的头部和元数据检查其语法结构。然后检查orderer的签名是否符合指定的策略。任何这些测试失败的区块都会被立即丢弃。

在初始校验之后，区块被推入队列，保证其添加到区块链。但是，在此之前，区块会依次执行两个验证步骤和最后一个提交步骤。

在第一个验证步骤中，将解压缩块中的所有交易，检查其语法并验证其认可。未通过此测试的交易将被视为无效，但会保留在块中。此时，只有善意创建的交易仍然有效。

在第二个验证步骤中，Peer确保有效交易之间的相互作用不会导致无效的世界状态。回想一下，每个交易都带有一组需要从世界状态数据库（它的读取集）读取的键，以及它将写入数据库（它的写集）的一组键和值，以及它们记录的版本号通过背书者。在第二个验证步骤中，交易的读写集中的每个键仍必须具有相同的版本号。从任何先前交易中写入该Key会更新版本号并使交易无效。这可以防止双重支出。

在最后一步中，Peer将块（现在包括其交易的验证标志）写入文件系统。

Fabric被虚拟化为多个通道，由通道ID识别。

Key及其值（即世界状态）将保留在LevelDB或CouchDB中，具体取决于应用程序的配置。此外，每个块及其交易的索引都存储在LevelDB中以加速数据访问。

## 3. 设计

本节介绍我们对Fabric 1.2版的体系结构和实现的更改。该版本于2018年7月发布，随后于2018年9月发布1.3版，2019年1月发布1.4版。但是，最近发布的版本中引入的更改不会影响我们的提议，因此我们预计在将我们的工作与新版本集成时不会有任何困难。重要的是，我们的改进使各个模块的接口和职责保持不变，这意味着我们的更改与现有的Peer或排序服务实现兼容。此外，我们的改进是相互正交的，因此可以单独实施。对于Orderer和Peer，我们按照从最小到最大的性能影响的升序来描述我们的提议，与其对Fabric的各自更改相比。

A. 准备

使用拜占庭容错（BFT）一致性算法是HyperLedger中的关键性能瓶颈[2]。这是因为BFT一致性算法不能很好地与参与者的数量成比例。在我们的工作中，我们选择超越这个明显的瓶颈有三个原因：

•可以说，在许可的区块链中使用BFT协议并不像在无权限的系统中那么重要，因为所有参与者都是已知的并且被激励以保持系统以诚实的方式运行。

•正在广泛研究BFT共识[12]，我们预计未来一两年内将出现更高吞吐量的解决方案。

•实际上，Fabric 1.2不使用BFT共识协议，而是依赖于Kafka进行交易排序，如前所述。

出于这些原因，我们的工作目标不是使用更好的BFT一致性算法来提高orderer性能，而是为了缓解当共识不再是瓶颈时出现的新问题。我们首先对排序服务提出两项改进，然后对Peer进行一系列改进。

B. Orderer改进I：从有效负载中分离交易头

在Fabric 1.2中，使用Apache Kafka的订购者将整个交易发送给Kafka进行订购。交易的长度可能是几千字节，导致高通信开销，从而影响整体性能。但是，就交易订单达成共识只需要交易ID，因此我们可以通过仅向Kafka集群发送交易ID来获得订货人吞吐量的显着改善。

具体而言，在从客户端接收交易时，我们的订货人从标题中提取交易ID并发布图1.新的订货人架构。传入的交易同时处理。他们的TransactionID被发送到Kafka集群进行订购。当接收到有序的TransactionID时，订货人用它们的有效负载重新组装它们并将它们收集到块中。



![img](https://pic3.zhimg.com/80/v2-78cd4a4fcf4c2008324b7bdb06e62c56_720w.jpg)



这个ID到Kafka集群。排序者将相应的有效负载分别存储在本地数据结构中，并在从Kafka收回ID时重新组装交易。随后，与Fabric一样，orderer将交易集分段为块并将它们传递给Peer。值得注意的是，我们的方法适用于任何共识实现，并且不需要对现有排序界面进行任何修改，从而允许我们利用现有的Fabric客户端和Peer代码。

C. Orderer改进II：消息流水线

在Fabric 1.2中，订购服务逐个处理来自任何给定客户端的传入交易。当交易到达时，识别其相应的信道，根据一组规则检查其有效性，并最终将其转发到共识系统，例如，卡夫卡;只有这样才能处理下一个交易。相反，我们实现了一个可以同时处理多个传入交易的流水线机制，即使它们来自使用相同gRPC连接的同一客户端。为此，我们维护一个线程池，它并行处理传入的请求，每个传入的请求都有一个线程。线程调用Kafka API来发布交易ID，并在成功时向客户端发送响应。订货人完成的剩余处理与Fabric 1.2相同。

图1总结了新的排序设计，包括交易ID与有效负载的分离以及由于并行消息处理导致的横向扩展。

D. Peer任务

回忆一下第II-C2节，在从排序者接收块时，Fabric Peer按顺序执行以下任务：

•验证收到消息的合法性•验证块中每个交易的块头和每个认可签名•验证交易的读写集•更新LevelDB或CouchDB中的世界状态•将区块链日志存储在文件系统中，与LevelDB中的相应索引

我们的目标是在交易流程的关键路径上最大化交易吞吐量。为此，我们进行了广泛的调用图分析，以确定性能瓶颈。



![img](https://pic3.zhimg.com/80/v2-f9ec86e83e82a39cc486f831d4ab3ac6_720w.jpg)



图2.新的Peer结构。快速Peer使用内存中的哈希表来存储世界状态。验证通道完全并发，并行验证多个块及其交易。代言人角色和持久存储被分成可扩展的集群，并由快速Peer给出经过验证的块。通道的所有部分都使用缓存中的未编组块。

我们做出以下观察。首先，验证交易的读写集需要快速访问world状态。因此，我们可以通过使用内存中的哈希表而不是数据库来加速该过程（第III-E节）。其次，交易流程不需要区块链日志，因此我们可以在交易流程结束时将其存储到专用存储和数据分析服务器（第III-F节）。第三，如果Peer也是背书者，则需要处理新的交易提案。但是，提交者和背书者角色是不同的，这使得为每项任务专用不同的物理硬件成为可能（第III-G节）。第四，必须在Peer验证和解决传入的块和交易。最重要的是，必须按顺序完成通过交易写入集验证状态更改，阻止所有其他任务。因此，尽可能加快这项任务非常重要（第III-H节）。

最后，通过缓存协议缓冲区[13]解块的结果（第III-I节），可以获得显着的性能提升。我们在图2中详细介绍了这种架构重新设计，包括其他提出的Peer改进。

E. Peer改进I：用哈希表替换世界状态数据库

必须为每个交易按顺序查找和更新世界状态数据库，以保证所有Peer的一致性。因此，对此数据存储的更新以尽可能高的交易速率发生是至关重要的。

我们认为，对于常见情况，例如跟踪分类帐上的钱包或资产，世界状态可能相对较小。即使需要存储数十亿个密钥，大多数服务器也可以轻松地将它们保存在内存中。因此，我们建议使用内存中的哈希表而不是LevelDB / CouchDB来存储世界状态。这样可以在更新世界状态时消除硬盘访问。它还消除了由于区块链本身的冗余保证而不必要的昂贵的数据库系统保证（即，ACID属性），进一步提高了性能。当然，由于使用易失性存储器，这种替换易受节点故障的影响，因此必须通过稳定存储来增强内存中的哈希表。我们在第III-F节讨论了这个问题。

F.Peer改进II：使用对等集群存储块

通过定义，块是不可变的。这使它们非常适合仅附加数据存储。通过将数据存储与对等任务的其余部分分离，我们可以设想用于块和世界状态备份的多种类型的数据存储，包括在其文件系统中存储块和世界状态备份的单个服务器，如Fabric目前所做的那样;数据库或键值存储，如LevelDB或CouchDB。为了实现最大扩展，我们建议使用分布式存储集群。请注意，使用此解决方案，每个存储服务器仅包含链的一小部分，从而激励使用分布式数据处理工具，如Hadoop MapReduce或Spark5。

G.Peer改进III：单独承诺和认可

在Fabric 1.2中，endorser peer也负责提交块。认可是一项昂贵的操作，承诺也是如此。虽然对一组背书者进行并发交易处理可能会提高应用程序性能，但在每个新节点上复制承诺的额外工作实际上无效了这一优势。因此，我们建议将这些角色分开。

具体而言，在我们的设计中，提交者节点执行验证通道，然后将经过验证的块发送给背书者群集，这些背书者仅将更改应用于其世界状态而无需进一步验证。此步骤允许我们释放Peer的资源。请注意，这种可以扩展以满足需求的背书者群集只会将对等方的认可角色分割为专用硬件。此群集中的服务器不等同于Fabric 1.2中的完整版本的背书节点。

H. Peer改进IV：并行化验证

块和交易头验证（包括检查发件人的权限，执行认可策略和语法验证）都是高度可并行化的。我们通过引入完整的验证通道来扩展Fabric 1.2的并发性。

具体而言，对于每个传入的块，一个go-routine被分配用于通过块验证阶段。随后，这些例程中的每一个都使用Fabric 1.2中已存在的goroutine池进行交易验证。因此，在任何给定时间，并行检查多个块及其交易的有效性。最后，所有读写集都由单个goroutine以正确的顺序依次验证。这使我们能够充分利用多核服务器CPU的潜力。但是，我们当前的实现不包括这样的存储系统。

I. Peer增强 V: 缓存解析的区块

Fabric使用gRPC在网络中让节点进行通信。Protocol Buffers被用来进行序列化。为了保证处理应用和软件升级，Fabric的区块结构是高度层级化的，每一个层级都是单独序列化和反序列化的。这将导致大量的内存将用来进行byte数组转成结构化的数据。而且，Fabric 1.2没有在缓存中存储之前解析的数据，因此当需要这些数据的时候，这些工作将被重复执行。

为了缓解这个问题，我们计划用一个临时的缓存来存放解析的数据。区块会别缓存，无论当在校验管道还是通过区块号接收的时候。一旦区块的任何部分被解析了，它将会被存储到区块中以便再利用。我们用循环的缓存池实现这一功能，这是一个足够大的校验管道。无论区块是否被提交，一个新的区块可以被自动放进管道覆盖已经存在的位置的区块。由于在提交后不需要缓存，并且保证新块只在旧块离开管道后到达，所以这是一个安全的操作。注意解析操作只会在缓存中追加数据，不会修改。所以在校验通道里可以进行多线程的无锁操作。在最差的场景中，许多线程去读取同一个未被解析的数据，所有的程序并发的执行解析操作。接着最后写入缓存的线程获得胜利，这是没有问题的，因为大家执行的结果都是一致的。

调用图标分析，即使进行了这些操作，由于解析操作，内存占用率在执行期间仍然非常高。这是继gRPC调用和加密计算之后的占据最大的一项操作。但是后面2个操作，不在本次工作范围内。

## 4. 结果

本节介绍了对我们的体系结构改进的实验性性能评估。我们使用了15台本地服务器，通过1 Gbit/s交换机连接。每台服务器配备两个2.10GHz的Intel R Xeon R CPU E5-2620 v2处理器，总共24个硬件线程和64 GB RAM。我们使用Fabric1.2作为基本情况，并逐步添加我们的改进以进行比较。默认情况下，fabric配置为使用leveldb作为对等状态数据库，排序服务将已完成的块存储在内存中，而不是磁盘上。此外，我们不使用Docker容器来运行整个系统，以避免额外的开销。

虽然我们确保我们的实现不会更改fabric的验证行为，但所有测试都是使用不冲突且有效的交易运行的。这是因为有效交易必须经过每个验证检查，并且它们的写集将在提交期间应用于状态数据库。相反，可以删除无效的交易。因此，我们的结果评估了最坏情况下的性能。

![img](https://pic3.zhimg.com/80/v2-f9835f8b6ceb2ce78e2371d06746219a_720w.jpg)



对于专门针对排序者或提交者实验，我们分离了各自的系统部分。在order实验中，我们从客户机向order发送预加载的背书交易，并让一个模拟提交者简单地丢弃创建的块。类似地，在提交者的基准测试期间，我们将预加载的块发送给提交者，并为背书者和丢弃已验证块的块存储创建mock。

然后，对于端到端的设置，我们实现了完整的系统：背书者根据来自提交人的已验证块的复制世界状态从客户端背书交易建议；订购人从背书交易创建块并将它们发送给提交人；提交者验证并提交对其内存中世界状态的更改，并将已验证的块发送给背书者和块存储；块存储使用结构1.2数据管理将块存储在其文件系统中，状态存储在leveldb中。但是，我们没有为可伸缩分析实现分布式块存储；这超出了本工作的范围。

为了进行公平的比较，我们对所有实验使用了相同的交易链码：每个交易模拟从一个帐户到另一个帐户的资金转移，读取并更改状态数据库中的两个键。这些交易的有效负载为2.9kb，这是典型的[3]。此外，我们使用默认的背书政策，即接受单一背书人签名。

A. 通过GRPC传输数据块

我们首先对GRPC的性能进行基准测试。我们预先创建了包含不同交易数的有效块，通过Fabric GRPC接口将它们从Orderer发送到Peer，然后立即丢弃它们。实验结果如图3所示。

我们发现，对于从10到250个交易的块大小（这是在以下部分中导致最佳性能的大小），交易吞吐量超过40000个交易/秒是可持续的。与第IV-D节中的端到端测试结果相比，很明显，在我们的环境中，网络带宽和服务器机架中使用的1 Gbit/s交换机不是性能瓶颈。



![img](https://pic1.zhimg.com/80/v2-cb4807bc5ad781955fb127a4c7c9f1ac_720w.jpg)



B. 作为消息大小函数的订购方吞吐量

在这个实验中，我们设置了多个客户机来向订购者发送事务，并监视发送100000个事务所需的时间。我们评估订单在订单1.2中的交易率，并将其与我们的改进进行比较：

•opt o-i：仅向卡夫卡发布事务ID（第三-B节）

•选择O-II：来自客户的并行传入交易建议（第III-C节）

图4显示了不同负载大小的交易吞吐量。在Fabric1.2中，由于向Kafka发送大消息的开销，交易吞吐量随着负载大小的增加而降低。然而，当我们只将Turac ID发送给卡夫卡（OPT-O-1）时，对于4096 kb的有效负载大小，我们几乎可以将平均吞吐量（2.8×）增加三倍。添加优化o-2后，平均吞吐量比基础结构1.2提高4倍。特别是，对于2kb负载大小，我们将排序服务性能从6215个交易/秒提高到21719个交易/秒，比率接近3.5x。

C. Peer实验

在这一节中，我们描述了在孤立的单个Peer节点上的测试（我们在第IV-D节中展示了端到端评估的结果）。在这里，我们预先计算块并将它们发送给Peer，就像我们在第IV-A节的GRPC实验中所做的那样。然后Peer完全验证并提交块。

与结构1.2相比，图中所示的三种配置累计包含了我们的改进（即，opt p-i i包含opt p-i，opt p-iii包含了先前的两种改进）：

•opt p-i leveldb替换为内存哈希表 •opt p-ii验证和提交完全并行化；块存储和背书分离，通过远程GRPC传输，分离存储服务器 •opt p-iii所有解析的数据可以在校验和提交管道访问



![img](https://pic2.zhimg.com/80/v2-7a47c0e70821142f64745ab7d3808ff9_720w.jpg)





![img](https://pic1.zhimg.com/80/v2-dda4fd92387bba46b2c965b8ace1c760_720w.jpg)



1）固定块大小的实验：图5和图6显示了一次运行100000个交易的验证和提交结果，重复1000次。交易是收集到100个交易块中。我们先讨论延迟，然后是吞吐量。由于批处理，我们显示每个块的延迟，而不是每个交易延迟。结果与我们自己设定的目标一致，即不因吞吐量增加而引入额外的延迟；事实上，我们的性能改进将对等延迟减少到原始值的三分之一（请注意，这些实验没有考虑网络延迟）。虽然opt p-ii中引入的管道生成了一些额外的延迟，但是其他的优化并不能补偿它。

通过使用状态存储哈希表（opt p-i），我们能够将Fabric1.2对等机的吞吐量从3200个交易/秒增加到7500多个交易/秒。并行化Val-idation（optp-ii）每秒增加了大约2000个交易操作。这是因为，如图2所示，只有前两个验证步骤可以并行化和缩放。因此，整个管道性能取决于读写集验证和提交的吞吐量。虽然使用opt p-i时承诺几乎是免费的，但是直到在opt p-iii中引入解组缓存，optp-ii才有了回报。高速缓存大大减少了cpu的工作量，释放了并行验证额外块的资源。将所有Peer优化结合在一起，我们将Peer的提交性能提高了7倍，从大约3200个交易/秒提高到超过21000个交易/秒。

6.我们实验确定在该块大小上Peer吞吐量最大化。



![img](https://pic1.zhimg.com/80/v2-e416fdc25461160355cd35c12c59ab58_720w.jpg)





![img](https://pic4.zhimg.com/80/v2-6c1a1eda989cbbe3a7405148f0e2959b_720w.jpg)



2）参数敏感性：如第IV-C节所述，在Peer并行化块和交易验证至关重要。但是，不清楚要使性能最大化需要多少并行性。因此，我们探索一个Peer的性能可以通过改变两个参数来调谐的程度：

•验证管道中同时引导块的go例程的数量

•同时验证交易处理的go例程的数量

我们使用信号量控制系统中活动go协程的数量，同时允许多个块同时进入验证管道。这允许我们通过两个独立的go例程池来控制块头验证和交易验证中的并行级别。

对于100个交易的块大小，图7显示了改变go例程数量时的吞吐量。验证管道中的线程总数由两个独立轴的总和给出。例如，我们为管道中的25个交易验证go例程和31个并发块实现了最大吞吐量，总共为管道提供了56个go协程。当有太多线程时，我们会看到通过线程管理开销导致的性能小幅度下降，但是用太少的并行执行来耗尽cpu的代价是巨大的。因此，我们建议默认情况下，在给定的机器中，go例程的数量至少是物理线程的两倍。

我们现在研究Peer吞吐量对块大小的依赖性。每个块大小实验是用先前测试的最佳调谐GO例程参数进行的。在24±2交易验证中使用的所有配置执行常规程序，并在管道中进行30±3个阻塞。同样，我们在给定大小的块中为一个基准测试运行分割100000个交易，并重复实验1000次。我们选择在对数尺度上扫描块大小空间以获得宽光谱的概述。

结果如图8所示。我们发现，块大小为100个交易/块时，以每秒21000多个交易的速度提供最佳吞吐量。我们还研究了与这个块大小的小偏差。我们发现块大小在50到500之间的性能差异非常小，因此我们选择将块大小固定为100个交易。

D.端到端吞吐量

我们现在讨论通过组合所有优化（即opt）实现的端到端吞吐量。o-ii与opt相结合。p-iii，与我们对未改性织物1.2的测量结果相比。

我们设置了一个使用三个zookeeper服务器和三个kafka服务器（默认主题复制因子为三）的集群的单个排序节点，并将其连接到Peer。来自此Peer的块被发送到单个数据存储服务器，该服务器将世界状态存储在leveldb中，并将块存储在文件系统中。对于扩展，五个背书者复制对等状态并提供足够的吞吐量来处理客户端背书负载。最后，客户机安装在自己的服务器上；该客户机从五个背书服务器请求背书，并将背书事务发送到排序服务。这总共使用15台服务器连接到本地数据中心的同一个1 Gbit/s交换机。我们从客户端向排序方发送总计100000个已背书的交易，排序者将这些交易批处理为100个大小的块，并将它们传递给Peer。为了估计吞吐量，我们测量Peer上提交的块之间的时间，并取一次运行的平均值。这些运行重复100次。表1显示，与我们的基线结构1.2基准相比，显著提高了6-7倍。

## 5. 相关工作

hyperledger fabric是一个最近才开发的系统，它的架构仍在快速发展和重大变化中。因此，对于系统的性能分析或架构改进的建议方面的工作相对较少。在这里，我们综述了提高Fabric性能的最新技术。

最接近我们的工作是由thakkar等人[6]谁研究了各种配置参数对Fabric性能的影响。他们发现，主要的瓶颈是在背书策略验证期间重复验证x.509证书，对块中的交易进行顺序策略验证，以及在提交阶段进行状态验证。它们引入了对已验证的认可证书的积极缓存（合并到Fabric版本1.1中，因此是我们评估的一部分）、认可策略的并行验证以及批处理状态验证和承诺。这些改进使总吞吐量增加了16倍。我们还对提交方的验证进行了并行化，并进一步将状态数据库替换为更有效的数据结构，即哈希表。 hyperledger fabric是一个最近才开发的系统，它的架构仍在快速发展和重大变化中。因此，对于系统的性能分析或架构改进的建议方面的工作相对较少。在这里，我们综述了提高Fabric性能的最新技术。

在最近的工作中，sharma等人[14]研究了使用数据库技术，即事务重新排序和提前中止，来提高fabric的性能。他们关于早期识别冲突交易的一些想法与我们的想法是正交的，可以纳入我们的解决方案。但是，有些想法，例如让排序服务删除冲突交易，与我们的解决方案不兼容。首先，我们故意不向排序者发送交易读写集，只发送交易id。其次，我们选择遵守fabric的设计目标，即将不同的任务分配给不同类型的节点，因此我们的排序服务不检查读写集的内容。未来工作的一个有趣方向是在不同的交易工作负载下比较这两种方法，以了解何时向排序服务发送完整交易详细信息的开销值得提前修剪冲突交易。

众所周知，由于拜占庭容错协议（bft）的消息通信开销，fabric的order组件可能成为性能瓶颈。因此，在排序服务中使用bft协议的有效实现非常重要。sousa等人[3]研究了著名的bft-smart[15]实现作为fabric的一部分的使用，并表明，在单个数据中心内使用此实现，可以实现高达30000个交易/秒的吞吐量。然而，与我们的工作不同，Committer组件没有基准测试，端到端的性能也没有得到解决。

Androulaki等人[16]研究了通道在Fabric上的应用。然而，这项工作并没有提出一个绩效评估，以定量地确定其方法的效益。

raman等人[17]研究了当区块链用于存储大型数据集分析产生的中间结果时，使用有损压缩来降低结构背书人和验证人之间共享状态的通信成本。然而，他们的方法只适用于对有损压缩不敏感的场景，这不是基于区块链的应用程序的一般情况。

一些研究已经检查了Fabric的性能，但没有提出内部结构的变化。例如，Dinh等人使用Blockbench[5]这一工具来研究私有区块链的性能，研究Fabric的性能，并将其与以太坊和奇偶校验的性能进行比较。他们发现，由于消息通道中的聚集，他们研究的结构版本没有扩展到超过16个节点。nasir等人[18]比较了fabric 0.6和1.0的性能，发现1.0版本的性能优于0.6版本，这并不奇怪。Baliga等人[19]表明，应用程序级参数（如交易的读写集大小、链码和事件负载大小）显著影响交易延迟。类似地，Pongnumkul等人[20]比较了Fabric和以太坊对于加密货币工作负载的性能，发现Fabric在所有指标上都优于以太坊。bergman[21]将fabric与apache cassandra在类似环境中的性能进行了比较，发现对于少量Peer节点，fabric在读重工作负载中的可线性化交易的延迟比cassandra低。另一方面，由于节点数量较多，或者写的工作负载很重，cassandra有更好的性能。

## 6. 总结

这项工作的主要贡献是展示如何对每一个错误的BuffStand框架，如HyffeDiger-Frand，可以重新设计，以支持每秒近20000个交易，一个比现有工作好7的因素。我们通过实现一系列独立的优化来实现这个目标，这些优化集中在I/O、缓存、并行性和高效的数据访问上。在我们的设计中，排序服务只接收交易id而不是完整的交易，并且Peer上的验证是高度并行化的。我们还使用主动缓存，并利用轻量级数据结构在关键路径上快速访问数据。在未来的工作中，我们希望通过以下方式进一步提高hyperledger fabric的性能： •结合有效的bft一致性算法，如rcanopus[22] •在不打开整个交易头的情况下，加快为排序服务提取交易ID•替换现有的加密计算库，提供更有效率的库 •通过分配一个单独的每个通道的排序和快速Peer服务器 •使用分布式框架，如apache spark[23]

## 引用

[1] V. Espinel, D. O’Halloran, E. Brynjolfsson, and D. O’Sullivan, “Deep shift, technology tipping points and societal impact,” in New York: World Economic Forum–Global Agenda Council on the Future of Software & Society (REF 310815), 2015.
[2] M. Vukolic ́, “The quest for scalable blockchain fabric: Proof-of-work vs. bft replication,” in International Workshop on Open Problems in Network Security. Springer, 2015, pp. 112–125.
[3] J. Sousa, A. Bessani, and M. Vukolic, “A Byzantine fault-tolerant ordering service for the hyperledger fabric blockchain platform,” in 2018 48th Annual IEEE/IFIP International Conference on Dependable Systems and Networks (DSN). IEEE, 2018, pp. 51–58.
[4] M. Yin, D. Malkhi, M. K. Reiter, G. Golan Gueta, and I. Abraham, “HotStuff: BFT Consensus in the Lens of Blockchain,” arXiv preprint arXiv:1803.05069, 2018.
[5] T. T. A. Dinh, J. Wang, G. Chen, R. Liu, B. C. Ooi, and K.-L. Tan, “BLOCKBENCH: A Framework for Analyzing Private Blockchains,” Proceedings of the 2017 ACM International Conference on Management of Data - SIGMOD ’17, pp. 1085–1100, 2017.
[6] P. Thakkar, S. Nathan, and B. Vishwanathan, “Performance Benchmarking and Optimizing Hyperledger Fabric Blockchain Platform,” arXiv, 2018.
[7] Hyperledger Fabric, “[FAB-12221] Validator/Committer refactor - Hyperledger JIRA.” [Online]. Available: [https://jira.hyperledger.org/](https://link.zhihu.com/?target=https%3A//jira.hyperledger.org/) browse/FAB- 12221?filter=12526
[8] C. Cachin, “Architecture of the hyperledger blockchain fabric,” in Workshop on Distributed Cryptocurrencies and Consensus Ledgers, vol. 310, 2016.
[9] E. Androulaki, A. Barger, V. Bortnikov, C. Cachin, K. Christidis, A. De Caro, D. Enyeart, C. Ferris, G. Laventman, Y. Manevich, S. Muralidharan, C. Murthy, B. Nguyen, M. Sethi, G. Singh, K. Smith, A. Sorniotti, C. Stathakopoulou, M. Vukolic ́, S. W. Cocco, and J. Yellick, “Hyperledger Fabric: A Distributed Operating System for Permissioned Blockchains,” Proceedings of the Thirteenth EuroSys Conference on - EuroSys ’18, pp. 1–15, 2018.
[10] ApacheFoundation,“ApacheKafka,ADistributedStreamingPlatform,” 2018. [Online]. Available: [https://kafka.apache.org/](https://link.zhihu.com/?target=https%3A//kafka.apache.org/) (Accessed 2018-12- 05).
[11] Cloud Native Computing Foundation, “gRPC: A high performance, open-source universal RPC framework,” 2018. [Online]. Available: [https://grpc.io/](https://link.zhihu.com/?target=https%3A//grpc.io/)
[12] S. Bano, A. Sonnino, M. Al-Bassam, S. Azouvi, P. McCorry, S. Meik- lejohn, and G. Danezis, “Consensus in the age of blockchains,” arXiv preprint arXiv:1711.03936, 2017.
[13] Google Developers, “Protocol Buffers — Google Developers,” 2018. [Online]. Available: [https://developers.google.com/protocol-buffers/?hl=](https://link.zhihu.com/?target=https%3A//developers.google.com/protocol-buffers/%3Fhl%3D) en
[14] A. Sharma, F. M. Schuhknecht, D. Agrawal, and J. Dittrich, “How to databasify a blockchain: the case of hyperledger fabric,” arXiv preprint arXiv:1810.13177, 2018.
[15] A. Bessani, J. Sousa, and E. Alchieri, “State machine replication for the masses with BFT-SMART,” in DSN, 2014, pp. 355–362.
[16] E. Androulaki, C. Cachin, A. De Caro, and E. Kokoris-Kogias, “Channels: Horizontal scaling and confidentiality on permissioned blockchains,” in European Symposium on Research in Computer Se-
curity. Springer, 2018, pp. 111–131.
[17] R. K. Raman, R. Vaculin, M. Hind, S. L. Remy, E. K. Pissadaki, N. K. Bore, R. Daneshvar, B. Srivastava, and K. R. Varshney, “Trusted multi- party computation and verifiable simulations: A scalable blockchain approach,” arXiv preprint arXiv:1809.08438, 2018.
[18] Q. Nasir, I. A. Qasse, M. Abu Talib, and A. B. Nassif, “Performance analysis of hyperledger fabric platforms,” Security and Communication Networks, vol. 2018, 2018.
[19] A. Baliga, N. Solanki, S. Verekar, A. Pednekar, P. Kamat, and S. Chat- terjee, “Performance Characterization of Hyperledger Fabric,” in Crypto Valley Conference on Blockchain Technology, CVCBT 2018, 2018.
[20] S. Pongnumkul, C. Siripanpornchana, and S. Thajchayapong, “Performance analysis of private blockchain platforms in varying workloads,” in 2017 26th International Conference on Computer Communications and Networks, ICCCN 2017. IEEE, 7 2017, pp. 1–6.
[21] S. Bergman, “Permissioned blockchains and distributed databases: A performance study,” Master’s thesis, Linkoping University, 2018.
[22] S. Keshav, W. M. Golab, B. Wong, S. Rizvi, and S. Gorbunov, “RCano- pus: Making canopus resilient to failures and byzantine faults,” arXiv preprint arXiv:1810.09300, 2018.
[23] Apache Foundation, “Apache Spark - Unified Analytics Engine for Big Data,” 2018. [Online]. Available: [https://spark.apache.org](https://link.zhihu.com/?target=https%3A//spark.apache.org)

> 本文翻译自：[https://arxiv.org/pdf/1901.00910.pdf](https://link.zhihu.com/?target=https%3A//arxiv.org/pdf/1901.00910.pdf) 版权归作者所有，商业使用请联系作者



# 去中心化架构介绍&Fabcar项目详解

在刚刚接触到区块链的时候，我对智能合约，区块链网络，区块链应用三者之间的关系，一直很不清楚，没有一个很宏观的轮廓，不知道他们究竟是怎样协同工作的。后来，在慢慢的摸索中，开始阅读HyperLedger的官方文档，以及几个简单的区块链应用系统的介绍。渐渐明白了它们之间的关系。所以，我决定在介绍ChainCode的开发之前，先来简单的把一个区块链应用的基本组成部分。总的来说，我们的学习策略是，由**全面的肤浅**到**片面的深刻**。

接下来的文章，将按照如下的结构来展开：

![img](https://pic2.zhimg.com/80/v2-4b8a3a1b90832c028ddd3a7ec1400ee5_720w.jpg)全文结构

------

## **一、传统应用系统与简单区块链应用系统的对比**

- **传统应用系统的组成与架构**

在我们日常的一个应用系统大体上是B/S 架构或者C/S架构，移动端或者前端来实现用户的操作界面，后台有应用服务器和数据库服务器来进行数据和事务逻辑的处理。B/S架构的系统组成如下图所示：

![img](https://pic1.zhimg.com/80/v2-01b4dd8829294a6ff858a99adc8661ec_720w.jpg)B/S架构示意图

在实际的业务逻辑开发中，传统的是使用分层架构来实现用户，事务与数据之间的关系。经典三次架构图如下图所示：

![img](https://pic1.zhimg.com/80/v2-0948c15f9dd92584f2a0c0a5a982c55c_720w.jpg)三层架构图

1：数据访问层：主要是对非原始数据（数据库或者文本文件等存放数据的形式）的操作层，而不是指原始数据，也就是说，是对数据库的操作，而不是数据，具体为业务逻辑层或表示层提供数据服务。

2：业务逻辑层：主要是针对具体的问题的操作，也可以理解成对数据层的操作，对数据业务逻辑处理，如果说数据层是积木，那逻辑层就是对这些积木的搭建。

3：界面层：主要表示WEB方式，也可以表示成WINFORM方式，WEB方式也可以表现成：aspx，如果逻辑层相当强大和完善，无论表现层如何定义和更改，逻辑层都能完善地提供服务。



- **区块链应用系统的组成与架构**

在HyperLedger的区块链应用系统中，从组成部分来说，和传统的应用区别不大。用来和用户交互的依然是浏览器或者客户端或者移动端。不同的是，在这里，我们是把数据存储到区块链网络中，而不是存在数据库服务器。下面是一个简单弹珠管理系统的组成部分示意图：

![img](https://pic4.zhimg.com/80/v2-1c66fe97a4859120b1dad685092c6267_720w.jpg)弹珠管理系统-系统组成部分

在上图所示的弹珠管理系统中。我们在浏览器端操作弹珠的管理。前端调用后台应用程序，而后台应用程序通过调用运行在区块链网络中的智能合约，来实现对数据的存储与更改等。所以，如果我们要开发一个区块链应用，需要编写**前端程序**，**后端程序**，还要编写**智能合约**，**部署Fabric区块链网络**等。

而在区块链应用系统中采用的系统架构，并非传统的分层架构，而是微服务架构。它的主要作用是将功能分解到离散的各个服务当中，从而降低系统的耦合性，并提供更加灵活的服务支持。在微服务架构中，是按照业务，而不是技术来划分组织。下图展示了Fabric网络的基本架构：

![img](https://pic1.zhimg.com/80/v2-9099b3b4870820e3581d2c49facf77dc_720w.jpg)Fabric架构图

------

## **二、FabCar的区块链应用实战**

**简介**

在接下来的部分，我们将运行并解读HyperLedger官方提供的一个小的例子FabCar。这个例子，很好的展示了后台程序对ChainCode的调用，并来操作区块链网络账本的。这一部分，主要是参考：[HyperLedger中文文档-编写第一个应用](https://link.zhihu.com/?target=https%3A//hyperledgercn.github.io/hyperledgerDocs/write_first_app_zh/)

区块链网络应用程序需要提供给用户**查询**账本（包含特定记录）以及**更新**账本（添加记录）的功能。我们的应用程序基于Javascript，通过Node.js SDK与（账本所在的）网络进行交互。这一部分将通过三步来编写第一个应用程序。

- **1. 启动一个Hyperledger Fabric区块链测试网络。** 在我们的网络中，我们需要一些最基本的组件来查询和更新账本。这些组件 —— peer节点、ordering节点以及证书管理 —— 是我们网络的基础。而CLI容器则用来发送一些管理命令。
- **2. 学习应用程序中所用到的智能合约例子的参数。** 智能合约包含的各种功能让我们可以用多种方式和账本进行交互。如，我们可以读取整体的数据或者某一部分详尽的数据。
- **3. 开发能够查询以及更新记录的应用程序。** 我们提供两个程序例子 —— 一个用于查询账本，另一个用户更新账本。我们的程序将使用SDK APIs来和网络进行交互，并最终调用这些功能。

完成这一部分实战后，我们应该会基本了解一个使用Hyperledger Fabric Node.js SDK并带有智能合约的应用程序，是如何与Hyperledger Fabric网络中的账本进行交互的。



**环境准备**

- **安装go,下载Fabric源码以及下载好docker镜像。**这一部分，在《[1-HyperLedger实战-快速搭建一个Fabric1.0环境](https://zhuanlan.zhihu.com/p/35063055)》这篇文章里有详细的介绍。
- **安装node.**因为这文是利用nodeSDK 来进行开发的，所以，我们需要安装node。具体命令如下：

```text
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get install -y nodejs
```

> **注意**！Fabric Node SDK支持的Node版本是v6，不支持最新的v8版本。

安装完成后我们可以使用以下两个命令来查看安装的Node版本和npm版本。

```text
node –v
npm -v
```



**实战步骤**

- **1-下载测试网络**

这里我们进行测试的代码，是官方托管在GitHub上的，下载并进入fabcar子目录。

```text
git clone https://github.com/hyperledger/fabric-samples.git
cd fabric-samples/fabcar
```

这个子目录 - `fabcar` - 包含运行示例程序的脚本以及程序代码。在该目录运行`ls`命令，您应该会看到以下内容：

```text
enrollAdmin.js  node_modules  query.js         startFabric.sh
invoke.js       package.json  registerUser.js
```

现在调用startFabric.sh来启动网络。

```text
./startFabric.sh
```

这个命令主要做了如下工作：

1.启动peer节点、Ordering节点，证书颁发机构，CLI容器等。

2.创建一个通道，并将peer加入该通道

3.将智能合约（即链码）安装到peer节点的文件系统上。

4.在通道上实例化该链码；实例化会启动链码容器。

出现如下界面，表示测试网络成功运行：

![img](https://pic4.zhimg.com/80/v2-eafd8a57fcc624de1469e5193b920eb7_720w.jpg)网络成功运行界面

上面的步骤已经成功运行了区块链网络，接下来就要着手应用开发了。这里，我们首先来介绍一下用NodeSDK来开发调用智能合约的基本步骤：

- **2-编写package.json并下载依赖模块：**

fabcar/package.json中的内容如下：

```text
{
    "name": "fabcar",
    "version": "1.0.0",
    "description": "Hyperledger Fabric Car Sample Application",
    "main": "fabcar.js",
    "scripts": {
        "test": "echo \"Error: no test specified\" && exit 1"
    },
    "dependencies": {
        "fabric-ca-client": "~1.1.0",
        "fabric-client": "~1.1.0",
        "grpc": "^1.6.0"
    },
    "author": "Anthony O'Dowd",
    "license": "Apache-2.0",
    "keywords": [
        "Hyperledger",
        "Fabric",
        "Car",
        "Sample",
        "Application"
    ]
}
```

我们就可以运行npm install命令来下载所有相关的依赖模块，但是由于npm服务器在国外，所以下载可能会很慢，感谢淘宝为我们提供了国内的npm镜像，使得安装npm模块快很多。运行的命令是：

```text
npm install --registry=https://registry.npm.taobao.org
```

运行完毕后我们查看一下fabcar目录，可以看到多了一个**node_modules**文件夹。这里就是使用刚才的命令下载下来的所有依赖包。

- **3-运行enrollAdmin.js**

运行下列命令：

```text
node enrollAdmin.js 
```

运行成功后界面如下：

![img](https://pic4.zhimg.com/80/v2-b69d32ac79ecefefdb3c245f1b4d3457_720w.jpg)

会在fabcar目录下生成一个存放key的文件夹：hfc-key-store

- **4-运行registerUser.js**

```text
node registerUser.js
```

成功运行后界面如下：

![img](https://pic3.zhimg.com/80/v2-633652fd4f445a3310cc492fae08d25a_720w.jpg)

- **5-运行query.js**

现在我们可以运行JavaScript程序。运行`query.js` 程序，返回账本上所有汽车列表。程序中预先加载了一个`queryAllCars`函数，用于查询所有车辆，因此我们可以简单地运行程序

```text
node query.js
```

运行成功后，界面如下：

![img](https://pic1.zhimg.com/80/v2-628a31e0c836fb77c274afc2c100a03c_720w.jpg)

- **6.关闭网络**

> 在本地开发测试过程中，当关闭一个项目是，一定要记得清除所有的允许容器。否则的话，未关闭的容器，可能会占用端口号，给下一个项目的运行带来干扰。

进入上一级目录，basic-network,运行关闭网络的脚本

```text
./teardown.sh
```

------

## **三、详解应用程序与网络交互过程**

**简介**

通过前面的工作，我们有了简单的网络以及一些代码，现在看看他们是怎么一起工作的。

应用程序使用**APIs**来调用智能合约(即“链码”)而API可通过软件开发工具包（SDK）访问。

在本练习中，我们将使用[Hyperledger Fabric Node SDK](https://link.zhihu.com/?target=https%3A//fabric-sdk-node.github.io/)，除此以外，Fabric还提供了Java SDK和CLI用于开发应用程序。



**链码的工作流程**

我们首先来介绍一下链码的工作流程：

![img](https://pic1.zhimg.com/80/v2-d52ee3484d0800a883c92bdd186b0f0c_720w.jpg)链码工作流程

我们知道，在Fabric中，链码运行在节点上的沙盒（Docker容器）中，被调用时的基本工作流程如上图所示。

- 首先，用户通过客户端（SDK或CLI），向Fabric的背书节点（endorser）发出调用链码的交易提案（proposal）。
- 然后，节点对提案进行包括ACL权限检查在内的各种检验，通过后则创建模拟执行这一交易的环境。
- 接着，背书节点和链码容器之间通过gRPC消息来交互，模拟执行交易并给出背书结论。
- 最后，客户端收到足够的背书节点的支持后，便可以将这笔交易发送给排序节点（orderer）进行排序，并最终写入区块链。

链码容器的**shim层**是节点与链码交互的中间层。当链码的代码逻辑需要读写账本时，链码会通过shim层发送相应操作类型的ChaincodeMessage给节点，节点本地操作账本后返回响应消息。

在上面的流程中，出现了一个重要的部分，shim层。这里需要再次强调一下，**链码容器的shim层是节点与链码交互的中间层。**



**ChainCode的编写简介**

每个chaincode程序都必须实现 ***chaincode接口\*** ，接口中的方法会在响应传来的交易时被调用。

`Init`（初始化）方法会在chaincode接收到`instantiate`（**实例化**）或者`upgrade`(**升级**)交易时被调用，进而使得chaincode顺利执行必要的初始化操作，包括初始化应用的状态。`Invoke`（调用）方法会在响应`invoke`（**调用）交易时被调用**以执行交易。



**ChainCode编写的步骤**

**1.引入相关包**

首先， 我们先进行准备工作。对于每一个chaincode，它都会实现预定义的***chaincode接口\***，特别是`Init`和`Invoke`函数接口。所以我们首先为我们的chaincode引入必要的依赖。这里的shim层是节点与链码交互的中间层。如下图所示：

```text
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)
```



**2.初始化Chaincode**

接下来，我们将实现`Init`函数。

```text
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}
```

值得留意的是chaincode升级同样会调用该函数。当我们编写的chaincode会升级现有chaincode时，需要确保适当修正Init函数。特别地，如果没有“迁移”操作或其他需要在升级中初始化的东西，那么就提供一个空的“Init”方法。我们这里仅仅提供了一个简单的Init方法。



**3.编写Chaincode接口函数**

首先，添加`Invoke`函数签名。

```text
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

}
```

我们需要调用`ChaincodeStubInterface`来获取参数。我们将调用*ChaincodeStubInterface*并以键值为参数传入。如果一切正常，那么我们会收到表明初始化成功的peer.Response返回对象。`Invoke`函数所需的传入参数正是应用想要调用的chaincode的名称。在我们的应用里面，我们有几个简单的功能函数：queryCar,initLedger,createCar,queryAllCars,changeCarOwner等。

下面，我们将使这几个函数名正式生效，并调用这些chaincode应用函数，经由`shim.Success`或`shim.Error`函数返回一个合理的响应。这两个`shim`成员函数可以将响应序列化为gRPC protobuf消息

```text
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "queryCar" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCar" {
		return s.createCar(APIstub, args)
	} else if function == "queryAllCars" {
		return s.queryAllCars(APIstub)
	} else if function == "changeCarOwner" {
		return s.changeCarOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}
```



**4.编写Chaincode的调用函数**

如上文所述，我们的chaincode应用实现了五个函数，并可以被`Invoke`函数调用。下面我们就来真正实现这些函数。注意，就像上文一样，我们调用chaincode shim API中的*ChaincodeStubInterface.PutState*和*ChaincodeStubInterface.GetState*函数来访问账本。

```text
func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	i := 0
	for i < len(cars) {
		fmt.Println("i is ", i)
		carAsBytes, _ := json.Marshal(cars[i])
		APIstub.PutState("CAR"+strconv.Itoa(i), carAsBytes)
		fmt.Println("Added", cars[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}

	carAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAR0"
	endKey := "CAR999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Owner = args[1]

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
```

这里可以看到，在最后关头我们写了main函数，它将调用*shim.Start* 函数，main函数的作用，是在容器里启动chaincode。



**5.整合全部代码**

将上面的代码整合在一起如下：

```text
/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCar" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCar" {
		return s.createCar(APIstub, args)
	} else if function == "queryAllCars" {
		return s.queryAllCars(APIstub)
	} else if function == "changeCarOwner" {
		return s.changeCarOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	i := 0
	for i < len(cars) {
		fmt.Println("i is ", i)
		carAsBytes, _ := json.Marshal(cars[i])
		APIstub.PutState("CAR"+strconv.Itoa(i), carAsBytes)
		fmt.Println("Added", cars[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}

	carAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAR0"
	endKey := "CAR999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Owner = args[1]

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
```



**query.js的源码分析**

回到我们实战部分的第五个步骤。运行node query.js,系统给我们是十辆车的记录。现在让我们来看看代码内容。使用编辑器（例如atom或visual studio）打开`query.js`程序。

以下是构建查询的代码块：

```text
const request = {
		chaincodeId: 'fabcar',
		fcn: 'queryAllCars',
		args: []
	};
```

我们将`chaincode_id`变量赋值为`fabcar`- 这让我们定位到这个特定的链码 - 然后调用该链码中定义的`queryAllCars`函数。

在前面，我们发出`node query.js`命令时，会调用了特定函数来查询账本。但是，这不是我们能够使用的唯一功能。

在fabric-sample目录下，上文我们编写的ChainCode的代码其实是学习在`chaincode`子目录中的`fabcar.go`。

从上面可知，我们可以调用下面的函数- `initLedger`、`queryCar`、`queryAllCars`、`createCar`和`changeCarOwner`。让我们仔细看看`queryAllCars`函数是如何与账本进行交互的。

该函数调用shim接口函数`GetStateByRange`来返回参数在`startKey`和`endKey`间的账本数据。这两个键值分别定义为`CAR0`和`CAR999`。因此，我们理论上可以创建1,000辆汽车（假设Keys都被正确使用），`queryAllCars`函数将会显示出每一辆汽车的信息。

下图演示了一个应用程序如何在链码中调用不同的功能。

![img](https://pic2.zhimg.com/80/v2-56d15f04d11d3b89fe2d39a09d99c63d_720w.jpg)应用程序调用ChainCode示意图



我们可以看到我们用过的`queryAllCars`函数，还有一个叫做`createCar`，这个函数可以让我们更新账本，并最终在链上增加一个新区块。但首先，让我们做另外一个查询。

现在我们返回`query.js`程序并编辑请求构造函数以查询特定的车辆。为达此目的，我们将函数`queryAllCars`更改为`queryCar`并将特定的“Key” 传递给args参数。在这里，我们使用`CAR4`。 所以我们编辑后的`query.js`程序现在应该包含以下内容：

```text
const request = {
    chaincodeId: options.chaincode_id,
    txId: transaction_id,
    fcn: 'queryCar',
    args: ['CAR4']
}
```

保存程序并返回`fabcar`目录。现在再次运行程序：

```text
node query.js
```

您应该看到以下内容：

```text
{"colour":"black","make":"Tesla","model":"S","owner":"Adriana"}
```

按照上面的步骤，我们可以继续调用别的slim接口函数，这里不再累赘。具体步骤可以参考：

[Fabric中文档-编写第一个应用](https://link.zhihu.com/?target=https%3A//hyperledgercn.github.io/hyperledgerDocs/write_first_app_zh/)

## **4.总结**

这篇文章，主要是介绍了ChainCode的编写，以及如何使用JavaScript代码来调用智能合约。将官方的fabcar运行过程介绍了一下。因为在官方中文文档中，有一些步骤，并不完全正确。这里参考了一些网上其他人的博客。

至此，ChainCode的编写基本学习总结完了。后面将继续片面的深刻学习Fabric。



参考文章：

《区块链原理、设计、应用》-杨保华-13.2.3　链码基本工作原理

[Fabric中文档-编写第一个应用](https://link.zhihu.com/?target=https%3A//hyperledgercn.github.io/hyperledgerDocs/write_first_app_zh/)

# FastFabric:提升Hyperledger Fabric性能到20000TPS

> 本文翻译自：https://arxiv.org/pdf/1901.00910.pdf
> 版权归作者所有，商业使用请联系作者

## 摘要

预计区块链技术将对各种行业产生重大影响。然而，阻碍它们的一个问题是它们有限的交易吞吐量，特别是与诸如分布式数据库系统之类的已建立的解决方在本文中，我们重新构建了一个现代许可的区块链系统Hyperledger Fabric，以将交易吞吐量从每秒3,000次增加到20,000次。我们专注于超出共识机制的性能瓶颈，我们提出架构更改，以减少交易排序和验证期间的计算和I / O开销，从而大大提高吞吐量。值得注意的是，我们的优化是完全即插即用的，不需要对Hyperledger Fabric进行任何界面更改。

## 1. 介绍

区块链等分布式账本技术提供了一种以安全和可验证的方式进行交易的方法，而无需可信任的第三方。因此，人们普遍认为区块链将对金融、房地产、公共管理、能源和交通行业产生重大影响[1]。但是，为了在实践中可行，区块链必须支持与现有数据库管理系统支持的交易率相当的交易率，这可以提供一些相同的交易保证。

与不限制网络成员资格的无权限区块链相比，我们专注于许可的区块链，其中所有参与节点的身份都是已知的。许可的区块链适用于许多应用领域，包括融资;例如，Ripple区块链旨在提供类似于当前SWIFT系统的货币兑换和跨银行交易的支付网络。

从技术角度来看，我们观察到这两种区块链之间的重要区别。无权区块链的无信任性质要求工作或利益证明，或昂贵的全球规模的拜占庭 - 容错共识机制[2]。最近的许多工作都集中在使这些更有效。另一方面，许可的区块链通常将共识和交易验证委托给选定的节点组，从而减轻了共识算法的负担。虽然在这种情况下，共识仍然是一个瓶颈，但最近正在解决这个问题https：//ripple.com工作[3]，[4]，激励我们超越共识来确定进一步的绩效改进。

在本文中，我们批判性地研究了Hyperledger Fabric 1.2的设计，因为据报道它是最快的开源许可区块链[5]。虽然已经有一些关于优化Hyperledger Fabric的工作，例如，使用积极的缓存[6]，但我们并不知道任何先前关于重新构建系统的工作2。因此，我们基于通用系统设计技术设计并实现了几种架构优化，这些技术将端到端交易吞吐量提高了近7倍，从每秒3,000到20,000个交易处理，同时减少了块延迟。我们的具体贡献如下：

1）从数据中分离元数据：Fabric中的共识层接收整个交易作为输入，但只需要交易ID来决定交易顺序。我们重新设计Fabric的交易排序服务，仅使用交易ID，从而大大提高了吞吐量。

2）并行和缓存：交易验证的某些方面可以并行化，而其他方面可以从缓存数据中受益。我们重新设计了Fabric的验证服务，通过积极缓存调度器中的未编组块并通过并行化尽可能多的验证步骤，包括认可策略验证和语法验证。

3）利用存储器层次结构在关键路径上进行快速数据访问：Fabric维护世界状态的键值存储可以用轻量级内存数据结构代替，其缺乏耐久性保证可以通过区块链本身进行补偿。我们围绕轻量级哈希表重新设计Fabric的数据管理层，该表可以更快地访问关键交易验证路径上的数据，从而将不可变块的存储推迟到写优化存储集群。 4）资源分离：提交者和背书者的对等角色争夺资源。我们介绍了一种将这些角色转移到单独硬件的架构。

重要的是，我们的优化不会违反Fabric的任何API或模块化边界，因此它们可以合并到Fabric版本2.0的计划版本中[7]。

2有关相关工作的详细调查，请参阅第V节。

我们还概述了未来工作的几个方向，与我们提出的优化一起，有可能达到Visa等信用卡公司所要求的每秒50,000笔交易[2]。

在本文的其余部分，第二部分简要概述了Hyperledger Fabric，第三部分介绍了我们的改进设计，第四部分讨论了实验结果，第五部分介绍了我们在先前工作中的贡献，第六部分总结了未来的发展方向工作。

## 2. Fabric架构

作为由Linux Foundation托管的开源Hyperledger项目的一个项目，Fabric是最活跃开发的许可区块链系统之一[8]。由于Androulaki等[9]详细描述了交易流程，因此我们仅提供一个简短的概要，重点介绍我们在第III节中提出改进的系统部分。

为了避免智能合约确定性的陷阱并允许即插即用的系统组件替换，Fabric的结构与其他常见的区块链系统不同。交易遵循执行顺序提交流模式而不是公共顺序执行提交模式。客户端交易首先在沙箱中执行以确定它们的读写集，即由交易读取和写入的健值对集。然后，交易由排序服务排序，并最终验证并提交给区块链。此工作流程由分配了特定角色的节点实现，如下所述。

A. 节点类型

客户端发起交易，即对区块链的读取和写入，发送到Fabric节点。节点是Peer或排序者Orderer;一些Peer也是背书者。所有Peer都将块提交到区块链的本地副本，并将相应的更改应用于维护当前世界状态快照的状态数据库。允许背书者节点根据链码，Fabric的智能合约版本中捕获的业务规则来证明交易有效。排序者仅负责决定交易顺序，而不是正确性或有效性。

B. 交易流程

客户将其交易发送给一些背书者。每个背书者在沙箱中执行交易，并计算相应的读写集以及访问的每个键的版本号。每个背书者还使用业务规则来验证交易的正确性。客户等待足够数量的认可，然后将这些响应发送给排序服务的Orderer。排序者首先就进入的交易的顺序达成共识，然后将消息队列分成块。块被传递给Peer节点，然后Peer验证并提交它们。由于其许可性质，所有节点必须已知并向成员服务提供商（MSP）注册，否则其他节点将忽略它们。

C. 实现细节

为了讨论第III节中的改进，我们现在仔细研究一下orderer和peer架构。

1）排序者：在收到来自背书者的回复后，客户端创建包含标题和有效负载的交易提议。标题包括交易ID和通道ID。有效负载包括读写集和相应的版本号，以及支持的Peer节点的签名。交易提议使用客户的凭证签署并发送到排序服务。

排序服务的两个目标是（a）在交易顺序上达成共识，以及（b）将包含有序交易的块交付给提交者Peer节点。Fabric目前使用Apache Kafka（基于ZooKeeper [10]）来实现容错崩溃的共识。

排序者收到交易提案时，它会检查客户是否有权提交交易。如果是，则orderer将交易提议发布到Kafka集群，其中每个Fabric通道都映射到Kafka主题以创建相应的不可变的交易序列顺序。然后，每个排序者根据每个块允许的最大交易数或块超时周期将从Kafka接收的交易组装成块。使用orderer的凭证对块进行签名，并使用gRPC [11]将其传递给Peer节点。

2）Peer：从排序服务接收消息时，Peer首先从区块的头部和元数据检查其语法结构。然后检查orderer的签名是否符合指定的策略。任何这些测试失败的区块都会被立即丢弃。

在初始校验之后，区块被推入队列，保证其添加到区块链。但是，在此之前，区块会依次执行两个验证步骤和最后一个提交步骤。

在第一个验证步骤中，将解压缩块中的所有交易，检查其语法并验证其认可。未通过此测试的交易将被视为无效，但会保留在块中。此时，只有善意创建的交易仍然有效。

在第二个验证步骤中，Peer确保有效交易之间的相互作用不会导致无效的世界状态。回想一下，每个交易都带有一组需要从世界状态数据库（它的读取集）读取的键，以及它将写入数据库（它的写集）的一组键和值，以及它们记录的版本号通过背书者。在第二个验证步骤中，交易的读写集中的每个键仍必须具有相同的版本号。从任何先前交易中写入该Key会更新版本号并使交易无效。这可以防止双重支出。

在最后一步中，Peer将块（现在包括其交易的验证标志）写入文件系统。

Fabric被虚拟化为多个通道，由通道ID识别。

Key及其值（即世界状态）将保留在LevelDB或CouchDB中，具体取决于应用程序的配置。此外，每个块及其交易的索引都存储在LevelDB中以加速数据访问。

## 3. 设计

本节介绍我们对Fabric 1.2版的体系结构和实现的更改。该版本于2018年7月发布，随后于2018年9月发布1.3版，2019年1月发布1.4版。但是，最近发布的版本中引入的更改不会影响我们的提议，因此我们预计在将我们的工作与新版本集成时不会有任何困难。重要的是，我们的改进使各个模块的接口和职责保持不变，这意味着我们的更改与现有的Peer或排序服务实现兼容。此外，我们的改进是相互正交的，因此可以单独实施。对于Orderer和Peer，我们按照从最小到最大的性能影响的升序来描述我们的提议，与其对Fabric的各自更改相比。

A. 准备

使用拜占庭容错（BFT）一致性算法是HyperLedger中的关键性能瓶颈[2]。这是因为BFT一致性算法不能很好地与参与者的数量成比例。在我们的工作中，我们选择超越这个明显的瓶颈有三个原因：

•可以说，在许可的区块链中使用BFT协议并不像在无权限的系统中那么重要，因为所有参与者都是已知的并且被激励以保持系统以诚实的方式运行。

•正在广泛研究BFT共识[12]，我们预计未来一两年内将出现更高吞吐量的解决方案。

•实际上，Fabric 1.2不使用BFT共识协议，而是依赖于Kafka进行交易排序，如前所述。

出于这些原因，我们的工作目标不是使用更好的BFT一致性算法来提高orderer性能，而是为了缓解当共识不再是瓶颈时出现的新问题。我们首先对排序服务提出两项改进，然后对Peer进行一系列改进。

B. Orderer改进I：从有效负载中分离交易头

在Fabric 1.2中，使用Apache Kafka的订购者将整个交易发送给Kafka进行订购。交易的长度可能是几千字节，导致高通信开销，从而影响整体性能。但是，就交易订单达成共识只需要交易ID，因此我们可以通过仅向Kafka集群发送交易ID来获得订货人吞吐量的显着改善。

具体而言，在从客户端接收交易时，我们的订货人从标题中提取交易ID并发布图1.新的订货人架构。传入的交易同时处理。他们的TransactionID被发送到Kafka集群进行订购。当接收到有序的TransactionID时，订货人用它们的有效负载重新组装它们并将它们收集到块中。

![img](http://img.mochain.info/topjohn/blog/post/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-09-19%20%E4%B8%8B%E5%8D%885.44.28.png)

这个ID到Kafka集群。排序者将相应的有效负载分别存储在本地数据结构中，并在从Kafka收回ID时重新组装交易。随后，与Fabric一样，orderer将交易集分段为块并将它们传递给Peer。值得注意的是，我们的方法适用于任何共识实现，并且不需要对现有排序界面进行任何修改，从而允许我们利用现有的Fabric客户端和Peer代码。

C. Orderer改进II：消息流水线

在Fabric 1.2中，订购服务逐个处理来自任何给定客户端的传入交易。当交易到达时，识别其相应的信道，根据一组规则检查其有效性，并最终将其转发到共识系统，例如，卡夫卡;只有这样才能处理下一个交易。相反，我们实现了一个可以同时处理多个传入交易的流水线机制，即使它们来自使用相同gRPC连接的同一客户端。为此，我们维护一个线程池，它并行处理传入的请求，每个传入的请求都有一个线程。线程调用Kafka API来发布交易ID，并在成功时向客户端发送响应。订货人完成的剩余处理与Fabric 1.2相同。

图1总结了新的排序设计，包括交易ID与有效负载的分离以及由于并行消息处理导致的横向扩展。

D. Peer任务

回忆一下第II-C2节，在从排序者接收块时，Fabric Peer按顺序执行以下任务：

•验证收到消息的合法性•验证块中每个交易的块头和每个认可签名•验证交易的读写集•更新LevelDB或CouchDB中的世界状态•将区块链日志存储在文件系统中，与LevelDB中的相应索引

我们的目标是在交易流程的关键路径上最大化交易吞吐量。为此，我们进行了广泛的调用图分析，以确定性能瓶颈。

![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-20%E4%B8%8B%E5%8D%884.22.49.png)

图2.新的Peer结构。快速Peer使用内存中的哈希表来存储世界状态。验证通道完全并发，并行验证多个块及其交易。代言人角色和持久存储被分成可扩展的集群，并由快速Peer给出经过验证的块。通道的所有部分都使用缓存中的未编组块。

我们做出以下观察。首先，验证交易的读写集需要快速访问world状态。因此，我们可以通过使用内存中的哈希表而不是数据库来加速该过程（第III-E节）。其次，交易流程不需要区块链日志，因此我们可以在交易流程结束时将其存储到专用存储和数据分析服务器（第III-F节）。第三，如果Peer也是背书者，则需要处理新的交易提案。但是，提交者和背书者角色是不同的，这使得为每项任务专用不同的物理硬件成为可能（第III-G节）。第四，必须在Peer验证和解决传入的块和交易。最重要的是，必须按顺序完成通过交易写入集验证状态更改，阻止所有其他任务。因此，尽可能加快这项任务非常重要（第III-H节）。

最后，通过缓存协议缓冲区[13]解块的结果（第III-I节），可以获得显着的性能提升。我们在图2中详细介绍了这种架构重新设计，包括其他提出的Peer改进。

E. Peer改进I：用哈希表替换世界状态数据库

必须为每个交易按顺序查找和更新世界状态数据库，以保证所有Peer的一致性。因此，对此数据存储的更新以尽可能高的交易速率发生是至关重要的。

我们认为，对于常见情况，例如跟踪分类帐上的钱包或资产，世界状态可能相对较小。即使需要存储数十亿个密钥，大多数服务器也可以轻松地将它们保存在内存中。因此，我们建议使用内存中的哈希表而不是LevelDB / CouchDB来存储世界状态。这样可以在更新世界状态时消除硬盘访问。它还消除了由于区块链本身的冗余保证而不必要的昂贵的数据库系统保证（即，ACID属性），进一步提高了性能。当然，由于使用易失性存储器，这种替换易受节点故障的影响，因此必须通过稳定存储来增强内存中的哈希表。我们在第III-F节讨论了这个问题。

F.Peer改进II：使用对等集群存储块

通过定义，块是不可变的。这使它们非常适合仅附加数据存储。通过将数据存储与对等任务的其余部分分离，我们可以设想用于块和世界状态备份的多种类型的数据存储，包括在其文件系统中存储块和世界状态备份的单个服务器，如Fabric目前所做的那样;数据库或键值存储，如LevelDB或CouchDB。为了实现最大扩展，我们建议使用分布式存储集群。请注意，使用此解决方案，每个存储服务器仅包含链的一小部分，从而激励使用分布式数据处理工具，如Hadoop MapReduce或Spark5。

G.Peer改进III：单独承诺和认可

在Fabric 1.2中，endorser peer也负责提交块。认可是一项昂贵的操作，承诺也是如此。虽然对一组背书者进行并发交易处理可能会提高应用程序性能，但在每个新节点上复制承诺的额外工作实际上无效了这一优势。因此，我们建议将这些角色分开。

具体而言，在我们的设计中，提交者节点执行验证通道，然后将经过验证的块发送给背书者群集，这些背书者仅将更改应用于其世界状态而无需进一步验证。此步骤允许我们释放Peer的资源。请注意，这种可以扩展以满足需求的背书者群集只会将对等方的认可角色分割为专用硬件。此群集中的服务器不等同于Fabric 1.2中的完整版本的背书节点。

H. Peer改进IV：并行化验证

块和交易头验证（包括检查发件人的权限，执行认可策略和语法验证）都是高度可并行化的。我们通过引入完整的验证通道来扩展Fabric 1.2的并发性。

具体而言，对于每个传入的块，一个go-routine被分配用于通过块验证阶段。随后，这些例程中的每一个都使用Fabric 1.2中已存在的goroutine池进行交易验证。因此，在任何给定时间，并行检查多个块及其交易的有效性。最后，所有读写集都由单个goroutine以正确的顺序依次验证。这使我们能够充分利用多核服务器CPU的潜力。但是，我们当前的实现不包括这样的存储系统。

I. Peer增强 V: 缓存解析的区块

Fabric使用gRPC在网络中让节点进行通信。Protocol Buffers被用来进行序列化。为了保证处理应用和软件升级，Fabric的区块结构是高度层级化的，每一个层级都是单独序列化和反序列化的。这将导致大量的内存将用来进行byte数组转成结构化的数据。而且，Fabric 1.2没有在缓存中存储之前解析的数据，因此当需要这些数据的时候，这些工作将被重复执行。

为了缓解这个问题，我们计划用一个临时的缓存来存放解析的数据。区块会别缓存，无论当在校验管道还是通过区块号接收的时候。一旦区块的任何部分被解析了，它将会被存储到区块中以便再利用。我们用循环的缓存池实现这一功能，这是一个足够大的校验管道。无论区块是否被提交，一个新的区块可以被自动放进管道覆盖已经存在的位置的区块。由于在提交后不需要缓存，并且保证新块只在旧块离开管道后到达，所以这是一个安全的操作。注意解析操作只会在缓存中追加数据，不会修改。所以在校验通道里可以进行多线程的无锁操作。在最差的场景中，许多线程去读取同一个未被解析的数据，所有的程序并发的执行解析操作。接着最后写入缓存的线程获得胜利，这是没有问题的，因为大家执行的结果都是一致的。

调用图标分析，即使进行了这些操作，由于解析操作，内存占用率在执行期间仍然非常高。这是继gRPC调用和加密计算之后的占据最大的一项操作。但是后面2个操作，不在本次工作范围内。

## 4. 结果

本节介绍了对我们的体系结构改进的实验性性能评估。我们使用了15台本地服务器，通过1 Gbit/s交换机连接。每台服务器配备两个2.10GHz的Intel R Xeon R CPU E5-2620 v2处理器，总共24个硬件线程和64 GB RAM。我们使用Fabric1.2作为基本情况，并逐步添加我们的改进以进行比较。默认情况下，fabric配置为使用leveldb作为对等状态数据库，排序服务将已完成的块存储在内存中，而不是磁盘上。此外，我们不使用Docker容器来运行整个系统，以避免额外的开销。

虽然我们确保我们的实现不会更改fabric的验证行为，但所有测试都是使用不冲突且有效的交易运行的。这是因为有效交易必须经过每个验证检查，并且它们的写集将在提交期间应用于状态数据库。相反，可以删除无效的交易。因此，我们的结果评估了最坏情况下的性能。
![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-21%E4%B8%8A%E5%8D%8811.37.01.png)

对于专门针对排序者或提交者实验，我们分离了各自的系统部分。在order实验中，我们从客户机向order发送预加载的背书交易，并让一个模拟提交者简单地丢弃创建的块。类似地，在提交者的基准测试期间，我们将预加载的块发送给提交者，并为背书者和丢弃已验证块的块存储创建mock。

然后，对于端到端的设置，我们实现了完整的系统：背书者根据来自提交人的已验证块的复制世界状态从客户端背书交易建议；订购人从背书交易创建块并将它们发送给提交人；提交者验证并提交对其内存中世界状态的更改，并将已验证的块发送给背书者和块存储；块存储使用结构1.2数据管理将块存储在其文件系统中，状态存储在leveldb中。但是，我们没有为可伸缩分析实现分布式块存储；这超出了本工作的范围。

为了进行公平的比较，我们对所有实验使用了相同的交易链码：每个交易模拟从一个帐户到另一个帐户的资金转移，读取并更改状态数据库中的两个键。这些交易的有效负载为2.9kb，这是典型的[3]。此外，我们使用默认的背书政策，即接受单一背书人签名。

A. 通过GRPC传输数据块

我们首先对GRPC的性能进行基准测试。我们预先创建了包含不同交易数的有效块，通过Fabric GRPC接口将它们从Orderer发送到Peer，然后立即丢弃它们。实验结果如图3所示。

我们发现，对于从10到250个交易的块大小（这是在以下部分中导致最佳性能的大小），交易吞吐量超过40000个交易/秒是可持续的。与第IV-D节中的端到端测试结果相比，很明显，在我们的环境中，网络带宽和服务器机架中使用的1 Gbit/s交换机不是性能瓶颈。

![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-21%E4%B8%8A%E5%8D%8811.56.15.png)

B. 作为消息大小函数的订购方吞吐量

在这个实验中，我们设置了多个客户机来向订购者发送事务，并监视发送100000个事务所需的时间。我们评估订单在订单1.2中的交易率，并将其与我们的改进进行比较：

•opt o-i：仅向卡夫卡发布事务ID（第三-B节）

•选择O-II：来自客户的并行传入交易建议（第III-C节）

图4显示了不同负载大小的交易吞吐量。在Fabric1.2中，由于向Kafka发送大消息的开销，交易吞吐量随着负载大小的增加而降低。然而，当我们只将Turac ID发送给卡夫卡（OPT-O-1）时，对于4096 kb的有效负载大小，我们几乎可以将平均吞吐量（2.8×）增加三倍。添加优化o-2后，平均吞吐量比基础结构1.2提高4倍。特别是，对于2kb负载大小，我们将排序服务性能从6215个交易/秒提高到21719个交易/秒，比率接近3.5x。

C. Peer实验

在这一节中，我们描述了在孤立的单个Peer节点上的测试（我们在第IV-D节中展示了端到端评估的结果）。在这里，我们预先计算块并将它们发送给Peer，就像我们在第IV-A节的GRPC实验中所做的那样。然后Peer完全验证并提交块。

与结构1.2相比，图中所示的三种配置累计包含了我们的改进（即，opt p-i i包含opt p-i，opt p-iii包含了先前的两种改进）：

•opt p-i leveldb替换为内存哈希表
•opt p-ii验证和提交完全并行化；块存储和背书分离，通过远程GRPC传输，分离存储服务器
•opt p-iii所有解析的数据可以在校验和提交管道访问

![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-21%E4%B8%8B%E5%8D%8812.49.19.png)

![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-21%E4%B8%8B%E5%8D%8812.49.27.png)

1）固定块大小的实验：图5和图6显示了一次运行100000个交易的验证和提交结果，重复1000次。交易是收集到100个交易块中。我们先讨论延迟，然后是吞吐量。由于批处理，我们显示每个块的延迟，而不是每个交易延迟。结果与我们自己设定的目标一致，即不因吞吐量增加而引入额外的延迟；事实上，我们的性能改进将对等延迟减少到原始值的三分之一（请注意，这些实验没有考虑网络延迟）。虽然opt p-ii中引入的管道生成了一些额外的延迟，但是其他的优化并不能补偿它。

通过使用状态存储哈希表（opt p-i），我们能够将Fabric1.2对等机的吞吐量从3200个交易/秒增加到7500多个交易/秒。并行化Val-idation（optp-ii）每秒增加了大约2000个交易操作。这是因为，如图2所示，只有前两个验证步骤可以并行化和缩放。因此，整个管道性能取决于读写集验证和提交的吞吐量。虽然使用opt p-i时承诺几乎是免费的，但是直到在opt p-iii中引入解组缓存，optp-ii才有了回报。高速缓存大大减少了cpu的工作量，释放了并行验证额外块的资源。将所有Peer优化结合在一起，我们将Peer的提交性能提高了7倍，从大约3200个交易/秒提高到超过21000个交易/秒。

6.我们实验确定在该块大小上Peer吞吐量最大化。

![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-21%E4%B8%8B%E5%8D%881.07.25.png)

![img](http://img.mochain.info/topjohn/blog/post/%E6%88%AA%E5%B1%8F2019-09-21%E4%B8%8B%E5%8D%881.08.06.png)

2）参数敏感性：如第IV-C节所述，在Peer并行化块和交易验证至关重要。但是，不清楚要使性能最大化需要多少并行性。因此，我们探索一个Peer的性能可以通过改变两个参数来调谐的程度：

•验证管道中同时引导块的go例程的数量

•同时验证交易处理的go例程的数量

我们使用信号量控制系统中活动go协程的数量，同时允许多个块同时进入验证管道。这允许我们通过两个独立的go例程池来控制块头验证和交易验证中的并行级别。

对于100个交易的块大小，图7显示了改变go例程数量时的吞吐量。验证管道中的线程总数由两个独立轴的总和给出。例如，我们为管道中的25个交易验证go例程和31个并发块实现了最大吞吐量，总共为管道提供了56个go协程。当有太多线程时，我们会看到通过线程管理开销导致的性能小幅度下降，但是用太少的并行执行来耗尽cpu的代价是巨大的。因此，我们建议默认情况下，在给定的机器中，go例程的数量至少是物理线程的两倍。

我们现在研究Peer吞吐量对块大小的依赖性。每个块大小实验是用先前测试的最佳调谐GO例程参数进行的。在24±2交易验证中使用的所有配置执行常规程序，并在管道中进行30±3个阻塞。同样，我们在给定大小的块中为一个基准测试运行分割100000个交易，并重复实验1000次。我们选择在对数尺度上扫描块大小空间以获得宽光谱的概述。

结果如图8所示。我们发现，块大小为100个交易/块时，以每秒21000多个交易的速度提供最佳吞吐量。我们还研究了与这个块大小的小偏差。我们发现块大小在50到500之间的性能差异非常小，因此我们选择将块大小固定为100个交易。

D.端到端吞吐量

我们现在讨论通过组合所有优化（即opt）实现的端到端吞吐量。o-ii与opt相结合。p-iii，与我们对未改性织物1.2的测量结果相比。

我们设置了一个使用三个zookeeper服务器和三个kafka服务器（默认主题复制因子为三）的集群的单个排序节点，并将其连接到Peer。来自此Peer的块被发送到单个数据存储服务器，该服务器将世界状态存储在leveldb中，并将块存储在文件系统中。对于扩展，五个背书者复制对等状态并提供足够的吞吐量来处理客户端背书负载。最后，客户机安装在自己的服务器上；该客户机从五个背书服务器请求背书，并将背书事务发送到排序服务。这总共使用15台服务器连接到本地数据中心的同一个1 Gbit/s交换机。我们从客户端向排序方发送总计100000个已背书的交易，排序者将这些交易批处理为100个大小的块，并将它们传递给Peer。为了估计吞吐量，我们测量Peer上提交的块之间的时间，并取一次运行的平均值。这些运行重复100次。表1显示，与我们的基线结构1.2基准相比，显著提高了6-7倍。

## 5. 相关工作

hyperledger fabric是一个最近才开发的系统，它的架构仍在快速发展和重大变化中。因此，对于系统的性能分析或架构改进的建议方面的工作相对较少。在这里，我们综述了提高Fabric性能的最新技术。

最接近我们的工作是由thakkar等人[6]谁研究了各种配置参数对Fabric性能的影响。他们发现，主要的瓶颈是在背书策略验证期间重复验证x.509证书，对块中的交易进行顺序策略验证，以及在提交阶段进行状态验证。它们引入了对已验证的认可证书的积极缓存（合并到Fabric版本1.1中，因此是我们评估的一部分）、认可策略的并行验证以及批处理状态验证和承诺。这些改进使总吞吐量增加了16倍。我们还对提交方的验证进行了并行化，并进一步将状态数据库替换为更有效的数据结构，即哈希表。
hyperledger fabric是一个最近才开发的系统，它的架构仍在快速发展和重大变化中。因此，对于系统的性能分析或架构改进的建议方面的工作相对较少。在这里，我们综述了提高Fabric性能的最新技术。

在最近的工作中，sharma等人[14]研究了使用数据库技术，即事务重新排序和提前中止，来提高fabric的性能。他们关于早期识别冲突交易的一些想法与我们的想法是正交的，可以纳入我们的解决方案。但是，有些想法，例如让排序服务删除冲突交易，与我们的解决方案不兼容。首先，我们故意不向排序者发送交易读写集，只发送交易id。其次，我们选择遵守fabric的设计目标，即将不同的任务分配给不同类型的节点，因此我们的排序服务不检查读写集的内容。未来工作的一个有趣方向是在不同的交易工作负载下比较这两种方法，以了解何时向排序服务发送完整交易详细信息的开销值得提前修剪冲突交易。

众所周知，由于拜占庭容错协议（bft）的消息通信开销，fabric的order组件可能成为性能瓶颈。因此，在排序服务中使用bft协议的有效实现非常重要。sousa等人[3]研究了著名的bft-smart[15]实现作为fabric的一部分的使用，并表明，在单个数据中心内使用此实现，可以实现高达30000个交易/秒的吞吐量。然而，与我们的工作不同，Committer组件没有基准测试，端到端的性能也没有得到解决。

Androulaki等人[16]研究了通道在Fabric上的应用。然而，这项工作并没有提出一个绩效评估，以定量地确定其方法的效益。

raman等人[17]研究了当区块链用于存储大型数据集分析产生的中间结果时，使用有损压缩来降低结构背书人和验证人之间共享状态的通信成本。然而，他们的方法只适用于对有损压缩不敏感的场景，这不是基于区块链的应用程序的一般情况。

一些研究已经检查了Fabric的性能，但没有提出内部结构的变化。例如，Dinh等人使用Blockbench[5]这一工具来研究私有区块链的性能，研究Fabric的性能，并将其与以太坊和奇偶校验的性能进行比较。他们发现，由于消息通道中的聚集，他们研究的结构版本没有扩展到超过16个节点。nasir等人[18]比较了fabric 0.6和1.0的性能，发现1.0版本的性能优于0.6版本，这并不奇怪。Baliga等人[19]表明，应用程序级参数（如交易的读写集大小、链码和事件负载大小）显著影响交易延迟。类似地，Pongnumkul等人[20]比较了Fabric和以太坊对于加密货币工作负载的性能，发现Fabric在所有指标上都优于以太坊。bergman[21]将fabric与apache cassandra在类似环境中的性能进行了比较，发现对于少量Peer节点，fabric在读重工作负载中的可线性化交易的延迟比cassandra低。另一方面，由于节点数量较多，或者写的工作负载很重，cassandra有更好的性能。

## 6. 总结

这项工作的主要贡献是展示如何对每一个错误的BuffStand框架，如HyffeDiger-Frand，可以重新设计，以支持每秒近20000个交易，一个比现有工作好7的因素。我们通过实现一系列独立的优化来实现这个目标，这些优化集中在I/O、缓存、并行性和高效的数据访问上。在我们的设计中，排序服务只接收交易id而不是完整的交易，并且Peer上的验证是高度并行化的。我们还使用主动缓存，并利用轻量级数据结构在关键路径上快速访问数据。在未来的工作中，我们希望通过以下方式进一步提高hyperledger fabric的性能：
•结合有效的bft一致性算法，如rcanopus[22]
•在不打开整个交易头的情况下，加快为排序服务提取交易ID•替换现有的加密计算库，提供更有效率的库
•通过分配一个单独的每个通道的排序和快速Peer服务器
•使用分布式框架，如apache spark[23]

## 引用

[1] V. Espinel, D. O’Halloran, E. Brynjolfsson, and D. O’Sullivan, “Deep shift, technology tipping points and societal impact,” in New York: World Economic Forum–Global Agenda Council on the Future of Software & Society (REF 310815), 2015.
[2] M. Vukolic ́, “The quest for scalable blockchain fabric: Proof-of-work vs. bft replication,” in International Workshop on Open Problems in Network Security. Springer, 2015, pp. 112–125.
[3] J. Sousa, A. Bessani, and M. Vukolic, “A Byzantine fault-tolerant ordering service for the hyperledger fabric blockchain platform,” in 2018 48th Annual IEEE/IFIP International Conference on Dependable Systems and Networks (DSN). IEEE, 2018, pp. 51–58.
[4] M. Yin, D. Malkhi, M. K. Reiter, G. Golan Gueta, and I. Abraham, “HotStuff: BFT Consensus in the Lens of Blockchain,” arXiv preprint arXiv:1803.05069, 2018.
[5] T. T. A. Dinh, J. Wang, G. Chen, R. Liu, B. C. Ooi, and K.-L. Tan, “BLOCKBENCH: A Framework for Analyzing Private Blockchains,” Proceedings of the 2017 ACM International Conference on Management of Data - SIGMOD ’17, pp. 1085–1100, 2017.
[6] P. Thakkar, S. Nathan, and B. Vishwanathan, “Performance Benchmarking and Optimizing Hyperledger Fabric Blockchain Platform,” arXiv, 2018.
[7] Hyperledger Fabric, “[FAB-12221] Validator/Committer refactor - Hyperledger JIRA.” [Online]. Available: https://jira.hyperledger.org/ browse/FAB- 12221?filter=12526
[8] C. Cachin, “Architecture of the hyperledger blockchain fabric,” in Workshop on Distributed Cryptocurrencies and Consensus Ledgers, vol. 310, 2016.
[9] E. Androulaki, A. Barger, V. Bortnikov, C. Cachin, K. Christidis, A. De Caro, D. Enyeart, C. Ferris, G. Laventman, Y. Manevich, S. Muralidharan, C. Murthy, B. Nguyen, M. Sethi, G. Singh, K. Smith, A. Sorniotti, C. Stathakopoulou, M. Vukolic ́, S. W. Cocco, and J. Yellick, “Hyperledger Fabric: A Distributed Operating System for Permissioned Blockchains,” Proceedings of the Thirteenth EuroSys Conference on - EuroSys ’18, pp. 1–15, 2018.
[10] ApacheFoundation,“ApacheKafka,ADistributedStreamingPlatform,” 2018. [Online]. Available: https://kafka.apache.org/ (Accessed 2018-12- 05).
[11] Cloud Native Computing Foundation, “gRPC: A high performance, open-source universal RPC framework,” 2018. [Online]. Available: https://grpc.io/
[12] S. Bano, A. Sonnino, M. Al-Bassam, S. Azouvi, P. McCorry, S. Meik- lejohn, and G. Danezis, “Consensus in the age of blockchains,” arXiv preprint arXiv:1711.03936, 2017.
[13] Google Developers, “Protocol Buffers — Google Developers,” 2018. [Online]. Available: https://developers.google.com/protocol-buffers/?hl= en
[14] A. Sharma, F. M. Schuhknecht, D. Agrawal, and J. Dittrich, “How to databasify a blockchain: the case of hyperledger fabric,” arXiv preprint arXiv:1810.13177, 2018.
[15] A. Bessani, J. Sousa, and E. Alchieri, “State machine replication for the masses with BFT-SMART,” in DSN, 2014, pp. 355–362.
[16] E. Androulaki, C. Cachin, A. De Caro, and E. Kokoris-Kogias, “Channels: Horizontal scaling and confidentiality on permissioned blockchains,” in European Symposium on Research in Computer Se-
curity. Springer, 2018, pp. 111–131.
[17] R. K. Raman, R. Vaculin, M. Hind, S. L. Remy, E. K. Pissadaki, N. K. Bore, R. Daneshvar, B. Srivastava, and K. R. Varshney, “Trusted multi- party computation and verifiable simulations: A scalable blockchain approach,” arXiv preprint arXiv:1809.08438, 2018.
[18] Q. Nasir, I. A. Qasse, M. Abu Talib, and A. B. Nassif, “Performance analysis of hyperledger fabric platforms,” Security and Communication Networks, vol. 2018, 2018.
[19] A. Baliga, N. Solanki, S. Verekar, A. Pednekar, P. Kamat, and S. Chat- terjee, “Performance Characterization of Hyperledger Fabric,” in Crypto Valley Conference on Blockchain Technology, CVCBT 2018, 2018.
[20] S. Pongnumkul, C. Siripanpornchana, and S. Thajchayapong, “Performance analysis of private blockchain platforms in varying workloads,” in 2017 26th International Conference on Computer Communications and Networks, ICCCN 2017. IEEE, 7 2017, pp. 1–6.
[21] S. Bergman, “Permissioned blockchains and distributed databases: A performance study,” Master’s thesis, Linkoping University, 2018.
[22] S. Keshav, W. M. Golab, B. Wong, S. Rizvi, and S. Gorbunov, “RCano- pus: Making canopus resilient to failures and byzantine faults,” arXiv preprint arXiv:1810.09300, 2018.
[23] Apache Foundation, “Apache Spark - Unified Analytics Engine for Big Data,” 2018. [Online]. Available: [https://spark.apache.org](https://spark.apache.org/)

# 企业级区块链对比：Quorum vs. Corda vs. Hyperledger Fabric

2019-08-17 19:41:02 阅读：95 来源： **互联网**



**标签：**[Fabric](https://www.icode9.com/tags-Fabric-0.html) [以太](https://www.icode9.com/tags-以太-0.html) [Corda](https://www.icode9.com/tags-Corda-0.html) [企业级](https://www.icode9.com/tags-企业级-0.html) [vs](https://www.icode9.com/tags-vs-0.html) [Hyperledger](https://www.icode9.com/tags-Hyperledger-0.html) [区块](https://www.icode9.com/tags-区块-0.html) [Quorum](https://www.icode9.com/tags-Quorum-0.html)



企业分布账本技术（Distributed Ledger Technology）需要解决5个方面的挑战：数据隐私性、技术正当性、可伸缩性、最终一致性和互操作性。本文将对企业以太坊（Quorum）、Hyperledger Fabric和Corda就这些环节进行比较。

![在这里插入图片描述](https://www.icode9.com/img/ll/?i=2019081719242658.png?,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NoZWJhbzMzMzM=,size_16,color_FFFFFF,t_70)

> 如果希望快速掌握区块链应用的开发，推荐汇智网的[**区块链应用开发系列教程**](http://www.hubwiz.com/course/?type=区块链&affid=csdn7878)， 内容涵盖比特币、以太坊、eos、超级账本fabric和tendermint等多种区块链，以及 java、go、c#、nodejs、python、php、dart等多种开发语言。

## 1、数据隐私性

应该说绝大多数机构都不会愿意与市场中的对手分享其竞争优势。如果竞争对手可以访问区块链上的企业私有数据，那么相信绝大多数企业都不会选择这种区块链技术。因此我们要对比的这几种企业区块链都采用了许可机制，以确保只有许可的节点才能够加入网络。

### 1.1 Quorum

Quorum是企业级的以太坊实现，它将世界状态区分为公共和私有两种，并且使用加密的P2P消息通信。

虽然Quorum提供了私有数据传输的选项，但是并没有考虑人为、商业逻辑或代码错误所带来的潜在的数据泄露风险。此外，Quorum依赖于使用zk-SNARKS实现的零知识证明安全层（ZSL），而这一密码学技术相对而言还缺少足够的测试，并且不能对抗量子计算，这是值得注意的一点 ，当然对于某些机构而言没太大影响，但是对于全国性的金融机构而言，就需要考虑量子计算的问题了。

### 1.2 Hyperledger Fabric

超级账本Fabric项目提出了**通道**的概念，用来保护数据的隐私。应当指出的是，随着网络的扩大，管理通道的创建和维护会变得越来越困难。

Hyperledger Fabric也提供了zk-SNARKS支持。通道和零知识数据证据的结合的确能够支持一个相对严谨的隐私模型，但是也带来了开发和处理方面的高成本。

由于固有的隐私设计方面的问题，Hyperledger Fabric也已经进行了若干次重构，因此当需要考虑你的应用的后向兼容性时，别忘了这一点。

### 1.3 Corda

Corda使用点对点的交易模型，天生支持隐私交易。数据仅暴露给需要了解的人，例如交易目标以及notary集群。

一个点对点的模型，在数据篡改方面会有更大的风险。Corda的威胁模型使用了Notary共识，Notary将拒绝被篡改资产的状态更新。一个攻击者需要勾结所有的相关参与方，同时也需要控制特定的notary基础架构。

## 2、技术正当性

企业分布账本技术应当尽量避免技术方面的摩擦，使用经过考验的技术，并考虑协议或开发语言的未来可能的变化。

### 2.1 quorum

Quorum使用Solidity进行智能合约开发，这是一把双刃剑。对于公有链来说，EVM和Solidity无疑是去中心化应用的最流行的选择，这意味着从公链到私有以太坊区块链的技能共享是一个独特的优势。

但是，基于个人经验以及众多声名狼藉的事件，例如DAO攻击、Parity多签钱包攻击等等，Solidity程序错误导致的后果可能是灾难性的。

应当说以太坊出色的核心开发团队已经意识到了这一点，并且正在使用eWASM来替换EVM。我希望Quorum也能跟随这一方向。

### 2.2 Hyperledger Fabric

Hyperledger Fabric中的智能合约被成为链码，支持Java、NodeJS、Golang和Python。有趣的是，Hyperledger Fabric支持以太坊，我猜测是为了在短期内利用开发者的EVM区块链上的开发技能。

### 2.3 Corda

Corda依赖于广泛经过考验的技术：Java（或Kotlin）、AMQP和SQL数据库。

Corda的智能合约使用Corda的流引擎开发，该产品正在进行重大升级，以便支持开发者能够聚焦于核心业务逻辑，并提供相对较短的学习曲线。

Corda的交易依赖于企业中的传统架构，在交易中可以包含法律条款（李嘉图合约）。在Corda区块链上，代码不是法律，法律才是法律。

## 3、可伸缩性

可伸缩性是区块链和分布账本技术的圣杯，在某些情况下也是大规模应用区块链的瓶颈所在。对于企业而言，可伸缩性和交易吞吐量（TPS）是需要考虑的重要因素。

可伸缩性三元悖论指的是，当可伸缩性提升时，去中心化或安全则会受损。企业分布账本技术已经找到了创新的技术来解决这个问题。

### 3.1 Quorum

Quorum通常可以达到100-200TPS的吞吐量，不过这一数据依赖于网络设计、硬件资源等因素。

### 3.2 Hyperledger Fabric

据报道，超级账本Fabric可以达到2000~20000 TPS的吞吐量，这令人印象非常深刻，已经达到企业级应用的水准了。

### 3.3 Corda

Corda企业版提供了多线程能力，可以非常显著的提升交易吞吐量。DTCC最近证实在170个云节点组成的区块链上达到了6300TPS的吞吐量，这一指标大约等价于每天处理1.15亿个交易。

## 4、最终一致性

无可争议的一点是，像比特币或以太坊这样的公链所使用的POW共识，不适合企业采用，在POW共识中，一旦区块提交到账本，随着链的增长，区块被反转的几率指数性减少，但是永远也达不到数学上的最终一致性 —— 你知道在公链上存在着51%攻击。

最终一致性基本上是二元的，一个交易要么被视为完成，要么没有。上面的三个项目都提供了实现最终一致性的共识算法。因此三者都通过了这一环节的测试。

## 5、互操作性

每个人都理解互联网有改变世界的潜力，但是这一点并不是立刻变得很清晰。互联网实际上是一些分立的网络，使用不同的编程语言，也没有标准的协议设计。

知道区块链有了它的TCP/IP时刻，并且达成了标准 —— Polkadot和Cosmos看起来有希望，还有
Jasper/Ubin的先驱性的工作 —— 跨链的互操作性现在看来，还很遥远。

因此为了便于比较，让我们对比同质区块链的互操作性，在同一个DLT网络中的应用如何进行互操作。

### 5.1 Quorum

企业以太坊，在企业以太坊联合会的指导下，所提出的客户端规范V3，是走在了正确的方向上。目前还没有谈到何时由谁来采纳这些标准，虽然我猜测EEA的成员将朝着这一方向前进。

Solidity允许合约彼此调用，这可能会带来重入的安全问题，因此应当小心对待。

### 5.2 Hyperledger Fabric

虽然通道能够提升速度并且显式地支持隐私，它同时也有一些不容忽视的缺点。

例如，在通道之间缺乏互操作性。虽然也有哈希时间锁定合约以及哈希时间锁协议这样的解决方案，但是公平的说一个通道里的资产还是需要在其他通道中重新发行。

就我个人认为（以及Gendal Brown，Corda的CTO），这是在设计缺陷上打补丁的做法。

### 5.3 Corda

Corda的核心原则之一就是互操作性，从一开始，Cord就是设计为可以在Corda应用间实现互操作性。Corda网络的涌现允许节点互联网来通过一个安全层交换资产或数据。

Corda Settler 甚至允许创建连接外部支付的模块。

![在这里插入图片描述](https://www.icode9.com/img/ll/?i=20190817192711753.png?,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3NoZWJhbzMzMzM=,size_16,color_FFFFFF,t_70)

## 6、结论

当选型企业区块链技术时，和其他任何软件一样，你必须考虑框架的各种优缺点，以及与你的应用的契合度。我希望本文能提供一些帮助。



# 饭后闲聊，Hyperledger Fabric性能测试Caliper

最近用Caliper测试了以下Hpyerledger Fabric的性能，测试环境用本地的byfn脚本

作为饭后闲谈，这里我们并不去探究如何提高tps，以及代码的角度为什么会有这样的区别。

话不多说，代码：

https://github.com/SamYuan1990/FabriccaliperSample

# Caliper性能很差么？

从观察到的结果来看，

如果我们仅仅用一个client发起交易的tx，那么其性能相对于https://github.com/guoger/stupid.git 直接通过GRPC发起请求的对比来看会低，

那么，真的是caliper用nodejs，以及fabric nodejs sdk的性能很低么？

我们增加client的数量，可以看到Caliper也可以达到200TPS的。

| Item              | TPS   |
| ----------------- | ----- |
| Caliper 1 client  | 56.6  |
| stupid            | 231   |
| Caliper 10 client | 194.3 |

# Client可以无限增加么？

我们把client增加到多个看一下结果

| Item              | TPS   |
| ----------------- | ----- |
| Caliper 10 client | 194.3 |
| Caliper 15 client | 191.4 |
| Caliper 20 client | 176.8 |
| Caliper 30 client | 130.1 |

我们可以看到随着client的增加，tps并不是无限增高的，Avg Latency不断提高。https://github.com/SamYuan1990/FabriccaliperSample/blob/master/result.md

 

# Calliper的rate？

目前我们都是用的fixed-rate, 那么fixed-feedback-rate，fixed-backlog中unfinished_per_client的选项呢？

从试验结果来看，个人感觉是unfinished_per_client是对于性能测试进行微调。正如https://hyperledger.github.io/caliper/vLatest/rate-controllers/中描述的那样

“100 unfinished transactions for each client”

允许每个client有多少个未完成的请求的样子。如果这个值小于send rate，那么发送的tps就会低于send rate。如果大于等于，那么这里tps就会略高于未设置的fixed-rate。

 

# 补充一个结论

使用cliaper做性能测试时，需要根据测试结果，谨慎的优化测试方案。来得到相对好的性能。

## fabric 网络性能测试

#### 项目介绍

```
主要测试fabric吞吐量和并发，基于fabric-sdk-go，测试工具使用wrk
```

#### 项目依赖

```
- fabric-sdk-go （它本身有很多依赖）
- git clone https://github.com/wg/wrk.git
```

#### 安装

```
go get -u learnergo/fabric-performance-test

cd $GOPATH/src/github.com/learnergo/fabric-performance-test
```

#### chaincode

测试链码（官网），主要做了存(put)取(get)操作，存的过程加入了加解密操作增加复杂度，cli 操纵示例：

```
peer chaincode query -C mychannel -n mycc -c '{"Args":["get","a"]}'
peer chaincode invoke -o orderer.example.com:7050  --tls true --cafile $ORDERER_CA -C mychannel -n mycc -c '{"Args":["put","a",b"]}'
```

#### 实现思路

```
因为主要侧重吞吐和并发测试，对通道和链码安装部分不是很侧重。
本项目fabric网络有mychannel通道和testcc链码（名字可以自己确定，并在程序和配置中对应修改）。
在用命令行创建链码时，先存入一个值对（"a":"b"）,取的测试是取a值；存的测试是存入当前时间戳。

**特别提示**：
为了避免日志打印对性能影响，只打印了error日志。运行正常的标志也就是没有日志打印
```

#### 操作步骤

- 配置fixtures下证书密钥和配置文(只配置一个peer和orderer即可)
- 运行main.go
- 在新窗口用wrk进行测试（调整-t 和-c 值即可，-d 越大越准确）

#### 测试环境

```
多机
Linux VM-0-17-ubuntu 4.4.0-91-generic #114-Ubuntu SMP Tue Aug 8 11:56:56 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
```

#### 测试结果

读写并发都在1000以上，读的tps在400以上，写在100以上

本人tps最佳参数：

```
./wrk -t4 -c150 -d10 --timeout 10 http://localhost:8026/v1/gettest

./wrk -t4 -c100 -d10 --timeout 30 http://localhost:8080/puttest
```

#### 影响因素

```
- 节点数量
- 服务器配置（cpu 内存 网络等等）
- 日志级别（级别越低性能越低）
- 是否启用tls（不启用tps高）
- solo or kafka （solo 高）
- leveldb couchdb选择（leveldb 高）
- orderer 出块配置（自己研究吧）
```

### [learnergo](http://github.com/learnergo)/[fabric-performance-test](http://github.com/learnergo/fabric-performance-test)

[18](http://github.com/learnergo/fabric-performance-test/watchers)[13](http://github.com/learnergo/fabric-performance-test/network/members)

基于go sdk 的性能测试工具 — [Read Mo](http://github.com/learnergo/fabric-performance-test#readme)



# 测试技术之Hyperledger fabric 性能压测

![img](http://iotekimg.zhizuobiao.com/zzb/190614111445/023B.jpg)白羽2018-06-21来源 ：网络阅读 1877评论 0

摘要：本文将带你了解测试技术之Hyperledger fabric 性能压测，希望对大家学测试技术有所帮助。







　　基础知识储备已达到一定阶段，所以开始投入生产研发，所以就对fabric进行了性能压测；

　　项目使用的是fabric-sdk-java，所以这里的压测也只java版本的结果；

　　部署情况：

```
　10.0.200.111机器：
　　    |-orderer1.lychee.com
　　    |-ca0.org1.lychee.com
　　10.0.200.113机器：
　　    |-orderer2.lychee.com
　　    |-peer0.org1.lychee.com
　　    |-z1   (zookeeper)
　　    |-k1   (kafka)
　　    |-couchdb0.org1.lychee.com
　　10.0.200.114机器：
　　    |-orderer3.lychee.com
　　    |-peer1.org1.lychee.com
　　    |-z2
　　    |-k2
　　    |-couchdb1.org1.lychee.com
　　10.0.200.115机器：
　　    |-peer0.org2.lychee.com
　　    |-z3
　　    |-k3
　　    |-couchdb0.org2.lychee.com
　　10.0.200.116机器：
　　    |-peer1.org2.lychee.com
　　    |-k4
　　    |-ca0.org2.lychee.com
　　    |-couchdb1.org2.lychee.com
   
　　机器配置如下：
　　2核4G内存200G硬盘
　　sdk虚拟处的channel只加入了2个peer，具体结果如下：
　　1、网络搭建过程，即channel创建过程
　　      10s左右
　　2、查询耗时
　　      同步使用10个线程，每个线程进行10个查询，总耗时在2s左右
　　      平均每个查询耗时在60～400毫秒之间
　　3、invoke相关，本案例测试的智能合约包含了getstate和putstate操作
　　      同步使用10个线程，每个线程进行10个查询，总耗时在6s左右
　　      平均每个查询耗时在60～400毫秒之间
   
　　总结：
　　1、client使用单例模式
　　2、查询接口在50TPS
　　3、数据处理接口在16TPS
   
　　附java性能测试类：
　　package com.lychee.fabric.sdk.demo;
　　import java.util.concurrent.CountDownLatch;
　　public class TestHarness {
　　       public static void main(String[] args) throws InterruptedException {  
　　            TestHarness testHarness = new TestHarness();  
　　            long timeTasks = testHarness.timeTasks(10, new Runnable() {  
　　                @Override  
　　                public void run() {  
　　                    try {  
　　                        Thread.sleep(1000);  
　　                    } catch (InterruptedException e) {  
　　                        e.printStackTrace();  
　　                    }  
　　                }  
　　            });  
　　            System.out.println(timeTasks);  
　　        }  
　　        public long timeTasks(int nThreads, final Runnable task) throws InterruptedException {  
　　//          //预热，编译  
　　//          for (int i = 0; i < 10000; i++) {  
　　//              task.run();  
　　//          }          
　　            // 真正的测试  
　　            // 使用同步工具类，保证多个线程同时（近似同时）执行
　　            final CountDownLatch startGate = new CountDownLatch(1);  
　　            // 使用同步工具类，用于等待所有线程都运行结束时，再统计耗时
　　            final CountDownLatch endGate = new CountDownLatch(nThreads);  
　　            for (int i = 0; i < nThreads; i++) {  
　　                final int j = i;
　　                Thread t = new Thread() {  
　　                    @Override  
　　                    public void run() {  
　　                        try {  
　　                            startGate.await();  
　　                            try {  
　　                                
　　                                task.run();  
　　                            } finally {  
　　                                endGate.countDown();  
　　                            }  
　　                        } catch (InterruptedException e) {  
　　                            e.printStackTrace();  
　　                        }  
　　                    }  
　　                };  
　　                t.start();  
　　            }  
　　            long start = System.currentTimeMillis();  
　　            startGate.countDown();  
　　            endGate.await();  
　　            long end = System.currentTimeMillis();  
　　            return end - start;  
　　        } 
　　}
```

  本文由职坐标整理并发布，希望对同学们有所帮助。了解更多详情请关注职坐标常用软件之Maya频道！



# 以docker启动fabric网络，高并发大规模数据插入账本时，容器磁盘占用率急速升高



# Hyperledger Caliper测试Hyperledger Fabric并用Prometheus Grafana监控

