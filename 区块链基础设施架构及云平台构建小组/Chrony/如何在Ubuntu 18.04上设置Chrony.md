# 如何在Ubuntu 18.04上设置Chrony

[2020年2月7日](https://sysadminjournal.com/chrony-ubuntu/) 4分钟阅读



分享

## 什么是Chrony？

[Chrony](https://chrony.tuxfamily.org/)是网络时间协议（NTP）的另一种实现。它用于同步来自不同NTP服务器的时间和时钟，参考时钟和手动输入。

可以将Chrony用作NTPv4服务器（RFC 5905），该服务器向网络中的其他计算机提供时间服务。它提供了两个不同的程序，chronyd是服务守护程序，chronyc是用于配置Chrony的命令行界面。

**我们将要做什么？**

在本教程中，您将学习如何在Ubuntu 18.04 Server上设置Chrony。您将学习如何将Chrony安装和配置为NTP服务器。另外，使用Chrony作为NTP客户端。

**先决条件**

对于本教程，请确保您具有具有root用户特权的Ubuntu 18.04服务器。

以下是有关Ubuntu 18.04服务器安装的教程。

[安装Ubuntu 18.04服务器](https://sysadminjournal.com/install-ubuntu/)

在继续之前，请登录Ubuntu服务器并键入sudo命令以获取系统的root特权。

苏多苏

### 步骤1 –在Ubuntu 18.04服务器上安装Chrony

首先，您将安装Chrony及其在Ubuntu系统上的依赖项。默认情况下，chrony软件包在官方的Ubuntu存储库中可用。

使用下面的apt命令更新Ubuntu存储库并安装chrony。

apt更新

apt安装年代

之后，启动chrony服务并将其添加到系统引导中。

systemctl开始时间

systemctl使能时间

计时服务已启动并正在运行，请使用以下命令对其进行检查。

chronyc活动

结果，您已经在Ubuntu 18.04 Server上安装了Chrony。

![在Ubuntu 18.04上安装Chrony](https://sysadminjournal.com/wp-content/uploads/2020/02/1-2-700x240.png)

### 第2步–设置Chrony NTP服务器

安装Chrony之后，您将更改默认的chrony NTP池，并将chrony设置为NTP服务器。

默认情况下，chrony配置'chrony.conf'位于'/ etc / chrony'目录中。使用vim编辑器编辑chrony配置'/etc/chrony/chrony.conf'。

vim / etc / chrony / chrony。conf

#### –更改默认NTP池

NTP池是志愿者提供的计算机服务器的集合，目的是通过网络时间协议（NTP）为全世界的客户端提供高度准确的时间同步。

建议使用距离您区域最近的NTP池，您可以在[此处](https://www.ntppool.org/)根据您的区域找到NTP池。

现在，使用您最近的区域更改默认的NTP池，并进行以下配置。

池0 JP 。池。ntp 。组织iburst maxsources 2

池1. jp 。池。ntp 。组织iburst maxsources 2

池2 JP 。池。ntp 。组织iburst maxsources 2

池3. jp 。池。ntp 。组织iburst maxsources 2

保存并关闭。

结果，您已使用最近的区域更改了默认NTP池。

#### –使用Chrony设置NTP服务器

接下来，您将使该服务器成为同一本地网络上所有客户端的NTP服务器。

添加以下配置并使用您自己的名称更改服务器IP地址。

允许10.5 。5.0 / 24

允许10.5 。5.3

保存并关闭。

现在，使用以下命令重新启动chrony服务。

systemctl重新启动时间

因此，您已经使用Chrony为本地网络创建了NTP服务器。

![更改默认池并设置NTP服务器](https://sysadminjournal.com/wp-content/uploads/2020/02/2-2-700x224.png)

### 步骤3 –使用Chrony设置NTP客户端

在此步骤中，您将使用Chrony设置NTP客户端，ane确保您已在客户端计算机上安装了chrony软件包。

要设置NTP客户端，请使用vim编辑器编辑chrony配置'/etc/chrony/chrony.conf'。

vim / etc / chrony / chrony。conf

现在，使用配置和以下内容更改默认的NTP池，并使用您自己的更改NTP服务器IP地址。

服务器10.5 。5.21 喜欢iburst

保存并关闭。

之后，重新启动Chrony服务。

systemctl重新启动时间

结果，您已经在Ubuntu系统上使用Chrony设置了NTP客户端。

![使用Chrony在Ubuntu系统上设置NTP客户端](https://sysadminjournal.com/wp-content/uploads/2020/02/3-2-700x309.png)

### 步骤4 –测试

在此步骤中，您将学习“ chronyc”命令来检查NTP服务器和客户端。

#### –检查NTP的时间源

要检查chronyd服务使用的NTP服务器的当前来源，请使用以下'sources'选项。

年代来源

以下是您将获得的结果。

![检查NTP服务器的来源](https://sysadminjournal.com/wp-content/uploads/2020/02/4-2-700x243.png)

有关详细信息状态，请添加“ -v”选项（详细说明）。

chronyc来源-v

以下是NTP来源的详细信息。

![检查NTP服务器详细信息的来源](https://sysadminjournal.com/wp-content/uploads/2020/02/5-2-700x392.png)

此外，以下是NTP客户端计算机上相同命令的结果。

![检查NTP服务器详细信息的来源](https://sysadminjournal.com/wp-content/uploads/2020/02/6-2-700x422.png)

#### –检查连接到NTP服务器的客户端

在为本地网络客户端运行NTP服务器时，请使用下面的“ chronyc”命令检查连接到NTP服务器的客户端的IP地址。

chronyc客户

结果，您的客户端已连接到刚刚在顶部设置的本地NTP服务器。

![检查长期客户](https://sysadminjournal.com/wp-content/uploads/2020/02/7-2-700x147.png)

最后，您已经在Ubuntu 18.04 Server上使用Chrony成功设置了NTP服务器和客户端。

# 使用Ansible配置时间服务

 **发表于\*2018年9月12日* *读了4分钟*

<details class="container entry-toc" open="" style="box-sizing: inherit; display: block; width: 740px; max-width: 740px; padding: 0px 20px; margin: 0px auto 2.5em 0px;"><summary class="title" style="box-sizing: inherit; display: block; cursor: pointer; font-family: &quot;Libre Baskerville&quot;, serif; font-weight: 700; line-height: 1.2; color: rgb(38, 50, 56); margin: 0px 0px 0.625em; padding-left: 0px; font-size: 16.2px;"><span>&nbsp;</span><span style="box-sizing: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;">目录</font></font></span></summary><nav id="TableOfContents" style="box-sizing: inherit;"><ul style="box-sizing: inherit; padding: 0px; margin: 0px; list-style: none;"><li style="box-sizing: inherit;"><a href="https://www.nathancurry.com/blog/20-configuring-chrony-with-ansible/#structure" style="box-sizing: inherit; background-color: transparent; color: rgb(33, 33, 33); text-decoration: none; transition: color 0.1s ease-in-out 0s;"><font style="box-sizing: inherit; vertical-align: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;">结构体</font></font></a><ul style="box-sizing: inherit; padding: 0px 0px 0px 1.5em; margin: 0px; list-style: none;"><li style="box-sizing: inherit;"><a href="https://www.nathancurry.com/blog/20-configuring-chrony-with-ansible/#configure-time-yml" style="box-sizing: inherit; background-color: transparent; color: rgb(33, 33, 33); text-decoration: none; transition: color 0.1s ease-in-out 0s;"><font style="box-sizing: inherit; vertical-align: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;">configure_time.yml</font></font></a></li><li style="box-sizing: inherit;"><a href="https://www.nathancurry.com/blog/20-configuring-chrony-with-ansible/#tasks-main-yml" style="box-sizing: inherit; background-color: transparent; color: rgb(33, 33, 33); text-decoration: none; transition: color 0.1s ease-in-out 0s;"><font style="box-sizing: inherit; vertical-align: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;">任务/ main.yml</font></font></a></li><li style="box-sizing: inherit;"><a href="https://www.nathancurry.com/blog/20-configuring-chrony-with-ansible/#templates" style="box-sizing: inherit; background-color: transparent; color: rgb(33, 33, 33); text-decoration: none; transition: color 0.1s ease-in-out 0s;"><font style="box-sizing: inherit; vertical-align: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;">范本</font></font></a></li></ul></li><li style="box-sizing: inherit;"><a href="https://www.nathancurry.com/blog/20-configuring-chrony-with-ansible/#testing" style="box-sizing: inherit; background-color: transparent; color: rgb(33, 33, 33); text-decoration: none; transition: color 0.1s ease-in-out 0s;"><font style="box-sizing: inherit; vertical-align: inherit;"><font style="box-sizing: inherit; vertical-align: inherit;">测验</font></font></a></li></ul></nav></details>

我使用chrony组合了Ansible角色来配置网络时间服务。

到目前为止，这是我迄今为止最出色的工作，尽管在进行打包之前，我想配置脚本以测试Debian vs. Red Hat，以及添加一些变量。

您可以在[此处](https://github.com/nathancurry/homelab.ansible/tree/1f437d2a033d77d9ab7335db67ad67b67931c00f)查看相关的git commit [。](https://github.com/nathancurry/homelab.ansible/tree/1f437d2a033d77d9ab7335db67ad67b67931c00f)

## 结构体

这是该项目中的文件。调用角色的任务，以及主/从角色。

```bash
tasks
├── configure_time.yml
roles/
├── chrony-master
│   ├── handlers
│   │   └── main.yml
│   ├── tasks
│   │   └── main.yml
│   ├── templates
│   │   ├── chrony.keys.j2
│   │   └── chrony-master.conf.j2
│   └── vars
│       └── main.yml
├── chrony-slave
    ├── handlers
    │   └── main.yml
    ├── tasks
    │   └── main.yml
    ├── templates
    │   ├── chrony.conf.j2
    │   └── chrony.keys.j2
    └── vars
       └── main.yml
```

## configure_time.yml

该脚本仅调用角色。我的时间服务器是Proxmox节点，因为VM时间不稳定。

```yml
---
- hosts: proxmox
  gather_facts: false
  user: root
# Debian masters
  roles:
    - chrony-master


- hosts: all:!proxmox
  gather_facts: false
  user: root
# CentOS slaves
  roles:
    - chrony-slave
```

## 任务/ main.yml

这些对于主服务器和从服务器基本上是相同的，因此这是主文件：

```yml
---
  - name: Include vars
    include_vars: ../vars/main.yml

  - name: Include vars
    include_vars: ~/0/vault/secrets.yml

  - name: Set timezone to America/Denver
    timezone:
      name: America/Denver
    notify: restart crond

  - name: ensure ntpdate isn't installed
    package:
      name: 'ntpdate'
      state: absent

  - name: install chrony
    package:
      name: chrony
      state: present

  # I mask instead of uninstall, because some services  
  # depend on NTP
  - name: Mask NTP service
    systemd:
      name: ntpd
      enabled: no
      masked: yes
      state: stopped

  - name: Ensure chronyd is active
    service:
      name: chronyd
      state: started

 # Debian config path.  Change for CentOS
  - name: install chrony.conf
    template:
      src: ../templates/chrony-master.conf.j2
      dest: /etc/chrony/chrony.conf
      owner: root
      group: root
      mode: 0644
    notify: restart chronyd

  - name: install chrony.keys
    template:
      src: ../templates/chrony.keys.j2
      dest: /etc/chrony/chrony.keys
      owner: root
      group: root
      mode: 0640
    notify: restart chronyd
```

## 范本

我基本上有两个模板。密钥文件实际上只是从Vault中提取变量。

主服务器和从服务器之间的配置文件略有不同，在此我将再次显示主服务器。差异变量是每个服务器本身都没有对等项。

我将IP用于服务器对等方，将主机名用于客户端。

```jinja
# Use public servers from the pool.ntp.org project.
# Please consider joining the pool (http://www.pool.ntp.org/join.html).
server 0.debian.pool.ntp.org iburst xleave
server 1.debian.pool.ntp.org iburst xleave
server 2.debian.pool.ntp.org iburst xleave
{% for host in groups['proxmox'] | difference([inventory_hostname]) %}
peer {{ hostvars[host].ansible_host }} iburst xleave
{% endfor %}

# Record the rate at which the system clock gains/losses time.
driftfile /var/lib/chrony/drift

# Allow the system clock to be stepped in the first three updates
# if its offset is larger than 1 second.
makestep 1.0 3

# This is disabled on the slaves, since VMs don't have an RTC
rtcsync

# This likely does nothing since my network interface is a bridge
# device.  However, if things change, this is good practice.
# Disabled on slaves
hwtimestamp *

# Increase the minimum number of selectable sources required to adjust
# the system clock.
#minsources 2

# Allow NTP client access from local network.
# This is of course not enabled on the slaves
allow {{ network }}
allow fe80::/64

# Serve time even if not synchronized to a time source.
# In case my network gets isolated, I still want
# Kerberos to work.
local stratum 10

# Specify file containing keys for NTP authentication.
keyfile /etc/chrony/chrony.keys

# Specify directory for log files.
logdir /var/log/chrony

# Select which information is logged.
#log measurements statistics tracking
```

## 测验

在我的笔记本电脑上测试：

```tcsh
chronyc> sources
210 Number of sources = 3
MS Name/IP address         Stratum Poll Reach LastRx Last sample               
===============================================================================
^* gold.lan.nathancurry.com     11   6   377    72  +1924us[+1974us] +/- 9101us
^- silver.lan.nathancurry.c>    12   6   377    75   +262us[ +311us] +/- 6356us
^+ bronze.lan.nathancurry.c>    10   6   377    74   +880us[ +930us] +/-   16ms
```

在时间服务器上：

```tcsh
chronyc> sources
210 Number of sources = 5
MS Name/IP address         Stratum Poll Reach LastRx Last sample               
===============================================================================
=- silver.lan.nathancurry.c>     3   7   377  1062   +865us[ +595us] +/-   60ms
=* bronze.lan.nathancurry.c>    10   7   377   158    -15us[+2100ns] +/-   15ms
^- h113.95.219.67.cable.sta>     2   7   377   183    +10ms[  +10ms] +/-   85ms
^- 12.167.151.2                  3   7   377    52    +10ms[  +10ms] +/-   59ms
^- startkeylogger.hungrycat>     3   6   377    52  +9268us[+9289us] +/-   33ms
```

真甜

您可以在[此处](https://github.com/nathancurry/homelab.ansible/tree/1f437d2a033d77d9ab7335db67ad67b67931c00f)找到相关的git commit [。](https://github.com/nathancurry/homelab.ansible/tree/1f437d2a033d77d9ab7335db67ad67b67931c00f)

 分类：[Ansible](https://www.nathancurry.com/categories/ansible/)，[Chrony](https://www.nathancurry.com/categories/chrony/)

[ 以前上一篇：本地仓库的Ansible角色](https://www.nathancurry.com/blog/19-ansible-role-for-local-repo/)

[下一篇：Ansible产生Kickstarts](https://www.nathancurry.com/blog/21-generate-kickstarts-with-ansible/)