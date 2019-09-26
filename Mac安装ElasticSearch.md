# Mac 安装 ElasticSearch  
[参考文档1](https://juejin.im/entry/5913d6132f301e006b82e068)</br>
[参考文档2](https://blog.csdn.net/wd2014610/article/details/82426863)</br>

>ElasticSearch安装及环境配置，Kibana安装及环境配置，中文分词器安装 

### 安装ElasticSearch 

[**参考安装**](https://www.jianshu.com/p/b128a880436d)

Elasticsearch 是一款稳定高效的分布式搜索和分析引擎，它的底层基于 Lucene，并提供了友好的 RESTful API 来对数据进行操作，还有比较重要的一点是， Elasticsearch 开箱即可用，上手也比较容易。

```
1.docker 搜索镜像
docker search elasticsearch

2.拉取elasticsearch镜像
docker pull elasticsearch:6.8.3

3.启动Elasticsearch
docker run -id --name esearch -p 9200:9200 -p 9300:9300 elasticsearch:6.8.3

4.检查启动后的ElasticSearch启动镜像状态
docker ps -a 

5.检查ElasticSearch的web服务运行情况
curl  http://宿主机IP:9200/

6.Elasticsearch 启动后，也启动了两个端口 9200 和 9300：
解释：
	1.9200 端口：HTTP RESTful 接口的通讯端口
	2.9300 端口：TCP 通讯端口，用于集群间节点通信和与 Java 客户端通信的端口
	
测试:浏览器访问-->http://localhost:9200/ 或如下访问
curl 'http://localhost:9200/?pretty'


7.导入数据:  取得本地数据[/Users/facebook/Downloads/logs.jsonl]
curl -H 'Content-Type: application/x-ndjson' -XPOST 'localhost:9200/_bulk?pretty' --data-binary  @/Users/facebook/Downloads/logs.jsonl


8.查看数据:
curl -XGET 'localhost:9200/_cat/indices?v&pretty'

```
[**数据导入参考文档**](https://www.jianshu.com/p/76cc57d46328)

如果ElasticSearch服务停止或是挂掉，先使用docker删除对应的进程：docker rm ae89feb13d62

### docker容器关闭异常处理
```
1.查看进程
	docker  ps  - a 
2.杀死没用，删除
	docker rm ae89feb13d62
3.查看删除
	docker ps  -l 
```

### 安装Kibana 6.8.3版本
>Kibana是ES的一个配套工具，让用户在网页中可以直接与ES进行交互。 Kibana的默认端口是5601

```
1.安装
brew install kibana
2.启动
kibana

```

### 安装httpie或使用curl
[参考文档](https://httpie.org/doc#installation)</br>

```
1.安装
brew install httpie
2.启动
port install httpie
```
