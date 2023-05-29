package com.example.kieritserver.entities.user;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestBody;

interface IUserRepository {
    User createUser(User user) throws Exception;
    User getUserByEmail(String email) throws Exception;
}

interface IUserService {
    LoginUserRes SignIn(LoginUserReq req) throws Exception;
    CreateUserRes SingUp(CreateUserReq req) throws Exception;
}

interface IUserController{
    ResponseEntity<CreateUserRes> Register(@RequestBody CreateUserReq request);
    ResponseEntity<LoginUserRes> Login(@RequestBody LoginUserReq request);
}
