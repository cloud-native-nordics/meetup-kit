# meetup-kit

A toolbox to manage Cloud Native Communities by Pull Request -- MeetOps!

## Usage

```console
$ meetup-kit generate
```

generates the READMEs for e.g. the https://github.com/cloud-native-nordics/meetups repo

```console
$ meetup-kit serve
```

serves GraphQL query requests to act as a backend for e.g. the https://cloudnativenordics.com website
(available at https://stats-api.cloudnativenordics.com)

## Building

```console
$ make
```

The build requires docker, or alternatively, Go.

## License

[Apache v2](LICENSE)
