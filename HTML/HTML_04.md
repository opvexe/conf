## DOM

```html
功能：控制html文档的内容
代码：获取页面的标签（元素对象） Element
	document.getElementById("id值")
要干的事情: 操作Element对象
	1.设置其属性
	2.修改标签体的内容
	3.查看API文档，要img的属性去设置
	
<img src="img/off.gif" id="light">
//修改标签体的内容
<h1 id="title">阿里巴巴</h1>

//推荐写在body后面
<script>
//通过id获取元素对象
var light = document.getElementById("light")
light.src = "img/on.gif"

//获取h1标签
var title = document.getElementById("title")
//修改标签体内容
title.innerHTML = "不识妻美刘强东"
</script>
</body>

```

## 事件

功能:某些组件被执行了某些操作后，触发某些代码的执行。

```html
1.如何绑定时间
	1.直接在Html标签上，指定事件的属性，属性值就是js代码。
	2.事件 onclick --单击事件
	3.通过js获取元素对象，指定事件属性，设置一函数
	
代码:

<img src="img/off.gif" id="light" onclick="alert("我被点了");">
<img src="img/off.gif" id="light" onclick="func();">

<img src="img/off.gif" id="light">

<script>
方法1：
function fun(){
	alert("我被点了");
}
方法2：
var light = document.getElementById("light")
light.onclick = fun;
light.onclick = function fun(){
	alert("我被点了");
}
</script>
```

