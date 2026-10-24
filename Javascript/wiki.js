// here lives the INCASE JS used in wiki.html
const titles = {
            all: "General",
            madhab: "Madhab Section"
        };

        

                const contents = {
           // here to write HTML we may need `` to run it.
                    all: "We're Researching to get the get the best for you! Please check back SOON!",
            madhab: "We're researching the best on the Madhab topic for you! Please check back Soon!"
        };

        function switchTab(key) {
            document.querySelectorAll('.tab-button').forEach(btn => {
                btn.classList.remove('active');
            });

            if (event && event.currentTarget) {
                event.currentTarget.classList.add('active');
            }

            const titleElem = document.getElementById('section-title');
            if (titleElem && titles[key]) {
                titleElem.textContent = titles[key];
            }

            const contentElem = document.getElementById('section-content');
            if (contentElem && contents[key]) {
                contentElem.innerHTML = contents[key];
            }
        }
