- [博客园](https://www.cnblogs.com/)
- [首页](https://www.cnblogs.com/opsprobe/)
- [新随笔](https://i.cnblogs.com/EditPosts.aspx?opt=1)
- [联系](https://msg.cnblogs.com/send/OpsDrip)
- [管理](https://i.cnblogs.com/)
- [订阅](javascript:void(0))[![订阅](https://www.cnblogs.com/skins/coffee/images/xml.gif)](https://www.cnblogs.com/opsprobe/rss/)

随笔- 106 文章- 0 评论- 21 

# [Ubuntu18.04下Ansible的基本使用](https://www.cnblogs.com/opsprobe/p/13163591.html)

这里使用四台Ubuntu18.04主机（一台充当Ansible管理端服务器，另外三台充当Ansible被管理端服务器），Ansible用apt安装，版本为2.9.10，下表是它们所使用的操作系统以及IP地址。

| 四台Ubuntu18.04主机所使用的操作系统以及IP地址 |             |               |
| --------------------------------------------- | ----------- | ------------- |
| 主机名称                                      | 操作系统    | IP地址        |
| Ansible管理端服务器                           | Ubuntu18.04 | 192.168.0.132 |
| Ansible被管理端服务器                         | Ubuntu18.04 | 192.168.0.133 |
| Ansible被管理端服务器                         | Ubuntu18.04 | 192.168.0.134 |
| Ansible被管理端服务器                         | Ubuntu18.04 | 192.168.0.135 |

Ansible官网文档：[https://docs.ansible.com](https://docs.ansible.com/)

Ansible服务特点说明：

（1）管理端不需要启动服务程序（no server）

（2）管理端不需要编写配置文件（/etc/ansible/ansible.cfg）

（3）受控端不需要安装软件程序（libselinux-python）

  　 被管理端selinux服务没有关闭 --- 影响ansible软件的管理

  　  libselinux-python让selinux开启的状态也可以使用ansible程序

（4）受控端不需要启动服务程序（no agent）

（5）服务程序管理操作模块众多（module）

（6）利用剧本编写来实现自动化（playbook）

**提示：以下操作均在root用户下进行，如在普通用户，请自行加上sudo！**

查看系统版本：

lsb_release -a

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619155246813-843182266.png)

## 一、安装Ansible

**注意：以下操作均在Ansible管理端服务器上！**

官网文档截图：

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619154553861-1375404479.png)

依次执行以下命令，安装ansible：

```
apt update

apt install software-properties-common

apt-add-repository --yes --update ppa:ansible/ansible

apt install ansible
```

安装完成后，查看ansible的版本：

ansible --version

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619154958118-1835167236.png)

## 二、创建“密钥对”，Ansible是基于SSH远程管理服务实现远程主机批量管理的

注意：这里使用的是root用户，Ubuntu18.04默认root用户没有密码，并且不能使用SSH远程，（需要设置root密码，使用命令：sudo passwd root，需要开启root用户SSH远程权限，在配置文件/etc/ssh/sshd_config里加入PermitRootLogin yes配置，重启SSH服务生效即可）

```
ssh-keygen（也可以使用-t选项指定秘钥的类型，如：ssh-keygen -t rsa或ssh-keygen -t dsa）
```

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619160701350-1744604034.png)

## 三、使用脚本将公钥批量分发至Ansible被管理端服务器

思考，批量分发公钥时有两个问题：

1、需要输入连接确认信息 yes/no

2、需要第一次连接输入密码

解决方法：

1、使用-o StrictHostKeyChecking=no选项，即可解决第一个问题

2、使用sshpass即可解决第二个问题，sshpass用于非交互的ssh密码验证，使用-p参数指定明文密码，然后直接登录远程服务器。

　 安装sshpass命令，apt install sshpass

实现脚本：

```
for i in `seq 3 5`; do sshpass -p123456 ssh-copy-id -i /root/.ssh/id_rsa.pub root@192.168.0.13$i "-o StrictHostKeyChecking=no"; done
```

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619165154967-1128279380.png)

查看软件安装到系统里面的文件目录信息

```
dpkg -L ansible|egrep -v "^/usr/(share|lib)"
```

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619162417883-1715283484.png)

说明：

/etc/ansible/ansible.cfg  # ansible服务配置文件

/etc/ansible/hosts      # 主机清单文件，定义可以管理的主机信息

/etc/ansible/roles      # 角色配置文件

## 四、配置主机清单文件，并测试是否可以管理多个主机

打开配置文件：vim /etc/ansible/hosts

在配置文件最后，输入所有Ansible被管理端服务器的IP地址

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619173312780-1648757762.png)

保存退出后，进行测试基于秘钥的SSH是否连接正常

执行下面的命令，查看Ansible被管理端服务器的主机名：

ansible all -a "hostname"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619203407760-769659834.png)

