# GmSSL项目

## 支持国密SM2 / SM3 / SM4 / SM9 / ZUC / SSL的密码工具箱

[快速入门](http://gmssl.org/docs/quickstart.html) [文档](http://gmssl.org/docs/docindex.html) [新闻](http://gmssl.org/news/newsindex.html) [下载](https://github.com/guanzhi/GmSSL/releases) [英文版](http://gmssl.org/english.html) [关于我们](http://gmssl.org/members.html)

## 快速上手

快速上手指南介绍GmSSL的编译，安装和`gmssl`命令行工具的基本指令。

1. 下载源代码（[zip](https://github.com/guanzhi/GmSSL/archive/master.zip)），解压缩至当前工作目录

   ```
   $ unzip GmSSL-master.zip
   ```

2. 编译与安装

   Linux平台（其他平台的安装过程见[编译与安装](http://gmssl.org/)）

   ```
   $ ./config no-saf no-sdf no-skf no-sof no-zuc
   $ make
   $ sudo make install
   ```

   安装之后可以执行`gmssl`命令行工具检查是否成功

   ```
   $ gmssl version
   GmSSL 2.0 - OpenSSL 1.1.0d
   ```

   如果编译遇到

   ```
   error： while loading shared libraries: libssl.so.1.1: cannot open shared object file: No such file or directory
   gmssl: symbol BIO_debug_callback version OPENSSL_1_1_0d not defined in file libcrypto.so.1.1 with link time reference
   ```

   此类问题，可通过如下命令解决

   ```
   $ export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
   ```

3. SM4加密文件

   ```
   $ gmssl sms4 -e -in <yourfile> -out <yourfile>.sms4
   enter sms4-cbc encryption password: <your-password>
   Verifying - enter sms4-cbc encryption password: <your-password>
   ```

   解密

   ```
   $ gmssl sms4 -d -in <yourfile>.sms4
   enter sms4-cbc decryption password: <your-password>
   ```

4. 生成SM3摘要

   ```
   $ gmssl sm3 <yourfile>
   SM3(yourfile)= 66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0
   ```

5. 生成SM2密钥并签名

   ```
   $ gmssl genpkey -algorithm EC -pkeyopt ec_paramgen_curve:sm2p256v1 \
                   -out signkey.pem
   $ gmssl pkeyutl -sign -pkeyopt ec_scheme:sm2 -inkey signkey.pem \
                   -in <yourfile> -out <yourfile>.sig
   ```

   将可以从公钥`signkey.pem`中导出并发发布给验证签名的一方

   ```
   $ gmssl pkey -pubout -in signkey.pem -out vrfykey.pem
   $ gmssl pkeyutl -verify -pkeyopt ec_scheme:sm2 -pubin -inkey vrfykey.pem \
                   -in <yourfile> -sigfile <yourfile>.sig
   ```

6. 生成SM2私钥及证书请求

   ```
   $ gmssl ecparam -genkey -name sm2p256v1 -text -out user.key
   $ gmssl req -new -key user.key -out user.req
   ```

   查看证书请求内容：

   ```
   $ gmssl req -in user.req -noout -text -subject
   ```

**[GmSSL项目](http://gmssl.org/)由[Zhi Guan](https://github.com/guanzhi)维护。**版权所有© [2014-2017 GmSSL项目](http://gmssl.org/)，保留所有权利。