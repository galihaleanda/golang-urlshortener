## 🌐 Frontend Repository

The frontend for this project is maintained separately:
- #### Repository: https://github.com/galihaleandaaa/url-shortener-frontend

## 🔹 Installation

### 1️⃣ Clone Repository
```bash
git clone https://github.com/galihaleandaaa/golang-urlshortener.git
cd golang-urlshortener
```

### 2️⃣ Install Dependencies
```bash
go mod tidy
```

## ⚙️ Configuration

### 1️⃣ Environment Variables
Create a .env file in root directory
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=url_shortener
DB_SSLMODE=disable
PORT=8080
```

## 🗄 Database Setup
Create the PostgreSQL database:

```bash
CREATE DATABASE url_shortener;
```

Main table structure:
```bash
CREATE TABLE urls (
    id UUID PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    click_count INT DEFAULT 0,
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## ▶️ Running the Server
```bash
go run main.go
```

```bash
Server will start at
http://localhost:8080
```

## 📡 API Endpoints

### 1️⃣ Shorten URL

POST /shorten

Request Body : 

```bash
{
  "url": "https://example.com/long-url",
  "expiration_hour": 24
}
```

Response : 
```bash
{
  "short_code": "a1B2c3",
  "expires_at": "2026-02-20T15:19:32Z"
}
```

### 2️⃣ Redirect Short URL

GET /:short_code

##### Behavior: 
Redirects to the original URL
Returns 410 Gone if link expired


### 3️⃣ Analytics

GET /analytics/:short_code

Response
```bash
{
  "original_url": "https://example.com/long-url",
  "click_count": 5,
  "created_at": "2026-02-19T14:59:29Z",
  "expires_at": "2026-02-20T14:59:29Z"
}
```

## 🧪 Example Usage

#### Shorten a URL
```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url": "https://chat.openai.com", "expiration_hour": 48}'
```

#### Visit Short Link
http://localhost:8080/a1B2c3

#### Check Analytics
curl http://localhost:8080/analytics/a1B2c3

## 📝 Notes

- Expiration is optional
- All timestamps use UTC
- Click count does not track unique visitors
- Backend does not store IP addresses by default








