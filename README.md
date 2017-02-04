# gyn

Gyn is a web application for form filling of examination of patients in a maternity hospital.

The backend of this project implemented with:

- https://github.com/labstack/echo as a web framework;
- https://github.com/spf13/viper to load configuration;
- https://github.com/asaskevich/govalidator to validate http parameters;
- https://github.com/pkg/errors to handle errors;
- https://github.com/dgrijalva/jwt-go to create and verify tokens;
- https://github.com/nguyenthenguyen/docx to fill docx template;
- https://github.com/fatih/structs as an utility to work with Go structs;
- https://godoc.org/golang.org/x/crypto/bcrypt to hash passwords;

The frontend implemented with:

- http://riotjs.com as elegant component-based UI library;
- http://getbootstrap.com for styling UI components;
- https://github.com/whatwg/fetch as implementation of the standart Fetch specification;
- https://momentjs.com to parse, validate, manipulate and display dates;
- https://babeljs.io as a compiler for writing next generation JavaScript;
- https://webpack.github.io to organize code, manage dependencies and build web UI;

## Getting started

### Prerequisites

To install developer tools run:

```
./install-dev-tools.sh
```

To build application run:

```
./build.sh [target operating system] [target architecture]
```

Target operating system and architecture parameters have values
as described in [specification] (https://golang.org/pkg/go/build/).
In the absence of one of them, the backend will build for current
operating system and architecture.

For example, to build:

- for Linux x64 `./build.sh linux amd64`

- for Linux x32 `./build.sh linux 386`

- for Windows x64 `./build.sh windows amd64`

- for Windows x32 `./build.sh windows 386`

- for current OS `./build.sh`