package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnrollService_Enroll(t *testing.T) {
	db := connectToDB()
	enrollService := NewEnroll(db)
	enrolledPupil, err := enrollService.Enroll(testPupil)
	require.Nil(t, err)
	assert.Equal(t, testPupil.Name, enrolledPupil.Name)
}
