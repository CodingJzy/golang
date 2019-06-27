package main

import (
	"fmt"
)

func main()  {
	fmt.Printf("go web 框架之echo安装")

	install := `
1、使用go get
	cd <PROJECT IN $GOPATH>
	go get -u github.com/labstack/echo/...
                    
2、使用dep
	cd <PROJECT IN $GOPATH>
	dep ensure -add github.com/labstack/echo@^3.1

3、使用glide
	cd <PROJECT IN $GOPATH>
	glide get github.com/labstack/echo#~3.1

4、使用govender
	cd <PROJECT IN $GOPATH>
	govendor fetch github.com/labstack/echo@v3.1

注意：
	安装过程可能会报错，少了一些包
	可以手动下载
	cd <PROJECT IN $GOPATH>/src/golang.org/x
	git clone git@github.com:golang/crypto.git
	git clone git@github.com:golang/sys.git
	git clone git@github.com:golang/net.git
	git clone git@github.com:golang/text.git
	` 
	fmt.Println(install)

}