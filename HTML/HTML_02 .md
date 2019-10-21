# HTML_02 

## 表单标签
概念:用于采集用户输入的数据，用于和服务器进行交互。

> form用于定义表单的，可以定义一个范围，范围代表采集用户数据的范围。

```html
<form>
用户名:<input></br>
密码：<input></br>
</form>
```

form标签的属性:

   + action: 指定提交数据的url  `<form action="#">`
   + method: 指定提交方式 `<form method="get">`
   + type :类型 `<input type="submit" value="登录">`
   + 表单的数据要想被提交，表单项。必须指定其name属性。
   + name `<input name="password">` 密码输入框
   + text `<input name="text">`  文本输入框 默认输入框
   + radio `<input name="radio" name="gender">` 单选框 实现单选选择框name属性值必须一样gender。选中"on" 
   + checkbox: 复选框 `<input name="checkbox" name="hobby" checked="checked"> 逛街 </br>` //默认被选中check
   + label: 用于显示输入项文字描述信息
   + placeholder:文本输入的占位文字 placeholder="请输入密码"
   + file:文件框 `<input type="file" name="file">`
   + 隐藏域: `<input type="hidden" name="id" value="aaa"></br>` 用于提交一些信息
   + button: `<input type="button" value="按钮">` 
   + `<input type="image" src="./image/logo.png">` //图片按钮 src：指定图片的路径
   +  date： 时间日期选择器
   +  datetime： 时间日期选择器
   +  email: 邮箱输入框


#### 下拉列表

   + select:选择框  `<select name="province"> <option value=""> 请选择</option> <option value="1" selected> 上海</option></select>` option：选择项 value:提交的数值 selected:默认选项


#### 文本域

 + textarea: 文本域  name:提交指定他的name
 + cols : 列数
 + rows : 默认多少行

```html
<textarea cols="20" rows="5" name="data"> 
</textarea>
```

## 注册页面

```html
<form action="#" method="post"> 
<table border="1" algin="center" width="500">
<tr> 
<td> <label for="username"></lable>用户名<td>  //for 指定username的id
<td> <input type="text" name="username" id="username"><td>
</tr>

<tr> 
<td> <label for="password"></lable>密码<td>  
<td> <input type="password" name="password" id="password"><td>
</tr>

<tr> 
<td> <label for="email"></lable>email<td>  
<td> <input type="email" name="email" id="email"><td>
</tr>

<tr> 
<td> <label for="name"></lable>姓名<td>  
<td> <input type="text" name="name" id="name"><td>
</tr>

<tr> 
<td> <label for="iphone"></lable>手机号<td>  
<td> <input type="text" name="iphone" id="iphone"><td>
</tr>


<tr> 
<td> <label ></lable>性别<td>  
<td> <input type="radio" name="gender" value="male">男<td>
<td> <input type="radio" name="gender" value ="female">女<td>
</tr>

<tr> 
<td> <label for="birthday"></lable>出生日期<td>  
<td> <input type="date" name="birthday" id="birthday"><td>
</tr>


<tr> 
<td> <label for="checkcode"></lable>验证码<td>  
<td> <input type="text" name="checkcode" id="checkcode">
<image src= "./image/checkcode/1.png">
<td>
</tr>

<tr> 
<td> <input type="submit" name="login" id="login">
<td>
</tr>

</table>
</form>
```

