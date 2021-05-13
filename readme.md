# ひとこと日記 (ローカル起動 Ver)

今の状況をひとことつぶやくアプリ。
※Golang を学習するために作成。

## デモ

![Untitled](https://user-images.githubusercontent.com/50258433/118069766-b6bf4780-b3df-11eb-89d5-79e1de9d19fd.gif)

## 機能

- サインアップ
- ログイン
- ひとこと日記投稿
- ログアウト

## 必須環境

- Docker

## インストール方法

### リポジトリクローン

```zsh
$ git clone https://github.com/AkiUnleash/jwt_api_go.git
$ cd jwt_api_go

```

### 起動

```zsh
$ docker-compose up

# 以下の状況になれば起動完了です。
api_1    | root:golang@tcp(mysql-container-jwt:3306)/golang_db
mysql-container-jwt | mbind: Operation not permitted
api_1    |
api_1    |    ____    __
api_1    |   / __/___/ /  ___
api_1    |  / _// __/ _ \/ _ \
api_1    | /___/\__/_//_/\___/ v4.3.0
api_1    | High performance, minimalist Go web framework
api_1    | https://echo.labstack.com
api_1    | ____________________________________O/_______
api_1    |                                     O\
api_1    | ⇨ http server started on [::]:8082
success Already up-to-date.
Done in 1.72s.
yarn run v1.21.1
$ webpack serve --open --config webpack.dev.js --host 0.0.0.0
app_1    | Failed to load ./.env.
app_1    | ℹ ｢wds｣: Project is running at http://0.0.0.0:5000/
app_1    | ℹ ｢wds｣: webpack output is served from /
app_1    | ℹ ｢wds｣: Content not from webpack is served from /usr/src/app/dist
app_1    | ℹ ｢wds｣: 404s will fallback to /index.html
app_1    | ℹ ｢wdm｣: asset main.js 5.16 MiB [emitted] (name: main)
app_1    | asset index.html 172 bytes [emitted]
app_1    | cached modules 1.89 MiB (javascript) 1.19 KiB (runtime) [cached] 355 modules
app_1    | webpack 5.37.0 compiled successfully in 2798 ms
app_1    | ℹ ｢wdm｣: Compiled successfully.
```

※ mySql の稼働が完了する前に、API が開始される場合エラーが発生する可能性がある。
その場合、再度コマンドを入力し起動する。

## システム・アーキテクチャ

### API (サーバーサイド)

- Golang
- echo
- Gorm
- uuid
- echo-swagger

※API 仕様書は起動後、以下 URL で確認できます。

[http://localhost:8082/swagger/index.html](http://localhost:8082/swagger/index.html)

### Database (サーバーサイド)

- MySQL

### app (フロントエンド)

- React
- Typescript

## 注意事項

- クローン後すぐに稼働できるよう、パスワード及び環境設定ファイルなども含めてコミットしています。

## ライセンス

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## 作者

[Aki Unleash - Akio Yano](https://github.com/AkiUnleash)
