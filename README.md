###blooming-belle - A simple implementation of Bloom Filters 

### What?
 Bloom filter is a space efficient, probabilistic data structure, designed to test the membership of elements to a set.
 
### Trade-offs? 
 Being a space efficient data structure is it may return false positives, but always returns definite negatives.
 
### Applications?
 Testing for non-membership saves resources such as calls to a web server, checking a proxy cache. Google Chrome uses bloom filters as a check for malicious URLs.
 
### blooming-bella

 > A bloom filter for integers.
 > Uses __mummur3__,__Super Fast Hash__ and __marvin32__ hashing algorithms
 
### New 

Added __Super Fast Hashing__ algorithm

### TODO
 - [ ]  Calculate "ideal" number of hash functions to use.
 - [ ]  Dynamically "generate" the hash functions. *_What does it mean to be alive?_*