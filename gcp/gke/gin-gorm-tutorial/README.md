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
