package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

// Rules 规则集合
type Rules struct {
	RulesSet map[string]*Rule `yaml:"rules"`
}

// Rule 单条规则
type Rule struct {
	Description string `yaml:"description"`
	Path        string `yaml:"path"`
	Expression  string `yaml:"expression"`
}

// LoadRules 从文件读取Rule内容
func LoadRules(fileName string) (*Rules, error) {
	rules := &Rules{}
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, rules)
	return rules, err
}

// CheckRule 检查http响应是否符合当前规则
func CheckRule(target string, rules *Rules) {
	c := NewEnvOption()
	env, err := NewEnv(&c)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 遍历规则
	for _, v := range rules.RulesSet {
		vars := make(map[string]interface{})
		// 拼接Url，并请求
		path := v.Path
		url := strings.Trim(target, "/") + path
		oResp, err := RequestURL(url)
		if err != nil {
			continue
		}
		resp, err := parseResponse(oResp)
		vars["response"] = resp
		if err != nil {
			fmt.Println(err)
		}

		// 判断是否符合规则
		out, err := Evaluate(env, v.Expression, vars)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if out.Value() == true {
			fmt.Printf("Found: %s\n", url)
		}
	}

}

func parseResponse(oResp *http.Response) (*Response, error) {
	var resp Response
	headers := make(map[string]string)
	resp.StatusCode = int32(oResp.StatusCode)
	for key := range oResp.Header {
		headers[key] = oResp.Header.Get(key)
	}
	resp.Headers = headers
	resp.ContentType = oResp.Header.Get("Content-Type")
	body, err := getRespBody(oResp)
	if err != nil {
		return nil, err
	}
	resp.Body = body
	return &resp, nil
}
