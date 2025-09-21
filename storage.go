package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return  &Storage[T]{ FileName: fileName }
}

func (st *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return  os.WriteFile(st.FileName, fileData, 0644)
}

func (st *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(st.FileName)

	if err != nil {
		return err 
	}

	return json.Unmarshal(fileData, data)
}