package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Detailed struct {
	ModelId   string `json:"modelId"`
	Version   string `json:"version"`
	ModelName string `json:"modelName"`
	ProjectId string `json:"projectId"`
	ModelType string `json:"modelType"`
	AssetType string `json:"assetType"`
}

// Create Detialed funtcion for MBSE Model and applying Private Data collection on org1
func (contract *SmartContract) CreateDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (*Detailed, error) {

	fmt.Printf("CreateDetailedMBSEModelPrivate start-->")

	stub := ctx.GetStub()

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}

	//checking the Detailed MBSE Model is already created or not with Model Id using PDC
	detailAsBytes, err := ctx.GetStub().GetPrivateData("_implicit_org_Org1MSP", detail.ModelId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get detail model:" + err.Error())
	} else if detailAsBytes != nil {
		return nil, fmt.Errorf("Detail MBSE already exist: " + detail.ModelId)
	}
	//create method for the Detailed MBSE Model with Model Id in PDC
	err = stub.PutPrivateData("_implicit_org_Org1MSP", detail.ModelId, []byte(MBSEData))
	if err != nil {
		return nil, fmt.Errorf("Failed to create MBSE in ledger. %s", err.Error())
	}
	fmt.Println("MBSE Detailed Model Created Successfully")
	return nil, nil
}

// Update Detialed funtcion for MBSE Model and applying Private Data collection on org1
func (contract *SmartContract) UpdateDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (err error) {

	fmt.Println("UpdateDetailedMBSEModelPrivate start-->")

	stub := ctx.GetStub()

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}

	//checking the Detailed MBSE Model is already exists or not with Model Id using PDC
	detailAsBytes, err := ctx.GetStub().GetPrivateData("_implicit_org_Org1MSP", detail.ModelId)
	if err != nil {
		return fmt.Errorf("Failed to get detailed model:" + err.Error())
	} else if detailAsBytes == nil {
		return fmt.Errorf("Detailed MBSE Model does not exist: " + detail.ModelId)
	}
	//update method for the Detailed MBSE Model with existing data
	detailToUpdate := Detailed{}
	err = json.Unmarshal(detailAsBytes, &detailToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return fmt.Errorf("failed to unmarshal Summary JSON: %s", err.Error())
	}

	detailToUpdate.ModelName = detail.ModelName //update the Model Name
	detailToUpdate.ProjectId = detail.ProjectId //update the Project ID

	detailJSONasBytes, _ := json.Marshal(detailToUpdate)

	err = stub.PutPrivateData("_implicit_org_Org1MSP", detail.ModelId, detailJSONasBytes)
	if err != nil {
		return fmt.Errorf("Failed to update MBGetStateSE in ledger. %s", err.Error())
	}
	fmt.Println("MBSE Detailed model updated successfully")
	return nil
}

// Delete Detialed funtcion for MBSE Model and applying Private Data collection on org1
func (contract *SmartContract) DeleteDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (err error) {

	fmt.Printf("DeleteDetailedMBSEModelPrivate start-->")
	stub := ctx.GetStub()

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}

	mbseBytes, err1 := stub.GetPrivateData("_implicit_org_Org1MSP", detail.ModelId)
	if err1 != nil {
		return fmt.Errorf("Failed to get the MBSE. %s", err1.Error())
	}

	if mbseBytes == nil {
		return fmt.Errorf("No MBSE record found with ModelId. %s", detail.ModelId)
	}
	//Delete method for the Detailed MBSE Model with Model Id PDC
	err = ctx.GetStub().DelPrivateData("_implicit_org_Org1MSP", detail.ModelId)
	if err != nil {
		return fmt.Errorf("Failed to delete MBSE in ledger. %s", err.Error())
	}

	fmt.Println("MBSE Detailed MBSE model deleted successfully")
	return nil

}

// Read Detialed funtcion for MBSE Model and applying Private Data collection on org1
func (contract *SmartContract) GetDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (*Detailed, error) {

	fmt.Printf("GetDetailedMBSEModelPrivate start-->")

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}

	//checking the Detailed MBSE Model is already exists or not with Model Id
	//Read method for the Detailed MBSE Model with Model Id PDC
	mbseBytes, err1 := ctx.GetStub().GetPrivateData("_implicit_org_Org1MSP", detail.ModelId)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to get the MBSE. %s", err1.Error())
	}

	if mbseBytes == nil {
		return nil, fmt.Errorf("No MBSE record found with mbse_id. %s", detail.ModelId)
	}

	fmt.Println("MBSE Detailed model fetched successfully")

	details := new(Detailed)
	_ = json.Unmarshal(mbseBytes, details)

	return details, nil

}