可以看到被管理端服务器主机名分别为：ceph3、ceph4、ceph2

## 五、Ansible模块的应用

模块的应用语法格式：

ansible 主机名称/主机组名称/主机地址信息/all -m(指定应用的模块信息) 模块名称 -a(指定动作信息) "执行什么动作"

## 命令类型模块

1、command模块（官方文档：https://docs.ansible.com/ansible/latest/modules/command_module.html#command-module）

简单用法：

查看所有在主机清单配置文件里主机的主机名

ansible all -m command -a "hostname"（command也是个Ansible默认模块，-m command可以不写）

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619203524348-1415429599.png)

扩展应用

官方截图：

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619205002677-117153702.png)

（1）参数：chdir，在执行命令之前对目录进行切换

以下命令意思是：在所有被管理端服务器的tmp目录下创建demo.txt文件

ansible all -m command -a "chdir=/tmp touch demo.txt"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619204840353-75723868.png)

可以看到有条警告信息，如果不想再看到，根据说明进入Ansible配置文件：vim /etc/ansible/ansible.cfg，去掉command_warnings = False（188行）配置前的#即可，注意要顶格。

（2）参数：creates，如果文件存在了，不执行命令操作

以下命令意思是：如果被管理端服务器/tmp目录下没有hosts文件（注意：如果不加chdir=/tmp，默认会将文件创建到家目录下），则创建demo123.txt文件

ansible all -m command -a "creates=/tmp/hosts chdir=/tmp touch demo123.txt"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619214233787-210983847.png)

（3）参数：removes，如果文件存在了, 这个步骤将执行

以下命令意思是：如果被管理端服务器/tmp目录下有hosts文件（注意：如果不加chdir=/tmp，默认会将文件创建到家目录下），则创建demo456.txt文件

ansible all -m command -a "removes=/tmp/hosts chdir=/tmp touch demo456.txt"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619214428704-1405431066.png)

注意事项：command模块有些符号信息无法识别，如："<"，">"，"|"，";"，"&"

2、shell模块（官方文档：https://docs.ansible.com/ansible/latest/modules/shell_module.html#shell-module）

shell模块参数功能和command模块类似

简单用法：

查看所有在主机清单配置文件里主机的主机名

ansible all -m shell -a "hostname"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200619215409880-1352778578.png)

shell模块被称为万能模块，command模块能干的不能干的它都能干

ansible all -m shell -a "chdir=/tmp touch demo.txt"

ansible all -m shell -a "creates=/tmp/hosts chdir=/tmp touch demo123.txt"

ansible all -m shell -a "removes=/tmp/hosts chdir=/tmp touch demo456.txt"

shell模块，这些"<"，">"，"|"，";"，"&"符合都可以识别：

ansible all -m shell -a "echo 123 >> /tmp/hosts.bak"

ansible all -m shell -a "grep 192 /etc/hosts"

3、script模块（官方文档：https://docs.ansible.com/ansible/latest/modules/script_module.html#script-module）

scripts模块参数功能也和command模块类似

新建demo.sh文件，写入以下脚本，脚本内容是计算1到100的累加和并将结果重定向到tmp目录的sum.txt文件：

vim /root/demo.sh

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

```
#!/bin/bash
sum=0
for i in `seq 1 100`
do
        sum=$[$i+$sum]
done
echo $sum > /tmp/sum.txt
```

[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);)

保存退出后，执行以下命令，命令意思是：在所有被管理端服务器执行demo.sh脚本

ansible all -m script -a "/root/demo.sh"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200620110057357-1783407362.png)

## 文件类型模块

1、copy模块，批量分发文件（官方文档：https://docs.ansible.com/ansible/latest/modules/copy_module.html#copy-module）

以下命令意思是：将管理端服务器/root目录下的demo.sh文件批量分发至所有被管理端服务器/tmp目录下

ansible all -m copy -a "src=/root/demo.sh dest=/tmp/"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200620112502147-685127824.png)

扩展用法：

