/*
Copyright 2024.

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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strings"

	myAppsv1 "mashibing.com/pkg/mashibing-deployment/api/v1"
)

// MSbDeploymentReconciler reconciles a MSbDeployment object
type MSbDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps.mashibing.com,resources=msbdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.mashibing.com,resources=msbdeployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps.mashibing.com,resources=msbdeployments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MSbDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *MSbDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx, "MsbDeployment", req.NamespacedName)

	// TODO(user): your logic here
	//1.获取资源对象
	logger.Info("Reconcile  is   started")
	md := new(myAppsv1.MSbDeployment)
	err := r.Client.Get(ctx, req.NamespacedName, md)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err) //保证意外正常退出而不是死循环

	}

	//防止污染缓存
	mdCopy := md.DeepCopy()

	//=====处理deployment======
	//2.获取deployment
	deploy := new(appsv1.Deployment)
	if err := r.Client.Get(ctx, req.NamespacedName, deploy); err != nil {
		if errors.IsNotFound(err) {
			//2.1不存在，则创建
			//2.1.1创建deployment
			r.createDeployment(ctx, mdCopy)
		} else {
			return ctrl.Result{}, err
		}

	} else {
		//2.2存在则更新
		r.updateDeployment(mdCopy)
	}

	//=====处理svc======
	//3.获取svc对象
	svc := new(corev1.Service)
	if err := r.Client.Get(ctx, req.NamespacedName, svc); err != nil {
		//3.1不存在svc则创建

		if errors.IsNotFound(err) {
			if mdCopy.Spec.Expose.Mode == myAppsv1.ModeIngress {
				//3.1.1mode如果是ingress
				//3.1.1.1创建普通svc
				r.createService(mdCopy)
			} else if mdCopy.Spec.Expose.Mode == myAppsv1.ModeNodePort {
				//3.1.2mode如果是nodePort
				//3.1.2.1创建nodePort模式的svc
				r.createNPService(mdCopy)
			}
		} else {
			return ctrl.Result{}, err
		}
	} else {
		//3.2.1mode如果是ingress
		if mdCopy.Spec.Expose.Mode == myAppsv1.ModeIngress {
			//3.2存在svc则更新
			//3.2.1.1更新普通的svc
			r.updateService(mdCopy)
		} else if mdCopy.Spec.Expose.Mode == myAppsv1.ModeNodePort {
			//3.2.2mode如果是nodeport
			//3.2.2.1创建nodeport的svc
			r.updateNPService(mdCopy)
		} else {
			return ctrl.Result{}, myAppsv1.ErrNotSupportMode
		}

	}

	//=====处理ingress======
	//4获取ingress资源
	ig := new(networkv1.Ingress)
	if err := r.Client.Get(ctx, req.NamespacedName, ig); err != nil {
		//4.1不存在ingress
		if errors.IsNotFound(err) {
			if strings.ToLower(mdCopy.Spec.Expose.Mode) == myAppsv1.ModeIngress {
				//4.1.1mode为ingress,创建ingress
				r.createIngress(mdCopy)
			} else if strings.ToLower(mdCopy.Spec.Expose.Mode) == myAppsv1.ModeNodePort {
				//4.1.2mode为nodePort,无须ingress
				//4.1.2.1退出
				return ctrl.Result{}, nil
			}

		} else {
			return ctrl.Result{}, err

		}
	} else {
		//4.2存在ingress
		if strings.ToLower(mdCopy.Spec.Expose.Mode) == myAppsv1.ModeIngress {
			//4.2.1更新ingress
			r.updateIngress(mdCopy)
		} else if strings.ToLower(mdCopy.Spec.Expose.Mode) == myAppsv1.ModeNodePort {
			//4.2.2mode为nodeport
			//4.2.2删除ingress，更新成nodeport
			r.deleteIngress(mdCopy)
		}

	}

	logger.Info("Reconcile  is   结束")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MSbDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&myAppsv1.MSbDeployment{}).
		Complete(r)
}

func (r *MSbDeploymentReconciler) createDeployment(ctx context.Context, md *myAppsv1.MSbDeployment) {
	deploy, err := NewDeployment(md)
	if err != nil {
		return
	}
	r.Client.Create(ctx, deploy)

}

func (r *MSbDeploymentReconciler) updateDeployment(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) createService(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) createNPService(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) updateService(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) updateNPService(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) createIngress(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) updateIngress(_ *myAppsv1.MSbDeployment) {

}

func (r *MSbDeploymentReconciler) deleteIngress(_ *myAppsv1.MSbDeployment) {

}
