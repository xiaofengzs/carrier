package resource

import (
	controllerv1 "github.com/carrier/kb/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconcilable interface {
	MakeResource(carrier *controllerv1.Carrier) client.Object
	MakeDefaultResource() client.Object
	Validate(carrier *controllerv1.Carrier) error
}
