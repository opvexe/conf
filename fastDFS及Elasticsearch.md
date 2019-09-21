

# fastDFS 文件图片上传 

#### 3.2.利用fastDFS存储图片

##### 3.2.1什么是FastDFS

FastDFS 是用 c 语言编写的一款开源的分布式文件系统。FastDFS 为互联网量身定制， 充分考虑了冗余备份、负载均衡、线性扩容等机制，并注重高可用、高性能等指标，使用 FastDFS 很容易搭建一套高性能的文件服务器集群提供文件上传、下载等服务。

优点： 

FastDFS 架构包括 Tracker server 和 Storage server。客户端请求 Tracker server 进行文 件上传、下载，通过 Tracker server 调度最终由 Storage server 完成文件上传和下载。 

Tracker server 作用是负载均衡和调度，通过 Tracker server 在文件上传时可以根据一些 方法找到 Storage server 提供文件上传服务。可以将 tracker 称为追踪服务器或调度服务 器。 

Storage server 作用是文件存储，客户端上传的文件最终存储在 Storage 服务器上， Storageserver 没有实现自己的文件系统而是利用操作系统 的文件系统来管理文件。可以将 storage 称为存储服务器。 

![1538157518483](./assets/1538157518483.png)

服务端两个角色: 

Tracker:管理集群，tracker 也可以实现集群。每个 tracker 节点地位平等。收集 Storage 集群的状态。 

Storage:实际保存文件 

Storage 分为多个组，每个组之间保存的文件是不同的。每 个组内部可以有多个成员，组成员内部保存的内容是一样的，组成员的地位是一致的，没有 主从的概念。

##### 3.2.2文件上传流程 

![1538157604016](./assets/1538157604016.png)

客户端上传文件后存储服务器将文件 ID 返回给客户端，此文件 ID 用于以后访问该文 件的索引信息。文件索引信息包括:组名，虚拟磁盘路径，数据两级目录，文件名。

![1538157690533](./assets/1538157690533.png)

**组名**: 文件上传后所在的 storage 组名称，在文件上传成功后有 storage 服务器返回， 需要客户端自行保存。 

**虚拟磁盘路径**: storage 配置的虚拟路径，与磁盘选项 store_path*对应。如果配置了 store_path0 则是 M00，如果配置了 store_path1 则是 M01，以此类推。 

**数据两级目录 **:storage 服务器在每个虚拟磁盘路径下创建的两级目录，用于存储数据 文件。 

**文件名** :与文件上传时不同。是由存储服务器根据特定信息生成，文件名包含:源存储 服务器 IP 地址、文件创建时间戳、文件大小、随机数和文件拓展名等信息。

##### 3.2.3文件下载流程

![1538157946453](./assets/1538157946453.png)

##### 3.2.4简易FastDFS架构

![1538157991436](./assets/1538157991436.png)

##### 3.2.5FastDFS安装

###### 3.2.5.1安装FastDFS依赖包

1. 解压缩libfastcommon-master.zip   
2. 进入到libfastcommon-master的目录中
3. 执行**./make.sh**
4. sudo apt-get install make
5. 执行**sudo ./make.sh install**

###### 3.2.5.2安装FastDFS

1. 解压缩fastdfs-master.zip
2. 进入到 fastdfs-master目录中
3. 执行 **./make.sh**
4. 执行 **sudo ./make.sh install**

###### 3.2.5.3配置跟踪服务器tracker

1. ```shell
   sudo cp /etc/fdfs/tracker.conf.sample /etc/fdfs/tracker.conf
   ```

2. 在/home/itcast/目录中创建目录 fastdfs/tracker      

   ```shell
   mkdir -p /home/itcast/fastdfs/tracker
   ```

3. 编辑/etc/fdfs/tracker.conf配置文件    sudo vim /etc/fdfs/tracker.conf

​        修改 base_path=/home/itcast/fastdfs/tracker

###### 3.2.5.4配置存储服务器storage 

1. ```
   sudo cp /etc/fdfs/storage.conf.sample /etc/fdfs/storage.conf
   ```

2. 在/home/itcast/fastdfs/ 目录中创建目录 storage

   ```shell
   mkdir –p /home/itcast/fastdfs/storage
   ```

3. 编辑/etc/fdfs/storage.conf配置文件  sudo vim /etc/fdfs/storage.conf

   修改内容：

   ```shell
   base_path=/home/itcast/fastdfs/storage
   store_path0=/home/itcast/fastdfs/storage
   tracker_server=自己ubuntu虚拟机的ip地址:22122
   ```

   

###### 3.2.5.5启动tracker和storage

进入到/etc/fdfs/下面执行以下两条指令

```shell
sudo  fdfs_trackerd  /etc/fdfs/tracker.conf
sudo fdfs_storaged  /etc/fdfs/storage.conf
```

###### 3.2.5.6测试是否安装成功

1. **sudo cp /etc/fdfs/client.conf.sample /etc/fdfs/client.conf **
2. 编辑/etc/fdfs/client.conf配置文件  **sudo vim /etc/fdfs/client.conf**

修改内容：

```shell
base_path=/home/itcast/fastdfs/tracker
tracker_server=自己ubuntu虚拟机的ip地址:22122
```

3. 上传文件测试(fastDHT)

   sudo fdfs_upload_file /etc/fdfs/client.conf 要上传的图片文件 

   如果返回类似**group1/M00/00/00/rBIK6VcaP0aARXXvAAHrUgHEviQ394.jpg **的文件id则说明文件上传成功

###### 3.2.5.7安装fastdfs-nginx-module 

1. 解压缩 nginx-1.8.1.tar.gz

2. 解压缩 fastdfs-nginx-module-master.zip

3. 进入nginx-1.8.1目录中

