package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang Introduction",
		Description: "A golang introduction repository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	//variable := request["field"].(bool)	=> reflection
	if request.Private {

	}

	//Marshal takes an input interface and attempts to create a vaid json string
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	var target CreateRepoRequest

	//Unmarshal takes an input byte array and a *pointer* that we're trying to fill using this json.
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, `{"name":"Golang Introduction","description":"A golang introduction repository","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`,
		string(bytes))
}
