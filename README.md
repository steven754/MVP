# MVP
首先说明下，进行开发此项目，将会按照MVP原则进行，摘选网上对于MVP的介绍

MVP 是 Minimum Viable Product （最小化可行性产品）的简称，如何更好的理解呢？譬如一款电商产品核心目标就是让用户在产品上下单买东西。那核心流程就可能是：进入产品——挑选商品——下单付款——查询物流信息。那就围绕这个流程，剥离多余的高级功能（分享啊，评论啊，个性化推荐啊，积分啊这些都不要做）做一款MVP产品。

所以有些功能将会在前期不会出现在项目中，在后期的话根据需要再决定是否添加。

项目模块：
基于项目的需求，我把项目大致分为以下几个模块

用户模块
文章模块
评论/回复模块
IM聊天模块（websocket）
前三类属于api部分，所以将会在一个程序中
IM聊天模块将会单独做一个程序
这样的好处是在于，如果其中一个程序挂了的话，不会影响另外一个程序

开发顺序：
开发模块
由上图不难看出，评论/回复模块依赖于文章模块，文章模块依赖于用户模块；IM聊天依赖于用户模块
所以开发顺序将会根据上图由左到右，由上到下的顺序来进行。

项目的开发环境和软件
电脑操作系统：macOS Mojava 10.14.4
Golang版本：1.13.4
数据库：MySQL 8.0.16
数据库界面管理：Sequel Pro
开发工具IDE：Goland（很好用的Golang开发工具，只不过是收费的，而且很贵。想白嫖的话，自己查找激活码吧）
接口测试工具：Postman https://www.getpostman.com/

在path路径中，创建项目文件夹，比如为api（下篇文章会介绍项目的创建）

api程序项目的文件夹树：
├── conf
├── middleware
├── models
├── pkg
│   ├── e
│   ├── file
│   ├── logging
│   ├── setting
│   ├── upload
│   └── util
├── routers
│   └── v1
├── runtime
│   ├── logs
│   └── upload
│       ├── apks
│       └── images
│           └── avatar
└── templates
文件夹名称	介绍
conf	项目的配置文件夹
middleware	中间件文件夹
models	模型文件夹
pkg/e	api状态码文件夹
pkg/file	关于文件操作的文件夹
pkg/logging	打印日志操作的文件夹
pkg/setting	项目设置文件夹
pkg/upload	上传文件/图片操作的文件夹
pkg/util	工具类文件夹
routers	路由文件夹
runtime	存放文件/资源文件夹
runtime/logs	存放项目打印的日志
runtime/upload	存放上传的资源文件夹
runtime/upload/apks	存放安卓更新的安装包，都是apk文件
runtime/upload/images	存放项目的图片，比如用户的头像，用户上传的图片都在这里
templates	静态资源文件夹，比如web页面等
