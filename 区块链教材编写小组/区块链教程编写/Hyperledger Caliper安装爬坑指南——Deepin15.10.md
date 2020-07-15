# Hyperledger Caliper安装爬坑指南——Deepin15.10

全篇安装主要参照这篇官方文档，但中间有很多坑，需要格外注意。本人系统Deepin15.10

# 必要组件安装

请参加这位大大的blog,还有这位的,里面有详细记录如何安装相关组件，建议对照官方文档中的查看，因为blog里面的组件不全。并且，自己亲测按照blog中的安装方法无法适应最新版本，因此建议按照官方的来。注意docker-ce的版本要在18.09以后，否则会报错；还有就是一定要取消docker必须使用超级账户的权限，因为后续安装caliper不能用sudo，会报错。详情请参见：
最新docker下载请看 http://www.voidcn.com/article/p-zbjykblf-byx.html 需要注意如果要删除已安装的docker，尤其是版本装错了重装的，请使用

```bash
sudo apt-get remove docker-ce
```

否则会报错，说没有安装docker
安装完后改国内源和权限等后续处理请看 https://blog.csdn.net/qq_36148847/article/details/79273591

# caliper安装

由于最新版caliper的原因，很多2019年以前的教程中caliper的用法已经不适用了，因此还是得按照官方的来。
采用官方文档中 Local NPM install 的安装方法。 有两个坑需要注意
1)

```bash
npx caliper bind \
    --caliper-bind-sut fabric:1.4.0
```

这个语句现在已经不能用了，请用下面的语句代替

```bash
npx caliper bind \
    --caliper-bind-sut fabric --caliper-bind-sdk 1.4.0
```

2）

```bash
npx caliper launch master \
    --caliper-workspace . \
    --caliper-benchconfig benchmarks/scenario/simple/config.yaml \
    --caliper-networkconfig networks/fabric/fabric-v1.4.1/2org1peergoleveldb/fabric-go.yaml
```

这个语句现在也执行不了，现在按照

```bash
npx caliper benchmark run --caliper-workspace ./ --caliper-benchconfig benchmarks/scenario/simple/config.yaml --caliper-networkconfig networks/fabric/fabric-v1.4.1/2org1peergoleveldb/fabric-go.yaml
```

来执行了

上面2）中的语句实际就是执行测试的语句了，但是有个前提，必须在caliper官方的workbench目录下，该目录配置有很多测试用的文件

```git
git clone https://github.com/hyperledger/caliper-benchmarks.git
cd caliper-benchmarks
git checkout <your Caliper version>
```

选择和caliper相同的版本，如本文选择的是caliper v0.2.0，则这里也填 v0.2.0
一定要注意切换版本，否则可能出现测试的标准无法识别的情况。我一开始没有切版本，结果报了一个无法监控cpu的错误（其实我也不知道是不是这个原因导致的，反正切了应该没错吧）

# 使用caliper

大部分教程讲到测试成功就没有下文了，这里多讲一点关于怎么用的问题。首先，从开始执行的测试中可以看到测试的配置文件主要有两个。网络方面注意是 networks/fabric/fabric-v1.4.1/2org1peergoleveldb/fabric-go.yaml， 测试配置为benchmarks/scenario/simple/config.yaml

打开测试配置文件 benchmarks/scenario/simple/config.yaml，可以见到如下语句

