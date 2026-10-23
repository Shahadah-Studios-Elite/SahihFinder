class SahihFinderApp {
    private appRoot: HTMLElement | null;

    constructor() {
        this.appRoot = document.getElementById("app");
    }

    public render(): void {
        if (!this.appRoot) return;

        this.appRoot.innerHTML = `
            <header>
                <h1>SahihFinder</h1>
                <p>A website for Muslims to verify the truth from the internet</p>
            </header>

            <main>
                <h2>Welcome to SahihFinder</h2>
                <p>A dedicated platform built to help clarify and verify information online in accordance with Islamic values.</p>
                
                <a href="Wiki.html" class="wiki-button">Explore The WIKI</a>
            </main>

            <footer>
                <p>&copy; 2026 SahihFinder - Shahadah Studios Elite</p>
            </footer>
        `;
    }
}

document.addEventListener("DOMContentLoaded", () => {
    const app = new SahihFinderApp();
    app.render();
});