4. 执行

   ```shell
   sudo ./configure  --prefix=/usr/local/nginx/ --add-module=fastdfs-nginx-module-master的解压后的目录的绝对路径/src
   ```

   注意：**这时候会报一个错，说没有PCRE库**

   ![1538183542474](./assets/1538183542474.png)

   下载缺少的库

   ```shell
   sudo apt-get install libpcre3 libpcre3-dev 
   ```

   + 首先你需要去更换源，因为ubuntu自带的源没有这个库

   + 更换下载源为阿里的源

   + 先把原来的源文件备份

     ```shell
     sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak
     ```

   + 编辑源文件

     ```shell
     sudo vim /etc/apt/sources.list
     ```

     把原来的内容全部删掉，粘贴一下内容：

     ```shell
     deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
     deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
      
     deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
     deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
      
     deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
     deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
      
     deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
     deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
      
     deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
     deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
     ```

     更换完源之后执行  

     ```shell
     sudo apt-get  update
     sudo apt-get install libpcre3 libpcre3-dev
     ```
     
     安装 zlib
     ```
     sudo apt-get install ruby
     sudo apt-get install zlib1g
     sudo apt-get install zlib1g.dev
     ```

     然后进入nginx-1.8.1目录中，再次执行：

     ```shell
     sudo ./configure  --prefix=/usr/local/nginx/ --add-module=/home/itcast/fastDFS/fastdfs-nginx-module-master/src
     ```

     然后编译：

     ```shell
     sudo make
     ```

     这时候还会报一个错（错误还真多），错误原因是因为nginx编译的时候把警告当错误处理，事实上这个警告并不影响（程序员忽略警告）：

     ![1538184263769](./assets/1538184263769.png)

     解决方法：

     找到objs目录下的Makefile

     vim Makefile

     删掉里面的-Werror(**如果没有修改权限，修改一下这个文件的权限,`chmod 777 Makefile`**)

     ![1538185926173](./assets/1538185926173.png)

     然后回到nginx-1.8.1目录中

     执行完成后继续执行**sudo make**

     执行**sudo make install** 

     

5. sudo cp fastdfs-nginx-module-master解压后的目录中src下mod_fastdfs.conf   /etc/fdfs/mod_fastdfs.conf

6. sudo vim /etc/fdfs/mod_fastdfs.conf

   修改内容：

   ```shell
   connect_timeout=10
   tracker_server=自己ubuntu虚拟机的ip地址:22122
   url_have_group_name=true
   store_path0=/home/itcast/fastdfs/storage
   ```

7. sudo cp 解压缩的fastdfs-master目录中的conf中的http.conf  /etc/fdfs/http.conf

8. sudo cp 解压缩的fastdfs-master目录中conf的mime.types /etc/fdfs/mime.types

9. sudo vim /usr/local/nginx/conf/nginx.conf

   在http部分中添加配置信息如下：

   ```shell
   server {
               listen       8888;
               server_name  localhost;
               location ~/group[0-9]/ {
                   ngx_fastdfs_module;
               }
               error_page   500 502 503 504  /50x.html;
               location = /50x.html {
               root   html;
               }
           }
   
   ```

10. 启动nginx

  sudo  /usr/local/nginx/sbin/nginx

##### 3.2.6使用go客户端上传文件测试

+ 下载包

  ```shell
  go get -u -v github.com/weilaihui/fdfs_client
  ```

  这时候会报一个错：

  ![1538186585971](./assets/1538186585971.png)

  这是因为我们的网络有长城防火墙，不能直接去google下载相应的包，所以就失败啦

  解决办法：

  + 在`~/workspace/go/src`目录下面创建一个golang.org/x目录

    ``` shell
    cd  ~/workspace/go/src
    mkdir -p golang.org/x
    ```

  + 进入golang.org/x下载两个包

    ```shell
    cd golang.org/x
    git clone https://github.com/golang/crypto.git
    git clone https://github.com/golang/sys.git
    ```

  + 然后再执行最初的下载命令

    ```shell
    go get github.com/weilaihui/fdfs_client
    ```

+ go操作fastDFS的方法

  + 先导包，把我们下载的包导入

    ```go
    import "github.com/weilaihui/fdfs_client"
    ```

  + 导包之后,我们需要指定配置文件生成客户端对象

    ```go
    client,_:=fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
    ```

  + 接着我们就可以通过client对象执行文件上传，上传有两种方法，一种是通过文件名，一种是通过字节流

    + 通过文件名上传**UploadByFilename **,参数是文件名（必须通过文件名能找到要上传的文件），返回值是fastDFS定义的一个结构体，包含组名和文件ID两部分内容

      ```go
      fdfsresponse,err := client.UploadByFilename("flieName")
      ```

    + 通过字节流上传**UploadByBuffer**,参数是字节数组和文件后缀，返回值和通过文件名上传一样。

      ```go
      fdfsresponse,err := client.UploadByBuffer(fileBuffer,ext)
      ```

#### 3.3.改用fastDFS上传文件

​	我们以前保存视图传递过来的文件方法是先用GetFile()获取文件相关信息，然后再用SaveToFile()把文件保存到 文件夹下面。以前保存图片的代码如下：

```go
func UploadImage(this*beego.Controller,filePath string)string{
	//1.那数据
	//那标题
	f,h,err:=this.GetFile(filePath)

	defer f.Close()
	//上传文件处理
	//1.判断文件格式
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png"&&ext != ".jpeg"{
		beego.Info("上传文件格式不正确")
		return ""
	}

	//2.文件大小
	if h.Size>5000000{
		beego.Info("文件太大，不允许上传")
		return ""
	}

	//3.不能重名
	fileName := time.Now().Format("2006-01-02 15:04:05")


	err2:=this.SaveToFile(filePath,"./static/img/"+fileName+ext)
	if err != nil{
		beego.Info("上传文件失败")
		return ""
	}

	if err2 != nil{
		beego.Info("上传文件失败",err2)
		return ""
	}
	return "/static/img/"+fileName+ext
}
```

我们现在要用fastDFS来存储图片，第一步也是先用GetFile拿到文件，但是第二步，我们用UploadByBuffer()把静态文件存到我们fastDFS文件系统中。

```go
//先导包
import "github.com/weilaihui/fdfs_client"

//通过GetFile获取文件信息
f,h,err := this.GetFile(filePath)
defer f.Close()
//然后对上传的文件进行格式和大小判断
//1.判断文件格式
ext := path.Ext(h.Filename)
if ext != ".jpg" && ext != ".png"&&ext != ".jpeg"{
	beego.Info("上传文件格式不正确")
	return ""
}

//2.文件大小
if h.Size>5000000{
	beego.Info("文件太大，不允许上传")
	return ""
}
//3.上传文件
//先获取一个[]byte
fileBuffer := make([]byte,h.Size)
//把文件数据读入到fileBuffer中
f.Read(fileBuffer)
//获取client对象
client := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
//上传
fdfsresponse,_:=client.UploadByBuffer(fileBuffer,ext[1:])
//返回文件ID
return fdfsresponse.RemoteFileId
```

