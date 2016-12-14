// Package godash provides tility functions for searching and manipulating slices in golang.
// Inspired by the Lodash library in Javascript.
package godash

// shared types

type validator func(interface{}) bool

type mutator func(interface{}) interface{}
