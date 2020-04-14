package clients

type IClient interface {
	Get(limit int, filter interface{}) []interface{}
}
