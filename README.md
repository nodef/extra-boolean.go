[Boolean] data type has two possible truth values to represent logic.<br>
:package: [GoDev](https://pkg.go.dev/github.com/golangf/extra-boolean),
:newspaper: [GoDoc](https://pkg.go.dev/github.com/golangf/extra-boolean#section-documentation),
:blue_book: [Wiki](https://github.com/golangf/extra-boolean/wiki).

Here is my implementation of digital logic gates in software. That includes
the basic gates [Not], [And], [Or], [Xor]; their complements [Nand], [Nor],
[Xnor]; and 2 propositional logic (taught in discrete mathematics) gates
[Imply], [Eq]; and their complements [Nimply], [Neq]. There is also a
multiplexer, called [Select], and a `true` counter, called [Count]. [Count]
can help you make custom gates, such as an *alternate* concept of **xnor**
which returns `true` only if all inputs are the same (standard [Xnor] returns
`true` if even inputs are `true`). All of them can handle upto 8 inputs.

[Parse] is influenced by ["boolean"] package, and is quite good at translating
`string` to `boolean`. It can also handle double negatives, eg. `not inactive`.
You know the [And] of 2-inputs, but what of 1-input? What of 0? And what of
the other gates? I answer them here.

> Stability: Experimental.

<br>

```javascript
import (
  boolean "github.com/golangf/extra-boolean"
)

boolean.Parse("1")
boolean.Parse("truthy")
boolean.Parse("not off")
// true

boolean.Parse("not true")
boolean.Parse("inactive")
boolean.Parse("disabled")
// false

boolean.Imply(true, false)
// false

boolean.Eq(false, false)
// true

boolean.Xor3(true, true, true)
// true

boolean.Select3(1, true, false, true)
// false                   ^

boolean.Count3(true, false, true)
// 2            ^            ^
```

<br>
<br>


## Index

| Name     | Action                                     |
| -------- | ------------------------------------------ |
| [Parse]  | Converts string to boolean.                |
| [Not]    | Checks if value is false.                  |
| [And]    | Checks if all values are true.             |
| [Or]     | Checks if any value is true.               |
| [Xor]    | Checks if odd no. of values are true.      |
| [Nand]   | Checks if any value is false.              |
| [Nor]    | Checks if all values are false.            |
| [Xnor]   | Checks if even no. of values are true.     |
| [Eq]     | Checks if antecedent ⇔ consequent (a ⇔ b). |
| [Neq]    | Checks if antecedent ⇎ consequent (a ⇎ b). |
| [Imply]  | Checks if antecedent ⇒ consequent (a ⇒ b). |
| [Nimply] | Checks if antecedent ⇏ consequent (a ⇏ b). |
| [Select] | Checks if ith value is true.               |
| [Count]  | Counts no. of true values.                 |

<br>
<br>


## References

- [How do I do a case insensitive regular expression in Go?](https://stackoverflow.com/a/15326471/1413259)
- [Convert string to integer type in Go?](https://stackoverflow.com/a/29841190/1413259)
- [Set a global variable only once in golang](https://stackoverflow.com/a/38657214/1413259)
- [GoDoc add newline character](https://stackoverflow.com/q/51641640/1413259)
- [Indentation is replaced with tabs from spaces on save ...](https://github.com/microsoft/vscode-go/issues/1930)
- [Testable Examples in Go](https://blog.golang.org/examples)
- [Godoc: documenting Go code](https://blog.golang.org/godoc)
- [Optional Parameters in Go?](https://stackoverflow.com/q/2032149/1413259)
- [Publishing Go Modules](https://blog.golang.org/publishing-go-modules)
- [Developing and publishing modules](https://golang.org/doc/modules/developing)
- [Using Go Modules](https://blog.golang.org/using-go-modules)
- [Go naming conventions for const](https://stackoverflow.com/a/37216200/1413259)
- [What are conventions for filenames in Go?](https://stackoverflow.com/a/25162021/1413259)
- [Go Project Structure Best Practices](https://tutorialedge.net/golang/go-project-structure-best-practices/)
- [Naming Conventions in Go: Short but Descriptive](https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a)
- [Effective Go](https://golang.org/doc/effective_go)
- [How to Write Go Code](https://golang.org/doc/code)
- [How to fix “go get: warning: modules disabled by GO111MODULE=auto in GOPATH/src”](https://stackoverflow.com/a/59684525/1413259)
- [Batch Rename](https://marketplace.visualstudio.com/items?itemName=JannisX11.batch-rename-extension)

<br>
<br>

[![](https://img.youtube.com/vi/6mMK6iSZsAs/maxresdefault.jpg)](https://www.youtube.com/watch?v=6mMK6iSZsAs)
![](https://ga-beacon.deno.dev/G-RC63DPBH3P:SH3Eq-NoQ9mwgYeHWxu7cw/github.com/nodef/extra-boolean.go)

[Boolean]: https://pkg.go.dev/builtin#bool
["boolean"]: https://www.npmjs.com/package/boolean
[Parse]: https://github.com/golangf/extra-boolean/wiki/Parse
[Xor]: https://github.com/golangf/extra-boolean/wiki/Xor
[Not]: https://github.com/golangf/extra-boolean/wiki/Not
[And]: https://github.com/golangf/extra-boolean/wiki/And
[Or]: https://github.com/golangf/extra-boolean/wiki/Or
[Nand]: https://github.com/golangf/extra-boolean/wiki/Nand
[Nor]: https://github.com/golangf/extra-boolean/wiki/Nor
[Xnor]: https://github.com/golangf/extra-boolean/wiki/Xnor
[Eq]: https://github.com/golangf/extra-boolean/wiki/Eq
[Imply]: https://github.com/golangf/extra-boolean/wiki/Imply
[Nimply]: https://github.com/golangf/extra-boolean/wiki/Nimply
[Select]: https://github.com/golangf/extra-boolean/wiki/Select
[Count]: https://github.com/golangf/extra-boolean/wiki/Count
[Neq]: https://github.com/golangf/extra-boolean/wiki/Neq
