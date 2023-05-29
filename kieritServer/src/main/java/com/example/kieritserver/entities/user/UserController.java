package com.example.kieritserver.entities.user;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

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
    @PostMapping("/signup")
    public ResponseEntity<CreateUserRes> Register(@RequestBody CreateUserReq request) {

        return ResponseEntity.status(HttpStatus.OK).body(
                new CreateUserRes("1", request.username, request.email)
        );
    }

    @Override
    @PostMapping("/login")
    public ResponseEntity<LoginUserRes> Login(@RequestBody LoginUserReq request) {
        return null;
    }
}
