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
>
> Please refer to [here](#Client Token) for how to create one

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

### Client Token

It can be created from https://annict.com/oauth/applications

The configuration is as follows

- Redirect URI : `urn:ietf:wg:oauth:2.0:oob`
- Scope : read + write

### Generate API Client Code

[Annict's personal access token](https://annict.com/settings/apps) (Scope : Read) to an environment variable and then

```
export ANNICT_KEY=<API Key>
```

Do the following

```
make generate
```
