/**
 * Load basic information about movies and store it in the dataset of each element
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a imdb property in their dataset
 * @returns {Promise<HTMLElement[]>} Elements that have been loaded
 */
function loadBasic(elements) {
    return loadBasicTMDB(elements)
        // Load elements that have not been loaded
        .then(async loaded => {
            var imdbLoaded = await loadBasicIMDB(loaded.filter(x => !loaded.includes(x)));
            return imdbLoaded.concat(...loaded);
        });
}

/**
 * Load trailer information about movies and store it in the dataset of each element
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a imdb property in their dataset
 * @returns {Promise<HTMLElement[]>} Elements that have been loaded
 */
function loadVideo(elements) {
    return loadVideoTMDB(elements);
}

/**
 * Load information about movies and store it in the dataset of each element using
 * The Movie Database
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a imdb property in their dataset
 *
 * @returns {Promise<HTMLElement[]>} A promise that is done when all the data has been loaded,
 *                                   with all the elements that have been loaded
 */
function loadBasicTMDB(elements) {
    var promises = [];

    for (let i of elements) {
        if (!i.dataset.imdb) {
            continue;
        }

        var prom = tmdb.find
        .getById(i.dataset.imdb, 'imdb_id')
        .then(res => {
            if (res.movie_results.length > 0) {
                setDataMovieTMDB(i, res.movie_results[0]);
                return i;
            } else if (res.tv_results.length > 0) {
                setDataShowTMDB(i, res.tv_results[0]);
                return i;
            }
        });

        promises.push(prom);
    }

    return Promise.all(promises);
}

/**
 * Load trailers and store their information on the dataset of each element using
 * The Movie Database
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a imdb property in their dataset
 *
 * @returns {Promise<HTMLElement[]>} A promise that is done when all the data has been loaded,
 *                                   with all the elements that have been loaded
 */
async function loadVideoTMDB(elements) {
    var promises = [];
    var loaded = [];

    for (let i of elements) {

        var res = await tmdb.find.getById(i.dataset.imdb, 'imdb_id');
        var loadFunc = null;

        var type = i.dataset.type;
        if (!type) {
            if (res.movie_results.length > 0) type = 'movie';
            if (res.tv_results.length > 0) type = 'show';
        }

        if (i.dataset.type === 'movie' && res.movie_results.length > 0) {
            loadFunc = () => tmdb.movies.getVideos(res.movie_results[0].id);
        } else if (type === 'show' && res.tv_results.length > 0) {
            loadFunc = () => tmdb.tv.getVideos(res.tv_results[0].id);
        }

        if (loadFunc == null) {
            setVideoTMDB(i, []);
            continue;
        }

        loaded.push(i);
        promises.push(
            loadFunc()
            .then(data => setVideoTMDB(i, data.results))
        );
    }

    return Promise.all(promises).then(() => loaded);
}

/**
 * Load information about movies and store it in the dataset of each element using
 * the Internet Movie Database
 *
 * @param {NodeList|HTMLElement[]} elements Elements that have a imdb property in their dataset
 *
 * @returns {Promise<HTMLElement[]>} A promise that is done when all the data has been loaded,
 *                                   resolving with all the elements that have been loaded
 */
function loadBasicIMDB(elements) {
    var promises = [];
    var loaded = [];

    for (let i of elements) {
        if (!i.dataset.imdb || i.dataset.dataLoaded) {
            continue;
        }

        var prom = omdb.getById(i.dataset.imdb)
            .then(data => {
                if (data.Error) {
                    return
                }

                loaded.push(i);
                setDataIMDB(i, data);
            });

        promises.push(prom);
    }

    return Promise.all(promises).then(() => loaded);
}

/**
 * Fills a dataset with movie data
 *
 * @param {HTMLElement} element The element to fill the dataset
 * @param {*} movie Object containing movie data
 */
function setDataMovieTMDB(element, movie) {
    var data = element.dataset;

    data.tmdb = movie.id;
    data.dataLoaded = true;
    data.title = movie.title;
    data.overview = movie.overview;
    data.screening = movie.release_date;
    data.type = 'movie';

    if (movie.poster_path !== null) {
        data.poster = tmdb.common.getImage('w342', movie.poster_path);
    } else {
        data.poster = '';
    }

    if (movie.imdb_id) {
        data.imdb = movie.imdb_id;
    }
}

/**
 * Fill a dataset with TV Show data
 *
 * @param {HTMLElement} element The element to fill the dataset
 * @param {*} series Object containing series data
 */
function setDataShowTMDB(element, series) {
    var data = element.dataset;

    data.tmdb = series.id;
    data.dataLoaded = true;
    data.title = series.name;
    data.overview = series.overview;
    data.screening = series.first_air_date;
    data.type = 'show';

    if (series.poster_path !== null) {
        data.poster = tmdb.common.getImage('w342', series.poster_path);
    } else {
        data.poster = '';
    }

    if (series.imdb_id) {
        data.imdb = series.imdb_id;
    }
}

/**
 * Fill an element's dataset with video information
 *
 * @param {HTMLElement} element The element to fill the dataset
 * @param {*} videos Array of videos
 */
function setVideoTMDB(element, videos) {
    var data = element.dataset;

    var trailers = videos.filter(x => x.type.toLowerCase() === 'trailer');
    var teasers = videos.filter(x => x.type.toLowerCase() === 'teaser');

    var youtube = trailers.find(x => x.site.toLowerCase() === 'youtube') ||
                  teasers.find(x => x.site.toLowerCase() === 'youtube');

    if (youtube) {
        data.youtube = youtube.key;
    } else {
        data.youtube = '';
    }
}

/**
 * Fill a dataset with data comming from IMDB
 *
 * @param {HTMLElement} element The element to fill the dataset
 * @param {*} media The media data
 */
function setDataIMDB(element, media) {
    var data = element.dataset;

    data.dataLoaded = true;
    data.title = media.Title;
    data.overview = media.Plot;
    data.screening = media.Released;
    data.poster = media.Poster;
}