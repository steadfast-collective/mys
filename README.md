# mys

A tiny cli tool to rapidly create, import and drop local MySQL databases.

## Installation
Clone this repo into your desired location.

Add said location to your `$PATH`

e.g Add the following to your `.zshrc / .bashrc`
```sh
$ export PATH="$PATH:$HOME/path/to/this/repo"
```

## Usage
```sh
$ mys <command> [parameter(s)]
```

For more info, please run:

```sh
$ mys --help
```

## Contributors
[Daniel Crewdson](https://www.github.com/crumb1e)

## Disclamer!
This tool is only intended to be used locally for speeding up common development tasks - please do not use this on a live/production environment.

We plan to add the ability to configure `mys` to work with other databases. But for now, we're sticking with MySQL whilst we iron out any issues.
