> 这个项目由于国内无法正常访问glang/x，所以安装很波折



# 一、正常的安装方式

请参考这个链接:https://www.jianshu.com/p/14b0a9da931e

```sh
$ go get -u github.com/Masterminds/glide
$ git clone https://github.com/btcsuite/btcd $GOPATH/src/github.com/btcsuite/btcd
$ cd $GOPATH/src/github.com/btcsuite/btcd
$ glide install
$ go install . ./cmd/...
```







# 二、粗暴方式

但是我的有问题，所以使用粗暴的方式



>  网上的教程其实很简单，但是我的电脑显示的错误与教程不同，我尝试解决，但是没有成功，最终我选择了最笨的方式，手动安装所有的依赖包



## 1. 下载代码

```sh
git clone https://github.com/btcsuite/btcd $GOPATH/src/github.com/btcsuite/btcd

cd $GOPATH/src/github.com/btcsuite/btcd
```



##2. 安装依赖包

> 所有要安装的包都在下面，手动克隆即可

==注意事项：一定要按照层级目录来下载==

比如说`git clone https://github.com/davecgh/go-spew`， 对于这个工程，一定要先创建好`davechgh`目录

然后在里面执行克隆命令

```sh
git clone https://github.com/btcsuite/btclog
git clone https://github.com/btcsuite/websocket
git clone https://github.com/btcsuite/go-socks
git clone https://github.com/btcsuite/winsvc
git clone https://github.com/btcsuite/btcutil
git clone https://github.com/btcsuite/goleveldb

git clone https://github.com/davecgh/go-spew
git clone https://github.com/jessevdk/go-flags
git clone https://github.com/jrick/logrotate
git clone https://github.com/golang/crypto

git clone https://github.com/aead/siphash
git clone https://github.com/kkdai/bstream
```



## 3. 编译安装

最后切换到btcd目录，执行下面的命令

```sh
go install . ./cmd/...
```



## 4. 查看运行

编译好的程序会安装在$GOPATH/bin下面：

- btcd,
- btcctl
- addblock
- findcheckpoint
- gencerts









# 三、使用教程

https://zhuanlan.zhihu.com/p/33175991







# 四、基本交互

## 1. btcd

1. 创建btcd.conf

   ```sh
   /Users/duke/Library/Application Support/Btcd/btcd.conf
   ```

   

2. 填入配置信息

   ```sh
   #testnet=1 #比特币的公共测试网络
   regtest=1  #单机版本地网络
   txindex=1
   rpcuser=test
   rpcpass=test
   #rpccert=/Users/duke/Library/Application Support/Btcd/rpc.cert
   ```

3. 启动， 这个端口自己指定，否则命令行交互时端口不对

   ```sh
   btcd --rpclisten=127.0.0.1:8334
   ```

4. log如下

   ```sh
    duke ~/go/bin$  btcd --rpclisten=127.0.0.1:8334
   2019-03-10 22:04:36.667 [INF] BTCD: Version 0.12.0-beta
   2019-03-10 22:04:36.667 [INF] BTCD: Removing regression test database from '/Users/duke/Library/Application Support/Btcd/data/regtest/blocks_ffldb'
   2019-03-10 22:04:36.668 [INF] BTCD: Loading block database from '/Users/duke/Library/Application Support/Btcd/data/regtest/blocks_ffldb'
   2019-03-10 22:04:36.672 [INF] BTCD: Block database loaded
   2019-03-10 22:04:36.678 [INF] INDX: Transaction index is enabled
   2019-03-10 22:04:36.678 [INF] INDX: Committed filter index is enabled
   2019-03-10 22:04:36.679 [INF] INDX: Catching up indexes from height -1 to 0
   2019-03-10 22:04:36.679 [INF] INDX: Indexes caught up to height 0
   2019-03-10 22:04:36.679 [INF] CHAN: Chain state (height 0, hash 0f9188f13cb7b2c71f2a335e3a4fc328bf5beb436012afca590b1a11466e2206, totaltx 1, work 2)
   2019-03-10 22:04:36.692 [INF] RPCS: RPC server listening on 127.0.0.1:8334
   2019-03-10 22:04:36.693 [INF] AMGR: Loaded 0 addresses from file '/Users/duke/Library/Application Support/Btcd/data/regtest/peers.json'
   2019-03-10 22:04:36.694 [INF] CMGR: Server listening on 0.0.0.0:18444
   2019-03-10 22:04:36.694 [INF] CMGR: Server listening on [::]:18444
   
   
   ```

   



## 2.btcctl

1. 创建btcctl.conf

   ```sh
   /Users/duke/Library/Application Support/Btcctl/btcd.conf
   ```

   

2. 填入配置信息

   ```sh
   rpcuser=test
   rpcpass=test
   ```

   

3. 执行测试命令

   ```sh
    duke ~/go/bin$  btcctl getinfo
   {
     "version": 120000,
     "protocolversion": 70002,
     "blocks": 0,
     "timeoffset": 0,
     "connections": 0,
     "proxy": "",
     "difficulty": 1,
     "testnet": false,
     "relayfee": 0.00001,
     "errors": ""
   }
   ```

4. 补充

   ```sh
   - btcctl默认端口是8334， 但是btcd根据不同网络启动端口不一致
   - 不清楚为何btcctl不能够指定启动端口
   ```

   