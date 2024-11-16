package controller

import (
	"bytes"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"

	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	myAppsv1 "mashibing.com/pkg/mashibing-deployment/api/v1"

	"k8s.io/apimachinery/pkg/util/yaml"
	"text/template"
)

func parseTemplate(md *myAppsv1.MSbDeployment, templateName string) ([]byte, error) {
	tmpl, err := template.ParseFiles(fmt.Sprintf("controller/templates/%", templateName))
	if err != nil {
		return nil, err
	}
	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, md); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// 创建deployment解析的方法
func NewDeployment(md *myAppsv1.MSbDeployment) (*appsv1.Deployment, error) {
	//获取内容
	context, err := parseTemplate(md, "deployment.yaml")
	if err != nil {
		return nil, err
	}
	deploy := new(appsv1.Deployment)

	//解析成yaml
	if err := yaml.Unmarshal(context, deploy); err != nil {
		return nil, err
	}
	return deploy, nil
}

func NewIngress(md *myAppsv1.MSbDeployment) (*networkv1.Ingress, error) {
	//获取内容
	context, err := parseTemplate(md, "ingress.yaml")
	if err != nil {
		return nil, err
	}
	ingress := new(networkv1.Ingress)

	//解析成yaml
	if err := yaml.Unmarshal(context, ingress); err != nil {
		return nil, err
	}
	return ingress, nil
}

func NewService(md *myAppsv1.MSbDeployment) (*corev1.Service, error) {
	//获取内容
	context, err := parseTemplate(md, "service.yaml")
	if err != nil {
		return nil, err
	}
	service := new(corev1.Service)

	//解析成yaml
	if err := yaml.Unmarshal(context, service); err != nil {
		return nil, err
	}
	return service, nil
}

func NewServiceNP(md *myAppsv1.MSbDeployment) (*corev1.Service, error) {
	//获取内容
	context, err := parseTemplate(md, "service-np.yaml")
	if err != nil {
		return nil, err
	}
	service := new(corev1.Service)

	//解析成yaml
	if err := yaml.Unmarshal(context, service); err != nil {
		return nil, err
	}
	return service, nil
}
