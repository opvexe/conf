## 1. Go语言的接口

### 1.1什么是接口

* 在Go中，接口是一组方法签名。当类型为接口中的所有方法提供定义时，它被称为实现接口。它与OOP非常相似。接口指定了类型应该具有的方法，类型决定了如何实现这些方法。

>它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
>接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了该接口。
>接口是一个或多个方法签名的集合
>接口只有方法声明没有实现，没有数据字段
>接口可以匿名嵌入其他接口或者嵌入结构中
>将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针
>只有当接口存储的类型和对象都为nil是，接口才等于nil
>接口同样支持匿名字段方法
>接口也可是实现类似oop中的多态
>空接口可以作为任何类型数据的容器

### 1.2接口的定义语法

```go
type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    var phone Phone
    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()
}

//输出:
I am Nokia, I can call you!
I am iPhone, I can call you!
```



### 1.3接口嵌套

```go
type USB interface{
  Name() string
  //接口嵌套
  Connecter
}

type Connecter interface{
   Connect()
}

type PhoneConnecter struct{
    name string
}
func (pc PhoneConnecter)Name() string{
    return pc.name
}

func (pc PhoneConnecter)Connect(){
    fmt.Printf("%s is connect \n",pc.name)
}

func DisConnect(pc USB){
    fmt.Printf("phone is disconnect \n")
}

func main(){
   var a USB = PhoneConnecter{name:"iphone"}
   a.Connect()
   DisConnect(a)
}

//输出:
iphone is connect
phone is disconnect

```



### 1.4类型断言

```go
type USB interface{
  Name() string
  //接口嵌套
  Connecter
}

type Connecter interface{
   Connect()
}

type PhoneConnecter struct{
    name string
}
func (pc PhoneConnecter)Name() string{
    return pc.name
}

func (pc PhoneConnecter)Connect(){
    fmt.Printf("%s is connect \n",pc.name)
}

func DisConnect(pc USB){
//这里对传进来的类型进行判断 采用多返回值，当ok参数为true时才执行
  if v,ok := pc.(PhoneConnecter);ok{
    fmt.Printf("phone is disconnect \n")
  } 
}

func main(){
   var a USB = PhoneConnecter{name:"iphone"}
   a.Connect()
   DisConnect(a)
}

```



### 1.5空接口可以作为任何接口的基类

```go
//空接口
type empty interface{

}

type USB interface{
  Name() string
  //接口嵌套
  Connecter
}

type Connecter interface{
   Connect()
}

type PhoneConnecter struct{
    name string
}
func (pc PhoneConnecter)Name() string{
    return pc.name
}

func (pc PhoneConnecter)Connect(){
    fmt.Printf("%s is connect \n",pc.name)
}
//将接受类型写成任意接口类型
func DisConnect(pc interface{}){
   switch pc.(type){
   case PhoneConnecter:
       fmt.Printf("phone name is %s", pc.(PhoneConnecter).name)
   default:
       fmt.Println("unknow phone")    
   }
}

func main(){
   var a USB = PhoneConnecter{name:"iphone"}
   a.Connect()
   DisConnect(a)
}
```



### 1.6 类型转换

```go
type USB interface{
  Name() string
  //接口嵌套
  Connecter
}

type Connecter interface{
   Connect()
}

type PhoneConnecter struct{
    name string
}
func (pc PhoneConnecter)Name() string{
    return pc.name
}

func (pc PhoneConnecter)Connect(){
    fmt.Printf("%s is connect \n",pc.name)
}
//将接受类型写成任意接口类型
func DisConnect(pc interface{}){
   switch pc.(type){
   case PhoneConnecter:
       fmt.Printf("phone name is %s", pc.(PhoneConnecter).name)
   default:
       fmt.Println("unknow phone")    
   }
}

func main(){
   var pcConnect = PhoneConnecter{"iphone"}
   //类型转换
   var a = Connecter(pcConnect)
   a.Connect()
   DisConnect(a)
}
```



## 2.Go语言多态

### 2.1 开闭原则

```go
//使用开闭原则
type IBanker interface {
	Bussiness()
}

//存款业务,存款业务可以写到一个独立的文件中
type SaveMoneny struct {

}

func (sm *SaveMoneny)Bussiness()  {
	fmt.Println("banker  save money")
}

//取款业务
type GetMoneny struct {

}

func (sm *GetMoneny)Bussiness()  {
	fmt.Println("banker get money")
}


//同一处理业务的函数

func DoBussiss(banker IBanker)  {
	banker.Bussiness()
}


func main()  {

	//存款业务
	DoBussiss(&SaveMoneny{})

	DoBussiss(&GetMoneny{})
}

```



### 2.2依赖倒置

```go 
//中间层面 --- 抽象层
// 汽车
type Car interface {
	Run()
}

//驾驶员
type Driver interface {
	Drive()
}

//实现层

type Benz1 struct {
}

func (benz *Benz1) Run() {
	fmt.Println("benz is runing ..")
}

type BWM1 struct {
}

func (bwm *BWM1) Run() {
	fmt.Println("bwm is runing ..")
}

//具体驾驶员
type Lily struct {

}

func (lily *Lily)Drive(car Car) {
	car.Run()
}

type Tom struct {
}

func (tom *Tom) Drive(car Car) {
	car.Run()
}

//业务层
func main() {

	benz := Benz1{}
	bwm := BWM1{}

	//准备驾驶员
	tom := Tom{}
	lily := Lily{}

	//具体驾驶动作
	tom.Drive(&benz)
	tom.Drive(&bwm)

	lily.Drive(&benz)
	lily.Drive(&bwm)
}

```



### 2.3练习

```go
//定义显卡接口
type ICard interface {
	Card()
}

//定义内存
type IMeomery interface {
	Meomery()
}

//定义cpu
type ICPU interface {
	CPU()
}

////定义电脑类
type Computer struct {
	card ICard
	mem  IMeomery
	cpu  ICPU
}

//要传入实现接口的对象 ----
func NewComputer(card ICard, men IMeomery, cpu ICPU) *Computer {
	return &Computer{card, men, cpu}
}

//运行电脑
func (cmp *Computer) ComputerWork() {
	//不同组件运行
	cmp.card.Card()
	cmp.mem.Meomery()
	cmp.cpu.CPU()
}

type InterCard struct {
}

func (cmp *InterCard) Card() {
	fmt.Println("Card ...")
}

type InterMeomery struct {
}

func (cmp *InterMeomery) Meomery() {
	fmt.Println("Meomery ...")
}

type InterCPU struct {
}

func (cmp *InterCPU) CPU() {
	fmt.Println("ICPU ...")
}

//组装
type KinstoneMenory struct {
}

func (cmp *KinstoneMenory) Meomery() {
	fmt.Println("KinstoneMenory Meomery ...")
}

type Navia struct {
}

func (cmp *Navia) CPU() {
	fmt.Println("Navia CPU ...")
}

func main() {

	//组装一台interner系列的电脑，并运行
	cp1 := NewComputer(&InterCard{}, &InterMeomery{}, &InterCPU{})
	cp1.ComputerWork()

	//组装一台intercpu,内存，显卡
	cp2 := NewComputer(&InterCard{}, &KinstoneMenory{}, &Navia{})
	cp2.ComputerWork()
}

```

