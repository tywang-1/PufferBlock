//Package main 碳交易链码
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func main() {
	err := shim.Start(new(CarbonCC))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

//CarbonCC 链码结构
type CarbonCC struct {
}

//CarbonInfo 账户信息结构
type CarbonInfo struct {
	Market string `json:"market"`
	Amount int    `json:"amount"`
}

//Init 链码初始化接口
func (c *CarbonCC) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

//Invoke 链码操作接口
func (c *CarbonCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()
	if function == "createCarbonInfo" {
		//创建账户
		return c.createCarbonInfo(stub, args)
	} else if function == "queryAllCarbonInfo" {
		//查询全部账户信息
		return c.queryAllCarbonInfo(stub)
	} else if function == "updateCarbonInfo" {
		//更新账户信息
		return c.updateCarbonInfo(stub, args)
	} else if function == "queryByOwner" {
		//查询指定账户信息
		return c.queryByOwner(stub, args)
	} else if function == "queryByMarket" {
		//查询指定类型账户信息
		return c.queryByMarket(stub, args)
	} else if function == "queryByAmount" {
		//查询指定额度账户信息
		return c.queryByAmount(stub, args)
	} else if function == "transfer" {
		//进行交易
		return c.transfer(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (c *CarbonCC) createCarbonInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	owner := args[0]
	market := args[1]
	amount, _ := strconv.Atoi(args[2])
	carbonInfo := &CarbonInfo{market, amount}

	carbonInfoCheckAsBytes, err := stub.GetState(owner)
	if err != nil {
		return shim.Error("Failed to get state.")
	} else if carbonInfoCheckAsBytes != nil {
		return shim.Error("Account already exists.")
	}

	carbonInfoAsBytes, err := json.Marshal(carbonInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(owner, carbonInfoAsBytes)

	//rich.comKeys{Market, amount}

	return shim.Success(nil)
}

func (c *CarbonCC) queryAllCarbonInfo(stub shim.ChaincodeStubInterface) peer.Response {

	startKey := "a"
	endKey := "zzzzzzzzzz"

	allCarbonInfoAsBytes, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer allCarbonInfoAsBytes.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	writtenFlag := false
	for allCarbonInfoAsBytes.HasNext() {
		queryResponse, err := allCarbonInfoAsBytes.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if writtenFlag == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Owner\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"CarbonInfo\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		writtenFlag = true
	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (c *CarbonCC) updateCarbonInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	owner := args[0]
	market := args[1]
	amount, _ := strconv.Atoi(args[2])
	carbonInfo := &CarbonInfo{market, amount}

	carbonInfoCheckAsBytes, err := stub.GetState(owner)
	if err != nil {
		return shim.Error("Failed to get state.")
	} else if carbonInfoCheckAsBytes == nil {
		return shim.Error("Account doesn't exist.")
	}

	carbonInfoAsBytes, err := json.Marshal(carbonInfo)
	if err != nil {
		return shim.Error(err.Error())
	}
	stub.PutState(owner, carbonInfoAsBytes)
	return shim.Success(nil)
}

func (c *CarbonCC) queryByOwner(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("")
	}

	owner := args[0]
	carbonInfoAsBytes, err := stub.GetState(owner)
	if err != nil {
		return shim.Error("Failed to get state.")
	} else if carbonInfoAsBytes == nil {
		return shim.Error("Account doesn't exist.")
	}
	return shim.Success(carbonInfoAsBytes)
}

func (c *CarbonCC) queryByMarket(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (c *CarbonCC) queryByAmount(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (c *CarbonCC) transfer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	transferor := args[0]
	transferee := args[1]
	transferAmount, _ := strconv.Atoi(args[2])

	carbonInfoTransfereeCheckAsBytes, err := stub.GetState(transferee)
	if err != nil {
		return shim.Error("Failed to get state.")
	} else if carbonInfoTransfereeCheckAsBytes == nil {
		return shim.Error("Transferee doesn't exist.")
	}

	carbonInfoTransferorCheckAsBytes, err := stub.GetState(transferor)
	if err != nil {
		return shim.Error("Failed to get state.")
	} else if carbonInfoTransferorCheckAsBytes == nil {
		return shim.Error("Transferor doesn't exist.")
	}

	carbonInfoTransferorCheck := &CarbonInfo{}
	json.Unmarshal(carbonInfoTransferorCheckAsBytes, carbonInfoTransferorCheck)
	if carbonInfoTransferorCheck.Amount < transferAmount {
		return shim.Error("Transferor Shortage.")
	}
	carbonInfoTransferor := &CarbonInfo{carbonInfoTransferorCheck.Market, carbonInfoTransferorCheck.Amount - transferAmount}
	carbonInfoTransferorAsBytes, _ := json.Marshal(carbonInfoTransferor)
	stub.PutState(transferor, carbonInfoTransferorAsBytes)

	carbonInfoTransfereeCheck := &CarbonInfo{}
	json.Unmarshal(carbonInfoTransfereeCheckAsBytes, carbonInfoTransfereeCheck)
	carbonInfoTransferee := &CarbonInfo{carbonInfoTransfereeCheck.Market, carbonInfoTransfereeCheck.Amount + transferAmount}
	carbonInfoTransfereeAsBytes, _ := json.Marshal(carbonInfoTransferee)
	stub.PutState(transferee, carbonInfoTransfereeAsBytes)

	return shim.Success(nil)
}