# Go言語の勉強リポジトリ

Go言語の触りながら「レイヤードアーキテクチャとは何ぞや？」の疑問を解消するための勉強リポジトリ。

## 作ったもの

ユーザー情報の簡単なCRUD処理をもったAPI。

## 使ってるもの

- Go 1.17
- PostgreSQL 14.2
- フレームワーク [Gin](https://github.com/gin-gonic/gin)
- ORM [GORM](https://github.com/go-gorm/gorm)
- ホットリロード [Air](https://github.com/cosmtrek/air)

## 使用方法

```bash
# コンテナ起動
docker-compose up -d
```

## 参考サイト

- https://gorm.io/ja_JP/docs/connecting_to_the_database.html
- https://gorm.io/ja_JP/docs/query.html
- https://zenn.dev/nobonobo/articles/0b722c9c2b18d5
- https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode#go-get-%E3%81%AF%E3%82%AA%E3%83%AF%E3%82%B3%E3%83%B3%EF%BC%9F
- https://qiita.com/tono-maron/items/345c433b86f74d314c8d
- https://qiita.com/takehanKosuke/items/4342ca544d205fb36eb0
- https://qiita.com/egawata/items/7c18bbc54dc353034f5f
- https://zenn.dev/diwamoto/articles/aba45dc2da36b8
- https://qiita.com/marnie_ms4/items/e51cc6d879cc9ad07af3
- https://medium.com/geekculture/easily-run-your-unit-test-with-golang-gin-postgres-8a402a29e3f6
- https://www.getto.systems/entry/2020/08/09/190650