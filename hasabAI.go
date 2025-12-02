package hasabai

import (
	"encoding/json"
	"fmt"

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
