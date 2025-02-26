package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const PORT = ":8080"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := "thanks for visiting!\n"
	fmt.Fprintf(w, "%s", body)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found: %s", r.URL.Path)
		return
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)
	phone := paramStr[2]
	err := deleteEntry(phone)

	if err != nil {
		fmt.Println(err)
		body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", body)
		return
	}

	body := phone + " deleted!\n"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", body)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := list()
	fmt.Fprintf(w, "%s", body)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := fmt.Sprintf("Total entries: %d\n", len(data))
	fmt.Fprintf(w, "%s", body)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 5 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not enough arguments: %s", r.URL.Path)
		return
	}

	name := paramStr[2]
	surname := paramStr[3]
	phone := paramStr[4]

	p := strings.ReplaceAll(phone, "-", "")

	if !matchTel(p) {
		fmt.Println("Not a valid phone number:", p, " | ", phone)
		return
	}

	temp := &Entry{Name: name, Surname: surname, Tel: p}
	err := insert(temp)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		body := "failed to add record\n"
		fmt.Fprintf(w, "%s", body)
	} else {
		log.Println("Serving:", r.URL.Path, "from", r.Host)
		body := "new record added successfully\n"
		fmt.Fprintf(w, "%s", body)
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found: %s", r.URL.Path)
		return
	}

	var body string
	phone := paramStr[2]

	p := search(phone)

	if p == nil {
		w.WriteHeader(http.StatusNotFound)
		body = fmt.Sprintf("could not be found: %s\n", phone)
	} else {
		w.WriteHeader(http.StatusOK)
		body = fmt.Sprintf("%s %s %s \n", p.Name, p.Surname, p.Tel)
	}

	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	fmt.Fprintf(w, "%s", body)
}

