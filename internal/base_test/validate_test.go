package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("nil pointer - returns error", func(t *testing.T) {
		t.Parallel()

		rsl := base.Validate(nil)

		expected := []string{"nil pointer is used in ValidateRequest"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("all fields empty - returns all errors", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "",
			Title:       "",
			Description: "",
		}
		rsl := base.Validate(req)

		expected := []string{
			"UserId is required",
			"Title is required",
			"Description is required",
		}
		assert.Equal(t, expected, rsl)
	})

	t.Run("only UserId empty - returns UserId error", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "",
			Title:       "Valid Title",
			Description: "Valid Description",
		}
		rsl := base.Validate(req)

		expected := []string{"UserId is required"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("only Title empty - returns Title error", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "user123",
			Title:       "",
			Description: "Valid Description",
		}
		rsl := base.Validate(req)

		expected := []string{"Title is required"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("only Description empty - returns Description error", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "user123",
			Title:       "Valid Title",
			Description: "",
		}
		rsl := base.Validate(req)

		expected := []string{"Description is required"}
		assert.Equal(t, expected, rsl)
	})

	t.Run("UserId and Title empty - returns two errors", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "",
			Title:       "",
			Description: "Valid Description",
		}
		rsl := base.Validate(req)

		expected := []string{
			"UserId is required",
			"Title is required",
		}
		assert.Equal(t, expected, rsl)
	})

	t.Run("UserId and Description empty - returns two errors", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "",
			Title:       "Valid Title",
			Description: "",
		}
		rsl := base.Validate(req)

		expected := []string{
			"UserId is required",
			"Description is required",
		}
		assert.Equal(t, expected, rsl)
	})

	t.Run("Title and Description empty - returns two errors", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "user123",
			Title:       "",
			Description: "",
		}
		rsl := base.Validate(req)

		expected := []string{
			"Title is required",
			"Description is required",
		}
		assert.Equal(t, expected, rsl)
	})

	t.Run("all fields valid - returns empty slice", func(t *testing.T) {
		t.Parallel()

		req := &base.ValidateRequest{
			UserId:      "user123",
			Title:       "Valid Title",
			Description: "Valid Description",
		}
		rsl := base.Validate(req)

		expected := []string{}
		assert.Equal(t, expected, rsl)
	})
}