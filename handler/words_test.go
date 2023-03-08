package handler

import (
	"net/http"
	"testing"

	"github.com/shoriwe/rollee-test-2/models"
)

func TestAddWord(t *testing.T) {
	t.Run("Repeated", func(tt *testing.T) {
		expect, terminate := NewTest(tt)
		defer terminate()
		for i := 0; i < 10; i++ {
			expect.POST(AddWordRoute).
				WithJSON(models.Word{Word: "Hello"}).
				Expect().
				Status(http.StatusOK)
		}
		obj := expect.GET(QueryWordRoute + "/H").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()
		obj.Value("word").
			String().
			IsEqual("Hello")
		obj.Value("repeated").
			Number().
			IsEqual(10)
	})
	t.Run("MostCommon", func(tt *testing.T) {
		expect, terminate := NewTest(tt)
		defer terminate()
		for i := 0; i < 10; i++ {
			expect.POST(AddWordRoute).
				WithJSON(models.Word{Word: "Hello"}).
				Expect().
				Status(http.StatusOK)
		}
		for i := 0; i < 100; i++ {
			expect.POST(AddWordRoute).
				WithJSON(models.Word{Word: "HelloSulcud"}).
				Expect().
				Status(http.StatusOK)
		}
		obj := expect.GET(QueryWordRoute + "/H").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()
		obj.Value("word").
			String().
			IsEqual("HelloSulcud")
		obj.Value("repeated").
			Number().
			IsEqual(100)
	})
	t.Run("InvalidWord", func(tt *testing.T) {
		expect, terminate := NewTest(tt)
		defer terminate()
		expect.POST(AddWordRoute).
			WithJSON(models.Word{Word: "Hello12"}).
			Expect().
			Status(http.StatusInternalServerError)
		expect.GET(QueryWordRoute + "/H").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Value("isNull").
			Boolean().
			IsEqual(true)
	})
}

func TestQueryWord(t *testing.T) {
	t.Run("Repeated", func(tt *testing.T) {
		expect, terminate := NewTest(tt)
		defer terminate()
		for i := 0; i < 10; i++ {
			expect.POST(AddWordRoute).
				WithJSON(models.Word{Word: "Hello"}).
				Expect().
				Status(http.StatusOK)
		}
		obj := expect.GET(QueryWordRoute + "/H").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()
		obj.Value("word").
			String().
			IsEqual("Hello")
		obj.Value("repeated").
			Number().
			IsEqual(10)
	})
	t.Run("MostCommon", func(tt *testing.T) {
		expect, terminate := NewTest(tt)
		defer terminate()
		for i := 0; i < 10; i++ {
			expect.POST(AddWordRoute).
				WithJSON(models.Word{Word: "Hello"}).
				Expect().
				Status(http.StatusOK)
		}
		for i := 0; i < 100; i++ {
			expect.POST(AddWordRoute).
				WithJSON(models.Word{Word: "HelloSulcud"}).
				Expect().
				Status(http.StatusOK)
		}
		obj := expect.GET(QueryWordRoute + "/H").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()
		obj.Value("word").
			String().
			IsEqual("HelloSulcud")
		obj.Value("repeated").
			Number().
			IsEqual(100)
	})
	t.Run("InvalidWord", func(tt *testing.T) {
		expect, terminate := NewTest(tt)
		defer terminate()
		expect.POST(AddWordRoute).
			WithJSON(models.Word{Word: "Hello12"}).
			Expect().
			Status(http.StatusInternalServerError)
		expect.GET(QueryWordRoute + "/H").
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			Value("isNull").
			Boolean().
			IsEqual(true)
	})
}
