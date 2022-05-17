# Gin Example With Generic - Gin Example 的泛型实现

[![Build Status](https://github.com/tiancheng92/gin_example_with_generic/workflows/Build/badge.svg)](https://github.com/tiancheng92/gin_example_with_generic/actions)


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

### 默认支持JSON、XML、YAML格式返回
* JSON:
    ```bash
    # Input
    curl http://127.0.0.1:8080/api/v1/user | jq
    ````
    ```json
    # Output
    {
      "data": {
        "items": [
          {
            "id": 1,
            "created_at": "2022-05-17T13:55:39.354+08:00",
            "updated_at": "2022-05-17T13:55:39.354+08:00",
            "name": "tiancheng92",
            "country": {
              "id": 1,
              "created_at": "2022-05-17T12:50:22.28+08:00",
              "updated_at": "2022-05-17T12:50:22.28+08:00",
              "name": "china",
              "name_cn": "中国",
              "short_name": "cn"
            },
            "email": "kiritoeva@icloud.com"
          }
        ],
        "paginate": {
          "total": 1,
          "page": 1,
          "page_size": 20
        }
      },
      "msg": "Success",
      "code": 100000
    }
    ```
  
* XML:
  ```bash
  # Input
  curl http://127.0.0.1:8080/api/v1/user  -H 'accept: application/xml'
  ```
  ```xml
  # Output
  <result>
      <data>
          <items>
              <id>1</id>
              <created_at>2022-05-17T13:55:39.354+08:00</created_at>
              <updated_at>2022-05-17T13:55:39.354+08:00</updated_at>
              <name>tiancheng92</name>
              <country>
                  <id>1</id>
                  <created_at>2022-05-17T12:50:22.28+08:00</created_at>
                  <updated_at>2022-05-17T12:50:22.28+08:00</updated_at>
                  <name>china</name>
                  <name_cn>中国</name_cn>
                  <short_name>cn</short_name>
              </country>
              <email>kiritoeva@icloud.com</email>
          </items>
          <paginate>
              <total>1</total>
              <page>1</page>
              <page_size>20</page_size>
          </paginate>
      </data>
      <msg>Success</msg>
      <code>100000</code>
  </result>
  ```
* YAML:
  ```bash
  # Input
  curl http://127.0.0.1:8080/api/v1/user  -H 'accept: application/x-yaml'
  ```
  ```yaml
  # Output
  data:
    items:
    - id: 1
      created_at: 2022-05-17T13:55:39.354+08:00
      updated_at: 2022-05-17T13:55:39.354+08:00
      name: tiancheng92
      country:
        id: 1
        created_at: 2022-05-17T12:50:22.28+08:00
        updated_at: 2022-05-17T12:50:22.28+08:00
        name: china
        name_cn: 中国
        short_name: cn
      email: kiritoeva@icloud.com
    paginate:
      total: 1
      page: 1
      page_size: 20
  msg: Success
  code: 100000
  ```

### 默认实现Restful接口
* 入参（request）只需实现[RequestInterface接口](https://github.com/tiancheng92/gin_example_with_generic/blob/main/generic/request.go)即可
* 数据库模型（model）只需实现[ModelInterface接口](https://github.com/tiancheng92/gin_example_with_generic/blob/main/generic/model.go)即可
* model、service、controller层只需New出泛型对象与interface，无需自己实现具体方法，参考country系列接口的实现
* 如有特殊需求，如：model层外键关联、service层特殊业务逻辑等，可以自行对底层泛型方法进行覆盖或添加实现，参考user系列接口的实现