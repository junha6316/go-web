import requests


urls = [
    ('GET', 'http://localhost:8080/users/1',),
    ('GET', 'http://localhost:8080/users/1/addresses/2',),
    ('POST', 'http://localhost:8080/users',),
    ('POST', 'http://localhost:8080/users/1/addresses',),
]

answer = [
    "retrieve user",
    "retrieve user's address",
    "create user",
    "create user's address",
]

for url_info, a in zip(urls, answer):
    method, url = url_info
    if method == "GET":
        res = requests.get(url)
    if method == "POST":
        res = requests.post(url)
    print(res.text)