package main

import (
	"errors"
	gql "github.com/graphql-go/graphql"
	"time"
)

type PersonGql struct {
	Schema *gql.Schema
}

func newContactType() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Contact",
		Fields: gql.Fields{
			"contactType": &gql.Field{Type: gql.String},
			"details": &gql.Field{Type: gql.String},
		},
	})
}

func newPersonType() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Person",
		Fields: gql.Fields{
			"id": &gql.Field{Type: gql.String},
			"firstName": &gql.Field{Type: gql.String},
			"lastName": &gql.Field{Type: gql.String},
			"birthdate": &gql.Field{Type: gql.DateTime},
			"contacts": &gql.Field{Type: gql.NewList(newContactType())},
		},
	})
}

func newQuery() *gql.Object {

	personType := newPersonType()

	return gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"person": &gql.Field{
				Type: personType,
				Description: "Get person by id",
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{ Type: gql.String	},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return &Person{
						Id:        "1",
						FirstName: "Ivan",
						LastName:  "Ivanov",
						Birthdate: time.Date(1980, 10, 30, 0, 0, 0, 0, time.Local),
						Contacts:  []*Contact{
							{
								ContactType: "email",
								Details:     "test@example.com",
							},
							{
								ContactType: "phone",
								Details:     "234-90-34",
							},
						},
					}, nil
				},
			},
			"list": &gql.Field{
				Type: gql.NewList(personType),
				Description: "Get list of persons",
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					var persons []*Person
					persons = append(persons, &Person{
						Id:        "1",
						FirstName: "Ivan",
						LastName:  "Ivanov",
						Birthdate: time.Date(1980, 10, 30, 0, 0, 0, 0, time.Local),
						Contacts:  []*Contact{
							{
								ContactType: "email",
								Details:     "test@example.com",
							},
							{
								ContactType: "phone",
								Details:     "234-90-34",
							},
						},
					})
					persons = append(persons, &Person{
						Id:        "2",
						FirstName: "Petr",
						LastName:  "Petrov",
						Birthdate: time.Date(1980, 10, 30, 0, 0, 0, 0, time.Local),
						Contacts:  []*Contact{
							{
								ContactType: "email",
								Details:     "petrov@example.com",
							},
							{
								ContactType: "phone",
								Details:     "234-90-54",
							},
						},
					})
					return persons, nil
				},
			},
		},
	})
}

func newMutation() *gql.Object {
	return nil
}

func NewPersonGql() (*PersonGql, error) {
	sch, err := gql.NewSchema(gql.SchemaConfig{
		Query:    newQuery(),
		Mutation: newMutation(),
	})
	if err != nil {
		return nil, err
	}

	personGql := &PersonGql{
		Schema: &sch,
	}

	return personGql, nil
}

func (p *PersonGql) ExecQuery(query string) (interface{}, error) {
	res := gql.Do(gql.Params{
		Schema: *p.Schema,
		RequestString: query,
	})
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return res.Data, nil
}
