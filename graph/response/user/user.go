package user

type Info struct {
	BusinessPhones    []string    `json:"businessPhones"`
	DisplayName       string      `json:"displayName"`
	GivenName         string      `json:"givenName"`
	JobTitle          interface{} `json:"jobTitle"`
	Mail              string      `json:"mail"`
	MobilePhone       interface{} `json:"mobilePhone"`
	OfficeLocation    interface{} `json:"officeLocation"`
	PreferredLanguage string      `json:"preferredLanguage"`
	Surname           string      `json:"surname"`
	UserPrincipalName string      `json:"userPrincipalName"`
	Id                string      `json:"id"`
}
