package sqlstore

import (
	"os"
	"testing"

	"github.com/mattermost/focalboard/server/services/store"
	"github.com/mattermost/focalboard/server/services/store/storetests"
	"github.com/stretchr/testify/require"
)

func SetupTests(t *testing.T) (store.Store, func()) {
	dbType := os.Getenv("FB_STORE_TEST_DB_TYPE")
	if dbType == "" {
		dbType = sqliteDBType
	}

	connectionString := os.Getenv("FB_STORE_TEST_CONN_STRING")
	if connectionString == "" {
		connectionString = ":memory:"
	}

	store, err := New(dbType, connectionString, "test_")
	require.Nil(t, err)

	tearDown := func() {
		err = store.Shutdown()
		require.Nil(t, err)
	}

	return store, tearDown
}

func TestBlocksStore(t *testing.T) {
	t.Run("BlocksStore", func(t *testing.T) { storetests.StoreTestBlocksStore(t, SetupTests) })
	t.Run("SharingStore", func(t *testing.T) { storetests.StoreTestSharingStore(t, SetupTests) })
}
