## 目标
- 路由(Routing)：将请求映射到函数，支持动态路由。例如'/hello/:name。
- 模板(Templates)：使用内置模板引擎提供模板渲染机制。
- 工具集(Utilites)：提供对 cookies，headers 等处理机制。
- 插件(Plugin)：Bottle本身功能有限，但提供了插件机制。可以选择安装到全局，也可以只针对某几个路由生效。


## 跟踪

- [X] 将路由(router)独立出来，方便之后增强。
- [X] 设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
- [X] 实现路由分组控制
- [X] 自定义中间件
- [X] 异常处理
- [X] 支持静态文件及模板