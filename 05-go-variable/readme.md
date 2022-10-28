

# go-variable


## 1、const-iota 常量&iota关键字

iota记住下面2点
- 遇见const关键字就置为0（所以只能在const内部使用）
- 每增加一行则iota的值顺次+1
- 注意：_为省略，但是会保留 iota+1的值
