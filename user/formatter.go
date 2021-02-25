package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NamaLengkap string `json:"nama_lengkap"`
	Foto        string `json:"foto"`
}

func FormatUser(user User) UserFormatter {
	userFormat := UserFormatter{}
	userFormat.ID = user.ID
	userFormat.Username = user.Username
	userFormat.Password = user.Password
	userFormat.NamaLengkap = user.NamaLengkap
	userFormat.Foto = user.Foto
	return userFormat
}
