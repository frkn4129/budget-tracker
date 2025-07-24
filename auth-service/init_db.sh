#!/bin/bash

# ENV ayarları (gerekirse burada düzenleyebilirsin)
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="elektriklihayatlar"  # veya authdb

# Komutları çalıştırmak için psql’e parola aktarımı
export PGPASSWORD=$DB_PASSWORD

echo "📦 Connecting to PostgreSQL and creating 'users' table in $DB_NAME..."

docker exec -i evhaber-postgres psql -U postgres -d elektriklihayatlar <<EOF
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
EOF


if [ $? -eq 0 ]; then
  echo "✅ Users table created successfully!"
else
  echo "❌ Failed to create users table."
fi
