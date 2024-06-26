package Merger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NovoNordisk-OpenSource/decentralized-tech-radar/Verifier"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MergeStrat interface {
	// A function that updates the buffer with the correct information
	// depending on that merge strategy
	MergeFiles(buffer *bytes.Buffer, filepaths ...string) error
}

func getHeader(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []byte{}, err // Propagate error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	headerBytes := scanner.Bytes()
	headerBytes = append(headerBytes, []byte("\n")...)

	return headerBytes, nil
}

func MergeFromFolder(folderPath string, start MergeStrat) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return errors.New("Folder does not exist or could not be found. \nError: " + err.Error())
	} else if err != nil {
		return err
	}

	cachedRepos, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	var cachePaths []string
	for _, repo := range cachedRepos {
		if filepath.Ext(repo.Name()) == ".csv" {
			cachePaths = append(cachePaths, filepath.Join(folderPath, repo.Name()))
		}
	}

	if len(cachePaths) == 0 {
		fmt.Println("There are currently no files in the cache.")
	}

	MergeCSV(cachePaths, start)

	return nil
}

func MergeCSV(filepaths []string, strat MergeStrat) error {
	os.Remove("Merged_file.csv") // Remove file in case it already exists

	// Run data verifier on files
	err := Verifier.Verifier(filepaths...)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	// Add header to buffer
	header, err := getHeader(filepaths[0])
	if err != nil {
		return err // Propagate error
	}
	buf.Write(header)

	// Read csv data which removes duplicates
	// This only adds non-duplicates to the buffer
	err = strat.MergeFiles(&buf, filepaths...)
	if err != nil {
		panic(err)
	}

	// Write combined files to one file
	err = os.WriteFile("Merged_file.csv", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func zapLogger(f *os.File) *zap.SugaredLogger {
	// https://stackoverflow.com/questions/50933936/zap-logger-print-both-to-console-and-to-log-file
	pe := zap.NewProductionEncoderConfig()

	//fileEncoder := zapcore.NewJSONEncoder(pe)

	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.LevelKey = zapcore.OmitKey
	cfg.EncoderConfig.TimeKey = zapcore.OmitKey

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(cfg.EncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(f), zap.InfoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel),
	)

	l := zap.New(core)

	return l.Sugar()

}

// Create the string and write it to the buffer
func bufferWriter(buffer *bytes.Buffer, blips map[string]map[string]byte) error {
	var sb strings.Builder
	for line, intMap := range blips {
		sb.WriteString(line)       // Write the main line
		for atag := range intMap { // Write the repos (doesn't run if no URLs in the internal map)
			sb.WriteString("<br>")
			sb.WriteString(atag)
		}
		sb.WriteRune('\n')
	}

	buffer.WriteString(sb.String())

	return nil
}
