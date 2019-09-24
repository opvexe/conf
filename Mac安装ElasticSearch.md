# Mac 安装 ElasticSearch

************ ElasticSearch ，Kibana ，中文分词器sego ************

#### 安装ElasticSearch [**链接**](https://www.jianshu.com/p/b128a880436d)
Elasticsearch 是一款稳定高效的分布式搜索和分析引擎，它的底层基于 Lucene，并提供了友好的 RESTful API 来对数据进行操作，还有比较重要的一点是， Elasticsearch 开箱即可用，上手也比较容易。
> **docker search elasticsearch**    ****搜索镜像****</br>
> **docker pull elasticsearch:5.1.1 **  **** 拉取elasticsearch镜像 *****</br>
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


## 使用ElasticSearch API 实现CRUD
添加索引：
```
PUT /lib/

{

  "settings":{
  
      "index":{
      
        "number_of_shards": 5,
        
        "number_of_replicas": 1
        
        }
        
      }
}

PUT  lib
```

查看索引信息:

```
GET /lib/_settings

GET _all/_settings


```

添加文档:

```

PUT /lib/user/1

{
    "first_name" :  "Jane",
    
    "last_name" :   "Smith",
    
    "age" :         32,
    
    "about" :       "I like to collect rock albums",
    
    "interests":  [ "music" ]
}

POST /lib/user/

{
    "first_name" :  "Douglas",
    
    "last_name" :   "Fir",
    
    "age" :         23,
    
    "about":        "I like to build cabinets",
    
    "interests":  [ "forestry" ]
    
}

```
查看文档:

```
GET /lib/user/1

GET /lib/user/

GET /lib/user/1?_source=age,interests

```


更新文档:

```
PUT /lib/user/1

{
    "first_name" :  "Jane",
    
    "last_name" :   "Smith",
    
    "age" :         36,
    
    "about" :       "I like to collect rock albums",
    
    "interests":  [ "music" ]
}

POST /lib/user/1/_update

{

  "doc":{
  
      "age":33
      
      }
}
```

删除一个文档

```

DELETE /lib/user/1

```


删除一个索引

```
DELETE /lib

```



## 批量获取文档

使用es提供的Multi Get API：

使用Multi Get API可以通过索引名、类型名、文档id一次得到一个文档集合，文档可以来自同一个索引库，也可以来自不同索引库

使用curl命令：

```
curl 'http://192.168.25.131:9200/_mget' -d '{

"docs"：[

   {
   
    "_index": "lib",
    
    "_type": "user",
    
    "_id": 1
    
   },
   
   {
   
     "_index": "lib",
     
     "_type": "user",
     
     "_id": 2
     
   }

  ]
}'


```

在客户端工具中：

```


GET /_mget

{
   
    "docs":[
       
       {
           "_index": "lib",
           "_type": "user",
           "_id": 1
       },
       {
           "_index": "lib",
           "_type": "user",
           "_id": 2
       },
       {
           "_index": "lib",
           "_type": "user",
           "_id": 3
       }
       
     ]
}

```