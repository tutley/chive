# chive

**_Status_**

I have refactored this project to use vuetify and it looks pretty now!

![Screenshot](https://i.imgur.com/51Tzimo.png)

---

A website starter with Go using the Chi router, MongoDB, and Vue for the client. This project compiles into one single binary that can be executed on whatever OS or architecture you build it for.

### Architecture

A key difference between this and other full stack projects is that I am not including server-side rendering. Instead, I am using templating on the server side to produce a unique HTML header for each path, then letting the javascript client load as per normal.

So if someone goes to http://website/ and then clicks on example 1234, they will simply make an API request to the server to fecth 1234. The server will respond with JSON data as you'd expect. However, if someone browses directly to http://website/examples/1234, the server will respond with the HTML inserted into the header that generates meta tags with pertinent information about the asset, and then the javascript app loads as usual and the vue-router loads up example 1234.

One thing I want to experiment with after I get this basically working is to improve efficiency by populating a javascript variable inside the header of the HTML document with the result of whatever the request is. So if the request is a list of examples, I will use the go template to write a javascript snippet in the HTML header that contains an array with a list of examples. Then the Vue app will have logic to check to see if a javascript object already exists, and if not go ahead and make the API call.

I beieve that this approach solves the use case of SSR without actually rendering the javascript app on the server.

### Instructions

##### Development

Unfortunately, right now you need two terminals open:

- Run the client by changing to the client directory and executing "npm run dev"
- Run the server by running "go run *.go" from the project root

##### Production

Make sure to do things in this order. The fully built client will be incorported into the binary file.

- Build the client by changing to the client directory and executing "npm run build"
- In the project root directory, execute:
    
```
rice embed-go
go build
```

PS - make sure to set your GOOS and GOARCH environment variables accordingly for the target server
