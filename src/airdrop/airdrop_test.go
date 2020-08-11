package airdrop

import (
	"context"
	"fmt"
	"testing"

	"github.com/bonedaddy/paymentchannels/testutils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/postables/Postables-Payment-Channel/src/airdrop/bindings"
	"github.com/stretchr/testify/require"
)

func TestAirdrop(t *testing.T) {
	auth, backend := testutils.NewBlockchain(t)

	addr, tx, ct, err := bindings.DeployAirdrop(auth, backend)
	backend.Commit()
	require.NoError(t, err)

	_, err = bind.WaitDeployed(context.Background(), backend, tx)

	require.NoError(t, err)

	fmt.Println("aidrdrop contract: ", addr.String())

	_ = ct

	defer require.NoError(t, backend.Close())
}
