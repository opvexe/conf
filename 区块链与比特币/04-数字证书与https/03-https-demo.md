```sh
git@github.com:AdugiBeyond/golang-https-example.git
```



# 一、单向认证

- client认证server
- client端事先保存server的签发机构根证书，与server建立联系时，server需要向client发送自己的证书
- client无需向server发送证书



server证书生成：pem格式

```sh
openssl req \
    -x509 \
    -nodes \
    -newkey rsa:2048 \
    -keyout server.key \
    -out server.crt \
    -days 3650 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=Global Security/OU=IT Department/CN=*"
```



## 1. 写server.go

### -分析

```go
//1. 创建http server，
//- Addr
//- Handler
//- TLSConfig

//2. server启动服务， 同时加载自己的证书

```

### - 代码

```go
package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	//1. 创建http server，
	server := http.Server{
		Addr:      ":1234",
		Handler:   &myHandler{},
		TLSConfig: nil,
	}

	//2. 启动服务， 同时加载自己的证书
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("myFunc called!")
	w.Write([]byte("hello world!!!!"))
}

```



浏览器访问!



### - 补充(可选)

如果handler使用nil，则程序会使用默认的处理器

```go
	server := http.Server{
		Addr:      ":1234",
		Handler:   nil,
		TLSConfig: nil,
	}
```

实现下面的函数即可

```go
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("myFunc called!\n")
		writer.Write([]byte("hello world!!!!"))
	})
```



## 2. 写client.go

### - 分析

```go
//1. 注册server的ca证书
//- 读取证书
//- 加载证书到ca池

//2. 配置tls， tls.Config
//- 把ca池注册进去 ===> RootCAs

//3. 创建client ===> Transport， 指定cfg
//4. 发送请求
//5. 接收并打印返回值
```

### - 代码

```go
package main

import (
	"io/ioutil"
	"log"
	"crypto/x509"
	"crypto/tls"
	"net/http"
	"fmt"
)

func main() {
	//1. 注册server的ca证书
	//- 读取证书
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	//- 加载证书到ca池
	certPool := x509.NewCertPool()

	ok := certPool.AppendCertsFromPEM(caCert)
	if !ok {
		log.Fatal(err)
	}

	//2. 配置tls
	cfg := tls.Config{
		//配置ca池
		RootCAs: certPool,
	}

	//3. 创建client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &cfg,
		},
	}

	//4. 发送请求
	res, err := client.Get("https://localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	info, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//5. 接收并打印返回值
	fmt.Printf("body : %s\n", info)
	fmt.Printf("status : %s\n", res.Status)
}
```





# 二、双向认证



生成server证书

```sh
openssl req \
    -x509 \
    -nodes \
    -newkey rsa:2048 \
    -keyout server.key \
    -out server.crt \
    -days 3650 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=Global Security/OU=IT Department/CN=*"
```



生成client证书

```sh
openssl req \
    -x509 \
    -nodes \
    -newkey rsa:2048 \
    -keyout client.key \
    -out client.crt \
    -days 3650 \
    -subj "/C=GB/ST=ITCAST/L=ITCAST/O=Global Security/OU=IT Department/CN=*"
```



## 1. server.go

### - 分析

```go
//1. 注册客户端的ca证书
//- 读取client证书
//- 添加到ca池

//2.创建tls.config， 配置tls通信
//- 指明需要验证客户端， ClientAuth:tls
//- 添加caPool， ClientCAs:

//3. 创建http的server，
//- 端口
//- 处理器  =》 自己实现ServeHTTP函数， 也可以设置nil
//- TLS配置

//4. 启动服务，指定证书
```



### - 代码

```go
package main

import (
	"io/ioutil"
	"log"
	"crypto/x509"
	"crypto/tls"
	"net/http"
	"fmt"
)

func main() {
	//1. 注册client的ca证书，自己签名的，所以只需要把client的证书注册即可
	//-读取
	//caCert, err := ioutil.ReadFile("client.crt")
	caCert, err := ioutil.ReadFile("client.crt")
	if err != nil {
		log.Fatal("err : ", err)
	}

	//-加载到ca池
	caCerPool := x509.NewCertPool()
	caCerPool.AppendCertsFromPEM(caCert)

	//2. 配置tls通信
	cfg := tls.Config{
		//需要客户端的证书并且验证
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCerPool,
	}

	//3. 根据tls配置，创建http的server
	server := http.Server{
		Addr:      ":8843",
		Handler:   &handler{},
		TLSConfig: &cfg,
	}

	//4. 启动server
	err = server.ListenAndServeTLS("server.crt", "server.key")
	fmt.Printf(" err : %s\n", err)
}

type handler struct {
}

func (h *handler) ServeHTTP(ResponseWriter http.ResponseWriter, Request *http.Request) {
	resInfo := []byte("hello world!!!")
	fmt.Println(string(resInfo))
	ResponseWriter.Write(resInfo)
}

```



## 2. client.go

### - 分析

```go
//1. 注册server的ca证书
//- 读取证书
//- 加载证书到ca池

//1.5 加载自己的证书和私钥tls.Load  《======= 修改了
//- 证书传递给server
//- 私钥解密数据


//2. 配置tls， tls.Config
//- 把ca池注册进去 ===> RootCAs
//- 提供client自己的证书信息  《===========修改了 
// Certificates: []tls.Certificate{cert},

//3. 创建client ===> Transport， 指定cfg
//4. 发送请求
//5. 接收并打印返回值
```



### - 代码

```go
package main

import (
	"io/ioutil"
	"log"
	"crypto/x509"
	"crypto/tls"
	"net/http"
	"fmt"
)

//client代码
func main() {
	//1. 注册能够识别server证书的ca, 这个证书是自签名的，所以自己就是ca
	//- 读取ca
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	//- 注册到证书链
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	//2. 加载client自己的证书和私钥
	//- 证书是要传递给server
	//- 私钥是为了解开server用client公钥加密的信息
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	//3. 配置https的client为tls
	cfg := tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	//4. 创建http client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &cfg,
		},
	}

	//5. 发起http请求
	resp, err := client.Get("https://localhost:8843")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//6. 打印应答信息
	fmt.Printf("data : %s\n", data)
	fmt.Printf("status code: %s\n", resp.Status)
}
```

