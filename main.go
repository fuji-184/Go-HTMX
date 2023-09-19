

package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "sync"
    "time"
    "os"
    "io"

    _ "github.com/go-sql-driver/mysql"
)

type Item struct {
    Id int `db:"id"`
    Judul string `db:"judul"`
    Isi string `db:"isi"`
    Ktg int `db:"ktg"`
    Tgl time.Time `db:"tgl"`
}

func main() {
    db, err := sql.Open("mysql", "root:1234@/adsense")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Mengatur parameter koneksi database
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)
    
    http.Handle("/gambar/", http.StripPrefix("/gambar/",
    http.FileServer(http.Dir("gambar"))))

    http.HandleFunc("/tes", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            var results []Item
            var wg sync.WaitGroup

            // Menggunakan goroutine untuk mengambil data dari database
            wg.Add(1)
            go func() {
                defer wg.Done()
                stmt, err := db.Prepare("SELECT * FROM post")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                defer stmt.Close()

                rows, err := stmt.Query()
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                defer rows.Close()

                for rows.Next() {
                    var data Item
                    var times string
                    err := rows.Scan(&data.Id, &data.Judul, &data.Isi,
                    &data.Ktg, &times)
                    if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                    }
                    // Mengkonversi string timestampValue menjadi time.Time
    data.Tgl, err = time.Parse("2006-01-02 15:04:05", times)
    if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
    }
                    results = append(results, data)
                }
            }()

            // Menunggu goroutine selesai
            wg.Wait()

            tmpl, err := template.ParseFiles("index.html")
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            err = tmpl.Execute(w, results)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else if r.Method == http.MethodPost {
          
         
          
    no := r.PostFormValue("no")
    


    // Menyimpan data ke database
    stmt, err := db.Prepare("INSERT INTO tes (no) VALUES (?)")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(no) // Menggunakan "no" karena "id" tidak didefinisikan
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    
    
    



// Menerima file dari frontend
    file, header, err := r.FormFile("file") // "gambar" adalah nama field pada form HTML
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Mendapatkan nama file dan ekstensinya
    fileName := header.Filename

    // Membuka file untuk penyimpanan
    newFile, err := os.Create("gambar/" + fileName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer newFile.Close()
    
// Menyalin isi file yang diterima ke file baru
    _, err = io.Copy(newFile, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    
    
    // route closer
}
// end route
        

  
 
    })

    err = http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}


