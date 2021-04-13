# Dscan

**特点**

* 基于规则的目录扫描工具，实现精准扫描
* 适用于漏洞自动化巡检场景

**安装**
```
go get -u github.com/aboutbo/Dscan
```


**使用**
```
Usage:
  Dscan [command]

Available Commands:
  help        Help about any command
  scan        Use accurate model or fuzz model to scan

Flags:
      --accurate      Whether to use accurate mode
  -f, --file string   target file to scan
  -h, --help          help for Dscan
      --rule string   rule file (default "rules/rules.yaml")

Use "Dscan [command] --help" for more information about a command.
```

`./Dscan scan --accurate -f your_target_urls_file --rule your_rules_file`

**规则编写**

规则支持函数如下：
* response.body.bcontains(b'your strings') ： 判断HTTP响应body是否包含某字符串
* response.status_code ：`response.status_code==200`
* response.headers['your header key'].contains('your header value') ：判断HTTP响应headers某key是否包含某value
* response.headers.contains_key('your header key') ：判断HTTP响应headers是否存在某key

example:
```
  # the number of this rule
  rule5: 
    # rule description
    description: kibana
    # rule path
    path: /app/kibana
    # rule expression
    expression: response.status_code == 200 && response.body.bcontains(b'kibanaWelcomeLogo')
```

**TODO**
* 并发扫描
* fuzz模式
* 结果输出
* 更丰富的规则库