package apiGenerator

type TypeTransformer interface {
	Transform(dataBaseType string) string
}
