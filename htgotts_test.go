package htgotts

import (
	"github.com/col3name/htgo-tts/voices"
	"log"

	"testing"
)

func TestSpeech_Speak(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.English, Volume: 0, Speed: 1}
	err := speech.Speak("Test")
	if err != nil {
		log.Fatal(err)
	}
}

func TestSpeech_Speak_voice_UkEnglish(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.EnglishUK, Volume: 0, Speed: 1}
	err := speech.Speak("Lancaster")
	if err != nil {
		log.Fatal(err)
	}
}

func TestSpeech_Speak_voice_Japanese(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.Japanese, Volume: 0, Speed: 1}
	err := speech.Speak("Test")
	if err != nil {
		log.Fatal(err)
	}
}

func TestSpeech_CreateSpeechFile(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.English, Volume: 0, Speed: 1}
	_, err := speech.CreateSpeechFile("Test", "testfilename")
	if err != nil {
		t.Fatalf("CreateSpeechFile fail %v", err)
	}
}

func TestSpeech_(t *testing.T) {
	speech := Speech{Folder: "audio", Language: voices.English, Volume: 0, Speed: 1}
	f, err := speech.CreateSpeechFile("Test", "testplay")
	if err != nil {
		t.Fatalf("CreateSpeechFile fail %v", err)
	}
	err = speech.PlaySpeechFile(f)
	if err != nil {
		log.Fatal(err)
	}
}
