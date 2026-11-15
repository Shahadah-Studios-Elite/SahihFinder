// JAVASCRIPT!@
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

function appendMessage(sender, text, type) {
    const chatBox = document.getElementById('chat-box');
    const msgDiv = document.createElement('div');
    msgDiv.className = 'msg ' + type;

    const label = document.createElement('span');
    label.className = 'msg-sender';
    label.textContent = sender;

    const body = document.createElement('div');
    body.textContent = text;

    msgDiv.appendChild(label);
    msgDiv.appendChild(body);
    chatBox.appendChild(msgDiv);
    chatBox.scrollTop = chatBox.scrollHeight;
    return msgDiv;
}

function getFallbackReply(userMsg) {
    const lower = userMsg.toLowerCase();
    if (lower.includes('salam') || lower.includes('hello') || lower.includes('hi')) {
        return 'Wa Alaykum Assalam! How can I help you verify Islamic information or check trusted Sunni references today?';
    }
    if (lower.includes('hadith') || lower.includes('bukhari') || lower.includes('muslim')) {
        return 'When verifying any Hadith, always check its exact wording and grading on Sunnah.com (such as Sahih al-Bukhari or Sahih Muslim) rather than relying on unverified social media posts.';
    }
    return 'Please wait 30 seconds between live AI questions, or use our Searcher and Wiki tools to check verified Sunni references directly on Sunnah.com.';
}

async function sendMessage() {
    const input = document.getElementById('user-input');
    const sendBtn = document.getElementById('send-btn');
    const message = input.value.trim();
    if (!message) return;

    appendMessage('You', message, 'user');
    input.value = '';
    input.disabled = true;
    sendBtn.disabled = true;

    const thinkingMsg = appendMessage('SahihAgent', 'Thinking...', 'agent');

    try {
        const response = await fetch('https://pollinations.ai', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                messages: [
                    {
                        role: 'system',
                        content: 'You are SahihAgent, a respectful, helpful Sunni Muslim AI assistant for SahihFinder. Keep answers clear, concise, and grounded in mainstream Sunni sources like the Quran, Sahih al-Bukhari, and Sahih Muslim. Always remind users to verify exact Hadith wording on Sunnah.com or with a trusted local scholar.'
                    },
                    {
                        role: 'user',
                        content: message
                    }
                ]
            })
        });

        thinkingMsg.remove();

        if (response.status === 402 || !response.ok) {
            appendMessage('SahihAgent', getFallbackReply(message), 'agent');
        } else {
            const data = await response.json();
            const replyText = data && data.choices && data.choices.message
                ? data.choices.message.content.trim()
                : getFallbackReply(message);
            appendMessage('SahihAgent', replyText, 'agent');
        }
    } catch (err) {
        thinkingMsg.remove();
        appendMessage('SahihAgent', getFallbackReply(message), 'agent');
    } finally {
        input.disabled = false;
        sendBtn.disabled = false;
        input.focus();
    }
}
