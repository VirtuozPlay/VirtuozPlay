package repository

import "virtuozplay/models"

func (rs *RepositorySuite) Test_Performances_FindByNanoID() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.FindByNanoID("test-1")

	rs.NoError(err)
	rs.Equal(int64(1), perf.ID)
	rs.Equal(models.NanoID("test-1"), perf.NanoID)
	rs.Equal(models.Base64Notes, perf.NotesEncoding)
	rs.Equal(21, len(perf.Notes))

	// Not in progress, should return error
	perf, err = rs.Performances.FindInProgressByNanoID("test-1")
	rs.Nil(perf)
	rs.Error(err)
}

func (rs *RepositorySuite) Test_Performances_FindById_inProgress() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.New()
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
	perf, err := rs.Performances.New(-1)
	rs.Nil(perf)
	rs.Error(err)
}

func (rs *RepositorySuite) Test_Performances_Create() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.FindByNanoID("in-progress-1")
	rs.Nil(perf)
	rs.Error(err)

	rs.Nil(rs.Performances.Create(&models.Performance{NanoID: "in-progress-1"}))

	perf, err = rs.Performances.FindByNanoID("in-progress-1")
	rs.NoError(err)
	rs.Equal(models.NanoID("in-progress-1"), perf.NanoID)
}

func (rs *RepositorySuite) Test_Performance_Update() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.FindByNanoID("test-1")
	rs.NoError(err)
	rs.NotNil(perf)

	perf.Notes = append(perf.Notes, models.Note{At: 0, Duration: 0, Value: "A"})
	rs.Nil(rs.Performances.Update(perf))

	perf, err = rs.Performances.FindByNanoID("test-1")
	rs.NoError(err)
	rs.Equal(22, len(perf.Notes))
}

func (rs *RepositorySuite) Test_Performance_Update_inProgress() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.New()
	rs.NotNil(perf)
	rs.NoError(err)

	perf.Notes = append(perf.Notes, models.Note{At: 0, Duration: 0, Value: "A"})
	rs.Nil(rs.Performances.Update(perf))

	perf, err = rs.Performances.FindInProgressByNanoID(perf.NanoID)
	rs.NoError(err)
	rs.Equal(1, len(perf.Notes))
}

func (rs *RepositorySuite) Test_Performance_MarkAsFinished() {
	rs.LoadFixture("performance_all_possible_notes")

	perf, err := rs.Performances.New()
	rs.NoError(err)
	rs.NotNil(perf)

	rs.Nil(rs.Performances.MarkAsFinished(perf))
	rs.Error(rs.Performances.MarkAsFinished(perf)) // can't mark as finished twice

	noPerf, err2 := rs.Performances.FindInProgressByNanoID(perf.NanoID)
	rs.Error(err2)
	rs.Nil(noPerf)

	perf2, err3 := rs.Performances.FindByNanoID(perf.NanoID)
	rs.NoError(err3)
	rs.Equal(perf.NanoID, perf2.NanoID)
}

func (rs *RepositorySuite) Test_Performance_MarkAsFinished_invalidPerformance() {
	perf, err := rs.Performances.New()
	rs.NoError(err)
	rs.NotNil(perf)

	perf.Notes = append(perf.Notes, models.Note{At: -888, Duration: 0, Value: "?"})

	rs.Error(rs.Performances.MarkAsFinished(perf))
}