上传图片方式修改完成之后我们添加类型就业务就结束了。

#### 3.4 类型展示

类型展示和案例中的展示业务一样，只是一些表名和字段名需要简单修改，代码如下：

```go
//1.读取类型表，显示数据
	o := orm.NewOrm()
	var goodsTypes[]models.GoodsType
	//查询
	_,err:=o.QueryTable("GoodsType").All(&goodsTypes)
	if err != nil{
		beego.Info("查询类型错误")
	}


	this.Data["title"] = "添加类型"
	this.Data["goodsTypes"] = goodsTypes
	this.Layout = "layout.html"
	this.TplName = "addType.html"
```

**到这里，我们添加类型界面就算完成了。但是我们后面还需要添加goodsSPU，添加商品，删除商品，修改商品等功能和我们以前写的业务基本一样，老师就不给你们写了，你们自己去实现一下相关业务。**

**我们这里就直接导入数据，不再一步步的添加数据**

#### 3.5导入商品有关数据

##### 3.5.1导入数据库数据

​	导入数据库数据，老师给你们的资料中，有一个文件是` dailyfresh.sql`,就是我以前导入的数据，你们可以把这个sql语句导入到你们的数据库中，直接拿来使用。

+ 先进入数据库中

  ```shell
  mysql -uroot -p123456
  ```

+ 选中项目中用到的数据库

  ```shell
  use pyg
  ```

+ 导入文件(保证dailyfresh.sql文件在你当前目录下面)

  ```shell
  source fresh.sql
  ```

+ 查看数据是否导入成功

  ```shell
  select * from goods_type;
  ```

##### 3.5.2导入图片数据

​	导入数据库数据之后我们还有很多图片内容是存在fastDFS中的，也需要我们手动导入

+ 删除` ~/fdfs/storage/data/00`目录下的00文件夹

+ 把课堂资料中的00.zip拷贝到我们存放图片的路径下面` ~/fdfs/storage/data/00`

  ![1538202366120](./assets/1538202366120.png)

+ 解压00.zip到当前目录

***这时候我们就把数据全部导入到了我们的开发环境当中。***

## 二.商品信息展示

前面我们导入了全部商品有关的数据，接着我们就来实现，商品展示有关的内容。

### 1.生鲜首页内容的展示

首先我们来看一下首页的页面。

![1557761180054](./assets/1557761180054.png)

商品页面分为如图所示的四部分，那我们就来分别获取这四部分的值，仍旧按照我们的四步骤来走。

**请求**

这里我们设定首页的请求路径为`\`

**路由**

然后我们去router.go中修改相应的路由适配

```go
 //显示主页面
 beego.Router("/",&controllers.GoodsController{},"get:ShowIndex")
```

**控制器 **

这里我们指定首页对应`GoodsController`控制器，获取首页的方法为`ShowIndex`,接着我们就来实现这个方法。

+ 获取类型

  ```go
  //查询商品类型
  	var goodsTypes []models.GoodsType
  	o.QueryTable("GoodsType").All(&goodsTypes)
  	this.Data["types"] = goodsTypes
  ```

+ 获取轮播商品数据

  ```go
  //查询轮播商品图片
  	var goodsBanner []models.IndexGoodsBanner
  	o.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&goodsBanner)
  	this.Data["goodsBanners"] = goodsBanner
  ```

+ 查询促销商品数据

  ```go
  //查询促销商品
  	var promotionBanner []models.IndexPromotionBanner
  	o.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&promotionBanner)
  	this.Data["proBanner"] = promotionBanner
  ```

+ 查询首页商品

  ![1539416660608](./assets/1539416660608.png)

  这个业务比较复杂，我们观察页面发现我们查询的时候要把类型对象和相应的商品对象切片放在一起，那我们用什么容器呢？分析如下图：

  ![1539418555357](./assets/1539418555357.png)

  分析之后我们发现，只能用interface这种类型，还需要给不同类型的数据加一个key值，所以我们需要用map[string]interface{}，这是一个 类型的内容，那所有类型的内容，我们应该给这个类型放到切片中，所以我们这里用**[]map[string]interface{}**这种类型来存储首页商品。

  确定了存储容器之后，我们开始查询数据来存储，首先是把商品类型存储到我们的容器当中：

  ```go
  goods := make([]map[string]interface{},len(goodsTypes))
  	for index,goodsType := range goodsTypes{
  		temp := make(map[string]interface{})
  		temp["type"] = goodsType
  		goods[index] = temp
  	}
  ```

  接着把商品存储到我们的容器中，注意这里是分为图片商品和文字商品，需要分开来存储：

  ```go
  var goodsImage []models.IndexTypeGoodsBanner
  var goodsText []models.IndexTypeGoodsBanner
  for _,temp := range goods{
    o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsSKU","GoodsType").Filter("GoodsType",temp["type"]).Filter("Display_Type",1).OrderBy("Index").All(&goodsImage)
    o.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsSKU","GoodsType").Filter("GoodsType",temp["type"]).Filter("Display_Type",0).OrderBy("Index").All(&goodsText)
      
  temp["goodsText"] = goodsText
  temp["goodsImage"] = goodsImage
  }
  ```

  然后再把数据传递给视图

  ```go
  this.Data["goods"] = goods
  ```

**视图 **

​	获取完数据就需要在页面里面显示，这里注意，我们获取的图片数据是FastDFS返回的ID值，这里我们访问的是通过nginx获取的图片数据，所以需要在图片地址上需要加上相应的ip和端口号,代码如下：

```html
	<div class="center_con clearfix">
		<ul class="subnav fl">
            <!--------循环展示类型数据---------->
			{{range .types}}
				<li><a href="" class="{{.Logo}}">{{.Name}}</a></li>
			{{end}}
		</ul>
		<div class="slide fl">
			<ul class="slide_pics">
                <!--------循环展示轮播图---------->
				{{range .goodsBanners}}
					<li><img src="http://192.168.110.81:8888/{{.Image}}" alt="幻灯片"></li>
				{{end}}
			</ul>
			<div class="prev"></div>
			<div class="next"></div>
			<ul class="points"></ul>
		</div>
		<div class="adv fl">
            <!--------循环展示促销商品数据---------->
			{{range .proBanner}}
				<a href="{{.Url}}"><img src="http://192.168.110.81:8888/{{.Image}}"></a>
			{{end}}
		</div>
	</div>
