package common

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func StringInSlice(slice []string, item string) int {
	for index, value := range slice {
		if value == item {
			return index
		}
	}
	return -1
}

func UintInSlice(slice []uint, item uint) int {
	for index, value := range slice {
		if value == item {
			return index
		}
	}
	return -1
}

func UintSliceDifference(leftSlice []uint, rightSlice []uint) (left []uint, right []uint) {
	left = make([]uint, 0)
	right = make([]uint, 0)
	for _, value := range leftSlice {
		if existedIx := UintInSlice(rightSlice, value); existedIx == -1 {
			left = append(left, value)
		}
	}
	for _, value := range rightSlice {
		if existedIx := UintInSlice(leftSlice, value); existedIx == -1 {
			right = append(right, value)
		}
	}
	return left, right
}

func GetRandomUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
