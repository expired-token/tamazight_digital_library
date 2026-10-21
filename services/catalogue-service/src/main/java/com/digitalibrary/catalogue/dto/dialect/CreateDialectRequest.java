package com.digitalibrary.catalogue.dto.dialect;

import jakarta.validation.constraints.NotBlank;

public record CreateDialectRequest(
        @NotBlank(message = "Code is required")
        String code,

        @NotBlank(message = "Value is required")
        String value,

        String description
) {}
