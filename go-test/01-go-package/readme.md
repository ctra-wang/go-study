
@[TOC](go-package 包管理)

# 01、go-package 包管理

> 可以参考文章：
> [包与依赖管理：https://www.liwenzhou.com/posts/Go/11-package/](https://www.liwenzhou.com/posts/Go/11-package/)



## 一、go init() 函数

### 1、go 在package内的init() 方法
这里在包下的init() 方法会被引入此包的时候优先执行
类似默认的无参构造函数（注意：在包内需要首字母小写）

全部代码如下图所示：
![img.png](file/img.png)


## 二、package 加载顺序

### 1、package 的加载顺序图

![img.png](file/img2.png)

### 2、import 的加载顺序图
![img.png](file/img1.png)

