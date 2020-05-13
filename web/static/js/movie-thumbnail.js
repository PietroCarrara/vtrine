(() => {
    var thumbElements = document.querySelectorAll('.movie-thumb');

    load(thumbElements)
        .then(() => {
            for (var i of thumbElements) {

                if (i.dataset.dataLoaded) {
                    var poster = i.querySelector('.movie-poster');
                    var title = i.querySelector('.movie-title');
                    var screening = i.querySelector('.movie-screening');

                    title.textContent = i.dataset.title;

                    if (i.dataset.screening) {
                        screening.textContent = new Date(i.dataset.screening).getFullYear();
                    }

                    if (i.dataset.poster) {
                        poster.style.backgroundImage = `url(${i.dataset.poster})`;
                    }
                }
            }
        });
})()