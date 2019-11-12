# 默认版

```go
package main

import (
	"flag"
	"fmt"
)

//person:
//--name
//--age

//book:
//--pages
//--new

func main() {
	name := flag.String("name", "Lily", "这是人的名字")
	age := flag.Int("age", 18, "这个是人的年龄")

	pages := flag.Int("pages", 500, "这是书的页数")
	new1 := flag.Bool("new1", false, "这个书的新旧状态")

	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("pages:", *pages)
	fmt.Println("new:", *new1)
	fmt.Println("++++++++++")

	//解析命令
	flag.Parse()
	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("pages:", *pages)
	fmt.Printf("new:%v\n", *new1)

}

```







# 分组版

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("./a.exe person --name <Name> --age <AGE>\n")
	fmt.Printf("./a.exe book --pages <PAGES> --new <NEW>\n")
}

func main() {
	if len(os.Args) < 3 {
		Usage()
		return
	}

	//创建命令集
	personCmd := flag.NewFlagSet("person", flag.ExitOnError)
	bookCmd := flag.NewFlagSet("book", flag.ExitOnError)

	//对相应的命令集添加子命令
	name := personCmd.String("name", "Lily", "这是人的名字")
	age := personCmd.Int("age", 18, "这个是人的年龄")

	pages := bookCmd.Int("pages", 500, "这是书的页数")
	new1 := bookCmd.Bool("new1", false, "这个书的新旧状态")

	switch os.Args[1] {
	case "person":
		personCmd.Parse(os.Args[2:])
		fmt.Printf("name:%v\n", *name)
		fmt.Printf("age:%v\n", *age)
	case "book":
		bookCmd.Parse(os.Args[2:])
		fmt.Printf("pages:%v\n", *pages)
		fmt.Printf("new:%v\n", *new1)
	default:
		Usage()
	}
}

```

a