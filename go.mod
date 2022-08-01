module github.com/meshplus/bitxhub-core

go 1.13

require (
	github.com/Shopify/sarama v1.27.0 // indirect
	github.com/binance-chain/tss-lib v1.3.3-0.20210411025750-fffb56b30511
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/bytecodealliance/wasmtime-go v0.34.0
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/ethereum/go-ethereum v1.10.4
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.5.0
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/hyperledger/fabric v2.1.1+incompatible
	github.com/hyperledger/fabric-amcl v0.0.0-20210603140002-2670f91851c8 // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20201028172056-a3136dde2354
	github.com/iancoleman/orderedmap v0.2.0
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/libp2p/go-libp2p-core v0.5.6
	github.com/looplab/fsm v0.2.0
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/meshplus/bitxhub-kit v1.2.1-0.20220325052414-bc17176c509d
	github.com/meshplus/bitxhub-model v1.2.1-0.20220801074028-c71727dacea0
	github.com/meshplus/go-lightp2p v0.0.0-20200817105923-6b3aee40fa54
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	go.uber.org/atomic v1.7.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4 // indirect
	google.golang.org/grpc v1.33.2 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/binance-chain/tss-lib => github.com/dawn-to-dusk/tss-lib v1.3.2-0.20220422023240-5ddc16a330ed

replace github.com/agl/ed25519 => github.com/binance-chain/edwards25519 v0.0.0-20200305024217-f36fc4b53d43
