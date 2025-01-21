package muiscaDomain

type Config struct {
	URL string
}

const (
	NaturalPersonType   = "NATURAL_PERSON"
	JuridicalPersonType = "JURIDICAL_PERSON"
)

type Result struct {
	NIT             string
	DV              string
	State           string
	ContributorType string
	NaturalPerson   NaturalPerson
	JuridicalPerson JuridicalPerson
}

type NaturalPerson struct {
	FirstName      string
	MiddleName     string
	LastName       string
	SecondLastName string
}

type JuridicalPerson struct {
	SocialReason string
}
