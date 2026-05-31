# learn_dev_backend

Minimal Go backend for the full-stack learning slice.

## Run

```powershell
go run .
```

Then open http://localhost:8080.

## API

- `POST /api/login` with `{"username":"admin","password":"password"}` returns a JWT.
- `GET /api/items` requires `Authorization: Bearer <token>`.

## Test

```powershell
go test ./...
```

## Docker

The Dockerfile pulls base images through the Xuanyuan mirror:

- `qp2pdj7hsiudzdknli.xuanyuan.run/library/golang:1.22-alpine`
- `qp2pdj7hsiudzdknli.xuanyuan.run/library/alpine:3.23`

```powershell
docker build -t learn-dev-backend .
docker run --rm -p 8080:8080 learn-dev-backend
```

## Deployment Note

For a first deployment, build this image and run it on any service that supports containers. In a real app, move `jwtSecret` to an environment variable before deploying.
