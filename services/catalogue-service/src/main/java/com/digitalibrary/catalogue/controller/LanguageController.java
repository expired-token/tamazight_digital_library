package com.digitalibrary.catalogue.controller;

import com.digitalibrary.catalogue.dto.language.CreateLanguageRequest;
import com.digitalibrary.catalogue.dto.language.LanguageResponse;
import com.digitalibrary.catalogue.dto.global.ApiResponse;
import com.digitalibrary.catalogue.dto.language.UpdateLanguageRequest;
import com.digitalibrary.catalogue.service.ILanguageService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Controller
@RequestMapping("api/v1/catalog/language")
public class LanguageController {

    private final ILanguageService service;

    public LanguageController(ILanguageService languageService) {
        service = languageService;
    }

    @GetMapping()
    public ResponseEntity<ApiResponse<List<LanguageResponse>>> searchAll() {
        List<LanguageResponse> response = service.searchAll();
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Successful operation", response));
    }

    @GetMapping("/{code}")
    public ResponseEntity<ApiResponse<LanguageResponse>> searchByCode(@PathVariable(value = "code") String code) {
        LanguageResponse response = service.searchByCode(code);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Record found", response));
    }

    @PostMapping()
    public ResponseEntity<ApiResponse<LanguageResponse>> create(@RequestBody CreateLanguageRequest request) {
        LanguageResponse response = service.create(request);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Language created successfully", response));
    }

    @PutMapping("/{id}")
    public ResponseEntity<ApiResponse<LanguageResponse>> update(@PathVariable(value = "id") Long id, @RequestBody UpdateLanguageRequest request) {
        LanguageResponse response = service.update(id, request);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Language information updated successfully", response));
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ApiResponse<?>> delete(@PathVariable(value = "id") Long id) {
        service.delete(id);
        return ResponseEntity.status(HttpStatus.OK).body(new ApiResponse<>(true, "Language deleted successfully", null));
    }
}
