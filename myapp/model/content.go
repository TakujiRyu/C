package model

import "myapp/dataStore/postgres"

type Content struct {
	ContentId string `json:"ctid"`
	Post      string `json:"post"`
}

const (
	queryContentInsert  = "INSERT INTO content(contentid, post) VALUES($1, $2);"
	queryContentGetUser = "SELECT contentid, post FROM content WHERE contentid=$1;"
	queryContentUpdate  = "UPDATE content SET contentid=$1, post=$2 WHERE contentid=$3 RETURNING contentid;"
	queryContentDelete  = "DELETE FROM content WHERE contentid=$1 RETURNING contentid;"
	queryContentSelect  = "SELECT * from content;"
)

func (c *Content) Create() error {
	_, err := postgres.Db.Exec(queryContentInsert, c.ContentId, c.Post)
	return err
}

func (c *Content) Read() error {
	return postgres.Db.QueryRow(queryContentGetUser, c.ContentId).Scan(&c.ContentId, &c.Post)
}

func (c *Content) Update(ctid string) error {
	row := postgres.Db.QueryRow(queryContentUpdate, c.ContentId, c.Post, ctid)
	err := row.Scan(&c.ContentId)
	return err
}

func (c *Content) Delete() error {
	row := postgres.Db.QueryRow(queryContentDelete, c.ContentId)
	err := row.Scan(&c.ContentId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllContent() ([]Content, error) {
	rows, getErr := postgres.Db.Query(queryContentSelect)
	if getErr != nil {
		return nil, getErr
	}

	contents := []Content{}

	for rows.Next() {
		var c Content
		dbErr := rows.Scan(&c.ContentId, &c.Post)
		if dbErr != nil {
			return nil, dbErr
		}

		contents = append(contents, c)
	}
	rows.Close()
	return contents, nil
}
