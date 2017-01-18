# gyn

Gyn is a web application for form filling of examination of patients in a maternity hospital.

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
