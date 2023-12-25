package request

type Indikator struct{
	Description string `json:"description" valid:"required~field description is required"`
	TypeID uint `json:"type_id" valid:"required~field type_id is required"`
}

type CreateInstrumenType struct{
	Name string `json:"name" valid:"required~field name is required"`
}

type CreateIndikatorType struct{
	Description string `json:"description" valid:"required~field description is required"`
	InstrumentID uint `json:"instrument_id" valid:"required~field instrument_id is required"`
}
