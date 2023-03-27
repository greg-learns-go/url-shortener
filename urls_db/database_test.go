package urls_db

import (
	"testing"
)

func prepareDb(t *testing.T) Connection {
	var testConn Connection = CreateConnection("file::memory:")
	testConn.EnsureDBExists()
	t.Cleanup(func() {
		testConn.Close()
	})
	return testConn
}

func TestInsertStoresEntry(t *testing.T) {
	testConn := prepareDb(t)
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

func TestFindOrInsertNewUrl(t *testing.T) {
	testConn := prepareDb(t)

	entry, er := testConn.FindOrInsert("http://example.com")

	if er != nil {
		t.Fatalf("FindOrInsert returned error: %v", er)
	}

	if entry.LongUrl != "http://example.com" {
		t.Errorf("Returned entry data don't match, expected\nhttp://example.com got\n%s", entry.LongUrl)
	}
}

func TestFindOrInsertExisti(t *testing.T) {
	testConn := prepareDb(t)

	testConn.Insert("example.com", "leg")

	beforeCount := countAllRows(testConn)

	_, er := testConn.FindOrInsert("example.com")
	if er != nil {
		t.Fatalf("FindOrInsert returned error: %v", er)
	}

	afterCount := countAllRows(testConn)

	if afterCount-beforeCount != 0 {
		t.Errorf("FindOrInsert changed nymber of rows by %v, expected 0", afterCount-beforeCount)
	}

}

func countAllRows(conn Connection) int {
	var count int
	res := conn.Db.QueryRow("SELECT COUNT(*) FROM urls", nil)
	if er := res.Scan(&count); er != nil {
		panic(er)
	}
	return count
}
