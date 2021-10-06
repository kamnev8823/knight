# <img height="25" src="https://images.prismic.io/lichess/5cfd2630-2a8f-4fa9-8f78-04c2d9f0e5fe_lichess-box-1024.png?auto=compress,format" width="25"/> knight

###Golang Library for the [Lichess Api](https://lichess.org/api).

##Usage

---

```go
    api := knight.NewApi("your_token")
    
    event, err := api.StreamIncomingEvents()
    if err != nil {
        // do something
    }
    
    for ev := range event.Stream() {
        if ev.Type == "gameStart" {
            // do something
        }       
    }
```