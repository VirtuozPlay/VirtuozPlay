package repository

import "virtuozplay/models"

func (rs *RepositorySuite) Test_Users_implements() {
	rs.Implements((*Repository[models.User])(nil), &Users{})
}

func (rs *RepositorySuite) Test_Users_FindByNanoID() {
	rs.LoadFixture("user_basic")

	user, err := rs.Users.FindByNanoID("user-5001")

	rs.NoError(err)
	rs.Equal(int64(5001), user.ID)
	rs.Equal(models.NanoID("user-5001"), user.NanoID)
}
