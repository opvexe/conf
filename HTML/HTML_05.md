## CSS与HTML结合

#### 内联样式 不推荐使用
```html
案例:
<div style="color:red;">hello css</div>
			  key:value 键值
```
#### 内部样式  比较常用
```html
在head标签内，定义style标签，style 标签的标签内容就是css 代码

案例:
head标签内:
	<style>
	div{
		color:red;
	}
	</style>
	
<div>hello css </div>

```
#### 外部样式
```html
1.在定义css资源文件
2.在head标签内，定义link标签，来引入外部的资源文件

案例:
div{
	color:red;
}

在head引入:
<link rel="stylesheet" href="css/a.css">

```

+ 问题: 
	+ 1,2,3种方式作用范围越来越大。
	+ 1 方式不常用，后期常用2，3方式
	+ 3种格式可以写为: `<style> @import "css/a.css" </style>`
	
	
	
## CSS语法

#### 格式
```html
选择器 {
	属性名:属性值;
	属性名:属性值;
	....
}

*** 选择器: 筛选具有相似特征的元素
注意:
	每一对属性需要使用; （分好隔开），最后一个可以不加分号
	
p {
	color:red;
	font-size:30px; 像素px 可以加分号，也可以不加分号
}
```

#### 选择器  (筛选具有相似特征的元素)

分类：

+ 基础选择器
	+ id选择器: 选择器聚义的id属性值的元素。
	+ 元素选择器:选择具有相同标签名称的元素
	+ 类选择器: 选择具有相同的class属性值的元素
	
	

id选择器案例:

```html
<div id="div1"> blog</div>

<style> #div1{ color:red;} </style>

#语法:属性名  #语法表示：选择id 为div1的标识
```

元素选择器案例：

```html
div{
	color:red;
}

<div> blog</div>

注意:id 选择器优先级比元素选择器高

```

类选择器案例：

```html
. class1{
	color:red;
}

<p class="class1"> 北京大学</p>

语法: .class属性值
注意： 类优先级的优先级高于id

```
优先级： 类选择器 > id选择器 > 元素选择器

**扩展选择器**

1. 选择所有元素:  语法: *{}
2. 并集选择器:  语法：选择器1,选择器2{}
3. 子选择器: 语法: 选择器1 选择器2{}  含义：筛选选择器1下的选择器2 
4. 父选择器: 语法：选择器1>选择器2{}  含义：筛选选择器2的父元素的选择器1
5. 属性选择器: 选择元素名称，属性名=属性值的元素 语法: 元素名称[属性名="属性值"]{}   一般用于input标签
6. 伪类选择器: 选择一些元素的状态 语法: 元素:状态{}

```html

<div>
	<p>北京大学</p>
</div>

子选择器案例:
<style>
	div p{
		color:red;
	}
</style>

父选择器案例:
<style>
	div > p{		//p元素的父标签的div
		bordor:1px solid;
	}
</style>

```



```html
属性选择器案例:
<input type="text"> 
<input type="password"> 

样式：
<style>
	input[type="text"]{
		bordor:5px solid;
	}
</style>

```



```html
伪类选择器案例:
如：状态
	*link:初始化状态
	*visited:被访问过的状态
	*active:正在访问的状态
	*hover:鼠标悬浮状态
	
	
<a href="#"> 超链接 </a>

<style>
	a:link{		//初始化状态  类似ios button点击状态，离开，按住时button按钮颜色的变化
		color:pink;
	}
	a:visited{		
		color:green;
	}
	a:active{		
		color:red;
	}
</style>

```

#### 属性

1.字体，文本

+ font-size:字体大小
+ color:文本颜色
+ text-align:对其方式
+ line-height:文本的高度，行高

```html
<style>
 p {
 	color:red;
 	font-size:30px;
 	text-align:center;
 	line-height:100px; 	//整个字占100个高度
 	//border 边框
 	border:1px solid red;  //边框1px,solid 实线
 	
 }
 
 div{
 	border:1px solid red; 
 	/*
 	 * 尺寸
 	 */
 	height:200px;
 	width:200px;
 	background:url("image/logo.png") no-repeat center;//no-repeat 不重复，center居中显示
 }
 
</style>

<p>北京大学</p>
<div></div>
```

2.背景

+ background: 背景颜色 background-color，背景图片 background-img

3.边框

+ border: 设置边框，符合属性

4.尺寸

+ width:宽度
+ height:高度

5.盒子模型  （一个一个的元素当做一个一个盒子来看待，然后控制盒子的大小和尺寸）主要作用是控制布局

+ margin:外边距
+ padding:内边距 **默认情况下，内边距会影响整个盒子的大小，设置盒子的属性，让iwdth和height就是最终盒子的大小
+ box-sizing: border-box;
+ float:浮动
+ 左浮动：left
+ 右浮动: right





























