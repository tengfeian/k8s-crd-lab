/*
Copyright 2025.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	customoperatork8comv1alpha1 "github.com/dilipkumardk/S3ObjectStoreController/api/v1alpha1"
)

// S3ObjStoreReconciler reconciles a S3ObjStore object
type S3ObjStoreReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=custom.operator.k8.com.learning.k8.com,resources=s3objstores,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=custom.operator.k8.com.learning.k8.com,resources=s3objstores/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=custom.operator.k8.com.learning.k8.com,resources=s3objstores/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the S3ObjStore object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *S3ObjStoreReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	instance := &customoperatork8comv1alpha1.S3ObjStore{}
	if err := r.Get(ctx, req.NamespacedName, instance); err != nil {
		logger.Error(err, "Unable to get resource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if instance.Status.State == "" {
		instance.Status.State = customoperatork8comv1alpha1.PENDING_STATE
		r.Status().Update(ctx, instance)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *S3ObjStoreReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&customoperatork8comv1alpha1.S3ObjStore{}).
		Complete(r)
}
