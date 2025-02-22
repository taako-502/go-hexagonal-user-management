# Backend

## 下準備

`dev`と言う名前のデータベースをあらかじめ作成しておく。

```sql
CREATE DATABASE dev
```

最初のマイグレーションを実行する。

```bash
sql-migrate up
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
