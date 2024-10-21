package entity

import "time"

func fixtureMetadata() Metadata {
	now := time.Now()

	return Metadata{
		Author: &Utf8Text{
			Text: "author",
			UTF8: true,
		},
		Creator: &Utf8Text{
			Text: "creator",
			UTF8: false,
		},
		Subject: &Utf8Text{
			Text: "subject",
			UTF8: true,
		},
		Title: &Utf8Text{
			Text: "title",
			UTF8: true,
		},
		CreationDate: &now,
		KeywordsStr: &Utf8Text{
			Text: "keyword",
			UTF8: true,
		},
	}
}
