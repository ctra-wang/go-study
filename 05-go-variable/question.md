# 1、逻辑题

- 1、ABD
- 2、BCD
- 3、BDE
- 4、BCE
- 4、AB


# 2、语⾔特性（Golang）

### 1、输出代码结构
> [0,0,1,2]
> 
> 说明：make切片后第一个参数为长度，第二参数不写为cap=len，int得初始化用0代替。
> 然后append 进来新的元素（同时进行了扩容）

### 2、AB

> 说明： init 类似java得无参构造函数，由于package加载顺序决定优先顺序，
> 同一个package下按上下顺序加载


### 3、b
### 4、c
### 5、B

# 3、计算机⽹络、操作系统、数据库等

1. TCP建⽴连接需要进⾏ 3 次握⼿，断开连接需要进⾏ 4 次分⼿。
2. 请简要解释Cookie与Session的差别（三句话以内）。
> cookie存放在客户端，session存放在服务端
> cookie无状态，seesion有状态
3. 解释下列http状态码的含义： 301：永久重定向 302：重定向 400：错误请求 500：服务端错误
4. grep：查询进程 top：查看资源使用
   find：查找 mv：移动复制
   cat：查看 du：root权限
   wc： chmod：更改权限 
   crontab：定时任务 awk：
   uniq： ps：当前进程状态
5.  rwx：read 4 write 2 execute 1
6. explain：sql优化，尽量让type命中索引，而不是all全文索引
   show processlist：查看那些语句正在访问mysql，判断是否被锁

7.请列举你所熟知的或者你认为常⻅的软件设计模式
> 模板方法：抽象出一套公共模板，供子模块继承，并融入新的代码
> 动态代理：需要反射，运行时进行变量的修改
> 工厂：无需自己定义，自带一套弗雷定义好得属性和方法，继承即可

# 4、算法与编程能⼒

```go
func main() {
	str := "2020-05-1619:20:34|user.login|name=Charles&location=Beijing&device=iPhone"
	arr := strings.Split(str, "|")

	res := map[string]string{}
	if len(arr) > 0 {
		data := strings.Split(arr[len(arr)-1], "&")
		if len(data) > 0 {
			for _, temp := range data {
				data1 := strings.Split(temp, "=")
				if len(data1) == 2 {
					res[data1[0]] = data1[1]
				}
			}
		}
	}
	fmt.Println(res)
}
```


```go
func Test11(arr []int) int {

	length := len(arr)
	if length <= 0 {
		return 0
	}
	start := 0

	// 首位大于0
	if arr[0] > 0 {
		return arr[0]
	}
	// 末位大于0
	if arr[length-1] < 0 {
		return arr[length-1]
	}

	for {
		mid := (start + length) / 2
		if arr[mid] > 0 {
			if arr[mid-1] > 0 {
				length = mid - 1
			} else {
				if arr[mid] > -arr[mid-1] {
					return arr[mid-1]
				} else {
					return arr[mid]
				}
			}
		} else if arr[mid] < 0 {
			if arr[mid-1] < 0 {
				length = mid + 1
			} else {
				if -arr[mid] > arr[mid+1] {
					return arr[mid+1]
				} else {
					return arr[mid]
				}
			}
		} else {
			return arr[mid]
		}
	}

}
```