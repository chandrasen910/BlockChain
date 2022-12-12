package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Summary struct {
	SummaryId string `json:"summaryId"`
	Version   string `json:"version"`
	AssetType string `json:"assetType"`
	ModelId   string `json:"modelId"`
	ProjectId string `json:"projectId"`
	ModelName string `json:"modelName"`
	ModelType string `json:"modelType"`
	Reviewer  string `json:"reviewer"`
}

// Create funtcion for Summary MBSE Model
func (contract *SmartContract) CreateSummaryMBSEModel(ctx contractapi.TransactionContextInterface, MBSEData string) (err error) {

	fmt.Printf("CreateDetailedMBSEModel start-->")

	stub := ctx.GetStub()
	var summary Summary

	err1 := json.Unmarshal([]byte(MBSEData), &summary)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if summary.AssetType != "MBSEModel" {
		return fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 2 is creating Summary MBSE Model or not
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
	//we are checking the Organization role is " Lead Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return fmt.Errorf("Insufficient Roles! LEAD Organization & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Summary or not
	if summary.ModelType != "Summary" {
		return fmt.Errorf("Model Type provided is incorrect")
	}
	//checking the Summary MBSE Model is already created or not with Summary Id
	mbseBytes, err1 := ctx.GetStub().GetState(summary.SummaryId)
	if err1 != nil {
		return fmt.Errorf("Failed to get the MBSE. %s", err1.Error())
	} else if mbseBytes != nil {
		return fmt.Errorf("record already exists with SummaryId. %s", summary.SummaryId)
	}

	//create method for the Summary MBSE Model with Summary Id
	err = stub.PutState(summary.SummaryId, []byte(MBSEData))
	if err != nil {
		return fmt.Errorf("Failed to insert MBSE in ledger. %s", err.Error())
	}
	fmt.Println("Summary MBSE Model created successfully")

	return nil
}

// Update funtcion for Summary MBSE Model
func (contract *SmartContract) UpdateSummaryMBSEModel(ctx contractapi.TransactionContextInterface, MBSEData string) (err error) {

	fmt.Printf("UpdateSummaryMBSEModel start-->")

	stub := ctx.GetStub()
	var summary Summary

	err1 := json.Unmarshal([]byte(MBSEData), &summary)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if summary.AssetType != "MBSEModel" {
		return fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 2 is creating Summary MBSE Model or not
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
	//we are checking the Organization role is " Lead Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return fmt.Errorf("Insufficient Roles! LEAD Organization & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Summary or not
	if summary.ModelType != "Summary" {
		return fmt.Errorf("Model Type provided is incorrect")
	}
	//checking the Summary MBSE Model is already exists or not with Summary Id
	summaryAsBytes, err := ctx.GetStub().GetState(summary.SummaryId)

	if err != nil {
		return fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if summaryAsBytes == nil {
		return fmt.Errorf("Summary %s does not exist", summary.SummaryId)
	}

	summaryData := new(Summary)
	_ = json.Unmarshal(summaryAsBytes, summaryData)

	summaryData.ProjectId = summary.ProjectId
	summaryData.Reviewer = summary.Reviewer

	summaryDAsBytes, _ := json.Marshal(summaryData)

	//Update method for the Summary MBSE Model with Summary Id
	err = stub.PutState(summary.SummaryId, summaryDAsBytes)
	if err != nil {
		return fmt.Errorf("Failed to update MBSE summary in ledger. %s", err.Error())
	}
	fmt.Println("Summary MBSE Model updated successfully")

	return nil
}

// Delete funtcion for Summary MBSE Model
func (contract *SmartContract) DeleteSummaryMBSEModel(ctx contractapi.TransactionContextInterface, MBSEData string) error {

	fmt.Printf("DeleteSummaryMBSEModel start-->")

	myMSPID, _ := contract.GetOrgName(ctx)

	fmt.Println("DeleteSummaryMBSEModel myMSPID-->", myMSPID)
	var summary Summary

	err1 := json.Unmarshal([]byte(MBSEData), &summary)
	if err1 != nil {
		return fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if summary.AssetType != "MBSEModel" {
		return fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 2 is creating Summary MBSE Model or not
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
	//we are checking the Organization role is " Lead Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return fmt.Errorf("Insufficient Roles! LEAD Organization & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Summary or not
	if summary.ModelType != "Summary" {
		return fmt.Errorf("Model Type provided is incorrect")
	}
	//checking the Summary MBSE Model is already exists or not with Summary Id
	mbseBytes, err1 := ctx.GetStub().GetState(summary.SummaryId)
	if err1 != nil {
		return fmt.Errorf("Failed to get the MBSE. %s", err1.Error())
	}

	if mbseBytes == nil {
		return fmt.Errorf("No MBSE record found with mbse_id. %s", summary.SummaryId)
	}

	//Delete method for the Summary MBSE Model with Summary Id
	err = ctx.GetStub().DelState(summary.SummaryId)
	if err != nil {
		return fmt.Errorf("Failed to Delete MBSE summary in ledger. %s", err.Error())
	}
	fmt.Println("Summary MBSE Model Deleted successfully")
	return nil
}

// Read funtcion for Summary MBSE Model
func (contract *SmartContract) GetSummaryMBSEModel(ctx contractapi.TransactionContextInterface, MBSEData string) (*Summary, error) {

	fmt.Printf("GetSummaryMBSEModel start-->")

	myMSPID, _ := contract.GetOrgName(ctx)

	fmt.Println("GetSummaryMBSEModel myMSPID-->", myMSPID)

	var summary Summary

	err1 := json.Unmarshal([]byte(MBSEData), &summary)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to parse MBSE argument. %s", err1.Error())
	}
	//we are checking the Blockchain Asset Type should be MBSE Model or not
	if summary.AssetType != "MBSEModel" {
		return nil, fmt.Errorf("Asset Type provided is incorrect")
	}
	// we are fetching Organization Name from input JSON file
	orgName, er1 := contract.GetOrgName(ctx)
	if er1 != nil {
		return nil, fmt.Errorf("Org Name not defined properly. %s", er1.Error())
	}
	// we are checking Organization 2 is creating Summary MBSE Model or not
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
	//we are checking the Organization role is " Lead Organization"  & User roles is "CSE"
	if orgRole != "LEAD" || userRole != "CSE" {
		return nil, fmt.Errorf("Insufficient Roles! LEAD Organization & CSE roles are required.")
	}
	// we are fetching Model Type from input JSON file and checking the Model Type is Summary or not
	if summary.ModelType != "Summary" {
		return nil, fmt.Errorf("Model Type provided is incorrect")
	}
	//checking the Summary MBSE Model is already exists or not with Summary Id

	//Read method for the Summary MBSE Model with Model Id
	mbseBytes, err1 := ctx.GetStub().GetState(summary.SummaryId)
	if err1 != nil {
		return nil, fmt.Errorf("Failed to get the MBSE in ledger. %s", err1.Error())
	}

	if mbseBytes == nil {
		return nil, fmt.Errorf("No MBSE record found with mbse_id. %s", summary.SummaryId)
	}

	fmt.Println("Summary MBSE Model fetched successfully")

	summary1 := new(Summary)
	_ = json.Unmarshal(mbseBytes, summary1)

	return summary1, nil

}
