FROM golang:1.24

# モジュールを使用して依存関係を管理
ENV GO111MODULE=on

# ワーキングディレクトリの設定
WORKDIR /app

# 依存関係のインストール
COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download

# ソースコードとairの設定ファイルのコピー
COPY backend/ .
COPY backend/.air.toml ./

# .envファイルのコピー
COPY .env ./

# airのインストール
RUN go install github.com/air-verse/air@latest

# ポートを開放
EXPOSE 8080

# airでアプリケーションを起動
# CMD ["air"]
CMD ["/go/bin/air"]
