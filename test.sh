curl -g 'http://localhost:8080/person?query={person(id:"1"){id,firstName,lastName,birthdate,contacts{contactType, details}}}'
curl -g 'http://localhost:8080/person?query={list{id,firstName,lastName,birthdate,contacts{contactType, details}}}'
