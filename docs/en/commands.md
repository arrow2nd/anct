# Commands

## auth

### login

```
anct auth login
```

Authentication with a Annict

### logout

```
anct auth logout
```

Log out of a Annict

## config

### client-token

```
anct config client-token
```

Change the API key used to connect to Annict

- If you are using pre-built binaries, you do not need to set this

## info

```
anct info [<query>] [flags]
```

Display information about the work

#### flags

- `-e` `--editor`
  - Use an external editor to enter text
- `-l` `--limit int`
  - Maximum number of results to fetch
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - Search within the library
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - Retrieve works for a given season

## status

```
anct status [<query>] [flags]
```

Update the watching status of work

### flags

- `--state {wanna_watch|watching|watched|on_hold|stop_watching|no_state}`
  - Update status state

---

- `-e` `--editor`
  - Use an external editor to enter text
- `-l` `--limit int`
  - Maximum number of results to fetch
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - Search within the library
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - Retrieve works for a given season

## review

```
anct review [<query>] [flags]
```

Review of the work

### flags

- `--overall-rating {great|good|average|bad}`
- `--movie-rating {great|good|average|bad}`
- `--character-rating {great|good|average|bad}`
- `--story-rating {great|good|average|bad}`
- `--music-rating {great|good|average|bad}`
- `--comment string`

---

- `-e` `--editor`
  - Use an external editor to enter text
- `-l` `--limit int`
  - Maximum number of results to fetch
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - Search within the library
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - Retrieve works for a given season

## record

```
anct record [<query>] [flags]
```

Record the watching of episode

- You can also record them all together.
  - In this case, comments are not recorded.

### flags

- `-r` `--rating {great|good|average|bad}`
- `-c` `--comment string`
- `-u` `--unwatch`
  - Select from the unwatched episodes of the work you are watching

---

- `-e` `--editor`
  - Use an external editor to enter text
- `-l` `--limit int`
  - Maximum number of results to fetch
- `-L` --library {wanna_watch|watching|watched|on_hold|stop_watching}`
  - Search within the library
- `-S` `--season YYYY-{spring|summer|autumn|winter}`
  - Retrieve works for a given season
