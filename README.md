
- [Slack Daily Digest](#slack-daily-digest)
  - [Status](#status)
  - [Getting started](#getting-started)
    - [Environment](#environment-variables)
    - [How to run](#how-to-run)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Slack Daily Digest

This Project solves the issue of missing out on important threads in slack 
channels. You get daily digests of all the important threads of the channels you subcribed to.

## Getting started

### Environment
Please see [.env.example](https://github.com/konrad-amtenbrink/slack-daily-digest/blob/main/.env.example)

### How to run
#### First, install [air](https://github.com/cosmtrek/air) via
```zsh
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```
#### then check if it is installed with
```zsh
make version
```
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

#### To run debug mode with hot reload just simply run
```zsh
make debug
```
