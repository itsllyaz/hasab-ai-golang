## Hasab AI Golang SDK 
Unofficial Golang SDK for Hasab AI API (Audio transcription, translation, TTS, and more)

### How to Use It
##### 1. Installation
```
    go get github.com/itsllyaz/hasab-ai-golang
```

##### 2. Setup

```go
    package main

    import (
        "os"
        "github.com/itsllyaz/hasab-ai-golang"
    )

    func main(){
        client := hasabai.New(os.Getenv("HASAB_API_KEY"))
        fmt.Println("Hasab client initialized:", client)
    }
```

##### 3. Upload Audio & Transcription
```go
    resp, err := client.TranscribeAudio("audio.mp3", "eng")
    if err != nil {
        fmt.Println("Transcription error:", err)
        return
    }

    fmt.Printf("Transcription result: %+v\n", resp)

```

##### 4. Upload Audio & Translation
```go
    resp, err := client.TranslateAudio("audio.mp3", "eng", "amh")
    if err != nil {
        fmt.Println("Translation error:", err)
        return
    }

    fmt.Printf("Translation result: %+v\n", resp)


```

##### 5. Fetch Hasab History
```go
    history, err := client.HasabHistory()
    if err != nil {
        fmt.Println("Error fetching history:", err)
        return
    }
    fmt.Printf("History: %+v\n", history)


```

##### 6. TTS & Analytics
```go
    ttsHistory, err := client.TTSHistory()
    if err != nil {
        fmt.Println("TTS history error:", err)
        return
    }
    fmt.Printf("TTS history: %+v\n", ttsHistory)

```


### Resources
- [Hasab API Docs](https://developer.hasab.ai/)




### Contributions
- wat r u waiting....
