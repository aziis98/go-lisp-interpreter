# Go Lisp Interpreter

A small lisp-like language interpreted in Golang 

```lisp
(fmt.Println "Hello, World!")
```

## Features

This language has the following features (those not checked are work in progress)

- [ ] Very easy interoperability from and with Golang

    To call Go code from this lisp you can use `interp.Use(lisp.Symbols.All)` or `interp.Use(lisp.Symbols.Fmt)` that just adds those symbols to the root scope, those variables are just collections of stdlib functions exported as a map.

    To mix lisp and Go code instead you can use `interp.EvalWithMap("(+ x 1)", lisp.ScopeMap{ "x": 1 })` that evaluates to `int64(2)`. Alternatively the method `Interpreter.EvalWithValues(string, any...)` just takes a series of values and binds them to `_0`, `_1` and so on. By the way these bindings don't override globals with the same name because new scope gets created.

- [x] Lists with `[1 2 3]`

- [x] Maps with `{ #name "John" #surname "Smith" }`

- [x] Property chaining with `.`

    For example `(math.Sqrt (+ (^ v.x 2) (^ v.y 2) (^ v.z 2)))` to compute vector length

- [x] Quasi-quotes forms with `#` for quoting and `$` for interpolation, for example `#(a b 1.0 [3 $(+ 3 1) 5] $(+ 1 2 3))` evaluates to `#(a b 1 [3 4 5] 6)`.

    For example an html DSL could look like this

    ```lisp
    #(div { #class "card" }
        (img { #src $author.profile-image-url #alt "example image" })
        (div { #class "title" } $author.full-name)
        (div { #class "description" } $author.bio)  
    ```  

- [ ] Lexical scoping (now there is no scoping yet)