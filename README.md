### Interfaces

#### Overview

We usually talk about interfaces in terms of behavior. Any object, for example, can implement the `fmt` package's `Stringer` interface, by implementing a `String()` method, which returns a string.

Let's begin with a quick look at how Go implements interfaces.

#### How Go Interfaces Work

See Russ Cox's blog post here: https://research.swtch.com/interfaces, and the Go FAQ https://go.dev/doc/faq#nil_error

Go is a typed language, and Go interfaces are types.

Interfaces are implemented as a struct that has two fields:

* A pointer to its underlying concrete type
* A pointer to that type's data

**Compile Time**

The Go compiler checks whether a concrete type implements the interface it purports to implement.

**Runtime**

At runtime, Go caches the lookup of the concrete type's function inside the interface in what Russ Cox calls an `itable`, so that subsequent lookups are faster.

**Assignment**

Interfaces are useful only where they contain concrete objects. You can create an interface, and never use it, and the compiler will not complain. The section in this repo called `assign-concrete-object-to-interface` shows the three ways we assign a concrete object to an interface.

#### Why Interfaces are Useful

Here are some of the ways interfaces are useful:

* **Create an application**, like a server or worker, whose base object has fields all of whose types are interface types. See `implementation-examples/application-server`. Using interfaces types for struct fields is useful to:
  * Constrain each field's behavior. Our base object's logger field may use a package that has 100 exported functions, we may prefer to create a Logger interface that contains only 5 of those functions, and therefore constrain our application to use only those 5.

  * Swap out the underlying implementation of each base object field, without changing our base object's implementation. Suppose our base object has a field of type Respository (interface) that represents all its required data access. We can swap out the underlying concrete object from a database, to a filestore, without changing our implementation, so long as our filestore implements the same list of function signatures contained in our Repository interface. This is of course crucial for mocking.

  * We can often describe a small application's behavior very concisely, merely through listing a set of well-defined and well-named interfaces at the top of the application.

* Implement behavior defined in other packages.
  * `implementation-examples/one-method-interface-package` shows a simplified version of the `fmt` package's `Stringer()` interface, where any object that has a `String()` method, will print nicely where passed to one of many `fmt` methods

  * `implementation-examples/multi-method-interface-package` is a simplified explanation of `sort.Interface` pattern, where any object that implements a collection of functions can be passed into a package's primary function, like `sort.Sort`, to provide powerful behavior to any object.

* Create a framework. Pulling together all the above, we can create frameworks that allow powerful functionality, by first building nearly all the functionality we need, exposing one or more interfaces we require to be implemented. See `./server-framework`
