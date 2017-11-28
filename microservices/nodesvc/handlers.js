//  @ts-check
"use strict";


const express = require("express");
module.exports = (mongoSession) => {
    if (!mongoSession) {
        throw new Error("provide a mongo session");
    }

    let router = express.Router();

    router.get("/v1/channels", (req, res) => {
        // Query mongo using mongo session
        res.json([{name: "general"}]);
    });

    return router;
}