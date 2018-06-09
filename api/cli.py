import click 
from click.testing import CliRunner
from converter import convert

@click.command()
@click.argument('amount')
@click.argument('currency')		
def cliConv(amount,currency):
	"""This script converts inputed amount of USD to the amount of chosen currency."""
	try:
		res=convert(amount,currency)
		print(res)
	except ValueError:
		click.echo('Supported currencies: UAH, EUR, GBP.\nAmount should be a number greater than 0.')	
	except:
	 	click.echo('Oops. Try again later...')


if __name__ == '__main__':
    testConverter()