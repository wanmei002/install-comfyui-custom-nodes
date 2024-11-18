package cf

import (
	"gitee.com/wanmei002/install-comfyui-custom-node/utils"
	"io"
	"os"
	"strings"
)

func (s *Service) repoList() ([]string, error) {
	fileInfo, err := os.Open(s.repoListPath)
	if err != nil {
		utils.Debug("open file error:", s.repoListPath, err)
		return nil, err
	}
	fileContent, err := io.ReadAll(fileInfo)
	if err != nil {
		utils.Debug("read file content error:", s.repoListPath, err)
		return nil, err
	}
	repoList := strings.Split(string(fileContent), "\n")
	ret := []string{}
	for _, repo := range repoList {
		tmp := strings.TrimSpace(repo)
		if len(tmp) == 0 {
			continue
		}
		ret = append(ret, tmp)
	}
	utils.Debug("repoList len :", len(ret))
	return repoList, nil
}
