# go-nagios-config

# Overview

Nagios has a lot of [Object Definitions](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/4/en/objectdefinitions.html), these definitions are what you use to define your configurations. There are a lot of options on each object definition, and many times you are referencing other definitions within a definition. This package encompasses the basic object definitions, templates, etc. The package was meant to be a useful way to build up your configuration with code and write out the rendered files, which then could be used in a Nagios deployment.

## Installation

```
go get github.com/yzguy/go-nagios-config
```

## Upgrading

```
go get -u github.com/yzguy/go-nagios-config
```

## Getting Started

TODO

## Contributing

Contributing pull requests are gladly welcomed for this repository.

## License

MIT License

Copyright (c) 2018 Adam Smith

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
