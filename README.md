# anct

📺 Unofficial CLI Client of [Annict](https://annict.com/)

[![release](https://github.com/arrow2nd/anct/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/release.yml)
[![test](https://github.com/arrow2nd/anct/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/test.yml)
[![CodeQL](https://github.com/arrow2nd/anct/actions/workflows/codeql.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/codeql.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/anct)](https://goreportcard.com/report/github.com/arrow2nd/anct)
[![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/anct/total)](https://github.com/arrow2nd/anct/releases)

> [English](./README_EN.md)

![anct-demo](https://user-images.githubusercontent.com/44780846/220039050-c19a0545-0028-4511-841d-cf4e930f2dea.gif)

## 推奨環境

- sixel に対応したターミナル (画像表示に必要です)
- tmux 内で Kitty 画像プロトコル対応ターミナル (Ghostty / kitty 等) を使う場合は、tmux の設定に `set -g allow-passthrough on` が必要です

## できること

- 作品の検索
- 視聴ステータスの更新
- エピソードの視聴記録
- レビューの投稿

## インストール

> [!Warning]
>
> 以下の方法以外でインストールした場合、クライアントトークンが内蔵されていません
>
> 作成方法は [こちら](#クライアントトークン) をご覽ください

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

[Releases](https://github.com/arrow2nd/anct/releases)
からお使いの環境にあったものをダウンロードしてください

## 初期設定

以下のコマンドを実行して Annict と連携してください

```
anct auth login
```

## ドキュメント

- [コマンド一覧](./docs/ja/commands.md)

## Develop

### クライアントトークン

https://annict.com/oauth/applications から作成することができます

設定は以下の通りです

- リダイレクト URI : `urn:ietf:wg:oauth:2.0:oob`
- スコープ : 読み込み + 書き込み

### API クライアントのコードを生成

[Annict の個人用アクセストークン](https://annict.com/settings/apps) (スコープ :
読み込み) を環境変数に設定して

```
export ANNICT_KEY=<API Key>
```

以下を実行してください

```
make generate
```
