package schema

type Organizations struct {
	Base
	Name	string
}

func (Organizations) TableName() string {
	return "organizations"
}

func (Organizations) Pk() string {
	return "id"
}

func (t Organizations) Ref() string {
	return t.TableName() + "(" + t.Pk() + ")"
}

func (t Organizations) AddForeignKeys() {
}

func (t Organizations) InsertDefaults() {
	Database.Exec(`INSERT INTO organizations (name)
	SELECT * FROM (SELECT 'xendit') AS tmp
		WHERE NOT EXISTS (
    		SELECT name FROM organizations WHERE name = 'xendit'
	) LIMIT 1;`)
	Database.Exec(`INSERT INTO organizations (name)
	SELECT * FROM (SELECT 'oracle') AS tmp
		WHERE NOT EXISTS (
    		SELECT name FROM organizations WHERE name = 'oracle'
	) LIMIT 1;`)
}
