from distutils.core import setup

from setuptools import find_packages

setup(
    name='agoraapi',
    version='0.17.0',
    packages=find_packages(include=["agoraapi", "agoraapi.*"]),
    author='kin',
    install_requires=[
        "protobuf"
    ]
)
