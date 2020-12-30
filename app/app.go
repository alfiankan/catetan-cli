package app

import (
	"catetan-cli/handler"
	"flag"
	"fmt"
)

func StartApp() {
	fmt.Println("Starting App....")
	/*check mode Arg
	-m i => Interaktif
	-m c => Inline
		Inline Flag
		-n <file_name> => Saving with File Name (Default is ctt_{timestamp}
		-u <file_name> => Updating Save With Same Name Before
		-ur <file_name> => Updating and rename
		-h => history
	*/

	flag_mode := flag.String("m", "i", "Mode")
	flag_newCtt := flag.String("n", "ctt_", "Ctt Name")
	flag.Parse()
	switch *flag_mode {
	case "i":
		//inline

		handler.Inline(*flag_newCtt)
		break
	case "c":
		//interaktif
		break
	case "h":

		handler.History()
		break

	}

	//keeping alive at terminal

}
