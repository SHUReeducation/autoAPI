package typeTransform

type TypeTransformer interface {
	Transform(dataBaseType string) string
}
