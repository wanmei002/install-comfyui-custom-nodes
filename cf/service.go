package cf

import (
	"encoding/json"
	"fmt"
	"gitee.com/wanmei002/install-comfyui-custom-node/utils"
	"strings"
)

type Service struct {
	customNodes  map[string]string
	basePath     string
	repoListPath string
}

func NewService(basePath string, repoListPath string) (*Service, error) {
	svc := &Service{
		basePath:     basePath,
		repoListPath: repoListPath,
	}
	err := svc.customNodeList()
	if err != nil {
		return nil, err
	}
	l, err := svc.repoList()
	if err != nil {
		return nil, err
	}
	if len(l) < 1 {
		return nil, fmt.Errorf("repo list is empty")
	}
	return svc, nil
}

type customNodes struct {
	CustomNodes []struct {
		Author      string   `json:"author"`
		Title       string   `json:"title"`
		ID          string   `json:"id"`
		Reference   string   `json:"reference"`
		Files       []string `json:"files"`
		InstallType string   `json:"install_type"`
		Description string   `json:"description"`
	} `json:"custom_nodes"`
}

func (s *Service) customNodeList() error {
	s.customNodes = make(map[string]string)
	list := customNodes{}
	err := json.Unmarshal([]byte(CustomNodeList), &list)
	if err != nil {
		fmt.Println("custom node list json unmarshal err: ", err)
		return err
	}
	utils.Debug("custom node list len: ", len(list.CustomNodes))

	for _, v := range list.CustomNodes {
		if v.InstallType != "git-clone" {
			continue
		}
		tmpList := strings.Split(v.Reference, "/")
		gitRepoName := tmpList[len(tmpList)-1]
		s.customNodes[gitRepoName] = v.Reference
	}
	if len(s.customNodes) < 1 {
		return fmt.Errorf("custom node list is empty")
	}
	return nil
}
