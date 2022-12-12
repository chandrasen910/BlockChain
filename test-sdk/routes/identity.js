const express = require("express");
const router = express.Router();
const {registerEnrollUser} = require('../src/handlers/userController');

router.post('/', registerEnrollUser);

module.exports = router;