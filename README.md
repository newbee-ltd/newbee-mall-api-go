![](static-files/newbee-mall.png)

### newbee-mall 项目的 go 语言版本

本项目为[新蜂商城后端接口 newbee-mall-api](https://github.com/newbee-ltd/newbee-mall-api) 的 go 语言版本，当前项目中的代码主要由 [@十三](https://github.com/newbee-mall)和 [@可乐](https://github.com/dalaohekele) 共同开发。newbee-mall 项目是一套电商系统，基于 Spring Boot 2.X 及相关技术栈开发。本项目采用了原版本的所有数据结构，技术栈为 Go + Gin，主要面向服务端开发人员，前端 Vue 页面源码在另外三个 Vue 仓库。

前端项目：

- [新蜂商城 Vue2 版本 newbee-mall-vue-app](https://github.com/newbee-ltd/newbee-mall-vue-app)
- [新蜂商城 Vue3 版本 newbee-mall-vue3-app](https://github.com/newbee-ltd/newbee-mall-vue3-app)
- [新蜂商城后台管理系统 Vue3 版本 vue3-admin](https://github.com/newbee-ltd/vue3-admin)

**如果觉得项目还不错的话可以给项目一个 Star 吧。**

## 联系作者

> 大家有任何问题或者建议都可以在 [issues](https://github.com/newbee-ltd/newbee-mall-api-go/issues) 中反馈给我。

### 项目讲解

- [【go商城】gin+gorm实现CRUD](https://blog.csdn.net/zxc19854/article/details/125267635)
- [【go商城】gin+mysql实现token登陆校验](https://blog.csdn.net/zxc19854/article/details/125352067)
- [【go商城】gin+vue跨域问题](https://blog.csdn.net/zxc19854/article/details/125464151)

### 本地启动

#### 后端项目启动

首先导入 static-files 中的 sql 文件。

```bash
# 克隆项目
git clone https://github.com/newbee-ltd/newbee-mall-api-go

# 使用 go mod 并安装go依赖包
go generate
# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )
# 运行二进制
./server (windows运行命令为 server.exe)
```

#### 前端项目启动

然后按照原项目的部署说明部署即可。

[后台管理项目](https://github.com/newbee-ltd/vue3-admin)

测试用户名：admin  测试密码：123456


[前台商城](https://github.com/newbee-ltd/newbee-mall-vue3-app)

直接注册账号就可以了。

## 页面展示

以下为新蜂商城 Vue 版本的页面预览：

- 登录页

![](static-files/登录.png)

- 首页

![](static-files/首页.png)

- 商品搜索

![](static-files/商品搜索.png)

- 商品详情页

![](static-files/详情页.png)

- 购物车

![](static-files/购物车.png)

- 生成订单

![](static-files/生成订单.png)

- 地址管理

![](static-files/地址管理.png)

- 订单列表

![](static-files/订单列表.png)

- 订单详情

![](static-files/订单详情.png)

## 感谢

- [newbee-ltd](https://github.com/newbee-ltd)

- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)
