def user [lastName firstName] {
    { lastName: '$lastName', firstName: '$firstName' }
}

http post http://localhost:8080/echo --content-type application/json user