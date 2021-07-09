package main

// Github Repo Struct
type Github struct {
	Username, Repo, Branch string
}

// Unikorn Command Struct
type Command struct {
	Name, Description, Usage string
	Handler                  func(params []string)
}
