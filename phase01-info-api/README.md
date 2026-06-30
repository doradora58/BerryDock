# Phase 1: System Info API

このディレクトリは、Goで作成したシンプルなシステム情報取得APIのフェーズ1用です。

## 1. 前提条件

- Go 1.20以降がインストールされていること
- `phase01-info-api` ディレクトリに移動して作業すること

## 2. 実行手順

1. ターミナルで作業ディレクトリに移動します。

```powershell
cd c:\work\01_source\BerryDock\phase01-info-api
```

2. Goアプリを実行します。

```powershell
go run .
```

3. サーバーが `:8080` で起動するので、別ターミナルで確認します。

## 3. 動作確認

### ヘルスチェック

```powershell
curl http://localhost:8080/health
```

期待値:

```json
{"status":"ok"}
```

### システム情報取得

```powershell
curl http://localhost:8080/info
```

期待値: JSON形式でOS/CPU情報が返ります。

例:

```json
{
  "hostname": "raspberrypi",
  "os": "...",
  "architecture": "arm",
  "cpu_model": "ARMv7 Processor rev 4 (v7l)",
  "cpu_count": 4
}
```

## 4. バイナリをビルドする場合

ローカル実行用バイナリを生成するには:

```powershell
go build -o info-api
```

生成された `info-api` を実行します。

```powershell
./info-api
```

## 5. Raspberry Pi（ARM）向けクロスコンパイル

Raspberry Pi向けにビルドする場合は次のように実行します。

```powershell
$env:GOOS = "linux"
$env:GOARCH = "arm"
$env:GOARM = "7"
go build -o info-api-armv7
```

Windows PowerShellでは、`$env:VARIABLE = "value"` の形式で環境変数を設定します。
