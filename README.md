# Unpredictable number generator #

This is a Go implementation of the same algorithm as implemented by
OpenBSD `arc4random()`.

## Naming ##

The package is called "unpredictable" to point out the fact that we
provide unpredictable numbers, not "random". My reasoning is that the
word "random" has been so badly misused within computer science that
it today means "perfectly predictable" with the added assumption that
each run of a program will produce the exact same sequence. Therefore
I chose the name "unpredictable" to point out that we are not
"random". The word "unpredictable" is hopefully less ambiguous and any
suggestion to make the returned numbers predictable can be deflected
by saying that the API is supposed to return unpredictable numbers,
it's right there in the name.

## API ##

Currently the only API provided is a math/rand.Source which we can get
by calling `unpredictable.NewMathRandSource()`. It provides the
interface that math/rand expects, but the Seed() function causes a
panic, since we want to provide unpredictable numbers, not random. I
want to do some thinking and preferably have some heavy user before I
start inventing an API. Maybe plugging it into math/rand is all the
API we need.

## Correctness ##

The implementation has been tested against OpenBSD arc4random. When
both generators are fed the exact same entropy they generate the same
byte sequence for at least 1680000 bytes (number chosen since the only
interesting special case in the OpenBSD arc4random happens at 1600000
bytes). Any bugs in here are likely to show up there too.

## Performance ##

When plugged into math.random each call to Int63 takes 40ns/op. This
compared to the default math.random source which take 10ns/op and
compared to buffered `C.arc4random` calls at 30ns/op. See [my test
repository for more details](https://github.com/art4711/randbench).


## ChaCha20 implementation ##

The chacha20 implementation is from
[https://github.com/codahale/chacha20](https://github.com/codahale/chacha20)
and heavily modified to make it do only what we need it to do.

## Entropy source ##

Our entropy comes from "crypto/rand" which is hopefully of high enough
quality (just slow).

## Additional safety ##

As opposed to the OpenBSD implementation we don't have fork detection
(not sure how to approach that in Go) and we don't have globally
exported functions so early startup is not ensured (might not be
possible) and locking is up to the user. Some, if not all, of these
problems will be solved when a proper API materializes.