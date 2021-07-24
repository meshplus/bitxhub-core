module github.com/meshplus/bitxhub-core

go 1.13

require (
	github.com/Shopify/sarama v1.27.0 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.3
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/hyperledger/fabric v2.0.1+incompatible
	github.com/hyperledger/fabric-amcl v0.0.0-20200424173818-327c9e2cf77a // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20200330074707-cfe579e86986
	github.com/looplab/fsm v0.2.0
	github.com/meshplus/bitxhub-kit v1.2.1-0.20210723100713-b8d99c166281
	github.com/meshplus/bitxhub-model v1.1.2-0.20210724105649-55b7bde6bc91
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/sykesm/zap-logfmt v0.0.3 // indirect
	github.com/wasmerio/go-ext-wasm v0.3.1
	go.uber.org/zap v1.15.0 // indirect
)

replace github.com/ultramesh/crypto-gm => git.hyperchain.cn/dmlab/crypto-gm v0.2.14
