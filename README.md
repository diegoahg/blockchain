# Learn basic of Block chain with cars case

This code heavily inspired from "Code your own blockchain in less than [200 lines of Go](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)"

### Technologies used

* negroni - For application split ups
* spew - Print formatted structs in  console

### Deployment steps:
- `git clone https://github.com/diegoahg/blockchain.git`
- fill the environment variables in `config/.env`
- install depenedences `go get`
- run de app `go run main.go`

## GET /api/blocks
First validate if chain is correct and then get all blocks

## POST /api/blocks
Send the car data

Request example:
```
{
	"license_plate": "JCBR87",
	"owner": "Diego3"
}
```

## PUT /api/hack
Can edit block values

Request example:
```
{
	"index":2,
	"owner":"Omar",
	"hash":"uyfluyfyf"
}
```