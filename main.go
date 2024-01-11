package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToDictionary(t *testing.T) {
	dict := NewDictionary()

	err := dict.Add("key1", "value1")
	assert.Nil(t, err)

	err = dict.Add("key1", "value2")
	assert.NotNil(t, err)
}

func TestGetFromDictionary(t *testing.T) {
	dict := NewDictionary()
	dict.Add("key1", "value1")

	value, err := dict.Get("key1")
	assert.Nil(t, err)
	assert.Equal(t, "value1", value)

	_, err = dict.Get("key2")
	assert.NotNil(t, err)
}

func TestRemoveFromDictionary(t *testing.T) {
	dict := NewDictionary()
	dict.Add("key1", "value1")

	err := dict.Remove("key1")
	assert.Nil(t, err)

	err = dict.Remove("key2")
	assert.NotNil(t, err)
}

func TestListDictionary(t *testing.T) {
	dict := NewDictionary()
	dict.Add("key1", "value1")
	dict.Add("key2", "value2")

	items := dict.List()
	assert.Equal(t, 2, len(items))
	assert.Contains(t, items, DictionaryItem{Key: "key1", Value: "value1"})
	assert.Contains(t, items, DictionaryItem{Key: "key2", Value: "value2"})
}
