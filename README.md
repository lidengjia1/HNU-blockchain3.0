# 区块链学习小组

## 具体任务

### 第一组：区块链基础设施架构及云平台构建小组

负责人：周中定<br>
成员：胡乎虎，熊元坤，黄浩宇，李登佳

<br>

目标：区块链基础架构技术及云平台技术学习，完成区块链云平台搭建

### 第二组：区块链底层及前端展示小组

负责人：匡先华
<br>
成员：王鑫，程林海，孙霖，陈勋奇

<br>目标：区块链底层技术及前端展示技术学习，掌握多种区块链架构体系。

### 第三组：区块链新风险及智能合约编写小组

负责人：李信儒
<br>
成员：杨婷，陈冬晴及硕士研究生

<br>目标：完善专利及掌握区块链智能合约编写，掌握区块链技术在供应链金融及管理领域的应用。

### 第四组：区块链与供应链金融学术研究小组

负责人：林子君，孔晓琳

<br>

目录：书籍，论文，白皮书

<br>

目标：区块链与供应链金融，非对称信息博弈理论（特定分支），区块链信息加密理论。

### 第五组：教材编写小组

学习资料（图书类）

<br>

主编：马超群，周中定，李信儒，兰秋军

<br>

序言:
<br>
第一章：
<br>
第二章 
<br>
第三章
<br>
第四章
<br>

### 学习资料目录

目录：网课链接，计算机软件分享，博客分享

### 湖南大学区块链3.0工程源码

目录：



### 使用手册
master分支的合并由管理员完成，每周按照每个小组的进度在各自目录下更新
<br>
目前创建了四个分支供项目组使用
<br>
第一组及第五组使用分支a
<br>
第二组使用分支b
<br>
第三组使用分支c
<br>
第四组使用分支d



### git操作

#### 注册git账号以及gitee账号

```
https://gitee.com/ 码云网址
```

```
https://github.com/ GitHub网址
```

#### 下载git的window客户端

```
https://git-scm.com/
```

![1594713379(1)](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\HNU-blockchain3.0\区块链教材编写小组\图片\1594713379(1).png)



#### 生成本地公钥

如何生成本地公钥
执行以下语句来生成本地公钥

```
ssh-keygen -t rsa -C " "
```


这个指令会要求你提供一个 位置和文件名 去存放键值对和密码，可以一直点击 Enter键 去使用默认值。

![20190811132917176](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\HNU-blockchain3.0\区块链教材编写小组\图片\20190811132917176.png)

提示1：最好的情况是一个密码对应一个ssh key，但也不是非得这样去做，就像上面我们跳过创建密码这个步骤。
提示2：设置的密码不能被修改，也不可以被获取。

此时你按照上述路径 C:/Users/Admin/.ssh，找到该文件夹，如下图所示

![20190811133557754](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\HNU-blockchain3.0\区块链教材编写小组\图片\20190811133557754.png)

```
将excel发送给我
```

#### 管理员的工作

- 立项：克隆远程仓库+配置身份信息+创建项目+推送项目到远程仓库

- 1.克隆远程仓库的命令

  ```
  git clone git@github.com:lidengjia1/HNU-blockchain3.0.git 使用ssh
  ```

  ![微信图片编辑_20200714160139](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\微信图片编辑_20200714160139.jpg)

- 2.配置管理员身份信息

  ```
    cd Desktop/manager/info/
    git config user.name 'kuangxianhua'
    git config user.email '4564641897@qq.com'
  ```

- 3.日常管理

  ```
  日常管理可以往这个项目中增加文件（PDF，图片，office三件套，代码，链接）
  注意git记录的是增删操作，如果建立的空文件夹等于没有更新操作
  ```

  ![image-20200714161921546](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\HNU-blockchain3.0\区块链教材编写小组\图片\image-20200714161921546.png)

- 4.推送项目到远程仓库

  ```
    #查看当前项目变更状态
    git status
    # 工作区添加到暂存区
    git add .
    # 暂存区提交到仓库区
    git commit -m '立项'
    # 推送到远程仓库
    git push
  ```

![微信图片编辑_20200714163547](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\HNU-blockchain3.0\区块链教材编写小组\图片\微信图片编辑_20200714163547.jpg)

- 5.在 push 的时候需要设置账号与密码，该密码则是 github 的账号与密码

  如果在每次 push 都需要设置账号与密码，那么可以设置记住密码

  ```
  设置记住密码（默认15分钟）：
  git config --global credential.helper cache
  如果想自己设置时间，可以这样做(1小时后失效)：
  git config credential.helper 'cache --timeout=3600'
  长期存储密码：
  git config --global credential.helper store
  ```


