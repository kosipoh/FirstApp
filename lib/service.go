package lib

import mod "FirstApp/model"

func CheckUser(chat_id int64) {
	db := Con()
	defer ConClose(db)
	req := `INSERT INTO public.users (
		ph, chat_id, fio)
		VALUES ($1, $2, $3)`

	mod := mod.CLient{Chat: 4544545545455, Ph: "78787878", Fio: "HUI"}

	InsDb(db, req, mod)

}
