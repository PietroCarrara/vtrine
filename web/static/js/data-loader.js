/**
 * Load information about movies and store it in the dataset of each element
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a tmdb or imdb property in their dataset
 * @returns {Promise<HTMLElement[]>} Elements that have not been loaded
 */
function load(elements) {
    return loadTMDB(elements);
}

/**
 * Load information about movies and store it in the dataset of each element using
 * The Movie Database
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a tmdb property in their dataset
 *
 * @returns {Promise<HTMLElement[]>} A promise that is done when all the data has been loaded,
 *                                   with all the elements that have not been loaded
 */
function loadTMDB(elements) {
    var promises = [];
    var notLoaded = [];

    for (let i of elements) {
        if (!i.dataset.tmdb || i.dataset.dataLoaded) {
            if (!i.dataset.dataLoaded) {
                notLoaded.push(i);
            }
            continue;
        }

        var prom;

        prom = tmdb.movies
            .getById(i.dataset.tmdb)
            .then(movie => {
                setDataMovieTMDB(i, movie);
                i.dataset.dataLoaded = true;
            })
            .catch(async err => {
                if (err.response.status == 404) {
                    await tmdb.tv
                        .getById(i.dataset.tmdb)
                        .then(series => {
                            setDataSeriesTMDB(i, series);
                            i.dataset.dataLoaded = true;
                        })
                        .catch(() => {
                            delete i.dataset.dataLoaded;
                            notLoaded.push(i);
                        });
                } else {
                    delete i.dataset.dataLoaded;
                    notLoaded.push(i);
                }
            })


        promises.push(prom);
    }

    return Promise.all(promises).then(() => notLoaded);
}

/**
 * Fills a dataset with movie data
 *
 * @param {HTMLElement} element The element to fill the dataset
 * @param {*} movie Object containing movie data
 */
function setDataMovieTMDB(element, movie) {
    var data = element.dataset;

    data.title = movie.title;
    data.overview = movie.overview;
    data.screening = movie.release_date;

    if (movie.poster_path !== null) {
        data.poster = tmdb.common.getImage('w342', movie.poster_path);
    }
}

/**
 * Fill a dataset with series data
 *
 * @param {HTMLElement} element The element to fill the dataset
 * @param {*} series Object containing series data
 */
function setDataSeriesTMDB(element, series) {
    var data = element.dataset;

    data.title = series.name;
    data.overview = series.overview;
    data.screening = series.first_air_date;

    if (series.poster_path !== null) {
        data.poster = tmdb.common.getImage('w342', series.poster_path);
    }
}