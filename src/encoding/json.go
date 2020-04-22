package encoding

import (
	"encoding/json"
	"fmt"
	"io"

	beerproto "github.com/RossMerr/beerprotobuf/src"
	"github.com/RossMerr/beerprotobuf/src/mapping"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func JSON(r io.Reader, beer *beerproto.Beerjson) error {

	m := make(map[string]json.RawMessage)
	err := json.NewDecoder(r).Decode(&m)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}

	m = mapping.Enum(m)

	data, err := json.Marshal(&m)
	err = protojson.Unmarshal(data, beer)
	if err != nil {
		return fmt.Errorf("json: %w", err)
	}

	return nil
}

type beerJSON struct {
	Beer *beerproto.Beerjson `json:"beerjson"`
}

func (s *beerJSON) ProtoReflect() protoreflect.Message {
	return nil
}