<!--------循环展示首页商品---------->
{{range .goods}}
	<div class="list_model">
		<div class="list_title clearfix">
			<h3 class="fl" id="model01">{{.type.Name}}</h3>
			<div class="subtitle fl">
				<span>|</span>
                <!--------循环展示文字商品数据---------->
				{{range .goodsText}}
				<a href="">{{.GoodsSKU.Name}}</a>
				{{end}}
			</div>
			<a href="" class="goods_more fr" id="fruit_more">查看更多 ></a>
		</div>

		<div class="goods_con clearfix">
			<div class="goods_banner fl"><img src="http://192.168.110.81:8888/{{.type.Image}}"></div>
			<ul class="goods_list fl">
                <!--------循环展示图片商品数据---------->
				{{range .goodsImage}}
				<li>
					<h4><a href="">{{.GoodsSKU.Name}}</a></h4>
					<a href=""><img src="http://192.168.110.81:8888/{{.GoodsSKU.Image}}"></a>
					<div class="prize">¥ {{.GoodsSKU.Price}}</div>
				</li>
				{{end}}
			</ul>
		</div>
	</div>
{{end}}
```

### 2.商品详情页的展示

首页展示之后，我们接着来实现商品详情页的展示！仍然是我们你的四步骤

**请求**

商品详情页的请求应该在首页展示的相应页面添加，不管是文字商品还是图片商品都应该添加。商品详情和我们实现案例中文章详情一样，都需要传递相应的Id，所以我们设计商品详情的路由为

`/goodsDetail?id={{.GoodsSKU.Id}}`

**路由**

然后到router.go中添加对应的路由

```go
beego.Router("/goodsDetail",&controllers.GoodsController{},"get:ShowDetail")
```

**控制器**

接着我们来实现ShowDetail方法。仍然是我们的四步骤。

+ 获取数据

  ```go
  id,err:= this.GetInt("id")
  ```

+ 校验数据

  ```go
  if err != nil{
  	beego.Error("获取数据不存在")
  	this.Redirect("/",302)
  	return
  }
  ```

+ 数据处理

  首先我们看一下，我们这个页面：

  ![1557761280595](.\assets\1557761280595.png)

  由页面可知，我们要获取类型数据和商品数据。代码如下：

  ```go
  //获取商品类型
  	var goodsTypes []models.GoodsType
  	o.QueryTable("GoodsType").All(&goodsTypes)
  	this.Data["types"] = goodsTypes
  //获取商品详情
  	var goods models.GoodsSKU
  	o.QueryTable("GoodsSKU").RelatedSel("GoodsType","Goods").Filter("Id",id).One(&goods)
  	this.Data["goods"] = goods
  ```

**视图**

我们把获取到的商品详细信息添加在视图中显示，代码如下：

```html
<div class="goods_detail_con clearfix">
	<div class="goods_detail_pic fl"><img src="http://192.168.110.81:8888/{{.goods.Image}}"></div>

	<div class="goods_detail_list fr">
		<h3>{{.goods.Name}}</h3>
		<p>{{.goods.Desc}}</p>
		<div class="prize_bar">
			<span class="show_pirze">¥<em>{{.goods.Price}}</em></span>
			<span class="show_unit">单  位：{{.goods.Unite}}</span>
		</div>
		<div class="goods_num clearfix">
			<div class="num_name fl">数 量：</div>
			<div class="num_add fl">
				<input type="text" class="num_show fl" value="1" name="goodsCount">
				<a href="javascript:;" class="add fr">+</a>
				<a href="javascript:;" class="minus fr">-</a>
			</div> 
		</div>
		<div class="total">总价：<em>{{.goods.Price}}</em></div>
		<div class="operate_btn">
			<a href="javascript:;" class="buy_btn">立即购买</a>
			<a href="javascript:;" skuid="{{.goods.Id}}" class="add_cart" id="add_cart">加入购物车</a>
		</div>
	</div>
