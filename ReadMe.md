This hands on lab demonstrates a client server setup and communication between them using domain names of services defined in kubernetes.

app2 is the server and app1 is the client. Server simply returns a random number for every new request made by the client. Client takes the response from the server and writes it to a file named "data.txt". 

Follow below steps to dockerize these apps and run them on minikube(local kubernetes cluster).

First build docker images for two apps by using below commands respectively.

cd app1

docker build -t app1:v1 .

cd app2

docker build -t app2:v2 .

It you are using minikube to test this, you need to load these images to minikube so that you can run pods using them.  To do this,

minikube image load app1:v1
minikube image load app2:v1

Now you have these images available for minikube.

Now create a deployment for app2, which is our server that listens to requests on 4040.

kubectl create deployment --image=app2:v1 app2 --port=4040 --replicas=3

Then expose the above deployment through a service app2-svc (*** This service is the same we used in app1, which is our client. Usually services in kubernetes are reachable by following name convention "service-name.namespace.svc.cluster.local". so, our server becomes "app2-svc.default.svc.cluster.local" inside the cluster.)

kubectl expose deploy app2 --port=4040 --name app2-svc

Now, create a pod for our client application which is app1.

kubectl run --image=app1:v1 app1-pod

We are writing the output received from the server to a file named data.txt. we can check if everything is working properly by checking the contents of the "data.txt" file.

we can sh into this app1-pod, which is our client using this command.

kubectl exec -it app1-pod -- sh

We can run "cat data.txt" and see that the file is getting updated every second.



