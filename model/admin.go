package model

type Admin struct {
	ID        uint   `database:"primary_key" json:"id"`
	Name      string `database:"unique_index" json:"name"`
	Password  string `database:"not null;char(60)" json:"password"`
	Introduce string `database:"varchar(50)" json:"introduce"`
	Picture   string `database:"not null;varchar(255)" json:"picture"`
	CreatedAt string `database:"not null;datetime" json:"picture"`
}

func (admin *Admin) TableName() string {
	return "admin"
}
