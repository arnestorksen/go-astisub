package astisub_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/arnestorksen/go-astisub"
	"github.com/asticode/go-astikit"
	"github.com/stretchr/testify/assert"
)

func TestSTL(t *testing.T) {
	// Init
	astisub.Now = func() (t time.Time) {
		t, _ = time.Parse("060102", "170702")
		return
	}

	// Open
	s, err := astisub.OpenFile("./testdata/example-in.stl")
	assert.NoError(t, err)
	assertSubtitleItems(t, s)
	// Metadata
	assert.Equal(t, &astisub.Metadata{Framerate: 25, Language: astisub.LanguageFrench, STLMaximumNumberOfDisplayableCharactersInAnyTextRow: astikit.IntPtr(40), STLMaximumNumberOfDisplayableRows: astikit.IntPtr(23), STLPublisher: "Copyright test", Title: "Title test"}, s.Metadata)

	// No subtitles to write
	w := &bytes.Buffer{}
	err = astisub.Subtitles{}.WriteToSTL(w)
	assert.EqualError(t, err, astisub.ErrNoSubtitlesToWrite.Error())

	// Write
	c, err := ioutil.ReadFile("./testdata/example-out.stl")
	assert.NoError(t, err)
	err = s.WriteToSTL(w)
	assert.NoError(t, err)
	assert.Equal(t, string(c), w.String())
}

func TestSTLTV2(t *testing.T) {
	// Init
	astisub.Now = func() (t time.Time) {
		t, _ = time.Parse("060102", "170702")
		return
	}

	// Open
	s, err := astisub.OpenFile("./testdata/D0459622-opn.stl")
	assert.NoError(t, err)

	// Metadata
	assert.Equal(t, 25, s.Metadata.Framerate)
	assert.Equal(t, astikit.IntPtr(37), s.Metadata.STLMaximumNumberOfDisplayableCharactersInAnyTextRow)
	assert.Equal(t, astikit.IntPtr(11), s.Metadata.STLMaximumNumberOfDisplayableRows)
	assert.Equal(t, astisub.LanguageEnglish, s.Metadata.Language)

	// No subtitles to write
	w := &bytes.Buffer{}
	err = s.WriteToTTML(w)
	//assert.EqualError(t, err, astisub.ErrNoSubtitlesToWrite.Error())

	// Write
	err = ioutil.WriteFile("./testdata/D0459622-opn.ttml", w.Bytes(), os.ModePerm)
	//assert.NoError(t, err)
	//	err = s.WriteToSTL(w)
	//	assert.NoError(t, err)
	//	assert.Equal(t, string(c), w.String())
	w = &bytes.Buffer{}
	err = s.WriteToWebVTT(w)
	//assert.EqualError(t, err, astisub.ErrNoSubtitlesToWrite.Error())

	// Write
	err = ioutil.WriteFile("./testdata/D0459622-opn.vtt", w.Bytes(), os.ModePerm)

}
