import shutil
from pathlib import Path

try:
    json2go = [parent for parent in Path.cwd().parents if parent.name == 'json2go'][0]
except IndexError:
    print('Please run this script from the json2go directory')
    exit(1)

src = json2go / "web/node_modules/monaco-editor/min/vs"
dst = json2go / "web/public/monaco-editor/min/vs"

dst.mkdir(parents=True, exist_ok=True)

try:
    shutil.copytree(src / "base", dst / "base")
except FileExistsError:
    pass

try:
    shutil.copytree(src / "basic-languages" / "go", dst / "basic-languages" / "go")
except FileExistsError:
    pass

try:
    shutil.copytree(src / "editor", dst / "editor")
except FileExistsError:
    pass

try:
    shutil.copytree(src / "language" / "json", dst / "language" / "json")
except FileExistsError:
    pass

try:
    shutil.copy(src / "loader.js", dst / "loader.js")
except FileExistsError:
    pass
