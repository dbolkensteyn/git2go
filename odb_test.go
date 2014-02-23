package git

import (
	"io"
	"os"
	"testing"
)

func TestOdbStream(t *testing.T) {
	repo := createTestRepo(t)
	defer os.RemoveAll(repo.Workdir())
	_, _ = seedTestRepo(t, repo)

	odb, error := repo.Odb()
	checkFatal(t, error)

	str := "hello, world!"

	stream, error := odb.NewWriteStream(len(str), ObjectBlob)
	checkFatal(t, error)
	n, error := io.WriteString(stream, str)
	checkFatal(t, error)
	if n != len(str) {
		t.Fatalf("Bad write length %v != %v", n, len(str))
	}

	error = stream.Close()
	checkFatal(t, error)

	expectedId, error := NewOidFromString("30f51a3fba5274d53522d0f19748456974647b4f")
	checkFatal(t, error)
	if stream.Id.Cmp(expectedId) != 0 {
		t.Fatal("Wrong data written")
	}
}

func TestOdbHash(t *testing.T) {

    repo := createTestRepo(t)
	defer os.RemoveAll(repo.Workdir())
	_, _ = seedTestRepo(t, repo)

	odb, error := repo.Odb()
	checkFatal(t, error)

	str := `tree 115fcae49287c82eb55bb275cbbd4556fbed72b7
parent 66e1c476199ebcd3e304659992233132c5a52c6c
author John Doe <john@doe.com> 1390682018 +0000
committer John Doe <john@doe.com> 1390682018 +0000

Initial commit.`;

	oid, error := odb.Hash([]byte(str), ObjectCommit)
	checkFatal(t, error)

	coid, error := odb.Write([]byte(str), ObjectCommit)
	checkFatal(t, error)

	if oid.Cmp(coid) != 0 {
		t.Fatal("Hash and write Oids are different")
	}
}