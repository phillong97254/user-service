# user-service  (:8082)

REST API `GET /v1/users/{id}` trả hồ sơ admin, cần token có scope `users:read`. Gateway gọi bằng token client_credentials của chính nó.

Cần auth-service (JWKS) để verify token: `oauth.jwks_url`.

## Chạy

Folder này là 1 project Go độc lập (module riêng, không cần `go.work`), mở folder lên rồi:

```bash
go run ./cmd/rest    # hoặc: make run
make test             # go test -race ./...
```

Cấu hình ở `config/config.yml`, ghi đè bằng biến môi trường theo quy tắc `.` → `_`
(ví dụ `REST_SERVER_ADDR=:9999`, `OAUTH_JWKS_URL=...`).

## Docker

```bash
make build            # docker build -t go-admin-wallet/user-service:dev-latest .
make up               # chạy riêng service này (gọi service khác qua host.docker.internal)
make ci-docker-build ci-docker-push REGISTRY=<registry>/<group> TAG=v1.0.0
```

## Cấu trúc

```
cmd/rest/main.go, routes/   # bootstrap + đăng ký route
config/                      # config.go (viper) + config.yml
pkg/authkit                  # verify JWT (JWKS), middleware REST, interceptor gRPC
internal/
    domain/
    internal/repository/profile
    internal/service/profile
    internal/handler/rest/v1/m2m/user
```
