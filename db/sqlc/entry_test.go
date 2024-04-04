package sqlc

import (
	"context"
	"github.com/bigusef/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	params := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(100, 5000),
	}

	entry, err := testQueries.CreateEntry(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	require.NotZero(t, entry.ID)
	require.NotEmpty(t, entry.AccountID)
	require.NotEmpty(t, entry.Amount)
	require.NotZero(t, entry.CreatedAt)
}
