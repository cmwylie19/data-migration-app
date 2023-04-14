# Auth

Build

```bash
GOOS=linux go build -o auth ./cmd
mv auth build/auth

docker build -t cmwylie19/auth:0.0.1 build/
docker push cmwylie19/auth:0.0.1
```
