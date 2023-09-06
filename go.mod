module github.com/meshplus/bitxhub-core

go 1.13

require (
	github.com/binance-chain/tss-lib v1.3.3-0.20210411025750-fffb56b30511
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/bytecodealliance/wasmtime-go v0.37.0
	github.com/deckarep/golang-set v1.8.0
	github.com/ethereum/go-ethereum v1.12.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/hyperledger/fabric v2.1.1+incompatible
	github.com/hyperledger/fabric-protos-go v0.0.0-20201028172056-a3136dde2354
	github.com/iancoleman/orderedmap v0.2.0
	github.com/libp2p/go-libp2p-core v0.5.6
	github.com/looplab/fsm v0.2.0
	github.com/meshplus/bitxhub-kit v1.28.0
	github.com/meshplus/bitxhub-model v1.28.0
	github.com/meshplus/go-lightp2p v0.0.0-20200817105923-6b3aee40fa54
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.1
	go.uber.org/atomic v1.7.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
)

require (
	github.com/Shopify/sarama v1.27.0 // indirect
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/hyperledger/fabric-amcl v0.0.0-20210603140002-2670f91851c8 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	google.golang.org/genproto v0.0.0-20221014213838-99cd37c6964a // indirect
	google.golang.org/grpc v1.50.1 // indirect
)

replace github.com/binance-chain/tss-lib => github.com/dawn-to-dusk/tss-lib v1.3.2-0.20220422023240-5ddc16a330ed

replace github.com/agl/ed25519 => github.com/binance-chain/edwards25519 v0.0.0-20200305024217-f36fc4b53d43

replace google.golang.org/grpc => google.golang.org/grpc v1.33.0

//replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2
