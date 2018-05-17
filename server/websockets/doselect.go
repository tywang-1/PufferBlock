//Package websockets ...
package websockets

import "myrepo/PufferBlock/server/action"

//分类解析请求
func (req *Request) doSelect() (action.Response, error) {
	//缺少类型
	if req.Type == "" {
		return action.Response{IfSuccessful: false, ErrInfo: "no such type", Result: ""}, nil
	}
	//初始化账户
	if req.Type == "initCC" {
		return action.InitCC(req.Name)
	}
	//查询账户信息
	if req.Type == "queryCC" {
		return action.QueryCC(req.Function,req.OpName)
	}
	//进行交易
	if req.Type == "invokeCC" {

		//拒接非法交易
		if req.Name != req.OpName {
			return action.Response{IfSuccessful: false, ErrInfo: "denied", Result: ""}, nil
		}
		return action.InvokeCC(req.Function, req.OpName, req.OpAmount)
	}
	//类型错误
	return action.Response{IfSuccessful: false, ErrInfo: "wrong Type", Result: ""}, nil
}
