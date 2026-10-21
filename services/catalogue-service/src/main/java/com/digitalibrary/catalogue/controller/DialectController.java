package com.digitalibrary.catalogue.controller;

import com.digitalibrary.catalogue.dto.dialect.CreateDialectRequest;
import com.digitalibrary.catalogue.dto.dialect.DialectResponse;
import com.digitalibrary.catalogue.dto.dialect.UpdateDialectRequest;
import com.digitalibrary.catalogue.dto.global.ApiResponse;
import com.digitalibrary.catalogue.service.IDialectService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Controller
@RequestMapping("api/v1/catalog/dialect")
public class DialectController {

    private final IDialectService service;

    public DialectController(IDialectService dialectService) {
        service = dialectService;
    }

    @GetMapping()
    public ResponseEntity<ApiResponse<List<DialectResponse>>> searchAll() {
        List<DialectResponse> response = service.searchAll();
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Successful operation", response));
    }

    @GetMapping("/{code}")
    public ResponseEntity<ApiResponse<DialectResponse>> searchByCode(@PathVariable(value = "code") String code) {
        DialectResponse response = service.searchByCode(code);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Record found", response));
    }

    @PostMapping()
    public ResponseEntity<ApiResponse<DialectResponse>> create(@RequestBody CreateDialectRequest request) {
        DialectResponse response = service.create(request);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Dialect created successfully", response));
    }

    @PutMapping("/{id}")
    public ResponseEntity<ApiResponse<DialectResponse>> update(@PathVariable(value = "id") Long id, @RequestBody UpdateDialectRequest request) {
        DialectResponse response = service.update(id, request);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Dialect information updated successfully", response));
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ApiResponse<?>> delete(@PathVariable(value = "id") Long id) {
        service.delete(id);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Dialect deleted successfully", null));
    }
}
