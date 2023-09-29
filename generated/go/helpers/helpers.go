package helpers

import (
	"context"
	"errors"
	"fmt"
	"io"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// UnmarshalContractObjectFromReader is a convenience method for returning a protobuf contract object from any
// io.Reader, such as a request body, and will work for any part of a contract, or contract revision
func UnmarshalContractObjectFromReader(r io.Reader, contract protoreflect.ProtoMessage) error {
	by, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("unable to read body for contract revision")
	}

	return UnmarshalContractObject(by, contract)
}

// UnmarshalContractObject is a convenience method for returning a protobuf contract object from bytes
// and will work for any part of a contract, or contract revision
func UnmarshalContractObject(by []byte, contract protoreflect.ProtoMessage) error {
	err := protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(by, contract)
	if err != nil {
		return fmt.Errorf("unable to convert bytes into contract: %w", err)
	}

	if contract == nil {
		return errors.New("supplied contract was nil")
	}

	return nil
}

// MarshalContract is a convenience function for sending contracts over the wire as bytes.
// This is commony used when forwarding requests through a pub sub system, and will support
// any nested contract object, or contract revision
func MarshalContractObject(ctx context.Context, pc protoreflect.ProtoMessage) ([]byte, error) {
	by, err := protojson.Marshal(pc)
	if err != nil {
		return nil, fmt.Errorf("unable to parse contract: %w", err)
	}
	return by, nil
}
