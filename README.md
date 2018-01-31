# Ernest.io fake connector

master:  [![CircleCI](https://circleci.com/gh/ernestio/all-all-fake-connector/tree/master.svg?style=shield)](https://circleci.com/gh/ernestio/all-all-fake-connector/tree/master)  
develop: [![CircleCI](https://circleci.com/gh/ernestio/all-all-fake-connector/tree/develop.svg?style=shield)](https://circleci.com/gh/ernestio/all-all-fake-connector/tree/develop)

This library aims to be a fake responder for all ernest.io providers. It allows you to bypass all external provider calls at the same time you can simulate errors.

## Configuration

This software is based on json file in order to determine the transitions it needs to respond to. The default file can be found at transitions/default.json, but you can simulate error environments with customized json files in that folder.

## Running it

To run vcloud faker:

```
NATS_URI=nats://localhost:4222 vcloud-fakery -transitions=transitions/default.json
```

If no transitions config is passed in, the default transition will be loaded (transitions/default.json).

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 ernest.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
