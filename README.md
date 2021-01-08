# Assignment

For this task you'll have to deploy a Go service to Kubernetes cluster, expose it to outside traffic and write a command-line client in bash.

## Objectives
- Kubernetes manifests for running the service and exposing public IP to access it
- service running on GKE
- client written for Bash

Create a branch in this repository and commit all your work products there.

### Service

You will find ready-made service code under `services` directory written in Go.

Prepare a `Dockerfile` for your server that can produce runnable service. Upload docker image to Google Cloud Container Registry of the project.  To be able to work with the registry you need to setup credentials for it:
```
gcloud auth configure-docker
```
For more information about working with Google Cloud Container Registry see [the documentation](https://cloud.google.com/container-registry/docs/pushing-and-pulling).

Tag the docker image with name.


### Connecting to the cluster

Install the [Google Cloud SDK](https://cloud.google.com/sdk/) for your platform.
Once that's done, configure it for project and zone.
You can find the SDK quick start guide
[here](https://cloud.google.com/sdk/docs/quickstarts).

Get credentials for cluster:
```
gcloud container clusters get-credentials marat --zone=ZONE_HERE
```
To interact with the cluster you need
[kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/).  Make sure you can connect to it:`kubectl cluster-info`

### CLI client
The service has two methods: `/echo` and `/now`. Take a look at the comments in the `service/main.go` to learn how they work.

Write a client in bash to interact with the service. Any common linux core tools are acceptable (curl, grep, whatever). Using any less common CLI tools is acceptable if there are steps included how others can install the needed dependency.

The script should work like this:
```
$ ./script ask-time
<your-name>: <greeting>!
Server: <greeting> <your-name>!
<your-name>: Could you please tell me the time?
Server: Yes I can! The time here is <human-readable-time>
<your-name>: Oh wow! Here in Australia it is <human-readable-time-in-australia-at-the-same-time> already
```
What is happening there is that it prints out human readable summary of the converstaion between your client and the server.
You might think you don't need a server for such script but in this task you do!

**your-name** is the name that server knows you by
**greetings** arbitarilty chosen greeting but the server needs to greet you back the same way!

To make sure that summary is the truthfull represantation of the HTTP traffic you should call both methods at least once.


### Kubernetes manifests

You should write a service and deployment manifest for your server, making sure that it gets a public IP assigned so that the CLI can access to it.

### Applying manifests and testing

Apply those manifests to the cluster! Let us know what you would do to know if everything went well 😈
