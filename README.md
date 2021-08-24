## 概要

Go(gqlgen)を利用した Graphql の dataloader サンプルです。

## 起動前提

ローカル mysql が存在して、
root ユーザがパスワード空文字列で設定されていること。

mysql で dataloader と言うデータベースを作成してください。

## 起動方法

User 起動

```
cd user
go run ./main.go
```

Todo 起動

```
cd todo
go run ./main.go
```

Gateway 起動

```
cd gateway
yarn
yarn run start
```

## お試し実行

以下 SQL を実行してテストデータを投入してください。

```
INSERT INTO `users` (`id`, `name`)
VALUES
	(1, 'taro'),
	(2, 'ao'),
	(3, 'namu');

INSERT INTO `todos` (`id`, `user_id`, `body`)
VALUES
	(1, 1, '明日やる'),
	(2, 1, '明後日やる'),
	(3, 2, '今日はやらない'),
	(4, 3, '多分やらない'),
	(5, 2, '１時間後にやる');
```

graphql_gateway の listen ポート(デフォルト 4000)に向けてクエリを実行すれば確認できます。

http://localhost:4000/

手元に GraphQL client がない場合は apollo studio から確認するのが容易です。

https://studio.apollographql.com/sandbox/explorer

```
query {
  todos{
    id
    userId
    body
    user{
      name
    }
  }
}
```
