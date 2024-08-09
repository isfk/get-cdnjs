package cdnjs

import (
	"testing"
)

// https://api.cdnjs.com/libraries/{:library}
// https://api.cdnjs.com/libraries/jquery

// https://api.cdnjs.com/libraries/{:library}/{:version}
// https://api.cdnjs.com/libraries/jquery/3.5.1

// https://cdnjs.cloudflare.com/ajax/libs/{:library}/{:version}/{:file}

func TestVersions(t *testing.T) {
	ret := &VersionsRet{}
	url := "https://api.cdnjs.com/libraries/jquery"
	_, err := Get[VersionsRet](url, "http://127.0.0.1:7897", ret)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(ret)
}

func TestFiles(t *testing.T) {
	ret := &FilesRet{}
	url := "https://api.cdnjs.com/libraries/jquery/3.7.1"
	_, err := Get[FilesRet](url, "http://127.0.0.1:7897", ret)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(ret)
}