</div>
```

**历史记录的添加和显示**

- 获取详情的时候我们需要添加历史浏览记录，我们介绍项目的时候就说过，我们的历史浏览记录存储在redis中。这里我们就需要做一个历史浏览记录的存储。在实现相关代码之前，我们先来分析一下，我们如何去存储 历史浏览记录。这里有以下几个问题。

  - 1.什么时候添加历史浏览记录

    在**登陆的情况下**查看商品详情的时候添加

  - 2.什么时候获取历史浏览记录

    在用户中心页获取历史浏览记录

  - 3.用什么格式来存储用户浏览记录

    这里我们要考虑一下历史浏览记录都存储哪些内容，历史浏览记录需要存储哪个用户浏览了哪些内容，而且还有先后顺序，所以要存储用户，商品内容，而且用户和商品是对应存储，并且有顺序，这里我们用redis数据库中的list来存储。存入的key值设定为`history_用户id`,value值为商品id,存储代码如下：

    ```go
    //添加历史浏览记录，需要先查询有没有登陆，只有登陆之后可以添加历史浏览记录
    	userName := this.GetSession("userName")
    	if userName != nil{
    		//查询用户信息
    		var user models.User
    		user.Name = userName.(string)
    		o.Read(&user,"Name")
    		conn,_:=redis.Dial("tcp",":6379")
    		//插入历史纪录
    		reply,err:=conn.Do("lpush","history"+strconv.Itoa(user.Id),id)
    		reply,_ = redis.Bool(reply,err)
    		if reply == false{
    			beego.Info("插入浏览数据错误")
    		}
    	}
    ```

    这时候需要注意，当用户多次浏览某个商品的时候，我们添加历史浏览记录只添加一条，所以在添加历史浏览记录之前需要把原来的记录给清空。修改之后的代码如下：

    ```go
    //添加历史浏览记录，需要先查询有没有登陆，只有登陆之后可以添加历史浏览记录
    	userName := this.GetSession("userName")
    	if userName != nil{
    		//查询用户信息
    		var user models.User
    		user.Name = userName.(string)
    		o.Read(&user,"Name")
    		conn,_:=redis.Dial("tcp",":6379")
    		//先清空以前的记录1
    		reply,err:=conn.Do("lrem","history"+strconv.Itoa(user.Id),0,id)
    		reply,_ = redis.Bool(reply,err)
    		if reply == false{
    			beego.Info("插入浏览数据错误")
    		}
    		//插入历史纪录
    		conn.Do("lpush","history"+strconv.Itoa(user.Id),id)
    	}
    ```

    添加完历史浏览记录，我们在用户中心页获取一下历史浏览记录。这里需要说明的是，我们只获取前五条历史浏览记录。代码如下：

    ```go
    //获取历史浏览记录
    	var goods []models.GoodsSKU
    	conn,_ :=redis.Dial("tcp",":6379")
    	reply,err := conn.Do("lrange","history"+strconv.Itoa(user.Id),0,4)
    	replyInts,_ := redis.Ints(reply,err)
    	for _,val := range replyInts{
    		var temp models.GoodsSKU
    		o.QueryTable("GoodsSKU").Filter("Id",val).One(&temp)
    		goods = append(goods, temp)
    	}
    	this.Data["goods"] = goods
    ```

    显示历史纪录的视图代码如下：

    ```html
    {{if .goods}}
        {{range .goods}}
             <li>
                 <a href="detail.html"><img src="http://192.168.110.81:8888/{{.Image}}"></a>
                 <h4><a href="detail.html">{{.Name}}</a></h4>
                 <div class="operate">
                 	<span class="prize">￥{{.Price}}</span>
                 	<span class="unit">{{.Price}}/{{.Unite}}</span>
                 	<a href="#" class="add_goods" title="加入购物车"></a>
                 </div>
             </li>
        {{end}}
    {{else}}
        <li>无历史浏览记录</li>
    {{end}}
    ```

### 3.列表页商品内容展示

**请求**

列表页展示的是相应类型的所有数据，所以请求应该是点击类型的时候跳转的页面，需要传递类型ID，所以我们要在类型超链接上添加请求路径为`/goodsList?typeId=1`

**路由**

有了请求之后需要我们在路由文件中添加相对应的控制器和方法，代码如下：

```go
//列表页展示
beego.Router("/goodsList",&controllers.GoodsController{},"get:ShowGoodsList")
```

**控制器**

接着我们来实现ShowGoodsList方法.

+ 获取数据

  首先我们要获取到路由传递过来的类型ID

  ```go
  //获取类型id
  typeId,err := this.GetInt("typeId")
  ```

+ 校验数据

  对获取的数据进行校验

  ```go
  //校验数据
  if err != nil{
  	beego.Info("获取类型ID错误")
  	this.Redirect("/",302)
  	return 
  }
  ```

+ 处理数据

  我们 这个页面主要是去查询相关的内容

  + 类型数据

    获取类型数据业务就比较简单，直接把所有类型数据从数据库中获取到。代码如下：

    ```go
    var types []models.GoodsType
    o.QueryTable("GoodsType").All(&types)
    this.Data["types"] = types
    ```

  + 获取当前类型的商品数据

    ```go
    //获取当前类型的商品
    var goodsSKus []models.GoodsSKU	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).All(&goodsSKus)
    this.Data["goods"] = goodsSKus
    ```

  + 获取新品数据

    我们这里获取的是同类型，时间靠前的两个商品数据，代码如下：

    ```go
    //获取两个新品
    var goodsNew []models.GoodsSKU
    o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).OrderBy("Time").Limit(2,0).All(&goodsNew)
    this.Data["goodsNew"] = goodsNew
    ```

**视图**

后台获取的数据，要在视图中循环显示：

```html
<div class="navbar_con">
		<div class="navbar clearfix">
			<div class="subnav_con fl">
				<h1>全部商品分类</h1>	
				<span></span>			
				<ul class="subnav">
					{{range .types}}
						<li><a href="/goodsList?id={{.Id}}" class="{{.Logo}}">{{.Name}}</a></li>
					{{end}}
				</ul>
			</div>
			<ul class="navlist fl">
				<li><a href="">首页</a></li>
				<li class="interval">|</li>
				<li><a href="">手机生鲜</a></li>
				<li class="interval">|</li>
				<li><a href="">抽奖</a></li>
			</ul>
		</div>
	</div>

	<div class="breadcrumb">
		<a href="#">全部分类</a>
		<span>></span>
		<a href="#">新鲜水果</a>
	</div>

	<div class="main_wrap clearfix">
		<div class="l_wrap fl clearfix">
			<div class="new_goods">
				<h3>新品推荐</h3>
				<ul>
					{{range .goodsNew}}
						<li>
							<a href="/goodsDetail?id={{.Id}}"><img src="http://192.168.110.81:8888/{{.Image}}"></a>
							<h4><a href="/goodsDetail?id={{.Id}}">{{.Name}}</a></h4>
							<div class="prize">￥{{.Price}}</div>
						</li>
					{{end}}
				</ul>
			</div>
		</div>

		<div class="r_wrap fr clearfix">
			<div class="sort_bar">
					<a href="/goodsList?id={{.typeId}}" class="active">默认</a>
					<a href="#">价格</a>
					<a href="#">人气</a>
			</div>

			<ul class="goods_type_list clearfix">
				{{range .goods}}
					<li>
						<a href="/goodsDetail?id={{.Id}}"><img src="http://192.168.110.81:8888/{{.Image}}"></a>
						<h4><a href="/goodsDetail?id={{.Id}}">{{.Name}}</a></h4>
						<div class="operate">
							<span class="prize">￥{{.Price}}</span>
							<span class="unit">{{.Price}}/{{.Unite}}</span>
							<a href="#" class="add_goods" title="加入购物车"></a>
						</div>
					</li>
				{{end}}
			</ul>

			<div class="pagenation">
				<a href="#">上一页</a>
						<a href="#" class="active">1</a>
				<a href="#">下一页></a>
			</div>
		</div>
	</div>
