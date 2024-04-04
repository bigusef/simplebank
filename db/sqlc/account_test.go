package sqlc

import (
	"context"
	"database/sql"
	"github.com/bigusef/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.Owner, args.Owner)
	require.Equal(t, acc.Balance, args.Balance)
	require.Equal(t, acc.Currency, args.Currency)

	return acc
}

func TestQueries_CreateAccount(t *testing.T) {
	account := createRandomAccount(t)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestQueries_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := UpdateAccountParams{ID: account.ID, Balance: 500}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.Equal(t, updatedAccount.Balance, args.Balance)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestQueries_ListAccount(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomAccount(t)
	}

	list_params := ListAccountParams{5, 0}

	accounts, err := testQueries.ListAccount(context.Background(), list_params)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
