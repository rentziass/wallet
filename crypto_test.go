package wallet

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestHashSHA1ForReader(t *testing.T) {
	s, err := hashSHA1ForReader(nil)
	if len(s) != 0 {
		t.Error("should return an empty string when providing a nil reader")
	}
	if err == nil {
		t.Error("should return an error when providing an empty string")
	}

	testCases := map[string]string{
		"icon.png":    "033536cd22a2a00416d8540ff26586a27505ef03",
		"icon@2x.png": "4c8845fc147c69f078d59ce9f0f2aba125e505e8",
		"logo.png":    "78959a9f8d490fe25d5bd75afb1f8fbc4e4f8a34",
		"logo@2x.png": "148f4d2af9c8965d3d4be56ee3aa9e35a691597d",
	}

	for filename, expectedOutput := range testCases {
		f, err := os.Open("./test_data/" + filename)
		if err != nil {
			t.Fatalf("failed to open file (%s): %s", filename, err)
		}

		out, err := hashSHA1ForReader(f)
		if err != nil {
			t.Errorf("no error expected, got: %s", err)
		}

		if out != expectedOutput {
			t.Errorf("expected output to be\n"+
				"%s,\t got %s", expectedOutput, out)
		}

		f.Close()
	}
}

func TestSignFile(t *testing.T) {
	_, err := signFile(nil, "", "", "", "")
	if err == nil {
		t.Error("expected error if reader is nil")
	}

	fileToSign, err := ioutil.ReadFile("./test_data/signature/file_to_sign")
	if err != nil {
		t.Fatal(err)
	}

	_, err = signFile(fileToSign, "", "./test_data/certs/cert.pem", "./test_data/certs/key.pem", "")
	if err == nil {
		t.Error("expected error if wwdrPath is empty")
	}

	_, err = signFile(fileToSign, "./test_data/certs/WWDR.pem", "", "./test_data/certs/key.pem", "")
	if err == nil {
		t.Error("expected error if certPath is empty")
	}

	_, err = signFile(fileToSign, "./test_data/certs/WWDR.pem", "./test_data/certs/cert.pem", "", "")
	if err == nil {
		t.Error("expected error if keyPath is empty")
	}

	_, err = signFile(fileToSign, "./test_data/certs/WWDR.pem", "./test_data/certs/cert.pem", "./test_data/certs/key.pem", "")
	if err != nil {
		t.Errorf("did not expect an error at this point, got %s", err)
	}
}
