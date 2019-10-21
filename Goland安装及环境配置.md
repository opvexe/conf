# Goland安装及环境配置

## 1.安装Go及Goland

### 1.1 下载Goland 

```json
// GoLand官网下载地址
https://www.jetbrains.com/go/download/#section=linux
```

### 1.2下载Go的包

```shell
# 下载Go的SDK包
$ sudo apt-get install golang-go
```

### 1.3liunx 文件目录结构
```shell
# liunx 文件目录介绍
用户下载软件：/usr/local/ 
环境变量：/etc/
日志：/var/log/
```



## 2. Go环境准备

### 2.1 配置环境变量
- 笔者是Mac系统，安装Go有多种方式，通过brew、下载源码安装go等方式可以安装go。

  在bash_profile中自定义GOPATH和GOBIN位置：

>1、打开终端输入cd ~进入用户主目录;
>2、输入ls -all命令查看是否存在.bash_profile;
>3、存在既使用vim .bash_profile 打开文件;
>4、输入 i 进入vim编辑模式；
>5、输入下面代码，

```shell
GOROOT=/usr/local/go
export GOPATH=/Users/user/aoho/go-workspace
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
```

- 安装完成之后，查看go的环境变量：

  `终端输入: go env`

```shell
GOARCH="amd64"
GOBIN="/usr/local/go/bin/go"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/user/aoho/go-workspace"
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/st/gkm45qzd2tv8mc32my38_n_00000gp/T/go-build646095787=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```


### 2.2GOPATH和GOROOT解释说明

* GOROOT不是必须要设置的。默认go会安装在/usr/local/go下，但也允许自定义安装位置，GOROOT的目的就是告知go当前的安装位置，编译的时候从GOROOT去找SDK的system libariry。

* GOPATH必须要设置，但并不是固定不变的。GOPATH的目的是为了告知go，需要代码的时候，去哪里查找。注意这里的代码，包括本项目和引用外部项目的代码。GOPATH可以随着项目的不同而重新设置。

GOPATH下会有3个目录：src、bin、pkg。

- src目录：go编译时查找代码的地方；
- bin目录：go get这种bin工具的时候，二进制文件下载的目的地；
- pkg目录：编译生成的lib文件存储的地方。
- 

[**安装Go文献**](https://juejin.im/post/5c6ac37cf265da2de7134242)

