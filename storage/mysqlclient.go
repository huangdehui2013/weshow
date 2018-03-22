

var db *sql.DB
 
func init() {
    db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
}


func pool(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM user limit 1")
    defer rows.Close() //Notice ! ! !
    checkErr(err)
 
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for j := range values {
        scanArgs[j] = &values[j]
    }
 
    record := make(map[string]string)
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
    }
 
    fmt.Println(record)
    fmt.Fprintln(w, "finish")
}
 
func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}

