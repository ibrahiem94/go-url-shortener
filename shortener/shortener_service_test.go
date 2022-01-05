package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {

	//given
	firstInitialLink := "https://www.link.com/test-link/potatos"
	SecondInitialLink := "https://www.link.com/test-link/carrots"
	ThirdInitialLink := "https://www.link.com/test-link/tomatos"
	//when
	firstShortLink := GenerateShortLink(firstInitialLink, UserId)
	SecondShortLink := GenerateShortLink(SecondInitialLink, UserId)
	ThirdShortLink := GenerateShortLink(ThirdInitialLink, UserId)
	//then
	assert.Equal(t, firstShortLink, "xn5k8Wr2")
	assert.Equal(t, SecondShortLink, "wUruSum5")
	assert.Equal(t, ThirdShortLink, "kPoTX7GE")
}
