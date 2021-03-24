package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/aminuolawale/bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t * testing.T)Account{
	arg:= CreateAccountParams{
		Owner: util.RandomOwner(), 
		Balance: util.RandomMoney(), 
		Currency: util.RandomCurrency(), 
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account) 
	require.Equal(t, arg.Owner, account.Owner) 
	require.Equal(t, arg.Balance , account.Balance) 
	require.Equal(t, arg.Currency, account.Currency) 

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T){
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T){
	arg := CreateRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(),arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account) 
	require.Equal(t, arg.Owner, account.Owner) 
	require.Equal(t, arg.Balance , account.Balance) 
	require.Equal(t, arg.Currency, account.Currency) 
	require.WithinDuration(t, arg.CreatedAt, account.CreatedAt, time.Second)
}


func TestUpdateAccount(t *testing.T){
	testAccount:= CreateRandomAccount(t)
	arg := UpdateAccountParams{ID:testAccount.ID, Balance:util.RandomMoney()}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account) 
	require.Equal(t, testAccount.ID, account.ID) 
	require.Equal(t, arg.Balance, account.Balance) 
	require.Equal(t, testAccount.Owner, account.Owner) 
	require.Equal(t, testAccount.Currency, account.Currency)
	require.WithinDuration(t, testAccount.CreatedAt, account.CreatedAt, time.Second)
}


func TestDeleteAccount(t *testing.T){
	testAccount:= CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), testAccount.ID)
	require.NoError(t, err)	
	account, err := testQueries.GetAccount(context.Background(), testAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}


func TestGetAccounts (t  *testing.T){
	for i:=0; i< 10; i++{
		CreateRandomAccount(t)
	}
	arg := GetAccountsParams{
		Limit:5, Offset:5,
	}
	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts{
		require.NotEmpty(t, account)
	}
}