# HTML——01 [参考资料](https://www.w3school.com.cn/tags/tag_hr.asp)
## 文件标签
 
   + htm:html 文档的根标签
   + head:头标签，引用外部资源
   + title:定义标题的标签
   + body: 体标签
   + <!DOCTYPE html>:定义文档类型 
   + <html lang="en"> :lang 属性 中文ch 
   +  <meta charset="UTF-8">: 指定字符集 utf-8 //保证不乱码

## 文本标签

   + `<h1> - <h6> `: 标题 自带换行效果 h1 - h6 字体大小逐渐递减
   + `<p> `	： 段落
   + `<br>  ` ：换行
   + `<hr>`  ：展示一条水平线 color颜色 width宽度 size 高度 align 对齐方式center 居中,left 左对齐，right 右对齐 
   + `<b> ` : 字体加粗
   + `<i> ` : 字体斜体
   + `<font> ` :对应字体标签 可以修改字体的

   + html属性  不支持 不推荐
    + `<hr color="red" widths="200" size="10" align="left">` size:文本的高度 
    + `<font color="red" size="5" face="楷体">`  face：字体字
  
属性的定义: 
	
  + color 
  	+ 1.英文单词: red,green blue 
  	+ 2.RGB(值1，值2，值3) 范围 0-255 eg:rgb(0,0,255) //红绿蓝  很多浏览器不支持了
  	+ 3.#值1值2值3： 值的范围 00-ff 之间。16进制的范围  **#**  eg:#FFFFFF 这样的写法
  
  + width
  	+ 1.数值 :width = "20" ，数值的单位，默认是px(像素)
  	+ 2.%数值: 占比相对于父元素的比例 50%  body 的相对比例 
  	
  + center  标签


## 图片标签

  + `<img src="image/banner.png" algin="right" alt="古镇" width="500" height="300"/>` :src = "指定图片存放的位置"  alt：加载失败展示文字
    + 一般写相对路径 "./":当前目录  eg: "./image/banner.png"
    + "../": 上一级目录

## 列表标签

+ 有序列表
	   + ol : order list   : type ="I"  ,type ="A" ... 
	   + li : list i
+ 无序列表
	+ ul:

// 默认 属性 digt,square ,cicle   ==== ** =====


```
//有序
eg:
	<ol type ="A" start="5">   FEGH 排序
    <li> 起床</li>
    <li> 刷牙</li>
    <li> 上班</li>
</ol>

//无序
eg:
	<ul type ="digt">   FEGH 排序
    <li> 起床</li>
    <li> 刷牙</li>
    <li> 上班</li>
</ul>

```

## 连接标签
	 
 + a: 超链接标签 
 	+ `href 属性:  eg: <a href="http://www.baidu.com">点我</a>   //href表示了点击超链接跳转的位置  ***本页面内部跳转 ***`
 	+ `_blank: eg: <a href="http://www.baidu.com" target="_blank">点我</a>   //点击打开新窗口跳转`
 	+ `_selft: 默认值`
 	+ `target 指定一个打开资源的方式 : 默认值_self,_blank在一个空白页打开`

eg: `<a href="http://www.baidu.com"> <img src="image/banner.png"></a>`

## 表格标签

  + table :定义表格
    + tr : 定义行
    + td : 定义单元格
    + th : 定义表头单元格
  + 属性
  	+ border :边框 
  	+ cellpadding : 单元格与内容之间的空白
  	+ cellspadding : 单元格之间的空白
  	+ bgcolor: 背景色
  	+ align:对齐样式
  	+ caption: 标题
  	+ <thead> : 表示表格的头部分  只是表示 后期配合css
  	+ <tbody> : 表示表格的体部分
  	+ <tfoot> : 表示表格的脚部分

```
<table border="1" width="50%" cellpadding="0" cellspadding="0">
    <caption>学生信息表</caption>
 <tr>
     <td>编号</td>
     <td>姓名</td>
     <td>成绩</td>
 </tr>
    //数据
    <tr>
        <td>1</td>
        <td>小龙女</td>
        <td>100</td>
    </tr>

    <tr>
        <td>2</td>
        <td>杨过</td>
        <td>10</td>
    </tr>
</table>
```


## 块标签

  + div :没有任何效果 ， 特点会换行  每一个div 占满一整行，块级标签
  + span : 本来没有任何效果，包裹效果  结合css使用,默认在一行展示，成为行内标签。不会换行

## 语义化标签

```
eg:htm4
<div id ="header"></div>
<div id ="footer"></div>
eg:html5		//增加了可读性  
<header>
<div id ="header">
</div>
<header>
```







