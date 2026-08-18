// Package main implements a Composition Function.
package main

import (
	"context"

	fnv1 "github.com/crossplane/function-sdk-go/proto/v1"
	"github.com/crossplane/function-sdk-go/resource/composed"
	"github.com/crossplane/function-sdk-go/response"
	"github.com/pkg/errors"

	"github.com/bakito/crossplane-function-inventory-sdk/inventory"

	examplev1 "dev.crossplane.io/models/io/crossplane/example/v1"
	nopv1alpha1 "dev.crossplane.io/models/io/crossplane/nop/v1alpha1"

	"github.com/crossplane/function-sdk-go/logging"
)

func init() {
	// Register once at startup, instead of on every RunFunction call.
	_ = nopv1alpha1.AddToScheme(composed.Scheme)
}

// Inventory declares the resources this Function reads and writes.
type Inventory struct {
	ObservedComposite *examplev1.XBuckets `crossplane:"observed-composite"`

	DesiredBuckets map[string]*nopv1alpha1.NopResource `crossplane:"desired-composed:xbuckets-"`
}

// Function is your composition function.
type Function struct {
	fnv1.UnimplementedFunctionRunnerServiceServer

	log logging.Logger
}

// RunFunction runs the Function.
func (f *Function) RunFunction(_ context.Context, req *fnv1.RunFunctionRequest) (*fnv1.RunFunctionResponse, error) {
	f.log.Info("Running function", "tag", req.GetMeta().GetTag())

	rsp := response.To(req, response.DefaultTTL)

	inv := &Inventory{}
	if err := inventory.BuildFromRequest(req, f.log, inv); err != nil {
		response.Fatal(rsp, errors.Wrap(err, "cannot build inventory"))
		return rsp, nil
	}

	if err := buildBuckets(inv); err != nil {
		response.Fatal(rsp, errors.Wrap(err, "cannot build desired buckets"))
		return rsp, nil
	}

	if err := inventory.ConvertToResponse(rsp, inv); err != nil {
		response.Fatal(rsp, errors.Wrap(err, "cannot convert inventory to response"))
		return rsp, nil
	}

	return rsp, nil
}
