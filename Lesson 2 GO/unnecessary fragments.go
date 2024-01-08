package gen

import (
	"bytes"
	"encoding/json"
	"flag"

func (h Header) String() string {
	return fmt.Sprintf(`// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
	// See https://github.com/exercism/go#synchronizing-tests-and-instructions
	// Source: %s
	// Commit: %s
	`, h.Origin, h.Commit)
}
//remove unnecessary fragments
func TestCreate(t *testing.T) {
	preTestValidate(t)

	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx
	cli.ForceSkipConfirm = true
