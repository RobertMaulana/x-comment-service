package query

const (
	GetOrganizationId = `SELECT id, name FROM organizations WHERE name LIKE $1`
	CreateComment = `INSERT INTO comments (comment, organization_id) VALUES ($1, $2) RETURNING id;`
	GetAllComments = `SELECT id, comment, created_at FROM comments WHERE organization_id = $1 AND deleted_at IS NULL;`
	CreateOrganization = `INSERT INTO organizations (name) VALUES ($1) RETURNING id;`
	DeleteComments = `UPDATE comments SET deleted_at=$1 WHERE organization_id = $2;`
)
