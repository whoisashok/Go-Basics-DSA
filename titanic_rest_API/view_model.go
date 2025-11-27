package main

import "gorm.io/gorm"

type Database struct {
	DB *gorm.DB
}

type FareData struct {
	Fare float64 `json:"fare"`
}

type Tbl_Titanic struct {
	PassengerId int     `gorm:"primaryKey;autoIncrement;not null"`
	Survived    int     `gorm:"unique;not null"`
	Pclass      int     `gorm:"unique;not null"`
	Name        string  `gorm:"unique;not null"`
	Sex         string  `gorm:"unique;not null"`
	Age         int     `gorm:"unique;not null"`
	SibSp       int     `gorm:"unique;not null"`
	Parch       int     `gorm:"unique;not null"`
	Ticket      string  `gorm:"unique;not null"`
	Fare        float64 `gorm:"unique;not null"`
	Cabin       string  `gorm:"unique;not null"`
	Embarked    string  `gorm:"unique;not null"`
}

type TitanicJSON struct {
	PassengerId int     `json:"PassengerId"`
	Survived    int     `json:"Survived"`
	Pclass      int     `json:"Pclass"`
	Name        string  `json:"Name"`
	Sex         string  `json:"Sex"`
	Age         int     `json:"Age"`
	SibSp       int     `json:"SibSp"`
	Parch       int     `json:"Parch"`
	Ticket      string  `json:"Ticket"`
	Fare        float64 `json:"fare"`
	Cabin       string  `json:"Cabin"`
	Embarked    string  `json:"Embarked"`
}
