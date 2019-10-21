## JavaScript

#### 基本语法

+ 基本语法
	+ 与html结合方式
		+ 内部的js 
			* 定义`<script>` ,标签体的内容就是js代码
		+ 外部的js 
		   * 定义`<script>`,通过src属性引入外部js文件
		   * 注意: `<script>` 可以定义在htm页面的任何地方，但定义的位置会影响执行的顺序。
		   * `<script>` 可以定义多个
	+ 注释
		+ 单行注释: //注释内容
		+ 多行注释:/* 注释内容*/ 
	+ 数据类型
		+ 原始数据类型(基本数据类型):
		   * number:整数，小数，NaN（一个不是数字的数字的类型）
		   * string:字符串 "abc",'abc'
		   * boolean:
		   * null:对象为空的占位符
		   * undefined:未定义。如果一个定义没有给初始化值则会被默认赋值为undefined
		+ 引用数据类型(对象): 
			*  
	+ 变量
	  * int a = 3
	  * 强类型：开辟变量存储空间时，定义了空间将来存储数据的数据类型。只能存储固定的数据类型。
	  * 弱类型：开辟变量存储空间时，不定义空间将来的存储数据类型。可以存放任意数据类型。
	  * 语法: var 变量名 = 初始化的值
	+ 运算符
	  * typeof :知道变量类型是什么样 语法:typeof(变量名) 
	+ 流程控制语句 	
	  * 其他类型转boolean
	  * 1.number:	0或NaN为假,其他为真
	  * 2.string:除了空字符串""，其他都是真
	  * 3.null&undefined: obj如果是null 则为假
	  * 4.对象 
	  		+  创建: var fun = new Function(形式参数，方法体) //了解
	  		+  function 方法名称(形式参数列表){ 方法体 }
	  		+  特点：方法定义时，形参的类型不用写，如： function fun2(a,b){ alert(a+b) } 调用:fun2(3,4)
	  		+  var fun3  = function(a,b){ alert(a+b) } 调用:fun3(3,4)
	+ js特殊语法

1.内部的js

```html
//在html内部定义script标签（可以写在html任意位置）
<script>
	alert("hello word")
</script>

注意: script 执行顺序根据放置的位置
```

2.定义变量

```html
案例一：
<script>
var a = 3;
alert(a);
a = "abc";
alert(a);

案例2：
var num = 1;
var num1 = 1.2;
var num2 = NaN;

//输出到页面显示
document.write(num+"</br>"); //页面换行+"</br>"
document.write(num1+"</br>");
document.write(num2+"</br>");

案例3:
var str  = "abc";
var str2 ='edf';

document.write(str+ typeof(str) +"</br>");
document.write(str2+"</br>");


案例4：
var obj  = null;
var obj2 =undefined;
var obj3;
document.write(obj +"</br>");
document.write(obj2 +"</br>");
document.write(obj3 +"</br>");

案例5:
for (var i :=1;i<=9;i++){
	for (var j:=1;j<i;j++){
		document.write(i*j+"&nbsp"); //空格&nbsp
	}
}

</script>

```

2.外部的js

```html
//使用src加载外部的文件----> 外部js文件一般放在js目录下 
<script src="js/a.js"> </script>

//a.js文件中
alert("hello word")
```

####基本对象

+ 函数对象 

```html
1.函数对象
<script>

//1.创建
1. var fun = new Function(形式参数列表，方法体);
2. function fun2(a,b) { 方法体 }	//方法定义时，形参的类型不用写
3. var fun1 = function(a,b) { 方法体 }
4. function add(a,b){ retrun a+b }
5. 求任意个数的和: function add(){ retrun alert(arguments[0])}  //arguments 内置数组

//方法


//属性
length:代表形参的个数

//特点
1.方法定义时，形参可以不用写
2.方法是一个对象，如果定义名称相同的方法，然后会覆盖 后来定义的会覆盖前面的
3.js中，方法的调用与方法名称有关，和参数列表无关。
4.在方法声明中，有一个隐藏的内置对象，叫做arguments,封装了所有的实际参数

//调用
方法名称：（实际参数列表）
fun2(1,2)
fun2(1)
fun2()
fun2(1,2,3)

function add(){ 
for (var i=0;i<arguments.length,i++){
	arguments[i]
}
	retrun alert(arguments[0])
} 

</script>
```


+ 数组对象 

```html
//创建 
1.var arr  = new Array(元素列表);
2.var arr = new Array(默认长度);
3.var arr = [元素列表];

//特点 
1.js中数组的类型是可变的。
2.js中数组的长度是可变的。

//属性
length : 数组的长度

//创建方式
var arr1 = new Array(1,2,3)
var arr2 = new Array(5)	//创建数组,表示数组的长度，而非数组的列表
var arr3 = [1,2,3,4];
var arr4 = new Array(); //创建空的数组

//打印
document.write(arr1+"</br>")
document.write(arr2+"</br>")
document.write(arr3+"</br>")

var arr5 = [1,"abc",true];
document.write(arr5+"<br>")

//访问数组元素
document.write(arr5[0]+"<br>")
arr5[10] = "hehe" //数组的长度是可变的
document.write(arr5+"<br>")

//方法
1.join() : 将数组中的元素拼接成字符串，按照指定的分割符拼接成字符串 arr.join() //无参数：默认逗号；有参join(-):按照-分割
2.push(): 想数组的尾部追加元素。 arr.push(11)
```

+ 日期对象

```html
//创建
1.var date = new date();

//方法
tolocaleString(): 返回当前date独享对应的时间本地字符格式
getTime(): 获取毫秒值，返回当前时间到1970.1.1毫秒值差

//调用
date.tolocaleString();
date.getTime();
```

+ Math对象

```html
//创建
	特点:不用创建直接使用. Math.方法名

//方法

//属性
Math.PI()
Math.Random() 
```









