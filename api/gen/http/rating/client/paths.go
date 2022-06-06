// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the rating service.
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"fmt"
)

// GetRatingPath returns the URL path to the rating service Get HTTP endpoint.
func GetRatingPath(id uint) string {
	return fmt.Sprintf("/resource/%v/rating", id)
}

// UpdateRatingPath returns the URL path to the rating service Update HTTP endpoint.
func UpdateRatingPath(id uint) string {
	return fmt.Sprintf("/resource/%v/rating", id)
}
