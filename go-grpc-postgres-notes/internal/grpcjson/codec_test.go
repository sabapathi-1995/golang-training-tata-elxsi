package grpcjson

import "testing"

func TestCodecName(t *testing.T) {
	var c Codec
	if c.Name() != Name {
		t.Fatalf("expected %q, got %q", Name, c.Name())
	}
}
