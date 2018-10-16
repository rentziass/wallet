package wallet

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os/exec"
)

func hashSHA1ForReader(r io.Reader) (string, error) {
	if r == nil {
		return "", errors.New("nil reader provided")
	}

	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func signFile(in []byte, wwdrPath, certPath, keyPath, password string) ([]byte, error) {
	if in == nil {
		return nil, errors.New("reader is nil")
	}

	if len(certPath) == 0 {
		return nil, errors.New("certPath is required for signature")
	}

	if len(keyPath) == 0 {
		return nil, errors.New("keyPath is required for signature")
	}

	return openssl(
		in,
		"smime",
		"-binary",
		"-sign",
		"-certfile", wwdrPath,
		"-signer", certPath,
		"-inkey", keyPath,
		"-outform", "DER",
		"-passin", "pass:"+password)
}

func openssl(stdin []byte, args ...string) ([]byte, error) {
	cmd := exec.Command("openssl", args...)

	in := bytes.NewReader(stdin)
	out := &bytes.Buffer{}
	errs := &bytes.Buffer{}

	cmd.Stdin, cmd.Stdout, cmd.Stderr = in, out, errs

	if err := cmd.Run(); err != nil {
		if len(errs.Bytes()) > 0 {
			return nil, fmt.Errorf("error running %s (%s):\n %v", cmd.Args, err, errs.String())
		}
		return nil, err
	}

	return out.Bytes(), nil
}
