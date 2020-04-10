# Install kubectl

```
➜ gcloud components install kubectl
```

or 

```
# mac

➜ curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes- release/release/stable.txt)/bin/darwin/amd64/kubectl"

➜ chmod +x ./kubectl

➜ sudo mv ./kubectl /usr/local/bin/kubectl

# linux

➜ curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes- release/release/stable.txt)/bin/linux/amd64/kubectl

➜ chmod +x ./kubectl

➜ sudo mv ./kubectl /usr/local/bin/kubectl

```

# Install minikube

```
➜ gcloud components install minikube 
```

or

```
# mac
➜ curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.21.0/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

# linux
➜ curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.21.0/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```

# Hello, there

start minikube
```
➜ minikube start

# flag -p
```

Check the status:
```
➜ minikube status
```

# Kubectl - command

check node
```
➜ kubectl get node
```

To see which cluster you point to
```
➜ kubectl config get-contexts
```

To choose a different context
```
➜ kubectl config set-context helloworld
```

run the webserver
```
➜ kubectl run webserver --image=raditarya/hello-there-web-server --port=8000
```

A webserver should have its port exposed
```
➜ kubectl expose deployment webserver --type=NodePort
```

In order to get the service URL, type
```
➜ minikube service webserver --url
```

See the content of the web page using
```
➜ curl $(minikube service webserver --url)
```

To show a summary of the running cluster run
```
➜ kubectl cluster-info

---
Kubernetes master is running at https://192.168.99.100:8443
KubeDNS is running at https://192.168.99.100:8443/api/v1/namespaces/kube- system/services/kube-dns:dns/proxy
To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

For more details
```
➜ kubectl cluster-info dump
```

We can also list the pods using
```
➜ kubectl get pods
```

And to access to the dashboard use
```
➜ minikube dashboard
```

proxy
```
➜ kubectl proxy
```

If we want to execute a command inside the container, get the pod id using:
```
➜ kubectl get pods

---
NAME                         READY   STATUS    RESTARTS   AGE
webserver-84f5bc569b-jmzg7   1/1     Running   0          4m10s
```

Then type your command after logging into your container
```
➜ kubectl exec webserver-84f5bc569b-jmzg7 -it -- /bin/sh
```

To finish, delete all deployments
```
➜ kubectl delete deployments --all
```

Delete all pods
```
➜ kubectl delete pods --all
```

stop Minikube
```
➜ minikube stop

---
✋ Stopping "helloworld" in virtualbox ...
🛑 "helloworld" stopped.
```