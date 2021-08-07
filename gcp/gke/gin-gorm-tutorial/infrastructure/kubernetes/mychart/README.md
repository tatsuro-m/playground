## これは何
https://helm.sh/docs/topics/charts/  
を参考にして汎用的なチャートを作ってみます。  

環境名等を外から入れることによって複数環境ですぐに展開できるような helm チャートの構築を目指します。  

公開はするものの他のユーザに使ってもらうことは想定していない勉強用なので、バージョニングに関しては適切に行っていない可能性が高いです。

---
```shell
$ helm install test-deploy -f mychart/stg-values.yaml ./mychart/ --create-namespace -n gin-gorm-tutorial
```
このようにチャートの install 時に名前空間を作成するようにして、マニフェストファイルからリソースとして名前空間を作成しないようにする。
よってチャート内での名前空間の参照は、
```yaml
namespace: {{ .Release.Namespace }}
```
のようになる。
