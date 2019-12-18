db.createUser(
    {
        user: "user1",
        pwd: "password1",
        roles:[
            {
                role: "readWrite",
                db:   "testing"
            }
        ]
    }
);