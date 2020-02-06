# IM BackEnd Package Control Design Doc

**Created Date: **2019.10.13

**Last Updated:** 2020.02.06

## Goal

We use `dep` to manage packages in our project.

## dep

One of the most attracting dependency management tool for Go. 

---

## go mod

A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module's module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.

---

## govendor

check-in the `vendor/vendor/json` file.

## Solution

Due to connection problems and reliability, we choose go mod as our tool for package control.

1.  `export GO111MODULE=on` on Linux or `set GO111MODULE=on` on Windows.
2. set proxy. We use [goproxy](https://goproxy.io/). `export GOPROXY=https://goproxy.io` on Linux or `set GOPROXY=https://goproxy.io` on Windows.

## Reference

1. [Golang Package Control - dep](https://blog.csdn.net/huwh_/article/details/81170900)
2. [Dep](https://www.jianshu.com/p/b02d9edb9ad0)
3. [dep -- Github source code](https://github.com/golang/dep)
4. [go vendor](https://www.cnblogs.com/liuzhongchao/p/9233177.html)
5. [go vendor -- Github source code](https://github.com/kardianos/govendor)
6. [go mod Usage1](https://segmentfault.com/a/1190000019314903)
7. [go mod Usage2](https://blog.csdn.net/jinyidong/article/details/87117892)
8. [go mod docs](https://blog.golang.org/using-go-modules)
9. [goproxy](https://goproxy.io/)