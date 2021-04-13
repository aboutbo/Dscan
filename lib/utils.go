package lib

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFromFile 读文件
func ReadFromFile(filename string) []string {
	var fileContents []string
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file fail", err)
		return nil
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		fileContents = append(fileContents, fileScanner.Text())
	}

	return fileContents
}

// CheckTargetAlive 检查目标是否存活
func CheckTargetAlive(url string) bool {
	_, err := RequestURL(url)
	if err != nil {
		return false
	}
	return true
}
