package data

func GetPhoneCallsDataGql() (gql []string, err error) {

	gql = []string{}

	companiesGql, err := GetCompaniesGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, companiesGql...)

	peopleGql, err := GetPeopleGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, peopleGql...)

	contractsGql, err := GetContractsGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, contractsGql...)

	callsGql, err := GetCallsGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, callsGql...)

	return gql, nil
}