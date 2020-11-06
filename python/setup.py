from distutils.core import setup

from setuptools import find_packages

setup(
    name='agora-api',
    version='0.23.0',
    description='Agora API for Python',
    author='Kik Engineering',
    author_email='engineering@kik.com',
    url='https://github.com/kinecosystem/agora-api',
    license='MIT',
    packages=find_packages(include=["agoraapi", "agoraapi.*"]),
    long_description=open("../README.md").read(),
    long_description_content_type="text/markdown",
    classifiers=[
        'License :: OSI Approved :: MIT License',
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Libraries :: Python Modules',
        'Programming Language :: Python',
        'Programming Language :: Python :: 3',
    ],
    install_requires=[
        "protobuf"
    ],
)
