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
        if (!i.dataset.imdb || i.dataset.dataLoaded) {
            if (!i.dataset.dataLoaded) {
                notLoaded.push(i);
            }
            continue;
        }

        var prom = tmdb.find
        .getById(i.dataset.imdb, 'imdb_id')
        .then(res => {
            if (res.movie_results.length > 0) {
                setDataMovieTMDB(i, res.movie_results[0]);
            } else if (res.tv_results.length > 0) {
                setDataSeriesTMDB(i, res.tv_results[0]);
            }
        });

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

    data.dataLoaded = true;
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

    data.dataLoaded = true;
    data.title = series.name;
    data.overview = series.overview;
    data.screening = series.first_air_date;

    if (series.poster_path !== null) {
        data.poster = tmdb.common.getImage('w342', series.poster_path);
    }
}