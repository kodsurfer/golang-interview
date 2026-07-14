## Разбор сниппетов кода («что выведет этот код?»)

**Пример** с defer и слайсами:

```go
func main() {
    s := []int{1, 2, 3}
    defer func() { fmt.Println(s) }()
    s = append(s, 4)
}
```
