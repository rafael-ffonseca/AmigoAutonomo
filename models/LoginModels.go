package models

type GetRequestLoginInput struct {
	username	string
	password	string
}

type GetRequestLoginOutput struct {
	id        			uint
	username 			string
	firstName			string
	lastName			string
	email				string
}
