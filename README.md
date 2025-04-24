# Go CRUD DevOps Project

Dự án CRUD đơn giản sử dụng Go, Gin, GORM với cấu hình DevOps hoàn chỉnh bao gồm Docker, Ansible, CI/CD, Prometheus và Grafana.

## Công nghệ sử dụng

- **Backend**: Go, Gin framework, GORM (với PostgreSQL)
- **Containerization**: Docker, Docker Compose
- **Monitoring**: Prometheus, Grafana, Node Exporter
- **CI/CD**: GitHub Actions
- **Deployment**: Ansible

## Cấu trúc dự án

```
go-crud-devops/
├── api/              # API handlers, middlewares, và routes
├── cmd/              # Entry points cho ứng dụng
├── config/           # Cấu hình ứng dụng
├── docker/           # Dockerfiles và cấu hình container
├── internal/         # Packages nội bộ (models, repositories)
├── ansible/          # Playbooks và roles cho Ansible
├── .github/          # GitHub Actions workflows
├── docker-compose.yml
└── README.md
```

## Chạy ứng dụng cục bộ

### Yêu cầu

- Go 1.21 +
- Docker và Docker Compose
- PostgreSQL (hoặc sử dụng container)

### Khởi động với Docker Compose

```bash
# Clone repository
git clone https://github.com/your-username/go-crud-devops.git
cd go-crud-devops

# Chạy với Docker Compose
docker-compose up -d
```

Ứng dụng sẽ chạy tại:
- API: http://localhost:8080
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (mặc định: admin/admin)

### Phát triển cục bộ

```bash
# Cài đặt dependencies
go mod download

# Chạy application
go run cmd/api/main.go
```

## API Endpoints

- `GET /api/users` - Lấy danh sách người dùng
- `GET /api/users/:id` - Lấy thông tin người dùng theo ID
- `POST /api/users` - Tạo người dùng mới
- `PUT /api/users/:id` - Cập nhật thông tin người dùng
- `DELETE /api/users/:id` - Xóa người dùng

## Giám sát

- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (user/pass: admin/admin)

## Triển khai

### Sử dụng Ansible

```bash
cd ansible
ansible-playbook -i inventory.yml playbook.yml
```

### Qua CI/CD Pipeline

Đẩy code lên nhánh `main` hoặc tạo tag để kích hoạt GitHub Actions workflow tự động build, test và triển khai.

## Giấy phép

MIT
