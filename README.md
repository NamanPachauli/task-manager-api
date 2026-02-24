# TASK MANAGER API

THIS IS MY GOLANG REST API PROJECT CONTAINERIZED WITH DOCKER AND DEPLOYED ON KUBERNETES.

A SIMPLE TODO/TASK MANAGER BACKEND API BUILT WITH GOLANG, EXPOSED VIA HTTP ENDPOINTS, AND RUNNING INSIDE CONTAINERS.

---

## FEATURES
- BUILT USING GOLANG  
- DOCKERIZED APPLICATION  
- KUBERNETES DEPLOYMENT  
- NODEPORT SERVICE FOR ACCESS  
- SIMPLE CRUD FOR TASKS

---

## REQUIREMENTS
- GO 1.22+  
- DOCKER  
- KUBERNETES (DOCKER DESKTOP OR MINIKUBE)  
- GIT

---

## FILES INCLUDED
├── server.go
├── go.mod
├── Dockerfile
├── .dockerignore
├── deployment.yaml
├── service.yaml
└── README.md


---

## API ENDPOINTS

### GET ALL TASKS

GET /tasks
### CREATE NEW TASK
POST /create
BODY:
{
"title": "TASK TITLE",
"done": false
}


---

## RUN LOCALLY (WITHOUT K8s)

BUILD DOCKER IMAGE:
docker build -t taskmanager-app .


RUN CONTAINER:
docker run -p 9091:9091 taskmanager-app


SERVER RUNS AT:
http://localhost:9091


---

## DEPLOY TO KUBERNETES

APPLY DEPLOYMENT:
kubectl apply -f deployment.yaml


APPLY SERVICE:
kubectl apply -f service.yaml

CHECK STATUS:
kubectl get pods
kubectl get services


ACCESS SERVICE:
http://localhost:30007/tasks


---

## NOTES
- KUBERNETES POD LABELS AND SERVICE SELECTOR MUST MATCH  
- IMAGE PULL POLICY SET TO NEVER IF USING LOCAL DOCKER IMAGE  
- WORKS WITH DOCKER DESKTOP K8S

---





                                                    




                                                         
