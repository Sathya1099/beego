package models

type Object struct {
	ObjectId   string `json:"object_id" orm:"pk"`
	Score      int    `json:"score" valid:"MaxSize(500)" orm:"column(score)"`
	PlayerName string `json:"player_name" valid:"MinSize(3)" orm:"column(playername)"`
}

func Insert(object *Object) error {
	_, err := ormer.Insert(object)
	return err
}

func InsertMulti(objects []Object) error {
	_, err := ormer.InsertMulti(100, objects)
	return err
}

func Read(objId string) (*Object, error) {
	b := &Object{ObjectId: objId}
	err := ormer.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ReadAll() ([]Object, error) {
	qs := ormer.QueryTable("object")
	objects := []Object{}
	_, err := qs.All(&objects)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func Update(object *Object) error {
	var fields []string
	if object.Score != 0 {
		fields = append(fields, "Score")
	}
	if object.PlayerName != "" {
		fields = append(fields, "PlayerName")
	}
	_, err := ormer.Update(&object, fields...)
	if err != nil {
		return err
	}
	return nil
}

func Delete(objId string) error {
	object := Object{ObjectId: objId}
	_, err := ormer.Delete(&object)
	if err != nil {
		return err
	}
	return nil
}
