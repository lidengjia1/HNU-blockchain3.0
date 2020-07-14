// hanxiaodong
// QQ群（专业Fabric交流群）：862733552
package main

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "fmt"
    "github.com/hyperledger/fabric/protos/peer"
    "encoding/json"
    "bytes"
)

type CouchDBChaincode struct {

}

func (t *CouchDBChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response  {
    return shim.Success(nil)
}

func (t *CouchDBChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response  {
    fun, args := stub.GetFunctionAndParameters()
    if fun == "billInit" {
        return billInit(stub, args)
    } else if fun == "queryBills" {
        return queryBills(stub, args)
    } else if fun == "queryWaitBills" {
        return queryWaitBills(stub, args)
    }

    return shim.Error("非法操作, 指定的函数名无效")
}

// 初始化票据数据
func billInit(stub shim.ChaincodeStubInterface, args []string) peer.Response  {
    bill := BillStruct{
        ObjectType:"billObj",
        BillInfoID:"POC101",
        BillInfoAmt:"1000",
        BillInfoType:"111",
        BillIsseDate:"20100101",
        BillDueDate:"20100110",

        HolderAcct:"AAA",
        HolderCmID:"AAAID",

        WaitEndorseAcct:"",
        WaitEndorseCmID:"",
    }

    billByte, _ := json.Marshal(bill)
    err := stub.PutState(bill.BillInfoID, billByte)
    if err != nil {
        return shim.Error("初始化第一个票据失败: "+ err.Error())
    }

    bill2 := BillStruct{
        ObjectType:"billObj",
        BillInfoID:"POC102",
        BillInfoAmt:"2000",
        BillInfoType:"111",
        BillIsseDate:"20100201",
        BillDueDate:"20100210",

        HolderAcct:"AAA",
        HolderCmID:"AAAID",

        WaitEndorseAcct:"BBB",
        WaitEndorseCmID:"BBBID",
    }

    billByte2, _ := json.Marshal(bill2)
    err = stub.PutState(bill2.BillInfoID, billByte2)
    if err != nil {
        return shim.Error("初始化第二个票据失败: "+ err.Error())
    }

    bill3 := BillStruct{
        ObjectType:"billObj",
        BillInfoID:"POC103",
        BillInfoAmt:"3000",
        BillInfoType:"111",
        BillIsseDate:"20100301",
        BillDueDate:"20100310",

        HolderAcct:"BBB",
        HolderCmID:"BBBID",

        WaitEndorseAcct:"CCC",
        WaitEndorseCmID:"CCCID",
    }

    billByte3, _ := json.Marshal(bill3)
    err = stub.PutState(bill3.BillInfoID, billByte3)
    if err != nil {
        return shim.Error("初始化第三个票据失败: "+ err.Error())
    }

    bill4 := BillStruct{
        ObjectType:"billObj",
        BillInfoID:"POC104",
        BillInfoAmt:"4000",
        BillInfoType:"111",
        BillIsseDate:"20100401",
        BillDueDate:"20100410",

        HolderAcct:"CCC",
        HolderCmID:"CCCID",

        WaitEndorseAcct:"BBB",
        WaitEndorseCmID:"BBBID",
    }

    billByte4, _ := json.Marshal(bill4)
    err = stub.PutState(bill4.BillInfoID, billByte4)
    if err != nil {
        return shim.Error("初始化第四个票据失败: "+ err.Error())
    }

    return shim.Success([]byte("初始化票据成功"))
}

// 根据持票人的证件号码批量查询持票人的持有票据列表
func queryBills(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("必须且只能指定持票人的证件号码")
    }
    holderCmID := args[0]

    // 拼装CouchDB所需要的查询字符串(是标准的一个JSON串)
    // "{\"key\":{\"k\":\"v\", \"k\":\"v\"[,...]}}"
    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"billObj\", \"HoldrCmID\":\"%s\"}}", holderCmID)

    // 查询数据
    result, err := getBillsByQueryString(stub, queryString)
    if err != nil {
        return shim.Error("根据持票人的证件号码批量查询持票人的持有票据列表时发生错误: " + err.Error())
    }
    return shim.Success(result)
}

// 根据待背书人的证件号码批量查询待背书的票据列表
func queryWaitBills(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("必须且只能指定待背书人的证件号码")
    }

    waitEndorseCmID := args[0]
    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"billObj\", \"WaitEndorseCmID\":\"%s\"}}", waitEndorseCmID)

    result, err := getBillsByQueryString(stub, queryString)
    if err != nil {
        return shim.Error("根据待背书人的证件号码批量查询待背书的票据列表时发生错误: " + err.Error())
    }
    return shim.Success(result)
}

// 根据指定的查询字符串查询批量数据
func getBillsByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

    iterator, err := stub.GetQueryResult(queryString)
    if err != nil {
        return nil, err
    }
    defer  iterator.Close()

    var buffer bytes.Buffer
    var isSplit bool
    for iterator.HasNext() {
        result, err := iterator.Next()
        if err != nil {
            return nil, err
        }

        if isSplit {
            buffer.WriteString("; ")
        }

        buffer.WriteString("key:")
        buffer.WriteString(result.Key)
        buffer.WriteString(", Value: ")
        buffer.WriteString(string(result.Value))

        isSplit = true

    }

    return buffer.Bytes(), nil

}

func main() {
    err := shim.Start(new(CouchDBChaincode))
    if err != nil {
        fmt.Errorf("启动链码失败: %v", err)
    }
}