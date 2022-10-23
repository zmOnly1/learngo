```shell
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```
### GOPATH、GOVENDOR、GO MODE
- GOPATH
  ```shell
  mkdir -p $GOPATH/src/project1
  mkdir -p $GOPATH/src/project2
  
  cd $GOPATH/src/project1
  go get -u github.com/abc
  $GOPATH/src/github.com/abc
  ```
- GOVENDOR
  ```shell
  mkdir -p $GOPATH/src/project1/vendor
    
  go get -u github.com/abc
  $GOPATH/src/project1/vendor/github.com/abc
  ```
- GO MOD
  ```shell
  go get -u github.com/abc@v1.1.1
  
  $GOPATH/pkg/mod/github.com/abc@v1.1.1
  
  go mod init <module  name>
  go mod tidy
  go build ./...
  go install ./... => $GOPATH/bin
  ```
### 目录
1. 常量枚举：const
2. 条件、循环：if exp {}, for {}
3. 指针：值传递，引用传递
4. 数组，切片，Map
5. 结构体组合，内嵌
6. 包管理：GOPATH,GOVENDOR,go mod
7. 接口，duck typing，接口组合
8. 函数式编程，闭包
9. defer,错误处理，panic,recover
10. http
11. 测试,覆盖测试和性能测试; 
```shell 
#代码覆盖率
go test .
go test -coverprofile c.out 
go tool cover -html=c.out
#性能测试
go test -bench .
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out
(pprof) help
(pprof) web
```