```yaml
---
test:
  clients:
    type: local
    number: 1
  rounds:
  - label: Change car owner.
    txNumber:
    - 100
    rateControl:
    - type: fixed-rate
      opts:
        tps: 50
    arguments:
      assets: 1000
    callback: benchmarks/scenario/fabcar/changeCarOwner.js
  - label: Query all cars.
    txNumber:
    - 100
    rateControl:
    - type: fixed-rate
      opts:
        tps: 50
    arguments:
      assets: 1000
      startKey: '1'
      endKey: '50'
    callback: benchmarks/scenario/fabcar/queryAllCars.js
  - label: Query a car.
    txNumber:
    - 100
    rateControl:
    - type: fixed-rate
      opts:
        tps: 50
    arguments:
      assets: 1000
    callback: benchmarks/scenario/fabcar/queryCar.js
  - label: Create a car.
     txNumber:
    - 100
    rateControl:
    - type: fixed-rate
      opts:
        tps: 50
    arguments:
    callback: benchmarks/scenario/fabcar/createCar.js
monitor:
  type:
  - docker
  - process
  docker:
    name:
    - all
  process:
  - command: node
    arguments: local-client.js
    multiOutput: avg
  interval: 1
```

test里面round，每个label对应一个测试项，实例中测试的主要是吞吐量，调用callback里面的js文件进行测试。随意打开一个js文件，如示例sample中的open.js:

```javascript
'use strict';

module.exports.info  = 'opening accounts';

let account_array = [];
let txnPerBatch;
let initMoney;
let bc, contx;
module.exports.init = function(blockchain, context, args) {
    if(!args.hasOwnProperty('money')) {
        return Promise.reject(new Error('simple.open - \'money\' is missed in the arguments'));
    }

    if(!args.hasOwnProperty('txnPerBatch')) {
        args.txnPerBatch = 1;
    }
    initMoney = args.money;
    txnPerBatch = args.txnPerBatch;
    bc = blockchain;
    contx = context;

    return Promise.resolve();
};

const dic = 'abcdefghijklmnopqrstuvwxyz';
/**
 * Generate string by picking characters from dic variable
 * @param {*} number character to select
 * @returns {String} string generated based on @param number
 */
function get26Num(number){
    let result = '';
    while(number > 0) {
        result += dic.charAt(number % 26);
        number = parseInt(number/26);
    }
    return result;
}

let prefix;
/**
 * Generate unique account key for the transaction
 * @returns {String} account key
 */
function generateAccount() {
    // should be [a-z]{1,9}
    if(typeof prefix === 'undefined') {
        prefix = get26Num(process.pid);
    }
    return prefix + get26Num(account_array.length+1);
}

/**
 * Generates simple workload
 * @returns {Object} array of json objects
 */
function generateWorkload() {
    let workload = [];
    for(let i= 0; i < txnPerBatch; i++) {
        let acc_id = generateAccount();
        account_array.push(acc_id);

        if (bc.bcType === 'fabric') {
            workload.push({
                chaincodeFunction: 'open',
                chaincodeArguments: [acc_id, initMoney.toString()],
            });
        } else {
            workload.push({
                'verb': 'open',
                'account': acc_id,
                'money': initMoney
            });
        }
    }
    return workload;
}

module.exports.run = function() {
    let args = generateWorkload();
    return bc.invokeSmartContract(contx, 'simple', 'v0', args, 100);
};

module.exports.end = function() {
    return Promise.resolve();
};

module.exports.account_array = account_array;
```

关键在于module.export.run中，返回的bc.invokeSmartContract即调用链码功能，其中第二项为链码名，第三项为版本，第四项即参数，最后一项是啥我还不清楚，后面在研究。
注意参数的设置在generateWorkload里面，这是用户自定义的一个函数，但是其实也可以不这么写。再看看另外一个测试查询借口的代码query.js:

```javascript
'use strict';

module.exports.info  = 'querying accounts';


let bc, contx;
let account_array;

module.exports.init = function(blockchain, context, args) {
    const open = require('./open.js');
    bc       = blockchain;
    contx    = context;
    account_array = open.account_array;

    return Promise.resolve();
};

module.exports.run = function() {
    const acc  = account_array[Math.floor(Math.random()*(account_array.length))];

    if (bc.bcType === 'fabric') {
        let args = {
            chaincodeFunction: 'query',
            chaincodeArguments: [acc],
        };

        return bc.bcObj.querySmartContract(contx, 'simple', 'v0', args, 10);
    } else {
        // NOTE: the query API is not consistent with the invoke API
        return bc.queryState(contx, 'simple', 'v0', acc);
    }
};

module.exports.end = function() {
    // do nothing
    return Promise.resolve();
};
```

