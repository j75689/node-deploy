## Cross-Chain Bridge
```bash
echo "12345678" | ./bin/tbnbcli bridge transfer-out --amount 500000000:BNB --expire-time $(date --date="+300 seconds" +%s) --to 0xF34273E326FF62bB1F7aD9b7bAf62DA9BE40ffe4 --from node0-delegator --chain-id Binance-Chain-Ganges --node http://172.22.42.13:26657 --home ./.local/bc/node0


curl -X POST "http://172.22.42.13:8545" -H "Content-Type: application/json"  --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xF34273E326FF62bB1F7aD9b7bAf62DA9BE40ffe4", "latest"],"id":1}' 
```