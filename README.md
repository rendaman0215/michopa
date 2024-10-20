# ディレクトリ構造

```sh
┣━cmd
┃  ┗━main.go
┃
┗━internal
   ┣━domain              // ドメイン層
   ┃  ┣━model            // ドメインモデル
   ┃  ┣━repository       // ドメインモデルの永続化
   ┃  ┗━service          // ドメインサービス
   ┃
   ┣━service             // サービス層（ユースケース）
   ┣━server              // connectについて
   ┣━infrastructure      // repoの実装など、インフラ周り
   ┣━interface_adapter   // データ変換系ロジック
   ┗━middleware          // logger, auth, trace, etc...
```
