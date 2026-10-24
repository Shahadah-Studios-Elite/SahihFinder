// here lives the INCASE JS used in wiki.html
const titles = {
            all: "General",
            madhab: "Madhab Section"
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
        }
