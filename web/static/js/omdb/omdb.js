class OMDB {
    constructor(key) {
        this._key = key;
    }

    /**
     * Returns information about the media that is associated with a given id
     *
     * @param {string} id The IMDB id of the media
     * @return {Promise} Promise that resolves to the media data
     */
    getById(id) {
        return this.fetch('/', {
            i: id
        });
    }

    /**
     * Send a request to the OMDB api
     *
     * @param {string} path API path to send request
     * @param {*=} params Parameters send to the GET request
     * @returns {*} The data returned by the API
     */
    fetch(path, params) {

        var options = {
            baseURL: 'http://www.omdbapi.com/',
            params: this.generateParams(params),
        };

        return axios.get(path, options)
        .then(r => r.data);
    }

    /**
     * Adjusts the parameters of the request
     *
     * @param {*=} params Parameters to adjust
     */
    generateParams(params) {
        params = params || {};

        params.apikey = this._key;

        return params;
    }
}