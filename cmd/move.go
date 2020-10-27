package cmd

import (
	"file-manage/helper"
	"file-manage/model"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move [source folder] [target folder]",
	Short: "Move files from source folder to target folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceFolder := args[0]
		targetFolder := args[1]

		source, err := filepath.Abs(sourceFolder)
		must(err)
		target, err := filepath.Abs(targetFolder)
		must(err)

		// Check if source folder and target folder exists
		if _, err := os.Stat(source); os.IsNotExist(err) {
			must(fmt.Errorf("Source folder %v does not exists.", sourceFolder))
		}
		if _, err := os.Stat(target); os.IsNotExist(err) {
			must(fmt.Errorf("Target folder %v does not exists.", targetFolder))
		}

		files, err := getFilteredFiles(source)
		must(err)

		for _, file := range files {
			newFileName, err := helper.GetNewFilePath(file, source, target)
			must(err)

			newFile := model.ConstructFile(newFileName)
			err = move(file, newFile)
			must(err)

			if enableLog {
				fmt.Printf("%v: Move %v --> %v\n", time.Now().Format(time.RFC3339), file.FullPath, newFileName)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(moveCmd)
}

func move(src, dest model.File) error {
	if _, err := os.Stat(dest.Folder); os.IsNotExist(err) {
		if err = os.MkdirAll(dest.Folder, os.ModePerm); err != nil {
			return err
		}
	}

	return os.Rename(src.FullPath, dest.FullPath)
}
