# sterr
Библиотека, позволяющая логировать ошибки с стектрейсом всего пути следования ошибки.

[enEN](README_EN.md)

## Установка
`go get -u github.com/Lechefer/sterr@latest`

## Обзор
### Пример
Предположим что у вас есть такая структура приложения
```
┌─ main.go
├─ handlers/ 
|  └─ handler.go
├─ services/ 
|  └─ service.go
└─ infrastructure/ 
   └─ somerepo/
      └─ repo.go
```

И такой код

#### handler.go
```go
1   func Handler(service Service) {
2       if err := service.Usecase(); err != nil {
3           log.Println(sterr.Wrapf(err, "usecase failed").Error())
4       }
5   }
```
#### service.go
```go
1   type Service struct {
2       repo Repo
3   }
4   
5   func (s Service) UseCase() error {
6       err := repo.SomeAction()
7       return sterr.Wrap(err)
8   }
```
#### repo.go
```go
1   type Repo struct {}
2   
3   func (s Service) SomeAction() error {
4       return sterr.New("failed some action")
5   }
```

В этом случае в лог упадёт следующее сообщение

```handlers.Handler:3 [usecase failed] -> services.Service.UseCase:7 -> somerepo.Repo.SomeAction:4 [failed some action]```

### Ещё
- Получение стектрейса происходит каждый раз в момент создания или оборачивания ошибки
- Оборачивать можно любые ошибки реализующие ```interface error```. Однако стектрейс будет распространяться только на те ошибки, что были созданы ```sterr```
- ```New``` и ```Wrapf``` поддерживают форматирование

### Производительность
|  ns/op | B/op | alloc/op |
|--------|------|----------|
|  5487  | 1249 | 21       |
