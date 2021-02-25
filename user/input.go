package user

type RegisterInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
	Foto        string `json:"foto"`
}
type UpdateInput struct {
	Password    string `json:"password"`
	NamaLengkap string `json:"nama_lengkap"`
	Foto        string `json:"foto"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
