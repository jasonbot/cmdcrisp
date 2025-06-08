// build: -darwin
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var (
		a string // application to use for opening the file
		b string // bundle identifier for the application to use when opening the file
		D bool   // Reveal the enclosing folder in finder
		e bool   // Open the file with /Applications/TextEdit
		t bool   // Open the file with the default text editor, as determined via Launch-Services
		f bool   // Read input from standard input and open the results in the default text editor
		F bool   // Open the application "fresh"
		W bool   // Wait until the applications exit
		n bool   // Open a new instance of the application(s)
		g bool   // Do not bring the application to the foreground
		h string // Search header locations for a header whose name matches the given string and then open it
	)

	// Parse flags
	flag.StringVar(&a, "a", "", "The application to use for opening the file")
	flag.StringVar(&b, "b", "", "The bundle identifier for the application to use when opening the file")
	flag.BoolVar(&D, "D", false, "Reveal the enclosing folder in finder")
	flag.BoolVar(&e, "e", false, "Open the file with /Applications/TextEdit")
	flag.BoolVar(&t, "t", false, "Open the file with the default text editor, as determined via Launch-Services")
	flag.BoolVar(&f, "f", false, "Read input from standard input and open the results in the default text editor")
	flag.BoolVar(&F, "F", false, "Open the application 'fresh'")
	flag.BoolVar(&W, "W", false, "Wait until the applications exit")
	flag.BoolVar(&n, "n", false, "Open a new instance of the application(s)")
	flag.BoolVar(&g, "g", false, "Do not bring the application to the foreground")
	flag.StringVar(&h, "h", "", "Search header locations for a header whose name matches the given string and then open it")
	flag.Parse()

	if a != "" || b != "" || D || f || F || W || n || g || h != "" {
		log.Fatal("I can't handle every arg")
	}

	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Nothing to do!")
	}

	command := "xdg-open"

	if runtime.GOOS == "Windows" {
		command = "explorer.exe"

		if e || t {
			command = "notepad.exe"
		}
	} else {
		if e || t {
			command = os.Getenv("VISUAL")
			if command == "" {
				command = os.Getenv("EDITOR")
			}
			if command == "" {
				command = "vi"
			}
		}
	}

	cmd := exec.Command(command, args...)
	cmd.Run()
}
