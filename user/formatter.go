package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NamaLengkap string `json:"nama_lengkap"`
	Foto        string `json:"foto"`
	Token       string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	userFormat := UserFormatter{}
	userFormat.ID = user.ID
	userFormat.Username = user.Username
	userFormat.Password = user.Password
	userFormat.NamaLengkap = user.NamaLengkap
	userFormat.Foto = user.Foto
	userFormat.Token = token
	return userFormat
}

type ProfileFormatter struct {
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	Foto        string `json:"foto"`
}

func FormatProfile(user User) ProfileFormatter {
	profileFormat := ProfileFormatter{}
	profileFormat.Username = user.Username
	profileFormat.NamaLengkap = user.NamaLengkap
	profileFormat.Foto = user.Foto
	return profileFormat
}
