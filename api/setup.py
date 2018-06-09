from setuptools import setup

setup(
	name='converter',
	version='1.0',
	py_modules=['cli'],
	install_requires=[
		'Click',
		'requests',
	],
	entry_points='''
		[console_scripts]
		convert=cli:cliConv
	'''
)