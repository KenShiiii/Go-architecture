package harddrive

import architecture "github.com/KenShiiii/Golang-architecture/02_code_organization"

type Db map[int]architecture.User

func (hd Db) Save(n int, u architecture.User) {
	hd[n] = u
}

func (hd Db) Retrieve(n int) architecture.User {
	return hd[n]
}
