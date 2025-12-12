package fn

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testError     = errors.New("testing error")
	addNameError  = errors.New("add name error")
	addIdError    = errors.New("add id error")
	addEmailError = errors.New("add email error")
	addJobError   = errors.New("add job error")
)

type myFirstType struct {
	name string
	id   int
}

type mySecondType struct {
	name  string
	id    int
	email string
}

type myThirdType struct {
	name  string
	id    int
	email string
	job   string
}

func addName(v myFirstType) (myFirstType, error) {
	return myFirstType{
		name: "jane doe",
	}, nil
}

func addId(v myFirstType) (myFirstType, error) {
	return myFirstType{
		name: v.name,
		id:   1234,
	}, nil
}

func addEmail(v myFirstType) (mySecondType, error) {
	return mySecondType{
		name:  v.name,
		id:    v.id,
		email: "jane_doe@work.com",
	}, nil
}

func addJob(v mySecondType) (myThirdType, error) {
	return myThirdType{
		name:  v.name,
		id:    v.id,
		email: v.email,
		job:   "sales",
	}, nil
}

func errorOut(v myFirstType) (mySecondType, error) {
	return mySecondType{}, testError
}

func TestTry(t *testing.T) {
	start := myFirstType{}

	withName := Try(addName, start, addNameError)

	withId := Try(addId, withName, addIdError)

	withEmail := Try(addEmail, withId, addEmailError)

	final, err := Try(addJob, withEmail, addJobError).Out()
	assert.NoError(t, err)
	assert.Equal(t, "jane doe", final.name)
	assert.Equal(t, 1234, final.id)
	assert.Equal(t, "jane_doe@work.com", final.email)
	assert.Equal(t, "sales", final.job)
}

func TestTryErrors(t *testing.T) {
	start := myFirstType{}

	withName := Try(addName, start, addNameError)

	withId := Try(addId, withName, addIdError)

	withEmail := Try(errorOut, withId, addEmailError)

	final, err := Try(addJob, withEmail, addJobError).Out()
	assert.Nil(t, final)
	assert.Error(t, err)
	assert.ErrorIs(t, err, testError)
	assert.ErrorIs(t, err, addEmailError)
}

func TestTryBadInput(t *testing.T) {
	start := myFirstType{}

	_ = Try(addName, start, addNameError)

	withId := Try(addId, "hello", addIdError)

	withEmail := Try(errorOut, withId, addEmailError)

	final, err := Try(addJob, withEmail, addJobError).Out()
	assert.Nil(t, final)
	assert.Error(t, err)
	assert.ErrorIs(t, err, addIdError)
}
