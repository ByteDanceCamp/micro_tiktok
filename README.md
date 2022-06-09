# micro_tiktok


## 一、项目简介
Micro_tiktok 是基于Kitex开发的一套极简版抖音API，包含用户、视频、评论、点赞、关系五个模块。
- 服务地址：http://mtt.dtpark.top/
- Golang 版本：1.16
- 说明文档地址：https://t5pn33186b.feishu.cn/docs/doccn5qygcgKFjejI3qDIj51Ukg
- apk: https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7
- API文档：https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

> 【注】安全起见，仓库中的七牛云密钥对已停用，加密用盐等信息也与服务器中不同

## 二、项目结构
![项目架构](https://s2.loli.net/2022/06/10/8xsHTzolRGW2Lfd.png)

## 三、数据库整体设计
![E-R 图](https://s2.loli.net/2022/06/10/NwuG4cJSFPQRHO2.jpg)

## 四、目录结构

### 4.1 整体结构
```bash
├── cmd                    # 应用目录
│   ├── api                # API 服务（Gin 实现的 http 接口，需要各自补充）
│   ├── comment            # 评论模块
│   ├── favorite           # 点赞模块
│   ├── relation           # 关系模块（关注、粉丝）
│   ├── user               # 用户模块
│   └── video              # 视频模块
             # 其他模块
├── docker-compose.yml     # docker compose 配置文件，包括Mysql、Etcd、Jarger、redis
├── go.mod                 # go 依赖管理
├── go.sum
├── idl                    # RPC 接口定义（可根据情况自行修改数据类型等）
│   ├── comment.proto      # 评论接口
│   ├── favorite.proto     # 点赞接口
│   ├── relation.proto     # 关系接口
│   ├── user.proto         # 用户接口
│   └── video.proto        # 视频接口
├── kitex_gen              # ketex 自动生成目录
│   ├── comment
│   ├── favorite
│   ├── relation
│   ├── user
│   └── video

└── pkg                    # 项目公共使用的代码
    ├── constants
    │   └── constant.go    # 项目中所有常量的定义
    ├── errno
    │   └── errno.go       # 项目自定义的 Error 状态码和信息
    └── tracer
        └── tracer.go      # Jarger 的初始化
```

### 4.2 API模块（Gin）结构

```bash
.
├── handlers                            # 各 API 的 handler
│   ├── comment_action.go               # 评论操作（评论/删除评论）
│   ├── comment_list.go                 # 获取评论列表
│   ├── common.go                       # 各 handler 需要使用的公共代码
│   ├── favorite_action.go              # 点赞/取消点赞
│   ├── favorite_video_list.go          # 喜欢列表
│   ├── feed.go                         # 视频流
│   ├── login.go                        # 登录
│   ├── register.go                     # 注册
│   ├── relation_action.go              # 关注/取消关注
│   ├── relation_follow_list.go         # 关注列表
│   ├── relation_follower_list.go       # 粉丝列表
│   ├── user_info.go                    # 用户信息
│   ├── video_list.go                   # 发布列表
│   └── video_publish.go                # 发布视频
├── main.go
├── middleware                          # 中间件
│   └── auth                            # 鉴权中间件
│       ├── auth.go                     # 继承 appleboy/gin-jwt 插件
│       ├── config.go                   # 实例化 jwt 时的配置
│       ├── form_auth.go                # 针对 gin-jwt 无法获取 form-data 中 token 做的扩展
│       └── select.go                   # 对 登录/未登录 用户视频流分别处理对中间件
├── router                              # 路由
│   └── router.go
├── rpc                                 # 封装其他微服务的 client 接口
│   ├── comment.go
│   ├── favorite.go
│   ├── init.go
│   ├── relation.go
│   ├── user.go
│   └── video.go
├── run.sh                              # 启动脚本
└── utils                               # 工具函数
    ├── upload.go                       # 上传视频方法
    └── validate_file.go                # 验证待上传文件类型&生成新的文件名
```

### 4.3 微服务目录结构（以关系模块为例）
```bash
.
├── Makefile                    # 根据 idl 构建微服务脚本
├── build.sh                    # 构建代码
├── dal                         # 数据层
│   ├── db                      # 数据库初始化及方法
│   │   ├── base.go             # 公共代码
│   │   ├── follow.go           # 关注表
│   │   ├── follower.go         # 粉丝表
│   │   ├── init.go             # 初始化
│   │   └── relation_count.go   # 关注&粉丝 数量表
│   ├── init.go
│   └── redis                  # Redis 初始化及操作方法
│       ├── action.go          # 关注/取关
│       ├── init.go            # 初始化 
├── handler.go                 # 视图层
├── main.go
├── pack                       # 数据格式转化
│   ├── resp.go                # 返回响应构建、转化
│   ├── user.go                # db 格式的 User 转化为 idl 定义的格式
│   └── utils.go               # 工具函数
├── rpc                        # 封装的调用其他 rpc 的客户端方法
│   ├── init.go
│   └── user.go
├── script                     # kitex 自动生成
│   └── bootstrap.sh
└── service                    # 逻辑层
    ├── action.go              # 关注/取关
    ├── count.go               # 关注&粉丝 总数
    └── list.go                # 关注/粉丝 列表

```

## 五、本地调试方法
1. 配置 七牛云 存储
```go
// 将 pkg/constants/constant.go 中下述七牛云相关参数替换为自己等
QiNiuAccessKey = "" 
QiNiuSecretKey = ""
QiNiuBucket    = ""
QiNiuServer    = ""
```
2. 启动 docker 环境
```bash
docker-compose up -d
```
2. 启动 API 服务
```bash 
cd cmd/api
sh run.sh
```
2. 启动其他服务
```bash
cd ../user
sh build.sh
sh output/bootstrap.sh

cd ../video
sh build.sh
sh output/bootstrap.sh

cd ../comment
sh build.sh
sh output/bootstrap.sh

cd ../favorite
sh build.sh
sh output/bootstrap.sh

cd ../relation
sh build.sh
sh output/bootstrap.sh
```