package repository

type StackOverFlowRepository interface {
	FindAll(url string) (interface{}, error)
	Single(url string, number int) (interface{}, error)
}
