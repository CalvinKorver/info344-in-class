//  @ts-check

"use strict"


// Load the express and morgan modueles
const express = require("express"); 
const morgan = require("morgan");   //Middleware handler that we can use many


const addr = process.env.ADDR || ":80";
const [host, port] = addr.split(":");
const portNum = parseInt(port);


const app = express();
// Use to write a response back to your client
// default stuff in dev
app.use(morgan(process.env.LOG_FORMAT || "dev")); 


// Request is eveything about the request, and res is the result
app.get("/", (req,res) => {
    res.set("Content-Type", "text/plain");
    res.send("Hello, Node.JS!");
});

app.get("/v1/users/me/hello", (req, res) => {
    let userJSON = req.get("X-User");
    if (!userJSON) {
        throw new Error("No X-User header provided");
    }
    
    let user = JSON.parse(userJSON);
    // Stringify and write it to the response
    res.json({
        message: `Hello, ${user.firstName} ${user.lastName}`
    });
});

app.use((err, req, res, next) => {
    console.error(err.stack); //Full stack trace
    res.set("Content-Type", "text/plain");
    res.send(err.message);
})

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}...`);
});