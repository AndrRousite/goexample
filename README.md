### Golang

####### 第一章
```$xslt
1. 安装的三种方式：源码安装、标准包安装、第三方工具安装
2. 环境变量配置：$GOROOT / $GOPATH
3. Go的常用命令：编译、安装、格式化、测试等
4. Go的开发工具：LiteIDE/Goglang/Sublime/VScode/Atom/vim/Emacs/Eclipse/IDEA
```
#######  第二章(基础)
```
25个关键字
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var

    
##### 基础
1. 包、变量、函数
2. 流程控制
    ```
    if/switch/for
    ```
3. 数据类型
    ```
    基本数据类型
        bool
        byte //uint8 的别名
        rune //int32 的别名 //代表一个Unicode码
        int int8 int16 int32 int64
        uint uint8 uint16 uint32 uint64 uintptr
        float32 float64
        complex64 complex128
    复杂类型
        struct
        slice
        map
    ```

##### 方法和接口
1. 方法
    ```
    Go 没有类。然而，仍然可以在结构体类型上定义方法。
    方法接收者 出现在 func 关键字和方法名之间的参数中。
    ```
2. 接口
    ```
    接口类型是由一组方法定义的集合。
    接口类型的值可以存放实现这些方法的任何值。
    ```

##### 并发(协程和管道)
1. goroutine
2. channel
```
####### 第三章(Web基础)
```$xslt
HTTP协议, DNS解析, go实现简单web server
```
####### 第四章(表单)
```$xslt
用户登录，正则匹配，预防跨站脚本和多次表单提交，文件上传
```
####### 第五章(数据库)
```$xslt
1. 数据库驱动（database/sql）：类似jdbc
2. ORM（ex:beego orm）
3. NoSql支持
```
####### 第六章(Session和Cookie)
```$xslt
1. go 原生支持cookie，但不支持session
2. session的生命周期：生成sessionId->session管理器以及session的存储（内存、文件、数据库）->将sessionId发送客户端
3. 防止session劫持
```
####### 第七章(文本文件处理)
```$xslt
1. 字符串处理
    1. strings(常用)
        Contains:查文档
        Join:
        Index:
        Repeat:
        Replace:
        Split:
        Trim:
        Fields:
    2. strconv
        Append系列函数将整数等转换为字符串后，添加到现有的字节数组中
        Format系列函数把其他类型的转换为字符串
        Parse系列函数把字符串转换为其他类型
2. xml和json编码和解码
3. 正则处理(Match/Compile/Expand)
4. 模板引擎(New(可以省略)->Parse->Execute)
5. 文件操作
    1. 目录操作（Mkdir/MkdirAll,Remove/RemoveAll）
    2. 文件操作
        1. 创建：Create/NewFile
        2. 打开：Open(只读)/OpenFile(可以设置权限)
        3. 写：Write/WriteAt(指定位置开始写入)/WriteString(写入字符串)
        4. 读: Read/ReadAt(指定位置开始读取数据)
        5. 删除: 和删除目录的函数一样，os.Remove/os.RemoveAll既可以删除文件也可以删除目录
```
####### 第八章(Web服务)
```$xslt
1. Socket 
2. WebSocket
3. REST
4. RPC(Remote Procedure Call Protocol)——远程过程调用协议
```
####### 第九章(安全与加密)
```$xslt
1. 预防CSRF攻击(跨站请求伪造)
2. 避免XSS攻击(跨站脚本攻击)
3. 避免SQL注入
4. 加密和解密数据，对称加密和非对称加密，单向加密或散列算法和双向加密
```
####### 第十章(国际化和本地化:i18n/L10N)
```$xslt
1. 设置默认地区
2. 本地化资源
3. 国际化站点
```
####### 第十一章(错误和异常)