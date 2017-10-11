# hostsconfig

An ultimate hosts file utility.


## Project Status

- This project is currently in its early stage, please backup your hosts file before using this tool.


## Try out

- Please download executable binaries for your architecture from [here](https://github.com/vishaltelangre/hostsconfig/releases).


## Usage

```
$ hostsconfig -h
```

```
  hostsconfig [OPTIONS]

  The OPTIONS are:
    -l, --list           List hosts file on standard output
    -b, --beautify       Display beautified hosts file on standard output
    -a=, --add=          Add a host entry
    -d=, --delete=       Delete a host entry
    -s, --save           Save changes
    -f=, --path=         Path to hosts file (Default: /etc/hosts)
    -h, --help           Show this usage help
    -v, --version        Display version
```


## Examples

- To list well-formatted hosts file:

```
hostsconfig -l
```

- To see a human readable, and beautified hosts file:

```
hostsconfig -b
```

- To use a different hosts file, (by default uses `/etc/hosts`):

```
hostsconfig --path="/tmp/hosts_copy"
```

- Adding new hosts:

```
hostsconfig -l -a="127.0.0.1 example1.com"
hostsconfig -l -a="192.168.0.3 example2.com example2.local"
hostsconfig -b -a="ff02::1 foo-ipv6.com"
```

- Delete existing hosts:

```
hostsconfig -l -d="127.0.0.1 example1.com"
hostsconfig -l -d="127.0.0.1 example1.com my-site.com foo.com"
hostsconfig -l -d="192.255.255.255"
hostsconfig -l -d="::ff02::3"
hostsconfig -l -d="example2.com bar.com"
```

- Saving changes:

```
sudo hostsconfig -b -a="ff02::1 foo-ipv6.com" -s
sudo hostsconfig -l -s -d="example2.com bar.com"
```


## TODO

- [x] Parse & validate hosts file
- [x] Preview hosts file in standard format, as well as in human understandable beautified format
- [x] Provision add/delete entry in hosts file
- [x] Allow saving modified, well-formatted changes in-place in hosts file
- [ ] Provision editing an existing entry
- [ ] Add unit tests to make it production ready
- [ ] Ship working cross-platform (Unix, Windows) executable
- [ ] And other features described [here](https://github.com/vishaltelangre/life/issues/1)


## Development

#### To build locally while development

```
go build
```

#### To generate cross-platform executables

```
env GOOS=darwin GOARCH=386 go build -o hostsconfig-darwin
env GOOS=linux GOARCH=arm GOARM=7 go build -o hostsconfig-linux
```


## Specifications

- Please RTFM [here](http://man7.org/linux/man-pages/man5/hosts.5.html).


## Copyright and License

Copyright (c) 2015, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).

<a target='_blank' rel='nofollow' href='https://app.codesponsor.io/link/PfwgcRiC73ERAe1WTDUo4DmM/vishaltelangre/hostsconfig'>
  <img alt='Sponsor' width='888' height='68' src='https://app.codesponsor.io/embed/PfwgcRiC73ERAe1WTDUo4DmM/vishaltelangre/hostsconfig.svg' />
</a>
