package util

type ErrUserExists string

func (e ErrUserExists) Error() string {
	return string(e)
}
