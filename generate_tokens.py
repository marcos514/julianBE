import jwt

import pandas as pd


import psycopg2


import base64
import jwt
from typing import Dict

emails = ['justin@smallsforsmalls.com']

def toBase64(payload_to_encode: Dict) -> str:
    return base64.b64encode(payload_to_encode)

def fromBase64(encoded_string: str) -> Dict:
    return base64.decodebytes(encoded_string)

def generate_token(email, shopify_id, recharge_id):
    seed = {
        "email": email,
        "shopify_id": shopify_id,
        "recharge_id": recharge_id,
    }
    PRIVATE_KEY = ''
    encodedJson = jwt.encode({"data": seed}, PRIVATE_KEY, algorithm='HS256')
    b64EncodedJson = toBase64(encodedJson)
    return b64EncodedJson

if __name__ == "__main__":
    try:
        connection = psycopg2.connect(user="",
                                    password="",
                                    host="",
                                    port="",
                                    database="")
        select_cursor = connection.cursor()
        str_email = "','".join(map(str, emails))
        select_query = f"select id, email, shopify_id, recharge_id from customers where email IN ('{str_email}') ORDER BY email"

        select_cursor.execute(select_query)
        print("Selecting customers")
        customers = select_cursor.fetchall()
        select_cursor.close()
        insert_token_sql = """INSERT INTO box_migration(customer_id, token)
             VALUES(%s, %s)"""
        print("Start for")
        insert_cursor = connection.cursor()
        i = 0
        for customer in customers:
            i += 1
            customer_id = customer[0]
            email = customer[1]
            shopify_id = customer[2]
            recharge_id = customer[3]
            print('Start Round ',i,' working with: ', email)
            token = generate_token(email, shopify_id, recharge_id)
            print('Token: ', token)
            insert_cursor.execute(insert_token_sql, (customer_id ,token.decode("utf-8")))
            # get the generated id back
            # commit the changes to the database
            connection.commit()
            



        #    generate_token
        # insert in table

    except (Exception, psycopg2.Error) as error :
        print ("Error while fetching data from PostgreSQL", error)

    finally:
        #closing database connection.
        if(connection):
            select_cursor.close()
            insert_cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")

