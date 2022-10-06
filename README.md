# hi-golang

## Courses

- Google资深工程师深度讲解Go语言
- Go高级工程师实战营二期(GoCN)
- Go进阶训练营第5期

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
