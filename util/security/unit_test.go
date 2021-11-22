package security_test

import (
	"testing"
	"x-men/test/testdata"
	"x-men/util/security"
)

func TestGenerateHash(t *testing.T) {

	for _, testD := range testdata.UtilGenerateHash() {
		hash := security.GenerateHash(testD.Data...)

		if hash != testD.Hash {
			t.Error(testD.Description)
		}
	}
}
