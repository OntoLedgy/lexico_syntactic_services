# Introduction

These services are a component of the go_bCLEARer toolkit. They provide
syntactic analysis of a set of strings in a dataset (e.g. a column). The
syntactic analysis is one of the first steps in the bCLEARer evolve
stage for string-based columns.

# Design

## Syntactic Check Types

The syntactic checker performs a series of checks based on the
configuration file. These checks can be performed across the entire
column or within a cell. When performed for the whole column, they
provide an overview of the characters used across all columns. When
performed within a cell they can cover the entire cell or just a part of
it. The full check type classification is as follows (this is a partial
taxonomy illustrating the broad structure):

![](code/object_model/2804908033.png)

## String Check Patterns

these can be defined as fixed strings, regexs or antlr grammar rules.

# Fixes

The syntactic checker can also implement fixes to the checked data.

for doing this, it uses the [string editor services](https://github.com/OntoLedgy/string_editing_services).
