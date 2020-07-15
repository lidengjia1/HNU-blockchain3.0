# Openstack Vitrage（根因分析）

发表于 2017-11-10 | 分类于 [Openstack op](https://www.backendcloud.cn/categories/Openstack-op/)

# Vitrage简介

Vitrage(平台问题定位分析服务)。Vitrage是一个OpenStack RCA(Root Cause Analysis)服务，用于组织、分析和扩展OpenStack的告警和事件，在真正的问题发生前找到根本原因。

众所周知，OpenStack平台最大的优势来自于架构的可扩展性，这也是OpenStack能够在基础架构曾一枝独秀的重要原因。分布式架构最大的优势在于扩展，但是过于灵活的扩展性为运维带来的极大的困难，所以Vitrage的出现在一定程度上缓解了OpenStack运维上的痛点。

# 功能简介：

1.物理-虚拟实体映射
2.推导告警和状态（例如：基于系统分析后产生告警或者修改状态，而不是直接监控）
3.对告警、事件进行根因分析
4.Horizon显示

# 架构

[![pic_4](https://www.backendcloud.cn/images/vitrage/4.png)](https://www.backendcloud.cn/images/vitrage/4.png)

Vitrage Data Source(s)：
负责从不同来源导入关于系统状态的信息。这些信息包括物理资源、虚拟资源的告警和状态，这些信息将会由vitrage Graph处理。目前，vitrage已经支持的数据源有：Nova、Cinder、Aodh、Nagios 告警以及静态物理资源等。

Vitrage Graph：
保存数据源收集的信息以及其内部关系。另外，它实现了一些vitrage evaluator使用的基本图算法(例如：sub-matching, BFS, DFS等)。

Vitrage Evaluator：
协调vitrage Graph进行分析（修改）并处理分析结果。负责执行vitrage中定义的不同类型模板的动作，例如添加一个告警的根因分析，产生一个推导后的告警或者设置一个推导后的状态。

Vitrage Notifiers：
用来通知外部系统Vitrage的告警和状态。目前支持用Aodh notifier来产生vitrage告警，以及用Nova notifier来标注主机down等。

# 服务

Vitrage-graph服务主程序：包含了in-memory 实体图、模板分析程序以及datasource分析程序等。

Vitrage-collector服务：负责从不同数据源获取资源信息，包括主动获取数据源的告警和状态等资源，被动接收数据源发送过来的信息。负责将取得的信息解析成entity信息传给Vitrage-graph使用。

Vitrage-notifier服务：用来通知外部系统Vitrage的告警或者状态变化。在Ocata版本中只支持通知Nova force-down API以及SNMP。(参考另一篇文章 [新增OpenStack NOVA API用于强制计算节点nova-compute服务down](https://www.backendcloud.cn/2017/06/08/force-down/))

Vitrage-api进程：Vitrage的API层，以进程形式表示Vitrage图形显示、vitrage CLI将调用vitrage-api。

# Entity Graph和Evaluator

Vitrage将系统状态（资源和告警）生成对应的图

- 实体在图上对应的是顶点，实体间的关系在图上对应的是顶点相互连接的边。
- 每一个顶点和边可以有额外的属性
- 每一个边对应一个“lable”，表示关系类型
- 展示数据关系和相互影响的直观模型
  [![pic_1](https://www.backendcloud.cn/images/vitrage/1.png)](https://www.backendcloud.cn/images/vitrage/1.png)

Vitrage Evaluator监听图上事件的变化

- 查找事件相关联的模板（脚本）
- 评估图上的状态以及状态改变
- 根据对图的状态以及状态改变的评估执行相应的动作

# Use Cases for Vitrage：新增Nova实例

[![pic_5](https://www.backendcloud.cn/images/vitrage/5.png)](https://www.backendcloud.cn/images/vitrage/5.png)
1) Nova datasource Driver查询所有的nova实例，或者获取消息队列通知，得知新增了一个nova实例。
2) Nova datasource Driver向Entity Queue发出对应的事件。
3) The Entity Processor从Entity Queue获得新增Nova实例事件。
4) The Entity Processor将事件传递给Nova Instance Transformer
5) 在图上显示最新的Nova实例，并和相应的Host建立连接关系。
[![pic_6](https://www.backendcloud.cn/images/vitrage/6.png)](https://www.backendcloud.cn/images/vitrage/6.png)

# Use Cases for Vitrage：新增Aodh告警

流程类似上面的新增Nova实例，不详细说明了，直接上图。
[![pic_7](https://www.backendcloud.cn/images/vitrage/7.png)](https://www.backendcloud.cn/images/vitrage/7.png)
[![pic_8](https://www.backendcloud.cn/images/vitrage/8.png)](https://www.backendcloud.cn/images/vitrage/8.png)

# Use Cases for Vitrage：Nagios主机物理网卡故障推导出主机上的实例故障

[![pic_9](https://www.backendcloud.cn/images/vitrage/9.png)](https://www.backendcloud.cn/images/vitrage/9.png)
[![pic_10](https://www.backendcloud.cn/images/vitrage/10.png)](https://www.backendcloud.cn/images/vitrage/10.png)

通过Nagios, Zabbix监视物理网卡的状态，一旦故障给出告警。

Vitrage通过监视工具得到故障告警，并将告警加入实体图，下图中的Host NIC

找到对应的脚本（模板）并执行以下的动作：

```
给出推导的有关主机的告警
    同样将此告警加入实体图
改变Vitrage的主机状态
    可以通过调用Nova API来改变主机的状态(参考另一篇文章  新增OpenStack NOVA API用于强制计算节点nova-compute服务down)
增加告警的因果关系
```

一旦推导的主机告警成立，相同的流程会应用到主机上的vm实例和相关的VNF
[![pic_2](https://www.backendcloud.cn/images/vitrage/2.png)](https://www.backendcloud.cn/images/vitrage/2.png)

# Use Cases for Vitrage：RCA根本原因分析

[![pic_11](https://www.backendcloud.cn/images/vitrage/11.png)](https://www.backendcloud.cn/images/vitrage/11.png)
1) Evaluator被通知一个新的告警Alarm-X（如下图的Instance At Risk1）
2) Evaluator评估模板和图，推导出在Alarm-X and Alarm-Y间有一个根本原因，在图上增加一个Alarm-Y指向Alarm-X的箭头。
[![pic_12](https://www.backendcloud.cn/images/vitrage/12.png)](https://www.backendcloud.cn/images/vitrage/12.png)