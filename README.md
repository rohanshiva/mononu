## How to run?

First, run `make install` to install all the packages for the react app and install the go packages for the server.

Next, run `make build` to build the react app.

Now, run `make run` to run the server that spits out the following routes:

- `/` : Simple message indicating that the server is running successfully.

- `/static` : Renders the static site located under 
`staticapp` folder.
    - `/static/hello` : Returns a message (emulates some sort of service that may be used by the static app)

- `/reactapp/` : Renders the built react app located under `reactapp` folder.

- `/api/hello` : Returns a message (emulates some sort of API server)
