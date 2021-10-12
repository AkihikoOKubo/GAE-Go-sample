# gae-go-sample
GAE+Goで簡単なAPIを立ち上げるサンプルです。

## WHAT
GAEで簡単なAPIを実行するサンプルです。

- RuntimeはGo116です
- Firestoreで簡単な読み書きを行うサンプルコードを含みます
- Cloud Buildでビルド&デプロイできるyamlのサンプルを含みます

## HOW TO RUN
GCP環境であることが前提です

1. gcloudコマンドを実行できるようにしてください
2. Firebase及びFirestoreを有効化してください
3. GAEを有効化してください
4. IAMを開き、GAEのサービスアカウントに「Firebase管理者」のロールを付与してください
5. GAEにデプロイします
   1. `make deploy`

## ディレクトリ構成
```
- /src
    - /handler
        - /input
            - request bodyをマッピングする構造体
        - /output
            - response bodyをマッピングする構造体
        - router.go
            - APIのルーティング
    - /domain
        - /model
            - ドメインモデル
        - /repository
            - infra層(DBやFirestoreなど)のinterface
    - /usecase
            - ビジネスロジックの実装
    - /infra
            - repositoryの実装
    - /registry
            - repositoryに実装をインジェクトして返す
    - /lib
            - logging、HTTPリクエストなど共通基盤
    - /main
        - app_{ENV}.yaml(ビルド時にenvconfigで出し分けます)
        - main.go 
```

## サンプルAPIの内容
ユーザオブジェクトの追加・取得のサンプルです

### POST /users
- ユーザをUpsertします
- Post body
  - `{"id":1,"name":"test user","age":20}`

### GET /users/{id}
- ユーザを取得します

## Firebase Authentication
Firebase Authenticationを使ってユーザ認証するサンプルがあります。router.goを書き換えて試してみてください。
フロントでidTokenを発行し`AppAuthorization`ヘッダにtokenを設定してリクエストしてください。

## IAP対応
GAEサービス間通信で、IAPを使ってセキュアにHTTPリクエストするサンプルを書いてあります。試すには、以下の作業が必要です。
- IAPを有効化し、/src/main/app_{ENV}.yamlの`APP_OTHERSERVICE_AUDIENCE`にCLIENT_IDを入力する
- IAMを開き、GAEを実行するサービアカウントにidTokenを発行できるロール(serviceAccountTokenCreator)を付ける

## テストのサンプル
テストのサンプルを書いてあります。
1. repositoryのインターフェースを変更した場合は、[mockgen](https://github.com/golang/mock)をインストールしてmockを生成し直してください
2. `make test`