# Mac搭建nginx

[**本文链接**](https://segmentfault.com/a/1190000016020328)

#### 步骤一：先更新homebrew
```
brew update
```

#### 步骤二： 查看nginx信息
```
brew search nginx
```

#### 步骤三：安装nginx
```
brew install nginx
```

#####  **** 对应的配置文件地址在  **/usr/local/etc/nginx/nginx.conf **  

#### 步骤四：运行nginx
```
nginx 
```

nginx默认使用8080端口 如果发现端口被占用了，可以杀掉使用使用改端口的进程，也可以修改/usr/local/etc/nginx/nginx.conf 下的

#### 重新启动nginx
```
nginx -s reload
```