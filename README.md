![HTGO-TTS](https://banners.beyondco.de/HTGO-TTS.png?theme=light&packageManager=&packageName=go+get+%22github.com%2Fhegedustibor%2Fhtgo-tts%22&pattern=bamboo&style=style_1&description=Text+to+Speech+Package+for+GoLang&md=1&showWatermark=0&fontSize=100px&images=volume-up)

# htgo-tts
[https://hegedustibor.github.io/htgo-tts/](https://hegedustibor.github.io/htgo-tts/)


### Install
```
go get "github.com/hegedustibor/htgo-tts"
```

### Update
```
go get -u "github.com/col3name/htgo-tts"
```

### Remove
```
go clean -i "github.com/col3name/htgo-tts"
```

### Add to ```go.mod``` file
```replace github.com/hegedustibor/htgo-tts latest => github.com/col3name/htgo-tts latest```

### Import
```go
import "github.com/col3name/htgo-tts"
import "github.com/col3name/htgo-tts/voices"
```

### Use
```go
speech := htgotts.Speech{Folder: "audio", Language: voices.English, Volume: 0, Speed: 1}
speech.Speak("Your sentence.")
```

### Use with Handlers
```go
import (
    htgotts "github.com/col3name/htgo-tts"
    handlers "github.com/col3name/htgo-tts/handlers"
    voices "github.com/col3name/htgo-tts/voices"
)

speech := htgotts.Speech{Folder: "audio", Language: voices.English, Volume: 0, Speed: 1}
speech.Speak("Your sentence.")
```

Have Fun!
