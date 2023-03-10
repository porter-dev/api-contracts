package helpers

import (
	"context"
	"errors"
	"fmt"
	"io"

	porterv1 "github.com/porter-dev/api-contracts/generated/go/porter/v1"

	"google.golang.org/protobuf/encoding/protojson"
)

// ContractFromReader is a convenience method for returning a protobuf contract from any
// io.Reader, such as a request body
func ContractFromReader(ctx context.Context, r io.Reader) (*porterv1.Contract, error) {
	var pc *porterv1.Contract

	by, err := io.ReadAll(r)
	if err != nil {
		return pc, fmt.Errorf("unable to read body for contract")
	}

	err = protojson.Unmarshal(by, pc)
	if err != nil {
		return pc, fmt.Errorf("unable to convert reader into contract")
	}

	if err != nil {
		return pc, errors.New("supplied contract was nil")
	}

	return pc, nil
}
