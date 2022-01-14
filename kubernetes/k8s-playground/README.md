## これは何
Kubernetes の挙動を確認したい時にサクッと立ち上げて使用できる playground 的なクラスターを構築するもの。  

構成は好きに変えて良いが、それなりに実践を意識した構成にする。  
- `kubectl apply` で Pod を操作するのではなく、 ArgoCD にするとか。  
- グローバル static な IP を取得して Ingress に関連付ける
- １クラスターマルチテナント方式
- イメージタグは latest を使わない
- stg, prod の２つ以上の環境を用意
- 環境差異の吸収には kustomize を使う

とか。
