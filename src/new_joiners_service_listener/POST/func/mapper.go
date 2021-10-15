package mapper

import (
	"encoding/json"
	"fmt"
	joiner "handler/subscriber/func/dataaccess"
	"strings"
)

type Mapper struct {
	Body string
}

func (m Mapper) Map(body string) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)

	if err != nil {
		panic(err)
	}

	j := new(joiner.Joiner)
	newJoinerAcknowledgement, err := j.CreateJoinerMessageAcknowledgement(strings.ReplaceAll(body, "'", ""))

	if err != nil {
		panic(err)
	}

	var stackTag, _ = json.Marshal(data["PRODUCT"])
	var personTag, _ = data["PERSON"].([]interface{})
	var languageTag, _ = json.Marshal(data["LANGUAGE"])
	nounPhraseTag := data["Noun_Phrase"].([]interface{})
	nounPhrases := make([]string, len(nounPhraseTag))
	persons := make([]string, len(personTag))

	for i, v := range nounPhraseTag {
		nounPhrases[i] = v.(string)
	}

	for i, v := range personTag {
		persons[i] = v.(string)
	}

	var roleIndex = FindKeyPosition(nounPhrases, "Role")
	fmt.Printf("Role Name Index: %d", roleIndex)

	var role string
	if roleIndex > 0 {
		role = nounPhrases[roleIndex]
	}

	var joinerIndex = FindKeyPosition(nounPhrases, "Name")
	fmt.Printf("Joiner Name Index: %d", joinerIndex)

	var name string
	if joinerIndex > 0 {
		name = nounPhrases[joinerIndex]
	}

	if name == "" {
		name = persons[0]
	}

	var doesJoinerExists, _ = j.GetJoiner(name)

	if name != "" && doesJoinerExists == 0 {
		j.Name = name
		j.Role = role
		j.Stack = string(stackTag)
		j.Languages = string(languageTag)
		j.JoinerMessageAcknowledgementId = newJoinerAcknowledgement
		j.CreateJoiner()
		j.UpdateJoinerMessageAcknowledgementStatus()
	}
}

func FindKeyPosition(nonPhrases []string, key string) int {

	for i, d := range nonPhrases {
		if d == key {
			return i + 1
		}
	}

	return -1
}
