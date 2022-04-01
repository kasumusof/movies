package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Comment is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Comment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Author    string    `json:"author" db:"author"`
	MovieID   int       `json:"movie_id" db:"movie_id"`
	Text      string    `json:"text" db:"text"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NewComment
func NewComment(comment Comment) (*Comment, error) {
	if err = tx.Create(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetComments
func GetComments(movieID int) (*Comments, error) {
	comment := Comments{}
	query := tx.Where("movie_id = ?", movieID).Order("created_at desc")
	err := query.All(&comment)
	if err != nil {
		return &Comments{}, err
	}
	return &comment, nil
}

// String is not required by pop and may be deleted
func (c Comment) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Comments is not required by pop and may be deleted
type Comments []Comment

// String is not required by pop and may be deleted
func (c Comments) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Comment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Author, Name: "Author"},
		&validators.IntIsPresent{Field: c.MovieID, Name: "MovieID"},
		&validators.StringIsPresent{Field: c.Text, Name: "Text"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Comment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Comment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
