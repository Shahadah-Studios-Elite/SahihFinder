const openAiKey = process.env.OPENAI_API_KEY;
const exaKey = process.env.EXA_API_KEY;

module.exports = async (req, res) => {
    if (req.method !== 'POST') {
        return res.status(405).send('Method Not Allowed');
    }

    try {
        const { message } = req.body || JSON.parse(req.body);

        if (!openAiKey || !exaKey) {
            return res.status(200).json({ text: "Error: API keys are missing in Vercel settings." });
        }

        const exaResponse = await fetch('https://exa.ai', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'x-api-key': exaKey
            },
            body: JSON.stringify({ 
                query: message, 
                numResults: 3,
                includeDomains: ["islamqa.info", "islamweb.net", "sunnah.com", "quran.com"]
            })
        });
        
        const exaData = await exaResponse.json();
        const searchResults = exaData.results || [];
        
        let contextText = "";
        searchResults.forEach(result => {
            contextText += `Source: ${result.url}\nTitle: ${result.title}\n\n`;
        });

        const openaiResponse = await fetch('https://openai.com', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${openAiKey}`
            },
            body: JSON.stringify({
                model: 'gpt-4o-mini',
                messages: [
                    { 
                        role: 'system', 
                        content: 'You are SahihAgent, a trustworthy Sunni Muslim AI assistant. Answer the user based on these trusted search results. Only trust the provided sunni sources.' 
                    },
                    { 
                        role: 'user', 
                        content: `Search Results:\n${contextText}\n\nUser Question: ${message}` 
                    }
                ]
            })
        });

        const openaiData = await openaiResponse.json();
        const replyText = openaiData.choices[0].message.content;

        return res.status(200).json({ text: replyText });

    } catch (error) {
        return res.status(200).json({ text: "Sorry, SahihAgent encountered a small processing issue." });
    }
};

