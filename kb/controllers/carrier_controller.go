/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/carrier/kb/controllers/reconciler"
	"github.com/carrier/kb/controllers/resource"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	controllerv1 "github.com/carrier/kb/api/v1"
)

// CarrierReconciler reconciles a Carrier object
type CarrierReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	Reconcilers []reconciler.Reconciler
}

func NewCarrierReconciler(client client.Client, Scheme *runtime.Scheme) *CarrierReconciler {
	cr := &CarrierReconciler{
		Client: client,
		Scheme: Scheme,
	}
	deploymentReconciler := reconciler.NewDeploymentReconciler(client, Scheme, resource.NewDeployment())
	serviceReconciler := reconciler.NewServiceReconciler(client, Scheme, resource.NewService())
	cr.Reconcilers = append(cr.Reconcilers, deploymentReconciler, serviceReconciler)
	return cr
}

//+kubebuilder:rbac:groups=controller.xiaofeng.cloud,resources=carriers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=controller.xiaofeng.cloud,resources=carriers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=controller.xiaofeng.cloud,resources=carriers/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps/v1,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Carrier object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
func (r *CarrierReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	carrier := &controllerv1.Carrier{}
	if err := r.Get(ctx, req.NamespacedName, carrier); err != nil {
		return ctrl.Result{}, err
	}

	for _, r := range r.Reconcilers {
		if result, err := r.Reconcile(ctx, req, carrier); err != nil {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CarrierReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&controllerv1.Carrier{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
