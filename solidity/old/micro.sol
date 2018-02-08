	function submitMicroPaymentProof(
		bytes32 _channelId,
		bytes32 _h,
		uint8   _v,
		bytes32 _r,
		bytes32 _s,
		uint256 _paymentId,
		uint256 _withdrawalAmount)
		public
		returns (bool)
	{
		// validate channel id
		require(channelIds[_channelId]);
		// validate channel state
		require(channels[_channelId].state == ChannelStates.opened);
		// make sure nobody else is trying to withdraw the funds
		require(msg.sender == channels[_channelId].destination);
		// prevent a micropayment from reducing the entire balance
		require(channels[_channelId].value > _withdrawalAmount && _withdrawalAmount > 0);
		// following two lines construct the proof, with prefix to validate _h
		bytes32  _proof = keccak256(_channelId, _paymentId, _withdrawalAmount);
		bytes32 proof = keccak256(prefix, _proof);
		// validate the proof, if it fails most likely malicious submitter, so lets waste their gas ;)
		assert(proof == _h);
		// make sure the proof hasn't already bee submitted
		require(!microPaymentHashes[_channelId][_h]);
		// mark proof as submitted
		microPaymentHashes[_channelId][_h] = true;
		// time to recover the signature
		address signer = ecrecover(_h, _v, _r, _s);
		// make sure its from the source, otherwise someones up to no good so lets waste their gas ;)
		assert(signer == channels[_channelId].source);
		// calculate remaining channel value
		uint256 remainingChannelValue = channels[_channelId].value.sub(_withdrawalAmount);
		// adjust channel value
		channels[_channelId].value = remainingChannelValue;
		// notify blockchain
		MicroPaymentWithdrawn(_channelId, _withdrawalAmount, remainingChannelValue);
		// withdraw funds
		msg.sender.transfer(_withdrawalAmount);
		return true;
	}