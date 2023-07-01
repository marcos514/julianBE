import json
import requests
import pandas as pd
import datetime
import jwt

import psycopg2

import logging

users = []

def filter_customers_by_list(customers, list_filter):
    filter_customers = []
    for customer in customers:
        if customer['customer_email'] not in list_filter:
            filter_customers.append(customer)
    logging.info(f'Users to migrate after filter array email: { len(filter_customers) }')
    return filter_customers
    # str_email = "','".join(map(str, emails))

    

def parse_customer(customer):
    customer_id = customer['Recharge Customer ID']
    customer_email = customer['Email']
    new_cadence = customer['Old Order Cadence']
    old_cadence = customer['Recharge Customer ID']
    old_wet, old_fd_12, old_fd_3, old_kib_40, old_kib_4 = customer['Current Order Configuration \n (Wet Food : FD 12oz : FD 3oz : Kib 40oz : Kib 4oz)'].split("-")
    new_wet, new_fd_12, new_fd_3, new_kib_40, new_kib_4 = customer['New Order Configuration \n(Wet Food : FD 12oz : FD 3oz : Kib 40oz : Kib 4oz)'].split("-")

    return {
        'customer_id': customer_id,
        'customer_email': customer_email,
        'old_config': {
            'old_cadence': old_cadence,
            'old_wet': old_wet,
            'old_fd_12': old_fd_12,
            'old_fd_3': old_fd_3,
            'old_kib_40': old_kib_40,
            'old_kib_4': old_kib_4
        },
        'new_config': {
            'new_cadence': int(new_cadence),
            'new_wet': int(new_wet),
            'new_fd_12': int(new_fd_12),
            'new_fd_3': int(new_fd_3),
            'new_kib_40': int(new_kib_40),
            'new_kib_4': int(new_kib_4)
        }
    }

def read_migration_users():
    df = pd.read_csv('migration_users.csv')
    customers = []
    for customer in df.to_dict(orient='records'):
        customers.append(parse_customer(customer))
    logging.info(f'There are {len(customers)} in migration file')
    return customers


def recharge_request(endpoint,method,url_params={},body={}):
    url = f'https://api.rechargeapps.com/{endpoint}'
    headers= {'X-Recharge-Access-Token': 'TOKEN'}

    if method == 'GET':
        r = requests.get(url = url,params = url_params,headers=headers)
    elif method == 'PUT':
        headers['Accept'] = "application/json"
        headers['Content-Type'] = "application/json"
        r = requests.put(url = url, data=json.dumps(body), headers=headers)

    data = r.json()
    return data
        
def get_customer(customer):
    endpoint = 'customers'
    method = 'GET'
    url_params_id = {'customer_id':customer['id']}
    url_params_email = {'email':customer['email']}
    logging.info(recharge_request(endpoint,method,url_params=url_params_email))

def get_subscriptions(customer):
    endpoint = 'subscriptions'
    method = 'GET'
    url_params_id = {'customer_id': customer['customer_id'], 'status': 'ACTIVE'}
    return recharge_request(endpoint,method,url_params=url_params_id)['subscriptions']

def update_subscription(subscriptions, customer):
    method = 'PUT'
    for subscription in subscriptions:
        endpoint = f'subscriptions/{ subscription["id"] }'
        body = {
            'order_day_of_week': 4
        }
        if subscription['order_day_of_week'] != 4:
            logging.info(f'Start update subscription: { subscription["shopify_variant_id"] }')
            recharge_request(endpoint, method, body=body)
            logging.info(f"From: { subscription['order_day_of_week'] }")

def charge_schedule_at(customer):
    method = 'GET'
    endpoint = f'charges/'
    url_params = {
        'status': 'queued',
        'customer_id': customer['customer_id']
    }
    ret = recharge_request(endpoint, method, url_params=url_params)['charges']
    if len(ret)>0:
        logging.info(f"Charge {ret[0]['id']} scheduled at {ret[0]['scheduled_at']}")


if __name__ == "__main__":
    logging.basicConfig(filename='check_migration_5.log',level=logging.DEBUG)
    try:
        customers = read_migration_users()
        customers = [customer for customer in customers if customer['customer_email'] in users ]
        migration_round = 1
        for customer in customers:
            logging.info(f'\n\n\n\n\nMigration Round { migration_round } for { customer["customer_email"] }')
            subscriptions = get_subscriptions(customer)
            # logging.error(customer)
            update_subscription(subscriptions, customer)
            # charge_schedule_at(customer)
            migration_round += 1

    except (Exception) as error :
        logging.error(f"Error exception type ({ type(error) })\nMessage: { error }")
