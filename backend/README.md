# Backend

## 下準備

`dev`と言う名前のデータベースをあらかじめ作成しておく。

sql-migrate のインストールが必要かもしれない。

```bash
go install github.com/rubenv/sql-migrate/...@latest
```

```sql
CREATE DATABASE dev
```

## sql-migrate

### 新しいマイグレーションファイルを作成する

```bash
sql-migrate new {マイグレーション名}
```

### マイグレーションを実行する

```bash
sql-migrate up
```

### マイグレーションをロールバックする

```bash
sql-migrate down
```
