package main

import (
	"testing"
)

func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn menneske ---V \\____/___________/ Ø---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
