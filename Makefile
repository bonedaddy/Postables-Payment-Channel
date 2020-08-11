.PHONY: gen-bindings
gen-bindings:
	abigen --pkg bindings --abi bin/solidity/PaymentChannels.abi --bin bin/solidity/PaymentChannels.bin --out bindings/payment_channels.go