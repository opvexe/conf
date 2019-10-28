# 设计模式

## 1. 单利模式

```go
// 创建全局对象
type singleton struct{
}

// 同步Once,保证每次调用时，只有第一次生效
var once sync.Once

// 定义一个包级别的Public实例变量
var SingleManager *singleton

//初始化单利对象
func NewSingleton() *singleton{
	once.Do(func() {
    SingleManager = new(singleton)
	})
  return SingleManager
}
```



## 2. 观察者模式

```go
/*
		定义观察者接口
 */
type Observer interface{
  Update(int)	//每个观察者要实现的变更自身的方法
}

// 被观察者接口 
type ObserverInterFace interface{
  Notify()									//当被观察者状态变更时，通知所有已注册的观察者，调用其自身的变更方法
  State() int								//查看被观察者此时的状态
  SetState(int)							//变更被观察者的状态
  AddObserve(ob Observer)		//注册观察者
  RemoveObserve(ob Observer)//删除观察者
}

/*
		实例化观察者
 */
type Object struct{
  	state     int  //观察者监听的状态
    observers []Observer  //存储已注册的观察者对象的容器
}

// 初始化
func InitWithObject(state int) *Object {
  return &Object{state:state,observers:make([]Observer,0)}
}


// 实现查看当前被观察者对象的状态
func (o *Object)State() int{
	return o.state
}

// 实现更改被观察者对象的状态
func (o *Object)SetState(state int){
  	o.state = state
}

 // 注册观察者
func (o *Object)AddObserve(ob Object){
  o.observers = append(o.observers,ob)
}

// 移除观察者
func (o *Object)RemoveObserve(ob Object){
  for i,o := range o.observers{
    if o == ob {
      	o.observers = append(o.observers[:i], o.observers[i+1:]...)
    }
  }
}

// 通知所有已注册的观察者，变更自身的状态
func (o *Object)Notify(){
  for _, ob := range o.observers {
       ob.Update(o.state)
    }
}

```

