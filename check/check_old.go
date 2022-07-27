package main

import "strings"

func numUniqueEmails(emails []string) int {
	if len(emails) < 1 || len(emails) > 100 {
		// handle error
		return -1
	}

	set := make(map[string]struct{})
	for _, email := range emails {
		if parse(email) {
			res := strings.Split(email, "@")
			localname, domainName := res[0], res[1]

			localname = strings.Replace(localname, ".", "", -1)
			plusSplitRes := strings.Split(localname, "+")
			localname = plusSplitRes[0]

			set[strings.Join([]string{localname, domainName}, "@")] = struct{}{}
		}
	}

	return len(set)
}

func parse(email string) bool {
	if len(email) < 1 || len(email) > 100 {
		return false
	}

	if strings.ToLower(email) != email {
		return false
	}

	res := strings.Split(email, "@")

	if len(res) != 2 {
		return false
	}

	localname, domainName := res[0], res[1]

	if localname == "" || domainName == "" {
		return false
	}

	if strings.HasPrefix(localname, "+") {
		return false
	}

	if strings.HasSuffix(domainName, ".com") == false {
		return false
	}

	return true
}
