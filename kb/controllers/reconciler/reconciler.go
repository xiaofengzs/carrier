package reconciler

import (
	"context"
	controllerv1 "github.com/carrier/kb/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Reconciler interface {
	Reconcile(ctx context.Context, req ctrl.Request, carrier *controllerv1.Carrier) (ctrl.Result, error)
}
