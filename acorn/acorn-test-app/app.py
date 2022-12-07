import logging as log
import os

import psycopg2
import redis
from flask import Flask, render_template_string

# HTML Jinja2 Template which will be shown in the browser
page_template = '''
        <div style="margin: auto; text-align: center;">
        <h1>{{ welcome_text }}</h1><br>
        You're visitor #{{ visitors }} to learn what squirrels love the most:<br>
        <ul>
            {%- for food in foods %}
            <li>{{ food }}</li>
            {%- endfor %}
        </ul>
        </div>
        '''

# Defining the Flask Web App
app = Flask(__name__)
cache = redis.StrictRedis(host='cache', port=6379)


# The website root will show the page_template rendered with
# - visitor count fetched from Redis Cache
# - list of food fetched from Postgres DB
# - welcome text passed in as environment variable
@app.route('/')
def root():
    visitors = cache_get_visitor_count()
    food = db_get_squirrel_food()

    return render_template_string(page_template, visitors=visitors, foods=food, welcome_text=os.getenv("WELCOME", "Hey Acorn user!"))


# Fetch the squirrel food from the Postgres database
def db_get_squirrel_food():
    conn = psycopg2.connect(
        host="db",
        database="acorn",
        user=os.environ['PG_USER'],
        password=os.environ['PG_PASS'],
    )

    cur = conn.cursor()
    cur.execute("SELECT food FROM squirrel_food;")

    return [x[0] for x in cur.fetchall()]  # Return the list of food items


# Increment the visitor count in the Redis cache and return the new value
def cache_get_visitor_count():
    return cache.incr('visitors')
    