# unleash-checkr

Checks if any flags have expired an notifies someware

![CI](https://github.com/unleash/unleash-checkr/workflows/CI/badge.svg)

# Install

### Pre-compiled binary

**shell script**

```sh
curl -sf https://gobinaries.com/unleash/unleash-checkr | sh
```

**manually**

Download the pre-compiled binaries from the [releases page](https://github.com/unleash/unleash-checkr/releases) and copy to the desired location

### Docker

```sh
$ docker run -it --rm unleashorg/unleash-checkr --help
```

### Compiling from source

Clone the repository

```sh
$ git clone git@github.com:unleash/unleash-checkr.git

$ cd gitignore
```

download dependencies

```sh
$ go mod download
```

build

```sh
$ go build -o unleash-checkr main.go
```

verify it works

```sh
$ unleash-checkr --help
```

# Usage

**check**

List overdue flags on terminal

```sh
$ unleash-checkr check -u "http://unleash.herokuapp.com"
```

**notify**

List overdue flags on terminal and notifies via Slack

```sh
$ unleash-checkr notify --channel "#checkr" --slack-token "my-token" --url "http://unleash.herokuapp.com" -e 60
```