看起来似乎要短很多，原因在于这里没有随机生成用户名字的代码，转而从open.js的执行结果中取出了生成名字列表，注意开始init的部分。open.js的列表导出在最后一句

```javascript
module.exports.account_array = account_array;
```

这些东西我还没有摸索清楚，因为不是很会JavaScript，后面在慢慢弄。
还有一点，所有要测试的链码都要在通道和peer上安装，不然你是测试不了的。测试代码会自动安装，但是需要配置要安装的链码和链码的路径。配置文件在 networks/fabric/fabric-v1.4.1/2org1peergoleveldb/fabric-go.yaml里面。

```yaml
channels:
  mychannel:
    configBinary: networks/fabric/config_solo/mychannel.tx
    created: false
    orderers:
    - orderer.example.com
    peers:
      peer0.org1.example.com:
        eventSource: true
      peer0.org2.example.com:
        eventSource: true

    chaincodes:
#     - id: marbles
#       version: v0
#       language: golang
#       path: fabric/samples/marbles/go
#       metadataPath: src/fabric/samples/marbles/go/metadata
#     - id: drm
#       version: v0
#       language: golang
#       path: fabric/scenario/drm/go
    - id: simple
      version: v0
      language: golang
      path: fabric/scenario/simple/go
#     - id: smallbank
#       version: v0
#       language: golang
#       path: fabric/scenario/smallbank/go
    - id: abstore
      version: v0
      language: golang
      path: fabric/scenario/abstore/go
```

chaincodes下面就是配置要安装的链码的，注意这个path，是从你指定的workspace下src目录开始算的，即完整地址为 你的workspace/src/path

大致就这些了，我也就研究了这么多，后续有进展还会继续记录。



本文链接: http://www.luyixian.cn/news_show_312965.aspx   分享到： 腾讯 新浪 人人网 邮件 收藏夹 复制网址 [更多](http://www.jiathis.com/share/?uid=90225)

# Hyperledger Caliper - Fabric区块链性能测试工具

