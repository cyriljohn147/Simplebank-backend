package db

import (
	"context"
	"testing"
	"time"

	"github.com/cyriljohn147/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	en := CreateRandomEntry(t)
	entry, err := testQueries.GetEntry(context.Background(), en.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, en.ID, entry.ID)
	require.Equal(t, en.AccountID, entry.AccountID)
	require.Equal(t, en.Amount, entry.Amount)
	require.WithinDuration(t, en.CreatedAt, entry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for range 10 {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
