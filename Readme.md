# chive

**_Status_**

This runs but doens't have much implemented, just starting up the idea.

---

A website starter with Go using the Chi router, MongoDB, and Vue for the client. This project compiles into one single binary that can be executed on whatever OS or architecture you build it for.

### Instructions

##### Development

Unfortunately, right now you need two terminals open:

- Run the client by changing to the client directory and executing "npm run dev"
- Run the server by running "go run *.go" from the project root

##### Production

Make sure to do things in this order. The fully built client will be incorported into the binary file.

- Build the client by changing to the client directory and executing "npm run build"
- In the project root directory, execute:
    rice embed-go
    go build