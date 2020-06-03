<h1 align="center"><img height="250" src="https://raw.githubusercontent.com/negrel/timer/master/.github/timer.svg"></h1>

<p align="center">
	<a href="https://pkg.go.dev/github.com/negrel/timer">
		<img src="https://godoc.org/github.com/negrel/timer?status.svg">
	</a>
	<a href="https://goreportcard.com/badge/github.com/negrel/timer">
		<img src="https://goreportcard.com/badge/github.com/negrel/ringo">
	</a>
	<a href="https://github.com/negrel/timer/raw/master/LICENSE">
		<img src="https://img.shields.io/badge/license-APACHE-blue">
	</a>
</p>

# :timer_clock: Timer - CLI timer
*Easy to use CLI timer*   
Timer is CLI timer built with [gnotify](https://github.com/negrel/gnotify/) (for cross-platform native notifications) and
[cobra](https://github.com/spf13/cobra) (cli framework).

## Features

- **Easy to use**: Check the [examples](https://github.com/negrel/timer#example)
- **Notifications**: A **native** notification show up when the time is up.

## Installation

Using **go get** :

```bash
go get -u github.com/negrel/timer.git
```

## Example
```
# Start a 25 seconds timer.
timer 25s

# Start a 5 minutes and 30 seconds timer.
timer 1m30s

# Start a 1 hour and 15 minutes timer.
timer 1h15m
```

## :stars: Show your support

Please give a :star: if this project helped you!

## :scroll: License

APACHE-2.0 © Alexandre Negrel
