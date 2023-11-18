from enum import Enum

from pydantic import BaseModel


class Language(str, Enum):
    JAVA = "Java"
    PYTHON = "Python"
    GO = "Go"
    TYPESCRIPT = "Typescript"
    SQL = "SQL"


class EntityInputModel(BaseModel):
    language: str = "Python"
    lib_name: str = "pydantic"
    references: str