6.将本地仓库关联不同的远程仓库

查看本地仓库关联远程仓库

```
git remote -v
```

![1594716868(1)](C:\Users\Administrator.DESKTOP-I1RC48K\Desktop\HNU-blockchain3.0\区块链教材编写小组\图片\1594716868(1).png)

## 方法1：每次`push`、`pull`时需分开操作

首先，查看本地仓库所关联的远程仓库：(假定最初仅关联了一个远程仓库)

```bash
$ git remote -v
origin  git@github.com:keithnull/keithnull.github.io.git (fetch)
origin  git@github.com:keithnull/keithnull.github.io.git (push)
```

然后，用`git remote add <name> <url>`添加一个远程仓库，其中`name`可以任意指定（对应上面的`origin`部分），比如：

```bash
$ git remote add coding.net git@git.coding.net:KeithNull/keithnull.github.io.git
```

再次查看本地仓库所关联的远程仓库，可以发现成功关联了两个远程仓库：

```bash
$ git remote -v
coding.net      git@git.coding.net:KeithNull/keithnull.github.io.git (fetch)
coding.net      git@git.coding.net:KeithNull/keithnull.github.io.git (push)
origin  git@github.com:keithnull/keithnull.github.io.git (fetch)
origin  git@github.com:keithnull/keithnull.github.io.git (push)
```

此后，若需进行`push`操作，则需要指定目标仓库，`git push <repo> <branch>`，对这两个远程仓库分别操作：

```bash
$ git push origin master
$ git push coding.net master
```

同理，`pull`操作也需要指定从哪个远程仓库拉取，`git pull <repo> <branch>`，从这两个仓库中选择其一:

```bash
$ git pull origin master
$ git pull coding.net master
```

## 方法2：`push`和`pull`无需额外操作

在方法1中，由于我们添加了多个远程仓库，在`push`和`pull`时便面临了仓库的选择问题。诚然如此较为严谨，但是在许多情况下，我们只需要保持远程仓库完全一致，而不需要进行区分，因而这样的区分便显得有些“多余”。

同样地，先查看已有的远程仓库：(假定最初仅关联了一个远程仓库)

```bash
$ git remote -v
origin  git@github.com:keithnull/keithnull.github.io.git (fetch)
origin  git@github.com:keithnull/keithnull.github.io.git (push)
```

然后，**不额外添加远程仓库，而是给现有的远程仓库添加额外的URL**。使用`git remote set-url -add <name> <url>`，给已有的名为`name`的远程仓库添加一个远程地址，比如：

```bash
$ git remote set-url --add origin git@git.coding.net:KeithNull/keithnull.github.io.git
```

再次查看所关联的远程仓库：

```bash
$ git remote -v
origin  git@github.com:keithnull/keithnull.github.io.git (fetch)
origin  git@github.com:keithnull/keithnull.github.io.git (push)
origin  git@git.coding.net:KeithNull/keithnull.github.io.git (push)
```

可以看到，我们并没有如方法1一般增加远程仓库的数目，而是给一个远程仓库赋予了多个地址（或者准确地说，多个用于`push`的地址）。

因此，这样设置后的`push` 和`pull`操作与最初的操作完全一致，不需要进行调整。

### git的命令指南

```
git checkout 123.py 恢复还没提交到缓存区的文件
```

```
git checkout dev 进入到dev分支
```

```
git checkout -b dev 创建并进入到dev分支
```

```
git reset --hard 版本号
```

```
git reset --HEAD file
```

```
git rerflog 记录的是每一次针对commit修改的操作，所以从仓库撤回时采用的是reflog的日志 git reset --hard 
```

```
git switch -c feature1
```

```
git branch -d feature1 删除分支 -D
```

```
git status 查看当前代码状态
```

```
git branch 查看当前分支
```

#### 注意点

- 查看远程库信息，使用`git remote -v`；
- 本地新建的分支如果不推送到远程，对其他人就是不可见的；
- 从本地推送分支，使用`git push origin branch-name`，如果推送失败，先用`git pull`抓取远程的新提交；
- 在本地创建和远程分支对应的分支，使用`git checkout -b branch-name origin/branch-name`，本地和远程分支的名称最好一致；
- 建立本地分支和远程分支的关联，使用`git branch --set-upstream branch-name origin/branch-name`；
- 从远程抓取分支，使用`git pull`，如果有冲突，要先处理冲突。
- 处理分支冲突的时候，每一次提交前，将本组分支pull下来进行更新操作，对于本地仓库的使用推荐使用
- git的原理及操作视频学习可以参考廖雪峰老师网站：https://www.liaoxuefeng.com/wiki/896043488029600




