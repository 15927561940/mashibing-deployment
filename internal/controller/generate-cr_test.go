package controller

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	myAppsv1 "mashibing.com/pkg/mashibing-deployment/api/v1"
	"reflect"
	"testing"
)

func TestNewDeployment(t *testing.T) {
	type args struct {
		md *myAppsv1.MSbDeployment
	}
	tests := []struct {
		name    string
		args    args
		want    *appsv1.Deployment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDeployment(tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDeployment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeployment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIngress(t *testing.T) {
	type args struct {
		md *myAppsv1.MSbDeployment
	}
	tests := []struct {
		name    string
		args    args
		want    *networkv1.Ingress
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIngress(tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIngress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIngress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService(t *testing.T) {
	type args struct {
		md *myAppsv1.MSbDeployment
	}
	tests := []struct {
		name    string
		args    args
		want    *corev1.Service
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewServiceNP(t *testing.T) {
	type args struct {
		md *myAppsv1.MSbDeployment
	}
	tests := []struct {
		name    string
		args    args
		want    *corev1.Service
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServiceNP(tt.args.md)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewServiceNP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceNP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
