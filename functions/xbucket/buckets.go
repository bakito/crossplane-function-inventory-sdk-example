package main

import (
	"github.com/pkg/errors"

	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"

	nopv1alpha1 "dev.crossplane.io/models/io/crossplane/nop/v1alpha1"
	metav1 "dev.crossplane.io/models/io/k8s/meta/v1"
)

// buildBuckets turns spec.region/spec.names of the observed XBuckets composite
// into one desired NopResource per name, tagged with region and external-name.
func buildBuckets(inv *Inventory) error {
	spec := inv.ObservedComposite.GetSpec()
	if spec == nil || spec.GetRegion() == nil {
		return errors.New("spec.region field of observed composite resource is required")
	}
	region := *spec.GetRegion()

	var names []string
	if n := spec.GetNames(); n != nil {
		names = *n
	}

	buckets := make(map[string]*nopv1alpha1.NopResource, len(names))
	for _, name := range names {
		annotations := map[string]string{meta.AnnotationKeyExternalName: name}
		fields := map[string]any{"region": region, "name": name}

		buckets[InventoryDesiredComposedXbuckets(name)] = &nopv1alpha1.NopResource{
			Metadata: &metav1.ObjectMeta{Annotations: &annotations},
			Spec: &nopv1alpha1.NopResourceSpec{
				ForProvider: &nopv1alpha1.NopResourceSpecForProvider{
					Fields: &fields,
				},
			},
		}
	}
	inv.DesiredBuckets = buckets
	return nil
}
