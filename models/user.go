package user

type Users struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//temp data for mimic db
var DummyUsers = []Users{
	{ID: "1", Name: "jane Doe", Email: "janedoe@email.com"},
	{ID: "2", Name: "mark", Email: "mark@email.com"},
}
