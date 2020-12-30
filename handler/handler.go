package handler

import (
	"bufio"
	"catetan-cli/controller"
	"catetan-cli/view"
	"fmt"
	"os"
	"strings"
	"time"
)

func Inline(ctt_name string) {
	view.Head("Inline")
	fmt.Println("CTT Name Is : " + ctt_name)
	fmt.Println("Put Your Note Here, Type 'done;' to Save")
	fmt.Println("###########################################")
	newTxt := bufio.NewScanner(os.Stdin)
	saveText := ""
	for newTxt.Scan() {

		if strings.ToUpper(newTxt.Text()) == "DONE;" {

			fmt.Println(saveText)
			//saving to database
			res, err := controller.NewCtt(ctt_name, time.Now().Unix(), time.Now().Unix(), saveText)
			fmt.Println(res)
			fmt.Println(err)
			break
		} else {
			saveText += newTxt.Text() + "\n"
		}
		fmt.Println("###########################################")
		//fmt.Println(newTxt.Text())
	}
}

func History() {
	fmt.Println("History")
	controller.GetCttAll()
}
