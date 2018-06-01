[![Godoc Reference](https://godoc.org/github.com/aead/blake2s?status.svg)](https://godoc.org/github.com/aead/blake2s)

## The BLAKE2s hash algorithm

BLAKE2s is a fast cryptographic hash function described in [RFC 7963](https://tools.ietf.org/html/rfc7693).
BLAKE2s can be directly keyed, making it functionally equivalent to a Message Authentication Code (MAC).

### Recommendation 
This BLAKE2s implementation was submitted to the golang x/crypto repo.
I recommend to use the official [x/crypto/blake2s](https://godoc.org/golang.org/x/crypto/blake2s) package if possible.

### Installation

Install in your GOPATH: `go get -u github.com/aead/blake2s`

### Performance

**AMD64**  
Hardware: Intel i7-6500U 2.50GHz x 2  
System: Linux Ubuntu 16.04 - kernel: 4.4.0-64-generic  
Go version: 1.8.0  
```
SSE4.1
name        speed           cpb
Write64-4  261MB/s ± 0%     9.14
Write1K-4  323MB/s ± 0%     7.38
Sum64-4    215MB/s ± 0%    11.09 
Sum1K-4    317MB/s ± 0%     7.52

SSSE3
name        speed           cpb
Write64-4  191MB/s ± 0%    12.48 
Write1K-4  220MB/s ± 0%    10.84
Sum64-4    165MB/s ± 0%    14.45
Sum1K-4    217MB/s ± 0%    10.99

SSE2
name        speed           cpb
Write64-4  177MB/s ± 0%    13.47
Write1K-4  201MB/s ± 0%    11.86
Sum64-4    154MB/s ± 0%    15.48
Sum1K-4    199MB/s ± 0%    11.98
```