```

**分页处理**

我们这里实现的分页是如下的效果：

![1539707064353](./assets/1539707064353.png)

所以这里我们还需要获取相应页码，而且这个页码数，无论页码怎么变化，都显示的五条，所以这里获取五条页码就成了实现分页重要的一个环节。

这个可以分四种情况分析。
第一种：总页码数不到五页

这种情况下，要显示所有页码

第二种：显示的是前三页的内容

这种显示的都是1 2 3 4 5

第三种：显示的是最后三页的内容

显示的都是后五页的页码

第四种：处于中间的页码

分析出来的代码如下:

```go
if pageCount < 5{
		pageIndexBuffer = make([]int,pageCount)
		for index,_ := range  pageIndexBuffer{
			pageIndexBuffer[index] = index + 1
		}
	}else if pageIndex < 3  {
		pageIndexBuffer = make([]int,5)
		for index,_ := range pageIndexBuffer{
			pageIndexBuffer[index] = index + 1
		}
	}else if pageIndex >= pageCount -3{
		pageIndexBuffer = make([]int,page)
		for index,_ := range  pageIndexBuffer{
			pageIndexBuffer[index] = pageCount - 5 + index
		}
	}else {
		pageIndexBuffer = make([]int, 5)
		for index,_ := range pageIndexBuffer{
			pageIndexBuffer[index] = pageIndex - 3 + index
		}
	}
```

由代码分析可知，这里我们需要获取总页码说pageCount,需要获取当前页码pageIndex。那我们需要先获取pageIndex和pageCount。

**获取总页码和当前页码**

获取总页码和当前页码的业务逻辑和以前我们获取总页码和当前页码的逻辑一样，代码如下:

```go
//分页处理
	//指定每页显示多少数据
	pageSize := 2
	pageIndex ,err:=this.GetInt("pageIndex")
	if err !=nil{
		pageIndex = 1
	}
	start := pageSize * (pageIndex - 1)
	//处理页码
	count,_ := o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).Count()

	//获取总页码
	pageCount := math.Ceil(float64(count)/ float64(pageSize))
```

**获取显示的页码**

这里我们可以封装一个获取相应页码的函数,需要把总页码，当前页码，和要显示多少个页码当参数传递，代码如下：

```go
//指定显示多少个页码
page := 5
//函数定义如下

//分页助手函数
func PageTool(pageCount int,page int,pageIndex int)map[string]interface{}{
	//获取应该显示的页码
	var pageIndexBuffer []int
	if pageCount < page{
		pageIndexBuffer = make([]int,pageCount)
		for index,_ := range  pageIndexBuffer{
			pageIndexBuffer[index] = index + 1
		}
	}else if pageIndex < ( page + 1)/2  {
		pageIndexBuffer = make([]int,page)
		for index,_ := range pageIndexBuffer{
			pageIndexBuffer[index] = index + 1
		}
	}else if pageIndex > pageCount -( page + 1)/2{
		pageIndexBuffer = make([]int,page)
		for index,_ := range  pageIndexBuffer{
			pageIndexBuffer[index] = pageCount - page + index
		}
	}else {
		pageIndexBuffer = make([]int, page)
		for index,_ := range pageIndexBuffer{
			pageIndexBuffer[index] = pageIndex - (page - 1)/2 + index
		}
	}

	//上一页页码
	pagePre := pageIndex - 1
	if pageIndex == 1{
		pagePre = 1
	}
	pageNext := pageIndex + 1
	if pageIndex == pageCount{
		pageNext = pageIndex
	}
	//把数据返回
	pageData := make(map[string]interface{})
	pageData["pagePre"] = pagePre
	pageData["pageNext"] = pageNext
	pageData["pageIndex"] = pageIndexBuffer

	return pageData
}
```

> 注意：这里上一页和下一页的处理也放到这个函数中实现

获取到数据之后，我们要把相应的数据传递给视图：

```go
//把数据传递给视图
	this.Data["pagePre"] = pageData["pagePre"]
	this.Data["pageNext"] = pageData["pageNext"]
	this.Data["pages"] = pageData["pageIndex"].([]int)
	this.Data["pageIndex"] = pageIndex
