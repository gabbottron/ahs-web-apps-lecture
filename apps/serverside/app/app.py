from flask import Flask, render_template, request, redirect, url_for
import psycopg2

app = Flask(__name__)

# Configure the PostgreSQL database connection
db_connection = psycopg2.connect(
    database="webapp",
    user="postgres",
    password="password",
    host="postgres",
    port="5432"
)

# Function to fetch items from the PostgreSQL database
def fetch_items():
    cursor = db_connection.cursor()
    cursor.execute("SELECT id, name, quantity FROM items")
    items = cursor.fetchall()
    cursor.close()
    return items

# Function to fetch item details by ID
def fetch_item_details(item_id):
    cursor = db_connection.cursor()
    cursor.execute("SELECT name, quantity, description FROM items WHERE id = %s", (item_id,))
    item = cursor.fetchone()
    cursor.close()
    return item

@app.route("/")
def index():
    items = fetch_items()
    return render_template("index.html", items=items)

@app.route("/items/<int:item_id>")
def item_details(item_id):
    item = fetch_item_details(item_id)
    return render_template("item_details.html", item=item)

if __name__ == "__main__":
    app.run(debug=True)

