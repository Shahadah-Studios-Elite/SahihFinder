let count = 0;
let target = 33;
let rounds = 0;
let currentPhrase = 'SubhanAllah';

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

function saveState() {
    localStorage.setItem('sahihfinder_tasbih', JSON.stringify({
        count,
        target,
        rounds,
        currentPhrase
    }));
}

function updateDisplay() {
    document.getElementById('count-value').textContent = count;
    document.getElementById('phrase-title').textContent = currentPhrase;
    document.getElementById('target-text').textContent = 'Target: ' + target + ' | Completed Rounds: ' + rounds;
    saveState();
}

function incrementCount() {
    count++;
    if (count >= target) {
        rounds++;
        count = 0;
        if (navigator.vibrate) {
            navigator.vibrate(120);
        }
    }
    updateDisplay();
}

function decrementCount() {
    if (count > 0) {
        count--;
        updateDisplay();
    }
}

function resetCount() {
    count = 0;
    rounds = 0;
    updateDisplay();
}

function setPhrase(phrase, btn) {
    currentPhrase = phrase;
    document.querySelectorAll('.phrase-btn').forEach(b => b.classList.remove('active'));
    if (btn) btn.classList.add('active');
    updateDisplay();
}

function setTarget(newTarget, btn) {
    target = newTarget;
    count = 0;
    document.querySelectorAll('.target-btn').forEach(b => b.classList.remove('active'));
    if (btn) btn.classList.add('active');
    updateDisplay();
}

document.addEventListener('DOMContentLoaded', () => {
    const saved = localStorage.getItem('sahihfinder_tasbih');
    if (saved) {
        try {
            const data = JSON.parse(saved);
            if (typeof data.count === 'number') count = data.count;
            if (typeof data.target === 'number') target = data.target;
            if (typeof data.rounds === 'number') rounds = data.rounds;
            if (typeof data.currentPhrase === 'string') currentPhrase = data.currentPhrase;

            document.querySelectorAll('.phrase-btn').forEach(b => {
                b.classList.toggle('active', b.textContent === currentPhrase);
            });
            document.querySelectorAll('.target-btn').forEach(b => {
                b.classList.toggle('active', Number(b.textContent) === target);
            });
        } catch (e) {}
    }
    updateDisplay();
});

document.addEventListener('keydown', (e) => {
    if (e.code === 'Space') {
        e.preventDefault();
        incrementCount();
    }
});
