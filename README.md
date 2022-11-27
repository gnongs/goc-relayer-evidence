# goc-relayer-evidence

1. Edit `main.go`

```go
const (
	rpc     = "YOUR_RPC_ENDPOINT"
	address = "YOUR_RELAYER_ADDRESS"
)
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
TOTAL_COUNT: 12 of 13
[287623 288069 288112 294001 295209 295935 296887 301010 305597 308727 314110 315114]
```
