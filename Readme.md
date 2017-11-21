# chive

**_Status_**

I have refactored this project to use vuetify and it looks pretty now!

![Screenshot](https://i.imgur.com/51Tzimo.png)

---

A website starter with Go using the Chi router, MongoDB, and Vue for the client. This project compiles into one single binary that can be executed on whatever OS or architecture you build it for.

### Architecture

A key difference between this and other full stack projects is that I am not including server-side rendering. Instead, I am using templating on the server side to produce a unique HTML header for each path, then letting the javascript client load as per normal.

So if someone goes to http://website/ and then clicks on example 1234, they will simply make an API request to the server to fecth 1234. The server will respond with JSON data as you'd expect. However, if someone browses directly to http://website/examples/1234, the server will respond with the HTML inserted into the header that generates meta tags with pertinent information about the asset, and then the javascript app loads as usual and the vue-router loads up example 1234.

I think that this approach is a little better than SSR for PWAs because it provides the SSO and the "crawlability" that SSR gives you, but it's a lot easier. And since the service worker is keeping all the static assets on the device anyway, you don't have to worry about load times on subsequent visits to your site by actual people.

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


Lastly, you can serve this site over TLS and http2 (including push) by setting up a caddy server and reverse-proxying to your app. Caddy does it all for you magically. https://caddyserver.com

Here's an example caddyfile:
```
yourwebsite.com www.yourwebsite.com {
  tls you@yourwebsite.com
  push
  proxy / http://localhost:6000
  log foundirl.access.log
}
```