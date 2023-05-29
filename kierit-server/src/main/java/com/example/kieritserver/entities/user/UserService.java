package com.example.kieritserver.entities.user;

import org.springframework.beans.factory.annotation.Autowired;

public class UserService implements IUserService{
    private final UserRepository userRepo;

    public UserService() {
        this.userRepo = new UserRepository();
    }

    @Override
    public LoginUserRes SignIn(LoginUserReq req) throws Exception {
        return null;
    }

    @Override
    public CreateUserRes SingUp(CreateUserReq req) throws Exception {
        return null;
    }
}
