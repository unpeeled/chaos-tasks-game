package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"net/http"
)

type Texts map[string]string

var texts = map[string]Texts{
	"German":Texts{
		"y_name": "Dein Name:",
		"button": "Los gehts!",
		"descr": "Herzlich Willkommen auf der Party. Hol dir die Aufgabe, die das Universum für dich bereit hält.",	},
	"English":Texts{
		"y_name": "Your Name:",
		"button":"Let's go!",
		"descr": "Welcome to the party. Take the task that you have to fulfill tonight.",
	},
}

var matcher = language.NewMatcher([]language.Tag{
	language.English, //fallback
	language.German})

func lang_trans(r *http.Request, entry string) string {
	lang, _ := r.Cookie("lang")
	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(matcher, lang.String(), accept)

	return texts[display.English.Tags().Name(tag)][entry]
}

