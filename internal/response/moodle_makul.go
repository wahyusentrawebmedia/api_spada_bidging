package response

type MoodleMakulRequest struct {
	IDNumberCourse string `json:"idnumber_course"`
	KodeMK         string `json:"kode_mk"`
	NamaMK         string `json:"nama_mk"`
	Tahun          string `json:"tahun"`
	Semester       string `json:"semester"`
	Kelas          string `json:"kelas"`
	// DosenNIDN       string `json:"dosen_nidn"`
	MahasiswaCohort string `json:"mahasiswa_cohort"`

	ShortNameMakul string `json:"short_name_makul"`

	IdKelasMakul  string `json:"id_kelas_makul"`
	IdKelasKuliah string `json:"id_kelas_kuliah"`
	KodeKelas     string `json:"kode_kelas"`
}
