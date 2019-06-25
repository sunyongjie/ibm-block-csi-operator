package apis

import (
	v1 "github.com/IBM/ibm-block-csi-driver-operator/pkg/apis/csi/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
