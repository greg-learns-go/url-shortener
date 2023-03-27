package urls_db

import (
	"testing"

	"github.com/greg-learns-go/url-shortener/shortener"
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

	entry, er := testConn.FindOrInsert("http://example.com", shortener.New())

	if er != nil {
		t.Fatalf("FindOrInsert returned error: %v", er)
	}

	if entry.LongUrl != "http://example.com" {
		t.Errorf("Returned entry data don't match, expected\nhttp://example.com got\n%s", entry.LongUrl)
	}
}

func TestFindOrInsertExisting(t *testing.T) {
	testConn := prepareDb(t)

	testConn.Insert("example.com", "leg")

	beforeCount := countAllRows(testConn)

	_, er := testConn.FindOrInsert("example.com", shortener.New())
	if er != nil {
		t.Fatalf("FindOrInsert returned error: %v", er)
	}

	afterCount := countAllRows(testConn)

	if afterCount-beforeCount != 0 {
		t.Errorf("FindOrInsert changed number of rows by %v, expected 0", afterCount-beforeCount)
	}
}
func TestFindOrInsertExistingConflict(t *testing.T) {
	testConn := prepareDb(t)

	testConn.Insert("example.com", "sHOrT")
	beforeCount := countAllRows(testConn)

	_, er := testConn.FindOrInsert("other.example.com", NewMockShortener("sHOrT"))
	if er == nil {
		t.Error("expected FindOrInsert to return an error")
	}

	expectedMessage := "URL other.example.com is in conflict with example.com, both produce sHOrT"
	if er.Error() != expectedMessage {
		t.Errorf(
			"Expected FindOrInsert to return an error with message \n%v, got \n%v",
			expectedMessage,
			er.Error(),
		)
	}

	afterCount := countAllRows(testConn)

	if afterCount-beforeCount != 0 {
		t.Errorf("FindOrInsert changed number of rows by %v, expected 0", afterCount-beforeCount)
	}
}

type mockShortener struct {
	Shorted string
}

func NewMockShortener(short string) Shortener {
	return mockShortener{Shorted: short}
}

func (s mockShortener) Shorten(url string) string {
	return s.Shorted
}

func countAllRows(conn Connection) int {
	var count int
	res := conn.Db.QueryRow("SELECT COUNT(*) FROM urls", nil)
	if er := res.Scan(&count); er != nil {
		panic(er)
	}
	return count
}
