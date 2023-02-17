# Commands

## auth

### login

```
anct auth login
```

Annict との認証を行います

### logout

```
anct auth logout
```

Annict との認証を解除します

---

### 検索共通 flags

- `-e` `--editor`
  - クエリの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数を指定する
- `--library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - ライブラリ内を検索対象にする
- `--season YYYY-{spring|summer|autumn|winter}`
  - 放送シーズンを指定する

## info

```
anct info [<query>] [flags]
```

作品の詳細を出力します

#### flags

## status

```
anct status [<query>] [flags]
```

作品の視聴ステータスを更新します

### flags

- `--state {wanna_watch|watching|watched|on_hold|stop_watching|no_state}`
  - 視聴ステータス

## review

```
anct review [<query>] [flags]
```

作品のレビューを投稿します

### flags

- `--overall-rating {great|good|average|bad}`
  - 全体の評価
- `--movie-rating {great|good|average|bad}`
  - 映像の評価
- `--character-rating {great|good|average|bad}`
  - キャラクターの評価
- `--story-rating {great|good|average|bad}`
  - ストーリーの評価
- `--music-rating {great|good|average|bad}`
  - 音楽の評価
- `--comment string`
  - コメント

## track

```
anct track [<query>] [flags]
```

エピソードの視聴記録を投稿します

- まとめて記録することもできます
  - コメントは投稿されません

### flags

- `--id string`
  - 作品 ID
- `--episodes int`
  - 話数 (複数指定可能)
- `--rating {great|good|average|bad}`
  - 評価
- `--comment string`
  - コメント
