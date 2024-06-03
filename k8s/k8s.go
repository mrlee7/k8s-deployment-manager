package k8s

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

var clientset *kubernetes.Clientset

func Initialize(cs *kubernetes.Clientset) {
	clientset = cs
}

func CreateDeployment() (*appsv1.Deployment, error) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:latest",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateDeployment() (*appsv1.Deployment, error) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(context.Background(), "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			return errors.Wrap(getErr, fmt.Sprintf("failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(1)
		result.Spec.Template.Spec.Containers[0].Image = "nginx:latest"
		_, updateErr := deploymentsClient.Update(context.Background(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return nil, errors.Wrap(retryErr, fmt.Sprintf("update failed: %v", retryErr))
	}
	result, _ := deploymentsClient.Get(context.Background(), "demo-deployment", metav1.GetOptions{})
	return result, nil
}

func ListDeployments() (*appsv1.DeploymentList, error) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	list, err := deploymentsClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func DeleteDeployment() error {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.Background(), "demo-deployment", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
