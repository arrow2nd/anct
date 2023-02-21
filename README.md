# anct

💊 Unofficial CLI Client of [Annict](https://annict.com/)

[![CodeQL](https://github.com/arrow2nd/anct/actions/workflows/codeql.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/codeql.yml)
[![release](https://github.com/arrow2nd/anct/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/release.yml)
[![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/anct/total)](https://github.com/arrow2nd/anct/releases)

> [English](./README_EN.md)

![anct-demo](https://user-images.githubusercontent.com/44780846/220039050-c19a0545-0028-4511-841d-cf4e930f2dea.gif)

## 推奨環境

- sixel に対応したターミナル (画像表示に必要です)

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

## ドキュメント

- [コマンド一覧](./docs/ja/commands.md)

## Develop

### Generate API Client Code

```
export ANNICT_KEY=<API Key>
make generate
```
