# Go Lisp Interpreter

A small lisp-like language interpreted in Golang 

```lisp
(fmt.Println "Hello, World!")
```

## Features

This language has the following features (those not checked are work in progress)

- [ ] Very easy interoperability from and with Golang

- [ ] Lists with `[1 2 3]`

- [ ] Maps with `{ #name "John" #surname "Smith" }`

- [x] Property chaining with `.`

    For example `(math.Sqrt (+ (^ v.x 2) (^ v.y 2) (^ v.z 2)))` to compute vector length

- [ ] Quasi-quotes forms with `#` for quoting and `$` for interpolation

    For example an html DSL could look like this

    ```lisp
    #(div { #class "card" }
        (img { #src $author.profile-image-url #alt "example image" })
        (div { #class "title" } $author.full-name)
        (div { #class "description" } $author.bio)  
    ```  
