# Git基本操作

## Git安装
1.检查是否安装
> git version

2.安装git
>sudo apt-get install git

3.配置环境
> git config --global user.name "Your Name"<br>
> git config --global user.email "email@example.com"<br>

** 注意git config命令的--global参数，用了这个参数，表示你这台机器上所有的Git仓库都会使用这个配置，当然也可以对某个仓库指定不同的用户名和Email地址。** <br> 

## Git 操作

#### 拉取远程仓库
> git clone https://github.com/bugfan/authx

#### 将本地代码关联到远端
1.项目目录下初始化本地仓库</br>
>git init</br>

2.推到远程仓库
>git remote add origin https://github.com/bugfan/authx

3.提交本地文件或代码  ** . 代表所有 **
>git add . 

4.将本地代码提交到远端
>git commit -m '内容'

5.推送到远端
>git push -u origin master


## 代码回滚



+ **情况1: 代码commit但未push，想回到commit状态（保留本次commit代码块）**

```
$ git log 
```


![Git](./assets/20191011825.png)

```
$ git reset bfd7b2698079fcdacb9df648345f07a9937ccfb8  
```

---


+ **情况2: 代码commit但未push，想回到commit状态（不保留本次commit代码块）**


```
$ git log 
```

![Git](./assets/20191011825.png)

```
$ git reset -hard bfd7b2698079fcdacb9df648345f07a9937ccfb8
```


## 工作区和暂存区











## 暂存区
***  将本地代码提交到暂存区 **** 

