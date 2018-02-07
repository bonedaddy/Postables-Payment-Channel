from web3 import Web3, IPCProvider
import json
import getpass

gethIpcPath = input('Enter the absolute path for your geth ipc file')
w3 = Web3(IPCProvider(gethIpcPath))

keyFilePath = input('Enter the absolute path to your key file')
keyFilePassword = getpass.getpass('Enter the password to decrypt your key file')
with open(keyFilePath, 'r') as fh:
	private_key = w3.eth.account.decrypt(json.load(fh), keyfilePassword)


messageTextToSign = input('What is the message you wish to sign')
signature = w3.eth.account.sign(message_text=messageTextToSign, private_key=private_key)
h = Web3.toHex(signature.messageHash)
v = signature.v
# zero-pad with 32 bytes, issue #617 in web3py
r = Web3.toHex(Web3.toBytes(signature.r).rjust(32, b'\0'))
s = Web3.toHex(Web3.toBytes(signature.s).rjust(32, b'\0'))



# make sure we checksum the address
contractAddress = Web3.toChecksumAddress(input('Enter the Postables Channel Hub Contract Address:'))
abiFilePath = input('Enter the absolute path to the Postbales Channel Hub Contract ABI file:')
with open(abiFilePath, 'r') as fh:
	abi = json.load(fh)

# load the contract
contract = w3.eth.contract(contractAddress, abi=abi)

choice = 2
while True:
	choice = input('0\tsubmit purchaser signature\n1\tsubmit vendor signature\n')
	if choice not in [0, 1]:
		print('Invalid choice, try again')
	else:
		break

if choice == 0: # purchaser
	pass
elif choice == 1: # vendor
	channelId = Web3.toHex(Web3.toBytes(input('enter channel id')))
	contract.submitVendorProof(h, v, r, s, channelId)
else:
	print('?')
	exit(9000)