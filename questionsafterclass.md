###Var和自动赋值的区别
var 是声明变量
自动赋值是一般在函数的返回值中体现。
var关键字声明全局变量，但是:=这种方式是不能用在全局变量中的。
使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。可以同时声明多个变量

###Benchmark
Benchmark 被称为基准测试，也叫性能测试和压力测试

func BenchmarkXxx(b *testing.B) 函数格式

执行命令 go test -v -bench=. -cpu=8 -benchtime="3s" -timeout="5s" -benchmem 
benchmem：输出内存分配统计
benchtime：指定测试时间
cpu：指定GOMAXPROCS
timeout：超时限制 

####执行结果                 
```go
//Benchmark的名字 - CPU  //循环的次数           //平衡每次执行时间        //表示每次执行分配的内存（字节）
BenchmarkRepeat-12      20000000              113 ns/op              16 B/op                    
//表示每次执行分配了多少次对象
4 allocs/op 
PASS      //在那个目录下执行的go test         //累计耗时
ok      tdd-practice/class3              2.407s 
```



