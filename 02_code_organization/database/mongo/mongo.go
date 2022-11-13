package mongo

import architecture "github.com/KenShiiii/Golang-architecture/02_code_organization"

type Db map[int]architecture.User

func (m Db) Save(n int, u architecture.User) {
	m[n] = u
}

func (m Db) Retrieve(n int) architecture.User {
	return m[n]
}
