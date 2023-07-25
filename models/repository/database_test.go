package repository

import "virtuozplay/models"

func (rs *RepositorySuite) Test_Database_implements() {
	rs.Implements((*Repository[models.Value])(nil), &DatabaseRepository[models.Value]{})
}
