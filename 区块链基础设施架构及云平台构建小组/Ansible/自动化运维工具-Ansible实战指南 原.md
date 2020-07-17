## 自动化运维工具-Ansible实战指南 原

[**家** 家住海边喜欢浪](https://my.oschina.net/u/3413282) 发布于 2017/04/10 16:55

 

字数 11366

 

阅读 4K

 

 收藏 2

 

点赞 0

 

[ 评论 0](https://my.oschina.net/u/3413282/blog/876231#comments)



[ECharts5.0版本即将上线，来说说我与ECharts的那些事吧！>>> ![img](https://www.oschina.net/img/hot3.png)](https://www.oschina.net/action/visit/ad?id=1199)



#  



# 一：Ansible基础介绍



## 1、背景

>   在日常服务器运维中，我们经常要配置相同的服务器配置，前期我们都是一台一台的去配置，这种方法操作主要应对于服务器数量不多且配置简单的情况还可以继续这样操作，如果我们后期维护几百服务器或者几万服务器呢？ 我应该怎样去快速配置服务器呢？如果需要手动的每台服务器进行安装配置将会给运维人员带来许多繁琐而又重复的工作同时也增加服务器配置的异常，至此自动化运维工具解决我们的瓶颈---Ansible工具。



## 2、Ansible简介

  Ansible是一款基于Python开发的自动化运维工具，主要是实现批量系统配置、批量程序部署、批量运行命令、批量执行任务等等诸多功能。Ansible是一款灵活的开源工具，能够很大程度简化运维中的配置管理与流程控制方式，它利用推送方式对客户系统加以配置，这样所有工作都可在主服务器端完成。Asible是基于模块工作的，其本身没有批量部署的能力，总之只要明白Ansible是一款运维自动化的神器就好了~！

**(****上面的自动化工具Ansible基本都可以实现，自动安装系统正在研发中。。。)**



## 3、Ansible特性

```
(1).no agents：不需要在被管控主机上安装任何客户端；
(2).no server：无服务器端，使用时直接运行命令即可；
(3).modules in any languages：基于模块工作，可使用任意语言开发模块；
(4).yaml，not code：使用yaml语言定制剧本playbook；
(5).ssh by default：基于SSH工作；
(6).strong multi-tier solution：可实现多级指挥。
```



## 4、Ansible优点

```
(1).轻量级，无需在客户端安装agent，更新时，只需在操作机上进行一次更新即可；
(2).批量任务执行可以写成脚本，而且不用分发到远程就可以执行；
(3).使用python编写，维护更简单，ruby语法过于复杂；
(4).支持sudo。
```



## 5、Ansible 与其它配置管理的对比

选择了目前几款主流的与 Ansible 功能类似的配置管理软件 Puppet、Saltstack，这里所做的对比不针对各个软件的性能作比较，只是对各个软件的特性做个对比。

|                              | **Puppet**                                              | **Saltstack**                                | **Ansible**                                     |
| ---------------------------- | ------------------------------------------------------- | -------------------------------------------- | ----------------------------------------------- |
| 开发语言                     | Ruby                                                    | Python                                       | Python                                          |
| 是否有客户端                 | 有                                                      | 有                                           | 无                                              |
| 是否支持二次开发             | 不支持                                                  | 支持                                         | 支持                                            |
| 服务器与远程机器是否相互验证 | 是                                                      | 是                                           | 是                                              |
| 服务器与远程机器通信是否加密 | 是，标准 SSL 协议                                       | 是，使用 AES 加密                            | 是，使用 OpenSSH                                |
| 平台支持                     | 支持 AIX、BSD、HP-UX、Linux、 MacOSX、Solaris、 Windows | 支持 BSD、Linux、Mac OS X、Solaris、 Windows | 支持 AIX、BSD、 HP-UX、 Linux、Mac OSX、Solaris |
| 是否提供 web                 | ui   提供                                               | 提供                                         | 提供，不过是商业版本                            |
| 配置文件格式                 | Ruby 语法格式                                           | YAML                                         | YAML                                            |
| 命令行执行                   | 不支持，但可通过配置模块实现                            | 支持                                         | 支持                                            |



## **6、Ansible****基本架构**

  Ansible 是一个模型驱动的配置管理器，支持多节点发布、远程任务执行。默认使用 SSH 进行远程连接。无需在被管理节点上安装附加软件，可使用各种编程语言进行扩展。

![img](https://static.oschina.net/uploads/space/2017/0410/122524_4dcg_3413282.png)



## 7、**Ansible** **工作机制**

Ansible 在管理节点将 Ansible 模块通过 SSH 协议（或者 Kerberos、LDAP）推送到被管理端执 行，执行完之后自动删除，可以使用 SVN 等来管理自定义模块及编排

![img](https://static.oschina.net/uploads/space/2017/0410/122544_hMKH_3413282.png)

```
1、管理端支持local 、ssh、zeromq 三种方式连接被管理端，默认使用基于ssh的连接－－－这部分对应基本架构图中的连接模块；
2、可以按应用类型等方式进行Host Inventory（主机群）分类，管理节点通过各类模块实现相应的操作－－－单个模块，单条命令的批量执行，我们可以称之为ad-hoc；
3、管理节点可以通过playbooks 实现多个task的集合实现一类功能，如web服务的安装部署、数据库服务器的批量备份等。playbooks我们可以简单的理解为，系统通过组合多条ad-hoc操作的配置文件 。
```

![img](https://static.oschina.net/uploads/space/2017/0410/122610_Btfb_3413282.png)

**ansible****的任务执行流程**

```
（1）读取配置
（2）抓取全量机器&分组列表
      可从多个静态文件、文件夹、脚本中读取机器，分组及其变关联量信息。
（3）使用host-pattern过滤机器列表
（4）根据参数确定执行模块和配置
       从modules目录动态读取，用户可以自行开发模块。
（5）Runner执行返回
       Connection环节定义连接方式 => Action阶段机器列表（Lookup plugin Action变量/文件等资源的获取）
       Callback plugin各阶段的钩子调用
（6）输出结束
       Filter plugin过滤算子
       Callback plugin各阶段的钩子调用
```



# 二：Ansible安装入门



## 1、Ansible安装

```
[root@Ansible ~]# wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-6.repo
[root@Ansible ~]# wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-6.repo
[root@Ansible ~]# yum -y install ansible
```



## 2、配置文件

```
主配置文件
/etc/ansible/ansible.cfg
主机清单：
/etc/ansible/hosts
角色目录:
 /etc/ansible/roles
插件目录：
/usr/share/ansible_plugins/
```



## 3、命令格式

```
#常用格式
ansible <host-pattern> [-f forks] [-m module] [-a args]
host-pattern # 可以是all，或者配置文件中的主机组名
-f forks # 指定并行处理的进程数
-m module # 指定使用的模块，默认模块为command
-a args # 指定模块的参数
#查看各模块
ansible-doc [options] [modules]
# 主要选项有：
-l或--list # 列出可用的模块
-s或--snippet #显示指定模块的简略使用方法
```



## 4、Ansible命令参数介绍

```
-v,–verbose   详细模式，如果命令执行成功，输出详细的结果(-vv –vvv -vvvv)
-i PATH,–inventory=PATH   指定host文件的路径，默认是在/etc/ansible/hosts 
-f NUM,–forks=NU  NUM是指定一个整数，默认是5，指定fork开启同步进程的个数。 
-m NAME,–module-name=NAME   指定使用的module名称，默认是command
-m DIRECTORY,–module-path=DIRECTORY   指定module的目录来加载module，默认是/usr/share/ansible, 
-a,MODULE_ARGS   指定module模块的参数 
-k,-ask-pass     提示输入ssh的密码，而不是使用基于ssh的密钥认证
-sudo                   指定使用sudo获得root权限
-K,-ask-sudo-pass       提示输入sudo密码，与–sudo一起使用 
-u USERNAME,-user=USERNAME  指定移动端的执行用户 
-C,-check               测试此命令执行会改变什么内容，不会真正的去执行
```



## 5、Ansible主要组成部分

```
(1)Ansibleplaybooks
   任务剧本（任务集），编排定义ansible任务集的配置文件，有ansible顺序依次执行，通常是JSON格式的YML文件
(2)INVENTORY
   ansible管理主机的清单，可以定义单独主机，也可以定义主机组。
(3)Modeules
   ansible执行命令的功能模块，多数为内置的核心模块，也可自定义。Ansible2.2版本提供模块有500多个，下面应用到核心模块有synchronize（备份）、copy（恢复）、shell（nas获取）、cron（定时任务）、user（密码修改）、zypper（软件包管理rpm）、setup（获取facts）；扩展模块有UPStartItem（启动项管理）、UPInstall（软件包管理）、UPUPload（文件上传）
(4)Plugins
   模块功能的补充，如连接类型插件、循环插件、变量插件、过滤插件等，该功能不常用。
(5)API
   供第三方程序调用的应用程序编程接口
     ANSIBLE
     该部分图中表示不明显，组合INVENTORY、API、MODULES、PLUGINS的绿框大家可以理解为是ansible命令工具，其为核心执行工具；
```



## 6、Ansible组件调用关系

![img](https://static.oschina.net/uploads/space/2017/0410/151635_dHB9_3413282.png)

> 使用者使用Ansible或Ansible-playbook（会额外读取Playbook文件）时，在服务器终端输入Ansible的Ad-Hoc命令集或Playbook后，Ansible会遵循预先编排的规则将Playbooks逐条拆解为Play，再将Play组织成Ansible可识别的任务（Task），随后调用任务涉及的所有模块（Module）和插件（Plugin），根据Inventory中定义的主机列表通过SSH（Linux默认）将任务集以临时文件或命令的形式传输到远程客户端执行并返回执行结果，如果是临时文件则执行完毕后自动删除。



## 7、Ansible命令集

```
/usr/bin/ansible   # Ansibe AD-Hoc 临时命令执行工具，常用于临时命令的执行
/usr/bin/ansible-doc  # Ansible 模块功能查看工具
/usr/bin/ansible-galaxy # 下载/上传优秀代码或Roles模块的官网平台，基于网络的
/usr/bin/ansible-playbook # Ansible 定制自动化的任务集编排工具
/usr/bin/ansible-pull # Ansible远程执行命令的工具（使用较少，海量机器时使用，对运维的架构能力要求较高）
/usr/bin/ansible-vault # Ansible 文件加密工具
/usr/bin/ansible-console  # Ansible基于Linux Consoble界面可与用户交互的命令执行工具
```



## 8、实现基于ssh密钥通信

```
[root@Ansible ~]# ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa):
Created directory '/root/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /root/.ssh/id_rsa.
Your public key has been saved in /root/.ssh/id_rsa.pub.
The key fingerprint is:
ad:c8:99:6b:70:b2:c0:10:2e:4b:c7:68:52:2d:17:e0 root@John
The key's randomart image is:
+--[ RSA 2048]----+
|  .o..           |
|..o o            |
|.oEo             |
|+= o     .       |
|=+.     S .      |
|. o o..+ .       |
|   . == .        |
|    . ..         |
|     ..          |
+-----------------+
[root@Ansible ~]# chmod 700 /root/.ssh
[root@Ansible ~]# cat /root/.ssh/id_rsa.pub > /root/.ssh/authorized_keys
[root@Ansible ~]# chmod 600 /root/.ssh/authorized_keys
#同步到其他机器
[root@Ansible ~]# scp .ssh/authorized_keys root@192.168.102.133:/root/.ssh/
[root@Ansible ~]# ssh 192.168.102.133
Last login: Mon Apr 10 23:09:16 2017 from 192.168.102.1
[root@node1 ~]#
```



## 9、简单测试Ansible

```
[root@Ansible ~]# vim /etc/ansible/hosts 
[webserver]
192.168.102.133
[root@Ansible ~]# ansible webserver -m ping
192.168.102.133 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
```



# 三：主机清单介绍hosts

Ansible 通过读取默认的主机清单配置/etc/ansible/hosts，可以同时连接到多个远程主机上执行任务，默认路径可以通过修改 ansible.cfg 的 hostfile 参数指定路径。

```
[dbserver]  []表示主机的分组名,可以按照功能,系统进行分类,便于进行操作
192.168.10.2 
one.example.com 
www.bds.com:5309         #支持指定ssh端口5309 
jumper ansible_ssh_port=5309 ansible_ssh_host=192.168.10.2   #设置主机别名为jumper
www[01:50].bds.com       #支持通配符匹配www01.bds.com www02.bds.com
[web]                    #提醒下这里面字母是随便定义的
web-[a:f].bds.com        #支持字母匹配 web-a.bds.com ..web-f.bds.com
```



## 1、为主机指定类型和连接用户

```
[bds]
Localhost  ansible_connection=local
other1.example.com ansible_connection=ssh ansible_ssh_user=deploy
other2.example.com ansible_connection=ssh ansible_ssh_user=deploy
ansible hosts配置文件中支持指令
```

注意: 前面如果不配置主机免密钥登录,可以在/etc/ansible/hosts中定义用户和密码,主机ip地址,和ssh端口,这样也可以进行免密码访问,但是这个/hosts文件要保护好,因为所有的密码都写在里面



## 2、hosts文件配置参数介绍

```
1, ansible_ssh_host :
指定主机别名对应的真实 IP，如：100 ansible_ssh_host=192.168.1.100，随后连接该主机无须指定完整 IP，只需指定 251 就行
2, ansible_ssh_port :
指定连接到这个主机的 ssh 端口，默认 22
3, ansible_ssh_user:
连接到该主机的 ssh 用户
4, ansible_ssh_pass:
连接到该主机的 ssh 密码（连-k 选项都省了），安全考虑还是建议使用私钥或在命令行指定-k 选项输入
5, ansible_sudo_pass:                              sudo 密码
6, ansible_sudo_exe:                                sudo 命令路径
7, ansible_connection :
连接类型，可以是 local、ssh 或 paramiko，ansible1.2 之前默认为 paramiko
8, ansible_ssh_private_key_file :               私钥文件路径
9, ansible_shell_type  :
目标系统的 shell 类型，默认为 sh,如果设置 csh/fish，那么命令需要遵循它们语法
10, ansible_python_interpreter  :
python 解释器路径，默认是/usr/bin/python，但是如要要连BSD系统的话，就需要该指令修改 python 路径
11, ansible__interpreter  :
这里的"*"可以是 ruby 或 perl 或其他语言的解释器，作用和 ansible_python_interpreter 类似
```

> Ansible 可同时操作属于一个组的多台主机,组和主机之间的关系通过 inventory 文件配置. 默认的文件路径为 /etc/ansible/hosts配置文件：**ansible.cfg**，**hosts**

```
[root@Ansible ~]# cat /etc/ansible/ansible.cfg  主配置文件，可以根据实际应用自行修改
[defaults]
# some basic default values...
hostfile       = /etc/ansible/hosts   \\指定默认hosts配置的位置
# library_path = /usr/share/my_modules/  
remote_tmp     = $HOME/.ansible/tmp
pattern        = *
forks          = 5
poll_interval  = 15
sudo_user      = root  \\远程sudo用户
#ask_sudo_pass = True  \\每次执行ansible命令是否询问ssh密码
#ask_pass      = True  \\每次执行ansible命令时是否询问sudo密码
transport      = smart
remote_port    = 22
module_lang    = C
gathering = implicit
host_key_checking = False    \\关闭第一次使用ansible连接客户端是输入命令提示
log_path    = /var/log/ansible.log \\需要时可以自行添加。chown -R root:root ansible.log
system_warnings = False    \\关闭运行ansible时系统的提示信息，一般为提示升级
   
# set plugin path directories here, separate with colons
action_plugins     = /usr/share/ansible_plugins/action_plugins
callback_plugins   = /usr/share/ansible_plugins/callback_plugins
connection_plugins = /usr/share/ansible_plugins/connection_plugins
lookup_plugins     = /usr/share/ansible_plugins/lookup_plugins
vars_plugins       = /usr/share/ansible_plugins/vars_plugins
filter_plugins     = /usr/share/ansible_plugins/filter_plugins
fact_caching = memory
   
[accelerate]
accelerate_port = 5099
accelerate_timeout = 30
accelerate_connect_timeout = 5.0
# The daemon timeout is measured in minutes. This time is measured
# from the last activity to the accelerate daemon.
accelerate_daemon_timeout = 30
[root@Ansible ~]# cat /etc/ansible/hosts  主机清单信息配置文件，可以自定义主机，支持IP，域名，支持分组

[host01]  \\分组名,[]表示主机的分组名,可以按照功能、系统等进行分类，便于对某些主机或者某一组功能相同的主机进行操作
10.11.8.21 ansible_ssh_user=root ansible_ssh_pass=Passwd     \\远程ip,ssh登录用户，密码 
10.11.8.28 ansible_ssh_user=root ansible_ssh_pass=GxwLaXOs&1SK
10.10.30.50                \\若主机间进行的秘钥通信，则只需要添加主机 ip 就行了
[host02]
10.11.2.28
   
[web]
10.11.0.25
10.11.0.26
[web:var]            \\统一对web组设置变量
ansible_ssh_port=22
ansible_ssh_user=root
ansible_ssh_pass=123456
  
[db]
10.11.1.10
test ansible_ssh_port=5555 ansible_ssh_host=10.11.15    \\设置主机别名为 test
10.11.1.11:2156                   \\指定ssh端口
www[001:006].example.com          \\支持通配符 www001 www002 ..
new-[a:f]-node.example.com        \\支持字母匹配 new-a-node.example.com new-b-node.example.com ...
   
[server:children]   \\组可以包含其它组
web
db
[test]        
host01
host02
 
hosts 文件支持一些特定指令，所有支持的指令如下：
 
ansible_ssh_host：指定主机别名对应的真实 IP，如：251 ansible_ssh_host=183.60.41.251，随后连接该主机无须指定完整 IP，只需指定 251 就行
ansible_ssh_port：指定连接到这个主机的 ssh 端口，默认 22
ansible_ssh_user：连接到该主机的 ssh 用户
ansible_ssh_pass：连接到该主机的 ssh 密码（连-k 选项都省了），安全考虑还是建议使用私钥或在命令行指定-k 选项输入
ansible_sudo_pass：sudo 密码
ansible_sudo_exe(v1.8+的新特性):sudo 命令路径
ansible_connection：连接类型，可以是 local、ssh 或 paramiko，ansible1.2 之前默认为 paramiko
ansible_ssh_private_key_file：私钥文件路径
ansible_shell_type：目标系统的 shell 类型，默认为 sh,如果设置 csh/fish，那么命令需要遵循它们语法
ansible_python_interpreter：python 解释器路径，默认是/usr/bin/python，但是如要要连*BSD系统的话，就需要该指令修改 python 路径
ansible_*_interpreter：这里的"*"可以是 ruby 或 perl 或其他语言的解释器，作用和 ansible_python_interpreter 类似
例子：
  some_host ansible_ssh_port=2222 ansible_ssh_user=manager
  aws_host ansible_ssh_private_key_file=/home/example/.ssh/aws.pem
  freebsd_host ansible_python_interpreter=/usr/local/bin/python
  ruby_module_host ansible_ruby_interpreter=/usr/bin/ruby.1.9.3
```



## 3、Inventory 参数的说明



### 3.1、主机连接：

| **参数**           | **说明**                                                     |
| ------------------ | ------------------------------------------------------------ |
| ansible_connection | 与主机的连接类型.比如:local, ssh 或者 paramiko. Ansible 1.2 以前默认使用 paramiko.1.2 以后默认使用 'smart','smart' 方式会根据是否支持 ControlPersist, 来判断'ssh' 方式是否可行. |



### 3.2、ssh连接参数：

| **参数**                     | **说明**                                                     |
| ---------------------------- | ------------------------------------------------------------ |
| ansible_ssh_host             | 将要连接的远程主机名.与你想要设定的主机的别名不同的话,可通过此变量设置. |
| ansible_ssh_port             | ssh端口号.如果不是默认的端口号,通过此变量设置.               |
| ansible_ssh_user             | 默认的 ssh 用户名                                            |
| ansible_ssh_pass             | ssh 密码(这种方式并不安全,我们强烈建议使用 --ask-pass 或 SSH 密钥) |
| ansible_ssh_private_key_file | ssh 使用的私钥文件.适用于有多个密钥,而你不想使用 SSH 代理的情况. |
| ansible_ssh_common_args      | 此设置附加到sftp，scp和ssh的缺省命令行                       |
| ansible_sftp_extra_args      | 此设置附加到默认sftp命令行。                                 |
| ansible_scp_extra_args       | 此设置附加到默认scp命令行。                                  |
| ansible_ssh_extra_args       | 此设置附加到默认ssh命令行。                                  |
| ansible_ssh_pipelining       | 确定是否使用SSH管道。 这可以覆盖ansible.cfg中得设置。        |



### 3.3、远程主机环境参数：

| **参数**                   | **说明**                                                     |
| -------------------------- | ------------------------------------------------------------ |
| ansible_shell_type         | 目标系统的shell类型.默认情况下,命令的执行使用 'sh' 语法,可设置为 'csh' 或 'fish'. |
| ansible_python_interpreter | 目标主机的 python 路径.适用于的情况: 系统中有多个 Python, 或者命令路径不是"/usr/bin/python",比如  *BSD, 或者 /usr/bin/python |
| ansible_*_interpreter      | 这里的"*"可以是ruby 或perl 或其他语言的解释器，作用和ansible_python_interpreter 类似 |
| ansible_shell_executable   | 这将设置ansible控制器将在目标机器上使用的shell，覆盖ansible.cfg中的配置，默认为/bin/sh。 |



# 四：Ansible常用模块



## 1、ansible使用帮助

```
[root@Ansible ~]# ansible-doc -l                #查询ansible的所有模块
[root@Ansible ~]# ansible-doc -s module_name      #查看模块的属性
[root@Ansible ~]# ansible-doc -s shell
- name: Execute commands in nodes.
  action: shell
      chdir                  # cd into this directory before running the command
      creates                # a filename, when it already exists, this step will *not* be run.
      executable             # change the shell used to execute the command. Should be an absolute
                               path to the executable.
      free_form=             # The shell module takes a free form command to run, as a string.
                               There's not an actual option named
                               "free form".  See the examples!
      removes                # a filename, when it does not exist, this step will *not* be run.
      warn                   # if command warnings are on in ansible.cfg, do not warn about this
                               particular line if set to no/false.
```



## 2、command模块

默认模块 ,用于在各个被管理节点运行指定的命令(不支持管道和变量)

相关选项如下：

```
    creates：一个文件名，当该文件存在，则该命令不执行
    free_form：要执行的Linux指令
    chdir：在执行指令之前，先切换到该目录
    removes：一个文件名，当该文件不存在，则该选项不执行
    executable：切换shell来执行指令，该执行路径必须是一个绝对路径
[root@Ansible ~]# ansible all -m command -a "hostname"
192.168.102.133 | SUCCESS | rc=0 >>
node1
```



## 3、shell模块

command模块功能相同，但比command的模块功能强大(支持管道和变量)

```
[root@Ansible ~]# ansible all -m shell -a "cat /etc/passwd | grep root"
192.168.102.133 | SUCCESS | rc=0 >>
root:x:0:0:root:/root:/bin/bash
operator:x:11:0:operator:/root:/sbin/nologin
```



## 4、user模块

用户模块,用于在各管理节点管理用户所使用

```
[root@Ansible ~]# ansible all -m user -a "name=zabbix uid=1001 home=/home shell=/sbin/nologin"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "comment": "", 
    "createhome": true, 
    "group": 1001, 
    "home": "/home", 
    "name": "zabbix", 
    "shell": "/sbin/nologin", 
    "state": "present", 
    "stderr": "useradd：警告：此主目录已经存在。\n不从 skel 目录里向其中复制任何文件。\n", 
    "system": false, 
    "uid": 1001
}
[root@Ansible ~]# ansible all -m shell -a "cat /etc/passwd | grep zabbix"
192.168.102.133 | SUCCESS | rc=0 >>
zabbix:x:1001:1001::/home:/sbin/nologin
#删除用户
[root@Ansible ~]# ansible all -m user -a "name=zabbix uid=1001 state=absent"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "force": false, 
    "name": "zabbix", 
    "remove": false, 
    "state": "absent"
}
```



## **5、****setup****模块**

收集ansible的facts信息

```
# ansiblestorm_cluster -m setup
```

用来测试远程主机的运行状态

```
[root@Ansible ~]# ansible all -m ping
192.168.102.133 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
```



## **6、****file****模块**

设定文件属性

```
force：需要在两种情况下强制创建软链接，一种是源文件不存在，但之后会建立的情况下；另一种是目标软链接已存在，需要先取消之前的软链，然后创建新的软链，有两个选项：yes|no
group：定义文件/目录的属组
mode：定义文件/目录的权限
owner：定义文件/目录的属主
path：必选项，定义文件/目录的路径
recurse：递归设置文件的属性，只对目录有效
src：被链接的源文件路径，只应用于state=link的情况
dest：被链接到的路径，只应用于state=link的情况
state：
      directory：如果目录不存在，就创建目录
      file：即使文件不存在，也不会被创建
      link：创建软链接
      hard：创建硬链接
      touch：如果文件不存在，则会创建一个新的文件，如果文件或目录已存在，则更新其最后修改时间
      absent：删除目录、文件或者取消链接文件
## 远程文件符号链接创建
[root@Ansible ~]# ansible all -m file -a "src=/etc/resolv.conf dest=/tmp/resolv.conf state=link"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "dest": "/tmp/resolv.conf", 
    "gid": 0, 
    "group": "root", 
    "mode": "0777", 
    "owner": "root", 
    "size": 16, 
    "src": "/etc/resolv.conf", 
    "state": "link", 
    "uid": 0
}
## 远程文件信息查看
[root@Ansible ~]# ansible all -m command -a "ls -al /tmp/resolv.conf"
192.168.102.133 | SUCCESS | rc=0 >>
lrwxrwxrwx 1 root root 16 4月  10 23:56 /tmp/resolv.conf -> /etc/resolv.conf
## 远程文件符号链接删除
[root@Ansible ~]# ansible all -m file -a "path=/tmp/resolv.conf state=absent"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "path": "/tmp/resolv.conf", 
    "state": "absent"
}
```



## **7、****copy****模块**

复制文件

```
1.   backup：在覆盖之前，将源文件备份，备份文件包含时间信息。有两个选项：yes|no
2.   content：用于替代“src”，可以直接设定指定文件的值
3.   dest：必选项。要将源文件复制到的远程主机的绝对路径，如果源文件是一个目录，那么该路径也必须是个目录
4.   directory_mode：递归设定目录的权限，默认为系统默认权限
5.   force：如果目标主机包含该文件，但内容不同，如果设置为yes，则强制覆盖.
6.   如果为no，则只有当目标主机的目标位置不存在该文件时，才复制。默认为yes
7.   others：所有的file模块里的选项都可以在这里使用
8.   src：被复制到远程主机的本地文件，可以是绝对路径，也可以是相对路径。
9.   如果路径是一个目录，它将递归复制。在这种情况下，如果路径使用“/”来结尾，则只复制目录里的内容.
10. 如果没有使用“/”来结尾，则包含目录在内的整个内容全部复制，类似于rsync。
```

实例：

```
[root@Ansible ~]# ansible all -m copy -a "src=/etc/ansible/ansible.cfg dest=/tmp/ansible.cfg owner=root group=root mode=0644"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "checksum": "5d8f0f588dc5240369222712d8bcfc806b548e77", 
    "dest": "/tmp/ansible.cfg", 
    "gid": 0, 
    "group": "root", 
    "md5sum": "7c1e77df8058b957b9411e37937822f5", 
    "mode": "0644", 
    "owner": "root", 
    "size": 14388, 
    "src": "/root/.ansible/tmp/ansible-tmp-1491840235.16-274530631209696/source", 
    "state": "file", 
    "uid": 0
}
[root@Ansible ~]# ansible all -m copy -a "content='Hello Ansible' dest=/tmp/test.ansible.cfg"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "checksum": "489ae199bd11e563e29d6240b7726c4bca5b9cac", 
    "dest": "/tmp/test.ansible.cfg", 
    "gid": 0, 
    "group": "root", 
    "md5sum": "2b1cfc2ab8af697ac47f9cbbf7ac5a93", 
    "mode": "0644", 
    "owner": "root", 
    "size": 13, 
    "src": "/root/.ansible/tmp/ansible-tmp-1491840289.85-171493637637784/source", 
    "state": "file", 
    "uid": 0
}
[root@Ansible ~]# ansible all -a "cat /tmp/test.ansible.cfg"
192.168.102.133 | SUCCESS | rc=0 >>
Hello Ansible
```



## 8、script模块

自动复制脚本到远程节点,并运行

测试脚本

```
[root@Ansible ~]# cat ansible_test.sh 
#bin/bash
echo "Hello Ansible" >> /tmp/ansible.log
```

运行脚本

```
[root@Ansible ~]# ansible all -m script -a 'ansible_test.sh'
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "rc": 0, 
    "stderr": "Shared connection to 192.168.102.133 closed.\r\n", 
    "stdout": "", 
    "stdout_lines": []
}
[root@node1 ~]# cat /tmp/ansible.log 
Hello Ansible
```



## 9、cron模块

计划定时任务,用于在各管理节点管理计划任务

```
[root@Ansible ~]# ansible all -m cron -a "name=time minute='*/2' job='/usr/sbin/ntpdate'"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "envs": [], 
    "jobs": [
        "time"
    ]
}
[root@node1 ~]# crontab -l
#Ansible: time
*/2 * * * * /usr/sbin/ntpdate
```



## 10、template模块

根据官方的翻译是：template使用了Jinjia2格式作为文件模板，进行文档内变量的替换的模块。他的每次使用都会被ansible标记为changed状态。

在定义模板之后可以实现各节点对应的变量来取代，表达式自身会根据当前节点所赋值做运算，之后生成的值则赋予这个参数，用于生成不同配置的配置文件，所以模板主要实现配置不同场景文本文件。而且这样使用模板语言来定义，模板语言中可以根据定义替换成特定主机的某些值。

```
backup=     备份
src=\'#\'" 源文件
dest=          目标路径
owner=      属主
group=       主组
mode=       权限
```



## 11、yum模块

用于管理节点安装软件所使用

```
ansible-doc -s yum action: yum
      state=       指定是安装(`present', `latest'), 还是卸载 remove (`absent')
      disablerepo= 如果有多个yum源可能禁用一个
      name=        要安装的包名(如果只写包名，就安装最新的包，如果想指定安装什么片本的就包名+版本号）
      enablerepo=  要启用的repo
      list=        列表(跟playbook相关）
      disable_gpg_check=  是否开启gpg_check
      conf_file=    可以用其它服务器上的配置文件，而不使用本地的配置文件
```

例：所有的主机都安装上corosync.

```
[root@Ansible ~]# ansible all -m yum -a "state=present name=corosync"
[root@Ansible ~]# ansible all -m yum -a 'name=ntp state=present'
```

卸载的软件只需要将 name=ntp state=absent

安装特定版本 name=nginx-1.6.2 state=present

指定某个源仓库安装软件包name=htop enablerepo=epel state=present

更新软件到最新版 name=nginx state=latest



## 12、service 模块

作用：管理服务

```
参数
       name=       服务名称
       state=        状态
           started       启动
           stopped      停止
           restarted    重启
           enabled=    [yes|no] 是否随系统启动
           runlevel=             运行级别
```

示例：

```
[root@Ansible ~]# ansible all -m service -a"name=httpd state=started enabled=yes runlevel=5"
192.168.102.133 | SUCCESS => {
    "changed": true, 
    "enabled": true, 
    "name": "httpd", 
    "state": "started"
}
```

记得针对Centos7就不要使用这个模块了。



## 13、实站案例

```
执行shell
获取web组里得eth0接口信息
ansible web -a "ifconfig eth0"
执行ifconfig eth0 命令，ansible模块 默认是command，它不会通过shell进行处理，
所以像$ HOME和像“<”，“>”，“|”，“;” 和“＆”将不工作（如果您需要这些功能，请使用shell模块）。
以shell解释器执行脚本
ansible web -m shell -a "ifconfig eth0|grep addr"
以raw模块执行脚本
ansible web -m raw -a "ifconfig eth0|grep addr
将本地脚本传送到远程节点上运行
ansible web -m script -a ip.sh
传输文件
拷贝本地的/etc/hosts 文件到web组所有主机的/tmp/hosts（空目录除外）
ansible web -m copy -a "src=/etc/hosts dest=/tmp/hosts"
拷贝本地的ntp文件到目的地址，设置其用户及权限，如果目标地址存在相同的文件，则备份源文件。
ansible web -m copy -a "src=/mine/ntp.conf dest=/etc/ntp.conf owner=root group=root mode=644 backup=yes force=yes"
file 模块允许更改文件的用户及权限
ansible web -m file -a "dest=/tmp/a.txt mode=600 owner=user group=user"
使用file 模块创建目录，类似mkdir -p
ansible web -m file -a "dest=/tmp/test mode=755 owner=user group=user state=directory"
使用file 模块删除文件或者目录
ansible web -m file -a "dest=/tmp/test state=absent"
创建软连接，并设置所属用户和用户组
ansible web -m file -a  "src=/file/to/link/to dest=/path/to/symlink owner=user group=user state=link"
touch 一个文件并添加用户读写权限，用户组去除写执行权限，其他组减去读写执行权限
ansible web -m file -a  "path=/etc/foo.conf state=touch mode='u+rw,g-wx,o-rwx'"
管理软件包
apt、yum 模块分表用于管理Ubuntu 系列和RedHat 系列系统软件包
更新仓库缓存，并安装"foo"
ansible web -m apt -a "name=foo update_cache=yes"
删除 "foo"
ansible web -m apt -a "name=foo state=absent"
安装  "foo"
ansible web -m apt -a "name=foo state=present"
安装  1.0版本的 "foo"
ansible web -m apt -a "name=foo=1.00 state=present"
安装最新得"foo"
ansible web -m apt -a "name=foo state=latest"
注释：Ansible 支持很多操作系统的软件包管理，使用时-m 指定相应的软件包管理工具模块，如果没有这样的模块，可以自己定义类似的模块或者使用command 模块来安装软件包
安装 最新的 Apache
ansible web -m yum -a  "name=httpd state=latest"
删除apache
ansible web -m yum -a  "name=httpd state=absent"
从testing 仓库中安装最后一个版本得apache
ansible web -m yum -a  "name=httpd enablerepo=testing state=present"
更新所有的包
ansible web -m yum -a  "name=* state=latest"
安装远程的rpm包
ansible web -m yum -a  "name=http://nginx.org/packages/centos/6/noarch/RPMS/nginx-release-centos-6-0.el6.ngx.noarch.rpm state=present"
安装  'Development tools' 包组
ansible web -m yum -a  "name='@Development tools' state=present"
用户和用户组
添加用户 'user'并设置其 uid 和主要的组'admin'
ansible web -m user -a "name=user  comment='I am user ' uid=1040 group=admin"
添加用户 'user'并设置其登陆shell，并将其假如admins和developers组
ansible web -m user -a "name=user shell=/bin/bash groups=admins,developers append=yes"
删除用户 'user '
ansible web -m user -a "name=user state=absent remove=yes"
创建 user用户得   2048-bit SSH key，并存放在 ~user/.ssh/id_rsa
ansible web -m user -a "name=user generate_ssh_key=yes ssh_key_bits=2048 ssh_key_file=.ssh/id_rsa"
设置用户过期日期
ansible web -m user -a "name=user shell=/bin/zsh groups=nobdy expires=1422403387"
创建test组，并设置git为1000
ansible web -m group -a "name=test gid=1000 state=present"
删除test组
ansible web -m group -a "name=test state=absent"
源码部署
Ansible 模块能够通知变更，当代码更新时，可以告诉Ansible 做一些特定的任务，比如从git 部署代码然后重启apache 服务等
ansible web-m git -a "repo=https://github.com/Icinga/icinga2.git dest=/tmp/myapp   version=HEAD"
服务管理
确保web组所有主机的httpd 是启动的
ansible web-m service -a "name=httpd state=started"
重启web组所有主机的httpd 服务
ansible web-m service -a "name=httpd state=restarted"
确保web组所有主机的httpd 是关闭的
ansible web-m service -a "name=httpd state=stopped"
后台运行
长时间运行的操作可以放到后台执行，ansible 会检查任务的状态；在主机上执行的同一个任
务会分配同一个job ID
后台执行命令3600s，-B 表示后台执行的时间
ansible all -B 3600 -a "/usr/bin/long_running_operation --do-stuff"
检查任务的状态
ansible all -m async_status -a "jid=123456789"
后台执行命令最大时间是1800s 即30 分钟，-P 每60s 检查下状态默认15s
ansible all -B 1800 -P 60 -a "/usr/bin/long_running_operation --do-stuff"
定时任务
每天5点，2点得时候执行 ls -alh > /dev/null
ansible test -m cron -a "name='check dirs' minute='0' hour='5,2' job='ls -alh > /dev/null'"
搜集系统信息
搜集主机的所有系统信息
ansible all -m setup
搜集系统信息并以主机名为文件名分别保存在/tmp/facts 目录
ansible all -m setup --tree /tmp/facts
搜集和内存相关的信息
ansible all -m setup -a 'filter=ansible_*_mb'
搜集网卡信息
ansible all -m setup -a 'filter=ansible_eth[0-2]'
```



# 五：Ansible的Play book的使用



## 1、playbook简介

  playbook是ansible用于配置、部署和管理被控节点的剧本，通过playbook的详细描述，执行其中的一系列tasks，可以让远程主机达到预期的状态，playbook就像Ansible控制器给被控节点列出的一系列to-do-list，而被控节点必须要完成。

  也可以这样理解，pplaybook是由一个或多个"play"组成的列表。play的主要功能在于将事先归并为一组的主机装扮成事先通过ansible中的task定义好的角色。从根本上来讲所谓task无非是调用ansible的一个module。将多个play组织在一个playbook中即可以让它们联同起来按事先编排的机制同唱一台大戏。



## 2、playbook使用场景

  执行一些简单的任务，使用ad-hoc命令可以方便的解决问题，但是有时一个设施太过于复制，需要大量的操作时，执行ad-hoc命令就不太合适，这时最好使用playbook，就像执行shell命令与写shell脚本一样，也可以理解为批处理任务，不过playbook有自己的语法格式，具体可以看下面介绍。

  使用playbook你可以方便的重用这些代码，可以移值到不同的机器上面，像函数一样，最大化的利用代码。在使用ansible的过程中，所处理的大部分操作都是需要编写playbook的。



## 3、Playbook的核心元素

```
         Hosts：               主机，部署目标
         Tasks：               任务，ansible，执行目的
         Varlables：         变量
         Templates：       包含了模板语法的文本文件；
         Handlers：                   有特定条件触发的任务
         Roles ：              角色  (特别介绍)
```



## 4、playbook使用

```
Usage: ansible-playbook playbook.yml
```

相对于ansible，增加了下列选项：

参数  说明

```
--flush-cache 清除fact缓存
--force-handlers  如果任务失败，也要运行handlers
--list-tags   列出所有可用的标签
--list-tasks  列出将要执行的所有任务
--skip-tags=SKIP_TAGS    跳过运行标记此标签的任务
--start-at-task=START_AT_TASK   在此任务处开始运行
--step 一步一步：在运行之前确认每个任务
-t TAGS, --tags=TAGS 只运行标记此标签的任务
```

示例：

```
ansible-playbook -i hosts ssh-addkey.yml                # 指定主机清单文件
ansible-playbook -i hosts ssh-addkey.yml  --list-tags   # 列出tags
ansible-playbook -i hosts ssh-addkey.yml  -T install    # 执行install标签的任务
```



## 5、playbook的基础组件

```
hosts: 运行指定任务的目标主机   
remote_user：在远程主机以哪个用户身份执行；
sudo_user：非管理员需要拥有sudo权限；
tasks：任务列表
  模块，模块参数,格式有如下两种：
    (1) action: module arguments
    (2) module: arguments 
```

示例1：

```
- hosts: all
remote_user: root
tasks:
 - name: install a group
   group: name=mygrp system=true 
 - name: install a user
   user: name=user1 group=mygrp system=true
```

示例2

```
- hosts: websrvs
remote_user: root
tasks:
 - name: install httpd package
  yum: name=httpd
 - name: start httpd service 
  service: name=httpd state=started
```

主要由三个部分组成。

1. **hosts部分**：使用hosts指示使用哪个主机或主机组来运行下面的tasks，每个playbook都必须指定hosts，hosts也可以使用通配符格式。主机或主机组在inventory清单中指定，可以使用系统默认的/etc/ansible/hosts，也可以自己编辑，在运行的时候加上-i选项，指定清单的位置即可。在运行清单文件的时候，--list-hosts选项会显示那些主机将会参与执行task的过程中。
2. **remote_user**：指定远端主机中的哪个用户来登录远端系统，在远端系统执行task的用户，可以任意指定，也可以使用sudo，但是用户必须要有执行相应task的权限。
3. **tasks**：指定远端主机将要执行的一系列动作。tasks的核心为ansible的模块，前面已经提到模块的用法。tasks包含name和要执行的模块，name是可选的，只是为了便于用户阅读，不过还是建议加上去，模块是必须的，同时也要给予模块相应的参数。



## 6、Play book变量的使用

```
（1）facts: 可直接调用
（2）ansible-playbook 命令的命令行中的自定义变量
    -e EXTRA_VARS, --extra-vars=EXTRA_VARS  #命令行中定义变量传递至yaml文件。
（3）通过roles传递变量
（4）Host Inventory
（a）向不同的主机传递不同的变量；
    IP/HOSTANME varable=value var2=value2
    在hosts 组ip后添加变量
（b）向组中的主机传递相同的变量
    [group:var]         
    arable=value
注意：Inventory参数：
 用于定义ansible远程连接目标主机时使用的参数，而非传递给playbook的变量。
    ansible_ssh_host   
    ansible_ssh_user
    ansible_ssh_port
    ansible_ssh_pass
    ansible_sudo_pass
		….
查看远程主机的全部系统信息
ansible all -m setup  #收集到的远程主机的变量
(1)变量的定义示例:         
	变量定义位置 /etc/ansible/hosts
	普通变量
		[web]
		172.16.250.240  http_port=80
		172.16.252.18   http_port=8080  
	组变量
		[web:var1]
		http_port=80
		[web]
		172.16.250.240  
		172.16.252.18    
	在playbook中定义变量的方法
		Vars：
	    - var1：value1
	    - var2：value2
	命令行指定变量
		nsible-playbook -e  调用
示例1：hosts定义变量使用方法
[root@centos7_1 ~]#vim /etc/ansible/hosts
[web]
172.16.250.90 hname=node1
[root@centos7_1 ~]# cd /apps/yaml/
[root@centos7_1 yaml]# vim hosname.yml
---
- hosts: web
  remote_user: root
  tasks:
  - name: sethostname
    hostname:name={{ hname }}
[root@centos7_1 yaml]# ansible-playbook  hosname.yml
示例2：在playbook中定义变量的方法
[root@centos7_1yaml]# vim user1.yml
---
- hosts: web
 remote_user: root
 vars:  #定义变量
 - username: testuser1   #变量列表
 - groupname: testgroup1
 tasks:
 - name: crete group
   group: name={{ groupname }} state=present
 - name: crate user
   user: name={{ username }} state=present                                                                                                                                                           
[root@centos7_1 yaml]#ansible-playbook  user1.yml
示例3：命令行参数传递
利用命令行定义变量传递参数至剧本安装memcached。
[root@centos7_1 yaml]#vim forth.yml
---
- hosts: web
 remote_user: root
 tasks:
 - name: install $pkname
yum: name={{pkname }} state=present       
[root@centos7_1yaml]# ansible-playbook -e pkname=memcached forth.yml
```



## 7、Play book中notifyh和handlers的使用

notify这个action可用于每个play的最后被触发，这样可以避免多次有改变发生时每次都执行指定的操作，取而代之，仅在所有的变化发生完成之后一次性地执行指定操作。

在notify中列出的操作称为handler，即notify中调用handler中定义的操作。

handler也是也写task的列表，通过名字来引用，他们和一般的task没有什么区别。

handler是由通知者进行notify，如果没有被notify，handler是不会执行的。

不管有多少个通知者进行了notify，等到play中的所有task执行完成之后，handler也只会被执行一次。

handlers最佳的应用场景是用来重启服务，或者触发系统重启操作的。除此之外很少会用到的。

 

示例：触发

利用notify、handlers触发式重启服务。

```
[root@centos7_1yaml]# vim web-2.yml
---
- hosts: web
  remote_user: root
  tasks:
  - name: install httpdpackage
    yum: name=httpdstate=present
  - name: install configurefile
copy: src=/apps/work/files/httpd.confdest=/etc/httpd/conf/ 
#该文件与目标主机文件不完全一致变回触发。
    notify: restart httpd
  - name: start httpd service
    service: name=httpdstate=started
  handlers:
  - name: restart httpd
service: name=httpd state=restarted
[root@centos7_1 yaml]#ansible-playbook  web-2.yml
```



## 8、Play book中tags的使用

  tags即标签，tags可以和一个play（就是很多个task）或者一个task进行捆绑。然后再执行play book时只需指定相应的tags即可仅执行与tags绑定的task。

示例：执行指定tags

```
[root@centos7_1yaml]# vim web-3.yml
---
- hosts: web
 remote_user: root
 tasks:
 - name: install httpd package
   yum: name=httpd state=present
 - name: install configure file
   copy: src=/apps/work/files/httpd.conf dest=/etc/httpd/conf/
   tags: instconf              #tags
 - name: start httpd service
   service: name=httpd state=started
[root@centos7_1 yaml]# ansible-playbook -tinstconf  web-3.yml 

#指定tags instconf 执行。
ansible-playbookweb-3.yml --tags=" instconf "  
执行此命令同样仅执行instconf 标签内容。
```



## 9、tepmplates 模板的使用

template是文本文件，嵌套有脚本(使用模板编程语言编写)的配置文件。

讲到template模板就不得不先介绍template使用语言 jinja2。



### 9.1、jinja2语言

  Jinja2是基于python的模板引擎，功能比较类似于PHP的smarty，J2ee的Freemarker和velocity。它能完全支持  unicode，并具有集成的沙箱执行环境，应用广泛。

```
Jinja2 语言：
  字面量：
    字符串：使用单引号或双引号；
    数字：整数，浮点数
    列表：[item1,item2 …..]
    元组：(item1item2…,)
    字典：{key1：value，key2：value….}
      布尔型： true/filase
    算数运算：
      +,- , * , / , // , % **
    比较操作：
      ==， != , >=  ,<=
    逻辑运算：
      and，or， not，
    流表达式
      For、IF、when
```

示例：模板安装nginx

模板配置文件nginx.conf.j2

```
Worker_porcesses {{ ansible_precossor_vcpus }}  #注意空格哦。
此变量执行ansible all -m setup  (收集到的远程主机的变量) 即可查看到
Worker_porcesses {{ ansible_precossor_vcpus +1 }}
此表达式也可。此处只为表示可支持算数运算。
[root@centos7_1 yaml]# vim nginx.yml
---
- hosts: web
 remote_user: root
 tasks:
 - name: install nginx
  yum: name=nginx state=present
 - name: install conf file
  template: src=/apps/work/files/nginx.conf.j2 dest=/etc/nginx/nginx.conf
  notify: restart nginx
  tags: instconf
 - name: start nginx service
  service: name=nginx state=started
 handlers:
 - name: restart nginx
    service:name=nginx state=restarted
[root@centos7_1 yaml]# ansible-playbook  nginx.yml
```



### 9.2、when条件判断  

when 语句：在task中使用。Jinja2的语法格式

```
tasks：
- name: install conf file to Centos7
  template:src=files/nginxconf.c7.j2 dest=/etc/nginx/nginx.conf
when: ansible_distribution_major_version==”7”
- name: install conf file to Centos6
  template:src=files/nginxconf.c6.j2 dest=/etc/nginx/nginx.conf
  when:ansible_distribution_major_version ==”6”
```

以上语法表示若查询远程主机系统为centos6则执行，install conf file to Centos6。

若为cenos7则执行install conf file to Centos7。



### 9.3、迭代with_items

  循环迭代，需要重复执行的任务；对迭代项引用，固定变量名为item，而后在task中使用with_items给定迭代的元素列表；

​     列表方法：

​       字符串

​       字典

示例1：

字符串方式

```
- name： install some package
  yum：name={{ item }}  state=present
   with_items:
  - nginx
   - memecached
   - php-fpm
```

示例2：

字典方式

```
    - name: add  some groups
    group: name={{ item }} state=present
    with_items:
    - group1
    - group2
    - group3
    - name: add some user
       user: name={{ item.name }} group={{item.group}} state=present
      with_items:
      - {name: 'user1',group: 'group1'}
      - {name: 'user2',group: 'group2'}
      - {name: 'user3',group: 'group3'}
```



## 10、Playbook执行结果解析

使用ansible-playbook运行playbook文件，得到如下输出信息，输出内容为JSON格式。并且由不同颜色组成，便于识别。一般而言

> 绿色代表执行成功，系统保持原样
>
> 黄色代表系统代表系统状态发生改变
>
> 红色代表执行失败，显示错误输出。



## 11、Playbook案例剖析

```
实例：
---
- hosts: all
  sudo: yes
 
  tasks:
   - name: 安装Apache
     yum: name={{ item }} state=present
     with_items:
     - httpd
     - httpd-devel
   - name: 复制配置文件
     copy:
       src=\'#\'" /tmp/httpd.conf",
         dest: "/etc/httpd/conf/httpd.conf" }
     - {
       src=\'#\'" /tmp/httpd-vhosts.conf",
       dest: "/etc/httpd/conf/httpd-vhosts.conf"
       }
   - name: 检查Apache运行状态，并设置开机启动
     service: name=httpd state=started enabled=yes
```



## 12、执行playbook文件

运行playbook，使用ansible-playbook命令

(1) 检测语法

```
ansible-playbook  --syntax-check  /path/to/playbook.yaml
```

(2) 测试运行

```
ansible-playbook -C /path/to/playbook.yaml
--list-hosts   # 列出主机
--list-tasks  # 列出任务
--list-tags   # 列出标签
```

 (3) 运行

```
ansible-playbook  /path/to/playbook.yaml
-t TAGS, --tags=TAGS
--skip-tags=SKIP_TAGS
--start-at-task=START_AT
```

在执行playbook前，可以做些检查

检查palybook语法

```
ansible-playbook -i hosts httpd.yml --syntax-check
```

列出要执行的主机

```
ansible-playbook -i hosts httpd.yml --list-hosts
```

列出要执行的任务

```
ansible-playbook -i hosts httpd.yml --list-tasks
```



## 13、debug你的playbook

```
检查语法：ansible-playbook --syntax-check playbook.yml 
查看host列表：ansible-playbook --list-hosts playbook.yml
查看task列表：ansible-playbook --list-tasks playbook.yml
检查模式(不会运行): ansible-playbook --check playbook.yml
diff模式(查看文件变化)： ansible-playbook --check --diff playbook.yml
从指定的task开始运行：ansible-playbook --start-at-task="install packages" playbook.yml
逐个task运行，运行前需要你确认：ansible-playbook --step playbook.yml
指定tags：ansible-playbook --tags=foo,bar playbook.yml
跳过tags：ansible-playbook --skip-tags=baz,quux playbook.yml
```



# 六：ROLES



## ROLES 角色

  对于以上所有的方式有个弊端就是无法实现复用假设在同时部署Web、db、ha 时或不同服务器组合不同的应用就需要写多个yml文件。很难实现灵活的调用。。
   roles 用于层次性、结构化地组织playbook。roles能够根据层次型结构自动装载变量文件、tasks以及handlers等。要使用roles只需要在playbook中使用include指令即可。简单来讲，roles就是通过分别将变量(vars)、文件(file)、任务(tasks)、模块(modules)及处理器(handlers)放置于单独的目录中，并可以便捷地include它们的一种机制。角色一般用于基于主机构建服务的场景中，但也可以是用于构建守护进程等场景中。



## 1、目录层级结构

  roles每个角色中，以特定的层级目录进行组织

```
Mysql/  角色
 Files/     #存放有copy或script模块等调用的文件；’
 Tepmlates/    #template模块查找所需要模板文件目录；
 Tasks/           #定义任务；至少应该包含一个名为main.yml的文件；其他的文件需要在此文件中通过include进行包含。
 Handlers/      #定义触发器；至少应该包含一个名为main.yml的文件；其他的文件需要在此文件中通过include进行包含。
 Vars/              #定义变量；至少应该包含一个名为main.yml的文件；其他的文件需要在此文件中通过include进行包含。
 Meta/             #定义变量；至少应该包含一个名为main.yml的文件；定义当前角色的特殊设定及其依赖
关系；其他的文件需要在此文件中通过include进行包含。
 Default/         #设定默认变量时使用此目录中的main.yml文件。
```



## 2、角色调用

```
[root@centos7_1 yaml]# vim roles.yml
   ---
    Hosts：web
   Remote_user：root
   Roles：
   - mysql
   - memchached
   - nginx
```



##  3、层级结构展示

示例1：利用ansible角色安装nginx
 

```
[root@centos7_1 ~]# mkdir/etc/ansible/roles/nginx/{files,tasks,templates,handlers,vars, \
default,mata} –pv
#创建固定目录结构
[root@centos7_1 ~]# tree  /etc/ansible/roles/nginx/
/etc/ansible/roles/nginx/
├── default
├── files
├── handlers
├── mata
├── tasks
├── templates
└── vars
[root@centos7_1 ~]# cd/etc/ansible/roles/nginx/
[root@centos7_1 nginx]# vimtasks/main.yml  #创建任务
- name: install nginx package
 yum: name=nginx state=present
- name: install conf file
 template: src=nginx.conf.j2 dest=/etc/nginx/nginx.conf 
 #此处源文件可不写绝对路径，系统自查找。
- name: start nginx
 service: name=nginx state=started
[root@centos7_1 ~]# cp/apps/work/files/nginx.conf.c6.j2 ../templates/nginx.conf.j2 
#将配置文件拷贝至templates目录内。
[root@centos7_1 ~]# cd /apps/yaml/
[root@centos7_1 yaml]# cat roles.yml #创建调用文件
---
- hosts: web
 remote_user: root
 roles:
 - nginx
[root@centos7_1 yaml]#ansible-playbook roles.yml  #利用ansible-playbook执行。
```

示例2：变量调用
利用定义变量使远程主机的nginx服务运行用户变更为daemon

```
[root@centos7_1 ~]# vim/etc/ansible/roles/nginx/vars/main.yml
username: daemon
[root@centos7_1 ~]# vim/etc/ansible/roles/nginx/templates/nginx.conf.j2
user {{ username }};  #  将此处原有用户修改为变量
[root@centos7_1 ~]# cd/apps/yaml/
[root@centos7_1 yaml]#ansible-playbook  roles.yml
[root@centos7_1 yaml]#ansible-playbook  -e"username=adm"  roles.yml
#也可以直接利用命令行传递变量参数给剧本文件。
```

示例3：在playbook调用角色方法:传递变量给角色

```
[root@centos7_1 yaml]vim roles.yml
---
 - hosts：web
  remote_user:root
  roles:
  - {role: nigix, username: nginx } 
  #在调用nginx角色是使用变量username:nginx时服务运行用户为nginx
   键role:用于指定角色名称；后续的键值对用户传递变量给角色
[root@centos7_1yaml]# ansible-playbook roles.yml
```

示例4：条件测试角色调用
  还可以基于条件测试实现角色调用；

```
[root@centos7_1yaml]vim roles.yml
---
- hosts：web
  remote_user: root
  roles:
 {role: nigix, username: nginx ,when: “ansible_distribution_major_version ==’7’”}
#基于条件测试调用变量赋予nginx。
[root@centos7_1 yaml]#ansible-playbook -t instconf  roles.yml
```

示例5：角色安装

```
[root@centos7_1 ~]# mkdir/etc/ansible/roles/memcached/tasks -pv
[root@centos7_1 ~]# vim  /etc/ansible/roles/memcached/tasks/main.yml
- name: install package
 yum: name=memcached state=present
- name: start memecached
 service: name=memcached state=started
    
[root@centos7_1 ~]# cd/apps/yaml/
[root@centos7_1 yaml]# cat mem.yml
---
- hosts: web
  remote_user: root
  roles:
  - { role: nginx,when:ansible_distribution_version == '7' }  
  #系统为centos7时调用执行nginx
  - { role: memcached,when: ansible_hostname =='memcached' }  
  #系统用户名为memcached的主机调用执行角色memcached。
```

示例6：角色变量调整memcached内存大小
利用变量使远程主机上的Memcahed的缓存大小占用系统内存大小的三分之一。

```
[root@centos7_1 ~]# cd/etc/ansible/roles/memcached/
[root@centos7_1 memcached]#ls
handlers/  tasks/    templates/
[root@centos7_1 memcached]#mkdir  templates
[root@centos7_1memcached]# scp 172.16.254.216:/etc/sysconfig/memcached \
    ./templates//memcached.j2
[root@centos7_1 memcached]#vim templates/memcached.j2
PORT="11211"
USER="memcached"
MAXCONN="1024"
CACHESIZE="{{ansible_memtotal_mb//3 }}"
 #变量设置内存的3分之一  此变量为远程主机的总内存//3 指除3取商
  便为远程主机的三分之一
[root@centos7_1 memcached]#mkdir handlers/
[root@centos7_1 memcached]#vim handlers/main.yml
- name: restart memcached
  service: name=memcached state=restarted
[root@centos7_1 memcached]#cd /apps/yaml/
root@centos7_1 yaml]#ansible-playbook   mem.yml  #执行剧本
```



## 4、时间计时模块

ansible中可以加入一个计时模块在执行ansible-playbook时显示执行时长。方便使用。

```
1、配置方法
  cd /etc/ansible
  mkdir callback_plugins
  cd callback_plugins
  wget https://raw.githubusercontent.com/jlafon/ansible- \ profile/master/callback_plugins/profile_tasks.py
注意：ansible2.0以上版本需在ansible.cdg中加入以下内容
  [defaults] 下面加入
  callback_whitelist= profile_tasks
  再次执行ansbile-playbook时显示执行时长
2、测试结果
```

 