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

function gregorianToHijri(year, month, day) {
    let y = year;
    let m = month;
    const d = day;

    if (m < 3) {
        y--;
        m += 12;
    }

    const a = Math.floor(y / 100);
    const b = 2 - a + Math.floor(a / 4);
    const jd = Math.floor(365.25 * (y + 4716)) + Math.floor(30.6001 * (m + 1)) + d + b - 1524;

    let l = jd - 1948440 + 10632;
    const n = Math.floor((l - 1) / 10631);
    l = l - 10631 * n + 354;
    const j = Math.floor((10985 - l) / 5316) * Math.floor((50 * l) / 17719) + Math.floor(l / 5670) * Math.floor((43 * l) / 15238);
    l = l - Math.floor((30 - j) / 15) * Math.floor((17719 * j) / 50) - Math.floor(j / 16) * Math.floor((15238 * j) / 43) + 29;

    const hMonth = Math.floor((24 * l) / 709);
    const hDay = l - Math.floor((709 * hMonth) / 24);
    const hYear = 30 * n + j - 30;

    return { hYear, hMonth, hDay };
}

function convertDate() {
    const dateInput = document.getElementById('date-input').value;
    const resultBox = document.getElementById('result-box');
    const hijriElem = document.getElementById('hijri-date');
    const gregElem = document.getElementById('gregorian-label');

    if (!dateInput) return;

    const parts = dateInput.split('-').map(Number);
    const year = parts;
    const month = parts;
    const day = parts;

    const months = [
        "",
        "Muharram",
        "Safar",
        "Rabi' al-Awwal",
        "Rabi' al-Thani",
        "Jumada al-Awwal",
        "Jumada al-Thani",
        "Rajab",
        "Sha'ban",
        "Ramadan",
        "Shawwal",
        "Dhu al-Qi'dah",
        "Dhu al-Hijjah"
    ];

    const gregMonths = [
        "January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"
    ];

    const { hYear, hMonth, hDay } = gregorianToHijri(year, month, day);
    hijriElem.textContent = `${hDay} ${months[hMonth]} ${hYear} AH`;
    gregElem.textContent = `Gregorian: ${gregMonths[month - 1]} ${day}, ${year}`;
    resultBox.style.display = 'block';
}

document.addEventListener('DOMContentLoaded', () => {
    const today = new Date();
    const yyyy = today.getFullYear();
    const mm = String(today.getMonth() + 1).padStart(2, '0');
    const dd = String(today.getDate()).padStart(2, '0');
    document.getElementById('date-input').value = `${yyyy}-${mm}-${dd}`;
    convertDate();
});

