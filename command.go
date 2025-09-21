package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	AddNote string
	Del int
	Edit string
	Favorite int
	List bool
	ListOne int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.AddNote, "create", "", "Create a new note: Specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a note by index and specify new title/body. id:new_title:new_body")
	flag.IntVar(&cf.Del, "del", -1, "Specify a note by index to delete.")
	flag.IntVar(&cf.Favorite, "favorite", -1, "Specify a note by index to set as favorite.")
	flag.BoolVar(&cf.List, "list", false, "List all notes.")
	flag.IntVar(&cf.ListOne, "select", -1, "List one note by index.")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(notes *Notes) {
	switch {
		case cf.List:
			notes.printAll()
		case cf.ListOne != -1:
			notes.printOne(cf.ListOne)
		case cf.AddNote != "":
			title := cf.AddNote
			notes.create(title)
		case cf.Edit != "":
			var title string
			parts := strings.SplitN(cf.Edit, ":", 2)
			if len(parts) == 2 {
				title = parts[1]
			}
			ind, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error, invalid index for edit")
				os.Exit(1)
			}
			notes.edit(ind, title)
		case cf.Favorite != -1:
			notes.setFavorite(cf.Favorite)
		case cf.Del != -1:
			notes.delete(cf.Del)
		default:
			fmt.Println("Invalid command")
	}
}