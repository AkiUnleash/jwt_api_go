## ディレクトリ構成

### config

設定値等の処理を行うための反り

### domain

モデル及び routing

### repositories

データ単体での処理

### until

汎用的な関数群

## 登録

```
curl -X POST -H \
"Content-Type: application/json" -d \
'{"name":"a", "email":"a@a.com", "password": "1"}' \
http://127.0.0.1:8001/api/register
```

## ログイン

```
curl -X POST -H \
"Content-Type: application/json" -d \
'{"email":"a@a.com", "password": "1"}' \
http://127.0.0.1:8001/api/login
```
