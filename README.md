# basic-contract-interaction

API example for talking with a smart-contract deployed on polygon 



## Config

Copy `example.env` to `app.env` and config the private key 

```bash
cp example.env app.env
```


## Run

First compile 

```bash
go  build
```

Then run 
```bash
./basic-contract-interaction
```


## Test 


```bash
curl http://localhost:8080/product/17
```

```json
{
  "Name": "newprod",
  "Status": 1,
  "Owner": "0x08f5f9a336aae6a72c795ddf307864b13d13f0aa",
  "NewOwner": "0x08f5f9a336aae6a72c795ddf307864b13d13f0aa"
}
```
