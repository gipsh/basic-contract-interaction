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

#### Get product by id

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

#### Get delegated products from a wallet

```bash
curl http://localhost:8080/products/0xCF6380c9B128941d20d9F812dA406A79424b4B7B
```

```json
{
    "address": "0xCF6380c9B128941d20d9F812dA406A79424b4B7B",
    "products": [
        "34",
        "35",
        "37"
    ]
}
```

