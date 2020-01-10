# Todo App
The todo app is a Golang application designed to be a example of Go, Docker and Kubernetes with a React Typrescript frontend.

## Getting Setup
Currently to get started, you will need to run go through the following
1. [Download Docker](https://docs.docker.com/v17.12/install/)
2. [Download Go](https://golang.org/doc/install)
3. [Install Minikube (including kubectl](https://kubernetes.io/docs/tasks/tools/install-minikube/)
4. [Install Skaffold](https://skaffold.dev/docs/quickstart/)

## Starting backend server
In order to start the backend server, run the command
```bash
skaffold dev --port-forward
```
This will build images of the database and the server, deploy them to minikube, expose the server on http://localhost:8080, watch for changes and auctomatically build when changes are detected.

## Starting the frontend server
Coming soon, probably when there is a frontend server
