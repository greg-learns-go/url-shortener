package urls_db

import (
	"testing"
)

var testConn Connection = CreateConnection("file::memory:")

func prepareDb(t *testing.T) {
	testConn.EnsureDBExists()
	t.Cleanup(func() {
		testConn.Close()
	})
}

func TestInsertStoresEntry(t *testing.T) {
	prepareDb(t)
	longUrl := "http://www.example.com"
	testConn.Insert(longUrl, "ex1")

	retrievedLongUrl, er := testConn.Find("ex1")
	if er != nil {
		t.Fatalf("Find returned error: %v", er)
	}

	if longUrl != retrievedLongUrl {
		t.Errorf("Expected %v got %v", longUrl, retrievedLongUrl)
	}
}
