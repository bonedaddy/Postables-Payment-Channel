# Postables-Payment-Channel

> **warning: none of this code is audited use at your own risk**

Postables-Payment-Channel is a collection of easy to use smart contracts that can be used to facilitate arbitrary ERC20 and ETH payment channels between two parties, as well as an extremely cheap method of distributing airdrops using signed messages instead of mass transaction broadcasting called "AirDropChannels". 

The payment channels are robust allowing for any ERC20 to be used, along with ETH itself. The channels support micro payments allowing a party to withdraw micro amounts from the channels available balance. This is useful in situations where the two parties may want the trust that comes with payment channels, but don't want to end the channel by withdrawing funds from it as is typical with the standard payment channel examples. This can be very useful if you are paying a contractor, and the contractor wants the assurance that all the funds they are being promised actually exist, but you don't want to pay them for work they haven't done yet.

AirDropChannels are a slightly modified payment channel concept, but instead of one-to-one unidirectional channels, it is a one-to-many unidirectional channel, allowing any number of parties to withdraw tokens from the channel! This can be used to facilitate stupidly cheap airdrops! Gone are the days of airdrops costing the token develop tens of thousands of dollars. One of the benefits to using this method of airdrops, is that unclaimed tokens are refunded back to the token developers wallet as soon as they close the channel! No  more wasted money from tokens floating around in unused addresses for ethernity!

In addition golang, as well as **very** out of date python bindings exist to interact with the contracts using go-ethereum (go) and web3py (python).

# Files

- solidity/AirDropChannels.sol (modified payment channel allowing for cheap airdrops)
- solidity/ChannelsV3.sol (base channel setup, no micropayment)
- solidity/ChannelsV4.sol (base channel setup, micropayments)
- solidity/ChannelsV5.sol (WIP)

# Bugs and Fixes

## Improper  address returned from signature verification using python bindings

> **note: this uses a very out of date version of web3py **

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
