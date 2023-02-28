# go-cadence-example

go-cadence-example is a playground for testing capabilities of [candence](https://github.com/uber/cadence)
focusing on sagas.

## Architecture overview

We are going to implement simple booking system tickets consisting of services:

- order - receving requests for a new bookings
- payment - processing the payment oprations
- reservation - keeping track of seats available

The execution chain will look as follow:
order -> payment -> reservation

## Toolsets

We will use the cadence as a saga orchastrator for the communication.
We will use the golang for services implementation, utilising the [go-kit](https://github.com/go-kit/kit)
library.

https://sultanov.dev/blog/orchestration-based-saga-using-cadence/
https://medium.com/stashaway-engineering/building-your-first-cadence-workflow-e61a0b29785