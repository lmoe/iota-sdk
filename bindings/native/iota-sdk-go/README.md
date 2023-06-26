This is a PoC for now and untested.

Memory leaks will most likely build up as no controlled freeing of memory is in place.
Wrapper is very low level for now and does not offer fancy easy to use classes.

The native dll function `internal_listen_wallet` is not ported correctly yet.
It requires a handler callback and an array as parameters which are not yet adjusted.

It will most likely crash the application once called.