package repository

import "virtuozplay/models"

func (rs *RepositorySuite) Test_Performances_implements() {
	rs.Implements((*Repository[models.Performance])(nil), &Performances{})
}

func (rs *RepositorySuite) Test_Performances_FindByNanoID() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.FindByNanoID("perf-1")

	rs.NoError(err)
	rs.Equal(int64(1), perf.ID)
	rs.Equal(models.NanoID("perf-1"), perf.NanoID)
	rs.Equal(models.Base64Notes, perf.NotesEncoding)
	rs.Equal(21, len(perf.Notes))

	// Not in progress, should return error
	perf, err = rs.Performances.FindInProgressByNanoID("perf-1")
	rs.Nil(perf)
	rs.Error(err)
}

func (rs *RepositorySuite) Test_Performances_FindById_inProgress() {
	rs.LoadFixture("song_empty_2")

	song, err := rs.Songs.FindByNanoID("song-1002")
	rs.NotNil(song)
	rs.NoError(err)

	perf, err := rs.Performances.New(song)
	rs.NotNil(perf)
	rs.NotEmpty(perf.NanoID)
	rs.NoError(err)

	perf2, err2 := rs.Performances.FindByNanoID(perf.NanoID)
	rs.NoError(err2)
	rs.Equal(perf.NanoID, perf2.NanoID)

	perf3, err3 := rs.Performances.FindInProgressByNanoID(perf.NanoID)
	rs.NoError(err3)
	rs.Equal(perf.NanoID, perf3.NanoID)
	rs.Equal(perf2.NanoID, perf3.NanoID)
}

func (rs *RepositorySuite) Test_Performances_New_badNanoId() {
	perf, err := rs.Performances.New(nil, -1)
	rs.Nil(perf)
	rs.Error(err)
}

func (rs *RepositorySuite) Test_Performances_Create() {
	rs.LoadFixture("song_empty_3")

	song, err := rs.Songs.FindByNanoID("song-1003")
	rs.NotNil(song)
	rs.NoError(err)

	perf, err := rs.Performances.FindByNanoID("in-progress-1")
	rs.Nil(perf)
	rs.Error(err)

	rs.NoError(rs.Performances.Create(&models.Performance{NanoID: "in-progress-1", Song: song, SongID: song.ID}))

	perf, err = rs.Performances.FindByNanoID("in-progress-1")
	rs.NoError(err)
	rs.Equal(models.NanoID("in-progress-1"), perf.NanoID)
}

func (rs *RepositorySuite) Test_Performance_Update() {
	rs.LoadFixture("performance_to_update")

	perf, err := rs.Performances.FindByNanoID("perf-5")
	rs.NoError(err)
	rs.NotNil(perf)

	perf.Notes = append(perf.Notes, models.Note{At: 0, Duration: 0, Value: "A"})
	rs.Nil(rs.Performances.Update(perf))

	perf, err = rs.Performances.FindByNanoID("perf-5")
	rs.NoError(err)
	rs.Equal(22, len(perf.Notes))
}

func (rs *RepositorySuite) Test_Performance_Update_inProgress() {
	rs.LoadFixture("song_empty_4")

	song, err := rs.Songs.FindByNanoID("song-1004")
	rs.NotNil(song)
	rs.NoError(err)

	perf, err := rs.Performances.New(song)
	rs.NotNil(perf)
	rs.NoError(err)

	perf.Notes = append(perf.Notes, models.Note{At: 0, Duration: 0, Value: "A"})
	rs.Nil(rs.Performances.Update(perf))

	perf, err = rs.Performances.FindInProgressByNanoID(perf.NanoID)
	rs.NoError(err)
	rs.Equal(1, len(perf.Notes))
}

func (rs *RepositorySuite) Test_Performance_MarkAsFinished() {
	rs.LoadFixture("performance_mark_as_finished")

	song, err := rs.Songs.FindByNanoID("song-4")
	rs.NotNil(song)
	rs.NoError(err)

	perf, err := rs.Performances.New(song)
	rs.NoError(err)
	rs.NotNil(perf)

	rs.NoError(rs.Performances.MarkAsFinished(perf))
	rs.Error(rs.Performances.MarkAsFinished(perf)) // can't mark as finished twice

	noPerf, err2 := rs.Performances.FindInProgressByNanoID(perf.NanoID)
	rs.Error(err2)
	rs.Nil(noPerf)

	perf2, err3 := rs.Performances.FindByNanoID(perf.NanoID)
	rs.NoError(err3)
	rs.Equal(perf.NanoID, perf2.NanoID)
}

func (rs *RepositorySuite) Test_Performance_MarkAsFinished_invalidPerformance() {
	perf, err := rs.Performances.New(&models.Song{ID: -999})
	rs.NoError(err)
	rs.NotNil(perf)

	perf.Notes = append(perf.Notes, models.Note{At: -888, Duration: 0, Value: "?"})

	rs.Error(rs.Performances.MarkAsFinished(perf))
}
