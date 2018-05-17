//Package action 定义了初始化和操作网络的方法
package action

import (
	"fmt"
	"os/exec"
	"strconv"
)

//初始化账户
func initCC(name string) (Response, error) {
	cmd := "bash initCC.sh"+" "+name
	outAsBytes, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//进行交易
func invokeCC(name string, opName string, opAmount int) (Response, error) {
	cmd := "bash invokeCC.sh"  + " " + name + " " + opName + " " + strconv.Itoa(opAmount)
	outAsBytes, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//查询账户信息
func queryCC(function string,opName string) (Response, error) {
	cmd := "bash queryCC.sh"  + " " + function + " " + opName
	outAsBytes, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}
