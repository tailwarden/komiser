package models

type DBInstance struct {
	Status           string
	StorageType      string
	AllocatedStorage int64
	InstanceClass    string
	Engine           string
}
