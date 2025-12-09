package hasab_test 

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/itsllyaz/hasab-ai-golang"
)

type dummyResp struct {
	Message string `json:"message"`
}

func setupMockServer(t *testing.T, expectedPath string, response interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header")
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	
	return server
}

func TestCallHasab(t *testing.T) {
	expected := dummyResp{Message: "ok"}
	server := setupMockServer(t, "/history", expected)
	defer server.Close()

	client := hasabai.New("test-api-key")
	var result dummyResp
	err := client.CallHasab(server.URL+"/history", &result)
	if err != nil {
		t.Fatalf("CallHasab error: %v", err)
	}
	if result.Message != "ok" {
		t.Errorf("Unexpected result: %+v", result)
	}
}

func TestHasabHistory(t *testing.T) {
	expected := hasabai.HasabResponse{Success: true}
	server := setupMockServer(t, "/history", expected)
	defer server.Close()

	

	client := hasabai.New("test-api-key")
	res, err := client.HasabHistory()
	if err != nil {
		t.Fatalf("HasabHistory failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Unexpected result: %+v", res)
	}
}

func TestHasabTranscriptionHistory(t *testing.T) {
	expected := hasabai.HasabTranscriptionResponse{Success: true}
	server := setupMockServer(t, "/transcriptions", expected)
	defer server.Close()

	client := hasabai.New("test-api-key")
	res, err := client.HasabTranscriptionHistory()
	if err != nil {
		t.Fatalf("HasabTranscriptionHistory failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Unexpected result: %+v", res)
	}
}

func TestHasabTranslationsHistory(t *testing.T) {
	expected := hasabai.HasabTranslationsResponse{Success: true}
	server := setupMockServer(t, "/translations", expected)
	defer server.Close()

	client := hasabai.New("test-api-key")
	res, err := client.HasabTranslationsHistory()
	if err != nil {
		t.Fatalf("HasabTranslationsHistory failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Unexpected result: %+v", res)
	}
}

func TestHasabSpeakers(t *testing.T) {
	expected := hasabai.HasabSpeakersResponse{Success: true}
	server := setupMockServer(t, "/speakers", expected)
	defer server.Close()

	client := hasabai.New("test-api-key")
	res, err := client.HasabSpeakers()
	if err != nil {
		t.Fatalf("HasabSpeakers failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Unexpected result: %+v", res)
	}
}

func TestTTSFunctions(t *testing.T) {
	historyExpected := hasabai.TTSHistoryResponse{Success: true}
	analyticsExpected := hasabai.TTSAnalyticsResponse{Success: true}
	recordExpected := hasabai.TTSRecordResponse{Success: true}

	historyServer := setupMockServer(t, "/tts/history", historyExpected)
	defer historyServer.Close()

	analyticsServer := setupMockServer(t, "/tts/analytics", analyticsExpected)
	defer analyticsServer.Close()

	recordServer := setupMockServer(t, "/tts/record/1", recordExpected)
	defer recordServer.Close()

	client := hasabai.New("test-api-key")

	hRes, err := client.TTSHistory()
	if err != nil || !hRes.Success {
		t.Fatalf("TTSHistory failed: %v", err)
	}

	aRes, err := client.TTSAnalytics()
	if err != nil || !aRes.Success {
		t.Fatalf("TTSAnalytics failed: %v", err)
	}

	rRes, err := client.TTSRecord(1)
	if err != nil || !rRes.Success {
		t.Fatalf("TTSRecord failed: %v", err)
	}
}

func TestUploadAudioFunctions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := hasabai.UploadAudioResponse{
			Success: true,
			Audio: struct {
				Transcription string `json:"transcription"`
				Translation   string `json:"translation"`
			}{
				Transcription: "dummy transcription",
				Translation:   "dummy translation",
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	tmpFile, _ := os.CreateTemp("", "dummy*.mp3")
	tmpFile.Write([]byte("dummy content"))
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	client := hasabai.New("test-api-key")

	res, err := client.TranscribeAudio(tmpFile.Name(), "eng")
	if err != nil {
		t.Fatalf("TranscribeAudioWithURL failed: %v", err)
	}
	if res.Audio.Transcription != "dummy transcription" {
		t.Errorf("Unexpected transcription: %s", res.Audio.Transcription)
	}

	
	res2, err := client.TranslateAudio(tmpFile.Name(), "eng", "amh")
	if err != nil {
		t.Fatalf("TranslateAudioWithURL failed: %v", err)
	}
	if res2.Audio.Translation != "dummy translation" {
		t.Errorf("Unexpected translation: %s", res2.Audio.Translation)
	}
}
