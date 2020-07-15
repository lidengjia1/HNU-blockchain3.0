# GO框架 Etcd-Raft模块

### 一. 写在前面:

- Etcd Version "3.4.7"

> Etcd is a strongly consistent, distributed key-value store that provides a reliable way to store data that needs to be accessed by a distributed system or cluster of machines.

Etcd是现在强一致性的分布式KV存储中最可信赖的一个,它建立在Raft共识理论之上,可谓风头无两. K8s用它作为集群配置中心,Hyperledger  Fabric也用它(只用其中Raft)作为排序节点一致性算法.

- Etcd从Github拉下来挺久的了,一直怯于它的复杂性.难得有个的假期,就想着系统性的看看.发现这个坑果然很大,4天时间远远不够,只能把Raft模块大致看一下吧.幸好网络上各种分析文章也很多,难以理解之处总能找到有用的解释.

### 二. rafpb包与gogoprotobuf

Raft需要通讯协议达到共识,通讯中传递的Message又是通过protobuf定义的.

1. gogoprotobuf补课

[https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/](https://links.jianshu.com/go?to=https%3A%2F%2Fcolobu.com%2F2019%2F10%2F03%2Fprotobuf-ultimate-tutorial-in-go%2F)
 [https://my.oschina.net/alexstocks/blog/387031](https://links.jianshu.com/go?to=https%3A%2F%2Fmy.oschina.net%2Falexstocks%2Fblog%2F387031)
 [https://my.oschina.net/alexstocks/blog/387058](https://links.jianshu.com/go?to=https%3A%2F%2Fmy.oschina.net%2Falexstocks%2Fblog%2F387058)

在raft.proto开始部分,



```proto
option (gogoproto.marshaler_all) = true;
option (gogoproto.sizer_all) = true;
option (gogoproto.unmarshaler_all) = true;
option (gogoproto.goproto_getters_all) = false;
option (gogoproto.goproto_enum_prefix_all) = false;
```

定义了Marshal()、MarshalTo(),Unmarshal()一组序列化方法,也不会生成GetXXXX方法,枚举变量名也不会包含前缀.

1. raft.proto

包含了Entry(记录到log和storage)机构,Message(事件消息通知),Snapshot(快照),HardState(要记log的状态),ConfState(集群中节点的分组状态),ConfChange(节点状态变化,还分了2个版本)

可以看出raft模块只定义了需要传递和记录的Message,但是不定义RPC调用,这里没有Service的定义.

1. confstate.go

给ConfState增加了Equivalent方法,用在newRaft和restore一个raft时候,其confstate也原来记录的一致.

1. confchange.go

定义了ConfChangeI,用来针对不同版本的ConfChange采用不同的Marshal()方法.

### 三. quorum包、tracker包、confstate包

参考代码,看下面大神的吧..
 由于版本变化，源码有些不同，etcd在tracker基础上又抽象出一个Changer的概念，由他来统一管理process,这里代码充斥着TODO,未来版本一定还要再改.

### 四. raft包

![img](https:////upload-images.jianshu.io/upload_images/5384456-1aece7ca47ff8cc7.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

核心结构数据流



具体的数据逻辑,看下面大神的吧..

### 各路大神文章:

#### raft动画

[http://thesecretlivesofdata.com/raft/](https://links.jianshu.com/go?to=http%3A%2F%2Fthesecretlivesofdata.com%2Fraft%2F)

#### etcd raft 设计与实现

[https://zhuanlan.zhihu.com/p/51063866](https://links.jianshu.com/go?to=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F51063866)
 [https://zhuanlan.zhihu.com/p/51065416](https://links.jianshu.com/go?to=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F51065416)

#### etcd Raft库解析

[https://www.codedump.info/post/20180922-etcd-raft/](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.codedump.info%2Fpost%2F20180922-etcd-raft%2F)

#### Go实现Raft

[https://mp.weixin.qq.com/s/5GJIGx7aeHpDdPCs4jyE2Q](https://links.jianshu.com/go?to=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F5GJIGx7aeHpDdPCs4jyE2Q)
 [https://mp.weixin.qq.com/s/zgLcBWuVzFsKkNngKW85Zw](https://links.jianshu.com/go?to=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FzgLcBWuVzFsKkNngKW85Zw)
 [https://mp.weixin.qq.com/s/FdAQlSsXCYOiBHKbgsGepQ](https://links.jianshu.com/go?to=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FFdAQlSsXCYOiBHKbgsGepQ)
 [https://mp.weixin.qq.com/s/azdjVpeTHkr4knajmdPtgw](https://links.jianshu.com/go?to=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FazdjVpeTHkr4knajmdPtgw)
 [https://github.com/eliben/raft](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Feliben%2Fraft)

#### etcd的raft实现 tracker quorum node log

[https://blog.csdn.net/weixin_42663840/article/details/100056484](https://links.jianshu.com/go?to=https%3A%2F%2Fblog.csdn.net%2Fweixin_42663840%2Farticle%2Fdetails%2F100056484)
 [https://blog.csdn.net/weixin_42663840/article/details/101039942](https://links.jianshu.com/go?to=https%3A%2F%2Fblog.csdn.net%2Fweixin_42663840%2Farticle%2Fdetails%2F101039942)
 [https://blog.csdn.net/weixin_42663840/article/details/100005978](https://links.jianshu.com/go?to=https%3A%2F%2Fblog.csdn.net%2Fweixin_42663840%2Farticle%2Fdetails%2F100005978)

#### 深入浅出 Raft - Membership Change

https://www.jianshu.com/p/99562bfec5c2

#### etcd raft如何实现成员变更

[https://zhuanlan.zhihu.com/p/27908888](https://links.jianshu.com/go?to=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F27908888)

#### Etcd raft源码阅读

[https://zhuanlan.zhihu.com/p/43282243](https://links.jianshu.com/go?to=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F43282243)

#### Etcd Raft架构设计和源码剖析

[https://lessisbetter.site/2019/08/19/etcd-raft-sources-arch/](https://links.jianshu.com/go?to=https%3A%2F%2Flessisbetter.site%2F2019%2F08%2F19%2Fetcd-raft-sources-arch%2F)
 [https://lessisbetter.site/2019/08/22/etcd-raft-source-data-flow/](https://links.jianshu.com/go?to=https%3A%2F%2Flessisbetter.site%2F2019%2F08%2F22%2Fetcd-raft-source-data-flow%2F)
 [https://lessisbetter.site/2019/09/05/etcd-raft-sources-structs/](https://links.jianshu.com/go?to=https%3A%2F%2Flessisbetter.site%2F2019%2F09%2F05%2Fetcd-raft-sources-structs%2F)

#### 官网Learner和翻译

[https://etcd.io/docs/v3.3.12/learning/learner/](https://links.jianshu.com/go?to=https%3A%2F%2Fetcd.io%2Fdocs%2Fv3.3.12%2Flearning%2Flearner%2F)
 [https://fuckcloudnative.io/posts/etcd-server-learner/](https://links.jianshu.com/go?to=https%3A%2F%2Ffuckcloudnative.io%2Fposts%2Fetcd-server-learner%2F)



作者：沉寂之舟
链接：https://www.jianshu.com/p/186fe51a3999
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。