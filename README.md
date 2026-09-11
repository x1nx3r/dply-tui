# dply
Custom k8s service deployment management

## k8s strategy
| Namespace  | Pods                          |
|------------|-------------------------------|
| production | auth-service, checker-service |
| staging    | auth-service, checker-service |


## How to create new environment
```
  kubectl create namespace <environment>
```

## How to connect with dply tools
1. Connect local machine with kube (means install kubectl etc.)
2. Install dply client on local machine (run at dply folder)
- ```go build -o dply```
- ```mkdir -p ~/bin```
- ```mv dply ~/bin/```
- ```echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc```
- ```source ~/.zshrc```
check dply already installed
```dply```
3. dply config set-dply-server <dply-server-ip>:<port>
4. dply login -e <email> -p <password>

## How to deploy new service
1. Add docker image
```dply image add -n <service-name> -i <repo@sha256:digest> -d "<desc>"```
Verify image already registered
```dply image list -n <service-name>```
2. Prepare environment variable
```dply spec envar-edit -e <environment> -n <service-name>```
Example
```
{
  "DBDATABASENAME": "core",
  "DBHOST": "127.0.0.1",
  "DBLOGENABLE": true,
  "DBLOGLEVEL": 3,
  "DBLOGTHRESHOLD": 1,
  "DBPASSWORD": "dbpassword",
  "DBPORT": 54330,
  "DBSCHEMA": "public",
  "DBTIMEZONE": "Asia/Jakarta",
  "DBUSERNAME": "postgres",
  "ENV": "dev",
  "ENVIRONMENT": "dev",
  "GRPCPORT": 29000,
  "JAEGERURL": "127.0.0.1:4318",
  "JWTAUDIENCE": "localhost",
  "JWTISSUER": "localhost",
  "JWTSECRET": "jwtsecret",
  "MAINTENANCE": false,
  "RABBITMQURL": "amqp://guest:guest@127.0.0.1:5672/",
  "RESTPORT": 28100,
  "SERVICENAME": "servicename",
  "TOKENDURATION": 1440
}
```
3. Prepare port configuration
```dply spec port-edit -e <environment> -n <service-name>```
example
```
{
  "access_type": "ClusterIP",
  "external_ip": "",
  "ports": [
      {
          "name": "rest",
          "port": 28100,
          "remote_port": 28100,
          "protocol": "TCP"
      },
      {
          "name": "grpc",
          "port": 29000,
          "remote_port": 29000,
          "protocol": "TCP"
      }
  ]
}
```
4. Set scaling strategy
```dply spec scaling-edit -e <environment> -n <service-name>```
example
```
{
  "min_replica": 1,
  "max_replica": 2,
  "min_cpu": 64,
  "max_cpu": 256,
  "min_memory": 256,
  "max_memory": 512,
  "target_cpu": 70
}
```
5. Deploy selected docker image
```dply deploy image <digest> -e <environment> -n <service-name>```
6. Verify app started
```kubectl get pod -n <environment> -l app=<service-name>```

## How to update service or rollback
1. Upload new image to specific environment
```dply image add -n <service-name> -i <repo@sha256:digest> -d "<desc>"```
2. Get image list
```dply image list -n <service-name>```
3. Deploy selected docker image
```dply deploy image <digest> -e <environment> -n <service-name>```