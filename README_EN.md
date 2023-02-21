# anct

💊 Unofficial CLI Client of [Annict](https://annict.com/)

[![CodeQL](https://github.com/arrow2nd/anct/actions/workflows/codeql.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/codeql.yml)
[![release](https://github.com/arrow2nd/anct/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/anct/actions/workflows/release.yml)
[![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/anct/total)](https://github.com/arrow2nd/anct/releases)

> [日本語](./README.md)

![anct-demo](https://user-images.githubusercontent.com/44780846/220039050-c19a0545-0028-4511-841d-cf4e930f2dea.gif)

## Recommended

- Terminal with sixel support (Required to display images)

## Features

- Searching for works
- Update your watching status
- Record your episode watchings
- Create a review

## How to Install

> **Warning**
>
> No built-in client token if installed by other than the following methods

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

### Binaries

Download the appropriate one for your environment from [Releases](https://github.com/arrow2nd/anct/releases)

## Initial Setup

Execute the following commands to link with Annict

```
anct auth login
```

## Documents

- [Commands](./docs/en/commands.md)

## Develop

### Generate API Client Code

```
export ANNICT_KEY=<API Key>
make generate
```
