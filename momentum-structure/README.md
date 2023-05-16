# Momentum GitOps Structure

This folder contains the general GitOps structure.

## Structure

The structure is as follows (generated with `tree -a -I`):

```bash
.
└── {root}
    ├── kustomization.yaml
    └── {application} # An application e.g. my-app (level 1)
        ├── _base
        │   ├── kustomization.yaml
        │   ├── ns.yaml
        │   ├── release.yaml
        │   ├── secrets.yaml
        │   └── values.yaml
        ├── _deploy
        │   └── {deployment}
        │       ├── kustomization.yaml
        │       ├── release.yaml
        │       ├── secrets.yaml
        │       └── values.yaml
        ├── _template
        │   ├── secrets.yaml # The helmrelease keys for secrets templated with go template
        │   └── values.yaml # The helmrelease values templated with go template
        ├── kustomization.yaml
        ├── ns.yaml # The namespace of the application (to store the helmrepository stuff)
        ├── repository.yaml # The helm repository of the application
        ├── {deployment}.yaml # A deployment of the application, which is a flux kustomization
        └── {stage} # A stage of the application e.g. cluster-test, cluster-prod (level 2)
            ├── _base
            │   ├── kustomization.yaml
            │   ├── release.yaml
            │   ├── secrets.yaml
            │   └── values.yaml
            ├── _deploy
            │   └── {deployment}
            │       ├── kustomization.yaml
            │       ├── release.yaml
            │       ├── secrets.yaml
            │       └── values.yaml
            ├── _template
            │   ├── secrets.yaml
            │   └── values.yaml
            ├── kustomization.yaml
            ├── {deployment}.yaml
            └── {stage} # A stage of the application e.g. test,prod (level 3)
                ├── _base
                │   ├── kustomization.yaml
                │   ├── release.yaml
                │   ├── secrets.yaml
                │   └── values.yaml
                ├── _deploy
                │   └── {deployment}
                │       ├── kustomization.yaml
                │       ├── release.yaml
                │       ├── secrets.yaml
                │       └── values.yaml
                ├── _template
                │   ├── secrets.yaml
                │   └── values.yaml
                ├── kustomization.yaml
                └── {deployment}.yaml
```

## Kustomize

The structure is based on [kustomize](https://kustomize.io/). The `kustomization.yaml` files are used to build the kubernetes manifests.

Usage:

```bash
kustomize build --load_restrictor LoadRestrictionsNone momentum
```
