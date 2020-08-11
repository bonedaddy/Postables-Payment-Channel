.PHONY: gen-bindings
gen-bindings:
	abigen --pkg payment --abi bin/solidity/PaymentChannels.abi --bin bin/solidity/PaymentChannels.bin --out bindings/payment/payment_channels.go
	abigen --pkg airdrop --abi bin/solidity/AirdropChannels.abi --bin bin/solidity/AirdropChannels.bin --out bindings/airdrop/aidrop_channels.go