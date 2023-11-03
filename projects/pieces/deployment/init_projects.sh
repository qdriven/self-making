#  https://github.com/DavidVujic/python-polylith.git
# https://davidvujic.github.io/python-polylith-docs/commands/
poetry init
poetry add poetry-polylith-plugin
poetry add poetry-multiproject-plugin

## create workspace
poetry poly create workspace --name spells-hub --theme loose
# poetry poly create workspace --name spellshub --theme tdd
poetry install
# create component
poetry poly create component --name base
# # create base
poetry poly create base --name my_example_base
# create project
poetry poly create project --name pdf-spell
# poly diff
poetry poly diff




poetry poly info
poetry poly libs
poetry poly check
poetry poly sync

# packaging
poetry build-project --directory path/to/project
poetry build-project --with-top-namespace my_custom_namespace

# deploying
pip install the-built-artifact.whl
