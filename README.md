# Check vendor

This is my solution. I decided not using an api framework (such as Gin, Echo, Fiber, etc) and created it using the standard library.

The endpoints are accepting all the HTTP method types. I haven't added a check on this because at this point, I feel that I have took too much time already working on the API.

For the tests, I like to use the `assert` and `mock` packages from `github.com/stretchr/testify` library. Using the `assert` package in my opnion, helps writing the  assertion for the test cases. As the `mock` package helps creating faster mock or stub structs. But, they are totally optional, and I could live without those packages.