- ![img](https://msd.misuland.com/images/v2/auth.svg)xuxiang
-  

- ![img](https://msd.misuland.com/images/v2/level.svg)7
-  

- ![img](https://msd.misuland.com/images/v2/time.svg)2020-03-10 21:45

在这个教程中，我们将学习如何使用Hyperledger Caliper对包含多个排序节点的Fabric网络进行基准测试，我们使用Docker Swarm作为容器编排工具。

>  
>
> Hyperledger Fabric[区块链](http://msd.misuland.com/pd/3691885236884346656)开发教程：
>
> - Fabric区块链Node.js开发详解
> - Fabric区块链Java开发详解
> - Fabric区块链Golang开发详解

## 1、待测Fabric网络的基本配置

测试环境使用3台虚拟机，配置如下：

- Google cloud VM instance (n1-standard-4)
- 4 vCPUs, 15 GB memory
- Ubuntu 18.04.2 LTS

每台虚拟机需要提前安装以下软件：

- Docker version 18.09.8, build 0dd43dd87f or above
- docker-compose version 1.17.1 or above
- Node.js v8.16.0
- NPM 6.4.1

## 2、待测Fabric网络的拓扑结构

需要测试的Hyperledger Fabric网络中包含3个参与机构，共3个排序节点采用Raft共识[算法](http://msd.misuland.com/pd/3691885236884346576)， 每个机构提供1个排序节点和1个对等节点，对等节点使用GoLevelDB作为状态[数据库](http://msd.misuland.com/pd/3691885236884346610)。

## 3、安装Hyperledger Caliper

首先参考官方文档在HOST1上安装Hyperledger Caliper。

然后在单机模式下运行基准测试以确认Hyperledger Caliper安装正确。

在所有参与Fabric网络的虚拟机上克隆Hyperledger Caliper软件仓库，或者创建HOST1的镜像。

在所有虚拟机上将本文提供的测试网络配置文件 克隆到本地以下目录：

```
~/caliper/packages/caliper-samples/network/fabric-v1.4.1/swarm-3org1peer-raft
```

现在，我们已经准备好了进行Hyperledger Caliper基准测试的Fabric网络环境。下图展示了本文 示例基准测试的架构(Fabric version 1.4.1, 3org1peergoleveldb, Fabric-CCP adaptor)：



查看虚拟机的内部IP，确保虚拟机之间可以彼此通信。你可能需要配置防火墙规则以放行Caliper使用的某些端口，例如7050,7051,7054,8051,8054等。

## 4、运行Caliper基准测试

我们使用Docker Swarm来管理容器。

### 4.1 创建Docker Swarm

首先在我们运行基准Caliper测试的虚拟机上创建一个Docker Swarm。如果你的所有主机在同一个网段，你也可以使用主机的内部IP。

```
$ docker swarm init --listen-addr HOST1-REACHABLE-PUBLIC-IP:2377
```

注意：在上面的命令中，需要将IP换成你自己的IP。

运行上述命令后应当可以看到如下输出显示：

```
# To add a worker to this swarm, run the following command:
#    docker swarm join --token xxxxxxxxxxxxx IP:2377
```

这意味着你的Host1已经作为管理者加入了Swarm。

### 4.2 将其他节点加入Swarm集群

将以下命令复制到其他主机并执行：

```
docker swarm join --token xxxxxxxxxxxxx IP:2377
```

上述命令执行后将输入如下内容，表示该主机已经加入Swarm集群：

```
# This node joined a swarm as a worker.
```

在所有需要运行Caliper基准测试的机器上运行上述命令以加入Swarm集群。

注意：所有参与Hyperledger Caliper基准测试的主机上，Caliper仓库的目录都应该和HOST1上的一致，因为系统需要在此路径定位密码学材料。例如：

```
#home/HOST1/caliper/packages/caliper-samples/network/fabric-v1.4.1/swarm-3org1peer-raft/
#home/HOST2/caliper/packages/caliper-samples/network/fabric-v1.4.1/swarm-3org1peer-raft/
#home/HOST3/caliper/packages/caliper-samples/network/fabric-v1.4.1/swarm-3org1peer-raft/
```

### 4.3 检查Docker Swarm状态

在HOST1上运行以下命令查看Swarm集群中的主机清单：

```
$ docker node ls
```

结果如下：

```
ID        HOSTNAME         STATUS   AVAILABILITY     MANAGER STATUS 
xxx *   caliper-latest      Ready     Active             Leader   
xxx     caliper-latest2     Ready     Active
xxx     caliper-latest3     Ready     Active
```

### 4.4 将Fabric服务容器分配到主机

在编辑器里打开 docker-swarm-compose-tls.yaml文件，该文件中Services下定义了Peer/Order/CA容器。在Service定义的末尾部分可以找到容器部署的细节：

```
deploy:
    placement:
        constraints: [node.hostname ==  YOUR-HOSTNAME]
```

将`YOUR-HOSTNAME`修改为你希望该服务部署的目标主机名。

### 4.5、运行Hyperledger Caliper基准测试

现在在HOST1的Caliper目录下运行如下命令启动基准测试：

```
$ caliper benchmark run -w ./packages/caliper-samples \
          -c benchmark/simple/config.yaml \
          -n network/fabric-v1.4.1/swarm-3org1peer-raft/fabric-ccp-go-tls.yaml
```

------

原文链接：Hyperledger Caliper多排序节点基准测试 — 汇智网