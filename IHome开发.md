# IHome开发

## 1.项目目录结构

首先是gin的下载，下载命令如下：

```shell
$ go get -u -v github.com/gin-gonic/gin
```

#### 1.1 创建工具类文件夹

```json
package utils

const (
	RECODE_OK         = "0"
	RECODE_DBERR      = "4001"
	RECODE_NODATA     = "4002"
	RECODE_DATAEXIST  = "4003"
	RECODE_DATAERR    = "4004"
	
	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERONERR  = "4104"
	RECODE_ROLEERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_USERERR    = "4107"
	RECODE_SMSERR     = "4108"
	RECODE_MOBILEERR  = "4109"

	RECODE_REQERR     = "4201"
	RECODE_IPERR      = "4202"
	RECODE_THIRDERR   = "4301"
	RECODE_IOERR      = "4302"
	RECODE_SERVERERR  = "4500"
	RECODE_UNKNOWERR  = "4501"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_USERONERR:  "用户已经注册",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
	RECODE_SMSERR:     "短信失败",
	RECODE_MOBILEERR:  "手机号错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
```

#### 1.2 创建数据库和表

- 打开数据库

```go
//将数据库时间转换  需要设置parseTime=true
//例如：2019:11:10 08:25:55 时间格式转换： +0000 UTC 2019-11-10 08:25:55 +0000 UTC <nil>} bj5q 123456}
db,err:=gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/ihome?parseTime=true")
	if err!=nil {
		fmt.Println("gorm Open:",err)
		return
	}
```

- ==gorm默认设置连接池，设置连接池属性==

```go
	//２.gorm设置链接池
	db.DB().SetMaxIdleConns(20) //设置最大空闲数量,设置初始化数量
	db.DB().SetMaxOpenConns(30) //最大数量
	db.DB().SetConnMaxLifetime(60*30) //设置最大生命周期
```

- ==建表，gorm默认表数据是复数 例如stu -->stus==

```go
db.SingularTable(true) //设置为单数
db.AutoMigrate(
	new(Stu),
)
//注意 gorm默认创建的数据为空，orm默认非空
```

- 数据插入

```go
var stu  Stu
	stu.Name = "bj5q"
	stu.PassWord = "123456"

	if err:=GlobalDB.Create(&stu).Error;err!=nil{
		fmt.Println("创建数据失败")
		return
	}
	fmt.Println(stu)
```

- 数据库查询

```go
	var stu Stu
	//select *from user where name =bj5q and password =123456
	if err:=GlobalDB.First(&stu).Where("name=?","bj5q").Where("pass_word=?","123456").Error;err!=nil{
		fmt.Println("数据库查询失败")
		return
	}
	fmt.Println(stu)


//First  查询到的第一条数据
//last	查询到的最后一条数据
//Find	查询到所有数据
```

 ==默认情况下已主键查询 stu.id = 1,如果主键没有数据，查询所有并获取第一条==

 根据name查询时，以where查询

- 更新数据库 (按照条件更新)

```go
	var stu Stu
	stu.ID = 1
	stu.Name = "itcast"
	stu.PassWord = "111"

	if err:=GlobalDB.Save(&stu).Error;err!=nil{
		fmt.Println("更新数据库失败")
		return
	}

	fmt.Println(stu)

//save 一般不常用，如果没有的话，会执行插入操作
```

注意: 是以主键为更新条件 ，如果主键没有赋值，则是插入操作。

- 更新一个字段

```shell
	if err:=GlobalDB.Model(&stu).Where("name=?","itcast").Update("pass_word","0000").Error;err!=nil{
		fmt.Println("更新数据库失败")
		return
	}
```

- 更新多个字段

```go
	if err:=GlobalDB.Model(&stu).Where("id=1").Updates(map[string]interface{}{"name":"itheim","pass_word":"12122"}).Error;err!=nil{
		fmt.Println("更新数据库失败")
		return
	}
```

==update或者updates都可以更新多个字段==

- 数据库删除 （软删除）（数据库中含有deteTimeAt字段）

```go
if err:=GlobalDB.Where("id=1").Delete(&stu).Error;err!=nil{
		fmt.Println("删除数据库失败")
		return
	}
```

