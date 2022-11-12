package stringer

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewReverseCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"testisawesome"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "emosewasitset\n" {
		t.Fatalf("expected \"%s\" got \"%s\"", "emosewasitset", string(out))
	}
}
