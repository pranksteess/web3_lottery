# How To Get Rich By Doing Nothing
This repo provides an easy way to realize your daydream

## Compile
You can skip this step via downloading the executable file on [release page](https://github.com/pranksteess/web3_lottery/releases/tag/v1.0.0), if you are using MacOS.
```
go build -o web3_lottery main/main.go
```

## Run
```
chmod +x ./web3_lottery
nohup ./web3_lottery >> out.log &
```

## Others
 * The data source of whale ETH addresses is coming from [etherscan.io](https://etherscan.io/accounts) and [debank.com](https://debank.com/ranking)
 * You can choose binary search or bloom filter by edit the field `search_method` in config/config.yaml file
