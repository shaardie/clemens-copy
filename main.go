package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// flags
	targetDir := flag.String("path", ".", "Pfad zum Ordner, der überprüft werden soll")
	deleteFolder := flag.String("delete", "delete", "Name des Ordners für verschobene Dateien")
	dryRun := flag.Bool("dry", false, "Wenn aktiviert, werden keine Dateien verschoben, sondern nur angezeigt")
	flag.Parse()

	// absolute path of the target directory
	absPath, err := filepath.Abs(*targetDir)
	if err != nil {
		fmt.Println("Ungültiger Pfad:", *targetDir)
		return
	}

	// create delete directory, if necessary
	deleteDir := filepath.Join(absPath, *deleteFolder)
	if !*dryRun {
		if _, err := os.Stat(deleteDir); os.IsNotExist(err) {
			err = os.Mkdir(deleteDir, 0755)
			if err != nil {
				fmt.Println("Fehler beim Erstellen des Delete-Ordners:", err)
				return
			}
		}
	}

	jpgFiles := map[string]bool{}
	dirEntries, err := os.ReadDir(absPath)
	if err != nil {
		fmt.Println("Fehler beim Lesen des Verzeichnisses:", err)
		return
	}

	// Find all jpgs and store them in a map
	for _, entry := range dirEntries {
		// Ignore dirs
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext == ".jpg" || ext == ".jpeg" {
			base := strings.TrimSuffix(entry.Name(), ext)
			jpgFiles[base] = true
		}
	}

	// iterate through files and check, if there is a jpg in the previously created map
	for _, entry := range dirEntries {
		// ignore dirs
		if entry.IsDir() {
			continue
		}
		// ignore jpgs
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext == ".jpg" || ext == ".jpeg" {
			continue
		}

		// check if name is present in jpg map and move otherwise.
		base := strings.TrimSuffix(entry.Name(), ext)
		if !jpgFiles[base] {
			srcPath := filepath.Join(absPath, entry.Name())
			dstPath := filepath.Join(deleteDir, entry.Name())
			if *dryRun {
				fmt.Println("[Dry Run] Würde verschieben:", entry.Name())
			} else {
				err := os.Rename(srcPath, dstPath)
				if err != nil {
					fmt.Println("Fehler beim Verschieben:", srcPath, err)
				} else {
					fmt.Println("Verschoben:", entry.Name())
				}
			}
		}
	}

	if *dryRun {
		fmt.Println("Dry Run abgeschlossen. Es wurden keine Dateien verschoben.")
	} else {
		fmt.Println("Fertig.")
	}
}
