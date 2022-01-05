package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = Init()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {

	// Given
	initialLink := "https://www.link.com/test-link/potatos"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// When
	SaveUrl(shortURL, initialLink, userUUId)
	retrievedUrl := RetrieveUrl(shortURL)

	// Then
	assert.Equal(t, initialLink, retrievedUrl)
}
