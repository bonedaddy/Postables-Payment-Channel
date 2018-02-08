from web3 import Web3
import json
import getpass

w3 = Web3()

keyFilePath = input('Enter the absolute path to your key file:\t')
keyFilePassword = getpass.getpass('Enter the password to decrypt your key file:\t')
with open(keyFilePath, 'r') as fh:
	private_key = w3.eth.account.decrypt(json.load(fh), keyFilePassword)


messageTextToSign = input('What is the message you wish to sign:\t')
hashStr = Web3.toHex(Web3.soliditySha3(['string'], [messageTextToSign]))
signature = w3.eth.account.sign(message_hexstr=hashStr, private_key=private_key)
h = Web3.toHex(signature.messageHash)
v = signature.v
# zero-pad with 32 bytes, issue #617 in web3py
r = Web3.toHex(Web3.toBytes(signature.r).rjust(32, b'\0'))
s = Web3.toHex(Web3.toBytes(signature.s).rjust(32, b'\0'))

print("h\t\t",h,"\nv\t\t",v,"\nr\t\t",r,"\ns\t\t",s,"\n")