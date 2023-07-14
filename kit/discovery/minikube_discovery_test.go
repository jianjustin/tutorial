package discovery

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"testing"
)

func TestForGetKubeConfigPath(t *testing.T) {
	kubeconfig := getKubeConfigPath()
	fmt.Sprintf("kube config = %s", kubeconfig)
}

func TestForClientSet(t *testing.T) {
	clientset := getClientset()

	version, err := clientset.ServerVersion()
	if err != nil {
		t.Fatalf("unable to get server version: %+v", err)
	}
	fmt.Printf("server version: %v", version)
}

func TestForCreateEndpoint(t *testing.T) {
	clientset := getClientset()
	_, err := clientset.CoreV1().Endpoints("default").Create(context.Background(), &v1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test1",
		},
		Subsets: []v1.EndpointSubset{
			{
				Addresses: []v1.EndpointAddress{
					{
						IP: "192.168.64.1",
					},
				},
				Ports: []v1.EndpointPort{
					{
						Name: "test1",
						Port: 8080,
					},
					{
						Name: "test1",
						Port: 8081,
					},
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return
	}
}

func getKubeConfigPath() string {
	home := homedir.HomeDir()
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}
	return kubeconfig
}

func getClientset() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", getKubeConfigPath())
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
