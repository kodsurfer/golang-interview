## Задача 4. Агрегация и нормализация событий (почти «микросервисная» задача)
**Условие**. Есть поток событий в виде слайса структур:
```go
type Event struct {
 ID        string
 UserID    string
 Type      string
 Timestamp time.Time
 Payload   map[string]any
}
```
