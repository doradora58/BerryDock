# Phase 2: Docker Containerization

このディレクトリは、Phase 1 の Go アプリを Docker コンテナ化するためのフェーズ2用です。

## 1. 前提条件

- Docker が実行中であること
- `phase02-docker` の親ディレクトリでコマンドを実行すること

## 2. Docker イメージのビルド

リポジトリのルートから以下を実行します。

```powershell
cd c:\work\01_source\BerryDock
docker build -t berrydock-info-api -f phase02-docker\Dockerfile .
```

このコマンドは、`phase02-docker/Dockerfile` を使ってルートをビルドコンテキストとして指定します。

## 3. コンテナの起動

```powershell
docker run --rm -p 8080:8080 berrydock-info-api
```

## 4. 動作確認

別ターミナルで次を実行します。

```powershell
curl http://localhost:8080/health
curl http://localhost:8080/info
```

`/health` では `{"status":"ok"}` が返り、`/info` では JSON のシステム情報が返ります。

## 5. 目的

- Go アプリをコンテナとして実行可能にする
- ビルドステージと実行ステージを分離し、最終イメージを最小化する
- `scratch` を使うことで余分なランタイムを含まない軽量イメージを生成する
