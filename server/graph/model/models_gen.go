// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Coords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Session struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type User struct {
	Name   string  `json:"name"`
	Radius int     `json:"radius"`
	Coords *Coords `json:"coords"`
}
