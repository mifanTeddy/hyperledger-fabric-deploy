package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

type AssertsExchangeCC struct{}

type History struct {
    Data string `json:"data"`
}

func setData(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
        return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err != nil {
        return "", fmt.Errorf("Failed to set asset: %s", args[0])
    }
    txId := stub.GetTxID()
    fmt.Println(txId)
    return txId, nil
}

func getData(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
        return "", fmt.Errorf("Incorrect arguments. Expecting a key")
    }

    value, err := stub.GetState(args[0])
    if err != nil {
        return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
    }
    if value == nil {
        return "", fmt.Errorf(" not found: %s", args[0])
    }
    return string(value), nil
}


func getHistory(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    keyInter, err := stub.GetHistoryForKey(args[0])
    if err != nil {
        return "", fmt.Errorf(" not found: %s", args[0])
    }
    history := make([]string, 0)
    for keyInter.HasNext(){
        response, interErr := keyInter.Next()
        if interErr != nil{
            return "", fmt.Errorf(" not found: %s", args[0])
        }
        history = append(history, string(response.Value))
    }
    historiesBytes, err := json.Marshal(history)
    return string(historiesBytes), nil
}


func (c *AssertsExchangeCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    funcName, args := stub.GetFunctionAndParameters()
    var result string
    var err error
    switch funcName {
    case "set":
        result, err = setData(stub, args)
    case "get":
        result, err = getData(stub, args)
    case "getHistory":
        result, err = getHistory(stub, args)
    default:
        return shim.Error(fmt.Sprintf("unsupported function: %s", funcName))
    }
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success([]byte(result))
}



func (c *AssertsExchangeCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func main() {
    err := shim.Start(new(AssertsExchangeCC))
    if err != nil {
        fmt.Printf("Error starting AssertsExchange chaincode: %s", err)
    }
}

