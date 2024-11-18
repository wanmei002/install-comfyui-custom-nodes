package cf

import (
	"bytes"
	"fmt"
	"gitee.com/wanmei002/install-comfyui-custom-node/utils"
	"github.com/go-git/go-git/v5"
	"io"
	"os"
	"path"
)

func (s *Service) GitClone() error {
	list, err := s.repoList()
	if err != nil {
		return err
	}
	for _, repoName := range list {
		repoAddress, ok := s.customNodes[repoName]
		if !ok {
			utils.Debug(fmt.Sprintf("%s not found in custom nodes", repoName))
			continue
		}
		filePath := path.Join(s.basePath, repoName)
		_, err := os.Stat(filePath)
		if err == nil {
			utils.Debug(fmt.Sprintf("%s already exists", repoName))
			continue
		}
		if !os.IsNotExist(err) {
			s.printError(err)
			continue
		}
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			s.printError(err)
			continue
		}
		utils.Debug(fmt.Sprintf("start git clone %s", repoAddress))
		_, err = git.PlainClone(filePath, false, &git.CloneOptions{
			URL: fmt.Sprintf("%s%s", "https://mirror.ghproxy.com/", repoAddress),
		})
		if err != nil {
			s.printError(err)
			os.RemoveAll(filePath)
		} else {
			fmt.Printf("git clone %s success!\n", repoAddress)
		}
	}

	return nil
}

func (s *Service) printError(strs ...any) {
	strs = append(strs, "\n")
	val := fmt.Sprint(strs...)
	io.Copy(os.Stderr, bytes.NewBufferString(val))
}

// %s/"https:\/\/github.com/"https:\/\/mirror.ghproxy.com\/https:\/\/github.com/g
