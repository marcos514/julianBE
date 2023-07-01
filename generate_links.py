import jwt

import pandas as pd


import psycopg2


import base64
import jwt
from typing import Dict

emails=[
    'justin@smallsforsmalls.com'
]

if __name__ == "__main__":
    try:
        connection = psycopg2.connect(user="",
                                    password="",
                                    host="",
                                    port="",
                                    database="")
        select_cursor = connection.cursor()
        select_query = "Select c.email, b.token from box_migration b left join customers c ON c.id = b.customer_id where c.email in ('justin@smallsforsmalls.com');"

        select_cursor.execute(select_query)
        customers = select_cursor.fetchall()
        select_cursor.close()
        for customer in customers:
            email = customer[0]
            token = customer[1]
            print(f'https://smalls.com/pages/migration?email={email}&token={token}')
            



    except (Exception, psycopg2.Error) as error :
        print ("Error while fetching data from PostgreSQL", error)

    finally:
        #closing database connection.
        if(connection):
            select_cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")

