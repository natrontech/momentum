<p align="center">
    <a href="https://momentum.natron.io">
        <img height="130px" src="assets/momentum-logo.png" />
    </a>
</p>

<p align="center">
  <strong>
    <a href="https://momentum.natron.io/">Momentum</a>
    <br />
		Propel your GitOps workflow
  </strong>
</p>

<p align="center">
  <a href="https://github.com/natrontech/momentum/issues"><img
    src="https://img.shields.io/github/issues/natrontech/momentum"
    alt="Build"
  /></a>
  <a href="https://github.com/natrontech/momentum"><img
    src="https://img.shields.io/github/license/natrontech/momentum"
    alt="License"
  /></a>
	<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/natrontech/momentum/main/momentum-backend?label=Go%20Version" />
	<img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/natrontech/momentum/ci.yml?label=CI" />
	<img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/natrontech/momentum/codeql.yml?label=CodeQL" />
	<img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/natrontech/momentum/docker-release.yml?label=Docker%20Release" />
</p>

<h2></h2>

Momentum is a next-generation GitOps as a Service platform, designed to simplify, accelerate, and automate your software delivery process. By embracing the power of GitOps, Momentum provides transparent, auditable, and easy-to-manage continuous deployment solutions.

*Powered by [Natron](https://natron.io)*

---

## Key Features

- **Flexible Deployments**: Manage deployments effortlessly across various stages in your pipeline.
- **Custom Stage Design**: Tailor your stage structure to align with your unique development and deployment procedures.
- **Application Catalogue Control**: Easily manage your application catalogue.
- **Personalized Helm Chart Integration**: Benefit from the flexibility to use your own Helm chart for tailored configurations.
- **Enhanced Observability**: Enhance your deployment oversight with comprehensive stage value tracking. Clearly visualize where deployment values originate from within various stages, promoting improved understanding and control.
- **GitOps Repository Integration**: Seamlessly integrate with your existing GitOps repositories for a smooth transition.
- **Efficient Rollouts with FluxCD**: Utilize FluxCD to enable efficient and reliable application rollouts.
- **Multi Cloud Compatibility**: Deploy applications across a variety of cloud providers or on-premise environments.

## Getting Started

### Prerequisites

- Kubernetes cluster (v1.23+)
- Helm (v3.0+)
- FluxCD (v2.0.0)

## Installation

### Helm

*tbd*

### Docker Compose

```bash
# clone the repo
git clone git@github.com:natrontech/momentum.git
cd momentum

# build the images
docker compose build

# in the docker-compose.yaml is a volume mounted for the backend pb_data
# this is where the database is stored, so you might want to create a directory for it or change the path
mkdir -p ./momentum-backend/pb_data

# start the containers
docker compose up -d
```

## Documentation

*tbd*

## Developing

### Dev Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Node.js](https://nodejs.org/en/download/) (v20+)
- [Go](https://golang.org/doc/install) (v1.20+)
- [modd](https://github.com/cortesi/modd/releases)

### Setup

Follow these steps CAREFULLY, or else it won't work. Also read the README files referred above before proceeding.

1. If using Docker then copy `.env.example` to `.env` and then edit it to match your environment. And then just run `docker compose up -d`. Without Docker, see below ...
2. Setup the backend in accordance with [./momentum-backend/README.md](./momentum-backend/README.md)
3. Setup the frontend in accordance with [./momentum-ui/README.md](./momentum-ui/README.md)

After you've done the setup in the above two README files, run
the backend and the frontend in dev mode (from `momentum-ui` directory).

```bash
# start the backend
npm run dev:backend
# and then start the frontend ...
npm run dev
```

Now visit http://localhost:5173 (sk) or http://localhost:8080 (core)

Now making changes in the Svelte code (frontend) or Go code (backend) will show
results (almost) immediately.

### Building

See the build process details in the README files for backend and frontend.

## Contributing

> **Note:** We currently have no contribution guidelines. This will be added in the future.

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests.
