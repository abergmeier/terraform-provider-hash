package datasources

import (
	"testing"
)

func TestGenerateSha256OfFiles(t *testing.T) {

	hash, err := generateSha256OfFiles([]string{
		"testdata/testme.tst",
		"testdata/foo.bar",
	})
	if err != nil {
		t.Fatal(err)
	}
	if hash != "09fd542731bded7a10813ce214447b685bfbc247546f1c4f363c9282c7d3a025" {
		t.Fatal("Expected hash 09fd542731bded7a10813ce214447b685bfbc247546f1c4f363c9282c7d3a025", "Received hash:", hash)
	}
}
