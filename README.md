
# hello-go-app 
Simple Go web app that responds with "Hello World" and current time.

### Build Status

![Docker Build](https://github.com/sirajudheenam/hello-app/actions/workflows/docker-build.yml/badge.svg?branch=main)
![Docker Image Version](https://img.shields.io/docker/v/sirajudheenam/hello-go-app?sort=semver)



```bash

go mod init hello-go-app

go get gopkg.in/natefinch/lumberjack.v2

go run main.go

# optionally 
go build -o hello-server main.go
./hello-server

# create local dir called logs
mkdir logs


# Build the container
docker build -t hello-go-app:latest .

# Run the container
docker run -p 8080:8080 hello-go-app:latest

# run with local logs directory
# docker run -p 8080:8080 -v $(pwd)/logs:/app/logs hello-go-app:latest

# 
# docker run -it --rm hello-go-app-alpine bash


docker build -t hello-go-app:v1.0.0 .

docker run -p 8080:8080 hello-go-app:v1.0.0

docker login

docker tag hello-go-app:v1.0.0 sirajudheenam/hello-go-app:v1.0.0

docker push sirajudheenam/hello-go-app:v1.0.0

# extract container ID from docker ps command
CONTAINER_ID=$(docker ps -q -f "ancestor=hello-go-app:latest" | head -n 1)
echo $CONTAINER_ID

# Make a Reques to our server on command line or Web Browser
curl localhost:8080

# To view logs the request
docker exec -it $CONTAINER_ID cat /app/logs/server.json.log

Starting server on :8080
{"timestamp":"2025-09-09T01:04:18Z","remote_addr":"173.194.69.82:22615","method":"GET","path":"/","status":200,"latency_ms":0,"response_size":40}
{"timestamp":"2025-09-09T01:04:18Z","remote_addr":"173.194.69.82:22615","method":"GET","path":"/favicon.ico","status":200,"latency_ms":0,"response_size":40}
{"timestamp":"2025-09-09T01:10:22Z","remote_addr":"173.194.69.82:33847","method":"GET","path":"/","status":200,"latency_ms":0,"response_size":40}


# create a file .github/workflows/docker-build.yml
name: Build and Push Docker Image

on:
  push:
    branches:
      # - main
      - 'v*'   # only run when pushing version tags like v1.0.0

jobs:
  docker:
    runs-on: ubuntu-latest
    # Secrets are located under an environment gh secret list -e DOCKER
    environment: DOCKER # 🔑 use DOCKER environment

    steps:
      # Checkout the repo
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set image tags
        id: vars
        run: |
          echo "date=$(date +'%Y%m%d')" >> $GITHUB_OUTPUT
          echo "sha=$(echo $GITHUB_SHA | cut -c1-7)" >> $GITHUB_OUTPUT
          echo "tag=${echo $GITHUB_REF_NAME}" >> $GITHUB_OUTPUT

      # Log in to Docker Hub (set secrets in repo settings)
      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      # Extract metadata (tags, labels) for Docker
      - name: Extract Docker metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: sirajudheenam/hello-go-app

      # Build and push the Docker image
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: |
            sirajudheenam/hello-go-app:latest
            sirajudheenam/hello-go-app:${{ steps.vars.outputs.date }}-${{ steps.vars.outputs.sha }}
            sirajudheenam/hello-go-app:${{ steps.vars.outputs.tag }}

# Generate SSH Keys on your macOS or compatible Systems
ssh-keygen -t ed25519 -C "your_email@example.com"

# on legacy systems use,
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"

# Copy the public key ed25519.pub or id_rsa.pub content to your github under 
# https://github.com/settings/profile
# https://github.com/settings/ssh/new 
# Title: <any title you like>
# Key type: Authentication Key
# Key: <PASTE your ed25519.pub or id_rsa.pub content>
# Click on Add SSH Key button to save it.

```


<!-- **Latest Docker Image:** `sirajudheenam/hello-go-app:{{TAG}}` -->


## Kubernetes Deployment

```

kubectl create -f ./k8s

deployment.apps/hello-app created
service/hello-app-service created


kubectl get deployments

NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
fortune-app-blue      0/3     1            0           32h
hello-app             2/2     2            2           7s
kubernetes-bootcamp   1/1     1            1           3d12h
source-ip-app         0/1     1            0           3d10h

kubectl get svc

NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
clusterip           ClusterIP   10.107.212.138   <none>        80/TCP    3d10h
hello-app-service   ClusterIP   10.110.121.100   <none>        80/TCP    16s
kubernetes          ClusterIP   10.96.0.1        <none>        443/TCP   3d12h

kubectl get pods -l app=hello-app

NAME                         READY   STATUS    RESTARTS   AGE
hello-app-77dd695876-fjqb6   1/1     Running   0          4m28s
hello-app-77dd695876-xtwqm   1/1     Running   0          4m28s

kubectl logs -l app=hello-app

Starting server on :8080
Starting server on :8080


# Add ingress on minikube
minikube addons enable ingress


💡  ingress is an addon maintained by Kubernetes. For any concerns contact minikube on GitHub.
You can view the list of minikube maintainers at: https://github.com/kubernetes/minikube/blob/master/OWNERS
💡  After the addon is enabled, please run "minikube tunnel" and your ingress resources would be available at "127.0.0.1"
    ▪ Using image registry.k8s.io/ingress-nginx/controller:v1.12.2
    ▪ Using image registry.k8s.io/ingress-nginx/kube-webhook-certgen:v1.5.3
    ▪ Using image registry.k8s.io/ingress-nginx/kube-webhook-certgen:v1.5.3
🔎  Verifying ingress addon...
🌟  The 'ingress' addon is enabled


kubectl get pods -n ingress-nginx

NAME                                       READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-h4jng       0/1     Completed   0          33s
ingress-nginx-admission-patch-7n9xf        0/1     Completed   1          33s
ingress-nginx-controller-67c5cb88f-ff4rc   1/1     Running     0          33s

kubectl create -f ingress.yaml

ingress.networking.k8s.io/hello-app-ingress created


minikube ip

192.168.49.2


sudo nano /etc/hosts

# Add an entry
192.168.49.2  hello-app.local


kubectl get ingress


minikube tunnel


 minikube addons list | grep ingress
| ingress                     | minikube | enabled ✅   | Kubernetes                     |
| ingress-dns                 | minikube | disabled     | minikube                       |


kubectl get endpoints hello-app-service
Warning: v1 Endpoints is deprecated in v1.33+; use discovery.k8s.io/v1 EndpointSlice
NAME                ENDPOINTS                           AGE
hello-app-service   10.244.0.26:8080,10.244.0.27:8080   24m



kubectl run curl --rm -it --image=radial/busyboxplus:curl --restart=Never -- curl -s http://hello-app-service.default.svc.cluster.local:80/


kubectl get pods -l app=hello-app -o wide


kubectl get svc hello-app-service
NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
hello-app-service   ClusterIP   10.110.121.100   <none>        80/TCP    34m


kubectl exec -it deploy/hello-app -- wget -qO- http://localhost:8080/
Hello World - SAM - 2025-09-09 02:18:10


kubectl get ingress hello-app-ingress
NAME                CLASS   HOSTS             ADDRESS        PORTS   AGE
hello-app-ingress   nginx   hello-app.local   192.168.49.2   80      27m

kubectl logs -n ingress-nginx deploy/ingress-nginx-controller


# Try DNS resolution
ping -c 3 hello-app.local

kubectl describe ingress hello-app-ingress

Name:             hello-app-ingress
Labels:           <none>
Namespace:        default
Address:          192.168.49.2
Ingress Class:    nginx
Default backend:  <default>
Rules:
  Host             Path  Backends
  ----             ----  --------
  hello-app.local
                   /   hello-app-service:80 (10.244.0.26:8080,10.244.0.27:8080)
Annotations:       kubernetes.io/ingress.class: nginx
                   nginx.ingress.kubernetes.io/rewrite-target: /
Events:
  Type    Reason  Age                From                      Message
  ----    ------  ----               ----                      -------
  Normal  Sync    17m (x3 over 37m)  nginx-ingress-controller  Scheduled for sync


kubectl run curlpod \
  --rm -it \
  --image=curlimages/curl:7.85.0 \
  --restart=Never \
  -- curl -v http://hello-app-service.default.svc.cluster.local:80/


*   Trying 10.105.100.158:80...
* Connected to hello-app-service.default.svc.cluster.local (10.105.100.158) port 80 (#0)
> GET / HTTP/1.1
> Host: hello-app-service.default.svc.cluster.local
> User-Agent: curl/7.85.0-DEV
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 09 Sep 2025 03:11:04 GMT
< Content-Length: 40
< Content-Type: text/plain; charset=utf-8
<
Hello World - SAM - 2025-09-09 03:11:04
* Connection #0 to host hello-app-service.default.svc.cluster.local left intact
pod "curlpod" deleted from default namespace










# Troubleshooting
kubectl port-forward svc/hello-app-service 8080:80
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
Handling connection for 8080


curl -v http://localhost:8080/
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.7.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Tue, 09 Sep 2025 02:23:14 GMT
< Content-Length: 40
< Content-Type: text/plain; charset=utf-8
<
Hello World - SAM - 2025-09-09 02:23:14
* Connection #0 to host localhost left intact





## NodePort ENABLED

apiVersion: v1
kind: Service
metadata:
  name: hello-app-service
  namespace: default
spec:
  type: NodePort
  selector:
    app: hello-app
  ports:
    - port: 80         # Service port, what clients use inside the cluster
      targetPort: 8080 # Port on the pod container
      nodePort: 30080  # Port on the Minikube/host node


Local/host (nodePort 30080) -> Service (port 80) -> Pod (targetPort 8080)


## PORT FORWARD
kubectl port-forward svc/hello-app-service <LOCAL_PORT>:<SERVICE_PORT_IN_CLUSTER>

kubectl port-forward svc/hello-app-service 8080:80

## Access it using cURL
curl -v http://localhost:8080

* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.7.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Tue, 09 Sep 2025 03:29:37 GMT
< Content-Length: 40
< Content-Type: text/plain; charset=utf-8
<
Hello World - SAM - 2025-09-09 03:29:37
* Connection #0 to host localhost left intact

## Example with NodePort

        ┌─────────────────────┐
        │     Your Host /     │
        │   Minikube Node     │
        │  (macOS / VM)      │
        │                     │
        │  NodePort: 30080    │  <-- Access from browser / curl
        └─────────┬───────────┘
                  │
                  ▼
        ┌─────────────────────┐
        │   Kubernetes Service │
        │  hello-app-service  │
        │                     │
        │   Port: 80          │  <-- Cluster-internal access
        │   TargetPort: 8080  │  <-- Pod container port
        └─────────┬───────────┘
                  │
                  ▼
        ┌─────────────────────┐
        │       Pod(s)        │
        │   hello-app         │
        │                     │
        │   ContainerPort: 8080│ <-- Your Go app listening here
        └─────────────────────┘



## Example with ingress

                ┌───────────────────────────┐
                │       Your Browser        │
                │ http://hello-app.local    │
                └────────────┬─────────────┘
                             │
                             ▼
                ┌───────────────────────────┐
                │   Ingress Controller      │
                │   (NGINX in Minikube)    │
                │                           │
                │ Routes host/path to svc   │
                └────────────┬─────────────┘
                             │
                             ▼
                ┌───────────────────────────┐
                │  Kubernetes Service       │
                │ hello-app-service         │
                │                           │
                │ Port: 80                  │
                │ TargetPort: 8080          │
                │ NodePort: 30080           │  <-- optional, used if you access via <minikube-ip>:30080
                └────────────┬─────────────┘
                             │
                             ▼
                ┌───────────────────────────┐
                │       Pod(s)              │
                │ hello-app                 │
                │                           │
                │ ContainerPort: 8080       │ <-- Your Go app
                └───────────────────────────┘



## Example to run from another Pod inside the same cluster
kubectl run curlpod \
  --rm -it \
  --image=curlimages/curl:7.85.0 \
  --restart=Never \
  -- curl http://hello-app-service:80/


## Example using busybox avoid separate curl image
kubectl run busyboxpod --rm -it --image=busybox:1.36 --restart=Never -- sh

All commands and output from this session will be recorded in container logs, including credentials and sensitive information passed through the command prompt.
If you don't see a command prompt, try pressing enter.
/ #

wget -qO- http://hello-app-service:80/


# Since NodePort and minikube ip is not responding, let us try with different driver
minikube delete
minikube start --driver=hyperkit 

# did not work on Apple Silicon
# Going back to docker

minikube start --driver=docker
echo "$(minikube ip) hello-app.local" | sudo tee -a /etc/hosts

# Tunnel properly 
minikube service hello-app-service

|-----------|-------------------|-------------|---------------------------|
| NAMESPACE |       NAME        | TARGET PORT |            URL            |
|-----------|-------------------|-------------|---------------------------|
| default   | hello-app-service |          80 | http://192.168.49.2:30080 |
|-----------|-------------------|-------------|---------------------------|
🏃  Starting tunnel for service hello-app-service.
|-----------|-------------------|-------------|------------------------|
| NAMESPACE |       NAME        | TARGET PORT |          URL           |
|-----------|-------------------|-------------|------------------------|
| default   | hello-app-service |             | http://127.0.0.1:50213 |
|-----------|-------------------|-------------|------------------------|
🎉  Opening service default/hello-app-service in default browser...
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.



curl -v http://127.0.0.1:50213/


minikube service hello-app-service --url
http://127.0.0.1:50232





ps -ef | grep docker@127.0.0.1

0  3294  3289   0  6:32AM ttys000    0:00.02 sudo ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o IdentitiesOnly=yes -N docker@127.0.0.1 -p 49983 -i /Users/sam/.minikube/machines/minikube/id_rsa -L 80:127.0.0.1:80 -L 443:127.0.0.1:443

0  3302  3294   0  6:32AM ttys000    0:00.01 ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o IdentitiesOnly=yes -N docker@127.0.0.1 -p 49983 -i /Users/sam/.minikube/machines/minikube/id_rsa -L 80:127.0.0.1:80 -L 443:127.0.0.1:443

501  3309  3289   0  6:32AM ttys000    0:00.02 ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o IdentitiesOnly=yes -N docker@127.0.0.1 -p 49983 -i /Users/sam/.minikube/machines/minikube/id_rsa -L 8080:10.96.248.122:8080
501  3893  1728   0  7:04AM ttys004    0:00.00 grep docker@127.0.0.1



     ┌──────────────────────────────┐
     │      Your Mac Host           │
     │  (localhost / 127.0.0.1)     │
     │                              │
     │ curl http://localhost:8080/  │ <-- NodePort
     │ curl http://hello-app.local  │ <-- Ingress
     └───────────────┬─────────────┘
                     │
     ┌───────────────┴──────────────────┐
     │      SSH Tunnels by Minikube      │
     │      (docker@127.0.0.1)          │
     │                                  │
     │ NodePort: 8080 → Pod 8080       │
     │ Ingress HTTP: 80 → ClusterIP 80 │
     │ Ingress HTTPS: 443 → ClusterIP 443 │
     └───────────────┬──────────────────┘
                     │
       ┌─────────────┴─────────────┐
       │  Minikube Kubernetes Cluster│
       │                              │
       │  Service hello-app-service  │
       │    ClusterIP: 80 → Pod 8080 │
       │                              │
       │  Pod(s) hello-app           │
       │  ContainerPort: 8080        │
       └─────────────┬─────────────┘
                     │
         ┌───────────┴───────────┐
         │ Temporary curl Pod     │
         │ (curlimages/curl)     │
         │                        │
         │ curl http://hello-app-service:80/  │
         └─────────────────────────┘




# Other way of running it using NodePort 

kubectl port-forward svc/hello-app-service 8080:80
curl http://localhost:8080/



### Using LoadBalancer - Works

minikube tunnel

## minikube tunnel will create port forwarding internally 
ps -ef | grep docker@127.0.0.1


kubectl expose deployment hello-app --type=LoadBalancer --port=8080

service/hello-app exposed

kubectl get svc

NAME                TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-app           LoadBalancer   10.96.248.122   127.0.0.1     8080:32237/TCP   9s
hello-app-service   NodePort       10.97.170.37    <none>        80:30080/TCP     37m
kubernetes          ClusterIP      10.96.0.1       <none>        443/TCP          38m

curl -v 127.0.0.1:8080

*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/8.7.1
> Accept: */*
>
* Request completely sent off
< HTTP/1.1 200 OK
< Date: Tue, 09 Sep 2025 05:33:53 GMT
< Content-Length: 40
< Content-Type: text/plain; charset=utf-8
<
Hello World - SAM - 2025-09-09 05:33:53
* Connection #0 to host 127.0.0.1 left intact



```






## Updates
v1.0.23 - Updated logging middleware to display latency, removed PR with update README.md