package com.example.kieritserver;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class KieritServerApplication {
	public static void main(String[] args) {
		SpringApplication.run(KieritServerApplication.class, args);
	}

}
