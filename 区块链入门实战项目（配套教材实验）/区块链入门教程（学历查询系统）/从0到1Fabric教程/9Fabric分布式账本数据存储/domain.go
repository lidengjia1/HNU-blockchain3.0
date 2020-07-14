package main

type BillStruct struct {
    ObjectType    string    `json:"docType"`
    BillInfoID    string    `json:"BillInfoID"`
    BillInfoAmt    string    `json:"BillInfoAmt"`
    BillInfoType string    `json:"BillInfoType"`

    BillIsseDate    string    `json:"BillIsseDate"`
    BillDueDate    string    `json:"BillDueDate"`

    HolderAcct    string    `json:"HolderAcct"`
    HolderCmID    string    `json:"HolderCmID"`

    WaitEndorseAcct    string    `json:"WaitEndorseAcct"`
    WaitEndorseCmID    string    `json:"WaitEndorseCmID"`

}