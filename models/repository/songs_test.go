package repository

import "virtuozplay/models"

func (rs *RepositorySuite) Test_Songs_implements() {
	rs.Implements((*Repository[models.Song])(nil), &Songs{})
}

func (rs *RepositorySuite) Test_Songs_FindByNanoID() {
	rs.LoadFixture("song_empty")

	song, err := rs.Songs.FindByNanoID("song-1001")

	rs.NoError(err)
	rs.Equal(int64(1001), song.ID)
	rs.Equal(models.NanoID("song-1001"), song.NanoID)
	rs.Equal("Empty Song of Silence", song.Title)
}
