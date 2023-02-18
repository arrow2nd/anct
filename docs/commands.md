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

## info

```
anct info [<query>] [flags]
```

作品の詳細を出力します

#### flags

- `-e` `--editor`
  - クエリの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数を指定する
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - ライブラリ内を検索対象にする
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - 放送シーズンを指定する

## status

```
anct status [<query>] [flags]
```

作品の視聴ステータスを更新します

### flags

- `--state {wanna_watch|watching|watched|on_hold|stop_watching|no_state}`
  - 視聴ステータス

---

- `-e` `--editor`
  - クエリの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数を指定する
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - ライブラリ内を検索対象にする
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - 放送シーズンを指定する

## review

```
anct review [<query>] [flags]
```

作品のレビューを作成します

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

---

- `-e` `--editor`
  - クエリの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数を指定する
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - ライブラリ内を検索対象にする
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - 放送シーズンを指定する

## record

```
anct record [<query>] [flags]
```

エピソードの視聴記録を作成します

- まとめて記録することもできます
  - コメントは投稿されません

### flags

- `-r` ``--rating {great|good|average|bad}`
  - 評価
- `-c` ``--comment string`
  - コメント
- `-u` `--unwatch`
  - 視聴中作品の未視聴エピソードから選択する

---

- `-e` `--editor`
  - クエリの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数を指定する
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - ライブラリ内を検索対象にする
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - 放送シーズンを指定する
