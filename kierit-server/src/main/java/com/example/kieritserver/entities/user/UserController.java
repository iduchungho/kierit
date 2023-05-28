package com.example.kieritserver.entities.user;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/user")
public class UserController implements IUserController{
    @GetMapping("/get")
    public ResponseEntity<List<User>> getUser() {
        return new ResponseEntity<List<User>>(
                List.of(new User(12, "duchung", "abc", "123")),
                HttpStatus.UNAUTHORIZED
        );
    }

    @Override
    public ResponseEntity<User> Register(@RequestBody CreateUserReq request) {
        return null;
    }

    @Override
    public ResponseEntity<User> Login(@RequestBody LoginUserReq request) {
        return null;
    }
}
