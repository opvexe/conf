#  Navicat Premium 连接远端数据库 

```shell
修改:
	sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf 
修改文件:
 	bind-address = 127.0.0.1 改成  bind-address = 0.0.0.0
 退出，保存，重启MySQL服务器
```

```shell
**解决方法:**

	1.使用Linux下文字界面登录数据库： mysql -uroot -p123456
	
	2.选择使用 mysql 数据库：mysql> use mysql
	
	3.执行 mysql> update user set host = '%' where user = 'root';
	
    4.查看 select host,user from user;
    
    5.刷新MySQL的权限相关表：flush privileges;
    
    6.无需重启MySQL数据库，直接重连 Navicat 即可
```

