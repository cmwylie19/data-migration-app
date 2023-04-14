# Demo

## Setup Cluster

```bash
cat <<EOF | kind create cluster --name=demo --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 31469 #NodePort of AUTH service
    hostPort: 8080
    protocol: TCP
  - containerPort: 31470 #NodePort of Keycloak service
    hostPort: 8081
    protocol: TCP
  - containerPort: 31471 #NodePort of Frontend service
    hostPort: 8082
    protocol: TCP
EOF
```

Deploy Keycloak

```bash
kubectl apply -f -<<EOF
apiVersion: v1
kind: Namespace
metadata:
  creationTimestamp: null
  name: keycloak
spec: {}
status: {}
---
apiVersion: v1
kind: Service
metadata:
  name: keycloak
  namespace: keycloak 
  labels:
    app: keycloak
spec:
  ports:
  - name: http
    port: 8080
    targetPort: 8080
    nodePort: 31470
  selector:
    app: keycloak
  type: NodePort 
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: keycloak
  namespace: keycloak
  labels:
    app: keycloak
spec:
  replicas: 1
  selector:
    matchLabels:
      app: keycloak
  template:
    metadata:
      labels:
        app: keycloak
    spec:
      containers:
      - name: keycloak
        image: quay.io/keycloak/keycloak:18.0.0
        args: ["start-dev"]
        env:
        - name: KEYCLOAK_ADMIN
          value: "admin"
        - name: KEYCLOAK_ADMIN_PASSWORD
          value: "admin"
        - name: KC_PROXY
          value: "edge"
        ports:
        - name: http
          containerPort: 8080
        readinessProbe:
          httpGet:
            path: /realms/master
            port: 8080
EOF
```

Wait for Keycloak to be ready

```bash
k wait --for=condition=Ready pod -l app=keycloak -n keycloak --timeout=180s
```

Go to [localhost:8081](http://localhost:8081) and login with `admin` / `admin`

- Add a realm `eda`
- Add a client

In the client, ensure Client Authetnication, Standard Flow, Direct Access Grants, Service Accounts are enabled.

the redirect URI should be `http://localhot:8080/*` (the auth app)
Access Type should be set to confidental

- add groups
- add users

## Deploy Auth

Wait for Auth to be ready

```bash
kubectl apply -f -<<EOF
apiVersion: v1
kind: Namespace
metadata:
  creationTimestamp: null
  name: data
spec: {}
status: {}
---
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: auth
  name: auth
  namespace: data
spec:
  replicas: 1
  selector:
    matchLabels:
      app: auth
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: auth
    spec:
      containers:
      - image: cmwylie19/auth:0.0.2
        command: ["./auth", "serve"]
        name: auth
        env:
        - name: KC_URL
          value: http://localhost:8081
        - name: KC_CLIENT_ID
          value: migration-app
        - name: KC_CLIENT_SECRET
          value: BHmt7LdCYMG6VrRQu9NWOMmMbMAxq8CA
        - name: KC_REALM
          value: eda
        resources: {}
status: {}
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: auth
  name: auth
  namespace: data
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
    nodePort: 31469
  selector:
    app: auth
  type: NodePort
status:
  loadBalancer: {}
EOF
```

Wait for auth to be ready

```bash
k wait --for=condition=Ready pod -l app=auth -n data --timeout=180s
```
