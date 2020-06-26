# go-starter

`golang-starter-template` is an opinionated [golang](https://golang.org/) backend development template that includes setup with GORM ORM (Postgres, etc), Echo web framework for API's, Go-playground for validation.

## Overview

## Table of Contents

- [go-starter](#go-starter)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Usage](#usage)
    - [Requirements](#requirements)
    - [Quickstart](#quickstart)
  - [Contributing](#contributing)
  - [Maintainers](#maintainers)

## Features

- Development environment using [Docker Compose](https://docs.docker.com/compose/install/) and [VSCode devcontainers](https://code.visualstudio.com/docs/remote/containers) that works with Linux, MacOS and Windows.
- Adheres to the project layout defined in [golang-standard/project-layout](https://github.com/golang-standards/project-layout)
- Provides database ORM and abstraction for connecting to different databases.

TODO Future
- Integrates [go-swagger](https://github.com/go-swagger/go-swagger) for compile-time autogeneration of `swagger.json` and request/response validation functions.
- Integrates [MailHog](https://github.com/mailhog/MailHog) for easy SMTP-based email testing
- Parallel jobs optimized `Makefile` and various convenience scripts, a full rebuild via `make build` only takes seconds

## Usage

### Requirements

Requires the following local setup for development:

- [Docker CE](https://docs.docker.com/install/) (19.03 or above)
- [Docker Compose](https://docs.docker.com/compose/install/) (1.25 or above)


### Quickstart

After your `git clone` you may do the following:

```bash

# $local
# Easily start the docker-compose
docker-compose up -d
# Defaults to Postgres
# Create golang-starter-template db in postgres via localhost:8080

sh ./run-local.sh

```

## Contributing

Pull requests are welcome. For major changes, please [open an issue](https://github.com/ralberts/golang-starter-template/issues/new/choose) first to discuss what you would like to change.

## Maintainers

- [Ryan Alberts - @ralberts](https://github.com/ralberts)

