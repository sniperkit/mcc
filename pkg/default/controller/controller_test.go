/*
Sniperkit-Bot
- Status: analyzed
*/

package controller

import (
	"testing"
)

func TestOptimiseIncompleteParamas(t *testing.T) {
	if err := NewController("0.9.5", "1.1.0", "../../../_examples/simple/example.yml", true); err != nil {
		t.Fatalf("error:%v", err)
	}
}
