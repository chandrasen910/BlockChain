package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Function to get extract the userId from ca identity.  It is required to for checking the minter
func (contract *SmartContract) GetOrgRoles(ctx contractapi.TransactionContextInterface) (string, error) {

	fmt.Printf("GetOrgRoles start-->")

	orgRole, found, _ := ctx.GetClientIdentity().GetAttributeValue("orgRole")
	if found == false {
		fmt.Println("orgRole not found!")
		return "", fmt.Errorf("orgRole not found!")
	}

	return orgRole, nil
}

// Function to get extract the userId from ca identity.  It is required to for checking the minter
func (contract *SmartContract) GetUserRoles(ctx contractapi.TransactionContextInterface) (string, error) {

	fmt.Printf("GetUserRoles start-->")

	userRole, ufound, _ := ctx.GetClientIdentity().GetAttributeValue("userRole")
	if ufound == false {
		return "", fmt.Errorf("userRole not found!")
	}

	return userRole, nil
}

// Function to get extract the OrgName from ca identity.  It is required to for checking the minter
func (contract *SmartContract) GetOrgName(ctx contractapi.TransactionContextInterface) (string, error) {

	fmt.Println("GetOrgName start-->")

	orgName, found, _ := ctx.GetClientIdentity().GetAttributeValue("organizationName")
	if found == false {
		fmt.Println("orgName not found!")
		return "", fmt.Errorf("orgName not found!")
	}

	return orgName, nil
}
