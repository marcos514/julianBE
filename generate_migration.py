import json
import requests
import pandas as pd
import datetime
import jwt

import psycopg2

import logging

no_migrate_users = []

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

def log_error(customers_updated,customer_id,subscriptions_updated,subscription_error,error):
    error_time = datetime.datetime.strftime('%Y-%m-%d%H:%M:%S')
    f = open(f'{error_time}.json','w')
    json_error = {
        'customers_updated':customers_updated,
        'customer_id': customer_id,
        'subscriptions_updated': subscriptions_updated,
        'subscription_error': subscription_error,
        'error': error
    }
    with open(f'{error_time}.json','w') as f:
        json.dump(json_error,f)

def recharge_request(endpoint,method,url_params={},body={}):
    url = f'https://api.rechargeapps.com/{endpoint}'
    headers= {'X-Recharge-Access-Token': 'TOKEN'}

    if method == 'GET':
        r = requests.get(url = url,params = url_params,headers=headers)
    elif method == 'POST':
        r = requests.post(url = url,data = body,headers=headers)
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

def separate_subscriptions(subscriptions):
    logging.info('Separating subscriptions')
    wet = []
    fd_12 = []
    fd_3 = []
    kib_40 = []
    kib_4 = []
    others = []

    wet_list = [31605186691171, 31605186789475, 31605186723939,31605186855011, 31605186953315, 31605186986083,31605187018851, 32851664863331, 31956521255011,31956508934243, 31956484882531, ]
    fd_12_list = [31605186658403, 31605186560099, 30875534753891, 31846613385315]
    fd_3_list = [31605186625635, 31605186527331, 30875534721123, 31926051012707, 31926053240931, 31926049865827, 31846613352547]
    kib_40_list = [31605185183843, 31605185347683, 31605185478755]
    kib_4_list = [32123774468195, 32123777777763, 32123776401507]
    for sub in subscriptions:
        if sub['shopify_variant_id'] in wet_list:
            wet.append({
                "id": sub['id'],
                "quantity": int(sub['quantity']),
                "shopify_variant_id": sub['shopify_variant_id'],
                "variant_title": sub['variant_title']
            })
        elif sub['shopify_variant_id'] in fd_12_list:
            fd_12.append({
                "id": sub['id'],
                "quantity": int(sub['quantity']),
                "shopify_variant_id": sub['shopify_variant_id'],
                "variant_title": sub['variant_title']
            })
        elif sub['shopify_variant_id'] in fd_3_list:
            fd_3.append({
                "id": sub['id'],
                "quantity": int(sub['quantity']),
                "shopify_variant_id": sub['shopify_variant_id'],
                "variant_title": sub['variant_title']
            })
        elif sub['shopify_variant_id'] in kib_40_list:
            kib_40.append({
                "id": sub['id'],
                "quantity": int(sub['quantity']),
                "shopify_variant_id": sub['shopify_variant_id'],
                "variant_title": sub['variant_title']
            })
        elif sub['shopify_variant_id'] in kib_4_list:
            kib_4.append({
                "id": sub['id'],
                "quantity": int(sub['quantity']),
                "shopify_variant_id": sub['shopify_variant_id'],
                "variant_title": sub['variant_title']
            })
        else:
            others.append({
                "id": sub['id'],
                "quantity": int(sub['quantity']),
                "shopify_variant_id": sub['shopify_variant_id'],
                "variant_title": sub['variant_title']
            })

    return {
        'wet' :wet,
        'fd_12': fd_12,
        'fd_3': fd_3,
        'kib_40': kib_40,
        'kib_4': kib_4,
        'others': others
    }

def subscription_arithmetic(subscriptions, new_max):
    quantity = 0
    for subscription in subscriptions:
        quantity += subscription['quantity']
    return 0 if new_max < quantity else new_max - quantity


def update_subscription(subscriptions, customer, sub_type=False):
    method = 'PUT'
    more_quantity = 0
    if sub_type != False:
        more_quantity = subscription_arithmetic(subscriptions, customer['new_config'][sub_type])
        logging.info(f'Customer subscription type { sub_type } need { more_quantity } ')
    else:
        logging.info(f'Customer subscription type other needs 0')

    for subscription in subscriptions:
        logging.info(f'Start update subscription: { subscription["shopify_variant_id"] }')
        endpoint = f'subscriptions/{ subscription["id"] }'
        body = {
            'order_interval_unit': 'week',
            'order_interval_frequency': customer['new_config']['new_cadence'],
            'charge_interval_frequency': customer['new_config']['new_cadence'],
        }
        if more_quantity > 0:
            logging.info(f'Subscription { subscription["shopify_variant_id"] } needs to change from {subscription["quantity"]} to { subscription["quantity"] + more_quantity } quantity')
            body['quantity'] = subscription['quantity'] + more_quantity
            more_quantity = 0

        recharge_request(endpoint, method, body=body)
        logging.info(f"Subscription { subscription['id'] } updated")
        logging.info(f"From: { subscription }")
        logging.info(f"To: { body }")


def update_customer_subscription(subscriptions, customer):
    sub_type = 'wet'
    new_sub_type = f'new_wet'
    update_subscription(subscriptions[sub_type], customer, new_sub_type)

    sub_type = 'fd_12'
    new_sub_type = f'new_fd_12'
    update_subscription(subscriptions[sub_type], customer, new_sub_type)

    sub_type = 'fd_3'
    new_sub_type = f'new_fd_3'
    update_subscription(subscriptions[sub_type], customer, new_sub_type)

    sub_type = 'kib_40'
    new_sub_type = f'new_kib_40'
    update_subscription(subscriptions[sub_type], customer, new_sub_type)

    sub_type = 'kib_4'
    new_sub_type = f'new_kib_4'
    update_subscription(subscriptions[sub_type], customer, new_sub_type)

    sub_type = 'others'
    update_subscription(subscriptions[sub_type], customer)



if __name__ == "__main__":
    logging.basicConfig(filename='check_migration.log',level=logging.DEBUG)
    try:
        customers = read_migration_users()
        connection = psycopg2.connect(user="",
                                    password="",
                                    host="",
                                    port="",
                                    database="")
        select_cursor = connection.cursor()
        customers = filter_customers_by_list(customers, no_migrate_users)
        str_email = "','".join(map(str, [customer['customer_email'] for customer in customers]))
        select_query = f"select c.email from customers c LEFT JOIN box_migration b ON b.customer_id = c.id where c.email IN ('{str_email}') and b.update_box = true;"

        select_cursor.execute(select_query)
        logging.info("Selecting customers")
        db_customers = [customer[0] for customer in select_cursor.fetchall()]
        select_cursor.close()
        customers = [customer for customer in customers if customer['customer_email'] in db_customers]
        logging.info(f'We have {len(customers)} to migrate\n\n\n')
        migration_round = 1
        for customer in customers:
            logging.info(f'\n\n\n\n\nMigration Round { migration_round }')
            logging.info(f'Start customer migration for: { customer["customer_email"] }')
            subscriptions = get_subscriptions(customer)
            subscriptions = separate_subscriptions(subscriptions)
            update_customer_subscription(subscriptions, customer)
            logging.info(f'Customer { customer["customer_email"] } migration finnished')
            migration_round += 1

    except (Exception) as error :
        logging.error(f"Error exception type ({ type(error) })\nMessage: { error }")

    finally:
        #closing database connection.
        if(connection):
            select_cursor.close()
            connection.close()
            logging.info("PostgreSQL connection is closed")
