# micro_tiktok


## 项目简介
Micro_tiktok 是基于Kitex开发的一套极简版抖音API，包含用户、视频、评论、点赞、关系五个模块。
- 项目地址：https://github.com/ByteDanceCamp/micro_tiktok
- Golang 版本：1.16
- 说明文档地址：https://t5pn33186b.feishu.cn/docs/doccn5qygcgKFjejI3qDIj51Ukg
- apk: https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7

> 【注】安全起见，仓库中的七牛云密钥对已停用，加密用盐等信息也与服务器中不同

## 本地调试方法
1. 启动 docker 环境
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