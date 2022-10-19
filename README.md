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
  ```