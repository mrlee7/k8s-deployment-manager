package controllers

import (
	"context"
	"k8s-custom-controller/api/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type CustomJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=custom.example.com,resources=customjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=custom.example.com,resources=customjobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=custom.example.com,resources=customjobs/finalizers,verbs=update

func (r *CustomJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Fetch the CustomJob instance
	customJob := &v1.CustomJob{}
	err := r.Get(ctx, req.NamespacedName, customJob)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Implement your logic here

	return ctrl.Result{}, nil
}

func (r *CustomJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.CustomJob{}).
		Complete(r)
}
