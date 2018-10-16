package wallet

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
)

type PassWriter struct {
	Pass                                     *Pass
	manifest                                 map[string]string
	buffer                                   *bytes.Buffer
	archive                                  *zip.Writer
	wwdrPath, certPath, keyPath, keyPassword string
}

func NewWriter(pass *Pass, wwdrPath, certPath, keyPass, keyPassword string) (*PassWriter, error) {
	w := &PassWriter{
		Pass:        pass,
		manifest:    make(map[string]string),
		buffer:      new(bytes.Buffer),
		wwdrPath:    wwdrPath,
		certPath:    certPath,
		keyPath:     keyPass,
		keyPassword: keyPassword,
	}

	w.archive = zip.NewWriter(w.buffer)

	passBytes, err := json.Marshal(pass)
	if err != nil {
		return nil, err
	}

	err = w.AddFile("pass.json", passBytes)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *PassWriter) AddFile(name string, b []byte) error {
	if len(name) == 0 {
		return errors.New("a file name is required")
	}

	h, err := hashSHA1ForReader(bytes.NewReader(b))
	if err != nil {
		return err
	}

	w.manifest[name] = h

	zipw, err := w.archive.Create(name)
	if err != nil {
		return err
	}

	_, err = zipw.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (w *PassWriter) Close() (*bytes.Buffer, error) {
	manifestBytes, err := json.Marshal(w.manifest)
	if err != nil {
		return nil, err
	}

	err = w.AddFile("manifest.json", manifestBytes)
	if err != nil {
		return nil, err
	}

	signature, err := signFile(
		manifestBytes,
		w.wwdrPath,
		w.certPath,
		w.keyPath,
		w.keyPassword,
	)
	if err != nil {
		return nil, err
	}

	err = w.AddFile("signature", signature)
	if err != nil {
		return nil, err
	}

	w.archive.Close()
	return w.buffer, nil
}
