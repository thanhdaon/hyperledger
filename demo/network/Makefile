SHELL := /bin/bash
PATH := $(PWD)/bin:$(PATH)

gen-genesis-block: export FABRIC_CFG_PATH=$(PWD)/devconfig
gen-genesis-block:
	configtxgen -profile SampleDevModeSolo \
		-channelID syschannel \
		-outputBlock genesisblock \
		-configPath $(FABRIC_CFG_PATH) \
		-outputBlock "$(PWD)/devconfig/genesisblock"

start-orderer: export FABRIC_CFG_PATH=$(PWD)/devconfig
start-orderer: export ORDERER_GENERAL_GENESISPROFILE=SampleDevModeSolo
start-orderer:
	orderer

start-peer: export FABRIC_CFG_PATH=$(PWD)/devconfig
start-peer: export FABRIC_LOGGING_SPEC=chaincode=debug
start-peer: export CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:7052
start-peer: export CORE_OPERATIONS_LISTENADDRESS=0.0.0.0:9444
start-peer:
	peer node start --peer-chaincodedev=true

create-channel: export FABRIC_CFG_PATH=$(PWD)/devconfig
create-channel:
	configtxgen \
		-channelID mychannel \
		-outputCreateChannelTx $(FABRIC_CFG_PATH)/mychannel.tx \
		-profile SampleSingleMSPChannel \
		-configPath $(FABRIC_CFG_PATH)

	peer channel create \
		-o 127.0.0.1:7050 \
		-c mychannel \
		-f "$(FABRIC_CFG_PATH)/mychannel.tx"

join-peer: export FABRIC_CFG_PATH=$(PWD)/devconfig
join-peer:
	peer channel join -b $(PWD)/mychannel.block

approve-amd-commit-chaincode-definition: export FABRIC_CFG_PATH=$(PWD)/devconfig
approve-amd-commit-chaincode-definition:
	peer lifecycle chaincode approveformyorg \
		-o 127.0.0.1:7050 \
		--channelID mychannel \
		--name mycc \
		--version 1.0 \
		--sequence 1 \
		--init-required \
		--signature-policy "OR ('SampleOrg.member')" \
		--package-id mycc:1.0

	peer lifecycle chaincode checkcommitreadiness \
		-o 127.0.0.1:7050 \
		--channelID mychannel \
		--name mycc \
		--version 1.0 \
		--sequence 1 \
		--init-required \
		--signature-policy "OR ('SampleOrg.member')"

	peer lifecycle chaincode commit \
		-o 127.0.0.1:7050 \
		--channelID mychannel \
		--name mycc \
		--version 1.0 \
		--sequence 1 \
		--init-required \
		--signature-policy "OR ('SampleOrg.member')" \
		--peerAddresses 127.0.0.1:7051

chaincode-invoke: export FABRIC_CFG_PATH=$(PWD)/devconfig
chaincode-invoke:
	peer chaincode invoke \
		--orderer localhost:7050 \
		--peerAddresses localhost:7051 \
		--channelID mychannel \
		--name mycc \
		--ctor '$(ctor)' 

chaincode-query: export FABRIC_CFG_PATH=$(PWD)/devconfig
chaincode-query:
	peer chaincode query \
		--orderer localhost:7050 \
		--peerAddresses localhost:7051 \
		--channelID mychannel \
		--name mycc \
		--ctor '$(ctor)'


clear-dev-env:
	rm -rf /tmp/hyperledger
	rm -f $(PWD)/devconfig/genesisblock
	rm -f $(PWD)/devconfig/mychannel.block
	rm -f $(PWD)/devconfig/mychannel.tx

demo: export FABRIC_CFG_PATH=$(PWD)/devconfig
demo:
	peer node rollback \
		--channelID mychannel \
		--blockNumber 2