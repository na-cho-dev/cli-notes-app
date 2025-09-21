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
	Favorite string
	List bool
	ListOne int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.AddNote, "create", "", "Create a new note: Specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a note by index and specify new title/body. id:new_title:new_body")
	flag.IntVar(&cf.Del, "del", -1, "Specify a note by index to delete.")
	flag.StringVar(&cf.Favorite, "favorite", "Favorite", "Specify a note by index to set as favorite.")
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
		case cf.Favorite != "":
			parts := strings.SplitN(cf.Favorite, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Error, invalid format for favorite. Please use id:true or id:false")
				os.Exit(1)
			}
			if parts[1] != "true" && parts[1] != "false" {
				fmt.Println("Error, invalid value for favorite. Please use id:true or id:false")
				os.Exit(1)
			}

			ind, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error, invalid index for favorite")
				os.Exit(1)
			}
			notes.toggleFavorite(ind, parts[1])
		case cf.Del != -1:
			notes.delete(cf.Del)
		default:
			fmt.Println("Invalid command")
	}
}