# Momentum GitOps Structure

This folder contains the general GitOps structure.

## Structure

The structure is as follows (generated with `tree -a -I`):

```bash
.
└── momentum # The root folder
    ├── [application] # An application e.g. my-app
    │   ├── [deployment].yaml # A deployment of the application
    │   ├── [stage] # A stage of the application e.g. dev, staging, prod
    │   │   ├── [deployment].yaml # A deployment of the application in a stage
    │   │   ├── _base
    │   │   │   ├── kustomization.yaml
    │   │   │   ├── release.yaml # The flux helmrelease
    │   │   │   ├── secrets.yaml # The helmrelease values for secrets, encrypted with SOPS
    │   │   │   └── values.yaml # The helmrelease values
    │   │   ├── _deploy
    │   │   │   └── [deployment]
    │   │   │       ├── kustomization.yaml
    │   │   │       ├── release.yaml
    │   │   │       ├── secrets.yaml
    │   │   │       └── values.yaml
    │   │   ├── _template
    │   │   │   ├── secrets.yaml
    │   │   │   └── values.yaml
    │   │   └── kustomization.yaml
    │   ├── _base
    │   │   ├── kustomization.yaml
    │   │   ├── ns.yaml
    │   │   ├── release.yaml
    │   │   ├── repository.yaml
    │   │   ├── secrets.yaml
    │   │   └── values.yaml
    │   ├── _deploy
    │   │   └── [deployment]
    │   │       ├── kustomization.yaml
    │   │       ├── release.yaml
    │   │       ├── secrets.yaml
    │   │       └── values.yaml
    │   ├── _template
    │   │   ├── secrets.yaml # The helmrelease keys for secrets templated with go template
    │   │   └── values.yaml # The helmrelease values templated with go template
    │   └── kustomization.yaml
    └── kustomization.yaml
```
