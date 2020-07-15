#  印地世界演示

Indy World Demo实现了[Indy入门中](https://github.com/hyperledger/indy-node/blob/stable/getting-started.md)描述的**Alice Demo**。该演示演示了如何使用自主权身份（“ SSI”）来从发行方获取凭据，并响应于证明请求将这些凭据提供给验证者，从而提供可验证的证明。

此仓库提供了动手教程以及带说明的[视频录制](https://www.youtube.com/watch?v=cz-6BldajiA) [（文件镜像）](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/video/IndyWorldVideo2.mp4)。

> 注意：文件镜像视频需要兼容mp4的视频播放器。

## 致谢

作为主动管理的一部分，[具有自我主权身份可验证组织](https://von.pathfinder.gov.bc.ca/)，政府为[不列颠哥伦比亚省](https://www2.gov.bc.ca/)（“BC政务”）的合作伙伴关系与[公共服务和采购加拿大](https://www.canada.ca/en/public-services-procurement.html)和[安大略政府](https://www.ontario.ca/page/government-ontario)，贡献了几个[开源](https://github.com/bcgov)有助于引导使用[Hyperledger Indy](https://wiki.hyperledger.org/projects/indy)（“ Indy”）和[可验证凭证的](https://www.w3.org/TR/verifiable-claims-data-model/)新应用程序的[项目](https://github.com/bcgov)。该Indy World Demo源自[allowify](https://github.com/bcgov/permitify)和[von-network](https://github.com/bcgov/von-network)存储库。

## 总览

Indy World Demo旨在作为有兴趣学习Hyperledger Indy环境中如何交换可验证凭证的任何人的免费教程。

### 爱丽丝的数字证书故事

了解您的客户（“ KYC”）是与银行和其他机构用来识别和验证其客户身份的监管活动相关的业务流程。大学毕业后，Alice希望在Thrift Bank开一个新帐户。该银行提供了新的自助服务，新客户可以快速无缝地使用数字凭证来满足银行的KYC要求。

幸运的是，爱丽丝最近利用了母校法伯学院提供的新数字笔录。在她的新雇主Acme Corporation（“ Acme”）的申请过程中，该数字笔录非常方便。被录用后，Acme向爱丽丝颁发了一份数字工作证书。爱丽丝使用她的包含教育和就业证明的数字钱包，充满信心，可以利用Thrift Bank的新自助服务产品。

### 可验证的凭证角色

给定“ [可验证凭据”](https://www.w3.org/TR/verifiable-claims-data-model/)规范中概述的角色和职责，可以使用以下抽象概念来描述Alice的KYC故事：

| 角色           | 描述                                                       | 例子                                                         |
| -------------- | ---------------------------------------------------------- | ------------------------------------------------------------ |
| *持有人*       | 控制一个或多个可验证凭据的实体。                           | 持有人可以是个人，组织或连接的设备。爱丽丝是持有人。         |
| *发行人*       | 创建可验证凭证，将其与特定主题相关联并传送给持有人的实体。 | 发行人包括公司，政府和个人。Faber College和Acme是发行人。    |
| *检验员*       | 接收一个或多个可验证凭证（由证明请求指定）进行处理的实体。 | 验证者包括雇主，安全人员和网站。Acme和Thrift Bank是验证者。  |
| *标识符注册表* | 调解主题标识符的创建和验证。                               | 标识符注册表包括公司员工数据库，政府ID数据库和分布式分类帐。出于演示目的，我们使用由Indy节点池组成的分布式分类帐。 |

![vc角色](https://w3c.github.io/vc-data-model/diagrams/ecosystem.svg)

## 建筑

该Indy World Demo利用BC-Gov提供的三个存储库，使一个人可以注册一个虚构的政府ID，然后使用该ID获得大学成绩单，申请并被公司雇用并申请贷款。

存储库是：

[VON网络](https://github.com/bcgov/von-network)：VON网络创建一个4节点的Indy节点池。每个节点都包含一个区块链分类帐。这些分类帐保持同步并使用共识算法进行验证。

[组织书](https://github.com/bcgov/TheOrgBook)：组织书是企业及其相关凭据和许可证的存储库。对其进行了修改，以持有个人可验证的凭据。这些证书是由发行人（如大学，公司和银行）发行的。

[许可](https://github.com/bcgov/permitify)：许可实例化演示的发行者。有四种许可发行人服务：

- 政府SSI应用，可发布*政府ID*
- 费伯学院SSI应用程序，可发出*学生成绩单*
- Acme SSI应用程序，可发布*就业证明*
- 储蓄银行SSI应用程序，可发行*银行帐户*

这些发行者中的每个发行者都为由凭证架构定义描述的个人创建可验证的凭证。在颁发证书之前，必须使用个人当前持有的证书填写证明请求，并由机构进行验证。

下图显示了该演示的各种组件。

![演示架构](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/architecture.png)

[Indy SDK](https://github.com/hyperledger/indy-sdk)：Indy是在分布式分类帐之上的自主权身份生态系统的实现。

[钱包](https://sovrin.org/wp-content/uploads/2018/03/Sovrin-Provisional-Trust-Framework-2017-06-28.pdf)：安全存储密钥，凭证和其他由持有者，发行者，检查者验证者或身份证明者注册表使用的私人数据。

[Indy节点](https://sovrin.org/wp-content/uploads/2018/03/Sovrin-Provisional-Trust-Framework-2017-06-28.pdf)：运行Sovrin开源代码（Indy）实例以维护Sovrin分类帐的计算机网络服务器。

[VON代理](https://github.com/PSPC-SPAC-buyandsell/von_agent)：BC-Gov创建的高级API，用于包装Indy SDK。

## 使用印地世界

### 依存关系

**Docker**和**Docker Compose**用于构建和运行演示。要查看系统上当前安装了哪些版本，请运行：

```
$ docker --version 
$ docker-compose --version
```

可在[此处](https://www.docker.com/get-docker)找到有关为各种平台下载和安装Docker的信息。

还需要使用**源到映像**（s2i）工具来构建演示中使用的docker映像。S2I可以在[这里](https://github.com/openshift/source-to-image)下载。如果您使用的是Mac，并且已安装Homebrew，则以下命令将安装s2i：

```
$ brew install源到映像
```

该网站提供了在其他平台上安装的说明。

验证**s2i**在您的PATH中。如果没有，请编辑PATH并添加**s2i**的安装目录。该**manage.sh**脚本将寻找**S2I**您PATH可执行文件。如果找不到，您将收到一条消息，要求您下载并在PATH上进行设置。

### 正在安装

该项目包括来自BC-Gov的三个git存储库（VON Network，TheOrgBook和Permitify），如体系结构部分所述。

克隆indy-world仓库将获取构建和运行Indy World演示所需的BC-Gov仓库定制版本的所有代码。

```
$ git clone https://github.com/IBM-Blockchain-Identity/indy-ssivc-tutorial.git 
$ cd indy-ssivc-tutorial
```

### 建造

您需要构建组件VON Network，TheOrgBook和Permitify。

要从网络构建：

```
$ cd from-network 
$ ./manage build
```

要构建TheOrgBook：

```
$ cd TheOrgBook / 
docker $ ./manage.sh构建
```

要构建许可：

```
$ cd allowify / 
docker $ ./manage build
```

您可以使用'rebuild'选项重建docker镜像。

### 跑步

VON Network需要首先启动。

要启动VON Network，请打开一个新终端并输入以下命令：

```
$ cd from-network 
$ ./manage start
```

等待直到所有节点都完全启动，然后再启动TheOrgBook。

接下来需要启动TheOrgBook。

要启动TheOrgBook，请打开一个新终端并输入以下命令：

```
$ cd TheOrgBook / 
docker $ ./manage.sh开始
```

等待直到TheOrgBook完全启动，然后再启动Permitify。

要启动Permitify，请打开一个新终端并输入以下命令：

```
$ cd allowify / 
docker $ ./manage开始全部
```

启动所有容器后，您可以通过在浏览器中访问以下网页来查看节点池和分类帐：

要查看节点池（VON网络）：

```
http：//本地主机：9000 /
```

VON Network Web UI可以检查分类帐。创世文件也可以查看。

![冯网络](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/von-network.png)

您可以通过在浏览器中访问以下网页来运行演示：

要运行演示：

```
http：// localhost：8080 /
```

### 正在停止

要停止应用程序，请按Ctrl-C。这应该停止命令启动的所有docker容器。

如果仍有容器在运行，则可以使用以下命令停止它们：

来自网络：

```
$ cd from-network 
$ ./manage stop
```

TheOrgBook：

```
$ cd TheOrgBook / 
docker $ ./manage.sh停止
```

允许：

```
$ cd allowify / 
docker $ ./管理停止
```

### 打扫干净

可以使用**manage rm**参数删除运行演示时创建的docker容器和卷。如果您在添加用户或验证凭据时遇到问题，也可以调用此命令来重置用于演示的数据。

来自网络：

```
$ cd from-network 
$ ./manage rm
```

TheOrgBook：

```
$ cd TheOrgBook / 
docker $ ./manage.sh rm
```

允许：

```
$ cd allowify / 
docker $ ./manage rm
```

## 演示版

要运行该演示，请在浏览器中加载主页：

```
http：// localhost：8080 /
```

主页以可验证凭据强制执行的顺序显示了演示步骤的概述。例如，在证明自己拥有学位之前，不能雇用一个人获得工作证书。学位的验证是通过使用先前获得并经过验证的凭证（法伯学院）的信息填写证明请求（Acme Corp）来完成的。

该演示包括5个步骤：

1. 注册一个人
2. 获得大学成绩单
3. 申请工作
4. 受雇用
5. 申请贷款

![主页](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/homepage.png)

要注册人员并创建新的在线政府身份，请单击“ **注册人员”**按钮。

注册新人员后，将显示证书发行者列表。要从这些发行人之一获取凭证，请单击相应的按钮。请注意，只有那些满足所有依赖性的发行者才能被成功调用。

### 注册一个人

通过单击“ **注册人”**按钮，爱丽丝可以获得政府ID 。这将显示在Gov ID网站上托管的表单，该表单可以由政府员工亲自审核爱丽丝后填写，也可以由爱丽丝在使用她已经用于访问其他政府服务的凭据登录Gov ID网站后填写。无论哪种方式，这都是Alice使用可验证凭据的入门。

![gov_id](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/gov_id.png)

按下“ **提交”**按钮后，将通过Gov ID生成凭据，并将其下载到Alice的钱包中。请注意，爱丽丝的凭据也保存在发行者（Gov ID）钱包中。

![list_gov](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/list_gov.png)

现在，爱丽丝拥有证书，她可以使用该证书向其他发行人证明自己是谁。

现在，该演示将显示Alice获得了哪些凭据，以及在从特定发行者请求新凭据之前需要哪些凭据。

该网页显示，爱丽丝具有政府ID凭证，并且由于她的钱包中所有凭证都**取决于**她，现在可以向Faber College要求成绩单。请注意，由于尚未满足笔录依存关系，如红色X所示，她还不能申请工作。

在爱丽丝（Alice）向费伯学院（Faber College）要求其成绩单之前，她想从政府ID中查看其证书。单击“政府ID”部分中的“ **查看证书”**，将显示Alice的凭据。可以通过单击“ **验证凭据”**按钮来验证此凭据。

![alice_gov_id](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/alice_gov_id.png)

这会从Alice创建一个证明请求，并使用Alice钱包中的Gov ID凭据中的信息来填充该请求。它经过密码验证并确定为有效。“ **验证凭证”**按钮下方的内容显示了凭证请求的详细信息。这仅是出于演示目的，通常不会向用户显示。

单击“ **返回到个人信息”**链接将显示爱丽丝已在其钱包中收集的各种凭据。此时，只有一个-Gov ID凭证。

![alice_gov_id](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/alice_one.png)

单击“上**一步”**返回到发行者列表。

### 获得大学成绩单

现在，在凭据列表中，爱丽丝单击“ **获取大学成绩单”**。此操作可能需要20秒或更长时间。在后台，Alice从Faber College获得了证明请求定义，自动填写了她的ID凭证，将其发送给Faber College，然后Faber College通过密码验证了证明内容。完成此过程后，将从Faber College网站显示一个表格。通常，此表格将由注册服务机构填写；但是，出于演示目的，我们现在将其填写。

注意表单底部显示的索赔数据。这是Alice的政府ID凭证从她的钱包提供的经过验证的数据，该凭证是响应Faber College的证明要求而提供的。

当**提交**被按下时，麦嘉华学院创建副本凭证，并将其发送给Alice。然后，爱丽丝将其保存在她的钱包中。

![faber_college](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/faber_college.png)

### 申请工作

要在Acme Corp申请工作，Alice单击“ **申请工作”**。这将显示“职位申请”表单，其中大部分信息是通过爱丽丝在其钱包中持有的凭据填写的。

单击**提交**将创建一个工作申请凭证，该凭证将同时保存在爱丽丝和Acme Corp的钱包中。Acme将使用此经过验证的凭据信息作为面试过程的一部分，以决定是否要雇用Alice。

![acme_application](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/acme_application.png)

### 受雇用

好消息！爱丽丝已经被Acme Corp聘用。来自Acme人力资源部门的人将单击“ **工作证书颁发证书”**按钮，以创建爱丽丝证明其受雇的证书。

可能会通过电子邮件或在其员工网站上向Acice通知Alice有可用的工作凭证。无论哪种方式，爱丽丝都可以将此凭据添加到自己的钱包中，以验证其工作身份。对于此演示，通过单击“ **提交”**按钮将作业凭据添加到爱丽丝的钱包中。

![acme_hired](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/acme_hired.png)

### 申请贷款

现在，爱丽丝有了工作，她可以向储蓄银行申请贷款。她单击“ **申请贷款”**以显示来自Thrift Bank的贷款申请。

贷款申请表中已完全预填充了来自爱丽丝钱包中多个凭证的经过验证的信息。

请注意，爱丽丝有机会仅编辑和披露她希望储蓄银行知道的信息。在实际的应用程序中，每个输入框都将允许Alice选择要用于该特定数据的凭据。

爱丽丝单击**提交**按钮以申请贷款。

![储蓄银行](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/thrift_bank.png)

### 查看所有凭证

爱丽丝已从演示中的发行者列表中获取了所有凭据。她可以通过单击“ **查看证书”**按钮来查看任何凭据。

![list_all](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/list_all.png)

单击**查看人员记录**按钮，可以查看爱丽丝钱包中所有的凭据。

如果需要，可以查看和验证各个凭据。

![list_alice](https://gitee.com/acrowise/indy-ssivc-tutorial/raw/master/docs/images/alice_all.png)

## 休息码

### BC-Gov代码存储库

- [准许](https://github.com/bcgov/permitify)
- [TheOrgBook](https://github.com/bcgov/TheOrgBook)
- [来自网络](https://github.com/bcgov/von-network)

### Indy代码存储库

- [印地节点](https://github.com/hyperledger/indy-node)
- [Indy SDK](https://github.com/hyperledger/indy-sdk)
- [印地](https://github.com/hyperledger/indy-plenum)
- [印地加密](https://github.com/hyperledger/indy-crypto)
- [印地·阿南·克雷德斯](https://github.com/hyperledger/indy-anoncreds)