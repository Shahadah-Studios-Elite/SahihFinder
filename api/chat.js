const openAiKey = process.env.OPENAI_API_KEY;
const exaKey = process.env.EXA_API_KEY;

export async function POST(request) {
    try {
        const { message } = await request.json();

        if (!message) {
            return Response.json({ text: "Please type a question for SahihAgent." }, { status: 200 });
        }

        if (!openAiKey || !exaKey) {
            return Response.json({ text: "Error: API keys are missing in Vercel settings." }, { status: 200 });
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
            return Response.json({ text: openaiData.choices[0].message.content }, { status: 200 });
        }

        return Response.json({ text: "SahihAgent could not compile a clean response." }, { status: 200 });

    } catch (error) {
        return Response.json({ text: `System Error: ${error.message}` }, { status: 200 });
    }
}

