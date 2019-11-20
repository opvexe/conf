# `casbin`角色权限控制

```json
//sub   "alice"// 想要访问资源的用户.
//obj  "data1" // 要访问的资源.
//act  "read"  // 用户对资源执行的操作.

# Request definition
[request_definition]
r = sub, obj, act

# Policy definition
[policy_definition]
p = sub, obj, act

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))

# Matchers
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
```

> `Request definition` 即为请求定义，代表了你可以传入什么样的参数来确定权限
>
> `Policy definition` 代表了规则的组成，这两处也就是我上面说的，这个模型规定了，谁 (sub) 可以对什么资源 (obj) 进行什么操作(act)。
>
> `Policy effect` 则表示什么样的规则可以被允许

