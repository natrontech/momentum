# Momentum GitOps Structure

This folder contains the general GitOps structure.

## Structure

The structure is as follows (generated with `tree -a -I`):

```bash
.
└── momentum-root # This is the root of the repository
    ├── kustomization.yaml # This is the root kustomization file
    └── {application}
        ├── _base
        │   ├── kustomization.yaml
        │   ├── ns.yaml
        │   ├── release.yaml # This is the helm release for base
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
        ├── ns.yaml
        ├── repository.yaml
        ├── {deployment}.yaml
        └── {stage}
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
            └── {stage}
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
