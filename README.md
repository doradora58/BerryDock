# Project BerryDock (プロジェクト・ベリードック)

3台のRaspberry Piを自宅ミニデータセンター（分散クラスター環境）に見立て、Go言語によるバックエンド開発と、モダンなコンテナ・クラウドネイティブ技術（Docker / Kubernetes / CI/CD / IaC）を完全無課金・ローカル完結で習得する実践プロジェクト。

---

## 🌐 1. クラスター構成（3台の役割分担）

* **1号機：司令塔 (Masterノード)**
  * Kubernetes管理、Argo CD（GitOps）、Prometheus/Grafana（監視インフラ）
* **2号機：実験環境 (Workerノード A)**
  * 開発・ステージング環境用のGoアプリ、PostgreSQLコンテナの配置先
* **3号機：本番環境 (Workerノード B)**
  * 本番環境用のGoアプリ、PostgreSQLコンテナの配置先

---

## 🛠️ 2. テクニカル・デプロイメント・ロードマップ

各段階のチェックボックスを埋めながら進捗を管理します。

- [ ] **Phase 1: システム情報取得APIサーバー**
  * **内容:** Goの `net/http` を使用し、ラズパイのOS/CPU情報をJSONで返す軽量APIの構築。
  * **成果物:** `main.go` / 各自のPCからのクロスコンパイル実績
- [ ] **Phase 2: 超軽量Dockerコンテナ化**
  * **内容:** Goのマルチステージビルドを駆使し、容量わずか数MBの極限まで最適化されたDockerイメージの作成。
  * **成果物:** `Dockerfile`
- [ ] **Phase 3: Docker Composeによる2層構成**
  * **内容:** GoアプリとPostgreSQLコンテナの連携。Volumeマウントによるデータの永続化。
  * **成果物:** `docker-compose.yml`
- [ ] **Phase 4: GitHub Actionsによる自動テスト (CI)**
  * **内容:** 変更を検知してクラウド上で `go test` を自動実行するパイプラインの構築。
  * **成果物:** `.github/workflows/test.yml`
- [ ] **Phase 5: ローカルRunnerによる自動ビルド**
  * **内容:** 1号機にGitHub ActionsのSelf-hosted Runnerを仕込み、ローカルでコンテナを自動ビルド。
  * **成果物:** 1号機内ビルド自動化設定
- [ ] **Phase 6: k3sによるマルチノードKubernetes環境**
  * **内容:** 軽量K8s（`k3s`）を用いて3台をネットワーク合体。複数ノードへのコンテナ自動分散配置。
  * **成果物:** K8sマニフェスト群（`deployment.yml`, `service.yml`）
- [ ] **Phase 7: Argo CDを用いたGitOps自動デプロイ (CD)**
  * **内容:** 1号機にArgo CDを導入。Git上のコード変更を2, 3号機の実環境へ自動同期。
  * **成果物:** Argo CDアプリケーション定義YAML
- [ ] **Phase 8: 環境変数を意識した複数環境管理**
  * **内容:** `Kustomize` を用いて、2号機（実験用）と3号機（本番用）の設定・DB接続先を安全に分離。
  * **成果物:** `kustomize/` ディレクトリ配下の差分コード
- [ ] **Phase 9: Prometheus/Grafanaによる3台一括監視**
  * **内容:** Goアプリにメトリクスを埋め込み。3台の温度・負荷・リクエスト数をGrafanaで可視化。
  * **成果物:** `prometheus.yml` / Grafanaダッシュボード設定
- [ ] **Phase 10: Terraformによる3台インフラのコード化 (IaC)**
  * **内容:** 3台にまたがるK8sやネットワーク環境をTerraformコード化。1行での環境構築・全削除。
  * **成果物:** `main.tf` / `variables.tf`

---

## 🥇 3. 接続するターゲット資格

本プロジェクトの実践を通して、以下の資格試験に直結するインフラ・開発知識を証明します。
1. **LinuC レベル1 / LPIC-1** (Linux操作・systemd・NW基礎)
2. **Go言語技術者認定試験 (Bronze / Silver)** (Go標準仕様・並行処理・WEBAPI)
3. **CKAD (Certified Kubernetes Application Developer)** (K8sリソース操作・コンテナ運用)
4. **AWS 認定ソリューションアーキテクト – アソシエイト (SAA)** (オンプレからクラウド設計への応用)

---

## 📂 4. ディレクトリ構造（Monorepo構成）

```text
berrydock/
├── .github/workflows/    # Phase 4-5: CI/CDパイプライン設定
├── phase01-info-api/     # Phase 1: Goソースコード
├── phase02-docker/       # Phase 2: Dockerfile
├── phase03-compose/      # Phase 3: docker-compose構成
├── phase06-k3s/          # Phase 6-7: K8sマニフェスト / ArgoCD設定
├── phase08-kustomize/    # Phase 8: 環境分離設定
└── phase10-terraform/    # Phase 10: IaCコード