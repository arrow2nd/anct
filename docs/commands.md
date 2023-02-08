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

## search

### works

```
anct search works [keyword] [flags]
```

作品を検索します

#### flags

- `-e` `--editor`
  - キーワードの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数
- `-s` --season strings`
  - 放送シーズン : `YYYY-{spring|summer|autumn|winter}`

### characters

```
anct season characters [keyword] [flags]
```

キャラクターを検索します

#### flags

- `-e` `--editor`
  - キーワードの入力に外部エディタを使用する
- `-l` `--limit int`
  - 表示件数

## library

```
anct library [flags]
```

### flags

- `-s` `--state string`
  - 視聴ステータス : `{wanna_watch|watching|watched|on_hold|stop_watching|no_state}`
