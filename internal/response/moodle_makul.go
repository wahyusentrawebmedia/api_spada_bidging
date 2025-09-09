package response

type MoodleMakulRequest struct {
	KodeMK   string `json:"kode_mk"`
	NamaMK   string `json:"nama_mk"`
	Tahun    string `json:"tahun"`
	Semester string `json:"semester"`
	Kelas    string `json:"kelas"`
	// DosenNIDN       string `json:"dosen_nidn"`
	MahasiswaCohort string `json:"mahasiswa_cohort"`
}
