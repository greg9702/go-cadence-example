# **go-cadence-example**

go-cadence-example is a playground for testing capabilities of [candence](https://github.com/uber/cadence)
focusing on using it as a saga ochestrator. 

> :warning: *This project is just playground/POC. It follows lot of bad practices (including lack of 
> interfaces, hard coded values etc.) which should  not be followed. Maybe one day I will find some time to fix*


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
