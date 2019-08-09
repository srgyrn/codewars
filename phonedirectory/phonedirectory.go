package phonedirectory

import (
	"fmt"
	"regexp"
	"strings"
)

type contact struct {
	phoneNumber string
	name        string
	address     string
}

func Phone(dir, num string) string {
	matches := regexp.MustCompile("\\+"+num).FindAllStringSubmatch(dir, -1)

	if len(matches) == 0 {
		return "Error => Not found: " + num
	}

	if len(matches) > 1 {
		fmt.Println(dir, matches)
		return "Error => Too many people: " + num
	}

	var record contact

	for _, line := range strings.Split(dir, "\n") {
		record.setPhoneNumberFromString(line, dir)

		if record.phoneNumber == "error" {
			return "Error => Too many people: " + num
		}

		record.setName(line)
		record.setAddress(line)

		if record.phoneNumber == num {
			return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", record.phoneNumber, record.name, record.address)
		}
	}

	return "Error => Not found: " + num
}

func (contactp *contact) setPhoneNumberFromString(s string, completeRecords string) {
	contactp.phoneNumber = regexp.MustCompile(`(\d{1,2}-\d{3}-\d{3}-\d{4})`).FindString(s)
}

func (contactp *contact) setName(s string) {
	matched := regexp.MustCompile(`\<(.*?)\>`).FindStringSubmatch(s)
	contactp.name = matched[1]
}

func (contactp *contact) setAddress(s string) {
	rx := regexp.MustCompile(`(\<(.*?)\>)|(\d{1,2}-\d{3}-\d{3}-\d{4})|[^\w\s\.-]|_`)
	address := rx.ReplaceAllString(s, " ")
	address = strings.TrimSpace(strings.TrimSpace(address))
	contactp.address = regexp.MustCompile(`\s+`).ReplaceAllString(address, " ")
}
