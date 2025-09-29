# Quiz-13 API

REST API untuk manajemen buku dan kategori menggunakan Go, Gin, dan PostgreSQL.

## Features

- Authentication dengan JWT
- CRUD Categories
- CRUD Books
- Database PostgreSQL
- Environment configuration

## Setup

### 1. Clone dan Install Dependencies

```bash
git clone <repository-url>
cd quiz-13
go mod tidy
```

### 2. Database Setup

1. Install PostgreSQL
2. Buat database:
   ```sql
   CREATE DATABASE db_quiz13;
   ```
3. Jalankan migration:
   ```bash
   # Gunakan tool migration atau jalankan file SQL langsung
   psql -U postgres -d db_quiz13 -f db/migrations/001_init.sql
   ```

### 3. Environment Configuration

1. Copy file `.env.example` ke `.env`
2. Sesuaikan konfigurasi database dan JWT secret

### 4. Jalankan Aplikasi

```bash
go run cmd/main.go
```

Server akan berjalan di `http://localhost:8090`

## API Endpoints

### Authentication

- `POST /api/users/login` - Login user

### Categories

- `GET /api/categories/` - Get all categories (requires auth)
- `POST /api/categories/` - Create category (requires auth)

### Books

- `GET /api/books/` - Get all books (requires auth)
- `POST /api/books/` - Create book (requires auth)

## Testing

### Login

```bash
curl -X POST http://localhost:8090/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"1234"}'
```

### Create Category (with token)

```bash
curl -X POST http://localhost:8090/api/categories/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"name":"Fiction"}'
```

### Create Book (with token)

```bash
curl -X POST http://localhost:8090/api/books/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "title":"Sample Book",
    "description":"A sample book",
    "release_year":2023,
    "total_page":150,
    "price":100000,
    "category_id":1
  }'
```

## Environment Variables

- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name (default: db_quiz13)
- `JWT_SECRET_KEY` - JWT secret key
- `PORT` - Server port (default: 8080)
