const openAiKey = process.env.OPENAI_API_KEY;
const exaKey = process.env.EXA_API_KEY;

export default async function handler(req, res) {
    if (req.method !== 'POST') {
        return res.status(405).json({ text: "Method Not Allowed" });
    }

    try {
        let message = "";
        
        if (req.body && req.body.message) {
            message = req.body.message;
        } else {
            const chunks = [];
            for await (const chunk of req) {
                chunks.push(chunk);
            }
            const buffer = Buffer.concat(chunks).toString();
            if (buffer) {
                const parsed = JSON.parse(buffer);
                message = parsed.message;
            }
        }

        if (!message) {
            return res.status(200).json({ text: "Please type a question for SahihAgent." });
        }

        if (!openAiKey || !exaKey) {
            return res.status(200).json({ text: "Error: API keys are missing in Vercel settings." });
        }

        let contextText = "";
        try {
            const exaResponse = await fetch('https://exa.ai', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'x-api-key': exaKey
                },
                body: JSON.stringify({ 
                    query: message, 
                    numResults: 2,
                    includeDomains: ["islamqa.info", "islamweb.net", "sunnah.com", "quran.com"]
                })
            });
            const exaData = await exaResponse.json();
            const searchResults = exaData.results || [];
            searchResults.forEach(result => {
                contextText += `Source: ${result.url}\nTitle: ${result.title}\n\n`;
            });
        } catch (e) {
            contextText = "No direct search results available.";
        }

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
                        content: 'You are SahihAgent, a trustworthy Sunni Muslim AI assistant. Answer the user based on these trusted search results. Only trust mainstream sunni sources.' 
                    },
                    { 
                        role: 'user', 
                        content: `Search Results:\n${contextText}\n\nUser Question: ${message}` 
                    }
                ]
            })
        });

        const openaiData = await openaiResponse.json();
        
        if (openaiData.choices && openaiData.choices[0] && openaiData.choices[0].message) {
            return res.status(200).json({ text: openaiData.choices[0].message.content });
        }

        return res.status(200).json({ text: "SahihAgent could not compile a clean response." });

    } catch (error) {
        return res.status(200).json({ text: `System Error: ${error.message}` });
    }
}
