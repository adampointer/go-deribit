import yaml
import os
from yamlinclude import YamlIncludeConstructor

dirname = os.path.dirname(__file__)
basedir = os.path.join(dirname, '../schema/')
root = os.path.join(basedir, 'deribit_api_v2.yaml')

YamlIncludeConstructor.add_to_loader_class(loader_class=yaml.Loader, base_dir=basedir)

with open(root) as f:
    data = yaml.load(f, Loader=yaml.Loader)

print yaml.dump(data)
