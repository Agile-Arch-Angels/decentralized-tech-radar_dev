package Merger

import (
	"bufio"
	"bytes"
	"os"
	"slices"
	"strings"
)

// Boyer-Moore Majority Vote inspired
// AKA finds the element N with more than N/2 appearances

type Bmmv struct{}

func (b Bmmv) MergeFiles(buffer *bytes.Buffer, filepaths ...string) error {
	// Map functions as a set (name -> quadrant)
	var blipSet = make(map[string][]string)

	// Map functions as a set (blipName -> list of rings)
	var blipRings = make(map[string][]string)

	// auxiliary buffer
	var auxBuffer *bytes.Buffer

	// For every file, use scanFile.
	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		b.scanFile(file, auxBuffer, blipSet, blipRings)

		err = file.Close()
		if err != nil {
			panic(err)
		}
	}

	buffer = b.majorityVote(auxBuffer, blipRings)

	return nil
}

func (b Bmmv) scanFile(file *os.File, buffer *bytes.Buffer, blipSet map[string][]string, blipRings map[string][]string) {
	scanner := bufio.NewScanner(file)

	// Skip header
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		name := ""
		index := strings.IndexByte(line, ',')
		if index != -1 {
			name = line[:index]
		}

		var blipRing = b.duplicateRemoval(name, line, buffer, blipSet)
		blipRings[name] = append(blipRings[name], blipRing)
	}
}

func (b Bmmv) duplicateRemoval(name, line string, buffer *bytes.Buffer, set map[string][]string) string {
	//TODO: Unmarshal the json file (or some other file based solution) to get the alternative names
	// Or just use a baked in str read line by line or combination
	//os.Stat("./Dictionary/alt_names.txt")

	realName := name
	if alt_names[name] != "" {
		//TODO: Figure out how to handle numbers in names
		name = alt_names[strings.ToLower(name)]
	}

	ringLen := len(line[len(realName)+1 : strings.IndexByte(line[len(realName)+1:], ',')+len(realName)+1])
	ringName := line[len(realName)+1 : ringLen+1]
	quadrant := line[len(realName)+ringLen+2 : strings.IndexByte(line[len(realName)+ringLen+2:], ',')+len(realName)+ringLen+2]

	if set[name] != nil {
		// Skips the name + ring + 2 commas and does the same forward search for next comma
		// Example of a line from a csv file:
		// 		Python,hold,language,false,0,Lorem ipsum dolor sit amet consectetur adipiscing elit.
		// Quadrant:
		//		language
		if !(slices.Contains(set[name], quadrant)) {
			set[name] = append(set[name], quadrant)
			buffer.Write([]byte(line + "\n"))
		}
	} else {
		set[name] = append(set[name], line[len(name)+ringLen+2:strings.IndexByte(line[len(name)+ringLen+2:], ',')+len(name)+ringLen+2])
		buffer.Write([]byte(line + "\n"))
	}

	return ringName
}

func (b Bmmv) majorityVote(auxBuffer *bytes.Buffer, blipRings map[string][]string) *bytes.Buffer {

	// buffer to be returned
	finalBuffer := new(bytes.Buffer)

	// For every blip...
	for blipName, listRings := range blipRings {
		currentMajor := ""
		highestCount := 0
		allRings := len(listRings)

		// map: ringName string -> slice of count in int
		// slice 'cause then we can compare []int to nil
		ringCount := make(map[string][]int)

		// for every ring in this blip's list of rings...
		for _, ring := range listRings {

			// if this ring does not exist in the set...
			// it hasn't been encCOUNTered (haha) before.
			if ringCount[ring] != nil {
				ringCount[ring][0]++
			} else {
				ringCount[ring] = make([]int, 0)
			}

			count := ringCount[ring][0]

			// Note: Won't handle ties, so first come, first served.
			if count > allRings/2 && count > highestCount {
				highestCount = count
				currentMajor = ring
			}
		}

		scanner := bufio.NewScanner(auxBuffer)
		lenName := len(blipName)
		lenRing := len(blipRings[blipName])

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if strings.ToLower(line[0:lenName]) == strings.ToLower(blipName) {
				newLine := blipName + "," + currentMajor + line[lenName+lenRing+2:]
				finalBuffer.Write([]byte(newLine + "\n"))
			}
		}
	}

	return finalBuffer
}
