package keyring

import (
	"testing"

	"github.com/stretchr/testify/require"

	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBech32KeysOutput(t *testing.T) {
	// TODO - update this test to not use bech32

	tmpKey := secp256k1.GenPrivKey().PubKey()
	tmpAddr := sdk.AccAddress(tmpKey.Address().Bytes())

	multisigPks := kmultisig.NewLegacyAminoPubKey(1, []types.PubKey{tmpKey})
	multiInfo := NewMultiInfo("multisig", multisigPks)
	accAddr := sdk.AccAddress(multiInfo.GetPubKey().Address().Bytes())
	require.True(t, accAddr.Equals(tmpAddr))

	expectedOutput, err := NewKeyOutput(multiInfo.GetName(), multiInfo.GetType(), accAddr, tmpKey)
	require.NoError(t, err)
	expectedOutput.Threshold = 1
	expectedOutput.PubKeys = []multisigPubKeyOutput{{tmpAddr.String(), bechTmpKey, 1}}

	outputs, err := Bech32KeysOutput([]Info{multiInfo})
	require.NoError(t, err)
	require.Equal(t, expectedOutput, outputs[0])
}
