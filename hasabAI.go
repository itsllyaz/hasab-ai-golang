package hasabai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"gopkg.in/resty.v1"
)


type HasabAI interface{
	HasabHistory() (*HasabResponse, error)
	HasabTranscriptionHistory() (*HasabTranscriptionResponse, error)
	HasabTranslationsHistory() (*HasabTranslationsResponse, error)
	HasabSpeakers() (*HasabSpeakersResponse, error)
	TTSHistory() (*TTSHistoryResponse, error)
	TTSAnalytics() (*TTSAnalyticsResponse, error)
	TTSRecord() (TTSRecordResponse, error)
	
}

type Client struct {
    ApiKey     string

}

func New(apiKey string) *Client {
    
    return &Client{
        ApiKey:apiKey,
    }
}
func (c *Client) CallHasab(url string, target any) error {
    resp, err := resty.New().R().
        SetHeader("Authorization", "Bearer "+c.ApiKey).
        SetHeader("Content-Type", "application/json").
        SetHeader("Accept", "application/json").
        Get(url)

    if err != nil {
        return fmt.Errorf("request error: %w", err)
    }

    if resp.StatusCode() >= 300 {
       return fmt.Errorf("bad status: %d, body: %s", resp.StatusCode(), resp.Body()) 
    }

    if err := json.Unmarshal(resp.Body(), target); err != nil {
        return fmt.Errorf("unmarshal error: %w", err)
    }

    return nil
}

func (c *Client) HasabHistory() (*HasabResponse, error) {
    var result HasabResponse
    if err := c.CallHasab(HasabHistoryV1URL, &result); err != nil {
        return nil, err
    }
    return &result, nil
}


func (c *Client) HasabTranscriptionHistory() (*HasabTranscriptionResponse, error) {
    var result HasabTranscriptionResponse
    if err := c.CallHasab(HasabTranscriptionHistoryV1URL, &result); err != nil {
        return nil, err
    }
    return &result, nil
}



func (c *Client) HasabTranslationsHistory() (*HasabTranslationsResponse, error) {
    var result HasabTranslationsResponse
    if err := c.CallHasab(HasabTranslationsHistoryV1URL, &result); err != nil {
        return nil, err
    }
    return &result, nil
}

func (c *Client) HasabSpeakers() (*HasabSpeakersResponse, error){
	var result HasabSpeakersResponse 
	if err := c.CallHasab(HasabSpeakersV1URL, &result); err != nil{
		return nil , err 
	}

	return &result, nil
}

func (c *Client) TTSHistory() (*TTSHistoryResponse, error){
	var result TTSHistoryResponse
	if err := c.CallHasab(TTSHistoryV1URL, &result); err != nil{
		return nil, err
	}
	return &result, nil
}

func (c *Client) TTSAnalytics() (*TTSAnalyticsResponse, error){
	var result TTSAnalyticsResponse 
	if err := c.CallHasab(TTSAnalyticsV1URL, &result); err != nil{
		return nil , err 
	}
	return &result, nil 
}


func (c *Client) TTSRecord(recordID int) (*TTSRecordResponse, error){
	var result TTSRecordResponse 
	if err := c.CallHasab(fmt.Sprintf("%s/%d", TTSRecordV1URL, recordID), &result); err != nil {
		return nil , err 
	}
			
	return &result, nil
}


func (c *Client) uploadAudio(filePath string, transcribe, translate bool, targetLanguage, sourceLanguage string) (*UploadAudioResponse, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("cannot open file: %w", err)
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    part, err := writer.CreateFormFile("file", filepath.Base(filePath))
    if err != nil {
        return nil, fmt.Errorf("create form file error: %w", err)
    }
    if _, err := io.Copy(part, file); err != nil {
        return nil, fmt.Errorf("copy file error: %w", err)
    }

    _ = writer.WriteField("transcribe", fmt.Sprintf("%t", transcribe))
    _ = writer.WriteField("translate", fmt.Sprintf("%t", translate))
    _ = writer.WriteField("summarize", "false")
    _ = writer.WriteField("source_language", sourceLanguage)
    _ = writer.WriteField("language", targetLanguage)
    _ = writer.WriteField("timestamps", "false")
    _ = writer.WriteField("is_meeting", "false")
    writer.Close()

    client := resty.New()
    resp, err := client.R().
        SetHeader("Authorization", "Bearer "+c.ApiKey).
        SetHeader("Content-Type", writer.FormDataContentType()).
        SetHeader("Accept", "application/json").
        SetBody(body).
        Post("https://hasab.co/api/v1/upload-audio")

    if err != nil {
        return nil, fmt.Errorf("request error: %w", err)
    }

    if resp.StatusCode() >= 300 {
        return nil, fmt.Errorf("bad status: %d, body: %s", resp.StatusCode(), resp.Body())
    }

    var result UploadAudioResponse
    if err := json.Unmarshal(resp.Body(), &result); err != nil {
        return nil, fmt.Errorf("unmarshal error: %w", err)
    }

    return &result, nil
}


// helper for transcription only
func (c *Client) TranscribeAudio(filePath, sourceLanguage string) (*UploadAudioResponse, error) {
    return c.uploadAudio(filePath, true, false, "", sourceLanguage)
}

// helper for transcription + translation
func (c *Client) TranslateAudio(filePath, sourceLanguage, targetLanguage string) (*UploadAudioResponse, error) {
    return c.uploadAudio(filePath, true, true, targetLanguage, sourceLanguage)
}
