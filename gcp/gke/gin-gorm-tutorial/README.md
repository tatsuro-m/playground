## これは何
`go/gin-gorm-tutorial` のところで作成したアプリケーションのコピーです。  
GKEの学習の為にアプリをデプロイしたかったが、GCP の提供しているアプリでは簡単なAPIだけでつまらないので以前作成した gin の API をそのまま利用しました。  
それが理由なのでコピー元と sync する必要はありませんし、ソースコードがずれても問題ありません。

コピー元は GCP AppEngine を使ってデプロイしていましたが今回は GKE なので関連する設定ファイルは全て削除してあります。  
terraform のディレクトリも削除したので新たに作成してください。  

## やりたいこと
本番環境のワークフローで想定されるようなものは一通り体験しておきたいです。  
k8s の学習が主目的ですが余裕があれば [DataDog](https://www.datadoghq.com/ja/?utm_source=Advertisement&utm_medium=GoogleAdsNon1stTierBrand&utm_campaign=GoogleAdsNon1stTierBrand-JPNLangEN&utm_content=Brand&utm_keyword=%2Bdatadog&utm_matchtype=b&gclid=CjwKCAjwo4mIBhBsEiwAKgzXOP0_C_DS7eFCP22B3uqQzLt-_vHMLzMBSgz-8PvLHKsUt4FqABs29hoCJAEQAvD_BwE) 
も触ってみてください。

以下に実現したいことのリストを箇条書きで示します。順番に意味はありません。
- 今回作成するインフラリソースは基本的に全て terraform でコード化する
  - 終わったら丸ごと消しますが、またコマンド1発で再構築する為にもコード化は重要です
- terraform は環境ごとにディレクトリを切る構成として、クラスタもそれぞれ分割する
- GKE クラスタはリージョンクラスタで作成する
- マルチテナンシーを考慮してクラスタ内は名前空間でアプリごとに論理的に区切るものとする
- Cloud SQL はパブリックIPを持つものとする
- GKE から Cloud SQL への接続は Cloud SQL Auth Proxy を利用する。その際の認証情報は `Workload Identity` で共有する
- クラスタの起動タイプは Autopilot ではなく Standard で行う
- GitHub Actions で CI/CD パイプラインを構築する
- CI では静的ツールやテストの実行など標準的なものは一通りやるようにしてください
- CD は feature ブランチの段階で stg 環境にデプロイして良い
- [ArgoCD](https://argoproj.github.io/argo-cd/) を利用して GitOps を実現する
  - `Git で管理されているマニフェストの状態 = GKE で動いている k8s の構成` となるようにする
- DB migration を手動ではなく完全自動化する
- Docker イメージタグに latest を使うのは禁止し、git のコミットハッシュを利用する
- VPA, HPA, CA などを駆使して適切なオートスケーリングを設定する
- DB への接続情報などは k8s の Secret リソースを使用しセキュアに扱えるようにする（パブリックリポジトリなんで気をつけて）
- コンテナイメージの脆弱性スキャン機能を利用する
- DataDog で監視系を入れる
- Docker イメージキャッシュを使って CI/CD の時間をなるべく短くする
- helm で複数環境にデプロイできるようにする

やりたいことは随時書き足していってOK。

## Kubernetes 関連の補足
- 複数環境に展開できるようにアプリケーションは helm chart で管理する
```shell
$ helm install test-deploy -f mychart/stg-values.yaml ./mychart/ --create-namespace -n gin-gorm-tutorial
```
のように名前空間はマニフェストファイルではなく helm コマンドレベルで作成することとする

- chart は Argo CD から更新の度に自動適用されるように設定する（手動ではなくデフォルトの自動設定を利用してOK）
- Argo CD のセットアップは https://argo-cd.readthedocs.io/en/stable/getting_started/ を参考に行い、自動化はしない。手元から apply してパスワード変更する必要がある。
  - https://argoproj.github.io/argo-cd/operator-manual/declarative-setup/ を参考に yml を使って設定したいが、後回しで良い
- Argo CD の画面へは Service を割り当てるのではなく `kubectl port-forward` を使ってローカルホスト経由でアクセスする方法を取る
  - 現状のお試しレベルだと一般公開するのは不安が残るため
  - `kubectl port-forward svc/argocd-server -n argocd 8080:443`

### デプロイ戦略
色々と考えられるし正解も無いだろうが、小さく継続的なリリースを優先して以下のような構成で作ってみる。
1. feature ブランチで作業して変更を積んでいく
2. レビュー後に PR をマージすると CI/CD が起動。マージコミットのハッシュ値でタグ付けされたイメージが作成されてイメージレジストリへ push される。
3. イメージの push が終わったら自動で main にコミットを積む（PR作成 + 自動マージでも良い）。内容は 2 で作成したイメージをデプロイするように「検証環境の」マニフェストファイルのタグを書き換えただけのもの。
4. 3と同時に同じ容量で「本番環境用の」マニフェストファイルにも同じ変更を施したPRを自動生成する。ただしこちらは main にマージしない。この間に検証環境には最新イメージ（= main ブランチ）がデプロイされているので動作確認
5. 検証環境で確認して問題なければ 4 で作成された PR のマージボタンを押して本番環境にも適用する
6. もし途中で問題があればマージコミットを `revert -m 1` して元に戻す。もし本番環境に適用してから問題が発覚した場合には作成された本番向けのPRを直ちにマージするイメージ。検証環境にだけ適用された状態で戻したいなら revert 後に古い本番向けの PR を close すれば良い

この場合意識しなくてはいけないのはチームの自己組織化と小さくリリースを行うというカルチャー。  
開発を行ったチームが検証環境での動作確認から問題があった時の対応まで責任を持つことが重要。  
開発者は複数人いるが最終的に動作確認をしてデプロイする責任者が別人である場合にはこの方法はすぐに破綻する。  

モノリシックなリリースとなると検証環境での動作確認に時間がかかることになり、その修正などでさらに時間がかかる。  
その結果、本番環境と検証環境のソースコードがズレる時間が長くなってしまいチーム全体の開発体験が悪くなる。

もちろん自動テストの整備や Docker を使って環境差異を無くすというのは大前提。
