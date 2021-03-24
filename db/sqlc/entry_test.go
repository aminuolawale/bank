package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/aminuolawale/bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t * testing.T)Entry{
	testAccount := CreateRandomAccount(t)
	arg:= CreateEntryParams{
		AccountID: testAccount.ID, 
		Amount: util.RandomMoney(), 
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry) 
	require.Equal(t, arg.AccountID, entry.AccountID) 
	require.Equal(t, arg.Amount , entry.Amount) 

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T){
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T){
	arg := createRandomEntry(t)
	entry, err := testQueries.GetEntry(context.Background(),arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry) 
	require.Equal(t, arg.AccountID, entry.AccountID) 
	require.Equal(t, arg.Amount , entry.Amount) 
	require.WithinDuration(t, arg.CreatedAt, entry.CreatedAt, time.Second)
}


func TestUpdateEntry(t *testing.T){
	testEntry:= createRandomEntry(t)
	arg := UpdateEntryParams{ID:testEntry.ID, Amount:util.RandomMoney()}
	entry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry) 
	require.Equal(t, testEntry.ID, entry.ID) 
	require.Equal(t, testEntry.AccountID, entry.AccountID) 
	require.Equal(t, arg.Amount , entry.Amount) 
	require.WithinDuration(t, testEntry.CreatedAt, entry.CreatedAt, time.Second)
}


func TestDeleteEntry(t *testing.T){
	testEntry:= createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), testEntry.ID)
	require.NoError(t, err)	
	entry, err := testQueries.GetEntry(context.Background(), testEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry)
}


func TestGetEntries (t  *testing.T){
	for i:=0; i< 10; i++{
		createRandomEntry(t)
	}
	arg := GetEntriesParams{
		Limit:5, Offset:5,
	}
	entries, err := testQueries.GetEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries{
		require.NotEmpty(t, entry)
	}
}