```

在视图中显示:

```html
<a href="/goodsList?pageIndex={{.pagePre}}">上一页</a>
{{range $index,$val := .pages}}
	{{if compare $.pageIndex $val}}
		<a href="/goodsList?pageIndex={{$val}}" class="active">{{$val}}</a>
	{{else}}
		<a href="/goodsList?pageIndex={{$val}}">{{$val}}</a>
	{{end}}
{{end}}
<a href="/goodsList?pageIndex={{.pageNext}}">下一页></a>
```

> 感觉好像成功了，有没有问题呢？我们接着往下看

**获取分页数据**

这时候我们可以去获取分页的数据了，获取分页的数据依旧用高级查询中的Limit函数，代码如下：

```go
o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).Limit(pageSize,start).All(&goodsSKus)
```

> 注意：这时候你会遇见一个问题，当直接访问列表页的时候，数据正确，但是点击上一页，下一页的时候，你会发现，获取数据不正确。什么原因呢？其实根据错误提示也能发现，我们点击上一页下一页以及相应页码的时候，并没有传递typeId值过来，所以对于后台来说，在新的请求中并不知道这次你要获取哪个类型的数据。这时候我们需要在上一次的访问中把typeId传递给视图。然后在分页的<a>标签传递回来。具体代码实现如下：

**先把typeId传递回视图**

```go
this.Data["typeId"] = typeId
```

**在视图中拼接<a>标签，传递回来typeId内容**

```html
<a href="/goodsList?typeId={{.typeId}}&pageIndex={{.pagePre}}">上一页</a>
{{range $index,$val := .pages}}
	{{if compare $.pageIndex $val}}
		<a href="/goodsList?typeId={{$.typeId}}&pageIndex={{$val}}" class="active">{{$val}}</a>
	{{else}}
		<a href="/goodsList?typeId={{$.typeId}}&pageIndex={{$val}}">{{$val}}</a>
	{{end}}
{{end}}
<a href="/goodsList?typeId={{.typeId}}&pageIndex={{.pageNext}}">下一页</a>
```

这时候，分页的功能就实现了 。

**根据不同的选项获取不同排序内容**

我们这个页面还要根据不同的选择，显示不同的商品排序，分别是默认排序，价格排序，人气（销量）排序。那就需要在我们查询的时候再加上排序的函数，这里我们用的排序函数是**OrderBy()**，也是高级查询函数，参数是要排序的字段名。那这时候，我们给这三个超链接就需要添加一个sort的值。代码如下：

```html
<a href="/goodsList?typeId={{.typeId}}" class="active">默认</a>
<a href="/goodsList?typeId={{.typeId}}&sort=price">价格</a>
<a href="/goodsList?typeId={{.typeId}}&sort=sale">人气</a>
```

在后台，我们需要根据不同的sort值获取不同顺序的数据，代码如下：

```go
//根据不同的选项获取不同的数据
sort := this.GetString("sort")
//如果sort等于空，就按照默认排序
if sort == ""{	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).Limit(pageSize,start).All(&goodsSKus)
}else if sort == "price"{	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).OrderBy("Price").Limit(pageSize,start).All(&goodsSKus)
}else{
o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).OrderBy("Sales").Limit(pageSize,start).All(&goodsSKus)
}
```

这时候我们看页面，貌似实现了相应的功能。**但是，**仔细看会发现三个超链接的选中状态会有些问题。所以，在我们显示相应页面数据的时候，我们也应该把排序数据传递过来，做一个比较，确定哪个是选中的排序类型。具体后台代码如下：

```go
//根据不同的选项获取不同的数据
sort := this.GetString("sort")
//如果sort等于空，就按照默认排序
if sort == ""{	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).Limit(pageSize,start).All(&goodsSKus)
}else if sort == "price"{	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).OrderBy("Price").Limit(pageSize,start).All(&goodsSKus)
}else{	o.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id",typeId).OrderBy("Sales").Limit(pageSize,start).All(&goodsSKus)
}
this.Data["sort"] = sort
```

视图代码如下：

```html
<a href="/goodsList?id={{.typeId}}" {{if compare .sort ""}}class="active"{{end}}>默认</a>
<a href="/goodsList?id={{.typeId}}&sort=price" {{if compare .sort "price"}}class="active"{{end}}>价格</a>
<a href="/goodsList?id={{.typeId}}&sort=sale" {{if compare .sort "sale"}}class="active"{{end}}>人气</a>
```

这时候，仔细观察我们会发现，我们的分页也应该传递排序数据，不然获取分页数据的时候就会出现错误的显示，所以分页这块的视图显示如下：

```html
<div class="pagenation">
	<a href="/goodsList?id={{.typeId}}&pageIndex={{.pagePre}}&sort={{.sort}}">上一页</a>
	{{range $index,$val := .pages}}
		{{if compare $.pageIndex $val}}
			<a href="/goodsList?id={{$.typeId}}&pageIndex={{$val}}&sort={{$.sort}}" class="active">{{$val}}</a>
		{{else}}
			<a href="/goodsList?id={{$.typeId}}&pageIndex={{$val}}&sort={{$.sort}}">{{$val}}</a>
		{{end}}
	{{end}}
	<a href="/goodsList?id={{.typeId}}&pageIndex={{.pageNext}}&sort={{.sort}}">下一页></a>
</div>
```

到这里，我们的商品列表就彻底实现了。

### 4.搜索页面商品内容展示（老版）

接着我们来实现商品模块的最后一个内容，商品搜索。

在实现相关内容之前需要给大家先介绍一个内容，**ORM过滤器的高级用法**

我们在前面用过滤器Filter的时候，只能过滤相等的情况，但是在实际业务开发中，我们不止会遇到相等的情况，还有大于，小于，大于等于，小于等于，包含，以某个字符开始，以某个字符结束等相关查询操作，那这些功能在我们过滤器中是怎么实现呢？我们通过下面的例子来说明，举例如下：

```go
qs.Filter("profile__age__gt", 18) // WHERE profile.age > 18   查询profile表中age属性大于18的值
```

通过上面的例子，我们能看到如果要使用过滤器的高级用法，需要在第一个参数后面再追加上`__`然后跟上相应的操作符号，来表示不同的需求。那这些单词都有哪些呢？我们通过下面的表来做一个简单的了解:

|         操作符         |                             作用                             |
| :--------------------: | :----------------------------------------------------------: |
|      exact/iexact      |             判断指定的字段是否等于第二个参数的值             |
|   contains/icontains   |             判断指定的字段是否包含第二个参数的值             |
|        gt / gte        |        判断指定的字段是否大于/大于等于第二个参数的值         |
|        lt / lte        |        判断指定的字段是否小于/小于等于第二个参数的值         |
| startswith/istartswith |          判断指定的字段是否是以第二个参数的值为开头          |
|   endswith/iendswith   |          判断指定的字段是否是以第二个参数的值为结尾          |
|         isnull         |                   判断指定的字段是否为null                   |
|           in           | 判断指定的字段是否在第二个参数内部（这时候第二个参数一般为切片，也可以多放几个参数） |

> 注意：这里成对出现，并且以**i**开头的表示：大小写不敏感 

了解了上面的知识点之后我们就可以用contains来查询我们的商品数据。

**请求**

搜索的请求需要把搜索框的数据发送给后台，所以我们需要先在两个input标签外边加上form标签，然后给form表单中给搜索添加action。代码如下:

```html
<form method="post" action="/searchGoods">
    <input type="text" class="input_text fl" name="goodsName" placeholder="搜索商品">
    <input type="submit" class="input_btn fr" name="" value="搜索">
