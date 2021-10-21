# knight 
## <img height="55" src="https://images.prismic.io/lichess/5cfd2630-2a8f-4fa9-8f78-04c2d9f0e5fe_lichess-box-1024.png?auto=compress,format" width="55"/> **Golang Library for the [Lichess Api](https://lichess.org/api).** 
### Usage ###


```go
    api := knight.NewApi("your_token")
    
    eventStream, err := api.StreamIncomingEvents() // get <-chan *Event
    if err != nil {
        // do something
    }
    
    for ev := range eventStream {
        if ev.Type == "gameStart" {
            // do something
        }       
    }

```
or 
```go
    api := knight.NewApi("your_token")

    tvStream, err := a.StreamCurrentTVGame() // get <-chan *TVStream
    if err != nil {
        fmt.Println(err)
    }
    
    for game := range tvStream { 
        // stream current tv game
    }
```

###API is available:

```go
    // account api
    api.GetProfile
    api.GetEmail
    api.GetPreference
    api.GetKidMode
    api.SetKidMode

   // users api 	
    api.GetUsersStatus
    api.GetUser
    api.GetUserHistory
    api.GetPerformance
    api.GetUserActivity
    api.GetUsersById
    api.GetMembersTeam
    api.GetLiveStreamer
    api.GetTop10
    api.GetOneLeadBoard


    //relations api
    api.GetFollowing
    api.GetFollowers
    api.FollowPlayer
    api.UnfollowPlayer

    // game api
    api.ExportGameJson
    api.ExportGamePgn
    api.ExportOngoingGameJson
    api.ExportOngoingGamePgn

    // puzzle api
    api.GetDailyPuzzle
    api.GetPuzzleActivity
    api.GetPuzzleDashboard
    api.GetStormDashboard


    // tv api
    api.GetCurrentTVGames
    api.StreamCurrentTVGame
    api.GetBestTVOngoingGames
    api.GetBestTVOngoingGamesPGN

    // board api
    api.StreamIncomingEvents

```