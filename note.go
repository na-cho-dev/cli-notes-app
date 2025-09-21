package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Note struct {
    Title      string    `json:"title"`
    Body       string    `json:"body"`
    IsFavorite bool      `json:"isFavorite"`
    CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Notes []Note

func (notes *Notes) create(title string) {
	body, err := openEditor("")
	if err != nil {
		fmt.Println("Error opening editor:", err)
		os.Exit(1)
	}

	note := Note{
		Title:      title,
		Body:       body,
		IsFavorite: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Time{},
	}

	*notes = append(*notes, note)
	fmt.Println("Note added.")
}

func (notes *Notes) edit(ind int, title string) error {
	nts := *notes
	if err := nts.ValidateIndex(ind); err != nil {
		fmt.Println(err)
		return err
	}
	ind = ind - 1

	currentBody := nts[ind].Body
	newBody, err := openEditor(currentBody)
	if err != nil {
		fmt.Println("Error opening editor:", err)
		os.Exit(1)
	}

	if title != "" {
    (*notes)[ind].Title = title
	}
	if newBody != "" {
		(*notes)[ind].Body = newBody
	}
	(*notes)[ind].UpdatedAt = time.Now()
	fmt.Println("Note edited.")
	return nil
}

func (notes *Notes) delete(ind int) error {
	nts := *notes
    if err := notes.ValidateIndex(ind); err != nil {
        fmt.Println(err)
        return err
    }
    index := ind - 1

	*notes = append(nts[:index], nts[index+1:]...)
	fmt.Println("Note deleted.")
	return nil
}

func (notes *Notes) toggleFavorite(ind int, isFavorite string) error {
    if err := notes.ValidateIndex(ind); err != nil {
        fmt.Println(err)
        return err
    }

	boolean, err := strconv.ParseBool(isFavorite)
	if err != nil {
		fmt.Println(err)
        return err
	}
    index := ind - 1
	(*notes)[index].IsFavorite = boolean
	fmt.Println("Note marked as favorite.")
	return  nil
}

func (notes *Notes) printAll() {
	// ANSI escape codes for color and style
	const (
		reset   = "\033[0m"
		bold    = "\033[1m"
		blue    = "\033[34m"
		yellow  = "\033[33m"
		cyan    = "\033[36m"
		magenta = "\033[35m"
		divider = "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
		favIcon = "🌟"
		noteIcon = "📝"
	)

	idx := 1
	for _, note := range *notes {
		if note.IsFavorite {
			fmt.Printf("%s%s %sNote #%d %s%s\n", bold, favIcon, yellow, idx, note.Title, reset)
			// fmt.Printf("%s%s%s\n", magenta, divider, reset)
			fmt.Printf("%sBody:%s\n%s\n", bold, reset, note.Body)
			fmt.Printf("%sCreated:%s %s\n", bold, reset, note.CreatedAt.Format("2006-01-02 15:04:05"))
			if !note.UpdatedAt.IsZero() {
				fmt.Printf("%sUpdated At:%s %s\n", bold, reset, note.UpdatedAt.Format("2006-01-02 15:04:05"))
			}
			fmt.Printf("%sFavorite:%s %s\n", bold, reset, favIcon)
			fmt.Printf("%s%s%s\n\n", magenta, divider, reset)
			idx++
		}
	}
	for _, note := range *notes {
		if !note.IsFavorite {
			fmt.Printf("%s%s Note #%d %s%s\n", bold, noteIcon, idx, cyan, note.Title)
			// fmt.Printf("%s%s%s\n", blue, divider, reset)
			fmt.Printf("%sBody:%s\n%s\n", bold, reset, note.Body)
			fmt.Printf("%sCreated:%s %s\n", bold, reset, note.CreatedAt.Format("2006-01-02 15:04:05"))
			if !note.UpdatedAt.IsZero() {
				fmt.Printf("%sUpdated At:%s %s\n", bold, reset, note.UpdatedAt.Format("2006-01-02 15:04:05"))
			}
			fmt.Printf("%s%s%s\n\n", blue, divider, reset)
			idx++
		}
	}
}

func (notes *Notes) printOne(ind int) error {
	// ANSI escape codes for color and style
	const (
		reset   = "\033[0m"
		bold    = "\033[1m"
		blue    = "\033[34m"
		yellow  = "\033[33m"
		cyan    = "\033[36m"
		magenta = "\033[35m"
		divider = "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
		favIcon = "🌟"
		noteIcon = "📝"
	)
	if err := notes.ValidateIndex(ind); err != nil {
		fmt.Println(err)
		return err
	}
	index := ind - 1
	nts := (*notes)[index]
	if nts.IsFavorite {
		fmt.Printf("%s%s %sNote #%d %s%s\n", bold, favIcon, yellow, index+1, nts.Title, reset)
		fmt.Printf("%s%s%s\n", magenta, divider, reset)
		fmt.Printf("%sBody:%s\n%s\n", bold, reset, nts.Body)
		fmt.Printf("%sCreated:%s %s\n", bold, reset, nts.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("%sFavorite:%s %s\n", bold, reset, favIcon)
		fmt.Printf("%s%s%s\n", magenta, divider, reset)
	} else {
		fmt.Printf("%s%s Note #%d %s%s\n", bold, noteIcon, index+1, cyan, nts.Title)
		fmt.Printf("%s%s%s\n", blue, divider, reset)
		fmt.Printf("%sBody:%s\n%s\n", bold, reset, nts.Body)
		fmt.Printf("%sCreated:%s %s\n", bold, reset, nts.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("%s%s%s\n", blue, divider, reset)
	}
	return nil
}

func (notes *Notes) ValidateIndex(ind int) error {
	if ind < 1 || ind > len(*notes) {
        return errors.New("invalid index")
    }
    return nil
}

func openEditor(initialContent string) (string, error) {
	tmpfile, err := os.CreateTemp(os.TempDir(), "note-*.tmp")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(initialContent)); err != nil {
		return "", err
	}
	tmpfile.Close()

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano" // fallback
	}
	cmd := exec.Command(editor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	edited, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		return "", err
	}
	return string(edited), nil
}