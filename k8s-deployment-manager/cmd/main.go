package main

import (
	"flag"
	"k8s-deployment-manager/handlers"

	"log"
	"net/http"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientset *kubernetes.Clientset

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	handlers.Initialize(clientset)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/create", handlers.CreateHandler)
	http.HandleFunc("/update", handlers.UpdateHandler)
	http.HandleFunc("/list", handlers.ListHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)

	log.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
