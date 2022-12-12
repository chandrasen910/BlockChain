const express = require("express");
const router = express.Router();
const MBSEController = require('../src/handlers/MBSE-controller');
const mbseController = new MBSEController();

// router.post("/", mbseController.addMBSEDetails );


router.post("/private", mbseController.addMBSEPrivateDetails );
router.delete("/private", mbseController.deleteMBSEPrivateDetails );
router.put("/private", mbseController.updateMBSEPrivateDetails );
router.get("/private", mbseController.getMBSEPrivateDetails );


router.post("/summary", mbseController.addMBSESummary );
router.put("/summary", mbseController.updateMBSESummary );
router.delete("/summary", mbseController.deleteMBSESummary );
router.get("/summary", mbseController.getMBSESummary );


module.exports = router;    