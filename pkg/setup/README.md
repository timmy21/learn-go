## 下载安装

下载与安装：https://go.dev/doc/install
安装多个版本：https://go.dev/doc/manage-install

## 设置代理

在使用 go 的过程中，经常需要访问 github/google 等网站，进行包下载。使用默认的代理，下载速度很慢，甚至无法下载。可以通过设置国内的代理来解决。
```
go env -w GOPROXY="https://goproxy.cn,direct"
```

设置私有仓库
```
go env -w GOPRIVATE="*.mycompany.com"
```

通过ssh获取私有仓库的包
```
git config --global url."ssh://git@git.mycompany.com/".insteadOf https://git.mycompany.com/scm/
```

## 初始化项目

```
mkdir learn-go
cd learn-go
go mod init github.com/timmy21/learn-go
```