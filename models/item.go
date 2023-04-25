package models

/**
 * @Author : bothinli
 * @Description:
 */

type Item interface {
	GetName() string
}

type Job struct {
	Name string
}

func (j Job) GetName() string {
	return j.Name
}
