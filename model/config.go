package model

import "shuhelperConfigStorage/infrastructure"

type Config struct {
	StudentId  string `json:"student_id"`
	Mode       string `json:"mode"`
	SaveToDoIn string `json:"saveTodoIn"`
}

func Save(config Config) error {
	_, err := infrastructure.DB.Exec(`
	INSERT INTO config(student_id, mode, savetodoin) 
	values ($1,$2,$3)
	on conflict(student_id) DO
	UPDATE SET mode=$2,
	           savetodoin=$3;
	`, config.StudentId, config.Mode, config.SaveToDoIn)
	return err
}

func Get(studentId string) (Config, error) {
	result := Config{StudentId: studentId}
	row := infrastructure.DB.QueryRow(`
	SELECT mode, savetodoin
	FROM config
	WHERE student_id=$1;
	`, studentId)
	err := row.Scan(&result.Mode, &result.SaveToDoIn)
	return result, err
}
