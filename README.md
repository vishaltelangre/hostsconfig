# hostsconfig

Yet another hosts file manager

## Project Status [IMPORTANT]
In very early development stage, please don't use it directly unless you know what it does.

## Usage

```
$ hostsconfig -h
```

```
Hosts file manager

Usage:
  hostsconfig [OPTIONS]

The OPTIONS are:
  -sp, --standard-preview       Standard preview of file
  -pp, --pretty-preview         Pretty preview of file
  -f, --path                    Path to hosts file (Default: /etc/hosts)
  -h, --help                    Show this usage help
  -v, --version                 Display version
```

## Development

```
go build
```

## TODO

- [x] Parse & validate hosts file
- [x] Preview hosts file in standard format, as well as in human understandable beautified format
- [ ] Provision add/edit/delete entry in hosts file
- [ ] Ship working cross-platform (Unix, Windows) executable
- [ ] And other features described [here](https://github.com/vishaltelangre/life/issues/1)

## Copyright and License

Copyright (c) 2015, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).