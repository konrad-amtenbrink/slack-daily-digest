- [Golang Template Project](#golang-template-project)
  - [About the project](#about-the-project)
    - [Status](#status)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [How to run](#how-to-run)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Slack Daily Digest

## About the project

This Project solves the issue of missing out on important threads in slack 
channels. You get daily digests of the channels you subcribed to.

### Status

The template project is in development status.

## Getting started

Below we describe the conventions or tools specific to Slack Daily Digest.

### Layout

```tree
📦slack-daily-digest
┣ 📂.github
┃ ┗ 📂workflows
┃ ┃ ┗ 📜main.yml
┣ 📂handlers
┃ ┗ 📜event.go
┣ 📂logic
┃ ┣ 📂_slack
┃ ┃ ┣ 📜command.go
┃ ┃ ┗ 📜message.go
┃ ┗ 📂cron
┃ ┃ ┗ 📜cron.go
┣ 📂templates
┃ ┗ 📜message.json
┣ 📜.env
┣ 📜.gitignore
┣ 📜README.md
┣ 📜fly.toml
┣ 📜go.mod
┣ 📜go.sum
┗ 📜main.go
```

## How to run
### First, install [air](https://github.com/cosmtrek/air) via
```zsh
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
```
### then check if it is installed with
```zsh
make version
```
### To build just simply run
```zsh
make build
```

### To start just simply run
```zsh
make run
```

### To run dev mode with hot reload just simply run
```zsh
make dev
```

### To run debug mode with hot reload just simply run
```zsh
make debug
```