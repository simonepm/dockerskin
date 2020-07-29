const expect = require('chai').expect
const hello = require('./hello')

describe('TestSayHello', () => {
    let result = hello.SayHello()
    let expected = "Hello World"
    it(`Should say ${expected}`, () => {
        expect(result).to.equal(expected)
    })
})