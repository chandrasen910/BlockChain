const logger = require('../logger')(module)
const {READ_ACCESS, WRITE_ACCESS, DELETE_ACCESS, ALL_ACCESS} = require("../utils/ErrorMessages.js");
// const User = require('../models/user');
const AssetType = require('../models/AssetType')
const AssetAccess = require('../models/AssetPermission')
const { READ, CREATE, UPDATE, DELETE, ALL, NO} = require('../utils/Constants')


const checkOrgAccess = async(req, res, next) => {

    try {

        logger.info({ userInfo: req.loggerInfo, method: 'checkOrgAccess' })
        // console.log(req) 
        console.log(req.method)
        console.log(req.query.hasOwnProperty('assetType'))
        let assetType = req.query.assetType || req.body.assetType

        console.log(assetType)

        let objAssetType = await AssetType.findOne({ assetType: assetType });
        console.log(objAssetType)
        
        let access 
        if (objAssetType && Object.keys(objAssetType).length != 0) {

            let objAssetAccess = await AssetAccess.findOne({ assetTypeId: objAssetType.assetTypeId });
            console.log('----------objAssetAccess------------')
            // console.log(objAssetAccess)
    
            if (objAssetAccess && Object.keys(objAssetAccess).length != 0 && objAssetAccess.role != 'ALL' ) {
                
                console.log('NOT Full Access!!!!!!!!!!!!!!!!')
                console.log(objAssetAccess.role)
                console.log(req.method)

                switch (objAssetAccess.role) {
                    case 'READ':
                        (READ.includes(req.method))? access = 'READ' : access =""
                        break;

                    case 'CREATE':
                        (CREATE.includes(req.method))? access = 'CREATE' : access =""
                        break;

                    case 'UPDATE':
                        (UPDATE.includes(req.method))? access = 'UPDATE' : access =""
                        break;

                    case 'DELETE':
                        (DELETE.includes(req.method))? access = 'DELETE' : access =""
                        break;

                    case 'ALL':
                        (ALL.includes(req.method))? access = 'ALL' : access =""
                        break;

                    default:
                        return new Error('Not a pre defined permission')
                }
        
            } else{
                access ='FULL'
            }

            if (access =="") {
                throw new Error('Insufficient Privilages!!')
            }

        } else {
            return res.status(404).send({
                status: false,
                message: "assetType Not defined."
            });
        
        }

        console.log(`${access} access provided. `)
        logger.info({userInfo: req.loggerInfo, method: 'checkOrgAccess', message: `${access} access provided. `})

        next();

    } catch(error) {
        console.log(error.message)
        logger.error({userInfo: req.loggerInfo, method: 'checkOrgAccess'})
        return res.status(500).send({ success: false,
            message: error.message });
    }

}

module.exports = {
    checkOrgAccess
}