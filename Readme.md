## Fetch Submodule
```
git submodule update --recursive --remote
```

## Cross-Chain Bridge
```bash
echo "12345678" | ./bin/tbnbcli bridge transfer-out --amount 500000000:BNB --expire-time $(date --date="+300 seconds" +%s) --to 0xF34273E326FF62bB1F7aD9b7bAf62DA9BE40ffe4 --from node0-delegator --chain-id Binance-Chain-Ganges --node http://172.22.42.13:26657 --home ./.local/bc/node0


curl -X POST "http://172.22.42.13:8545" -H "Content-Type: application/json"  --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xF34273E326FF62bB1F7aD9b7bAf62DA9BE40ffe4", "latest"],"id":1}' 
```


## ABI GENERATION
```bash
abigen --abi=./migrate_validator_tool/abi/stakehub/stakehub.abi --pkg=stakehub --out=./migrate_validator_tool/abi/stakehub/stakehub.go
abigen --abi=./migrate_validator_tool/abi/validatorset/validatorset.abi --pkg=validatorset --out=./migrate_validator_tool/abi/validatorset/validatorset.go
abigen --abi=./migrate_validator_tool/abi/crosschain/crosschain.abi --pkg=crosschain --out=./migrate_validator_tool/abi/crosschain/crosschain.go
abigen --abi=./migrate_validator_tool/abi/bscgovernor/bscgovernor.abi --pkg=bscgovernor --out=./migrate_validator_tool/abi/bscgovernor/bscgovernor.go
abigen --abi=./migrate_validator_tool/abi/govtoken/govtoken.abi --pkg=govtoken --out=./migrate_validator_tool/abi/govtoken/govtoken.go
abigen --abi=./migrate_validator_tool/abi/tokenrecoverportal/tokenrecoverportal.abi --pkg=tokenrecoverportal --out=./migrate_validator_tool/abi/tokenrecoverportal/tokenrecoverportal.go
abigen --abi=./migrate_validator_tool/abi/govhub/govhub.abi --pkg=govhub --out=./migrate_validator_tool/abi/govhub/govhub.go
abigen --abi=./migrate_validator_tool/abi/tokenhub/tokenhub.abi --pkg=tokenhub --out=./migrate_validator_tool/abi/tokenhub/tokenhub.go


```

## Extract Bytecode
```bash
bash ./extract_bytecode.sh <genesis.json> ./bsc_fyenman_bytecode
```