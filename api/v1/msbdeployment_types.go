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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MSbDeploymentSpec defines the desired state of MSbDeployment
type MSbDeploymentSpec struct {
	//我们写的
	//Image镜像存储地址
	Image string `json:"image"`
	//Port  端口
	Port int32 `json:"port"`

	//Replicas副本数
	//+optional
	Replicas int32 `json:"replicas,omitempty"`
	//StartCmd开始的命令
	//+optional
	StartCmd string `json:"start_cmd,omitempty"`
	//Args存储启动命令行参数
	//+optional
	Args []string `json:",omitempty"`
	//Environments环境变量
	//+optional
	Environments []corev1.EnvVar `json:"environments,omitempty"`

	//Expose要暴露服务的模式，ingress还是nodeport等
	Expose *Expose
}

// Expose  //我们写的
type Expose struct {
	//Mode 模式，svc通过那个模式暴露， ingress模式  nodeport模式...之类的
	Mode string `json:"mode"`
	//IngressDomain 服务的ingress的域名，在Mode为ingress的时候此项必填,现在是可选的
	//+optional
	IngressDomain string `json:"ingress_domain,omitempty"`
	//NodePort  端口,在Mode为NodePort时此项必填,现在是可选的
	//+optional
	NodePort int32 `json:"node_port,omitempty"`
	//ServicePort  服务端口,可选，不填写则默认和上面的Port一样
	//+optional
	ServicePort int32 `json:"service_port,omitempty"`
}

// MSbDeploymentStatus defines the observed state of MSbDeployment
type MSbDeploymentStatus struct {
	//Phase 处于什么阶段
	//+optional
	Phase string `json:"phase,omitempty"`
	//Message  这个阶段的信息
	//+optional
	Message string `json:"message,omitempty"`
	//Reason  处于这个阶段的原因
	//+optional
	Reason string `json:",omitempty"`
	//Conditions  处于这个阶段的原因
	//+optional
	Conditions []Condition `json:"conditions,omitempty"`
}

// defines the observed state of Condition
type Condition struct {
	//Type 子资源类型
	//+optional
	Type string `json:"type,omitempty"`
	//Message 子资源的信息
	//+optional
	Message string `json:"message,omitempty"`
	//Status 子资源的状态
	//+optional
	Status string `json:"status,omitempty"`
	//Reason 处于这个状态的原因
	//+optional
	Reason string `json:"reason,omitempty"`

	//LastTransitionTime  最近更新时间
	//+optional
	LastTransitionTime metav1.Time `json:"last_transition_time,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MSbDeployment is the Schema for the msbdeployments API
type MSbDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MSbDeploymentSpec   `json:"spec,omitempty"`
	Status MSbDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSbDeploymentList contains a list of MSbDeployment
type MSbDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSbDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MSbDeployment{}, &MSbDeploymentList{})
}
