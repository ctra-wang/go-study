# go-study
go基础学习相关


## 01-go-loop 循环

从多个角度说明在循环中，如果不对变量复制，则会根据引用一直得出最后一个值
- array-iterator.go
- goroutine-iterator.go
- int-iterator.go

[go-package-README](https://github.com/Holyson/go-study/blob/master/01-go-package/readme.md)

## 02-go-switch-case 用法

这里只针对多条件的一个说明
在代码中超过3层以上的if判断，推荐使用switch-case代理if判断

[深入理解 CPU 的分支预测(Branch Prediction)模型](https://blog.csdn.net/qq_40227064/article/details/124326907)

[一文告诉你CPU分支预测对性能影响有多大](https://xindoo.blog.csdn.net/article/details/101762198?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-101762198-blog-124326907.pc_relevant_3mothn_strategy_recovery&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-101762198-blog-124326907.pc_relevant_3mothn_strategy_recovery&utm_relevant_index=2)

最后附赠一道算法题：
question2：求数组 [1, 3, 5, 7, 8] 中2个元素之和等于8的所有坐标 即：[0 3] [1 2]