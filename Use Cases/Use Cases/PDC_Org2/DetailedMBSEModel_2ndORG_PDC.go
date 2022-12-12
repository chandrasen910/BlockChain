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

// Create Detialed funtcion for MBSE Model and applying Private Data collection on org2
func (contract *SmartContract) CreateDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (*Detailed, error) {

	fmt.Printf("CreateDetailedMBSEModelPrivate start-->")

	stub := ctx.GetStub()

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if detail.AssetType != "MBSEModel" {
		return nil, fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return nil, fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 1 is creating Detailed MBSE Model or not
	if orgName != "org2" {
		return nil, fmt.Errorf("Given Organization doesn't support this function. %s")
	}
	// we are fetching Organization Roles from input JSON file
	orgRole, er1 := contract.GetOrgRoles(ctx)
	if er1 != nil {
		return nil, fmt.Errorf("Org role not defined properly. %s", er1.Error())
	}
	// we are fetching User Roles from input JSON file
	userRole, err := contract.GetUserRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("User role not defined properly. %s", err.Error())
	}
	//we are checking the Organization role is " LEAD Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return nil, fmt.Errorf("Insufficient Roles! LEAD & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Detailed or not
	if detail.ModelType != "Detailed" {
		return nil, fmt.Errorf("Model Type provided is incorrect")
	}

	//checking the Detailed MBSE Model is already created or not with Model Id using PDC
	detailAsBytes, err := ctx.GetStub().GetPrivateData("_implicit_org_Org2MSP", detail.ModelId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get detail model:" + err.Error())
	} else if detailAsBytes != nil {
		return nil, fmt.Errorf("Detail MBSE already exist: " + detail.ModelId)
	}

	//create method for the Detailed MBSE Model with Model Id in PDC
	err = stub.PutPrivateData("_implicit_org_Org2MSP", detail.ModelId, []byte(MBSEData))
	if err != nil {
		return nil, fmt.Errorf("Failed to create MBSE in ledger. %s", err.Error())
	}
	fmt.Println("Detailed MBSE Model Created Successfully")
	return nil, nil

}

// Update Detialed funtcion for MBSE Model and applying Private Data collection on org2
func (contract *SmartContract) UpdateDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (err error) {

	fmt.Println("UpdateDetailedMBSEModelPrivate start-->")

	stub := ctx.GetStub()

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if detail.AssetType != "MBSEModel" {
		return fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 1 is creating Detailed MBSE Model or not
	if orgName != "org2" {
		return fmt.Errorf("Given Organization doesn't support this function. %s")
	}

	// we are fetching Organization Roles from input JSON file
	orgRole, er1 := contract.GetOrgRoles(ctx)
	if er1 != nil {
		return fmt.Errorf("Org role not defined properly. %s", er1.Error())
	}

	// we are fetching User Roles from input JSON file
	userRole, err := contract.GetUserRoles(ctx)
	if err != nil {
		return fmt.Errorf("User role not defined properly. %s", err.Error())
	}

	//we are checking the Organization role is " LEAD Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return fmt.Errorf("Insufficient Roles! LEAD_Org & CSE roles are required.")
	}

	// we are fetching Model Type from input JSON file and checking the Model Type is Detailed or not
	if detail.ModelType != "Detailed" {
		return fmt.Errorf("Model Type provided is incorrect")
	}

	//checking the Detailed MBSE Model is already exists or not with Model Id using PDC
	detailAsBytes, err := ctx.GetStub().GetPrivateData("_implicit_org_Org2MSP", detail.ModelId)
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

	err = stub.PutPrivateData("_implicit_org_Org2MSP", detail.ModelId, detailJSONasBytes)
	if err != nil {
		return fmt.Errorf("Failed to update MBGetStateSE in ledger. %s", err.Error())
	}
	fmt.Println("Detailed MBSE Model updated successfully")
	return nil
}

// Delete Detialed funtcion for MBSE Model and applying Private Data collection on org2
func (contract *SmartContract) DeleteDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (err error) {

	fmt.Printf("DeleteDetailedMBSEModelPrivate start-->")
	stub := ctx.GetStub()

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if detail.AssetType != "MBSEModel" {
		return fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 1 is creating Detailed MBSE Model or not
	if orgName != "org2" {
		return fmt.Errorf("Given Organization doesn't support this function. %s")
	}
	// we are fetching Organization Roles from input JSON file
	orgRole, er1 := contract.GetOrgRoles(ctx)
	if er1 != nil {
		return fmt.Errorf("Org role not defined properly. %s", er1.Error())
	}
	// we are fetching User Roles from input JSON file
	userRole, err := contract.GetUserRoles(ctx)
	if err != nil {
		return fmt.Errorf("User role not defined properly. %s", err.Error())
	}
	//we are checking the Organization role is " LEAD Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return fmt.Errorf("Insufficient Roles! LEAD_Org & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Detailed or not
	if detail.ModelType != "Detailed" {
		return fmt.Errorf("Model Type provided is incorrect")
	}
	//checking the Detailed MBSE Model is already exists or not with Model Id using PDC
	mbseBytes, err1 := stub.GetPrivateData("_implicit_org_Org2MSP", detail.ModelId)
	if err1 != nil {
		return fmt.Errorf("Failed to get the MBSE. %s", err1.Error())
	}

	if mbseBytes == nil {
		return fmt.Errorf("No MBSE record found with ModelId. %s", detail.ModelId)
	}

	//Delete method for the Detailed MBSE Model with Model Id PDC
	err = ctx.GetStub().DelPrivateData("_implicit_org_Org2MSP", detail.ModelId)
	if err != nil {
		return fmt.Errorf("Failed to delete MBSE in ledger. %s", err.Error())
	}
	fmt.Println("Detailed MBSE Model deleted successfully")
	return nil

}

// Read Detialed funtcion for MBSE Model and applying Private Data collection on org2
func (contract *SmartContract) GetDetailedMBSEModelPrivate(ctx contractapi.TransactionContextInterface, MBSEData string) (*Detailed, error) {

	fmt.Printf("GetDetailedMBSEModelPrivate start-->")

	var detail Detailed

	err1 := json.Unmarshal([]byte(MBSEData), &detail)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if detail.AssetType != "MBSEModel" {
		return nil, fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return nil, fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 1 is creating Detailed MBSE Model or not
	if orgName != "org2" {
		return nil, fmt.Errorf("Given Organization doesn't support this function. %s")
	}
	// we are fetching Organization Roles from input JSON file
	orgRole, er1 := contract.GetOrgRoles(ctx)
	if er1 != nil {
		return nil, fmt.Errorf("Org role not defined properly. %s", er1.Error())
	}
	// we are fetching User Roles from input JSON file
	userRole, err := contract.GetUserRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("User role not defined properly. %s", err.Error())
	}
	//we are checking the Organization role is " LEAD Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return nil, fmt.Errorf("Insufficient Roles! LEAD & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Detailed or not
	if detail.ModelType != "Detailed" {
		return nil, fmt.Errorf("Model Type provided is incorrect")
	}
	//checking the Detailed MBSE Model is already exists or not with Model Id

	//Read method for the Detailed MBSE Model with Model Id PDC
	mbseBytes, err1 := ctx.GetStub().GetPrivateData("_implicit_org_Org2MSP", detail.ModelId)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to get the MBSE. %s", err1.Error())
	}
	if mbseBytes == nil {
		return nil, fmt.Errorf("No MBSE record found with mbse_id. %s", detail.ModelId)
	}
	fmt.Println("Detailed MBSE model fetched successfully")

	details := new(Detailed)
	_ = json.Unmarshal(mbseBytes, details)

	return details, nil

}
