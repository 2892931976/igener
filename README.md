[![Build Status](https://travis-ci.org/cloudaice/igener.svg?branch=master)](https://travis-ci.org/cloudaice/igener)

# Id生成器

### 产生原因

在实际业务中，往往需要产生一个唯一的Id来标示某些对象或者数据。例如订单号，数据库记录，某些队列数据的标示等等。在一些需要将并发转成串行处理的模型中，也需要生成一个唯一id去标识一个对象。

### 设计思路

在设计之初，我便联想到Mongodb的`_id`设计方式。于是就借用了它的思想。具体如下：

产生出来的`id`是一个由24个字符组成的字符串，本质上是一个16进制的数值。前八位是时间，接着6位是机器名的哈希，在接着4位是`pid`，最后6位是自增序号。这样保证在时间，主机，进程这单个维度上是唯一的，而且还保留了2的24次方的自增序号，这样即使在分布式环境下每一个id生成器产生的id都是唯一的。

### 例子

    import (
        "igener"
        "fmt"
    ) 
    
    func main(){
        ig := igener.NewIGener()
        id := <-ig
        fmt.Println(id)
    }

### 安装

    go get "github.com/cloudaice/igener"
