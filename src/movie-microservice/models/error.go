/*
 * @File: models.token.go
 * @Description: Defines Error information will be returned to the clients
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package models

// Error defines the response error
type Error struct {
	Code    int    `json:"code" example:"27"`
	Message string `json:"message" example:"Error message"`
}
