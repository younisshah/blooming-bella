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

### Example

```go
 bella, err := blooming_bella.NewBella(1000, 0.01)

	if err != nil {
		log.Fatal(err)
	}
	bella.Add(10)
	bella.Add(121)
	bella.Add(13)
	bella.Add(111)

	fmt.Println(bella.Test(10)) // => true
	fmt.Println(bella.Test(104)) // => false
	fmt.Println(bella.Test(110)) // => false
	fmt.Println(bella.Test(13)) // => true
 ```
 
### New 

Added __Super Fast Hashing__ algorithm

### TODO
 - [ ]  Calculate "ideal" number of hash functions to use.
 - [ ]  Dynamically "generate" the hash functions. *_What does it mean to be alive?_*
