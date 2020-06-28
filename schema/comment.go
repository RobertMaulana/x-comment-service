package schema

type Comments struct {
	Base
	Comment	string
	OrganizationId int
}

func (Comments) TableName() string {
	return "comments"
}

func (Comments) Pk() string {
	return "id"
}

func (t Comments) Ref() string {
	return t.TableName() + "(" + t.Pk() + ")"
}

func (t Comments) AddForeignKeys() {
	Database.Model(&t).AddForeignKey("organization_id", Organizations{}.Ref(), "CASCADE", "RESTRICT")
}

func (t Comments) InsertDefaults() {
}
