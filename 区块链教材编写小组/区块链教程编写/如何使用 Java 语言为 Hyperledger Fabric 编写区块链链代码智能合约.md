# 如何使用 Java 语言为 Hyperledger Fabric 编写区块链链代码智能合约

### 面向 Java 开发人员的链代码简介

您或许听说过区块链，但可能不确定它对 Java™ 开发人员有何用。本教程将帮助大家解惑。我将分步展示如何使用 Hyperledger Fabric v0.6 来构建、运行和执行使用 Java 语言编写的智能合约或链代码。您将安装一些工具，定义本地区块链网络，构建并运行一个链代码智能合约。

有关区块链的概述，请参阅 developerWorks 博客中的 “[区块链是什么？分布式账本技术入门](https://www.ibm.com/developerworks/community/blogs/3302cc3b-074e-44da-90b1-5055f1dc0d9c/entry/what-is-blockchain-hyperledger-fabric-distributed-ledger?lang=zh)”。

### 前提条件

本教程假设您满足以下前提条件：

- 听说过区块链或 Hyperledger Fabric 框架
- 拥有中等水平的 Java 编程知识，以及使用该语言和平台的经验
- 熟悉或（最好）精通使用：
  - Eclipse IDE
  - Docker 和 Docker Compose
  - Gradle
  - Linux 命令行
  - SoapUI 或另一种 HTTP 客户端软件，比如 Postman

您还应该能够在最少的指导下在计算机上安装软件。由于篇幅所限，我不会介绍所有必要软件的详细安装说明；为您提供软件的网站应该提供了安装说明。

深入介绍本教程之前，我想简单说说区块链。

### 区块链基础

尽管关于区块链存在大量炒作，但它确实物有所值。该技术不仅本身很有趣，而且它是颠覆性的，有潜力彻底变革互联网上的业务执行方式。

如何变革？让我们来想想一次成功的业务交易的基本属性：

- 信任：我们达成了协议，但我能够真的相信您会遵守它（或尊重我）吗？
- 透明性：允许查看 “幕后情况”（这既能建立信任，又能减少对信任的需求）。
- 责任性：用来定义确定所有各方是否都认为该协议得以遵守的条件。
   两方或多方之间的任何业务关系的健康程度对应于上述 3 种属性的不同级别（例如，更信任意味着需要的透明度更少，反之亦然），但其中一些属性必须存在，否则就会出现问题。

> “区块链技术正被快速应用到您身边的软件开发项目中。您准备好了吗？”

区块链有何帮助？首先，通过使用通用框架，业务合作伙伴可以提前建立一个信任网络。然后，通过使用对所有交易方可见的账本，区块链提供了透明性。最后，通过采用所有各方的一致意见（使用智能合约或链代码的形式），建立了责任制。

这对 Java 开发人员有何意义？

Hyperledger 社区和 Hyperledger Fabric 的快速发展，意味着区块链技术正快速被应用到您身边的软件开发项目中。您准备好了吗？

### 区块链技术发展形势

有时，开发技术可能会妨碍业务问题的解决。本教程的主要目的是展示如何编写 Java 链代码，所以我选择了最简单的开发技术组合来实现此目的。

也就是说，该组合中的组件还有其他选择。在本教程中，我将使用 [Docker](https://www.docker.com/) 作为网络容器环境，而另一个选择是带 [VirtualBox](https://www.virtualbox.org/) 的 [Vagrant](https://www.vagrantup.com/)。如果从未使用过 Vagrant，您至少应该尝试一下。

Docker 是一个容器环境，而 Vagrant 使用了虚拟化。在与 VirtualBox 结合使用时，虚拟化环境对计算环境进行了不同程度的控制，这一点备受一些开发人员的青睐（使其成为 fabric 开发人员的理想选择）。

如果想进一步了解容器化与虚拟化孰优孰劣，请参阅 developerWorks 博客上的 “[Docker 是什么？容器对应用程序有哪些好处的简介”](https://www.ibm.com/developerworks/community/blogs/3302cc3b-074e-44da-90b1-5055f1dc0d9c/entry/what-is-docker-containers?lang=zh)。

如果开发人员只想编写代码，而不必担心容器、虚拟化或任何基础架构，那么可以选择 IBM® Bluemix®。尽管 Bluemix 支持运行完整的 IBM 区块链网络，但它目前还不支持使用 Java 语言开发链代码。预计此情况很快就会发生变化，所以请时刻关注。

如果在您的印象中，区块链技术当时的发展情况非常不稳定，您是对的。但是，这意味着您在恰当的时机（从一开始）就开始接触区块链和链代码。随着该技术发展成熟，您早期阶段学习该技术的投资将会不断得到丰厚的回报。

区块链是能够彻底改变每个人的业务执行方式的颠覆式技术之一。这类技术不仅包括 B2B，还包括 B2C，甚至还有 C2C。这的确是一个非常激动人心的时刻。

让我们开始吧！

### 设置开发环境

要运行链代码，首先需要设置开发环境。

完成本节后，您就可以运行一个 Hyperledger Java 链代码示例了，在该示例中，您将在真实链代码上部署和调用交易。然后，我将展示如何（几乎）从头编写一个新链代码程序。

在本节中，您将：

- 设置网络环境 — 用于运行您的本地区块链网络。
- 安装构建软件 — 用于构建您的链代码。
- 安装一个 HTTP 客户端 — 用于在您的链代码上调用交易。
- 启动区块链网络。
- 构建 Java shim 客户端 JAR。

实话说，要编写链代码，有许多设置工作要做。但是，如果您按照这些说明进行操作并稍微勤奋一点，您的付出将是值得的。

#### 1.设置网络环境

本教程将使用 Docker 以及来自 Docker Hub 的预构建区块链网络组件镜像来运行本地区块链网络。如果愿意的话，可以从头构建 fabric（毕竟它是开源的），但在此阶段，使用 Docker Hub 中提供的预构建的 Hyperledger Fabric 镜像更容易一些。

我在介绍中已经提到过，另一个选择（您可能在 Hyperledger 文档中看到过）是使用 Vagrant 和 VirtualBox。Vagrant 是 fabric 开发人员的一个不错选择，但作为链代码开发人员，与处理 fabric 本身相比，我们更关心链代码的构建、运行和测试。

如果已经安装 Docker 1.12 版或更高版本，可以跳到下一节（“安装构建软件”）。在下面的操作说明中，假设您尚未安装 Docker（也就是说，您不是从以前的 Docker 版本进行升级）。安装 Docker 的过程中也会安装 Docker Compose，该工具用于定义和运行需要多个容器的应用程序，比如本教程中将运行的本地 Hyperledger 区块链网络。

**安装 Docker**
 可在这里找到针对 Mac、Windows 和 Linux 的安装说明：

[将 Docker 安装在 Mac、Windows 和 Linux 上](https://docs.docker.com/engine/getstarted/step_one)

**验证 Docker 安装**
 要测试 Docker 安装，可打开一个终端窗口（或 Windows 上的命令提示符）并键入以下命令：



```undefined
docker -v
docker-compose -v
```

您会获得以下输出：



```ruby
$ docker -v
Docker version 1.13.1, build 092cba3
$ docker-compose -v
docker-compose version 1.11.1, build 7c5d5e4
```

如果想查看 Docker 的实际运行效果，您可运行 hello-world 镜像，如下所示：



```jsx
$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
78445dd45222: Pull complete 
Digest: sha256:c5515758d4c5e1e838e9cd307f6c6a0d620b5e07e6f927b07d05f6d12a1ac8d7
Status: Downloaded newer image for hello-world:latest
 

Hello from Docker! This message shows that your installation appears to be working correctly.

 
To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.
 
To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash
 
Share images, automate workflows, and more with a free Docker ID:
 https://cloud.docker.com/
 
For more examples and ideas, visit:
 https://docs.docker.com/engine/userguide/
```

#### 2.安装构建软件

对于构建系统，Hyperledger Fabric 使用了 Gradle，本教程也将使用它。Gradle 是一个构建自动化系统，它结合使用了指定构建组件的简单语法与 Apache Ant 和 Apache Maven 的最佳特性，创建了一个容易使用的强大构建系统。如此多开发人员正将他们的项目切换到 Gradle，这不足为奇。请注意，本文使用的是 Gradle 3.3 版本。使用其他版本的 Gradle 可能会导致项目搭建失败。

可以在 [Gradle 主页](https://gradle.org/) 上进一步了解 Gradle（和它的一些知名度高的用户）。

**安装 Gradle**
 要安装 Gradle，请按照下面的说明进行操作：

[将 Gradle 安装在 Mac、Windows 和 Linux 上](https://gradle.org/install)

**验证 Gradle 安装**
 要验证 Gradle 安装，可打开一个终端窗口并执行此命令：



```undefined
gradle -v
```

您会看到以下输出：



```csharp
$ gradle -v
 
------------------------------------------------------------
Gradle 3.3
------------------------------------------------------------
 
Build time:   2017-01-03 15:31:04 UTC
Revision:     075893a3d0798c0c1f322899b41ceca82e4e134b
 
Groovy:       2.4.7
Ant:          Apache Ant(TM) version 1.9.6 compiled on June 29 2015
JVM:          1.8.0_102 (Oracle Corporation 25.102-b14)
OS:           Mac OS X 10.12.3 x86_64
```

#### 3.安装 HTTP 客户端

接下来安装 HTTP 客户端软件，它允许链代码与 Hyperledger 区块链结构的 REST 接口进行通信。您的浏览器可以发出 HTTP GET，但要与 fabric 进行交互，您需要能够通过 POST 发出消息。这意味着您需要一个 HTTP 客户端。

我为本教程选择的 HTTP 客户端是 SoapUI，它提供了一个强大的、容易使用的、包含许多功能的免费社区版本。

**安装 SoapUI**
 要安装 SoapUI，请按照下面的说明进行操作：

[为 Mac OS、Windows 和 Linux 安装 SoapUI](https://www.soapui.org/getting-started/installing-soapui.html)

**验证 SoapUI 安装**
 要确认 SoapUI 已安装，可在计算机上启动该应用程序。在 Mac OS 上，打开 SoapUI 后会显示 SoapUI Starter Page，如图 1 所示。

Mac OS X 上的 SoapUI



![img](https:////upload-images.jianshu.io/upload_images/11831773-b809b7068f365ded.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

Mac OS X 上的 SoapUI

#### 4.启动区块链网络

现在您已安装开发和测试链代码所需的软件，是时候启动本地区块链网络了。第一步是定义网络的配置。

首先，创建一个目录，用它作为链代码开发过程中使用的所有源代码的 root 目录。在本教程中，我将使用 ~/home/mychaincode（或 Windows 上的 C:\home\chaincode）。

接下来，将 GOPATH 环境变量设置为此路径。我们不会编译任何 Go 代码，也不会构建 Golang 包或其他二进制文件，但 Golang 术语已融合到 Hyperledger 中，所以熟悉按 Go 语言和 GOPATH 的方式进行思考是一个不错的主意。

在 Linux 上，执行以下命令：



```jsx
export GOPATH=~/home/mychaincode
```

或者在 Windows 上，可以使用以下命令：



```undefined
SET GOPATH=C:\home\mychaincode
```

接下来，必须告诉 Docker Compose 如何创建和运行区块链对等网络。该网络是使用 YAML 定义的，应该将它命名为 docker-compose.yml。可以将文件命名为其他名称，但在启动 Docker Compose 时，必须指定 -f 标志。建议坚持使用默认名称，也就是 docker-compose.yml。

在 GOPATH 的 root 目录中创建 docker-compose.yml 文件。粘贴以下内容：



```bash
membersrvc:
  image: hyperledger/fabric-membersrvc
  ports:
    - "7054:7054"
  command: membersrvc
vp0:
  image: hyperledger/fabric-peer:x86_64-0.6.0-preview
  ports:
    - "7050:7050"
    - "7051:7051"
    - "7053:7053"
  environment:
    - CORE_PEER_ADDRESSAUTODETECT=true
    - CORE_VM_ENDPOINT=unix:///var/run/docker.sock
    - CORE_LOGGING_LEVEL=DEBUG
    - CORE_PEER_ID=vp0
    - CORE_PEER_PKI_ECA_PADDR=membersrvc:7054
    - CORE_PEER_PKI_TCA_PADDR=membersrvc:7054
    - CORE_PEER_PKI_TLSCA_PADDR=membersrvc:7054
    - CORE_SECURITY_ENABLED=false
    - CORE_SECURITY_ENROLLID=test_vp0
    - CORE_SECURITY_ENROLLSECRET=MwYpmSRjupbT
  links:
    - membersrvc
  command: sh -c "sleep 5; peer node start --peer-chaincodedev"
```

这里涉及许多内容，其中大部分内容都不属于本教程的讨论范围，但我想稍微解释一下。

- 此文件告诉 Docker Compose 定义两个服务：
  - membersrvc：提供成员服务的成员服务节点，具体来讲，它是一个证书签发机构 (CA)，负责处理所有加密解密工作（比如颁发和撤销证书）。将用于此目的的预构建 Docker 镜像命名为 hyperledger/fabric-membersrvc。
  - vp0：网络中的单独验证对等节点。对于开发目的，我们不需要奢侈地执行对等网络验证，只需要一个对等节点即可。将用于此目的的预构建 Docker 镜像命名为 hyperledger/fabric-peer： x86_64-0.6.0-preview。请注意，由于 Fabric 版本升级，如果省略掉冒号及之后的字符，会导致构建错误。此处指定使用 Fabric 0.6 版本。hyperledger/fabric-peer 等同于 hyperledger/fabric-peer：latest。
- 一些环境变量由 vp0 对等节点设置。请注意，CORE_LOGGING_LEVEL 变量被设置为 DEBUG。这会生成大量输出，这些输出有时很方便。但是，如果想要更少的输出，可将该级别更改为 INFO。请参阅 Hyperledger 设置文档中的 “[日志控制](https://github.com/hyperledger/fabric/blob/v0.6/docs/Setup/logging-control.md)” 了解日志级别的更多信息。
   可以在 [Docker 网站](https://docs.docker.com/compose/compose-file/) 上找到有关 Docker Compose YML 文件定义的更多信息。

接下来请注意，CORE_SECURITY_ENABLED 的值为 false。这意味着 fabric 不需要您发送任何类型的最终用户凭证。安全性不属于本教程的讨论范围，但是如果您有兴趣了解更多信息，可以根据您的链代码请求来查阅[这篇安全功能说明](https://github.com/hyperledger/fabric/blob/v0.6/docs/Setup/Chaincode-setup.md#note-on-security-functionality)。

最后提醒一句：对任何这些值的默认设置（尤其是端口值）的任何改动都有可能导致本教程中的示例无法运行。区块链网络是一组分布式软件组件，它们需要经过精确协调的通信。强烈建议在理解 fabric 的所有组件如何相互作用之前，**不要**更改端口值的默认值。

完成区块链的定义后，就可以启动本地区块链网络了。为此，请运行 Docker Compose。导航到您的 $GOPATH 并执行此命令：



```undefined
docker-compose up
```

您会在终端窗口中获得以下输出：



```rust
$ docker-compose up
.
.
Pulling membersrvc (hyperledger/fabric-membersrvc:latest)...
latest: Pulling from hyperledger/fabric-membersrvc
.
.
Status: Downloaded newer image for hyperledger/fabric-membersrvc:latest
Pulling vp0 (hyperledger/fabric-peer:latest)...
latest: Pulling from hyperledger/fabric-peer
.
.
Status: Downloaded newer image for hyperledger/fabric-peer:latest
Creating mychaincode_membersrvc_1
Creating mychaincode_vp0_1
Attaching to mychaincode_membersrvc_1, mychaincode_vp0_1
vp0_1         | 19:30:03.773 [logging] LoggingInit -> DEBU 001 Setting default logging level to DEBUG for command 'node'
vp0_1         | 19:30:03.773 [nodeCmd] serve -> INFO 002 Running in chaincode development mode
.
.
.
vp0_1         | 19:30:04.146 [peer] chatWithSomePeers -> DEBU 07c Starting up the first peer of a new network
vp0_1         | 19:30:04.146 [consensus/statetransfer] verifyAndRecoverBlockchain -> DEBU 07d Validating existing blockchain, highest validated block is 0, valid through 0
vp0_1         | 19:30:04.146 [consensus/statetransfer] blockThread -> INFO 07e Validated blockchain to the genesis block
vp0_1         | 19:30:04.146 [consensus/handler] 1 -> DEBU 07f Starting up message thread for consenter
vp0_1         | 19:30:04.146 [nodeCmd] serve -> INFO 080 Starting peer with ID=name:"vp0" , network ID=dev, address=172.17.0.3:7051, rootnodes=, validator=true
vp0_1         | 19:30:04.146 [rest] StartOpenchainRESTServer -> INFO 081 Initializing the REST service on 0.0.0.0:7050, TLS is disabled.
vp0_1         | 19:30:04.147 [peer] ensureConnected -> DEBU 082 Starting Peer reconnect service (touch service), with period = 6s
.
.
```

此输出告诉您该网络在正常运行，已准备好接受链代码注册请求。

**备注**：突出显示的行应该仅在第一次运行区块链网络时出现，因为 Docker 需要从 Docker Hub 下载镜像。镜像下载到计算机后，仅在来自 Docker Hub 的镜像比您计算机上的镜像更新时，Docker 才会拉入它们。

现在已准备好构建 Java shim 客户端 JAR，它允许 Java 语言链代码与 Hyperledger Fabric 框架进行通信。

#### 5.构建 Java shim 客户端 JAR

在运行链代码示例前，需要从 Hyperledger 的 GitHub 存储库获取最新的源代码。

首先，需要将 Hyperledger Fabric 克隆到本地机器上，以便构建链代码（**备注**：这是一项临时措施；在以后某个时刻，应该能从主要的 Maven 存储库访问 Java shim 客户端 JAR）。

**备注**：回想一下，您之前已将 GOPATH 设置为 Linux（或 Mac）上的 ~/home/mychaincode 或 Windows 上的 C:\home\mychaincode。

执行此命令来创建结构构建脚本所期望的目录结构：



```bash
mkdir -p $GOPATH/src/github.com/hyperledger
```

接下来，导航到已创建的新目录结构的底部：



```bash
cd $GOPATH/src/github.com/hyperledger
```

您需要从这里获取 Hyperledger 源代码，以便构建 Java shim 客户端 JAR。

可通过两种方式获得 Hyperledger 源代码。

- 不使用 git：
   导航到 [Hyperledger GitHub 镜像](https://github.com/hyperledger/fabric)，选择v0.6分支，并单击 Clone or download 按钮，然后单击 Download ZIP（参见图 2）。一个名为 fabric-master.zip 的 ZIP 文件被下载到您的计算机，您可以将它解压到 $GOPATH/src/github.com/hyperledger。备注：请确保在解压该文件时，将 root 目录的名称从 fabric-master 更改为 fabric。
- 使用 git：
   导航到 $GOPATH/src/github.com/hyperledger，将文本字段中的 URL 复制到 “Clone with HTTPS” 框中（参见图 2 中的箭头），然后使用复制的 URL 执行此命令：



```php
git clone https://github.com/hyperledger/fabric.git
```

或



```php
git clone –b v0.6 
https://gerrit.hyperledger.org/r/fabric
```

您会看到 git 命令返回了以下终端窗口输出：



```bash
$ git clone https://github.com/hyperledger/fabric.git
Cloning into 'fabric'...
remote: Counting objects: 26976, done.
remote: Compressing objects: 100% (406/406), done.
remote: Total 26976 (delta 172), reused 0 (delta 0), pack-reused 26558
Receiving objects: 100% (26976/26976), 43.68 MiB | 4.85 MiB/s, done.
Resolving deltas: 100% (15114/15114), done.
```

![img](https:////upload-images.jianshu.io/upload_images/11831773-9d70a5feac97f4e8.jpg?imageMogr2/auto-orient/strip|imageView2/2/w/1045/format/webp)

Hyperledger GitHub 镜像

现在您已准备好构建 Java 链代码 shim 客户端 JAR。导航到 $GOPATH/src/github.com/hyperledger/fabric/core/chaincode/shim/java 并运行以下两个命令：



```css
gradle -b build.gradle clean
gradle -b build.gradle build
```

Gradle 构建输出应如下所示：



```ruby
$ cd $GOPATH/src/github.com/hyperledger/fabric/core/chaincode/shim/java
$ gradle -b build.gradle clean
Starting a Gradle Daemon (subsequent builds will be faster)
:core:chaincode:shim:java:clean
 
BUILD SUCCESSFUL
 
Total time: 5.422 secs
$ gradle -b build.gradle build
:core:chaincode:shim:java:copyProtos UP-TO-DATE
:core:chaincode:shim:java:extractIncludeProto
:core:chaincode:shim:java:extractProto UP-TO-DATE
:core:chaincode:shim:java:generateProto UP-TO-DATE
:core:chaincode:shim:java:compileJava
:core:chaincode:shim:java:processResources
:core:chaincode:shim:java:classes
:core:chaincode:shim:java:jar
:core:chaincode:shim:java:assemble
:core:chaincode:shim:java:extractIncludeTestProto
:core:chaincode:shim:java:extractTestProto UP-TO-DATE
:core:chaincode:shim:java:generateTestProto UP-TO-DATE
:core:chaincode:shim:java:compileTestJava UP-TO-DATE
:core:chaincode:shim:java:processTestResources UP-TO-DATE
:core:chaincode:shim:java:testClasses UP-TO-DATE
:core:chaincode:shim:java:test UP-TO-DATE
:core:chaincode:shim:java:check UP-TO-DATE
:core:chaincode:shim:java:build
:core:chaincode:shim:java:copyToLib
:core:chaincode:shim:java:generatePomFileForMavenJavaPublication
:core:chaincode:shim:java:publishMavenJavaPublicationToMavenLocal
:core:chaincode:shim:java:publishToMavenLocal
 
BUILD SUCCESSFUL
 
Total time: 4.521 secs
```

构建过程中执行的最后一件事是，将 shim 客户端 JAR 添加到本地 Maven 存储库。现在您已准备好构建链代码。除非在未来某个时刻要更新结构源代码，或者出于某种原因想要再次重新构建 shim 客户端 JAR，否则不需要再次运行 Java shim 客户端 JAR。

### 部署并运行 Java 链代码示例

您已经定义并启动了本地区块链网络，而且已构建 Java shim 客户端 JAR 并安装到本地 Maven 存储库中，现在已准备好在之前下载的 Hyperledger Fabric 附带的一个 Java 链代码示例上构建、注册和调用交易。

**部署并运行链代码**

您将执行以下步骤：

- 使用 Gradle 构建示例。
- 通过运行 Gradle 构建软件为您创建的脚本，向验证对等网络注册该示例。
- 使用 SoapUI 将示例部署到本地区块链网络。
- 使用 SoapUI 在示例链代码上调用交易。

#### 1.构建示例

导航到 $GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/java/Example 目录。

接下来，通过命令行，使用此命令启动 Gradle 构建软件：



```css
gradle -b build.gradle build
```

您会看到以下输出：



```ruby
$ cd GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/java/Example
$ gradle -b build.gradle build
Starting a Gradle Daemon (subsequent builds will be faster)
:examples:chaincode:java:Example:compileJava
:examples:chaincode:java:Example:processResources UP-TO-DATE
:examples:chaincode:java:Example:classes
:examples:chaincode:java:Example:jar
:examples:chaincode:java:Example:startScripts
:examples:chaincode:java:Example:distTar
:examples:chaincode:java:Example:distZip
:examples:chaincode:java:Example:assemble
:examples:chaincode:java:Example:compileTestJava UP-TO-DATE
:examples:chaincode:java:Example:processTestResources UP-TO-DATE
:examples:chaincode:java:Example:testClasses UP-TO-DATE
:examples:chaincode:java:Example:test UP-TO-DATE
:examples:chaincode:java:Example:check UP-TO-DATE
:examples:chaincode:java:Example:build
:examples:chaincode:java:Example:copyToLib
 
BUILD SUCCESSFUL
 
Total time: 6.935 secs
```

该构建过程通过两种形式创建了一个位于目录 build/distributions 中的独立发行版：TAR 文件和 ZIP 文件，每个文件都包含运行链代码所需的所有资源，其中包括一个用于驱动链代码的名为 Example 的脚本。

Example 链代码现在已准备好向本地区块链网络注册。

#### 2.注册示例

确保本地区块链网络正在运行。如果未运行，则需要启动它。如果需要温习一下相关内容，请参阅“[启动区块链网络](https://www.ibm.com/developerworks/cn/java/j-chaincode-for-java-developers/index.html#blockchainnetwork)”部分。

如果您未在 $GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/java/Example 目录下，请导航到这里。

接下来，将 Example.zip（或 Example.tar）解压到 build/distributions 目录中：



```ruby
$ cd $GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/java/Example
$ cd build/distributions/
$ unzip Example.zip 
Archive:  Example.zip
  inflating: Example/lib/chaincode.jar  
  inflating: Example/lib/grpc-all-0.13.2.jar  
  inflating: Example/lib/commons-cli-1.3.1.jar  
  inflating: Example/lib/shim-client-1.0.jar  
  inflating: Example/lib/grpc-netty-0.13.2.jar  
  inflating: Example/lib/grpc-auth-0.13.2.jar  
  inflating: Example/lib/grpc-protobuf-nano-0.13.2.jar  
  inflating: Example/lib/grpc-core-0.13.2.jar  
  inflating: Example/lib/grpc-protobuf-0.13.2.jar  
  inflating: Example/lib/grpc-okhttp-0.13.2.jar  
  inflating: Example/lib/grpc-stub-0.13.2.jar  
  inflating: Example/lib/protobuf-java-3.0.0.jar  
  inflating: Example/lib/netty-tcnative-boringssl-static-1.1.33.Fork21-osx-x86_64.jar  
  inflating: Example/lib/netty-codec-http2-4.1.0.CR3.jar  
  inflating: Example/lib/google-auth-library-oauth2-http-0.3.0.jar  
  inflating: Example/lib/guava-18.0.jar  
  inflating: Example/lib/protobuf-javanano-3.0.0-alpha-5.jar  
  inflating: Example/lib/jsr305-3.0.0.jar  
  inflating: Example/lib/okio-1.6.0.jar  
  inflating: Example/lib/okhttp-2.5.0.jar  
  inflating: Example/lib/netty-codec-http-4.1.0.CR3.jar  
  inflating: Example/lib/netty-handler-4.1.0.CR3.jar  
  inflating: Example/lib/google-auth-library-credentials-0.3.0.jar  
  inflating: Example/lib/google-http-client-1.19.0.jar  
  inflating: Example/lib/google-http-client-jackson2-1.19.0.jar  
  inflating: Example/lib/netty-codec-4.1.0.CR3.jar  
  inflating: Example/lib/netty-buffer-4.1.0.CR3.jar  
  inflating: Example/lib/netty-transport-4.1.0.CR3.jar  
  inflating: Example/lib/httpclient-4.0.1.jar  
  inflating: Example/lib/jackson-core-2.1.3.jar  
  inflating: Example/lib/netty-common-4.1.0.CR3.jar  
  inflating: Example/lib/netty-resolver-4.1.0.CR3.jar  
  inflating: Example/lib/httpcore-4.0.1.jar  
  inflating: Example/lib/commons-logging-1.1.1.jar  
  inflating: Example/lib/commons-codec-1.3.jar  
  inflating: Example/bin/Example     
  inflating: Example/bin/Example.bat
```

您可能想知道 “为何有如此多的文件？”该发行版包含（在独立进程中）单独运行链代码所需的一切资源，以及所有依赖 JAR 文件。

要注册链代码示例，可在 build/distributions 文件夹中执行以下脚本：



```undefined
./Example/bin/Example
```

这会运行一个独立进程来向本地区块链网络注册链代码示例。您会看到以下终端窗口输出：



```ruby
$ ./Example/bin/Example
Hello world! starting [Ljava.lang.String;@7ef20235
Feb 22, 2017 10:05:08 AM example.Example main
INFO: starting
Feb 22, 2017 10:05:08 AM org.hyperledger.java.shim.ChaincodeBase newPeerClientConnection
INFO: Inside newPeerCLientConnection
Feb 22, 2017 10:05:08 AM io.grpc.internal.TransportSet$1 call
INFO: Created transport io.grpc.netty.NettyClientTransport@3dd7b80b(/127.0.0.1:7051) for /127.0.0.1:7051
Feb 22, 2017 10:05:14 AM io.grpc.internal.TransportSet$TransportListener transportReady
INFO: Transport io.grpc.netty.NettyClientTransport@3dd7b80b(/127.0.0.1:7051) for /127.0.0.1:7051 is ready
```

查看本地区块链网络的控制台，您会看到以下输出行：



```ruby
.
.
vp0_1         | 16:05:14.048 [chaincode] HandleChaincodeStream -> DEBU 06d Current context deadline = 0001-01-01 00:00:00 +0000 UTC, ok = false
vp0_1         | 16:05:14.065 [chaincode] processStream -> DEBU 06e []Received message REGISTER from shim
vp0_1         | 16:05:14.065 [chaincode] HandleMessage -> DEBU 06f []Handling ChaincodeMessage of type: REGISTER in state created
vp0_1         | 16:05:14.065 [chaincode] beforeRegisterEvent -> DEBU 070 Received REGISTER in state created
vp0_1         | 16:05:14.065 [chaincode] registerHandler -> DEBU 071 registered handler complete for chaincode hello
vp0_1         | 16:05:14.065 [chaincode] beforeRegisterEvent -> DEBU 072 Got REGISTER for chaincodeID = name:"hello" , sending back REGISTERED
.
.
```

记下注册日志输出中的 chaincodeID name（示例中为 hello；如上面 **第 8 行** 所示）。以后在通过结构的 REST 接口部署 Example 链代码时，JSON 消息中需要使用此信息。

上面的输出表明 Example 链代码正在运行，而且已向本地区块链验证对等网络注册，并做好了部署准备。

#### 3.部署示例

Hyperledger Fabric 提供了一个用于与该结构交互的 REST Web 服务接口。与 fabric 的第一次交互是部署链代码。确保本地区块链网络正在运行，然后启动 SoapUI，单击 REST 按钮创建一个新的 REST 项目。您会看到一个类似图 3 的对话框，在其中输入用于所有 REST 请求的基础 URL：

![img](https:////upload-images.jianshu.io/upload_images/11831773-b1379870dac0e603.png?imageMogr2/auto-orient/strip|imageView2/2/w/547/format/webp)

SoapUI New REST Project 对话框

输入 http://localhost:7050 作为 URL，然后单击 OK。端口 7050 是 fabric 使用的默认 REST 端口，而且因为区块链网络是在本地计算机上运行的，所以将使用 localhost 作为主机名。

在 SoapUI 启动后，可以执行一次快速冒烟测试，以确保它能与本地区块链网络进行通信。展开刚创建的新的 REST 资源，直到看到 Request 1，然后在 Editor 窗口中打开它。使用 GET 方法，在 resource 下输入 /chain。确保单击了 output 选项卡上的 JSON 选项，然后运行请求（通过单击 arrow 图标）。执行此请求时，会在 Editor 窗口右侧的输出选项卡中返回当前区块的哈希值，如图 4 所示：



![img](https:////upload-images.jianshu.io/upload_images/11831773-48bb0a4cfae8635b.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

区块链冒烟测试

如果看到一条类似图 4 的 JSON 消息（当然您的网络的 currentBlockHash 值会有所不同），那么您已准备好部署 Example 链代码。

右键单击 REST Project 1 (http://localhost:7050) 下的端点并选择 New Resource；您会看到一个包含 Resource Path 字段的 “New REST Resource” 对话框（参见图 5）：

![img](https:////upload-images.jianshu.io/upload_images/11831773-1283767811a019cf.png?imageMogr2/auto-orient/strip|imageView2/2/w/615/format/webp)

 SoapUI New Resource 对话框


 输入 /chaincode 作为 resource path，然后单击 OK，您会看到 SoapUI Projects 面板中显示了新资源。打开对此资源的请求（默认情况下该请求名为 Request 1），将方法更改为 POST，并将此 JSON 粘贴到请求编辑器窗口左下角的请求区域：





```json
{
"jsonrpc": "2.0",
  "method": "deploy",
  "params": {
    "type": 1,
    "chaincodeID":{
        "name": "hello"
    },
    "CtorMsg": {
        "args": [""]
    }
  },
  "id": 1
}
```

有 3 点需要注意：

**第 3 行**：method 值必须为 deploy。
 **第 6-7 行**：JSON 消息中的 chaincodeID.name 必须与您在上一节中注册 Example 链代码时所用的 chaincodeID 匹配（在 Example 链代码中，该值为 hello）。
 **第 13 行**：id 值用于协调请求。本教程不需要过多地考虑它，但要注意的是，在响应中始终会发送回该值（参见下一个清单）。
 提交此请求时，JSON 输出应如下所示：



```json
{
   "jsonrpc": "2.0",
   "result":    {
      "status": "OK",
      "message": "hello"
   },
   "id": 1
}
```

图 6 给出了 SoapUI 中的输出的屏幕截图。JSON 输出消息会显示在输出选项卡中，该选项卡位于请求编辑器的右侧。



![img](https:////upload-images.jianshu.io/upload_images/11831773-2252d760ec81ef34.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

SoapUI 链代码部署请求

终端窗口中的网络日志输出应包含以下行：



```ruby
.
.
vp0_1         | 20:48:39.482 [rest] ProcessChaincode -> INFO 0c4 REST processing chaincode request...
vp0_1         | 20:48:39.482 [rest] processChaincodeDeploy -> INFO 0c5 REST deploying chaincode...
vp0_1         | 20:48:39.483 [devops] Deploy -> DEBU 0c6 Creating deployment transaction (hello)
vp0_1         | 20:48:39.483 [devops] Deploy -> DEBU 0c7 Sending deploy transaction (hello) to validator
vp0_1         | 20:48:39.483 [peer] sendTransactionsToLocalEngine -> DEBU 0c8 Marshalling transaction CHAINCODE_DEPLOY to send to local engine
vp0_1         | 20:48:39.483 [peer] sendTransactionsToLocalEngine -> DEBU 0c9 Sending message CHAIN_TRANSACTION with timestamp seconds:1487796519 nanos:483661510  to local engine
vp0_1         | 20:48:39.483 [consensus/noops] RecvMsg -> DEBU 0ca Handling Message of type: CHAIN_TRANSACTION 
vp0_1         | 20:48:39.483 [consensus/noops] broadcastConsensusMsg -> DEBU 0cb Broadcasting CONSENSUS
vp0_1         | 20:48:39.483 [peer] Broadcast -> DEBU 0cc Broadcast took 1.135s
vp0_1         | 20:48:39.483 [consensus/noops] RecvMsg -> DEBU 0cd Sending to channel tx uuid: hello
vp0_1         | 20:48:39.483 [rest] processChaincodeDeploy -> INFO 0ce Successfully deployed chainCode: hello
vp0_1         | 20:48:39.484 [rest] ProcessChaincode -> INFO 0cf REST successfully deploy chaincode: {"jsonrpc":"2.0","result":{"status":"OK","message":"hello"},"id":1}
.
.
```

**第 3-4 行**显示了输出，表明网络已收到部署消息，并且该结构正在部署链代码。**第 13-14 行**表明链代码已成功部署。

在运行链代码的终端窗口中，可以注意到以下输出：



```ruby
$ ./build/distributions/Example/bin/Example
Hello world! starting [Ljava.lang.String;@7ef20235
Feb 22, 2017 2:44:43 PM example.Example main
INFO: starting
Feb 22, 2017 2:44:43 PM org.hyperledger.java.shim.ChaincodeBase newPeerClientConnection
INFO: Inside newPeerCLientConnection
Feb 22, 2017 2:44:43 PM io.grpc.internal.TransportSet$1 call
INFO: Created transport io.grpc.netty.NettyClientTransport@46adccd3(/127.0.0.1:7051) for /127.0.0.1:7051
Feb 22, 2017 2:44:48 PM io.grpc.internal.TransportSet$TransportListener transportReady
INFO: Transport io.grpc.netty.NettyClientTransport@46adccd3(/127.0.0.1:7051) for /127.0.0.1:7051 is ready
Feb 22, 2017 2:48:40 PM example.Example run
INFO: In run, function:
Feb 22, 2017 2:48:40 PM example.Example run
```

我包含了所有上下文输出，在向区块链网络发送部署消息时，您会看到类似第 **11-13 行**的消息。

#### 4.在示例上调用交易

最后，将会调用 hello 方法，可以看到它会在运行链代码的终端窗口的日志消息中显示出来。

在 SoapUI 中的 chaincode 资源下，右键单击 Method 1 并选择 Clone Method。将该方法命名为 Invoke，然后单击 OK。打开新的 Invoke 方法下的 Request 1，并粘贴到以下 JSON 请求中：



```json
{
"jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "type": 1,
    "chaincodeID":{
        "name": "hello"
    },
    "CtorMsg": {
        "args": ["hello"]
    }
  },
  "id": 2
}
```

运行该请求时，会看到以下 JSON 响应：



```json
{ "jsonrpc":"2.0", "result":    { "status":"OK", "message":"1c1811d0-a958-4c58-ab1d-e1df550c18a3" }, "id":2 }
```

图 7 给出了 SoapUI 中的输出的屏幕截图



![img](https:////upload-images.jianshu.io/upload_images/11831773-093a9a27ee8178dc.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

SoapUI 链代码调用请求

网络日志输出应包含以下行：



```bash
.
.
vp0_1         | 21:44:35.143 [rest] ProcessChaincode -> INFO 555 REST processing chaincode request...
vp0_1         | 21:44:35.143 [rest] processChaincodeInvokeOrQuery -> INFO 556 REST invoke chaincode...
vp0_1         | 21:44:35.143 [devops] invokeOrQuery -> INFO 557 Transaction ID: 1c1811d0-a958-4c58-ab1d-e1df550c18a3
vp0_1         | 21:44:35.143 [devops] createExecTx -> DEBU 558 Creating invocation transaction (1c1811d0-a958-4c58-ab1d-e1df550c18a3)
vp0_1         | 21:44:35.143 [devops] invokeOrQuery -> DEBU 559 Sending invocation transaction (1c1811d0-a958-4c58-ab1d-e1df550c18a3) to validator
vp0_1         | 21:44:35.143 [peer] sendTransactionsToLocalEngine -> DEBU 55a Marshalling transaction CHAINCODE_INVOKE to send to local engine
vp0_1         | 21:44:35.143 [peer] sendTransactionsToLocalEngine -> DEBU 55b Sending message CHAIN_TRANSACTION with timestamp seconds:1487799875 nanos:143438691  to local engine
vp0_1         | 21:44:35.143 [consensus/noops] RecvMsg -> DEBU 55c Handling Message of type: CHAIN_TRANSACTION 
vp0_1         | 21:44:35.143 [consensus/noops] broadcastConsensusMsg -> DEBU 55d Broadcasting CONSENSUS
vp0_1         | 21:44:35.143 [peer] Broadcast -> DEBU 55e Broadcast took 1.249s
vp0_1         | 21:44:35.143 [consensus/noops] RecvMsg -> DEBU 55f Sending to channel tx uuid: 1c1811d0-a958-4c58-ab1d-e1df550c18a3
vp0_1         | 21:44:35.143 [rest] processChaincodeInvokeOrQuery -> INFO 560 Successfully submitted invoke transaction with txid (1c1811d0-a958-4c58-ab1d-e1df550c18a3)
vp0_1         | 21:44:35.143 [rest] ProcessChaincode -> INFO 561 REST successfully submitted invoke transaction: {"jsonrpc":"2.0","result":{"status":"OK","message":"1c1811d0-a958-4c58-ab1d-e1df550c18a3"},"id":2}
.
.
```

链代码日志输出如下所示：



```ruby
$ ./build/distributions/Example/bin/Example
Hello world! starting [Ljava.lang.String;@7ef20235
Feb 22, 2017 3:26:57 PM example.Example main
INFO: starting
Feb 22, 2017 3:26:57 PM org.hyperledger.java.shim.ChaincodeBase newPeerClientConnection
INFO: Inside newPeerCLientConnection
Feb 22, 2017 3:26:57 PM io.grpc.internal.TransportSet$1 call
INFO: Created transport io.grpc.netty.NettyClientTransport@765e4953(/127.0.0.1:7051) for /127.0.0.1:7051
Feb 22, 2017 3:27:02 PM io.grpc.internal.TransportSet$TransportListener transportReady
INFO: Transport io.grpc.netty.NettyClientTransport@765e4953(/127.0.0.1:7051) for /127.0.0.1:7051 is ready
Feb 22, 2017 3:27:24 PM example.Example run
INFO: In run, function:
Feb 22, 2017 3:27:24 PM example.Example run
SEVERE: No matching case for function:
Feb 22, 2017 3:30:55 PM example.Example run
INFO: In run, function:hello
hello invoked
```

我再次给出了所有链代码输出。您可以看到哪些地方（**第 16 行**）调用了 hello 函数。

现在您已知道如何在本地区块链网络上构建、部署和运行 Java 链代码。在下一节中，将会使用 Eclipse IDE（几乎）从头编写一个链代码程序，使用 Gradle 构建该链代码程序，然后使用 SoapUI 体验它。

### 编写第一个 Java 链代码程序

在上一节中，您已经熟悉了如何构建、运行、部署和调用链代码，但尚未编写任何 Java 代码。

在本节中，将会使用 Eclipse IDE、一个用于 Eclipse 的 Gradle 插件，以及一个名为 ChaincodeTutorial 的 Java 链代码框架项目，编写第一个 Java 链代码程序。您将从我为此教程创建的 GitHub 存储库中获取框架代码，将该代码导入 Eclipse 中，添加代码来让链代码智慧合同按要求生效，然后在 Eclipse IDE 内使用 Gradle 构建该代码。

您将执行的步骤如下：

- 安装适用于 Eclipse 的 Gradle Buildship 插件。
- 从 GitHub 克隆 ChaincodeTutorial 项目。
- 将该项目导入 Eclipse 中。
- 探索该链代码框架项目。
- 编写 Java 链代码。
- 构建 Java 链代码。
   完成本节后，您的链代码就可以在本地区块链网络上运行了。

#### 1.安装适用于 Eclipse 的 Gradle Buildship 插件

您使用自己喜欢的任何 IDE，但本教程中的说明是针对 Eclipse 的。**备注**：Buildship Gradle 插件有助于将 Gradle 与 Eclipse 集成，但仍然需要将 Gradle 安装在计算机上。

如果您一直在按照教程进行操作，那么您应该已经将 Gradle 安装在计算机上；如果尚未安装它，请立即安装。请参阅 “安装构建软件” 部分，了解如何将 Gradle 安装在计算机上。



![img](https:////upload-images.jianshu.io/upload_images/11831773-dcc1940026c9b835.png?imageMogr2/auto-orient/strip|imageView2/2/w/1046/format/webp)

Eclipse Marketplace：Gradle Buildship 插件

在 **Buildship Gradle Integration** 下，单击 **Install** 按钮并按照提示进行操作。单击 **Finish** 后，将安装适用于 Eclipse 的 Buildship Gradle 插件，而且会要求您重启 Eclipse。

重新打开 Eclipse 后，Gradle 应该已经与 Eclipse IDE 全面集成。您现在已准备好从 GItHub 克隆 ChaincodeTutorial 存储库。

#### 从 GitHub 克隆 ChaincodeTutorial 项目

配置 Eclipse IDE 和 Gradle集成后，将从 GitHub 克隆 ChaincodeTutorial 代码并将其导入 Eclipse 中。打开一个命令提示符或终端窗口，导航到 $GOPATH 并执行以下命令：



```php
git clone https://github.com/makotogo/ChaincodeTutorial.git
```

命令输出应类似于：



```bash
$ export GOPATH=/Users/sperry/home/mychaincode
$ cd $GOPATH
$ git clone https://github.com/makotogo/ChaincodeTutorial.git
Cloning into 'ChaincodeTutorial'...
remote: Counting objects: 133, done.
remote: Compressing objects: 100% (90/90), done.
remote: Total 133 (delta 16), reused 118 (delta 1), pack-reused 0
Receiving objects: 100% (133/133), 9.39 MiB | 1.95 MiB/s, done.
Resolving deltas: 100% (16/16), done.
$ cd ChaincodeTutorial
$ pwd
/Users/sperry/home/mychaincode/ChaincodeTutorial
```

此命令将 Blockchain ChaincodeTutorial 存储库从 GitHub 克隆到 $GOPATH。它包含一个 Java 链代码框架项目，您可以在本地区块链网络中构建、运行和测试它。

但在执行所有这些操作之前，需要将该代码导入 Eclipse 中。

#### 3.将该项目导入 Eclipse 中

在 Eclipse 中，转到 **File > Import...> Gradle > Existing Gradle Project**。这会打开一个向导对话框（参见图 9）。

![img](https:////upload-images.jianshu.io/upload_images/11831773-81385640a216ac36.png?imageMogr2/auto-orient/strip|imageView2/2/w/524/format/webp)

Eclipse Import Wizard：Gradle Project

单击 **Next**。在向导中随后出现的对话框中（参见图 10），浏览到 $GOPATH/ChaincodeTutorial，然后单击 **Finish** 导入该项目。

![img](https:////upload-images.jianshu.io/upload_images/11831773-6965c56fc19374b9.png?imageMogr2/auto-orient/strip|imageView2/2/w/643/format/webp)

Eclipse Import Wizard：Gradle Project（项目的 root 目录）



完成项目导入后，确保选择了 **Java Perspective**，您刚导入的 ChaincodeTutorial 项目会显示在 Project Explorer 视图中。

将代码导入 Eclipse 工作区后，就可以编写链代码了。

#### 4.探索该链代码框架项目

在本节中，将探索该链代码项目，以便理解在编写任何 Java 代码前它应该如何运行。

作为开发人员，我们喜欢编写代码，所以我不想让您失去编写 Java 代码的机会。但是，项目设置可能很复杂，我不想让这些设置阻碍实现本教程的主要目的。为此，我提供了您所需的大部分代码。

首先让我们快速查看一下基类 AbstractChaincode，它位于 com.makotojava.learn.blockchain.chaincode 包中，如清单 1 所示。

清单 1. AbstractChaincode 类



```tsx
package com.makotojava.learn.blockchain.chaincode;
 
import java.util.Arrays;
 
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.hyperledger.java.shim.ChaincodeBase;
import org.hyperledger.java.shim.ChaincodeStub;
 
public abstract class AbstractChaincode extends ChaincodeBase {
 
  private static final Log log = LogFactory.getLog(AbstractChaincode.class);
 
  public static final String FUNCTION_INIT = "init";
  public static final String FUNCTION_QUERY = "query";
 
  protected abstract String handleInit(ChaincodeStub stub, String[] args);
  protected abstract String handleQuery(ChaincodeStub stub, String[] args);
  protected abstract String handleOther(ChaincodeStub stub, String function, String[] args);
 
  @Override
  public String run(ChaincodeStub stub, String function, String[] args) {
    String ret;
    log.info("Greetings from run(): function -> " + function + " | args -> " + Arrays.toString(args));
    switch (function) {
    case FUNCTION_INIT:
      ret = handleInit(stub, args);
      break;
    case FUNCTION_QUERY:
      ret = handleQuery(stub, args);
    default:
      ret = handleOther(stub, function, args);
      break;
    }
    return ret;
  }
 
  @Override
  public String query(ChaincodeStub stub, String function, String[] args) {
    return handleQuery(stub, args);
  }
 
}
```

我想指出的第一点是，AbstractChaincode 是 ChaincodeBase 的子类，后者来自该结构的 shim 客户端（**第 7、10 行**）。

**第 17-19 行**显示了需要在 ChaincodeLog 类（AbstractChaincode 的子类）中实现的方法，这些方法分别用于实现初始化、账本查询和日志功能。

**第 22-36 行**显示了 ChaincodeBase 类（来自链代码 shim 客户端）的 run() 方法，我们可以在其中查看调用了哪个函数，以及该调用应委托给哪个处理函数。该类是可扩展的，因为 init 和 query 以外的其他任何函数（比如 log 函数）都由 handleOther() 处理，所以您还必须实现它。

现在打开 com.makotojava.learn.blockchain.chaincode 包中的 ChaincodeLog 类。

我只提供了一个框架供您填充 — 也就是说，我仅提供了编译它所需的代码。您需要编写剩余代码。您应该执行 JUnit 测试，然后会看到测试失败（因为还未编写实现）和失败的原因。换句话说，可以使用 JUnit 测试作为指导来正确地实现代码。

现在，如果感觉难以理解，不要担心；我在 com.makotojava.learn.blockchain.chaincode.solution 中提供了解决方案，以防您遇到阻碍（或者想根据参考来帮助完成实现）。

### 编写 Java 链代码

首先介绍一下在 ChaincodeLog 中实现链代码方法需要了解的一些背景。Java 链代码通过 ChaincodeStub 类与 Hyperledger Fabric 框架进行通信，另外需要记住，账本是区块链技术的透明性方面的核心。让智能合约（责任性）发挥其作用的是账本的状态，而链代码是通过 ChaincodeStub 来评估账本的状态。通过访问账本状态，可以实现一个智能合约（也即链代码）。

ChaincodeStub 上有许多方法可用于在账本的当前状态中存储、检索和删除数据项，但本教程仅讨论两个方法，它们用于存储和检索账本状态：

putState(String key, String value)— 将指定的状态值存储在账本中，该值被相应映射到指定的键。

getState()— 获取与指定键关联的状态值，并以字符串形式返回它。

为本教程编写代码时，只需在账本中存储或检索状态值，就会使用 putState() 或 getState() 函数。ChaincodeLog 类仅在账本中存储和检索值来实现其智能合约，所以实现这些方法只需知道该值即可。更复杂的链代码将使用 ChaincodeStub 中的其他一些方法（但这些方法不属于本教程的介绍范畴）。

我非常喜欢测试驱动开发 (TDD)，所以按照 TDD 的方式，我首先编写单元测试。继续运行它们，并观察它们的失败过程。在这之后，编写符合规范的代码，直到单元测试得到通过。单元测试的工作是确保能够获得预期的行为，通过研究单元测试，您将获得实现这些方法所需的足够信息。

但是，我还在每个方法顶部编写了 javadoc 注释，这可能有所帮助（以防您不熟悉 TDD 或 JUnit）。在学完本节的内容后，在 JUnit 测试中的代码与框架 ChaincodeLog 中的 javadoc 注释之间，你应该知道有实现链代码所需的所有信息。

从 Project Explorer 视图（在 Java 透视图中），导航到 ChaincodeLogTest 类，右键单击它并选择 Run As > Gradle Test。在它运行时，您会看到如图 11 所示的结果，其中显示了运行的所有 Gradle 任务的树结构。成功完成的任务在旁边会用一个复选标记进行指示。

![img](https:////upload-images.jianshu.io/upload_images/11831773-d7ececc487f0dc3a.png?imageMogr2/auto-orient/strip|imageView2/2/w/724/format/webp)

Eclipse：Gradle Executions 视图

Gradle Executions 选项卡中的感叹号表示与失败的单元测试对应的 Gradle 任务（跟我们期望的一样，所有 4 个单元测试都失败了）。

由于我们编写 JUnit 测试案例的方式，每个测试方法对应于 ChaincodeLog 中的一个方法，您需要在本教程中正确实现它们。

**实现 getChaincodeID()**
 首先，需要实现 getChaincodeID()。它的合约要求返回链代码的唯一标识符。我在 ChaincodeLog 类的顶部定义了一个名为 CHAINCODE_ID 的常量，您会用到它。可以自由更改它的值，但是，如果要更改 getChaincodeID() 返回的链代码 ID，请确保它在您的网络中是唯一的，而且不要忘记更改 JSON 消息的 ChaincodeID.name 属性。



```tsx
/**
 * Returns the unique chaincode ID for this chaincode program.
 */
@Override
public String getChaincodeID() {
  return null;// ADD YOUR CODE HERE
}
```

练习：完成 getChaincodeID() 方法。如果需要一个参考，请参见 com.makotojava.learn.blockchain.chaincode.solution 包。

**实现 handleInit()**

接下来将实现 handleInit() 方法。它的合约要求处理链代码程序的初始化，在本例中，这意味着它将向账本添加一条（由调用方指定的）消息，并在调用成功时将该消息返回给调用方。



```dart
/**
 * Handles initializing this chaincode program.
 *
 * Caller expects this method to:
 *
 * 1. Use args[0] as the key for logging.
 * 2. Use args[1] as the log message.
 * 3. Return the logged message.
 */
@Override
protected String handleInit(ChaincodeStub stub, String[] args) {
  return null;// ADD YOUR CODE HERE
}
```

练习：完成 handieInit() 方法。如果需要一个参考，请参见 com.makotojava.learn.blockchain.chaincode.solution 包。

**实现 handleQuery()**

接下来将实现 handleQuery() 方法。它的合约要求查询账本，为此，它会获取指定的键，在账本中查询与这个（这些）键匹配的值，然后将该（这些）值返回给调用方。如果指定了多个键，应该使用逗号分隔返回的值。



```dart
/**
 * Handles querying the ledger.
 *
 * Caller expects this method to:
 * 
 * 1. Use args[0] as the key for ledger query.
 * 2. Return the ledger value matching the specified key
 *    (which should be the message that was logged using that key).
 */
@Override
protected String handleQuery(ChaincodeStub stub, String[] args) {
  return null;// ADD YOUR CODE HERE
}
```

确保编写了代码来输出查询调用的结果，以便可以在控制台输出中查看结果（如果想了解我是如何做的，请参阅解决方案）。

练习：完成 handleQuery() 方法。如果需要一个参考，请参见 com.makotojava.learn.blockchain.chaincode.solution 包。

**实现 handleOther()**

最后需要实现 handleOther() 方法，它的合约要求处理其他消息（这是完全开放的，但正因如此它才是可扩展的）。您将在这里实现 log 函数，它的合同要求将调用方指定的一条消息添加到账本中，并在调用成功时将该消息返回给调用方。这看起来与 init 函数中发生的事非常相似，所以或许您可以在该实现中利用此函数。



```dart
/**
 * Handles other methods applied to the ledger.
 * Currently, that functionality is limited to these functions:
 * - log
 *
 * Caller expects this method to:
 * Use args[0] as the key for logging.
 * Use args[1] as the log message.
 * Return the logged message.
 */
@Override
protected String handleOther(ChaincodeStub stub, String function, String[] args) {
  // TODO Auto-generated method stub
  return null;// ADD YOUR CODE HERE
}
```

练习：完成 handleOther() 方法。如果需要一个参考，请参见 com.makotojava.learn.blockchain.chaincode.solution 包。

如果您为前面的每个练习编写的代码满足本节（以及代码注释中）为它们设定的要求，JUnit 测试应该都能通过，而且将链代码部署在本地区块链网络中并运行时，它们应该能够正常工作。

请记住，如果遇到阻碍，我提供了一个解决方案（但是在查看解决方案之前，您必须自行实现这些方法）。

### 构建 Java 链代码

现在您已编写 Java 链代码且通过了所有 JUnit 测试，是时候使用 Eclipse 和用于 Eclipse 的 **Gradle Buildship** 插件构建链代码了。通过转到 **Window > Show View > Other...** 调出 **Gradle Tasks** 视图，然后搜索 **gradle**，选择 **Gradle Tasks**，并单击 OK。（参见图 12。）

![img](https:////upload-images.jianshu.io/upload_images/11831773-fba410e627dba20d.png?imageMogr2/auto-orient/strip|imageView2/2/w/305/format/webp)

Eclipse：Show View：Gradle Tasks 视图

**Gradle Tasks** 视图打开后，展开 **ChaincodeTutorial > build** 节点，选择 **build** 和 **clean**。（参见图 13。）

![img](https:////upload-images.jianshu.io/upload_images/11831773-fe629bcba3422292.png?imageMogr2/auto-orient/strip|imageView2/2/w/651/format/webp)

Eclipse：Gradle Tasks 视图

右键单击 **build** 和 **clean**，然后选择 **Run Gradle Tasks**（Gradle 将确定运行它们的正确顺序）。您的 **Gradle Executions** 视图应该显示一个干净的构建版本，如图 14 所示，其中每项的旁边仅有一个复选标记。

![img](https:////upload-images.jianshu.io/upload_images/11831773-9203badbad380f6b.png?imageMogr2/auto-orient/strip|imageView2/2/w/1009/format/webp)

Eclipse：Gradle Executions 视图：干净构建



完成构建后，$GOPATH/ChaincodeTutorial 目录（您之前已从 GitHub 将代码克隆到这里）下有一个子目录 build/distributions，它包含您的链代码（这应该看起来很熟悉，因为本教程前面的 hello 示例中已经这么做过）。

构建 Java 链代码后，就可以在本地区块链网络中部署和运行它，并在它之上调用交易。

### 部署并运行 Java 链代码

在本节中，将会启动并注册您的链代码，部署它，并通过 Hyperledger Fabric REST 接口在链代码之上调用交易，就像本教程前面对 hello 示例所做的一样。确保本地区块链正在运行（如想温习一下相关内容，请参阅 “[启动区块链网络](https://www.ibm.com/developerworks/cn/java/j-chaincode-for-java-developers/blockchainnetwork)” 部分）。

您将执行以下步骤：

- 注册 Java 链代码。
- 部署 Java 链代码。
- 在 Java 链代码上调用交易。

#### 1.注册 Java 链代码

您需要提取 build/distributions/ChaincodeTutorial.zip 文件并运行链代码脚本，就像本教程前面运行 hello 示例时一样（参见 “[注册示例](https://www.ibm.com/developerworks/cn/java/j-chaincode-for-java-developers/registertheexample)” 部分）。

运行 ChaincodeTutorial 脚本时，输出应如下所示：



```ruby
$ ./ChaincodeTutorial/bin/ChaincodeTutorial
Feb 28, 2017 4:18:16 PM org.hyperledger.java.shim.ChaincodeBase newPeerClientConnection
INFO: Inside newPeerCLientConnection
Feb 28, 2017 4:18:16 PM io.grpc.internal.TransportSet$1 call
INFO: Created transport io.grpc.netty.NettyClientTransport@10bf86d3(/127.0.0.1:7051) for /127.0.0.1:7051
Feb 28, 2017 4:18:21 PM io.grpc.internal.TransportSet$TransportListener transportReady
INFO: Transport io.grpc.netty.NettyClientTransport@10bf86d3(/127.0.0.1:7051) for /127.0.0.1:7051 is ready
```

现在您的 Java 链代码已向本地区块链网络注册，您已准备好部署和测试链代码了。

#### 2.部署 Java 链代码

就像对 hello 示例链代码执行的操作一样，将会使用该结构的 REST 接口部署 Java 链代码，并在它之上调用交易。

打开 SoapUI。如果愿意的话，可以自行创建一个新 REST 项目和它的所有请求，或者可以导入我包含在之前克隆的 GitHub 项目中的 SoapUI REST 项目。该 SoapUI 项目位于 $GOPATH/ChaincodeTutorial 目录中。

要部署链代码，可以导航到 ChaincodeLog Deploy 请求（如图 15 所示）并提交该请求。

![img](https:////upload-images.jianshu.io/upload_images/11831773-7a0f43576a5e33c3.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

SoapUI：ChaincodeLog Deploy 请求

如果没有使用来自 GitHub 的 SoapUI 项目（或者使用不同的 HTTP 客户端），那么应该提交的 JSON 请求如下所示：



```json
{
"jsonrpc": "2.0",
  "method": "deploy",
  "params": {
    "type": 4,
    "chaincodeID":{
        "name": "ChaincodeLogSmartContract"
    },
    "ctorMsg": {
        "args": ["init", "KEY-1", "Chaincode Initialized"]
    }
  },
  "id": 1
}
```

提交请求。如果请求被成功处理，您会获得以下 JSON 响应：



```json
{
   "jsonrpc": "2.0",
   "result":    {
      "status": "OK",
      "message": "ChaincodeLogSmartContract"
   },
   "id": 1
}
```

现在您的链代码已部署并准备好运行。

#### 3.在 Java 链代码上调用交易

部署并初始化 Java 链代码后，就可以在它之上调用交易了。在本节中，将会调用 log 和 query 函数作为交易。

要调用 log 函数，可以打开 ChaincodeLog Log 请求并提交它。（参见图 16。）

![img](https:////upload-images.jianshu.io/upload_images/11831773-5145f693421a94e1.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

SoapUI：ChaincodeLog Log 请求

如果没有使用来自 GitHub 的 SoapUI 项目（或者使用不同的 HTTP 客户端），那么应该提交的 JSON 请求如下所示：



```json
{
"jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "type": 1,
    "chaincodeID":{
        "name": "ChaincodeLogSmartContract"
    },
    "CtorMsg": {
        "args": ["log", "KEY-2", "This is a log message."]
    }
  },
  "id": 2
}
```

如果请求被成功处理，您会获得以下 JSON 响应：



```json
{
   "jsonrpc": "2.0",
   "result":    {
      "status": "OK",
      "message": "a6f7a4fc-2980-4d95-9ec2-114dd9d0e4a5"
   },
   "id": 2
}
```

要调用 query 函数，可以打开 ChaincodeLog Query 请求并提交它。（参见图 17。）



![img](https:////upload-images.jianshu.io/upload_images/11831773-b4090ef3107cba28.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

SoapUI：ChaincodeLog Query 请求

如果没有使用来自 GitHub 的 SoapUI 项目（或者使用不同的 HTTP 客户端），那么应该提交的 JSON 请求如下所示：



```json
{
"jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "type": 1,
    "chaincodeID":{
        "name": "ChaincodeLogSmartContract"
    },
    "ctorMsg": {
        "args": ["query", "KEY-1", "KEY-2"]
    }
  },
  "id": 3
}
```

如果请求被成功处理，您会获得以下 JSON 响应：



```json
{
   "jsonrpc": "2.0",
   "result":    {
      "status": "OK",
      "message": "84cbe0e2-a83e-4edf-9ce9-71ae7289d390"
   },
   "id": 3
}
```

解决方案代码的终端窗口输出类似于：



```bash
$ ./ChaincodeTutorial/bin/ChaincodeTutorial
Feb 28, 2017 4:18:16 PM org.hyperledger.java.shim.ChaincodeBase newPeerClientConnection
INFO: Inside newPeerCLientConnection
Feb 28, 2017 4:18:16 PM io.grpc.internal.TransportSet$1 call
INFO: Created transport io.grpc.netty.NettyClientTransport@10bf86d3(/127.0.0.1:7051) for /127.0.0.1:7051
Feb 28, 2017 4:18:21 PM io.grpc.internal.TransportSet$TransportListener transportReady
INFO: Transport io.grpc.netty.NettyClientTransport@10bf86d3(/127.0.0.1:7051) for /127.0.0.1:7051 is ready
Feb 28, 2017 4:34:52 PM com.makotojava.learn.blockchain.chaincode.AbstractChaincode run
INFO: Greetings from run(): function -> init | args -> [KEY-1, Chaincode Initialized]
Feb 28, 2017 4:34:52 PM com.makotojava.learn.blockchain.chaincode.solution.ChaincodeLog handleLog
INFO: *** Storing log message (K,V) -> (ChaincodeLogSmartContract-CLSC-KEY-1,Chaincode Initialized) ***
Feb 28, 2017 4:50:27 PM com.makotojava.learn.blockchain.chaincode.AbstractChaincode run
INFO: Greetings from run(): function -> log | args -> [KEY-2, This is a log message.]
Feb 28, 2017 4:50:27 PM com.makotojava.learn.blockchain.chaincode.solution.ChaincodeLog handleLog
INFO: *** Storing log message (K,V) -> (ChaincodeLogSmartContract-CLSC-KEY-2,This is a log message.) ***
Feb 28, 2017 5:02:13 PM com.makotojava.learn.blockchain.chaincode.AbstractChaincode run
INFO: Greetings from run(): function -> query | args -> [KEY-1, KEY-2]
Feb 28, 2017 5:02:13 PM com.makotojava.learn.blockchain.chaincode.solution.ChaincodeLog handleQuery
INFO: *** Query: For key 'ChaincodeLogSmartContract-CLSC-KEY-1, value is 'Chaincode Initialized' ***
Feb 28, 2017 5:02:13 PM com.makotojava.learn.blockchain.chaincode.solution.ChaincodeLog handleQuery
INFO: *** Query: For key 'ChaincodeLogSmartContract-CLSC-KEY-2, value is 'This is a log message.' ***
```

恭喜您！您已向未来迈出了第一步。

鼓励您执行以下操作：修改 ChaincodeTutorial 项目，向它添加方法，更改实现，等等。您也可以自由地编写链代码。祝您好运，编码愉快！

### 结束语

本教程简要概述了区块链技术和智能合约（实现为链代码程序），以及最新的区块链技术的发展形势。

我们介绍了设置 Java 链代码开发环境的步骤，包括需要安装的软件，如何定义和运行本地区块链网络，以及如何部署来自 GitHub 中的 Hyperledger Fabric 项目的一个 Java 链代码示例程序并在它之上调用交易。

您学习了如何使用 Eclipse、JUnit 和 Gradle 编写和构建第一个 Java 链代码程序，然后部署该 Java 链代码程序并在它之上调用交易。

您亲自查看了区块链技术和智能合约，随着区块链技术发展日渐成熟和市场规模逐渐扩大，您会掌握更多的技巧来编写更复杂的 Java 链代码。

那么您接下来会怎么做？

### 后续行动

以下建议可帮助您在目前所学知识的基础上继续进行研究：

[深入研究 Hyperledger Fabric 架构](http://hyperledger-fabric.readthedocs.io/en/latest/arch-deep-dive.html)

#### 致谢

非常感谢杜婧细心评审本文，提供建设性意见并进行校正。

> 如果你希望**高效的**学习以太坊DApp开发，可以访问汇智网提供的**最热门**在线互动教程：
>
> - [适合区块链新手的以太坊智能合约和DApp实战入门教程](http://xc.hubwiz.com/course/5a952991adb3847553d205d1?affid=20180521jianshuw)
> - [区块链+IPFS+Node.js+MongoDB+Express去中心化以太坊电商应用开发实战](http://xc.hubwiz.com/course/5abbb7acc02e6b6a59171dd6?affid=20180521janshuw)
>
> 其他更多内容也可以访问[这个以太坊博客](http://blog.hubwiz.com)。

原文：
 https://www.ibm.com/developerworks/cn/java/j-chaincode-for-java-developers/index.html

作者： [J Steven Perry](https://developer.ibm.com/author/steve.perry/)



作者：编程狂魔
链接：https://www.jianshu.com/p/2f7c11d0c493
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。