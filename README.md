# Check vendor

This is my solution. I decided not using an api framework (such as Gin, Echo, Fiber, etc) and created it using the standard library.

The endpoints are accepting all the HTTP method types. I haven't added a check on this because at this point, I feel that I have took too much time already working on the API.

For the tests, I like to use the `assert` and `mock` packages from `github.com/stretchr/testify` library. In my opinion, using the `assert` package helps writing the  assertion for the test cases. The `mock` package helps creating faster mock or stub structs. But, they are totally optional, and I could live without those packages.

## Project structure

I'm not sure what you will evaluate about the overall structure, but I created the frontend application inside the same repository. For sure, in a real application, I wouldn't do that, I would have a separate git repository for that. Also, I haven't added dockerfiles, or docker compose files. It would be welcome to start local environment faster, but I haven't done that in order to try to obey the time constraint on the PDF.

## Frontend project

The frontend project I did as quick as possible, I didn't want to take too much time. I tried to be simple. 

The URL of the API is hardcoded, it wasn't parameterized and I'm aware of that. Again, just to try to obey the time constraint on the PDF.  

I've picked the `date-fns` package to format the timestamp returned from the API. 

I haven't added the `all-status` endpoint in the frontend application, just because it was a different response format, and I would need to make a non technical decision. 
As I created a `VendorCard` component, to display the data of each vendor, my idea was to have another component, something like `VendorCardList`, where this one would handle the `all-status` response, and create a `VendorCard` for each vendor in the response. It would take me to refactor the `VendorCard`, since it handles the request to the API internally.

## Running the project

To start the api server, run the following command in the root folder:

```bash
go run main.go
```

The API server will run at port 8000. 

To start the UI application, run the following command in the `frontend` folder:

```bash
npm i
npm start
```

The frontend will run at port 3000.





