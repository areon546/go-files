# How Do I want these libraries to be designed?

1. Ease of Use:
Bad:
```go
file := fileIO.NewTextFile(fileIO.ConstructFilePath("documents","files","text.txt"))
```
The above example has too many calls that would rely on the same package repeatedly, not fun. 

Good:
```go
file := fileIO.NewTextFile("documents/files/text.txt")
``
The above example is intuitive and lets users conveniently use the NewTextFile method. 