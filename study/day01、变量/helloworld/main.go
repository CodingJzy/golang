package main

import (
	"fmt"
)

func install() {
	fmt.Println(`
	在自己的gopath下操作
	1、mkdir -p ./src/github.com
	2、cd ./src/giuthub.com/
	3、mkdir acroca cweill derekparker go-delve josharian karrick mdempsky pkg ramya-rao-a rogpeppe sqs uudashr stamblerre
	4、分别进去到每个创建的文件夹clone
		git clone https://github.com/acroca/go-symbols.git
		git clone https://github.com/cweill/gotests.git
		git clone https://github.com/derekparker/delve.git
		git clone https://github.com/go-delve/delve.git
		git clone https://github.com/josharian/impl.git
		git clone https://github.com/karrick/godirwalk.git
		git clone https://github.com/mdempsky/gocode.git
		git clone https://github.com/pkg/errors.git
		git clone https://github.com/ramya-rao-a/go-outline.git
		git clone https://github.com/rogpeppe/godef.git
		git clone https://github.com/sqs/goreturns.git
		git clone https://github.com/uudashr/gopkgs.git
		git clone https://github.com/stamblerre/gocode.git
	5、 mkdir -p ./src/golang.org/x && cd ./src/golang.org/x
	6、 git clone https://github.com/golang/tools.git
	7、 git clone https://github.com/golang/lint.git
	8、 cd $GOPATH
	9、 手动安装插件
		go install github.com/mdempsky/gocode
		go install github.com/uudashr/gopkgs/cmd/gopkgs
		go install github.com/ramya-rao-a/go-outline
		go install github.com/acroca/go-symbols
		go install github.com/rogpeppe/godef
		go install github.com/sqs/goreturns
		go install github.com/derekparker/delve/cmd/dlv
		go install github.com/cweill/gotests
		go install github.com/josharian/impl
		go install golang.org/x/tools/cmd/guru
		go install golang.org/x/tools/cmd/gorename
		go install golang.org/x/lint/golint
	`)
}

func main() {
	fmt.Println("Hello Go")
	install()
}
