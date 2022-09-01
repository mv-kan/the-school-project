package service

import (
	"testing"

	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnrollService_Enroll(t *testing.T) {
	db := connectToDB()
	enrollService := NewEnroll(db)
	enrolledPupil, err := enrollService.Enroll(testingdb.TestPupil)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestPupil.Name, enrolledPupil.Name)
}
