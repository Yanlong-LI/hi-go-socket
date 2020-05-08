package db

import "errors"

func (_deleteBuilder *deleteBuilder) Delete() (uint64, error) {

	if len(_deleteBuilder.builder.getWhere()) == 0 {
		return 0, errors.New("删除条件不能为空")
	}

	result, err := db.Exec(_deleteBuilder.Sql(), _deleteBuilder.builder.args...)
	if err != nil {
		return 0, err
	}
	ra, err := result.RowsAffected()
	return uint64(ra), err
}
