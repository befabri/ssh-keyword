from setuptools import setup

readme = ""
with open("README.md") as f:
    readme = f.read()

setup(
    name="ssh_keyword",
    version='1.0',
    author="befabri",
    url="https://github.com/befabri/ssh-keyword",
    download_url="https://github.com/befabri/ssh-keyword/archive/1.0.tar.gz",
    packages=["ssh_keyword"],
    license="MIT",
    description="A keywords ssh connection",
    long_description=readme,
    long_description_content_type="text/markdown"
)
