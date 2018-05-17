//Package action 定义了初始化和操作网络的方法
package action

import (
	"fmt"
	"os/exec"
)

//生成配置文件
func generate() {
	cmd := "make generate"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	fmt.Println(string(out))
}

//启动网络
func networkUp() {
	cmd := "make networkup"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	fmt.Println(string(out))
}
