# Momentum UI

## Setup

```bash
npm install # install dependencies
npm run build # compile frontend
```

The above produces `build` output directory which is then used by PocketBase to serve the frontend of your app.

## Live Development

```bash
# start the backend, if not already running ...
npm run dev:backend
# and then start the frontend ...
npm run dev
```

Now visit http://localhost:5173 (ui) or http://localhost:8090 (pb)

## Generated Types

In `src/lib/momentum-core-client` you can find the generated client for the momentum-core backend. If there is a new version and you have to recreate the client, just call `npm run generate:api` which first generates the latest openapi spec and the generates a typescript client from the spec.

## Building

To create a production version of your app (static HTML/JS app):

_NOTE_: The build below will fail unless the backend has at least 1
post created. So please create a "posts" record using the app UI or
the admin UI before running build below.

```bash
# compile frontend
npm run build
# and then serve the backend
npm run backend
```

The above generates output in the `build` folder. Now you can serve production compiled version of the frontend using the backend (with `--publicDir ../frontend/build`), any static file web server, or `npm preview`.
