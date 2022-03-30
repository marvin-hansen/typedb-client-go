package data

import (
	"bytes"
	"encoding/json"
	"io"
	"text/template"
)

type Company struct {
	Name string `json:"name"`
}

var insertCompanyTemplate = template.Must(template.New("InsertCompany").Parse(`
insert 
	$_ isa company, 
		has name "{{.Name}}";
`))

var insertCompaniesTemplate = template.Must(template.New("InsertCompanies").Parse(`
insert
{{range .}}
	$_ isa company, has name "{{.Name}}";
{{end}}
`))

func GetCompanies() (companies []Company, err error) {
	data := getCompanyData()
	err = json.Unmarshal([]byte(data), &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func GetCompanyGql(company *Company, wr io.Writer) (err error) {
	err = insertCompanyTemplate.Execute(wr, company)
	if err != nil {
		return err
	}
	return err
}

func GetCompaniesGql() (gql []string, err error) {
	gql = []string{}
	companies, err := GetCompanies()
	if err != nil {
		return nil, err
	}

	templateRendered := bytes.Buffer{}
	err = insertCompaniesTemplate.Execute(&templateRendered, companies)
	if err != nil {
		return nil, err
	}
	gql = append(gql, templateRendered.String())
	return gql, err
}
