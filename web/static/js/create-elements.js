/**
 * Creates an element and fills it's dataset
 *
 * @param {string} html The HTML of the element to be created
 * @param {*} movie The movie data, to be saved into the element dataset
 * @returns {Element} The created element
 */
async function createMovieTMDB(html, movie) {
    var element = document.createElement('div');
    element.innerHTML = html;
    element = element.firstElementChild;

    if (!movie.imdb_id) {
        var ids = await tmdb.movies.getExternalIds(movie.id);
        movie.imdb_id = ids.imdb_id;
    }

    setDataMovieTMDB(element, movie);

    return element;
}

/**
 * Creates an element and fills it's dataset
 *
 * @param {string} html The HTML of the element to be created
 * @param {*} show The show data, to be saved into the element dataset
 * @returns {Element} The created element
 */
async function createShowTMDB(html, show) {
    var element = document.createElement('div');
    element.innerHTML = html;
    element = element.firstElementChild;

    if (!show.imdb_id) {
        var ids = await tmdb.tv.getExternalIds(show.id);
        show.imdb_id = ids.imdb_id;
    }

    setDataShowTMDB(element, show);

    return element;
}