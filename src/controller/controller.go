package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/danieldwipaska/muxserver/src/utils"
	"github.com/gorilla/mux"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("This is home route"))
}

func GetMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movies []utils.Movie
	movies = append(movies, utils.Movie{ID: "1", Isbn: "86753", Title: "Avengers: End Game", Director: &utils.Director{Firstname: "James", Lastname: "Wann"}})
	movies = append(movies, utils.Movie{ID: "2", Isbn: "62354", Title: "Inception", Director: &utils.Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, utils.Movie{ID: "3", Isbn: "74287", Title: "Cinta Brontosaurus", Director: &utils.Director{Firstname: "Raditya", Lastname: "Dika"}})
	json.NewEncoder(res).Encode(movies)
}

func CreateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movies []utils.Movie
	movies = append(movies, utils.Movie{ID: "1", Isbn: "86753", Title: "Avengers: End Game", Director: &utils.Director{Firstname: "James", Lastname: "Wann"}})
	movies = append(movies, utils.Movie{ID: "2", Isbn: "62354", Title: "Inception", Director: &utils.Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, utils.Movie{ID: "3", Isbn: "74287", Title: "Cinta Brontosaurus", Director: &utils.Director{Firstname: "Raditya", Lastname: "Dika"}})

	var movie utils.Movie
	_ = json.NewDecoder(req.Body).Decode(&movie) // assign the req.Body to variable movie
	movie.ID = strconv.Itoa(rand.Intn(100000))

	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movies)
}

func DeleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movies []utils.Movie
	movies = append(movies, utils.Movie{ID: "1", Isbn: "86753", Title: "Avengers: End Game", Director: &utils.Director{Firstname: "James", Lastname: "Wann"}})
	movies = append(movies, utils.Movie{ID: "2", Isbn: "62354", Title: "Inception", Director: &utils.Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, utils.Movie{ID: "3", Isbn: "74287", Title: "Cinta Brontosaurus", Director: &utils.Director{Firstname: "Raditya", Lastname: "Dika"}})

	vars := mux.Vars(req)

	for i, movie := range movies {
		if movie.ID == vars["id"] {
			movies = append(movies[:i], movies[i+1:]...)
		}
	}

	json.NewEncoder(res).Encode(movies)
}

func GetMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movies []utils.Movie
	movies = append(movies, utils.Movie{ID: "1", Isbn: "86753", Title: "Avengers: End Game", Director: &utils.Director{Firstname: "James", Lastname: "Wann"}})
	movies = append(movies, utils.Movie{ID: "2", Isbn: "62354", Title: "Inception", Director: &utils.Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, utils.Movie{ID: "3", Isbn: "74287", Title: "Cinta Brontosaurus", Director: &utils.Director{Firstname: "Raditya", Lastname: "Dika"}})

	vars := mux.Vars(req)
	var result interface{}

	for _, movie := range movies {
		if movie.ID == vars["id"] {
			result = movie
		}
	}

	if result != nil {
		json.NewEncoder(res).Encode(result)
	} else {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(utils.Error{Message: "Error", StatusCode: http.StatusNotFound})
	}

}

func UpdateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movies []utils.Movie
	movies = append(movies, utils.Movie{ID: "1", Isbn: "86753", Title: "Avengers: End Game", Director: &utils.Director{Firstname: "James", Lastname: "Wann"}})
	movies = append(movies, utils.Movie{ID: "2", Isbn: "62354", Title: "Inception", Director: &utils.Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, utils.Movie{ID: "3", Isbn: "74287", Title: "Cinta Brontosaurus", Director: &utils.Director{Firstname: "Raditya", Lastname: "Dika"}})

	vars := mux.Vars(req)

	var result interface{}
	var updatedMovie utils.Movie

	for i, movie := range movies {
		if movie.ID == vars["id"] {
			result = movie
			updatedMovie = movie
			movies = append(movies[:i], movies[i+1:]...)
		}
	}

	// if movie not found
	if result == nil {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(utils.Error{Message: "Error", StatusCode: http.StatusNotFound})
		return
	}

	var movie utils.Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = updatedMovie.ID

	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movies)
}
