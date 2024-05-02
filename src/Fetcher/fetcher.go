package Fetcher

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
)

var token sync.Mutex
var finished int
var CSV_errs []string

func FetchFiles(url, branch, whitelist_file string, ch chan error) {
	// Pulls files and returns the paths to said files
	seenFolders := make(map[string]string)
	paths, err := puller(url, branch, whitelist_file)
	if err != nil {
		ch <- err
	}
	splitUrl := strings.Split(url, "/")
	repoName := splitUrl[len(splitUrl)-1]
	for _, path := range paths {
		var fileNamePath []string
		if runtime.GOOS == "windows" {
			fileNamePath = strings.Split(path, "\\")
		} else {
			fileNamePath = strings.Split(path, "/")
		}

		if _, ok := seenFolders[fileNamePath[0]]; !ok {
			seenFolders[fileNamePath[0]] = ""
		}

		// Handle renaming even when two files have the same name, by adding a number to the end
		i := 0
		var newFileName string
		token.Lock()
		for {
			fileName := fileNamePath[len(fileNamePath)-1]
			if i == 0 { // Don't add a 0 to filename
				newFileName = "cache/" + fileName
			} else {
				newFileName = fmt.Sprintf("cache/"+fileName[:len(fileName)-4]+"(%d)"+fileName[len(fileName)-4:], i)
			}

			if _, err := os.Stat(newFileName); os.IsNotExist(err) {
				err := os.Rename(path, newFileName)
				if err != nil {
					panic(err)
				}
				token.Unlock()
				break
			}
			i++
		}

		// Runs data integrity verifier on downloaded file
		// file := "./cache/"+fileName[len(fileName)-1]
		err = Verifier.Verifier(newFileName)
		if err != nil {
			token.Lock()
			CSV_errs = append(CSV_errs, newFileName)
			token.Unlock()
		}

		// Append URL to CSV
		appendUrlToCSV(newFileName, url, repoName)
		
	}

	ch <- nil
}

func appendUrlToCSV(filename, url, repoName string) {

		// Open file and create buffer
		var buf bytes.Buffer
		var scanner *bufio.Scanner
		openfile, err := os.OpenFile(filename, os.O_RDWR, 0644)
		if err != nil {
			log.Printf("Error in opening %s. Error: %v | Continuing...\n", filename, err)
			// Instead of erroring out, we just skip the file and continue
			goto skip
		}
		
		scanner = bufio.NewScanner(openfile)

		// Read first line and write to buffer
		scanner.Scan()
		_, err = buf.Write([]byte(scanner.Text()+"\n"))
		if err != nil {
			log.Printf("Error in writing to buffer. Error: %v | Continuing...\n", err)
			goto skip
		}
		
		// Read rest of the file and append URL
		for scanner.Scan() {
			line := strings.Trim(scanner.Text(),"\n") + fmt.Sprintf("<br>Repos:<br> <a href=%s>%s</a>\n", url, repoName)
			_, err = buf.Write([]byte(line))
			if err != nil {
				log.Printf("Error in writing to buffer. Error: %v | Continuing...\n", err)
				goto skip
			}
		}

		// Write buffer to file
		err = os.WriteFile(filename, buf.Bytes(), 0644)
		if err != nil {
			log.Printf("Error in writing file %s. Error: %v | Continuing...\n", filename, err)
		}
		
		// Skip to close file if error occurs
		skip:
		openfile.Close()
}


func ListingReposForFetch(repos []string) error {
	// Create cache dir if it doesn't exist
	if _, err := os.Stat("cache"); os.IsNotExist(err) {
		err := os.Mkdir("cache", 0700)
		errHandler(err)
	}

	// Create temp folder for .git folders
	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		err := os.Mkdir("temp", 0700)
		errHandler(err)
	}
	defer os.RemoveAll("temp")

	channel := make(chan error)
	for i := 0; i < len(repos); i += 3 {
		go FetchFiles(repos[i], repos[i+1], repos[i+2], channel)
	}
	finished = 0

	// Make sure we print 100% when everything is fetched
	defer func(repos int) {
		progressBar := make([]string, 20)
		for i := 0; i < 20; i++ {
			progressBar[i] = "#"
		}
		fmt.Printf("\r| [%s] %d%%", strings.Join(progressBar, ""), 100)
		// print files with errors
		if len(CSV_errs) == 1 {
			fmt.Println("\n" + "CSV file contains incorrectly formatted content: \n\t" + CSV_errs[0])
		} else if len(CSV_errs) > 1 {
			fmt.Println("\n" + "CSV files contain incorrectly formatted content: \n\t" + strings.Join(CSV_errs, "\n\t"))
		}
	}(len(repos) / 3)
	go progressBar(len(repos) / 3)

	for i := 0; i < len(repos)/3; i++ {
		err := <-channel
		if err != nil {
			return err
		}
		finished++
	}
	return nil
}

// https://stackoverflow.com/questions/39544571/golang-round-to-nearest-0-05
func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func progressBar(numOfFiles int) {
	progressBar := []string{}
	for i := 0; i < 20; i++ {
		progressBar = append(progressBar, ".")
	}
	for {
		for _, r := range `-\|/` {
			percent_fin := float32(finished) / float32(numOfFiles) * 100.0
			rounded_percent := round((float64(finished) / float64(numOfFiles)), 0.05) * 100
			for i := 0; i < int(rounded_percent)/5; i++ {
				progressBar[i] = "#"
			}
			fmt.Printf("\r%c [%s] %d%%", r, strings.Join(progressBar, ""), int32(percent_fin))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func errHandler(err error, params ...string) {
	if err != nil {
		panic(err.Error() + strings.Join(params, " "))
	}
}

func executer(cmd *exec.Cmd, folder string) error {
	//TODO: Figure out a way to take strings as input and build cmd
	cmd.Dir = folder
	_, err := cmd.CombinedOutput()

	return err
}

func puller(url, branch, whitelist_file string) ([]string, error) {
	paths := []string{}

	// Create temp folder for git in the system temp folder
	var randomNum int
	var tempFolder string
	for {
		randomNum = rand.Int()
		tempFolder = fmt.Sprintf("temp/%d", randomNum)
		if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
			os.Mkdir(tempFolder, 0700)
			break
		}
	}

	// Create dummy repo
	cmd := exec.Command("git", "init")
	err := executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	//Enable sparse Checkout
	cmd = exec.Command("git", "config", "core.sparseCheckout", "true")
	err = executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	// Add whitelist to sparse-checkout
	fileData, err := os.ReadFile(whitelist_file)
	if err != nil {
		return paths, err
	}

	err = os.WriteFile(tempFolder+"/.git/info/sparse-checkout", fileData, 0644)
	if err != nil {
		return paths, err
	}

	// Add remote repo
	cmd = exec.Command("git", "remote", "add", "origin", url)
	err = executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	// git pull from remote repo
	cmd = exec.Command("git", "pull", "origin", branch, "--depth=1")
	err = executer(cmd, tempFolder)
	if err != nil {
		return paths, err
	}

	// https://stackoverflow.com/questions/55300117/how-do-i-find-all-files-that-have-a-certain-extension-in-go-regardless-of-depth
	// This function recursively walks the directors inside the workdir and checks for csv files
	// These then get added to the cache later
	filepath.WalkDir(tempFolder, func(str string, dir fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		path_seg := strings.Split(str, "/")
		if path_seg[0] != "cache" {
			if filepath.Ext(dir.Name()) == ".csv" {
				paths = append(paths, str)
			}
		}
		return nil
	})

	return paths, nil
}