（1）在传输文件时修改文件的属主和属组信息，注意：指定的属主和属组必须在被管理端服务器存在，否则会报错

ansible all -m copy -a "src=/root/demo.sh dest=/tmp/ owner=zhangsan group=zhangsan"

（2）在传输文件时修改文件的权限信息

ansible all -m copy -a "src=/root/demo.sh dest=/tmp/ mode=777"

（3）在传输数据文件信息时对远程主机源文件进行备份，注意：文件必须有改动后再次传输才会备份，被管理端服务器的文件备份名类似这样，demo.sh.8328.2020-06-20@03:51:21~ 

ansible all -m copy -a "src=/root/demo.sh dest=/tmp/ backup=yes"

（4）创建一个文件并直接编辑文件的信息

ansible all -m copy -a "content='123456' dest=/tmp/rsync.password"

复制文件时加不加“/”的区别：

ansible all -m copy -a "src=/tmp dest=/tmp/"

src后面目录没有/：将目录本身以及目录下面的内容都进行远程传输复制

ansible all -m copy -a "src=/tmp/ dest=/tmp/"

src后面目录有/：只将目录下面的内容都进行远程传输复制

2、file模块，设置文件属性信息（官方文档：https://docs.ansible.com/ansible/latest/modules/file_module.html#file-module）

基本用法：

ansible all -m file -a "dest=/tmp/demo.sh owner=zhangsan group=zhangsan mode=666"

扩展用法:

可以利用模块创建数据信息 (文件 目录 链接文件)

state 参数

=absent   # 缺席/删除数据信息

=directory  # 创建一个目录信息

=file      # 检查创建的数据信息是否存在，绿色存在，红色不存在

=hard    # 创建一个硬链接文件

=link     # 创建一个软链接文件

=touch    # 创建一个文件信息

创建目录信息：

ansible all -m file -a "dest=/tmp/demo/ state=directory"

ansible all -m file -a "dest=/tmp/demo/demo01/demo02 state=directory"

创建文件信息：

ansible all -m file -a "dest=/tmp/demo.py state=touch"

创建链接文件信息：

ansible all -m file -a "src=/tmp/demo.txt dest=/tmp/demo_hard.txt state=hard"

ansible all -m file -a "src=/tmp/demo.txt dest=/tmp/demo_link.txt state=link"

在被管理端服务器上验证是否为硬链接，也就是查看inode索引节点编号是否相同，相同即为硬链接：ll -i /tmp/demo_hard.txt /tmp/demo.txt

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200620153512517-265962481.png)

可以利用模块删除数据信息

ansible all -m file -a "dest=/tmp/demo.txt state=absent"

ansible all -m file -a "dest=/tmp/demo/ state=absent"

3、 fetch模块，批量拉取数据（官方文档：https://docs.ansible.com/ansible/latest/modules/fetch_module.html#fetch-module）

以下命令意思是：将所有被管理端服务器/tmp目录下的demo.txt文件批量拉取到管理端服务器的/tmp目录下

ansible all -m fetch -a "src=/tmp/demo.txt dest=/tmp"

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200620162744210-1748530884.png)

拉取成功后的目录，默认以主机IP命名

## 包类型模块

1、apt模块（官方文档：https://docs.ansible.com/ansible/latest/modules/apt_module.html#apt-module）

批量安装apache2 ：

ansible all -m apt -a "name=apache2 state=present"

ansible all -m apt -a "name=apache2 state=latest"

批量卸载apache2 ：

ansible all -m apt -a "name=apache2 state=absent"

说明：

name # 指定安装软件名称

state  # 指定是否安装软件

　　absent　　# 卸载

　　build-dep

　　latest　　 # 安装最新版本

　　present（默认）

　　fixed

## 系统类型模块

1、service模块，被管理端服务器上服务的运行状态，停止，开启，重启（官方文档：https://docs.ansible.com/ansible/latest/modules/service_module.html#service-module）

批量停止apache2服务，并禁止开机自启动

ansible all -m service -a "name=apache2 state=stopped enabled=no"

说明：

name　　# 指定管理的服务名称

state　　 # 指定服务状态

　　reloaded　　# 重新加载

　　restarted　　# 重新启动

　　started　　  # 启动

　　stopped　   # 停止

enabled　　# 指定服务是否开机自启动

2、cron模块，批量设置多个主机的定时任务信息（官方文档：https://docs.ansible.com/ansible/latest/modules/cron_module.html#cron-module）

