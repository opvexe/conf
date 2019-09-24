# Mac 安装 ElasticSearch

#### 安装ElasticSearch [**链接**](https://www.jianshu.com/p/b128a880436d)
Elasticsearch 是一款稳定高效的分布式搜索和分析引擎，它的底层基于 Lucene，并提供了友好的 RESTful API 来对数据进行操作，还有比较重要的一点是， Elasticsearch 开箱即可用，上手也比较容易。
> **docker search elasticsearch**    ****搜索镜像****</br>
> **docker pull elasticsearch 5.1.1**  **** 拉取elasticsearch镜像 *****</br>
> **docker run -id --name esearch -p 9200:9200 -p 9300:9300 elasticsearch** **** 完成后，启动Elasticsearch*** </br>
> **docker ps***    **** 检查启动后的ElasticSearch启动镜像状态****</br>


```
**通过宿主机的IP访问  如：http://宿主机IP:9200/****检查ElasticSearch的web服务运行情况**
```


Elasticsearch 启动后，也启动了两个端口 9200 和 9300：

 + 9200 端口：HTTP RESTful 接口的通讯端口
 + 9300 端口：TCP 通讯端口，用于集群间节点通信和与 Java 客户端通信的端口
 

现在，让我们做一些测试。在浏览器访问链接 http://localhost:9200/ ，或使用 curl 命令：

```
 curl 'http://localhost:9200/?pretty'
```

#### 安装Kibana
Kibana是ES的一个配套工具，让用户在网页中可以直接与ES进行交互。
> **brew install kibana**</br>
> **kibana** </br>
```
Kibana的默认端口是5601
```
####  安装 httpie   [**链接**](https://httpie.org/doc#installation) curl也可以
> **brew install httpie**</br>
> **port install httpie**</br>