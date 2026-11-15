// took a while to debug...
function toggleTheme() {
    const body = document.body;
    const icon = document.getElementById('theme-icon');
    if (body.getAttribute('data-theme') === 'dark') {
        body.removeAttribute('data-theme');
        icon.innerHTML = '<path d="M12 3c-4.97 0-9 4.03-9 9s4.03 9 9 9 9-4.03 9-9c0-.46-.04-.92-.1-1.36-.98 1.37-2.58 2.26-4.4 2.26-2.98 0-5.4-2.42-5.4-5.4 0-1.81.89-3.42 2.26-4.4-.44-.06-.9-.1-1.36-.1z"/>';
    } else {
        body.setAttribute('data-theme', 'dark');
        icon.innerHTML = '<path d="M12 7c-2.76 0-5 2.24-5 5s2.24 5 5 5 5-2.24 5-5-2.24-5-5-5zM2 13h2c.55 0 1-.45 1-1s-.45-1-1-1H2c-.55 0-1 .45-1 1s.45 1 1 1zm18 0h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1zM11 2v2c0 .55.45 1 1 1s1-.45 1-1V2c0-.55-.45-1-1-1s-1 .45-1 1zm0 18v2c0 .55.45 1 1 1s1-.45 1-1v-2c0-.55-.45-1-1-1s-1 .45-1 1zM5.99 4.58c-.39-.39-1.03-.39-1.41 0s-.39 1.03 0 1.42l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.42L5.99 4.58zm12.37 12.37c-.39-.39-1.03-.39-1.41 0s-.39 1.03 0 1.42l1.06 1.06c.39.39 1.03.39 1.41 0s.39-1.03 0-1.42l-1.06-1.06zm1.06-12.37c-.39-.39-1.03-.39-1.41 0l-1.06 1.06c-.39.39-.39 1.03 0 1.42s1.03.39 1.41 0l1.06-1.06c.39-.38.39-1.02 0-1.42zM5.99 16.95l-1.06 1.06c-.39.39-.39 1.03 0 1.42s1.03.39 1.41 0l1.06-1.06c.39-.39.39-1.03 0-1.42z"/>';
    }
}

function buildHadithLink() {
    const select = document.getElementById('book-select');
    const selectedOption = select.options[select.selectedIndex];
    const slug = select.value;
    const maxNum = Number(selectedOption.getAttribute('data-max'));
    const bookName = selectedOption.getAttribute('data-name');
    const num = Number(document.getElementById('hadith-num').value);
    const panel = document.getElementById('result-panel');

    panel.innerHTML = '';
    panel.style.display = 'block';

    if (!num || num < 1 || !Number.isInteger(num)) {
        const err = document.createElement('p');
        err.className = 'error-msg';
        err.textContent = 'Please enter a valid positive Hadith number.';
        panel.appendChild(err);
        return;
    }

    if (num > maxNum) {
        const err = document.createElement('p');
        err.className = 'error-msg';
        err.textContent = bookName + ' only goes up to Hadith #' + maxNum + '.';
        panel.appendChild(err);
        return;
    }

    const targetUrl = new URL('/' + slug + ':' + num, 'https://sunnah.com');

    const title = document.createElement('p');
    title.className = 'result-title';
    title.textContent = bookName + ' #' + num;

    const link = document.createElement('a');
    link.className = 'sunnah-link';
    link.href = targetUrl.href;
    link.target = '_blank';
    link.textContent = 'Open Verified Narration on Sunnah.com ->';

    panel.appendChild(title);
    panel.appendChild(link);
}
