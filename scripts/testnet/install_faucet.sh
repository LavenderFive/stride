
cd /stride 
git clone https://github.com/tendermint/faucet.git
cd faucet 
make install
/stride/go/bin/faucet --cli-name strided --denoms ustrd --account-name val2 --keyring-backend test --home /stride/.stride --credit-amount 1000000000 --max-credit 3000000001

# curl -X POST -d '{"address": "stride159atdlc3ksl50g0659w5tq42wwer334ajl7xnq"}' stride-node2.testnet-vishal.stridenet.co:8000