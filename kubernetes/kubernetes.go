package kubernetes

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JobsSpec contains all the configuration needed to run Jobs on the Kubernetes cluster.
type JobsSpec struct {
	Prefix    string
	Namespace string
	Image     string
}

// JobSpecRunner defines the methods for a Kubernetes Jobs manager. It have been
// created for testing purposes.
type JobSpecRunner interface {
	Create(string) error
	Delete(string) error
}

// Create creates a Kubernetes JobsSpec with the given image on the given
// namespace.
//
// Returns nil if succeed creating the JobsSpec, an error otherwise.
func (j *JobsSpec) Create(name string) error {
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprint(j.Prefix, "-", name), Namespace: j.Namespace},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					RestartPolicy: "Never",
					Containers: []apiv1.Container{
						{
							Name:  fmt.Sprint(j.Prefix, "-", name),
							Image: j.Image,
						},
					},
				},
			},
		},
	}

	client, err := getClientSet()
	if err != nil {
		return err
	}

	jobs := client.BatchV1().Jobs(j.Namespace)
	if _, err = jobs.Create(job); err != nil {
		return err
	}

	return nil
}

// Delete remove a Kubernetes JobsSpec with the given name on the given
// namespace.
//
// Returns nil if succeed deleting the JobsSpec, an error otherwise.
func (j *JobsSpec) Delete(name string) error {
	client, err := getClientSet()
	if err != nil {
		return err
	}

	policy := metav1.DeletePropagationForeground

	jobs := client.BatchV1().Jobs(j.Namespace)
	if err := jobs.Delete(fmt.Sprint(j.Prefix, "-", name), &metav1.DeleteOptions{
		PropagationPolicy: &policy,
	}); err != nil {
		return err
	}

	return nil
}

// getClientSet returns a kubernetes.Clientes configured to perform Kubernetes
// API calls inside the cluster.
//
// The ClientSet would beconfigured to run inside a cluster, that's why no auth
// or  certificate are provided, it would use the default Kubernetes
// ServiceAccount for the provided Namespace.

func getClientSet() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}
