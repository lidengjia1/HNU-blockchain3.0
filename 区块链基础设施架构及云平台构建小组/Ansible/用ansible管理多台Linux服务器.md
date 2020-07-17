# 用ansible管理多台Linux服务器

Ansible是个很流行的主机管理工具，不象其它类似工具，如Chef, Puppet需要安装客户端，Ansible只需要主机支持SSH即可。

下面以Dibian/Ubuntu为例

安装(本地Linux)

```
sudo apt install ansible
```

远程主需要从本地拷贝ssh public key, 这也是常规操作了。对于不熟悉的同学，这个是要免去每次ssh登录输入密码，一是麻烦，二是安全。特别是管理多个主机时，你不能每台主机都查一遍密码表吧

ssh-keygen #按照提示输入即可，一路回车也没问题，它会生成~/.ssh/id_rsa.pub 文件ssh-copy-id <远程主机> #这个会把上一步生成的pub文件放到远程主机，按提示输入ssh密码即可

然后我们修改ansible配置文件 /etc/ansible/hosts， [clients] 代表一组主机的别名，这样就可以按组来操作。有些云平台象ec2 是需要用pem文件登录的。我们只要按第二台的格式写就了。ansible_user是登录用户名

```
[clients]192.168.1.250 111.111.111.111 ansible_user=ubuntu ansible_ssh_private_key_file=~/my.pem
```

修改好后，我们测试下。-m ping是指调用ansible ping模块， 就是类似我们常用的ping来测试主机是否连接成功，当看到SUCCESS字样时即代表模块执行成功。

```
ansible clients -m ping192.168.1.250 | SUCCESS => {"changed": false, "ping": "pong"}
```

通常我们登录远程主机时，第一步都是跑下apt更新下软件，这个是需要sudo权限和输入密码的。而用ansible的话就是加-K参数表示输入密码， 同时要加-b --become-user=root 代表sudo为root

```
ansible clients -b --become-user=root -K -m apt -a 'name=ruby'
```

有的同学就有问题了，为什么我们还要输入密码啊，那上百台主机怎么办啊。当然我们可以去掉sudo了，只要修改/etc/sudoers，加上NOPASSWD就好了。既然是远程操作，这一步自然可以用ansible批量操作。不过我们就要学习playbook了，这个是把一个常见命令封装起来，以便按需要调用。网上也有大量别人写的Ansible可供使用。

好的，下面就是一个写好的playbook，很简单，用正则查找suders文件里以%sudo打头的行，然后替换掉整个行。我们需要把它保存在一个yml文件里。

```
---- hosts: alltasks: - lineinfile: path: /etc/sudoers state: present regexp: '^%sudo' line: '%sudo ALL=(ALL) NOPASSWD: ALL' validate: 'visudo -cf %s'
```

然后用ansible-playbook调用它

```
ansible-playbook set_sudoer.yml -b -K
```

这样我们就可以放心使用所有命令了，比如看下系统log

```
ansible clients -a "tail /var/log/syslog" -f 10
```

