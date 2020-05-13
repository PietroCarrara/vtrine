(() => {
    var thumbElements = document.querySelectorAll('.movie-thumb');

    load(thumbElements)
        .then(() => {
            for (var i of thumbElements) {
                var poster = i.querySelector('.movie-poster');
                var noPoster = i.querySelector('.movie-no-poster');
                var title = i.querySelector('.movie-title');
                var screening = i.querySelector('.movie-screening');
                var link = i.querySelector('.movie-link');

                if (i.dataset.title) {
                    title.textContent = i.dataset.title;
                } else {
                    title.textContent = 'Unknown';
                }

                if (i.dataset.screening) {
                    screening.textContent = new Date(i.dataset.screening).getFullYear();
                }

                if (i.dataset.poster) {
                    poster.style.backgroundImage = `url(${i.dataset.poster})`;
                } else {
                    noPoster.classList.remove('invisible');
                }

                if (i.dataset.imdb) {
                    link.href = `/media/${i.dataset.imdb}`;
                }
            }
        });
})()