# anct

💊 Unofficial CLI Client of [Annict](https://annict.com/)

## できること

- 作品の検索
- 視聴ステータスの更新
- エピソードの視聴記録
- レビューの投稿

## インストール

> **Warning**
>
> 以下の方法以外でインストールした場合、クライアントトークンが内蔵されていません

### Homebrew

```
brew tap arrow2nd/tap
brew install anct
```

### Scoop

```
scoop bucket add arrow2nd https://github.com/arrow2nd/scoop-bucket.git
scoop install arrow2nd/anct
```

### バイナリ

[Releases](https://github.com/arrow2nd/anct/releases) からお使いの環境にあったものをダウンロードしてください

## 初期設定

以下のコマンドを実行して Annict と連携してください

```
anct auth login
```

## Develop

### Generate API Client Code

```
export ANNICT_KEY=<API Key>
make generate
```
