Go安装包下载：https://studygolang.com/dl

**安装步骤：**

1、下载go安装包到本地（下载地址：https://studygolang.com/dl）

2、在ubuntu根目录下，使用xshell上传本地包

```
rz       #上传后，使用ll命令查看当前目录是否有go安装包
```

3、解压本地包到<font color=red>/usr/local/go</font>目录下

```
sudo tar -C /usr/local -xzf go1.14.1.darwin-amd64.tar.gz
```

4、设置GOPATH和GOROOT环境变量

```
sudo vim /etc/profile
```

在profile文件最后添加如下内容:

```
export GOPATH=$HOME/go
export GOROOT=/usr/local/go      #此处需跟go语言包的安装路径一致（参考上述3）
export PATH=$GOROOT/bin:$PATH
```

5、使用 source 命令，使刚刚添加的配置信息生效：

```
source /etc/profile
```

6、使用go env查看所有go相关配置

```
go env
```

7、查看go语言包是否成功：

```
go version
```



