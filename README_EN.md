# sterr
A convenient library that allows you to log errors with a stack trace of the entire error path.

## Quick Start
`go get -u github.com/Lechefer/sterr@latest`

## Overview

### Example
Suppose you have such an application structure
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

And there is such a code

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

In this case, the following message will fall into the log

```handlers.Handler:3 [usecase failed] -> services.Service.UseCase:7 -> somerepo.Repo.SomeAction:4 [failed some action]```

### More
- A stack trace is received every time an error is created or wrapped
- You can wrap any errors implementing ```interface error```. However, the stack trace will only apply to those errors that were created by ```sterr```
- ```New``` and ```Wrapf``` support formatting

### Bechmarks
|  ns/op | B/op | alloc/op |
|--------|------|----------|
|  5487  | 1249 | 21       |
