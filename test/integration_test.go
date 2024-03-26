package test

import (
	"os"
	"strings"
	"testing"

	html "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/HTML"
	Reader "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
)

// Tests
// Integration test
func TestReaderAndWriter(t *testing.T) {
	// Set up
	createCsvFile()
	defer cleanUp()

	// Read test file
	specs := Reader.ReadCsvSpec(testFileName + ".csv")
	html.GenerateHtml(specs)

	assertIndexHTML(t)
}

// Test Set up
var testFileName string = "ForTesting"

var csvTestString string = `name,ring,quadrant,isNew,moved,description
TestBlip1,Assess,Language,true,1,This is a description
TestBlip2,Adopt,Tool,false,0,Also a description`

var correctHTML string = `<html>
	<head>
		<title>Header 1</title>
	</head>
	<body>
		<h1 class="pageTitle">Header 1</h1>
		<ul>
			
					<li>Name: TestBlip1</li>
					<li>Quadrant: Language</li>
					<li>Ring: Assess</li>
					<li>Is new: true</li>
					<li>Moved: 1</li>
					<li>Desc: This is a description</li>
			
					<li>Name: TestBlip2</li>
					<li>Quadrant: Tool</li>
					<li>Ring: Adopt</li>
					<li>Is new: false</li>
					<li>Moved: 0</li>
					<li>Desc: Also a description</li>
			
		</ul>
	</body>
</html>`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createCsvFile() {
	err := os.WriteFile(testFileName+".csv", []byte(csvTestString), 0644)
	check(err)
}

func cleanUp() {
	os.Remove(testFileName + ".csv")
	os.Remove("index.html")
	//Works on Unix and Windows
	os.Remove("tech_radar.exe")
}

func assertIndexHTML(t *testing.T) {
	//check if the index.html was created
	_, err := os.Stat("index.html")
	if os.IsNotExist(err) {
		t.Fatal("Expected HTML file was not created.")
	}

	//read content of the HTML file
	content, err := os.ReadFile("index.html")
	if err != nil {
		t.Fatalf("Could not read the generated HTML file: %v", err)
	}
	contentStr := string(content)

	//check if content contains exptected string
	if !strings.Contains(contentStr, correctHTML) {
		t.Errorf("HTML doesn't contain the expected data\nContained:\n%s", contentStr)
	}
}