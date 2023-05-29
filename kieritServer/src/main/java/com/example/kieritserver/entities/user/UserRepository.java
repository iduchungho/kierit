package com.example.kieritserver.entities.user;

import org.springframework.data.mongodb.repository.MongoRepository;

public class UserRepository implements IUserRepository{
    @Override
    public User createUser(User user) throws Exception {
        return null;
    }

    @Override
    public User getUserByEmail(String email) throws Exception {
        return null;
    }
}
