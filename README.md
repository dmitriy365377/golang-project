# Golang Chat Application

A modern chat application built with Go, featuring authentication, real-time messaging, and REST API services.

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL
- Docker (optional)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/dmitriy365377/golang-project.git
   cd golang-project
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Build services**
   ```bash
   make build-all
   ```

## 📁 Project Structure

```
golang-chat/
├── cmd/                    # Application entry points
│   ├── auth-service/      # Authentication service
│   ├── chat-service/      # Chat service
│   ├── rest-auth-service/ # REST authentication service
│   └── chat-client/       # Chat client
├── internal/              # Internal packages
│   ├── auth/             # Authentication logic
│   ├── chat/             # Chat logic
│   └── rest-auth/        # REST auth logic
├── proto/                 # Protocol Buffer definitions
├── configs/               # Configuration files
├── scripts/               # Utility scripts
└── docs/                  # Documentation
```

## 🔧 Build Commands

### Available Make Targets

- `make build-all` - Build all services
- `make build-auth` - Build authentication service
- `make build-chat` - Build chat service  
- `make build-rest-auth` - Build REST auth service
- `make build-client` - Build chat client
- `make clean` - Clean build artifacts
- `make deps` - Install dependencies
- `make test` - Run tests
- `make fmt` - Format code
- `make help` - Show available commands

### Building Individual Services

```bash
# Build specific service
make build-auth

# Build all services
make build-all

# Clean build artifacts
make clean
```

## ⚠️ Important Notes

### Binary Files
**NEVER commit binary files to Git!** They are automatically ignored by `.gitignore`.

- Binary files are built to the `bin/` directory
- Use `make clean` to remove build artifacts
- Binary files are platform-specific and should be built locally

### Development Workflow

1. **Write code** in the appropriate packages
2. **Build locally** using `make build-*` commands
3. **Test** with `make test`
4. **Commit source code** (not binaries)
5. **Push to remote** - Git will ignore binary files automatically

## 🐳 Docker Support

```bash
# Start all services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f
```

## 📚 Documentation

- [Architecture Overview](docs/architecture.md)
- [Project Structure](docs/project-structure.md)
- [Implementation Notes](docs/implementation-notes.md)
- [Quick Start Guide](docs/quickstart.md)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Commit and push your changes
6. Create a pull request

## 📄 License

This project is licensed under the MIT License.
