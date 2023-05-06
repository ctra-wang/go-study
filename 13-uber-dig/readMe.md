首先给出代码示例，供大家更直观地感受通过 dig 实现依赖注入、路径梳理、bean 复用的能力：

- 存在 bean A、bean B，其中 bean A 依赖于 bean B
- 声明 bean A 和 bean B 的构造器方法，A 对 B 的依赖关系需要在构造器函数 NewA 的入参中体现
- 通过 dig.New 方法创建一个 dig container
- 通过 container.Provide 方法，分别往容器中传入 A 和 B 的构造器函数
- 同归 container.Invoke 方法，传入 bean A 的获取器方法 func(_a *A)，其中需要将获取器函数的入参类型设置为 bean A 的类型
- 在获取器方法运行过程中，入参通过容器取得 bean A 实例，此时可以通过闭包的方式将 bean A 导出到方法外层
