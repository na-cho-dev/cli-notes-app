package main

func main() {
	notes := Notes{}
	storage := NewStorage[Notes]("note.json")
	storage.Load(&notes)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&notes)
	storage.Save(notes)
}