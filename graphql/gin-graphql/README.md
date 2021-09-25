- GraphQL Server (gqlgen)
- Gin API Server

を組み合わせたマイクロサービス構成を作ってみる。  
余裕があれば簡単なフロントエンドアプリも。

Gin API はクライアントから直接叩くことはせず、 GraphQL 経由で叩いてフロントに返す。  
認証は全て GraphQL で行い、 API には jwt トークンを渡して認可を行う。
