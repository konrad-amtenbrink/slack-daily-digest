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
### To build just simply run
```zsh
make build
```

### To start just simply run
```zsh
make run
```