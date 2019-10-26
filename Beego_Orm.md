# ORM

## 一对多查询

- RelatedSel() 
  - 有参数:RelatedSel("ArticleType")  左内联
  - 无参数：RelatedSel() 左关联

```go
//1.获取orm对象
o :=orm.NewOrm()
//2. select *from Article inner join ArticleType on Article.id=ArticleType.id;
var article []Article
o.QueryTable("Article").RelatedSel("ArticleType").Limit(size,start).All(&article)
```



## 多对多插入

- 多对多插入五步骤（固定写法）

```go
//1.获取orm对象
o :=orm.NewOrm()
//2.获取被插入的对象
var article Article
article.Id = id
//3.读取
o.Read(&article,"Id")
//4.获取多对多插入对象
// 在mysql中会生成一个article_users第三张表
m2m:=o.QueryM2M(article,"Users") //article的多对多字段是Users
//5.获取插入对象
var usr User
usr.Name= UserName
o.Read(&usr,"Name")
//6.多对多插入
m2m.Add(usr) 
```



## 多对多查询

- LoadRelated :加载关系表

```go
//1.获取orm对象
o :=orm.NewOrm()
//2.获取被插入的对象
var article Article
article.Id = id
//3.读取
o.Read(&article,"Id")
//4.多对多关联 
o.LoadRelated(&article,"Users")
//5.多对多User表去重
//mysql去重distinct
//.Filter("Articles__Article__Id",id) 多对多查询 Articles__ 1对多查询不需要
// select *from user where Article.id = id 
o.QueryTable("User").Filter("Articles__Article__Id",id).Distinct().All(&usrs)
```

