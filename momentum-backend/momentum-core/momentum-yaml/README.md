# YAML support for the Go language

This fork of [go-yaml/yaml](https://github.com/go-yaml/yaml) provides 
features we need in natrium. Natrium is a gitops automation tool. In the
Following section we will explain when you should use this fork instead
of the official package (with high probability it doesn't make sense
for you to use our fork, but the original package instead). For Further
information about the parser implementation, we would like to refer to the 
documentation of the official go-yaml package. This documentation will
only contain the differences to the official documentation and what are
their intentions.

## Why use natrium-yaml and not go-yaml?

Our problem is that we want to have a fully dynamic tree-structure at runtime 
which mirrors any yaml document. go-yaml on the other hand assumes you have 
a clearly specified yaml structure which doesn't change at runtime or that a
certain process is applied to only one specific type of yaml configuration.

Of course you can implement different adapters which then handle all relevant
yaml structures but our idea is that if this fork here can handle its task
good enough we can spare this and work on a plain tree-structure.

## Use-Cases

Say we configure an aspect (e.g. hostname) of a service and have different configuration
structures which depend on the underlying technology (helm, flux or other cool stuff).

With this tool it shall be made easy to set the hostname in such yaml configurations, without
messing with different tools specifications.

## Our extension - Growing trees

The [YAML-Spec](https://yaml.org/spec/1.2.2/) defines a grammar which handles five different node types:
1. Scalar
2. Alias
3. Mapping
4. Sequence
5. Document

Copied from the YAML-Spec: 

```
A YAML node represents a single native data structure. Such nodes have content of one of three kinds: scalar, sequence or mapping. In addition, each node has a tag which serves to restrict the set of possible values the content can have.

Scalar

    The content of a scalar node is an opaque datum that can be presented as a series of zero or more Unicode characters.

Sequence

    The content of a sequence node is an ordered series of zero or more nodes. In particular, a sequence may contain the same node more than once. It could even contain itself.

Mapping

    The content of a mapping node is an unordered set of key/value node pairs, with the restriction that each of the keys is unique. YAML places no further restrictions on the nodes. In particular, keys may be arbitrary nodes, the same node may be used as the value of several key/value pairs and a mapping could even contain itself as a key or a value.

```

This grammar is implemented by go-yaml.





