package service

import "database/sql"

// DBFill fills corresponding DB with some data. For testing purposes only.
func DBFill(db *sql.DB) {
	_, _ = db.Exec("INSERT INTO globaltag (category) VALUES ('Спорт'), ('Музыка'), ('Гейминг') ")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Футбол', (SELECT globaltag_id FROM globaltag WHERE category='Спорт'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Баскетбол', (SELECT globaltag_id FROM globaltag WHERE category='Спорт'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Тениис', (SELECT globaltag_id FROM globaltag WHERE category='Спорт'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Пианино', (SELECT globaltag_id FROM globaltag WHERE category='Музыка'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Гитара', (SELECT globaltag_id FROM globaltag WHERE category='Музыка'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Контрабас', (SELECT globaltag_id FROM globaltag WHERE category='Музыка'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('CS: GO', (SELECT globaltag_id FROM globaltag WHERE category='Гейминг'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Minecraft', (SELECT globaltag_id FROM globaltag WHERE category='Гейминг'))")
	_, _ = db.Exec("INSERT INTO tag (activity, globaltag_id) VALUES ('Waaaaar Thundeeer', (SELECT globaltag_id FROM globaltag WHERE category='Гейминг'))")
	_, _ = db.Exec("INSERT INTO team (name, rate, rules, logo, media, reg_date, place, tags[])")
}
