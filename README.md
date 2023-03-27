# URL-shortener

[![CI](https://github.com/greg-learns-go/url-shortener/actions/workflows/ci.yml/badge.svg)](https://github.com/greg-learns-go/url-shortener/actions/workflows/ci.yml)

This is a project to learn GO.

The objective: self-contained app, including a server, that can

## Goals

- [x] save a new URL in the database with a short identifier
- [ ] retrieve a shortened URL on GET, and redirect(*)
- [ ] covered with unit tests
- [ ] flags specifying domain/port

## WIP

- [ ] The Shrotener interface/struct is still a mess and requires some refactoring
- [ ] redirect is not implemented yet (for manual testing convenience), will be implemented when main module is covered with tests

## Quirks

The "hash function" is a silly approach, with unknown probability for collisions (probably bad)
This is not a decision I'd make in a production code.

## Bonus points

- [ ] compile an executable for other platforms? (Mac, Windows?)
- [ ] deploy the executable in some hosting?

## Used concepts

- ✓ Go basics (slices, loops, pointers etc.)
- ✓ Modules/packages (internal and external)
- ✓ net/http server
- ✓ database/sql
- ✓ func init
- ✓ defer
- ✓ go routines
- ✓ html templates
- ✓ dependency injection (using interfaces)
