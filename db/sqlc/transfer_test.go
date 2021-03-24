package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/aminuolawale/bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t * testing.T)Transfer{
	testAccount1 := CreateRandomAccount(t)
	testAccount2 := CreateRandomAccount(t)
	arg:= CreateTransferParams{
		FromAccountID: testAccount1.ID, 
		ToAccountID: testAccount2.ID, 
		Amount: util.RandomMoney(), 
	}
	entry, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry) 
	require.Equal(t, arg.FromAccountID, entry.FromAccountID) 
	require.Equal(t, arg.ToAccountID, entry.ToAccountID) 
	require.Equal(t, arg.Amount , entry.Amount) 

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T){
	arg := createRandomTransfer(t)
	entry, err := testQueries.GetTransfer(context.Background(),arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry) 
	require.Equal(t, arg.FromAccountID, entry.FromAccountID) 
	require.Equal(t, arg.ToAccountID, entry.ToAccountID)  
	require.Equal(t, arg.Amount , entry.Amount) 
	require.WithinDuration(t, arg.CreatedAt, entry.CreatedAt, time.Second)
}


func TestUpdateTransfer(t *testing.T){
	testTransfer:= createRandomTransfer(t)
	arg := UpdateTransferParams{ID:testTransfer.ID, Amount:util.RandomMoney()}
	entry, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry) 
	require.Equal(t, testTransfer.ID, entry.ID) 
	require.Equal(t, testTransfer.FromAccountID, entry.FromAccountID) 
	require.Equal(t, testTransfer.ToAccountID, entry.ToAccountID) 
	require.Equal(t, arg.Amount , entry.Amount) 
	require.WithinDuration(t, testTransfer.CreatedAt, entry.CreatedAt, time.Second)
}


func TestDeleteTransfer(t *testing.T){
	testTransfer:= createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), testTransfer.ID)
	require.NoError(t, err)	
	entry, err := testQueries.GetTransfer(context.Background(), testTransfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry)
}


func TestGetTransfers (t  *testing.T){
	for i:=0; i< 10; i++{
		createRandomTransfer(t)
	}
	arg := GetTransfersParams{
		Limit:5, Offset:5,
	}
	entries, err := testQueries.GetTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries{
		require.NotEmpty(t, entry)
	}
}