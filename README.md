# LOCAL-CHAT-MESSENGER

ソケット通信を利用して、クライアントサーバー間で情報をやり取りするシンプルなアプリケーション。
クライアントはサーバーにメッセージを送信し、サーバーは github.com/jaswdr/faker パッケージを使用してランダムに生成された文章で応答する。

## Prerequisites

- Go 1.21.1+
- [github.com/jaswdr/faker](https://github.com/jaswdr/faker)パッケージ

## Installation

1. 本リポジトリをクローン
1. 依存関係をインストール

```zsh
$ go mod tidy
```

## Usage

### Running the Server

サーバー側を起動する

```zsh
cd server
go run main.go
```

### Running the Client

クライアント側を起動する

```zsh
cd client
go run main.go
```

### Interaction

1. クライアントのターミナルにメッセージを入力。
2. サーバーはそのメッセージを受信し、ランダムに生成した文章で返信する。
3. クライアントはサーバーからの返信を表示する。

## License

This project is licensed under the MIT License.
