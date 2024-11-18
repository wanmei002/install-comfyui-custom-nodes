package main

import (
	"flag"
	"fmt"
	"gitee.com/wanmei002/install-comfyui-custom-node/cf"
)

var path string
var repoListPath string
var help bool

var helpInfo = `
--repo:  repo list file path
--path:  git clone base path
--help:  display help
`

func main() {
	flag.StringVar(&path, "path", "", "path")
	flag.StringVar(&repoListPath, "repo", "", "repo list file path")
	flag.BoolVar(&help, "help", false, "help")
	flag.Parse()

	if help {
		fmt.Println(helpInfo)
		return
	}

	if repoListPath == "" {
		fmt.Println(helpInfo)
		panic("repo list file path is empty")
	}
	if path == "" {
		fmt.Println(helpInfo)
		panic("git clone base path is empty")
	}
	svc, err := cf.NewService(path, repoListPath)
	if err != nil {
		panic(err)
	}
	err = svc.GitClone()
	if err != nil {
		panic(err)
	}
	fmt.Println("success")
}
