package main

import "fmt"

func main() {
	done := make(chan bool, 3)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v //v 取消注释，则为正确输出

		go func() {
			// defer 延迟操作，相当于压栈，在方法的}的之前再次出栈执行
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println(fmt.Sprintf("当前变量v=%v，捕捉到异常了", v))
				}
			}()
			if v == "a" {
				panic("pppppppppppppppppppp")
			}
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
