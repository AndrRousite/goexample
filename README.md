### Golang

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

##### 并发
1. goroutine
2. channel
