# How To Get Rich By Doing Nothing
This library provides an easy way to realize your daydream

## Compile & Run
```
go build -o web3_lottery main/main.go
nohup ./web3_lottery >> out.log &
```

## Others
 * The data source of whale ETH addresses is coming from [etherscan.io](https://etherscan.io/accounts) and [debank.com](https://debank.com/ranking)
 * You can choose binary search or bloom filter by edit the field `search_method` in config/config.yaml file
