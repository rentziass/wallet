package wallet

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestNewWriter(t *testing.T) {
	p := &Pass{FormatVersion: 1}
	w, err := NewWriter(
		p,
		"wwdr",
		"cert",
		"key",
		"keypass",
	)

	if err != nil {
		t.Error("expected no error at this point")
	}

	if w.Pass != p {
		t.Error("expected to set provided pass as w.Pass")
	}

	if w.buffer == nil {
		t.Error("expected to initialize buffer")
	}

	if w.archive == nil {
		t.Error("expected to initialize zip archive")
	}

	if w.manifest == nil {
		t.Error("expected to initialize manifest map")
	}

	if w.manifest["pass.json"] != "24fc6eeef670572f5480d8dbc705e3ac6612061a" {
		t.Errorf("expected to SHA1 the contents of Pass on initializaton."+
			"Expected %s, got %s.",
			"24fc6eeef670572f5480d8dbc705e3ac6612061a",
			w.manifest["pass.json"])
	}
}

func TestPassWriter_AddFile(t *testing.T) {
	p := &Pass{FormatVersion: 1}
	w, err := NewWriter(
		p,
		"wwdr",
		"cert",
		"key",
		"keypass",
	)

	// ensure empty buffer
	w.buffer = new(bytes.Buffer)
	w.archive = zip.NewWriter(w.buffer)

	if err != nil {
		t.Error("expected no error at this point")
	}

	// requires a file name
	err = w.AddFile("", make([]byte, 0))
	if err == nil {
		t.Error("should return an error if no name is provided")
	}

	// adds every new file to the manifest
	iconBytes, err := ioutil.ReadFile("./test_data/icon.png")
	if err != nil {
		t.Fatal("could not open file: ./test_data/icon.png")
	}

	err = w.AddFile("icon.png", iconBytes)
	if err != nil {
		t.Error("error should have been nil")
	}

	sha, present := w.manifest["icon.png"]
	if !present {
		t.Error("icon.png should have been added to the manifest")
	}
	if sha != "033536cd22a2a00416d8540ff26586a27505ef03" {
		t.Errorf("expected the SHA1 for icon.png to be %s, got %s",
			"033536cd22a2a00416d8540ff26586a27505ef03",
			sha)
	}

	// adds every new file to the writer of the zip archive
	// and that writer writes to w.buffer
	buf := new(bytes.Buffer)
	archive := zip.NewWriter(buf)
	zipw, err := archive.Create("icon.png")
	if err != nil {
		t.Fatal(err)
	}

	_, err = zipw.Write(iconBytes)
	if err != nil {
		t.Fatal(err)
	}

	archive.Close()
	w.archive.Close()

	if bytes.Compare(buf.Bytes(), w.buffer.Bytes()) != 0 {
		t.Error("should have written archive bytes to w.buffer")
	}
}

func TestPassWriter_Close(t *testing.T) {
	//p := &Pass{FormatVersion: 1}
	//w, err := NewWriter(p)
	//if err != nil {
	//	t.Error("expected no error at this point")
	//}

	// returns a bytes.Buffer of zip file
	// containing the manifest
	// containing the signature file
}
