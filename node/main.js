'use strict'

const express = require('express')
const hello = require('./hello')
const app = express()
const port = 8080

app.get('/', (req, res) => {
    res.send(hello.SayHello() + '\n')
})

console.log(`Listening on port ${port}...`)

app.listen(port, '0.0.0.0')