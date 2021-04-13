package lib

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Scan starts from here
func Scan(cmd *cobra.Command, args []string) {
	// 获取参数
	accurateModel, _ := cmd.Flags().GetBool("accurate")
	targetFile, _ := cmd.Flags().GetString("file")
	target, _ := cmd.Flags().GetString("target")
	ruleFile, _ := cmd.Flags().GetString("rule")

	// 目标集合
	var targetSet []string
	switch {
	case target != "":
		targetSet = append(targetSet, target)
	case targetFile != "":
		targetSet = ReadFromFile(targetFile)
	default:
		fmt.Println("No targets!")
		return
	}

	if accurateModel {
		accurateScan(&targetSet, ruleFile)
	}
}

// 精确扫描
func accurateScan(targetSet *[]string, ruleFile string) {
	rules, err := LoadRules(ruleFile)
	if err != nil {
		fmt.Println(err)
	}
	for _, target := range *targetSet {
		// url存活性检查
		if !CheckTargetAlive(target) {
			continue
		}
		CheckRule(target, rules)
	}

}
