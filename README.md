                                                         TASK MANAGER REST API PROJECT 🚀
                                                         
 A SIMPLE REST API BUILT USING GOLANG, DOCKER, AND KUBERNETES.               

 
                                                         PROJECT OVERVIEW
                                                         
THIS PROJECT DEMONSTRATES BUILDING A BACKEND API USING GO, CONTAINERIZING IT WITH DOCKER, AND DEPLOYING IT ON KUBERNETES WITH A NODEPORT SERVICE.   


                                                         TECHNOLOGY STACK

                                                    -GOLANG
                                                    -DOCKER
                                                    -KUBERNETES
                                                    -GIT
                                                    -GITHUB

                                                          PROJECT STRUCTURE
                                                          
                                                    task-manager/
                                                    │
                                                    ├── server.go
                                                    ├── go.mod
                                                    ├── Dockerfile
                                                    ├── .dockerignore
                                                    ├── deployment.yaml
                                                    ├── service.yaml
                                                    └── README.md      

                                                    API ENDPOINTS
HOME ROUTE

GET /

RETURNS SERVER STATUS MESSAGE.

GET ALL TASKS

GET /tasks

RETURNS ALL TASKS IN JSON FORMAT.

CREATE NEW TASK

POST /create

REQUEST BODY:
{
  "title": "Learn Kubernetes",
  "done": false
}
DOCKER CONFIGURATION
BUILD IMAGE:
docker build -t taskmanager-app .

RUN CONTAINER:
docker run -p 9091:9091 taskmanager-app

APPLICATION RUNS ON:
http://localhost:9091

KUBERNETES DEPLOYMENT

APPLY DEPLOYMENT:
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

DEPLOYMENT DETAILS

REPLICAS: 2
IMAGE: taskmanager-app
IMAGEPULLPOLICY: NEVER
CONTAINER PORT: 9091

SERVICE DETAILS

TYPE: NODEPORT
PORT: 9091
NODEPORT: 30007

ACCESS APPLICATION AT:
http://localhost:30007/tasks



                                                    




                                                         
