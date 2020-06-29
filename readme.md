## Comment service
This service will handle data like comment & organization name and this service has own database.
There are 2 service on this project, 1 is this service (comment-service) and the other 1 is in different repo called (user-service).
Comment service will connect with User service through inter-communication using GRPC

### Pre-requisite
- Minikube
- Docker
- Kubectl

### Technologies
- Golang
- GRPC
- Gorm
- Postgre
- Docker
- Kubernetes

### How to deploy
``$ minikube start``
- Start minikube

``$ kubectl apply -f deployment/secret.yaml``
- Deploy initial secret containing env first before deploy the pod

``$ sh deploy.sh``
- In this deployment command, will execute build docker image, cleanup existing deployment & pod service, and deploy it to minikube

```$ kubectl port-forward service/comment-svc 8080:8080 6060:6060```
- After pod is deployed on minikube, need to forward the port using above command because by default service port is not accessible by public. Here is using clusterIp, to make it similar like in production microservices. After that service will serve and can access using ``localhost:8080``. `:8080` is port for rest & `:6060` is port for grpc

### Api Doc
- HOST: `localhost:8080`
- Create new comment
    - method: `POST`
    - organization_name: `xendit`
    - endpoint: `{HOST}/orgs/{organization_name}/comments`
    - example request: 
    `
    // POST localhost:8080/orgs/xendit/comments
    {
         "comment": "comment dari xendit 1"
    }
    `
    - example response:
    `
    {
        "status": 201,
        "data": {
            "id": 1,
            "comment": "comment dari xendit 1",
            "organization_id": 1
        },
        "message": "comment is successful created"
    }
    `
- Get all comments
    - method: `GET`
    - organization_name: `xendit`
    - endpoint: `{HOST}/orgs/{organization_name}/comments`
    - example response:
    `
        {
            "status": 200,
            "data": [
                {
                    "id": 1,
                    "comment": "comment dari xendit 1",
                    "created_at": "2020-06-28T21:24:30.366562Z"
                }
            ],
            "message": "success"
        }
    `
- Delete all comments
    - method: `DELETE`
    - organization_name: `xendit`
    - endpoint: `{HOST}/orgs/{organization_name}/comments`
    - example response:
    `
        {
            "status": 200,
            "message": "all comments are successful removed"
        }
    `
### DB design
https://dbdiagram.io/d/5ef9238b0425da461f03fa12

### How to run testing
- `$ cd x-comment-service/test/comment`
- `$ go command_test.go`