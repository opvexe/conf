# Makefile

==`${ }` 与`$( )`区别==

```shell
# $( )与` `（反引号）都是用来作命令替换的。

# ${ }  变量替换 就是把变量的真实值带入
```



- 描述：
  - ==很多时候，我们需要运行多个命令完成一件事==

- 解决方案：
   - 写shell 脚本
     	- 创建 build.sh
     	- 比较复杂，不通用
   - makefile
     	- 使用简单
     	- 可执行多条命令
- 语法

```makefile
target ... : prerequisites ...
    command
    ...
    ...
```

> PHONY (伪目标)

```makefile
# 不关心有没有旧文件，都会重新编译
.PHONY: clean

clean:
      rm *.db
```

>  target（目标）

```makefile
# 一般是执行文件名 比如: make  clean (执行文件名)
clean:
      rm *.db
```

>  prerequisites(前置条件)

```makefile
.PHONY: clean

# 依赖执行 ---> 先执行clean  ---> 然后执行 
install: clean
	go run *.go

clean:
      rm *.db
```

> commands(命令)

```makefile
# 由一行或多行的Shell命令组成
install: clean
 `Tab`
	go run *.go
	
#*************   处理多行换行问题 *********************#

# 办法1 ：是在换行符前加反斜杠转义。
clean:
   export foo=bar; \
   echo "foo=[$$foo]"
   
# 方法2：.ONESHELL:命令。
.ONESHELL:
clean:
    export foo=bar; 
    echo "foo=[$$foo]"
```

- demon

```makefile
# PHONY： 忽略 (不管是否已经执行生成blockChain，clean) -->依然会执行 make blockChain命令
.PHONY:blockChain clean

# 起别名
app:blockChain clean # app 依赖后面，后面先执行 blockChain，clean

# 编译命令
blockChain: clean  # 后面可以依赖执行 clean
	go build -o blockChain *.go
	# 执行顺序 先执行clean 在执行 go build -o blockChain *.go 

clean:
	rm -rf *.db blockChain

test:
	go test -v *_test.go

```

- Bee

```makefile
VERSION = $(shell grep 'const version' cmd/commands/version/version.go | sed -E 's/.*"(.+)"$$/v\1/')

.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	go build $(GOFLAGS) ./...

install:
	go get $(GOFLAGS) ./...

test: install
	go test $(GOFLAGS) ./...

bench: install
	go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	go clean $(GOFLAGS) -i ./...

publish:
	mkdir -p bin/$(VERSION)
	cd bin/$(VERSION)
	xgo -v -x --targets="windows/*,darwin/*,linux/386,linux/amd64,linux/arm-5,linux/arm64" -out bee_$(VERSION) github.com/beego/bee
	cd ..
	ghr -u beego -r bee $(VERSION) $(VERSION)
```

