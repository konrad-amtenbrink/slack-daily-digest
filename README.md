<!-- markdownlint-configure-file {
  "MD013": {
    "code_blocks": false,
    "tables": false
  },
  "MD033": false,
  "MD041": false
} -->
<div align="center">

[![Go Test](https://github.com/konrad-amtenbrink/slack-daily-digest/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/konrad-amtenbrink/slack-daily-digest/actions/workflows/go.yml)
# Slack Daily Digest

This project solves the issue of missing out on important threads in slack 
channels.

> :warning: **FYI**: Even though slack daily digest is basically finished, I never deployed it. I just used this project to learn the basics of go.

[Getting started](#getting-started) •
[Installation](#installation) •
[Development](#development)

</div>

## Getting started

### Environment
Please see [.env.example](https://github.com/konrad-amtenbrink/slack-daily-digest/blob/main/.env.example)

### Installation
#### First, install [air](https://github.com/cosmtrek/air) via
```zsh
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```

#### then check if it is installed with
```zsh
make version
```

## Run
### Development
#### To build just simply run
```zsh
make build
```

#### To start just simply run
```zsh
make run
```

#### To run dev mode with hot reload just simply run
```zsh
make dev
```
### Debug
#### To run debug mode with hot reload just simply run
```zsh
make debug
```
