# Gin Example With Generic - Gin Example 的泛型实现
### 代码目录结构
```text
├── cmd                             // 程序入口
├── config                          // 读取配置文件方法
├── config_file                     // 配置文件
├── controller                      // 控制器层
│   └── api                         // 控制器层的api
│       ├── universal               // 通用api（404、健康检测等接口）
│       └── v1                      // 具体的api
├── generic                         // 泛型实现
│   ├── controller.go               // 控制器层的泛型实现
│   ├── model.go                    // 重写gorm的Model结构体（为了实现特定方法）
│   ├── paginate.go                 // 分页器
│   ├── repository.go               // 数据库操作的泛型实现
│   ├── request.go                  // 请求体的接口
│   └── service.go                  // service层的泛型实现
├── pkg                             // 通用包
│   ├── ecode                       // 错误码
│   ├── errors                      // 重写pkg/errors包（结合错误码）
│   ├── http                        // 对Gin的一些方法的封装
│   │   ├── bind                    // 数据绑定
│   │   ├── middleware              // 中间件
│   │   │   ├── cross_domain        // 跨域
│   │   │   ├── handle_error        // 统一的错误处理
│   │   │   └── logging             // 日志（access log）
│   │   └── render                  // 数据渲染
│   ├── json                        // json序列化（fast）
│   ├── log                         // 日志（app log）
│   ├── mysql                       // mysql连接函数
│   └── validator                   // 参数校验
├── router                          // 路由
├── server                          // 服务
├── service                         // 服务层
├── store                           // 数据层
│   ├── db_engine.go                // 数据库引擎
│   ├── model                       // 模型
│   └── repository                  // 数据库操作
├── tools                           // 第三方工具（ecode生成）
└── types                           // 类型
    ├── config                      // 配置文件结构体
    ├── const.go                    // 常量 
    ├── paginate                    // 分页结构体
    ├── request                     // 入参结构体
    └── result                      // 返回结构体
```