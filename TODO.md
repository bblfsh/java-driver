This driver is in alpha stage,
there are still many things that need to be done.

This document aims to give an overview of things
that still need to be implemented to have a complete driver.
It will hopefully get empty at some point in the future.

During the normalization stage,
the driver uses [JDT](driver/normalizer/jdt/jdt.go)
to analyze and annotate the native AST.

Ideally, we should have all these elements both properly annotated and tested.
Annotations are handled by [annotations code](driver/normalizer/annotation.go).
The status of the annotations support is documented within that code.
We also have some samples in the [tests](tests) dir.
Since there's no code where to document the status of these samples,
we're docummenting it here.

## Node types

Currently we have samples that generate most of the native node types,
but there are some missing:

* BlockComment
* LineComment
* NameQualifiedType
* ParameterizedType
* QualifiedType
* SingleMemberAnnotation
* TypeMethodReference
* UnionType
* WildcardType

Comments are not provided by default by the native driver. This should eventually be fixed.

## Structural properties

TBD

## Keywords

TBD
