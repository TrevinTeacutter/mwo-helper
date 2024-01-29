# -*- coding: utf-8 -*-

# Learn more: https://github.com/kennethreitz/setup.py

from setuptools import setup, find_packages


with open('README.md') as f:
    readme = f.read()

with open('LICENSE.md') as f:
    license = f.read()

setup(
    name='mwo-helper',
    version='0.1.0',
    description='Meant to be a one-stop shop for MWO comp related activities, without the need for a website.',
    long_description=readme,
    author='Trevin Teacutter',
    author_email='tjteacutter1@cougars.ccis.edu',
    url='https://github.com/trevinteacutter/mwo-helper',
    license=license,
    packages=find_packages(exclude=('tests', 'docs'))
)