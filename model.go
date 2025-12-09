package hasabai

type HasabResponse struct {
    History []ChatHistory `json:"history"`
	Success bool		  `json:"success"`
}

type ChatHistory struct {
    ID       int       `json:"id"`
    Title    string    `json:"title"`
    Messages []Message `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type HasabTranscriptionResponse struct {
    ID               int    `json:"id"`
    Filename         string `json:"filename"`
    OriginalFilename string `json:"original_filename"`
    MimeType         string `json:"mime_type"`
    DurationSeconds  string `json:"duration_in_seconds"`
    Description      string `json:"description"`
	Success			 bool    `json:"success"`
    Transcription    string `json:"transcription"`
    Translation      string `json:"translation"`
    Summary          string `json:"summary"`
    AudioType        string `json:"audio_type"`
    CreatedAt        string `json:"created_at"`
    User             User   `json:"user"`
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}


type HasabTranslationsResponse struct {
    Success         bool   `json:"success"`
    Message         string `json:"message"`
    ChargeMessage   string `json:"charge_message"`
    RemainingTokens int    `json:"remaining_tokens"`
    TokensCharged   int    `json:"tokens_charged"`
    Summary         string `json:"summary"`
    Timestamp       []any  `json:"timestamp"`

    Audio struct {
        ID               int    `json:"id"`
        UserID           int    `json:"user_id"`
        Filename         string `json:"filename"`
        OriginalFilename string `json:"original_filename"`
        Path             string `json:"path"`
        MimeType         string `json:"mime_type"`
        FileSize         string `json:"file_size"`
        DurationSeconds  int    `json:"duration_in_seconds"`
        Description      string `json:"description"`
        IsMeeting        bool   `json:"is_meeting"`
        Summary          string `json:"summary"`
        Transcription    string `json:"transcription"`
        Translation      string `json:"translation"`
        CreatedAt        string `json:"created_at"`
        UpdatedAt        string `json:"updated_at"`
    } `json:"audio"`

    Metadata struct {
        TokensCharged    int    `json:"tokens_charged"`
        RemainingTokens  int    `json:"remaining_tokens"`
        ChargeMessage    string `json:"charge_message"`
    } `json:"metadata"`
}




type HasabSpeakersResponse struct {
    Languages struct {
        Amh []string `json:"amh"`
        Orm []string `json:"orm"`
        Tir []string `json:"tir"`
    } `json:"languages"`
    Success       bool `json:"success"`
    TotalSpeakers int  `json:"total_speakers"`
}



type TTSHistoryResponse struct {
    Records []struct {
        ID          int    `json:"id"`
        Text        string `json:"text"`
        Language    string `json:"language"`
        SpeakerName string `json:"speaker_name"`
        Status      string `json:"status"`
        AudioURL    string `json:"audio_url"`
        TokensUsed  int    `json:"tokens_used"`
        CreatedAt   string `json:"created_at"`
    } `json:"records"`
    Total     int `json:"total"`
    Limit     int `json:"limit"`
    Success   bool`json:"success"`
    Offset    int `json:"offset"`
}


type TTSAnalyticsResponse struct {
    TotalRequests           int     `json:"total_requests"`
    SuccessfulRequests      int     `json:"successful_requests"`
    Success                 bool    `json:"success"`
    FailedRequests          int     `json:"failed_requests"`
    TotalTokensUsed         int     `json:"total_tokens_used"`
    AverageTokensPerRequest float64 `json:"average_tokens_per_request"`
    LanguageBreakdown       map[string]int `json:"language_breakdown"`
    DailyUsage              []struct {
        Date       string `json:"date"`
        Requests   int    `json:"requests"`
        TokensUsed int    `json:"tokens_used"`
    } `json:"daily_usage"`
}


type TTSRecordResponse struct {
    Success bool  `json:"success"`
    Record struct {
        ID          int    `json:"id"`
        Text        string `json:"text"`
        Language    string `json:"language"`
        SpeakerName string `json:"speaker_name"`
        Status      string `json:"status"`
        AudioURL    string `json:"audio_url"`
        TokensUsed  int    `json:"tokens_used"`
        CreatedAt   string `json:"created_at"`
    } `json:"record"`
}


type UploadAudioResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Audio   struct {
        Transcription string `json:"transcription"`
        Translation   string `json:"translation"`
    } `json:"audio"`
}