</form>
```

**路由**

有了请求，接着去router.go文件中给请求添加对应的控制器和方法。代码如下：

```go
//搜索功能
beego.Router("/searchGoods",&controllers.GoodsController{},"post:HandleSearch")
```

**控制器**

接着我们去实现HandleSearch方法，一就是我们的四步骤：

+ 获取数据

  ```go
  goodsName := this.GetString("goodsName")
  ```

+ 校验数据

  这里需要注意的是，如果我们搜索框传递过来的数据为空，这时候我们应该怎么处理？暂时我们处理的是如果搜索框传递数据为空，那么我们就获取全部商品数据。

  ```go
  if goodsName == ""{
  	beego.Info("查找的数据为空")
  	o.QueryTable("GoodsSKU").All(&goods)
  	this.Data["goods"] = goods
  	ShowLayout(&this.Controller)
  	this.TplName = "search.html"
  }
  ```

+ 处理数据

  这时候我们拿到了搜索关键字，那我们就根据这个关键字进行搜索。代码如下：

  ```go
  //根据拿到的数据进数据库查询
  o := orm.NewOrm()
  var goods []models.GoodsSKU
  o.QueryTable("GoodsSKU").Filter("Name__icontains",goodsName).All(&goods)
  ```

+ 返回视图

  这时候需要注意的时候，我们给的页面里面没有搜索结果页面，所以需要我们自己创建一个搜索页面。这里我们修改一下商品列表页，充当我们的搜索结果页。修改之后的代码如下：

  ```html
  <div class="breadcrumb">
  	<a href="#">全部分类</a>
  	<span>></span>
  	<a href="#">搜索界面</a>
  </div>
  
  <div class="main_wrap clearfix">
  
  	<div class="r_wrap fr clearfix">
  		<ul class="goods_type_list clearfix">
  			{{range .goods}}
  				<li>
  					<a href="/goodsDetail?id={{.Id}}"><img src="http://192.168.110.66:8888/{{.Image}}"></a>
  					<h4><a href="/goodsDetail?id={{.Id}}">{{.Name}}</a></h4>
  					<div class="operate">
  						<span class="prize">￥{{.Price}}</span>
  						<span class="unit">{{.Price}}/{{.Unite}}</span>
  						<a href="#" class="add_goods" title="加入购物车"></a>
  					</div>
  				</li>
  			{{end}}
  		</ul>
  
  	</div>
  </div>
  ```

  后台代码处理：

  ```go
  //指定视图
  ShowLayout(&this.Controller)
  this.TplName = "serach.html"
  ```

  
  
### 新版搜索

 上面的搜索底层用的是关系数据库中like关键字实现的，但是like关键字的效率极低，而且查询需要在多个字段中进行，使用like关键字也不方便。

所以我们引入搜索引擎来实现全文检索。

全文检索即在指定的任意字段中进行检索查询。

#### 1.搜索引擎原理

通过搜索引擎进行数据查询时，搜索引擎并不是直接在数据库中进行查询，而是搜索引擎会对数据库中的数据进行一遍预处理，单独建立起一份索引结构数据。

我们可以将索引结构数据想象成是字典书籍的索引检索页，里面包含了关键词与词条的对应关系，并记录词条的位置。

  ![1563355469656](./assets/1563355469656.png)

我们在通过搜索引擎搜索时，搜索引擎将关键字在索引数据中进行快速对比查找，进而找到数据的真实存储位置。

# 2.Elasticsearch介绍

开源的 [Elasticsearch ](https://www.elastic.co/)是目前全文搜索引擎的首选。

它可以快速地储存、搜索和分析海量数据。维基百科、Stack Overflow、Github 都采用它。

Elasticsearch 的底层是开源库 [Lucene](https://lucene.apache.org/)。但是，你没法直接用 Lucene，必须自己写代码去调用它的接口。Elastic 是 Lucene 的封装，提供了 REST API 的操作接口，开箱即用。

Elasticsearch 是用Java实现的。

搜索引擎在对数据构建索引时，需要进行分词处理。分词是指将一句话拆解成多个单字或词，这些字或词便是这句话的关键词。如

```python
我是中国人。
```

'我'、'是'、'中'、'国'、'人'、'中国'等都可以是这句话的关键词。

Elasticsearch 不支持对中文进行分词建立索引，需要配合扩展**elasticsearch-analysis-ik**来实现中文分词处理。

#### 3.Elasticsearch及其扩展安装

+ 安装java环境

  > 在线安装,需要下载二百多兆安装数据，进公司可以这么安装

  ```bash
  sudo apt install default-jdk
  ```

  > 离线安装
  >
  > 1.下载jdk（java安装包,已下载）
  >
  > 下载连接：http://www.oracle.com/technetwork/java/javase/downloads/index.html
  >
  > 2.解压并移动到指定目录
  >
  > tar -zxvf jdk-12.0.2_linux-x64_bin.tar.gz 
  >
  > sudo mv  jdk-12.0.2  /usr/local/java
  >
  > 3.进行配置
  >
  > 打开家目录下面的.bashrc文件,在文件末尾加上如下配置。
  >
  > ```bash
  > export JAVA_HOME=/usr/local/java
  > export CLASSPATH=JAVA_HOME/lib:{JRE_HOME}/lib  
  > 
  > export PATH=JAVA_HOME/bin:PATH
  > ```
  >
  > 4.让配置生效
  >
  >  source .bashrc
  >
  > 5.测试是否安装成功
  >
  > java -version

+ 安装elasticsearch

  > 获取elasticsearch

  wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.2.0-amd64.deb

  > 解压文件

  sudo dpkg -i elasticsearch-7.2.0-amd64.deb

  > 重新生成ubuntu之间的依赖关系

  sudo systemctl daemon-reload

  > 设置开机自启动

  sudo systemctl enable elasticsearch.service

  > 开启elasticsearch服务

  sudo systemctl start elasticsearch



+ 编辑配置项

  sudo vim /etc/elasticsearch/elasticsearch.yml 

  去掉下面三行的注释

  ```shell
  bootstrap.memory_lock: true  
  network.host: 192.168.0.1  
  http.port: 9200 
  ```

  然后把 network.host对应的改成 localhost

+ 再次重启

  ```shell
  sudo systemctl daemon-reload  
  sudo systemctl enable elasticsearch.service  
  sudo systemctl start elasticsearch 
  ```

+ 查看服务是否开启成功

  netstat -plntu 

  ![1563454315174](./assets/1563454315174.png)

+ 重新启动并测试

  ```shell
  sudo systemctl restart elasticsearch  
  curl 'localhost:9200'   
  ```

  ![1563454421483](./assets/1563454421483.png)

+ 安装ik分词器

  不同版本的elasticsearch对应不同版本的ik分词器，对应如下：

  ![1563454741492](./assets/1563454741492.png)

  安装

  ```shell
  sudo /usr/share/elasticsearch/bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.2.0/elasticsearch-analysis-ik-7.2.0.zip 
  ```

s