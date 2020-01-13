# Todo App
The todo app is a Golang application designed to be a example of Go, Docker and Kubernetes with a React Typrescript frontend.

## Getting Setup
Currently to get started, you will need to run go through the following
1. [Download Docker](https://docs.docker.com/v17.12/install/)
2. [Download Go](https://golang.org/doc/install)
3. [Install Minikube (including kubectl](https://kubernetes.io/docs/tasks/tools/install-minikube/)
4. [Install Skaffold](https://skaffold.dev/docs/quickstart/)
5. :construction: TODO: Add secrets key addition here
6. Run the command (for more detals see starting backed server below)
```bash
skaffold dev --port-forward
```
7. Get the name of the MySQL pod using 
```bash
kubectl get pods
```
8. Open a bash shell in the above by running
```bash
kubectl exec -it [POD_NAME_HERE]  -- /bin/bash
```
9. In the shell, open MySQL using, entering the password from above at the prompt
```bash
mysql -u [secret_root_user] -p
``` 

## Starting backend server
In order to start the backend server, run the command
```bash
skaffold dev --port-forward
```
This will build images of the database and the server, deploy them to minikube, watch for changes and auctomatically build when changes are detected. The databasw will be accessible via localhost:3307 (however it might also be 3306 depending if you have MySQL running on the local machine) and the server available at http://localhost:8080.

## Starting the frontend server
Coming soon, probably when there is a frontend server
