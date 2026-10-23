export default async function handler(req) {
    if (req.method !== 'POST') {
        return new Response('Method Not Allowed', { status: 405 });
    }

    try {
        const { message } = await req.json();
        
        const openAiKey = process.env.OPENAI_API_KEY;
        const exaKey = process.env.EXA_API_KEY;

        if (!openAiKey || !exaKey) {
            return new Response(JSON.stringify({ text: "Error: API keys are missing in Vercel settings." }), { status: 200 });
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

        return new Response(JSON.stringify({ text: replyText }), {
            status: 200,
            headers: { 'Content-Type': 'application/json' }
        });

    } catch (error) {
        return new Response(JSON.stringify({ text: "Sorry, SahihAgent encountered an unexpected server issue." }), { status: 200 });
    }
}


