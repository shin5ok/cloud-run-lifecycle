# ベースイメージを指定
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係をコピーし、インストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# 実行するコマンドを指定
CMD ["./main"]

