package main

import (
	"testing"

	ft "github.com/bakito/crossplane-function-inventory-sdk/testing"

	"github.com/crossplane/function-sdk-go/logging"
)

func TestRunFunctionGolden(t *testing.T) {
	cases := map[string]ft.GoldenFileCase{
		"NewBuckets": {
			Reason: "It should create one NopResource per name in spec.names, " +
				"tagged with region and external-name.",
			Composite:       "testdata/xr.yaml",
			DesiredLocation: "testdata/golden/",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc.Run(t, name, &Function{log: logging.NewNopLogger()})
		})
	}
}
