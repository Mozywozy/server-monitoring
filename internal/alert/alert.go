package alert

import "fmt"

func SendAlert(server string) {
	fmt.Printf("Alert: Server %s is DOWN!\n", server)
}
