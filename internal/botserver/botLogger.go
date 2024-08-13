package botserver

import "log"

func botServLog(t string, e error) {
	if e != nil {
		log.Println("[bot server] ", t, ", error: ", e.Error())
		return
	}

	log.Println("[bot server] ", t, "  success.")
}
