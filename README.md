# Learn Go with Tests

<p align="center">
  <img src="red-green-blue-gophers-smaller.png" />
</p>

[Art by Denise](https://twitter.com/deniseyu21)

[![Go Report Card](https://goreportcard.com/badge/github.com/wilsonify/learn-go-with-tests)](https://goreportcard.com/report/github.com/wilsonify/learn-go-with-tests)

## Formats

- [Gitbook](https://quii.gitbook.io/learn-go-with-tests)
- [EPUB or PDF](https://github.com/wilsonify/learn-go-with-tests/releases)

## Translations

- [中文](https://studygolang.gitbook.io/learn-go-with-tests)
- [Português](https://larien.gitbook.io/aprenda-go-com-testes/)
- [日本語](https://andmorefine.gitbook.io/learn-go-with-tests/)
- [한국어](https://miryang.gitbook.io/learn-go-with-tests/)
- [Türkçe](https://halilkocaoz.gitbook.io/go-programlama-dilini-ogren/)

## Support me

I am proud to offer this resource for free, but if you wish to give some appreciation:

- [Tweet me @quii](https://twitter.com/quii)
- [Buy me a coffee :coffee:](https://www.buymeacoffee.com/quii)
- [Sponsor me on GitHub](https://github.com/sponsors/quii)

## Why

* Explore the Go language by writing tests
* **Get a grounding with TDD**. Go is a good language for learning TDD because it is a simple language to learn and testing is built-in
* Be confident that you'll be able to start writing robust, well-tested systems in Go
* [Watch a video, or read about why unit testing and TDD is important](why.md)

## Table of contents

### Go fundamentals

1. [Install Go](install-go.md) - Set up environment for productivity.
2. [Hello, world](hello-world.md) - Declaring variables, constants, if/else statements, switch, write your first go program and write your first test. Sub-test syntax and closures.
3. [Integers](integers.md) - Further Explore function declaration syntax and learn new ways to improve the documentation of your code.
4. [Iteration](iteration.md) - Learn about `for` and benchmarking.
5. [Arrays and slices](arrays-and-slices.md) - Learn about arrays, slices, `len`, varargs, `range` and test coverage.
6. [Structs, methods & interfaces](structs-methods-and-interfaces.md) - Learn about `struct`, methods, `interface` and table driven tests.
7. [Pointers & errors](pointers-and-errors.md) - Learn about pointers and errors.
8. [Maps](maps.md) - Learn about storing values in the map data structure.
9. [Dependency Injection](dependency-injection.md) - Learn about dependency injection, how it relates to using interfaces and a primer on io.
10. [Mocking](mocking.md) - Take some existing untested code and use DI with mocking to test it.
11. [Concurrency](concurrency.md) - Learn how to write concurrent code to make your software faster.
12. [Select](select.md) - Learn how to synchronise asynchronous processes elegantly.
13. [Reflection](reflection.md) - Learn about reflection
14. [Sync](sync.md) - Learn some functionality from the sync package including `WaitGroup` and `Mutex`
15. [Context](context.md) - Use the context package to manage and cancel long-running processes
16. [Intro to property based tests](roman-numerals.md) - Practice some TDD with the Roman Numerals kata and get a brief intro to property based tests
17. [Maths](math.md) - Use the `math` package to draw an SVG clock
18. [Reading files](reading-files.md) - Read files and process them
19. [Templating](html-templates.md) - Use Go's html/template package to render html from data, and also learn about approval testing
20. [Generics](generics.md) - Learn how to write functions that take generic arguments and make your own generic data-structure
21. [Revisiting arrays and slices with generics](revisiting-arrays-and-slices-with-generics.md) - Generics are very useful when working with collections. Learn how to write your own `Reduce` function and tidy up some common patterns.

### Build an application

Now that you have hopefully digested the _Go Fundamentals_ section you have a solid grounding of a majority of Go's language features and how to do TDD.

This next section will involve building an application.

Each chapter will iterate on the previous one, expanding the application's functionality as our product owner dictates.

New concepts will be introduced to help facilitate writing great code but most of the new material will be learning what can be accomplished from Go's standard library.

By the end of this, you should have a strong grasp as to how to iteratively write an application in Go, backed by tests.

* [HTTP server](http-server.md) - We will create an application which listens to HTTP requests and responds to them.
* [JSON, routing and embedding](json.md) - We will make our endpoints return JSON and explore how to do routing.
* [IO and sorting](io.md) - We will persist and read our data from disk and we'll cover sorting data.
* [Command line & project structure](command-line.md) - Support multiple applications from one code base and read input from command line.
* [Time](time.md) - using the `time` package to schedule activities.
* [WebSockets](websockets.md) - learn how to write and test a server that uses WebSockets.

### Questions and answers

I often run in to questions on the internets like

> How do I test my amazing function that does x, y and z

If you have such a question raise it as an issue on github and I'll try and find time to write a short chapter to tackle the issue. I feel like content like this is valuable as it is tackling people's _real_ questions around testing.

* [OS exec](os-exec.md) - An example of how we can reach out to the OS to execute commands to fetch data and keep our business logic testable/
* [Error types](error-types.md) - Example of creating your own error types to improve your tests and make your code easier to work with.
* [Context-aware Reader](context-aware-reader.md) - Learn how to TDD augmenting `io.Reader` with cancellation. Based on [Context-aware io.Reader for Go](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer)
* [Revisiting HTTP Handlers](http-handlers-revisited.md) - Testing HTTP handlers seems to be the bane of many a developer's existence. This chapter explores the issues around designing handlers correctly.

### Meta / Discussion

* [Why](why.md) - Watch a video, or read about why unit testing and TDD is important
* [Anti-patterns](anti-patterns.md) - A short chapter on TDD and unit testing anti-patterns

## Contributing

* _This project is work in progress_ If you would like to contribute, please do get in touch.
* Read [contributing.md](https://github.com/wilsonify/learn-go-with-tests/tree/842f4f24d1f1c20ba3bb23cbc376c7ca6f7ca79a/contributing.md) for guidelines
* Any ideas? Create an issue

## Background

I have some experience introducing Go to development teams and have tried different approaches as to how to grow a team from some people curious about Go into highly effective writers of Go systems.

### What didn't work

#### Read _the_ book

An approach we tried was to take [the blue book](https://www.amazon.co.uk/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) and every week discuss the next chapter along with the exercises.

I love this book but it requires a high level of commitment. The book is very detailed in explaining concepts, which is obviously great but it means that the progress is slow and steady - this is not for everyone.

I found that whilst a small number of people would read chapter X and do the exercises, many people didn't.

#### Solve some problems

Katas are fun but they are usually limited in their scope for learning a language; you're unlikely to use goroutines to solve a kata.

Another problem is when you have varying levels of enthusiasm. Some people just learn way more of the language than others and when demonstrating what they have done end up confusing people with features the others are not familiar with.

This ends up making the learning feel quite _unstructured_ and _ad hoc_.

### What did work

By far the most effective way was by slowly introducing the fundamentals of the language by reading through [go by example](https://gobyexample.com/), exploring them with examples and discussing them as a group. This was a more interactive approach than "read chapter x for homework".

Over time the team gained a solid foundation of the _grammar_ of the language so we could then start to build systems.

This to me seems analogous to practicing scales when trying to learn guitar.

It doesn't matter how artistic you think you are, you are unlikely to write good music without understanding the fundamentals and practicing the mechanics.

### What works for me

When _I_ learn a new programming language I usually start by messing around in a REPL but eventually, I need more structure.

What I like to do is explore concepts and then solidify the ideas with tests. Tests verify the code I write is correct and documents the feature I have learned.

Taking my experience of learning with a group and my own personal way I am going to try and create something that hopefully proves useful to other teams. Learning the fundamentals by writing small tests so that you can then take your existing software design skills and ship some great systems.

## Who this is for

* People who are interested in picking up Go.
* People who already know some Go, but want to explore testing with TDD.

## What you'll need

* A computer!
* [Installed Go](https://golang.org/)
* A text editor
* Some experience with programming. Understanding of concepts like `if`, variables, functions etc.
* Comfortable with using the terminal

## Feedback

* Add issues/submit PRs [here](https://github.com/wilsonify/learn-go-with-tests) or [tweet me @quii](https://twitter.com/quii)

[MIT license](LICENSE.md)

[Logo is by egonelbre](https://github.com/egonelbre) What a star!


## Go Directories

### `/cmd`

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/myapp`).

Don't put a lot of code in the application directory. 
If you think the code can be imported and used in other projects, then it should live in the `/pkg` directory.

If the code is not reusable or if you don't want others to reuse it, 
put that code in the `/internal` directory. 

You'll be surprised what others will do, 
so be explicit about your intentions!

It's common to have a small `main` function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

See the [`/cmd`](cmd/README.md) directory for examples.

### `/internal`

Private application and library code. 
This is the code you don't want others importing in their applications or libraries. 

Note that this layout pattern is enforced by the Go compiler itself. 

See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. 

Note that you are not limited to the top level `internal` directory. 

You can have more than one `internal` directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. 

It's not required (especially for smaller projects), 
but it's nice to have visual clues showing the intended package use. 

Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

### `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). 

Other projects will import these libraries expecting them to work, 
so think twice before you put something here :-) Note that the `internal` directory is a better way to ensure your private packages are not importable because it's enforced by Go. 

The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. 

The [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) blog post by Travis Jeffery provides a good overview of the `pkg` and `internal` directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: 
[`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) and [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

See the [`/pkg`](pkg/README.md) directory if you want to see which popular Go repos use this project layout pattern. 

This is a common layout pattern, but it's not universally accepted and some in the Go community don't recommend it.

It's ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value (unless you really want to :-)). 

Think about it when it's getting big enough and your root directory gets pretty busy (especially if you have a lot of non-Go app components).

The `pkg` directory origins: The old Go source code used to use `pkg` for its packages and then various Go projects in the community started copying the pattern (see [`this`](https://twitter.com/bradfitz/status/1039512487538970624) Brad Fitzpatrick's tweet for more context).

### `/vendor`

Application dependencies (managed manually or by your favorite dependency management tool like the new built-in [`Go Modules`](https://github.com/golang/go/wiki/Modules) feature). The `go mod vendor` command will create the `/vendor` directory for you. Note that you might need to add the `-mod=vendor` flag to your `go build` command if you are not using Go 1.14 where it's on by default.

Don't commit your application dependencies if you are building a library.

Note that since [`1.13`](https://golang.org/doc/go1.13#modules) Go also enabled the module proxy feature (using [`https://proxy.golang.org`](https://proxy.golang.org) as their module proxy server by default). 

Read more about it [`here`](https://blog.golang.org/module-mirror-launch) to see if it fits all of your requirements and constraints. 

If it does, then you won't need the `vendor` directory at all.

## Service Application Directories

### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

See the [`/api`](api/README.md) directory for examples.

## Web Application Directories

### `/web`

Web application specific components: static web assets, server side templates and SPAs.

## Common Application Directories

### `/configs`

Configuration file templates or default configs.

Put your `confd` or `consul-template` template files here.

### `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

### `/scripts`

Scripts to perform various build, install, analysis, etc operations.

These scripts keep the root level Makefile small and simple (e.g., [`https://github.com/hashicorp/terraform/blob/master/Makefile`](https://github.com/hashicorp/terraform/blob/master/Makefile)).

See the [`/scripts`](scripts/README.md) directory for examples.

### `/build`

Packaging and Continuous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.

Put your CI (travis, circle, drone) configurations and scripts in the `/build/ci` directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the `/build/ci` directory linking them to the location where the CI tools expect them (when possible).

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh). Note that in some repos (especially apps deployed with kubernetes) this directory is called `/deploy`.

### `/test`

Additional external test apps and test data. Feel free to structure the `/test` directory anyway you want. For bigger projects it makes sense to have a data subdirectory. For example, you can have `/test/data` or `/test/testdata` if you need Go to ignore what's in that directory. Note that Go will also ignore directories or files that begin with "." or "_", so you have more flexibility in terms of how you name your test data directory.

See the [`/test`](test/README.md) directory for examples.

## Other Directories

### `/docs`

Design and user documents (in addition to your godoc generated documentation).

See the [`/docs`](docs/README.md) directory for examples.

### `/tools`

Supporting tools for this project. Note that these tools can import code from the `/pkg` and `/internal` directories.

See the [`/tools`](tools/README.md) directory for examples.

### `/examples`

Examples for your applications and/or public libraries.

See the [`/examples`](examples/README.md) directory for examples.

### `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).

### `/githooks`

Git hooks.

### `/assets`

Other assets to go along with your repository (images, logos, etc).

### `/website`

This is the place to put your project's website data if you are not using GitHub pages.

See the [`/website`](website/README.md) directory for examples.

## Directories You Shouldn't Have

### `/src`

Some Go projects do have a `src` folder, but it usually happens when the devs came from the Java world where it's a common pattern. If you can help yourself try not to adopt this Java pattern. You really don't want your Go code or Go projects to look like Java :-)

Don't confuse the project level `/src` directory with the `/src` directory Go uses for its workspaces as described in [`How to Write Go Code`](https://golang.org/doc/code.html). The `$GOPATH` environment variable points to your (current) workspace (by default it points to `$HOME/go` on non-windows systems). This workspace includes the top level `/pkg`, `/bin` and `/src` directories. Your actual project ends up being a sub-directory under `/src`, so if you have the `/src` directory in your project the project path will look like this: `/some/path/to/workspace/src/your_project/src/your_code.go`. Note that with Go 1.11 it's possible to have your project outside of your `GOPATH`, but it still doesn't mean it's a good idea to use this layout pattern.

