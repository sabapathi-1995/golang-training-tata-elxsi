package grpcjson

import (
	"encoding/json"

	"google.golang.org/grpc/encoding"
)

// Name is the codec name used in content-subtype.
// We'll use "json" and force it in both server & client.
const Name = "json"

type Codec struct{}

func (Codec) Name() string { return Name }

func (Codec) Marshal(v any) ([]byte, error) { return json.Marshal(v) }

func (Codec) Unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }

func Register() {
	encoding.RegisterCodec(Codec{})
}
