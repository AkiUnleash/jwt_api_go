# React-Typescript Template

## 概要

React-Typescrpt 開発環境構築を簡略化するために、テンプレートを作成。

## 内容

- React
- Typescript
- Redux-toolkit
- Webpack
- Sass
- Dotenv

## 使い方

ローカルサーバーの起動

```
yarn start
```

ビルド

```
yarn build
```

## 環境変数の使い方

プロジェクトの直下に[ .env ]ファイルを作成。

```
TESTDATA="環境変数のテストデータ"
```

Typescript で以下のように記載して、データを取得できる。

```
process.env.TESTDATA
```
