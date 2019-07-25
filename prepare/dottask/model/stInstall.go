package model

import "fmt"

type StInstall struct {
	BaseMode
}

type hourInstallResult struct {
	server_id  int
	pfrom_id   int
	inst_date  string
	inst_hour  int
	inst_count int
	via        string
}

func HourInstallStatis(table string) interface{} {
	var result []hourInstallResult
	sql := "SELECT pf AS pfrom_id,server_id, DATE(regdate) AS inst_date, HOUR (regdate) AS inst_hour, count(*) AS inst_count,via  FROM  " + table + " WHERE isnewaccount=? GROUP BY pfrom_id,server_id, inst_date, inst_hour, via"

	rows, _ := db_role.Raw(sql, 1).Rows()
	defer rows.Close()

	fmt.Println(rows)
	for rows.Next() {
		var tm hourInstallResult
		rows.Scan(
			&tm.server_id,
			&tm.pfrom_id,
			&tm.inst_date,
			&tm.inst_hour,
			&tm.inst_count,
			&tm.via)

		result = append(result, tm)
	}
	return result
}
