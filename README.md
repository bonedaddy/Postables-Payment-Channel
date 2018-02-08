# Postables-Payment-Channel (WIP)

# Introduction

Postables Payment Channels is a set of payment channel contract whichs can be reused across multiple different transactions, and parties. Currently payment channels are limited to two participating entities. It is still very much so a work in progress and will be updated over time to include a bridge written in python.

# Roadmap

- Add reusable channels
- Add micro payments
- Add multi-party channels
- Add ERC20 channels
- Add a terminal based client



# Contribution guides:
TO DO


# Bugs and Fixes:

## Improper  address returned from signature verification
```
>>> private
b'\xd8Y\xe3T\xb8\x15\xaaJ\xe6i\xaf\xf9\x02\x8a\xbf\x89#\xf9D\x03\xdd\tc\xf2\xed\xc3\x9eQ\r\x98\xb7|'
>>> signature = w3.eth.account.sign(message_text="hello", private_key=private) 

AttributeDict({'message': HexBytes('0x68656c6c6f'), 'messageHash': HexBytes('0x50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750'), 'r': 53665784189267139251820747046956462498355115993943239404465477346598658281299, 's': 6629064374810355767282561830725699037469490635953226814818888019755260187756, 'v': 27, 'signature': HexBytes('0x76a5c1e7f682df3375a2bdd6f72ad2171b0cf826fc8a8a3209c33a4e57e88f530ea7eadf8603ba2c4a5c5006571be7665020812aa7e403a614587cfe7a18146c1b')})
>>> h = Web3.toHex(signature.messageHash)
>>> v = signature.v
>>> r = Web3.toHex(signature.r)
>>> s = Web3.toHex(signature.s)
>>> h
'0x50b2c43fd39106bafbba0da34fc430e1f91e3c96ea2acee2bc34119f92b37750'
>>> v
27
>>> r
'0x76a5c1e7f682df3375a2bdd6f72ad2171b0cf826fc8a8a3209c33a4e57e88f53'
>>> s
'0xea7eadf8603ba2c4a5c5006571be7665020812aa7e403a614587cfe7a18146c'
>>> Web3.toHex(Web3.toBytes(signature.s))
'0x0ea7eadf8603ba2c4a5c5006571be7665020812aa7e403a614587cfe7a18146c' (returns the correct address)
```
