# Go Lisp Interpreter

This repository contains a small Lisp-like language interpreter implemented in Golang.

```lisp
(fmt.Println "Hello, World!")
```

## Features

This Go Lisp Interpreter offers the following features:

- [ ] **Easy Interoperability with Golang:** You can easily call Go code from within this Lisp language. Simply use `interp.Use(lisp.Symbols.All)` or `interp.Use(lisp.Symbols.Fmt)` to add symbols to the root scope. These example symbols are collections of stdlib functions exported as a map of functions. To mix Lisp and Go code, use `interp.EvalWithMap("(+ x 1)", lisp.ScopeMap{ "x": 1 })`, which evaluates to `int64(2)`. Alternatively, the method `Interpreter.EvalWithValues(string, any...)` allows you to provide a series of values that will be bound to `_0`, `_1`, and so on. It's important to note that these bindings do not override globals with the same name, as a new scope is created.

- [x] **Lists:** The language supports lists represented as `[1 2 3]`.

- [x] **Maps:** Maps can be created using the syntax `{ #name "John" #surname "Smith" }`.

- [x] **Property Chaining:** Property chaining is supported using the dot notation. For example, you can compute the length of a vector with `(math.Sqrt (+ (^ v.x 2) (^ v.y 2) (^ v.z 2)))`.

- [x] **Quasi-quotes:** Quasi-quotes are denoted by the symbols `#` for quoting and `$` for interpolation. For instance, `#(a b 1.0 [3 $(+ 3 1) 5] $(+ 1 2 3))` evaluates to `#(a b 1 [3 4 5] 6)`. This feature allows for convenient construction of data structures or DSLs. As an example, an HTML DSL could be defined as follows:

    ```lisp
    #(div { #class "card" }
        (img { #src $author.profile-image-url #alt "example image" })
        (div { #class "title" } $author.full-name)
        (div { #class "description" } $author.bio)  
    ```  

- [ ] **Lexical Scoping:** Lexical scoping is not yet implemented in the current version.

