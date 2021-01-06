# Todo App
The todo app is a Golang application designed to be a example of Go, Docker and Kubernetes with a React Typrescript frontend.

## Getting Setup
Currently to get started, you will need to run go through the following
1. [Download Docker](https://docs.docker.com/v17.12/install/)
2. [Download Go](https://golang.org/doc/install)
3. [Install Minikube (including kubectl)](https://kubernetes.io/docs/tasks/tools/install-minikube/)
4. [Install Skaffold](https://skaffold.dev/docs/quickstart/)
6. Create a file named secrets.yaml and insert the following. The data is base64 encoded values which you can update yourself.
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-config
type: Opaque
data:                                       # plaintext values
  dbRootUser: cm9vdA==                      # root
  dbRootPw: cm9vdA==                        # root
  dbName: dG9kb2xpc3Q=                      # todolist
  dbUser: dXNlcg==                          # user
  dbPw: cGFzc3dvcmQ=                        # password
```
7. Run the following to create persistent volume claim so our database has somewhere to persist
```bash
kubectl apply -f k8s-persistent-volume-claim.yaml
kubectl apply -f secrets.yaml
kubectl apply -f k8s-todo-database-deployment.yaml
```
8. Get the name of the MySQL pod using 
```bash
kubectl get pods
```
9. Open a bash shell in the above by running
```bash
kubectl exec -it [POD_NAME_HERE]  -- /bin/bash
```
10. In the shell, open MySQL using, entering the password from above at the prompt
```bash
mysql -u [inser the dbRootUser from secrets] -p
``` 
11. Run this [SQL Script](https://github.com/someDevDude/todo-server/blob/master/database/sql/01-richmond/init.sql)
12. Run the command
```bash
skaffold dev --port-forward
```
13. Open http://localhost:8080 in the browser and you should see pong.

## Starting backend server
In order to start the backend server, run the commands
```bash
kubectl apply -f k8s-persistent-volume-claim.yaml
kubectl apply -f secrets.yaml
kubectl apply -f k8s-todo-database-deployment.yaml
skaffold dev --port-forward
```
This will build images of the database and the server, deploy them to minikube, watch for changes and auctomatically build when changes are detected. The database will be accessible via localhost:3307 (however it might also be 3306 depending if you have MySQL running on the local machine) and the server available at http://localhost:8080.

## Stoping the backend server
Exit the Skaffold in the terminal using CTRL+c then, run the following
```bash
kubectl delete -f k8s-persistent-volume-claim.yaml
kubectl delete -f secrets.yaml
kubectl delete -f k8s-todo-database-deployment.yaml
```

## Starting the frontend server
Coming soon, probably when there is a frontend server
