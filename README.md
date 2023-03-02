# Application API REST 
## fait en Golang, avec MongoDB pour base de données

Cet API supporte ces 5 fonctions :
- Create POST "/add/users" avec pour body un fichier JSON
- Login POST "/login/:id/:pwd"
- Delete DELETE "user/:id"
- Read GET "/users/list" et "/user/:id"
- Update PUT "/user/:id"

# Structure JSON du fichier que prend à paramètre la requete create

[
    {
        "id": "abcdebffeeeejekfeekkkkkkkkkkkkkkkkkkfoezjdddddddddddddddddddddddddddddddddddddddzfjzfzofjzfGZJ",
    "password":"est crypte un peu plus tard",
    "isActive": false,
    "balance": "$3,896.60",
    "age": 35,
    "name": "sujet test",
    "gender": "male",
    "company": "a name",
    "email": "testsubj@qmet.com",
    "phone": "+1 (992) 455-45852",
    "address": "567 Box Street, Courtland, Alabama, 7462",
    "about": "Consectetur cupidatat Lorem duis pariatur reprehenderit anim dolor ea sunt sunt nostrud. Mollit aute aliqua do velit exercitation ex mollit eiusmod et dolor tempor. Veniam adipisicing enim consectetur labore qui ipsum Lorem commodo mollit magna nostrud et.\r\n",
    "registered": "2016-10-10T10:11:34 -02:00",
    "latitude": 58.262631,
    "longitude": 72.846112,
    "tags": [
      "incididunt",
      "exercitation",
      "eu",
      "nulla",
      "ut",
      "sit",
      "amet"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Bates Fletcher"
      },
      {
        "id": 1,
        "name": "Sanders Hampton"
      },
      {
        "id": 2,
        "name": "Eliza Francis"
      }
    ],
    "data":"contenu contenant contenu contenantcontenu contenantcontenu contenant etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc...etc..."
       },
    ...
]
# Structure JSON du fichier que prend à paramètre la requete update
## L'id n'est pas modifiable
{
    "attribut": "valeur",
    ...
}
<!--go get go.mongodb.org/mongo-driver
go get -u github.com/gin-gonic/gin

go run main.go

TO DO
- refaire la classe user model (ok)
- validation des champs
- mot de passe a update (ok)
    dehashing
- Changer le parametre user en user [] (ok)
- controle de user list (ok)
- Recuperer la liste (ok)
- creer des fichiers (ok)
- update:
    if age not null (ok...?) -->