#!/usr/bin/env node  
"use strict";
// This is going to send messages to the queue

const amqp = require("amqplib");
const qName = "testQ";
const mqAddr = process.env.MQADDR || "localhost:5672";
const mqURL = `amqp://${mqAddr}`;


// Next line of code will nt execute
(async function() {
    console.log("connecting to %s", mqURL);
    let connection = await amqp.connect(mqURL);
    let channel = await connection.createChannel();
    // Since its durable it will write mqmessages to disc if true
    let qConf = await channel.assertQueue(qName, {durable: false});

    console.log("starting to send messages...");
    setInterval(() => {
        let message = "Message sent at: " + new Date().toLocaleTimeString();  
        channel.sendToQueue(qName, Buffer.from(msg));
    }, 1000);
})();