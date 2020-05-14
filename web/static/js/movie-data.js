(() => {
    var thumbElements = document.querySelectorAll('.movie-data');

    load(thumbElements)
        .then(() => {
            for (var i of thumbElements) {
                var poster = i.querySelector('.movie-poster');
                var noPoster = i.querySelector('.movie-no-poster');
                var title = i.querySelector('.movie-title');
                var screening = i.querySelector('.movie-screening');
                var link = i.querySelector('.movie-link');
                var overview = i.querySelector('.movie-overview');

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
                    poster.style.backgroundImage = `url(${i.dataset.poster})`;
                } else if (noPoster) {
                    noPoster.classList.remove('invisible');
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
            }
        });
})()