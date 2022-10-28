package interfaces

type ContactInterface interface {
	Add(...map[string]interface{})
	Edit(...map[string]interface{})
	Delete()
}
