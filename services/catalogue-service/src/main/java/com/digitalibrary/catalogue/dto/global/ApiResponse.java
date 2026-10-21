package com.digitalibrary.catalogue.dto.global;

public record ApiResponse<T>(
        boolean success,
        String message,
        T data
) {}
