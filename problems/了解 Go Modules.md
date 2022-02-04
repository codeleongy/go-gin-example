# 了解 Go Modules

## 

## 了解历史

在过去，Go 的依赖包管理在工具上混乱且不统一，有 dep，有 glide，有 govendor…甚至还有因为外网的问题，频频导致拉不下来包，很多人苦不堪言，盼着官方给出一个大一统做出表率。

而在 Go modules 正式出来之前还有一个叫 dep 的项目，我们在上面有提到，它是 Go 的一个官方实验性项目，目的也是为了解决 Go 在依赖管理方面的问题，当时社区里面几乎所有的人都认为 dep 肯定就是未来 Go 官方的依赖管理解决方案了。

但是万万没想到，半路杀出个程咬金，Russ Cox 义无反顾地推出了 Go modules，这瞬间导致一石激起千层浪，让社区炸了锅。大家一致认为 Go team 实在是太霸道、太独裁了，连个招呼都不打一声。我记得当时有很多人在网上跟 Russ Cox 口水战，各种依赖管理解决方案的专家都冒出来发表意见，讨论范围甚至一度超出了 Go 语言的圈子触及到了其他语言的领域。

当然，最后，推成功了，Go modules 已经进入官方工具链中，与 Go 深深结合，以前常说的 GOPATH 终将会失去它原有的作用，而且它还提供了 GOPROXY 间接解决了国内访问外网的问题。

## 了解 Russ Cox

在上文中提到的 Russ Cox 是谁呢，他是 Go 这个项目目前代码提交量最多的人，甚至是第二名的两倍还要多（从 2019 年 09 月 30 日前来看）。

Russ Cox 还是 Go 现在的掌舵人（大家应该知道之前 Go 的掌舵人是 Rob Pike，但是听说由于他本人不喜欢特朗普执政所以离开了美国，然后他岁数也挺大的了，所以也正在逐渐交权，不过现在还是在参与 Go 的发展）。

Russ Cox 的个人能力相当强，看问题的角度也很独特，这也就是为什么他刚一提出 Go modules 的概念就能引起那么大范围的响应。虽然是被强推的，但事实也证明当下的 Go modules 表现得确实很优秀，所以这表明一定程度上的 “独裁” 还是可以接受的，至少可以保证一个项目能更加专一地朝着一个方向发展。

## 配置

Go自从Go1.11开始就正式发布了Go Modules，但仍保留GOPATH模式，因此需要手动设置

```bash
go env

go env -w GO111MODULE=on

go env -w GOPROXY=https://goproxy.cn,direct
```

- `go env`：查看当前go配置信息

- `go env -w GO111MODULE=xxx`：配置Go Modules开启，总共有三个参数 `on`、`auto`、`off`，分别代表开启，自动选择和关闭`Go Module`

- `go env -w GOPROXY=xxx`：配置代理（提高依赖下载速度，防止速度过慢）

  常见的国内代理包括：

  1. 阿里云：https://mirrors.aliyun.com/goproxy/
  2. 官方：https://goproxy.io/
  3. 七牛云：https://goproxy.cn,direct

## 基础操作

```bash
go mod init github.com/leong-y/go-gin-example

go get -u -v github.com/gin-gonic/gin

```

- `go mod init [MODULE_PATH]`：初始化一个`GoModule`项目，会生成一个`go.mod`管理项目依赖
- 用`go get`拉取新的依赖
  - 拉取最新的版本(优先择取 tag)：`go get golang.org/x/text@latest`
  - 拉取 `master` 分支的最新 commit：`go get golang.org/x/text@master`
  - 拉取 tag 为 v0.3.2 的 commit：`go get golang.org/x/text@v0.3.2`
  - 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：`go get golang.org/x/text@342b2e`
  - 用 `go get -u` 更新现有的依赖
  - 用 `go mod download` 下载 go.mod 文件中指明的所有依赖
  - 用 `go mod tidy` 整理现有的依赖
  - 用 `go mod graph` 查看现有的依赖结构
  - 用 `go mod init` 生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)
- 用 `go mod edit` 编辑 go.mod 文件
- 用 `go mod vendor` 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)
- 用 `go mod verify` 校验一个模块是否被篡改过



参考自： [煎鱼大佬的博客](https://eddycjy.com/posts/go/gin/2018-02-10-install/)

​		[GoModule官方文档](https://github.com/golang/go/wiki/Modules)

