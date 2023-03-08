package db

type MedicationData interface {
	Data() MedicationData
}
type PatientData interface {
	Data() PatientData
}

func (m *Medication) Data() MedicationData {
	return &Medication{
		RxID:         m.RxID,
		PatientID:    m.PatientID,
		RxName:       m.RxName,
		Administered: m.Administered,
	}
}

func (p *Patient) Data() PatientData {
	return &Patient{
		PatientID: p.PatientID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
}
