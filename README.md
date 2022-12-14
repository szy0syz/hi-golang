# hi-golang

## Demos

- P01 - [Golang Tutorial: upload bulk files to S3 bucket aws golang][1]

## Courses

- Google资深工程师深度讲解Go语言
- Go高级工程师实战营二期(GoCN)
- Go进阶训练营第5期
  - 说实话，课程有点粗糙 
- ZTM - Go Programming
  - 感觉备课细节💯
  - 真不愧是"from an industry expert"👍
- Working with Concurrency in Go
  - 并发高性能的课暂时先不看
- Backend Master Class [Golang + Postgres + Kubernetes + gRPC]
- Let's Go Further! Advanced patterns for building APIs and web applications in Go
  - 细节慢慢的书📚

<img width="989" alt="image" src="https://user-images.githubusercontent.com/10555820/196014564-8aead27c-08c3-4547-97c0-d39b36daf211.png">
<img width="996" alt="image" src="https://user-images.githubusercontent.com/10555820/196014572-53b8f8dd-c2a8-4490-89f3-0f124279eea8.png">
<img width="1316" alt="image" src="https://user-images.githubusercontent.com/10555820/196014506-0fb88fa4-d0a4-4dd7-b82d-fac03dee044c.png">
<img width="990" alt="image" src="https://user-images.githubusercontent.com/10555820/196014593-9b7878cc-a3bf-4834-b942-0bff62fddc4f.png">
<img width="996" alt="image" src="https://user-images.githubusercontent.com/10555820/196014615-568d2a01-a1c2-4d9b-b0a3-09821a20d99e.png">

## Go进阶训练营第5期

### 基础语法

#### 基础语法——main函数

- 无参数、无返回值
- main 方法必须要在 main 包里面
- `go run main.go` 就可以执行
- 如果文件不叫 `main.go`，则需要 `go build` 之后再 `go run`

#### 基础语法——包申明

- package xxxx
- 字母和下划线的组合
- 可以和文件夹不同名字
- 通一文件夹下的**声明需一致**
- 引入包语法形式：`import [alias] xxxx`
- 如果一个包引入了但没有使用会报错
- 匿名引入时前面会多一个下划线

#### 基础语法——string

- string
  - 双引号引起来，内部双引号需用`\`转义
  - 也可以使用 ```
- string 的长度很特殊
  - 字节长度：和编码无关，用len(str)获取
  - 字符数量：和编码无关，用编码库来计算

#### 基础语法——strings包

- string 的拼接直接使用 + 号就可以。注意的是， 某些语言支持 string 和别的类型拼接，但是 golang 不可以
- strings 主要方法(你所需要的全部都可以找 到):
  - 查找和替换
  - 大小写转换
  - 子字符串相关
  - 相等

#### 基础语法——rune类型

- rune，直观理解，就是字符
- rune不是byte!
- rune本质是int32，一个rune四个字节
- rune在很多语言里面是没有的，与之对应的 是，golang 没有 char 类型。rune 不是数字， 也不是 char，也不是 byte!
- 实际中不太常用

```go
type rune = int32
```

#### 基础语法——bool,int,uint,float

- bool:true,false
- int8,int16,int32,int64,int
- uint8,uint16,uint32,uint64,uint
- float32,float64

#### 基础语法——byte类型

- byte，字节，本质是uint8
- 对应的操作包在bytes上

#### 基础语法——类型总结

- golang的数字类型明确标注了长度、有无符号
- golang不会帮你做类型转换，类型不同无法通过编译。也因此，string只能和string拼接
- golang有一个很特殊的rune类型，接近一般语言的char或者character的概念，非面试情况下，可以 理解为 “rune = 字符”
- string遇事不决找strings包

#### 基础语法——变量声明 var

- var，语法: `var {name} {type} = {value}`
  - 局部变量
  - 包变量
  - 块声明
- 驼峰命名
- 首字符是否大写控制了访问性: 大写包外可访问
- golang支持类型推断

#### 基础语法——变量声明 :=

- 只能用于局部变量，即方法内部
- golang使用类型推断来推断类型。数字会被理 解为 int 或者 float64。(所以要其它类型的数 字，就得用 var 来声明)

#### 基础语法——变量声明易错点

- 变量声明了没有使用
- 类型不匹配
- 同作用域下，变量只能声明一次

#### 基础语法——常量声明 const

• 首字符是否大写控制了访问性: 大写包外可访问
• 驼峰命名
• 支持类型推断
• 无法修改值

#### 基础语法——方法声明

• 四个部分:
  • 关键字func
  • 方法名字: 首字母是否大写决定了作用域
  • 参数列表: `[<nametype>]`
  • 返回列表: `[<nametype>]`

#### 基础语法——方法声明(推荐写法)

<img width="413" alt="image" src="https://user-images.githubusercontent.com/10555820/194209634-a0866776-1ff9-4f95-906a-aa5ed4b0a74d.png">

- 参数列表含有参数名
- 返回值不具有返回值名

#### 基础语法——方法声明(看看就好)

<img width="507" alt="image" src="https://user-images.githubusercontent.com/10555820/194209729-d805da1f-00fa-4b48-8072-ad3ac5bfd6af.png">

#### 基础语法——方法调用

<img width="366" alt="image" src="https://user-images.githubusercontent.com/10555820/194209022-9927be73-d8d8-45ed-a185-a29d85542949.png">

- 使用_忽略返回值

#### 基础语法——方法声明与调用总结

- `golang` 支持多返回值，这是一个很大的不同点
- `golang` 方法的作用域和变量作用域一样，通过大小写控制
- `golang` 的返回值是可以有名字的，可以通过给予名字让调用方清楚知道你返回的是什么

[1]: https://programmingeeksclub.com/golang-tutorial-upload-bulk-files-to-s3-bucket-aws-golang/