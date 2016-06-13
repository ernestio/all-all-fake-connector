# VCLOUD FAKER

This library aims to be a fake responder for all vcloud microservices. It allows you to bypass all vcloud calls at the same time you can simulate errors.

## Build status

* Master:  [![CircleCI Master](https://circleci.com/gh/ErnestIO/all-all-fake-connector/tree/master.svg?style=svg&circle-token=627e89c447fe342aff9815ca146b081a37c075ad)](https://circleci.com/gh/ErnestIO/all-all-fake-connector/tree/master)
* Develop: [![CircleCI Develop](https://circleci.com/gh/ErnestIO/all-all-fake-connector/tree/develop.svg?style=svg&circle-token=627e89c447fe342aff9815ca146b081a37c075ad)](https://circleci.com/gh/ErnestIO/all-all-fake-connector/tree/develop)

## Configuration

This software is based on a config.json file, where you define the transitions to bypass.

## Running it

To run vcloud fakery:

```
NATS_URI=nats://localhost:4222 vcloud-fakery -transitions=transitions/default.json
```

If no transitions config is passed in, the default transition will be loaded (default.json).

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

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
