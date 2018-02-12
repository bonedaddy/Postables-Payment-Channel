from web3 import Web3
import json
import getpass

# micropayment
# keccak256(channelId, paymentString, withdrawalAmount)


def sigParts(h,v,r,s):
	print("h\t\t",h,"\nv\t\t",v,"\nr\t\t",r,"\ns\t\t",s,"\n")

def genSigParts(signature):
	h = Web3.toHex(signature.messageHash)
	v = signature.v
	# zero-pad with 32 bytes, issue #617 in web3py
	r = Web3.toHex(Web3.toBytes(signature.r).rjust(32, b'\0'))
	s = Web3.toHex(Web3.toBytes(signature.s).rjust(32, b'\0'))
	return (h, v, r, s)
w3 = Web3()

keyFilePath = input('Enter the absolute path to your key file:\t')
keyFilePassword = getpass.getpass('Enter the password to decrypt your key file:\t')
with open(keyFilePath, 'r') as fh:
	private_key = w3.eth.account.decrypt(json.load(fh), keyFilePassword)


print("Options")
choice = int(input("0\tSign message\n1\tsign micro pay\t\t"))

if choice == 0:
	messageTextToSign = input('What is the message you wish to sign:\t')
	hashStr = Web3.toHex(Web3.soliditySha3(['string'], [messageTextToSign]))
	signature = w3.eth.account.sign(message_hexstr=hashStr, private_key=private_key)
	h, v, r, s = genSigParts(signature)
	sigParts(h,v,r,s)
elif choice == 1:
	channelId = Web3.toHex(0x021e26844f552d2d1c76aac6d9324ae42f498b73a5c09b9871c9fea9e3ef9559)
	paymentId = int(input('Enter payment id\t'))
	withdrawalAmount = int(input('Enter withdrawal amount\t'))
	rawMessageHash = Web3.toHex(Web3.soliditySha3(['bytes32', 'uint256', 'uint256'], [channelId, paymentId, withdrawalAmount]))
	print(rawMessageHash)
	signature = w3.eth.account.sign(message_hexstr=rawMessageHash, private_key=private_key)
	h, v, r, s = genSigParts(signature)
	sigParts(h,v,r,s)
else:
	print('try again')
	exit(1)




