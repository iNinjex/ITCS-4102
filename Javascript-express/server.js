const express = require("express");

const app = express();
const PORT = 3000;

app.use(express.json());

let languages = [
    {
        id: 1,
        name: "Python",
        paradigm: "Multi-paradigm",
        typing: "Dynamic",
        year_created: 1991
    },
    {
        id: 2,
        name: "JavaScript",
        paradigm: "Event-driven",
        typing: "Dynamic",
        year_created: 1995
    },
    {
        id: 3,
        name: "Go",
        paradigm: "Concurrent",
        typing: "Static",
        year_created: 2009
    }
];

// Home
app.get("/", (req, res) => {
    res.json({
        message: "Programming Languages API is working!"
    });
});

// GET all languages
app.get("/languages", (req, res) => {
    res.json(languages);
});

// GET one language
app.get("/languages/:id", (req, res) => {

    const id = parseInt(req.params.id);

    const language = languages.find(lang => lang.id === id);

    if (!language) {
        return res.status(404).json({
            message: "Language not found"
        });
    }

    res.json(language);
});

// POST
app.post("/languages", (req, res) => {

    const language = req.body;

    if (languages.find(lang => lang.id === language.id)) {
        return res.status(400).json({
            message: "Language ID already exists"
        });
    }

    languages.push(language);

    res.status(201).json(language);

});

// PUT
app.put("/languages/:id", (req, res) => {

    const id = parseInt(req.params.id);

    const index = languages.findIndex(lang => lang.id === id);

    if (index === -1) {
        return res.status(404).json({
            message: "Language not found"
        });
    }

    languages[index] = req.body;

    res.json(languages[index]);

});

// DELETE
app.delete("/languages/:id", (req, res) => {

    const id = parseInt(req.params.id);

    const index = languages.findIndex(lang => lang.id === id);

    if (index === -1) {
        return res.status(404).json({
            message: "Language not found"
        });
    }

    const deleted = languages.splice(index, 1);

    res.json({
        message: "Language deleted",
        deleted: deleted[0]
    });

});

app.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}`);
});