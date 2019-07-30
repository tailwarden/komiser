# Deploy komiser on kubernetes

This deployment will deploy:

1. pod with komiser and redis.
2. It will set up a service with node port for linking with AWS ALB Ingress controller
3. Setup appropriate [ALB Ingress controller](https://github.com/kubernetes-sigs/aws-alb-ingress-controller)

Optional to change the service type to  `LoadBalncer` or use any other Ingress type.

To deploy:

* First create a secret for aws credentials

```
kubectl create secret generic -n monitoring komiser --from-literal key='aws_key' --from-literal secret='aws_secret'
```

* Apply `yaml` files

```
kubectl apply -f ./deployment/
```