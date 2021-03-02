golang interface

---

在 Go 语言中，接口是一组方法的集合，但不包含方法的实现、是抽象的，接口中也不能包含变量。当一个类型 T 提供了接口中所有方法的定义时，就说 T 实现了接口。接口指定类型应该有哪些方法，类型决定如何去实现这些方法。

```go
type Shape interface {
    Area() float32
}

func main() {
        var s Shaper
        fmt.Println("value of s is ", s) // value of s is  <nil> ???
        fmt.Printf("type of s is %T\n", s)  // type of s is <nil> ???
}
```

**接口类型值**

- 静态类型
- 动态类型

​    变量的类型在声明时指定、且不能改变，称为静态类型。接口类型的静态类型就是接口本身。接口没有静态值，它指向的是动态值。接口类型的变量存的是实现接口的类型的值。该值就是接口的动态值，实现接口的类型就是接口的动态类型。

​    动态类型在上面已经讲过，动态值是实际分配的值。记住一点：**<u>当且仅当*动态值*和*动态类型*都为nil 时，接口类型值才为 nil</u>**。





五个关键点

- interface 是一种类型
- interface变量存储的是实现者的值
- 如何判断 interface 变量存储的是哪种类型
- 空的interface
- interface 的实现者的 receiver 如何选择

Languages with methods typically fall into one of two camps: **prepare tables for all the method calls statically (as in C++ and Java)**, or **do a method lookup at each call (as in Smalltalk and its many imitators, JavaScript and Python included) and add fancy caching to make that call efficient**. Go sits halfway between the two: **it has method tables but computes them at run time**. I don't know whether Go is the first language to use this technique, but it's certainly not a common one. (I'd be interested to hear about earlier examples; leave a comment below.)



