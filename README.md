# demo-minio

This is a demo for using minio object storage for Go client.

## Pre-requisite

- Install `minikube` or `kind` cluster locally. Cluster should be accesible via `kubectl`
- Install Go

## Install Minio

Run below command to deploy Minio on the local cluster

```
kubectl apply -f minio-dev.yaml
```

## Expose Minio

Run below command to port forward minio console and s3 port access to the localhost

```
kubectl  port-forward svc/minio -n minio-dev 9000 9090
```
Leave the terminal open so that ports remain open on localhost

## Verification in Go

Open another terminal and run below command.

```
go run main.go
```
Above command should create a **demo-bucket** and copy the **testfile** in the bucket

Connect to minio console at `localhost:9090/buckets` from any browser and verify if the bucket is created and the sample testfile exists.

Credentials for minio console - `minio|minio123`