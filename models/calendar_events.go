package models

type CalendarEventAPI struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	DateStart   string `json:"date_start"`
	DateEnd     string `json:"date_end"`
	Description string `json:"description"`
	Location    string `json:"location"`
	SourceId    int    `json:"source_id"`
	EntityType  string `json:"entity_type"`
	CreatedBy   int    `json:"created_by"`
	ModifiedBy  int    `json:"modified_by"`
	DateCreate  string `json:"date_create"`
	DateUpdate  string `json:"date_update"`
}

type CalendarEventDB struct {
	Nested
	Id          int    `json:"id"`
	Title       string `json:"title"`
	DateStart   int64  `json:"dateStart"`
	DateEnd     int64  `json:"dateEnd"`
	Description string `json:"description"`
	Location    string `json:"location"`
	SourceId    int    `json:"sourceId"`
	EntityType  string `json:"entityType"`
	CreatedBy   int    `json:"createdBy"`
	ModifiedBy  int    `json:"modifiedBy"`
	DateCreate  int64  `json:"dateCreate"`
	DateUpdate  int64  `json:"dateUpdate"`
}

/*func (u CalendarEventDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u CalendarEventDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

func (u *CalendarEventDB) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}*/

type ListOfCalendarEventDB []*CalendarEventDB

/*func (u ListOfCalendarEventDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u ListOfCalendarEventDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

func (u *ListOfCalendarEventDB) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &u)
}*/
