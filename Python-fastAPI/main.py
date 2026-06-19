from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI(title="Programming Languages API")


class Language(BaseModel):
    id: int
    name: str
    paradigm: str
    typing: str
    year_created: int


languages = [
    Language(id=1, name="Python", paradigm="Multi-paradigm", typing="Dynamic", year_created=1991),
    Language(id=2, name="JavaScript", paradigm="Event-driven", typing="Dynamic", year_created=1995),
    Language(id=3, name="Go", paradigm="Concurrent", typing="Static", year_created=2009),
]


@app.get("/")
def home():
    return {"message": "Programming Languages API is working"}


@app.get("/languages")
def get_languages():
    return languages


@app.get("/languages/{language_id}")
def get_language(language_id: int):
    for language in languages:
        if language.id == language_id:
            return language

    raise HTTPException(status_code=404, detail="Language not found")


@app.post("/languages", status_code=201)
def create_language(language: Language):
    for existing_language in languages:
        if existing_language.id == language.id:
            raise HTTPException(status_code=400, detail="Language ID already exists")

    languages.append(language)
    return language


@app.put("/languages/{language_id}")
def update_language(language_id: int, updated_language: Language):
    for index, language in enumerate(languages):
        if language.id == language_id:
            languages[index] = updated_language
            return updated_language

    raise HTTPException(status_code=404, detail="Language not found")


@app.delete("/languages/{language_id}")
def delete_language(language_id: int):
    for index, language in enumerate(languages):
        if language.id == language_id:
            deleted_language = languages.pop(index)
            return {
                "message": "Language deleted",
                "deleted": deleted_language
            }

    raise HTTPException(status_code=404, detail="Language not found")