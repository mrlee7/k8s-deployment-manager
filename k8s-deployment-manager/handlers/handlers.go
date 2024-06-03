package handlers

import (
	"encoding/json"
	"k8s-deployment-manager/k8s"
	"log"
	"net/http"
	"text/template"

	"k8s.io/client-go/kubernetes"
)

var clientset *kubernetes.Clientset

func Initialize(cs *kubernetes.Clientset) {
	clientset = cs
	k8s.Initialize(cs)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "error parsing template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "error executing template", http.StatusInternalServerError)
	}
}

func errorHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/home.html", nil)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	deployment, err := k8s.CreateDeployment()
	if err != nil {
		errorHandler(w, err)
		return
	}
	json.NewEncoder(w).Encode(deployment)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	deployment, err := k8s.UpdateDeployment()
	if err != nil {
		errorHandler(w, err)
		return
	}
	json.NewEncoder(w).Encode(deployment)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	deployments, err := k8s.ListDeployments()
	if err != nil {
		errorHandler(w, err)
		return
	}
	json.NewEncoder(w).Encode(deployments)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	err := k8s.DeleteDeployment()
	if err != nil {
		errorHandler(w, err)
		return
	}
	log.Println(w, "deployment deleted")
}
