import { generateText, tool } from 'ai';
import { openai } from '@ai-sdk/openai';
import { z } from 'zod';

export default async function handler(req) {
    if (req.method !== 'POST') {
        return new Response('Method Not Allowed', { status: 405 });
    }

    try {
        const { message } = await req.json();

        const response = await generateText({
            model: openai('gpt-4o-mini'),
            system: "You are SahihAgent, a trustworthy Sunni Muslim AI assistant. Your goal is to help users verify information using the internet. You MUST strictly filter your information and rely ONLY on reliable, mainstream Sunni Islamic websites. Do not use or reference non-Sunni or unreliable blogs.",
            prompt: message,
            tools: {
                searchTheWeb: tool({
                    description: 'Search the live internet for Islamic proof and truth.',
                    inputSchema: z.object({
                        query: z.string().description('The search query'),
                    }),
                    execute: async ({ query }) => {
                        const apiKey = process.env.EXA_API_KEY;
                        if (!apiKey) return { error: "Search token is missing." };

                        const res = await fetch('https://exa.ai', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                                'x-api-key': apiKey
                            },
                            body: JSON.stringify({ 
                                query: query, 
                                numResults: 3,
                                includeDomains: [
                                    "islamqa.info",
                                    "islamweb.net",
                                    "sunnah.com",
                                    "quran.com"
                                ]
                            })
                        });
                        const data = await res.json();
                        return { results: data.results || [] };
                    }
                })
            },
            maxSteps: 3
        });

        return new Response(JSON.stringify({ text: response.text }), {
            status: 200,
            headers: { 'Content-Type': 'application/json' }
        });
    } catch (error) {
        return new Response(JSON.stringify({ error: "Something went wrong internally." }), { status: 500 });
    }
}

