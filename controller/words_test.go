package controller

import (
	"testing"

	"github.com/shoriwe/rollee-test-2/common/sqlite"
	"github.com/shoriwe/rollee-test-2/models"
	"github.com/stretchr/testify/assert"
)

func TestAddWord(t *testing.T) {
	t.Run("Valid", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Equal(tt, "hello", word.Word)
		assert.Equal(tt, 1, word.Repeated)
	})
	t.Run("Repeated", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Equal(tt, "hello", word.Word)
		assert.Equal(tt, 5, word.Repeated)
	})
	t.Run("MostCommon", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "helloRelloo"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "helloRelloo"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "helloRelloo"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Equal(tt, "helloRelloo", word.Word)
		assert.Equal(tt, 3, word.Repeated)
	})
	t.Run("InvalidWord", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.NotNil(tt, c.AddWord(&models.Word{Word: "hello Wrong"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Nil(tt, word)
	})
}

// TestQueryWord is mostly the same as AddWord because I use Query in both function to check insertions
func TestQueryWord(t *testing.T) {
	t.Run("Repeated", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Equal(tt, "hello", word.Word)
		assert.Equal(tt, 5, word.Repeated)
	})
	t.Run("MostCommon", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "hello"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "helloRelloo"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "helloRelloo"}))
		assert.Nil(tt, c.AddWord(&models.Word{Word: "helloRelloo"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Equal(tt, "helloRelloo", word.Word)
		assert.Equal(tt, 3, word.Repeated)
	})
	t.Run("InvalidWord", func(tt *testing.T) {
		c := New(sqlite.NewTest())
		defer c.Close()
		assert.NotNil(tt, c.AddWord(&models.Word{Word: "hello Wrong"}))
		word, err := c.QueryWord("he")
		assert.Nil(tt, err)
		assert.Nil(tt, word)
	})
}
