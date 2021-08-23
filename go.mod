module github.com/ChainSafe/ChainBridge

go 1.13

require (
	github.com/ChainSafe/chainbridge-substrate-events v0.0.0-20200715141113-87198532025e
	github.com/ChainSafe/chainbridge-utils v1.0.5
	github.com/ChainSafe/log15 v1.0.0
	github.com/centrifuge/go-substrate-rpc-client v2.0.0+incompatible
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/ethereum/go-ethereum v1.10.7
	github.com/klauspost/compress v1.10.1 // indirect
	github.com/prometheus/client_golang v1.4.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.4.5
	gopkg.in/Graylog2/go-gelf.v2 v2.0.0-20191017102106-1550ee647df0
)

replace github.com/ChainSafe/chainbridge-utils v1.0.5 => github.com/equilibrium-eosdt/chainbridge-utils v0.0.0-20210517160540-c9475f17dc29
