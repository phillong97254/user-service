# user-service  (:8082)

REST API `GET /v1/users/{id}` trả hồ sơ admin, cần token có scope `users:read`. Gateway gọi bằng token client_credentials của chính nó.

Cần auth-service (JWKS) để verify token: `oauth.jwks_url`.

## Chạy

Folder này là 1 project Go độc lập (repo git, module riêng), mở folder lên rồi:

```bash
make init             # lần đầu: tạo .env, bật Postgres, migrate
go run ./cmd/rest    # hoặc: make run
make test             # go test -race ./...
```

Cấu hình ở `config/config.yml`, ghi đè bằng biến môi trường theo quy tắc `.` → `_`
(ví dụ `REST_SERVER_ADDR=:9999`, `OAUTH_JWKS_URL=...`).

## Database (Postgres)

Service có Postgres riêng (database-per-service), bảng: `profiles`. Repository dùng gorm (`go-lib/database`).

```bash
make init             # lần đầu: tạo .env, bật Postgres (docker, cổng 54324) và migrate
make run              # service tự migrate khi khởi động (db.auto_migrate: true)
make migrate          # chạy tay: migrate up bằng go-tools (giống dự án cũ)
make redo-migrate     # xoá hết rồi migrate lại (mất dữ liệu)
```

- Migration: `migration/scripts/000000001_init.*.sql` (schema), `000000002_seed.*.sql` (dữ liệu demo).
  Thêm thay đổi DB thì tạo cặp file số tiếp theo (`000000003_<tên>.up.sql` / `.down.sql`).
- Kết nối: `DB_URL` trong `.env` (`postgres://users:<DB_PASSWORD>@localhost:54324/user_db?sslmode=disable`), không để trong `config.yml`.
- Dữ liệu Postgres lưu ở `./tmp/postgres_data` (đã gitignore). Xoá thư mục này để làm lại từ đầu.
- Test cần Postgres: `TEST_DB_URL=...user_db_test... make test` (database riêng cho test, sẽ bị xoá dữ liệu); không đặt thì test DB được bỏ qua.

## Secret và cấu hình

- `config/config.yml` **được commit**: chỉ chứa giá trị mặc định, không chứa secret.
- Secret đặt trong `.env` (đã có trong `.gitignore` và `.dockerignore`, không commit, không vào image).
  `make env` tạo `.env` từ `.env.example`, sau đó điền giá trị. `make run` tự nạp `.env`.
- Thiếu secret thì service dừng ngay lúc khởi động và báo tên biến cần đặt.
- Trên server/CI: truyền secret bằng biến môi trường (GitHub Actions secrets, Docker/K8s secret...), không dùng `.env` commit lên.

## Thư viện (make mod)

Code dùng chung (`authkit`, `apperr`) nằm ở repo **go-lib** (`github.com/phillong97254/go-lib`), giống `go-library`
ở các dự án trước. Thư viện được vendor vào `vendor/` nên build/Docker không cần mạng hay repo khác.

```bash
make mod                          # go get -u thư viện + go mod tidy + go mod vendor
make mod GO_LIB_VERSION=v0.2.0    # chốt go-lib ở 1 version (khi dùng remote)
make mod-local                    # dùng ../go-lib (replace) khi đang sửa go-lib cùng lúc
make mod-remote                   # bỏ replace, lấy go-lib theo version từ GitHub (mặc định hiện tại)
```

## Docker

```bash
make build            # docker build -t ghcr.io/phillong97254/user-service:dev-latest .
make up               # chạy riêng service này (gọi service khác qua host.docker.internal)
make ci-docker-build ci-docker-push TAG=v1.0.0   # push lên ghcr.io/phillong97254 (GitHub Container Registry)
```

## Cấu trúc

```
cmd/rest/main.go, routes/   # bootstrap + đăng ký route
config/                      # config.go (viper) + config.yml
vendor/                      # thư viện (kể cả go-lib), tạo bằng make mod
internal/
    domain/
    internal/repository/profile
    internal/service/profile
    internal/handler/rest/v1/m2m/user
```
