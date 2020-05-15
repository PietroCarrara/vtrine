/**
 * Using the data present in the datasets, applies them to the layout
 *
 * @param {HTMLElement[]|NodeList} elements The elements that contain data
 */
function applyData(elements) {
    for (var i of elements) {
        var poster = i.querySelector('.movie-poster');
        var noPoster = i.querySelector('.movie-no-poster');
        var title = i.querySelector('.movie-title');
        var screening = i.querySelector('.movie-screening');
        var link = i.querySelector('.movie-link');
        var overview = i.querySelector('.movie-overview');
        var youtubeButton = i.querySelector('.movie-trailer-youtube-button');
        var youtubeIframe = i.querySelector('.movie-trailer-youtube-iframe');

        if (title) {
            if (i.dataset.title) {
                title.textContent = i.dataset.title;
            } else {
                title.textContent = 'Unknown';
            }
        }

        if (screening && i.dataset.screening) {
            screening.textContent = new Date(i.dataset.screening).getFullYear();
        }

        if (poster && i.dataset.poster) {
            poster.src = i.dataset.poster;
        } else if (noPoster) {
            noPoster.classList.remove('invisible');
            if (poster) {
                poster.classList.add('invisible');
            }
        }

        if (link && i.dataset.imdb) {
            link.href = `/media/${i.dataset.imdb}`;
        }

        if (overview) {
            if (i.dataset.overview) {
                overview.textContent = i.dataset.overview;
            } else {
                overview.textContent = 'No overview';
            }
        }

        if (youtubeIframe && i.dataset.youtube) {
            youtubeIframe.src = `https://www.youtube.com/embed/${i.dataset.youtube}`;
            if (youtubeButton) {
                youtubeButton.classList.remove('invisible');
            }
        }
    }
}