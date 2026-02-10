package cdnjs

import (
	"testing"
)

// https://data.jsdelivr.com/v1/package/npm/{:package}
// https://data.jsdelivr.com/v1/package/npm/jquery

// https://data.jsdelivr.com/v1/package/npm/{:package}@{:version}
// https://data.jsdelivr.com/v1/package/npm/jquery@3.7.1

// https://cdn.jsdelivr.net/npm/{:package}@{:version}/{:file}

func TestVersions(t *testing.T) {
	ret := &JSDelivrVersionsRet{}
	url := "https://data.jsdelivr.com/v1/package/npm/jquery"
	_, err := Get[JSDelivrVersionsRet](url, "http://127.0.0.1:7897", ret)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(ret)
}

func TestFiles(t *testing.T) {
	ret := &JSDelivrFilesRet{}
	url := "https://data.jsdelivr.com/v1/package/npm/jquery@3.7.1"
	_, err := Get[JSDelivrFilesRet](url, "http://127.0.0.1:7897", ret)
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Log(ret)
}
