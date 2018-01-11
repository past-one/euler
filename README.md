# euler
euler tour trees in golang :+1:

## algorithm description
- [EN](https://en.wikipedia.org/wiki/Euler_tour_technique#Euler_tour_trees "en.wikipedia.org")
- [RU](https://neerc.ifmo.ru/wiki/index.php?title=%D0%94%D0%B5%D1%80%D0%B5%D0%B2%D1%8C%D1%8F_%D0%AD%D0%B9%D0%BB%D0%B5%D1%80%D0%BE%D0%B2%D0%B0_%D0%BE%D0%B1%D1%85%D0%BE%D0%B4%D0%B0 "neerc.ifmo.ru")

## usage

```golang


trees := CreateEuler()

fmt.Println(trees.IsConnected(1, 2)) // false
fmt.Println(trees.Link(1, 2)) // true
fmt.Println(trees.Link(2, 3)) // true
fmt.Println(trees.IsConnected(1, 2)) // true
fmt.Println(trees.IsConnected(1, 3)) // true

fmt.Println(trees) // 1-2-3-2-1 - euler tour

fmt.Println(trees.Link(1, 2)) // false - already connected
fmt.Println(trees.Cut(1, 2)) // true
fmt.Println(trees.IsConnected(1, 2)) // false
fmt.Println(trees.IsConnected(1, 3)) // false
fmt.Println(trees.IsConnected(2, 3)) // true
fmt.Println(trees.Cut(1, 2)) // false - no edge
fmt.Println(trees.Cut(1, 3)) // false - no edge

fmt.Println(trees)
// 1
// 2-3-2

// order of link params affect euler tour
fmt.Println(trees.Link(2, 1)) // true

fmt.Println(trees) // 2-3-2-1-2
```

## tests
`go test`

## benchmarks
`go test -bench=.`


