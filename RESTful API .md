# RESTful API 

> REST全称是Representational State Transfer


## 使用 
> 1. bee new project </br>
> 2. cd $GOPATH/src/project </br>
> 3. bee run -gendoc=true -downdoc=true  //运行项目并获得API自动化文档 </br>
> 4. 打开  http://localhost:8080/swagger/#

## 报错

1.打开swagger文件

```
{
  	window.onload = function() {
  	const ui = SwaggerUIBundle({
    url: "swagger.json",
    dom_id: '#swagger-ui',
    validatorUrl: null,
}
```

## 状态码

```
GET

200（OK） - 表示已在响应中发出
204（无内容） - 资源有空表示
301（Moved Permanently） - 资源的URI已被更新
303（See Other） - 其他（如，负载均衡）
304（not modified）- 资源未更改（缓存）
400 （bad request）- 指代坏请求（如，参数错误）
404 （not found）- 资源不存在
406 （not acceptable）- 服务端不支持所需表示
500 （internal server error）- 通用错误响应
503 （Service Unavailable）- 服务端当前无法处理请求

POST
200（OK）- 如果现有资源已被更改
201（created）- 如果新资源被创建
202（accepted）- 已接受处理请求但尚未完成（异步处理）
301（Moved Permanently）- 资源的URI被更新
303（See Other）- 其他（如，负载均衡）
400（bad request）- 指代坏请求
404 （not found）- 资源不存在
406 （not acceptable）- 服务端不支持所需表示
409 （conflict）- 通用冲突
412 （Precondition Failed）- 前置条件失败（如执行条件更新时的冲突）
415 （unsupported media type）- 接受到的表示不受支持
500 （internal server error）- 通用错误响应
503 （Service Unavailable）- 服务当前无法处理请求
```