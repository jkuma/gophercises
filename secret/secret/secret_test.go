package secret

import (
	"encoding/hex"
	"testing"
)

func Test_decrypt(t *testing.T) {
	want := "some plaintext"
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	encrypted, _ := hex.DecodeString("ceca6dc9197bc25c93e5ef1b952bfe6d947a51d4a43f0715497edff392d6")

	if decrypted := decrypt(key, encrypted); decrypted != want {
		t.Errorf("The decoded value should be %s and not %v", want, decrypted)
	}
}

func Test_encrypt(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	val := "some plaintext"

	if val != decrypt(key, encrypt(key, val)) {
		t.Errorf("The encrypted value does not contain %v", val)
	}
}
