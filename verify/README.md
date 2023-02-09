# 参数校验工具

提供便捷全面的参数校验，验证您的结构体参数是否符合，可以安场景，使用指定校验方式进行校验

### tag 配置

tag 名称为 verify 然后其中包含所需的参数，内部所有参数都以 name=value 的形式进行多个参数之间使用 ; 进行分割

1. scene 场景，指定校验场景，可以多个场景公用
2. method 校验方法， required 为校验的方法 圆括号内设置校验参数，参数使用 'name':value 的方式设置
3. message 错误信息，指定校验失败时输出的信息

```go
type Model struct {
	ID        int64  `verify:"scene=UPDATE,DELETE;method=required('zero':false);message=ID不能为空"`
}
```

### 调用

在 verify 包内有提供 ParamsVerify 方法

方法定义 ParamsVerify(obj, scenes...) error

```go
func main(){
  m := Model{
    ID: 0
  }
  if err := verify.ParamsVerify(&m, verify_scene.UPDATE); err != nil {
    painc(err)
  }
}
```
