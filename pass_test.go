package wallet

import "testing"

func TestNewPassBarcode(t *testing.T) {
	b := NewPassBarcode("the_message", "the_format")

	if b.Message != "the_message" {
		t.Errorf("expected message to be %s, got %s", "the_message", b.Message)
	}

	if b.Format != "the_format" {
		t.Errorf("expected message to be %s, got %s", "the_format", b.Format)
	}

	if b.MessageEncoding != "iso-8859-1" {
		t.Errorf("expected message to %s, got %s", "iso-8859-1", b.MessageEncoding)
	}
}
