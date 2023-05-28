package com.example.kieritserver.entities.user;

import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestBody;

@Data
@Document
public class User {
    @Id
    private int id;
    private String username;
    @Indexed(unique = true)
    private String email;
    private String password;

    public User(int id, String username, String email, String password) {
        this.id = id;
        this.username = username;
        this.email = email;
        this.password = password;
    }
}

class LoginUserReq {
    public String email;
    public String password;
}

class LoginUserRes {
    public String accessToken;
    public String id;
    public String username;
}

class CreateUserReq {
    public String username;
    public String email;
    public String password;
}

class CreateUserRes {
    public String id;
    public String username;
    public String email;
}

interface IUserRepository {
    User createUser(User user);
    User getUserByEmail(String email);
}

interface IUserService {
    LoginUserRes SignIn(LoginUserReq req);
    CreateUserRes SingUp(CreateUserReq req);
}

interface IUserController{
    ResponseEntity<User> Register(@RequestBody CreateUserReq request);
    ResponseEntity<User> Login(@RequestBody LoginUserReq request);
}