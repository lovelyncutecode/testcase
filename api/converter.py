import requests

def convert(amount,currency):
	currency = currency.upper()
	if currency not in ['UAH','EUR','GBP']: 
		raise ValueError
	
	try:
		amount = float(amount)
	except:
		raise ValueError

	if amount < 0:
		raise ValueError
	URL = "https://free.currencyconverterapi.com/api/v5/convert"
	PARAMS = {'q':'USD_'+currency,'compact':'ultra'}

	try:
		r = requests.get(url = URL, params = PARAMS)
	except Exception as e:
		raise e
	else:
		data = r.json()
		rate = data['USD_'+currency]
		result = amount*rate 
		return result
