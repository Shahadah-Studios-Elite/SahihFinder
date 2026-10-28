// here lives the INCASE JS used in wiki.html

        let currentTab = 'all';

        const titles = {
            all: "General",
            madhab: "Madhab Section"
        };
     const contents = {
            all: `
                
                <details style="text-align: left; margin: 0 auto; max-width: 450px; cursor: pointer;">
                    <summary style="font-weight: bold; color: #054432;">Verified resources</summary>
                    <ul style="line-height: 1.8; padding-left: 20px; margin-top: 10px;">
                        <li><a href="https://sunnah.com" target="_blank" style="color: #c28100; text-decoration: none; font-weight: 500;">Sunnah.com</a><br>A website where verified <br>Sahih Hadiths are listed.</li>
                        <li><a href="https://islamqa.info" target="_blank" style="color: #c28100; text-decoration: none; font-weight: 500;">IslamQA.info</a><br>A website where verified Islamic Shaikhs answer <br>your questions related to Deen and all <br>answers are verified.</li>
                        <li><a href="https://quran.com" target="_blank" style="color: #c28100; text-decoration: none; font-weight: 500;">Quran.com</a><br>A Quran website which has a Quran app-store <br>and many other Islamic items.</li>
                    </ul>
                </details>
            `,
            madhab: `
                <p style="color: #054432; font-weight: bold; margin-bottom: 15px;">Explore your Madhab!</p>
                <ul style="text-align: left; margin: 0 auto; max-width: 450px; line-height: 1.8; padding-left: 20px;">
                    <li><strong style="color: #c28100;">Hanafi</strong></li>
                    <li><strong style="color: #c28100;">Maliki</strong></li>
                    <li><strong style="color: #c28100;">Shafi'i</strong></li>
                    <li><strong style="color: #c28100;">Hanbali</strong></li>
                </ul>
            `
        };


        function switchTab(key) {
            currentTab = key;
            document.getElementById('search-input').value = '';

            document.querySelectorAll('.tab-button').forEach(btn => {
                btn.classList.remove('active');
            });

            if (key === 'all') document.getElementById('tab-all').classList.add('active');
            if (key === 'madhab') document.getElementById('tab-madhab').classList.add('active');

            const titleElem = document.getElementById('section-title');
            if (titleElem && titles[key]) {
                titleElem.textContent = titles[key];
            }

            const contentElem = document.getElementById('section-content');
            if (contentElem && contents[key]) {
                contentElem.innerHTML = contents[key];
            }
            
            document.getElementById('content-box').style.display = 'block';
        }

        function filterContent() {
            const query = document.getElementById('search-input').value.toLowerCase();
            const contentBox = document.getElementById('content-box');
            const originalText = contents[currentTab].toLowerCase();

            if (originalText.includes(query)) {
                contentBox.style.display = 'block';
            } else {
                contentBox.style.display = 'none';
            }
        }

        function toggleTheme() {
            const body = document.body;
            const btn = document.getElementById('theme-btn');
            
            if (body.getAttribute('data-theme') === 'dark') {
                body.removeAttribute('data-theme');
                btn.textContent = '🌙 Dark';
            } else {
                body.setAttribute('data-theme', 'dark');
                btn.textContent = '☀️ Light';
            }
        }
