# Postables-Payment-Channel (WIP) --- DEVELOP

# Introduction

Postables Payment Channels is a set of payment channel contract whichs can be reused across multiple different transactions, and parties. Currently payment channels are limited to two participating entities. It is still very much so a work in progress and will be improved over time to include my currently closed-source EVM bridge to enable payment channels across side-chains. There will be a fee charged each time a channel is opened up to deter malicious channel openers/abuse. In future releases each channel will be re-usable between the same party to avoid having to pay multiple channel opening fees. This will then be migrated to a "user acccount" system to establish a level of trust so that repeated channel openers are exempted from any and all fees.

*NOTE* The fee system isn't totally solidified yet and may very well be abandoned in favour for another system to deter abuse that is more fair to everyone and doesn't impose usage fees.

# Who is this for?

These payment channels are suitable for any transaction involving two parties, but as you may have guessed from some of the variable names, it is primarily intended for businesses to pay their employees out. For example if you want to pay your employees every 2 weeks, but want to give them the option to withdraw their pay every 24 hours, without having to pay for multiple transactions (ie waiting 2-3 days and doing a single withdrawal for the total amount of 2-3 days), or to act as a sort of escrow for your employee to have ensurance that they will be paid

# Roadmap ()

1) Introduce the ability for rolling over daily incremental withdrawals
2) Introduce re-usable channels
3) Introduce channel factories to allow  for the existence of multiple PostableHubs (ChannelsV2.sol), each being owned by a single entity, so that each Hub could be used by a business to pay a particular department and so on.
4) Introduce Postables-Bridge to allow for payment channels between mainnet and other EVM chains (side chains)
	> Use PPCT (Postables Payment Channel Tokens) as the collateral mechanism (or ether, design is open at time point)
5) Add web3 desktop or browser Dapp for ease of  use

# Contribution guides:
TO DO
