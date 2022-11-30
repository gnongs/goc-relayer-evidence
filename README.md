# goc-relayer-evidence

**_If your relayer set 0 fee, you con't check heights with this program_**

1. Edit `config.json`

```json
[
    {
        "chain": "sputnik",
        "address": "RELAYER_ADDRESS",
        "rpc_endpoint": "API_ENDPOINT"
    },
    {
        "chain": "gopher",
        "address": "RELAYER_ADDRESS",
        "rpc_endpoint": "API_ENDPOINT"
    }
]
```

2. build binary
```sh
go build
```

3. Start binary
```sh
./goc-relayer-evidence
```

4. Print your evidence height
```sh
hero 's total_count: 23 of 25
[62882 63334 63377 69408 70645 71387 72359 76578 81283 84498 90055 91093 93501 94089 95066 97244 97270 97295 97993 98534 106299 106453 106547]
```
