package testutilsmocks

type StructWithPrivateField struct {
	privateField string
}

func (s *StructWithPrivateField) GetPrivateField() string {
	return s.privateField
}

func NewStructWithPrivateField(privateField string) *StructWithPrivateField {
	return &StructWithPrivateField{privateField}
}