基本用法：

ansible all -m cron -a "minute=0 hour=2 job='sync && echo 3 > /proc/sys/vm/drop_caches >/dev/null 2>&1'"

说明：

 \*  *  *  *  * 定时任务动作

分 时 日 月 周

ansible-doc -s cron　　# 查看使用说明

minute　　　 　# 设置分钟信息

hour　　　　 　# 设置小时信息

day　　　　  # 设置日期信息

month　　　 　 # 设置月份信息

weekday　　 　# 设置周信息

job　　　　　 　# 用于定义定时任务需要干的事情

扩展用法：

（1）给定时任务设置注释信息

ansible all -m cron -a "name='time sync' minute=0 hour=2 job='sync && echo 3 > /proc/sys/vm/drop_caches >/dev/null 2>&1'"

（3）批量注释定时任务，disabled=yes为注释，disabled=no为取消注释

ansible all -m cron -a "name='time sync' minute=0 hour=2 job='sync && echo 3 > /proc/sys/vm/drop_caches >/dev/null 2>&1' disabled=yes"

（2）删除指定定时任务

ansible all -m cron -a "name='time sync' state=absent"

注意：ansible可以删除的定时任务，只能是ansible设置好的定时任务

3、mount模块，批量进行挂载操作（官方文档：https://docs.ansible.com/ansible/latest/modules/mount_module.html#mount-module）

基本用法：

ansible all -m mount -a "src=192.168.0.130:/data path=/mnt fstype=nfs state=mounted"

参数：

src：需要挂载的存储设备或文件信息

path：指定目标挂载点目录

fstype：指定挂载时的文件系统类型

state

　　present/mounted   # 进行挂载

　　present：不会实现立即挂载，修改fstab文件，实现开机自动挂载

　　mounted：会实现立即挂载，并且会修改fstab文件，实现开机自动挂载 *****

　　absent/unmounted  # 进行卸载

　　absent：会实现立即卸载，并且会删除fstab文件信息，禁止开机自动挂载

　　unmounted：会实现立即卸载，但是不会删除fstab文件信息 *****

4、user模块，实现批量创建用户（官方文档：https://docs.ansible.com/ansible/latest/modules/user_module.html#user-module）

基本用法：

ansible all -m user -a "name=wangwu"

扩展用法：

（1）指定用户uid信息

ansible all -m user -a "name=wangwu01 uid=6666"

（2）指定用户组信息，group为指定主要组，groups为指定次要组

ansible all -m user -a "name=wangwu02 group=wangwu01"

ansible all -m user -a "name=wangwu03 groups=wangwu01"

（3）批量创建虚拟用户，create_home=no为不创建家目录

ansible all -m user -a "name=wangwu04 create_home=no shell=/sbin/nologin"

（4）给指定用户创建密码

ansible all -i localhost, -m debug -a "msg={{ 'mypassword' | password_hash('sha512', 'mysecretsalt') }}"

ansible all -i localhost, -m debug -a "msg={{ '密码信息' | password_hash('sha512', '加密校验信息') }}"

ansible all -i localhost, -m debug -a "msg={{ '123456' | password_hash('sha512', 'abcdefg') }}"

ansible all -m user -a 'name=wangwu05 password=$6$abcdefg$SVqgF31LWsg8o0zQyy63h/NIcgk3RP516ZKxuWJasoxHYEkq0YWf.WgIhFz5uK19zJrEK61oiA.TFD5ROrPxH0'

![img](https://img2020.cnblogs.com/blog/1404518/202006/1404518-20200620213710487-1961265531.png)

## 六、Ansible剧本的编写方法

学习Ansible剧本编写前，需要先了解下YAML，YAML是专门用来写配置文件的语言，非常简洁和强大，远比JSON格式方便。

YAML入门教程：https://www.runoob.com/w3cnote/yaml-intro.html

如何执行剧本：

第一个步骤：检查剧本的语法格式

ansible-playbook --syntax-check rsync_server.yaml

第二个步骤：模拟执行剧本

ansible-playbook -C rsync_server.yaml

第三个步骤：直接执行剧本

ansible-playbook rsync_server.yaml

 

## 扩展

ansible软件输出颜色说明：

绿色信息：查看主机信息/对主机未做改动

黄色信息：对主机数据信息做了修改

红色信息：命令执行出错了

粉色信息：忠告信息

蓝色信息：显示ansible命令执行的过程

ansible spark -m setup