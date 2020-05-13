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
            notLoaded.push(i);
            continue;
        }

        var prom = tmdb.movies
            .getById(i.dataset.tmdb)
            .then((movie) => {
                setDataTMDB(i, movie);
                i.dataset.dataLoaded = true;
            })
            .catch(() => i.dataset.dataLoaded = false);

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
function setDataTMDB(element, movie) {
    var data = element.dataset;

    data.title = movie.title;
    data.overview = movie.overview;
    data.screening = movie.release_date;

    if (movie.poster_path !== null) {
        data.poster = tmdb.common.getImage('w342', movie.poster_path);
    }
}