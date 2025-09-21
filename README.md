# CLI Notes App

A modern, interactive command-line notes application written in Go. Effortlessly create, edit, organize, and view your notes directly from the terminal, with a beautiful and user-friendly interface.

## Features

- **Add Notes**: Create new notes with a title and a rich, multi-line body using your favorite text editor.
- **Edit Notes**: Update the title and body of any note, with changes tracked by timestamps.
- **Delete Notes**: Remove notes by their index.
- **List Notes**: View all notes, with favorites shown first and a modern, colorized display.
- **Mark as Favorite**: Highlight important notes to always see them at the top.
- **Persistent Storage**: All notes are saved in a JSON file (`note.json`) for easy backup and portability.
- **Interactive Editing**: Uses your `$EDITOR` (or nano by default) for comfortable note writing and editing.

## Usage

### Commands

- `-add-note "Title"` : Add a new note. The body will be entered in your editor.
- `-edit N[:new_title]` : Edit note at index N (1-based). Optionally provide a new title.
- `-del N` : Delete note at index N.
- `-favorite N` : Mark note at index N as favorite.
- `-list` : List all notes, favorites first.

### Example

```sh
# Add a note
$ go run . -create "Meeting Notes"

# List notes
$ go run . -list

# List one single note
$ go run . -select 5

# Edit the second note title or body
$ go run . -edit "2:Updated Title"

# Edit the second note body
$ go run . -edit 2

# Delete the third note
$ go run . -del 3

# Mark the first note as favorite
$ go run . -favorite 1
```

## Requirements

- Go 1.18+
- A terminal with ANSI color support

## How It Works

- Notes are stored in `note.json` in the project directory.
- The app uses your `$EDITOR` for editing note bodies (defaults to nano if not set).
- Notes are indexed starting from 1 for user-friendliness.
- Favorite notes are always displayed first in the list.

## Customization

- You can change the default editor by setting the `EDITOR` environment variable:
  ```sh
  export EDITOR=vim
  ```

## Contributing

Pull requests and suggestions are welcome! Please open an issue to discuss your ideas or report bugs.

## License

MIT License
