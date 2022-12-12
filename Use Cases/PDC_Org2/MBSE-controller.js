/*
 * Copyright. All Rights Reserved.
 *
 * IBM-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');
const { getContractObject } = require('../../utils/util.js');
const { NETWORK_PARAMETERS, DOCTYPE } = require('../../utils/Constants');
const logger = require('../../logger')(module);
// const { formatReferences, formatAssetInput } = require('../../utils/FormatStruct');
// const AssetType = require('../../models/AssetType')

class TokenController {

	constructor() {
	}

	async addMBSEPrivateDetails(req, res, next) {
		try {
			console.log('*******Creating Detailed MBSE Model *******')
	
			let orgId = req.body.ownerOrgId;
			let orgName = "org" + orgId
			let user = req.body.userId;
			let tokenDef = req.body.data;
	
			const gateway = new Gateway();
			let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
			// let tokenData = JSON.parse(tokenDef);
			console.log('----------Creating Detailed MBSE Model------------\n', tokenDef)
			let stateTxn = contract.createTransaction('CreateDetailedMBSEModelPrivate');
			stateTxn.setEndorsingOrganizations('Org2MSP');
			let tx = await stateTxn.submit(JSON.stringify(tokenDef));
			console.log('*** Detailed MBSE Model created: committed');
			return res.status(200).send({
				status: true,
				message: "Detailed MBSE Model Created Successfully",
				txid: tx.toString()
			});
		} catch (error) {
			console.log(error.message)
			logger.error({ userInfo: req.loggerInfo, method: 'CreateDetailedMBSEModel', error })
			return res.status(500).send({
				status: false,
				message: error.message
			});
		}
	}
	

	async updateMBSEPrivateDetails(req, res, next) {
		try {
			console.log('*******Update Detailed MBSE Model  *******')
	
			let orgId = req.body.ownerOrgId;
			let orgName = "org" + orgId
			let user = req.body.userId;
			let tokenDef = req.body.data;
	
			const gateway = new Gateway();
			let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
			// let tokenData = JSON.parse(tokenDef);
			console.log('----------Updating Detailed MBSE Model------------\n', tokenDef)
			let stateTxn = contract.createTransaction('UpdateDetailedMBSEModelPrivate');
			stateTxn.setEndorsingOrganizations('Org2MSP');
			let tx = await stateTxn.submit(JSON.stringify(tokenDef));
			console.log('*** Detailed MBSE Model Updated: committed');
			// let tx ='xxxxxxxxxxxxxxxxx'
			return res.status(200).send({
				status: true,
				message: "Detailed MBSE Model Updated Successfully",
				txid: tx.toString()
			});
		} catch (error) {
			console.log(error.message)
			logger.error({ userInfo: req.loggerInfo, method: 'CreateDetailedMBSEModel', error })
			return res.status(500).send({
				status: false,
				message: error.message
			});
		}
	}

async deleteMBSEPrivateDetails(req, res, next) {
	try {
		console.log('*******Delete Detailed MBSE Model *******')

		let orgId = req.body.ownerOrgId;
		let orgName = "org" + orgId
		let user = req.body.userId;
		let tokenDef = req.body.data;
		
		console.log("data----------", JSON.stringify(tokenDef))
		const gateway = new Gateway();
		console.log('----------Deleting Detailed MBSE Model------------\n', JSON.stringify(tokenDef))
		let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
		let stateTxn = contract.createTransaction('DeleteDetailedMBSEModelPrivate');
		stateTxn.setEndorsingOrganizations('Org2MSP');
		let tx = await stateTxn.submit(JSON.stringify(tokenDef));
		
		console.log('*** Detailed MBSE Model deleted: committed');
		return res.status(200).send({
			status: true,
			message: "Detailed MBSE Model deleted Successfully",
			txid: tx.toString()
		});
	} catch (error) {
		console.log(error.message)
		logger.error({ userInfo: req.loggerInfo, method: 'DeleteDetailedMBSEModel', error })
		return res.status(500).send({
			status: false,
			message: error.message
		});
	}
}

async getMBSEPrivateDetails(req, res, next) {
	try {
		console.log('*******Read Detailed MBSE Model *******')

		let orgId = req.body.ownerOrgId;
		let orgName = "org" + orgId
		let user = req.body.userId;
		let tokenDef = req.body.data;
		
		console.log("mbseId----------", tokenDef )
		const gateway = new Gateway();
		console.log('----------Fetching Detailed MBSE Model------------\n', tokenDef )
		let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
		// let stateTxn = contract.createTransaction('GetDetailedMBSEModelPrivate');
		// stateTxn.setEndorsingOrganizations('Org2MSP');
		let result = await contract.evaluateTransaction('GetDetailedMBSEModelPrivate', JSON.stringify(tokenDef) );
		result = JSON.parse(result.toString())
		return res.status(200).send({
			success: true,
			message: `Detailed MBSE Model fetching successfully`,
			payload: result
		});
	} catch (error) {
		console.log(error.message)
		logger.error({ userInfo: req.loggerInfo, method: 'getMBSEPrivateDetails', error })
		return res.status(500).send({
			status: false,
			message: error.message
		});
	}
}


async addMBSESummary(req, res, next) {
	try {
		console.log('*******Add MBSE Sumamry *******')

		let orgId = req.body.ownerOrgId;
		let orgName = "org" + orgId
		let user = req.body.userId;
		let tokenDef = req.body.data;

		const gateway = new Gateway();
		let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
		// let tokenData = JSON.parse(tokenDef);
		console.log('----------Creating Summary MBSE Model ------------\n', tokenDef)
		let stateTxn = contract.createTransaction('CreateSummaryMBSEModel');
		let tx = await stateTxn.submit(JSON.stringify(tokenDef));
		console.log('*** Summary MBSE Model created: committed');
		// let tx ='xxxxxxxxxxxxxxxxx'
		return res.status(200).send({
			status: true,
			message: "Summary MBSE Model created successfully",
			txid: tx.toString()
		});
	} catch (error) {
		console.log(error.message)
		logger.error({ userInfo: req.loggerInfo, method: 'CreateSummaryMBSEModel', error })
		return res.status(500).send({
			status: false,
			message: error.message
		});
	}
}



async updateMBSESummary(req, res, next) {
	try {
		console.log('*******updateMBSESummary  *******')

		let orgId = req.body.ownerOrgId;
		let orgName = "org" + orgId
		let user = req.body.userId;
		let tokenDef = req.body.data;

		const gateway = new Gateway();
		let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
		// let tokenData = JSON.parse(tokenDef);
		console.log('----------Updating MBSE Summary Model ------------\n', tokenDef)
		let stateTxn = contract.createTransaction('UpdateSummaryMBSEModel');
		let tx = await stateTxn.submit(JSON.stringify(tokenDef));
		console.log('*** MBSE Summary Model Updated: committed');
		// let tx ='xxxxxxxxxxxxxxxxx'
		return res.status(200).send({
			status: true,
			message: "MBSE Summary Model Updated Successfully",
			txid: tx.toString()
		});
	} catch (error) {
		console.log(error.message)
		logger.error({ userInfo: req.loggerInfo, method: 'updateMBSESummary', error })
		return res.status(500).send({
			status: false,
			message: error.message
		});
	}
}

//new
async getMBSESummary(req, res, next) {
	try {
		console.log('*******Read MBSE Summary MBSE Model *******')

		let orgId = req.body.ownerOrgId;
		let orgName = "org" + orgId
		let user = req.body.userId;
		let tokenDef  = req.body.data;
		
		console.log("tokenDef ----------", tokenDef)
		const gateway = new Gateway();
		console.log('----------Fetching MBSE Summary MBSE Model------------\n', tokenDef)
		let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
		let result = await contract.evaluateTransaction('GetSummaryMBSEModel', JSON.stringify(tokenDef));
		result = JSON.parse(result.toString())
		return res.status(200).send({
			success: true,
			message: `MBSE Summary Model fetching successfully`,
			payload: result
		});
	} catch (error) {
		console.log(error.message)
		logger.error({ userInfo: req.loggerInfo, method: 'getMBSESummary', error })
		return res.status(500).send({
			status: false,
			message: error.message
		});
	}
}


async deleteMBSESummary(req, res, next) {
	try {
		console.log('*******Delete Summary MBSE Model *******')

		let orgId = req.body.ownerOrgId;
		let orgName = "org" + orgId
		let user = req.body.userId;
		let tokenDef = req.body.data;
		
		console.log("tokenDef----------", tokenDef)
		const gateway = new Gateway();
		console.log('----------Deleting Summary MBSE Model------------\n', tokenDef)
		let contract = await getContractObject(orgName, user, NETWORK_PARAMETERS.CHANNEL_NAME, NETWORK_PARAMETERS.CHAINCODE_NAME, gateway)
		let stateTxn = contract.createTransaction('DeleteSummaryMBSEModel');
		let tx = await stateTxn.submit(JSON.stringify(tokenDef));
		
		console.log('*** Detailed MBSE Model deleted: committed');
		// let tx ='xxxxxxxxxxxxxxxxx'
		return res.status(200).send({
			status: true,
			message: "Detailed MBSE Model deleted Successfully",
			txid: tx.toString()
		});
	} catch (error) {
		console.log(error.message)
		logger.error({ userInfo: req.loggerInfo, method: 'deleteMBSESummary', error })
		return res.status(500).send({
			status: false,
			message: error.message
		});
	}
}

}
module.exports = TokenController;
